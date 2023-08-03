package nodeutil

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"

	"bytes"

	"github.com/freeconf/yang/meta"
)

const QUOTE = '"'

type JSONWtr struct {

	// stream to write contents.  contents will be flushed only at end of operation
	Out io.Writer

	// adds extra indenting and line feeds
	Pretty bool

	// otherwise enumerations are written as their labels.  it may be
	// useful to know that json reader can accept labels or values
	EnumAsIds bool

	// Namespaces pollute JSON with module name similar to XML namespaces
	// rules
	//    { "ns:key" : {...}}
	// where you add the module name to top-level object then qualify any
	// sub objects when the module changes. Not only does it make JSON even more
	// illegible, it means you cannot move common meta to a common yang module w/o
	// altering your resulting JSON.  #IETF-FAIL #rant
	//
	// See https://datatracker.ietf.org/doc/html/rfc7951
	//
	// I realize this is to protect against 2 or more keys in same line from different
	// modules but maybe if someone is insane enough to do that, then, and only then, do
	// you distinguish each key with ns
	//
	// To disable this, make this true and get simple JSON like this
	//
	//    { "key": {...}}
	QualifyNamespace bool

	_out *bufio.Writer
}

func WriteJSON(s node.Selection) (string, error) {
	buff := new(bytes.Buffer)
	wtr := &JSONWtr{Out: buff}
	err := s.InsertInto(wtr.Node()).LastErr
	return buff.String(), err
}

func WritePrettyJSON(s node.Selection) (string, error) {
	buff := new(bytes.Buffer)
	wtr := &JSONWtr{Out: buff, Pretty: true}
	err := s.InsertInto(wtr.Node()).LastErr
	return buff.String(), err
}

func (wtr JSONWtr) JSON(s node.Selection) (string, error) {
	buff := new(bytes.Buffer)
	wtr.Out = buff
	err := s.InsertInto(wtr.Node()).LastErr
	return buff.String(), err
}

func NewJSONWtr(out io.Writer) *JSONWtr {
	return &JSONWtr{Out: out}
}

func (wtr *JSONWtr) Node() node.Node {
	// JSON can begin at a container, inside a list or inside a container, each of these has
	// different results to make json legal
	wtr._out = bufio.NewWriter(wtr.Out)
	return &Extend{
		Base: wtr.container(0),
		OnBeginEdit: func(p node.Node, r node.NodeRequest) error {
			if err := wtr.beginObject(); err != nil {
				return err
			}
			if meta.IsList(r.Selection.Meta()) && !r.Selection.InsideList {
				ident := wtr.ident(r.Selection.Path)
				if err := wtr.beginList(ident); err != nil {
					return err
				}
			}
			return nil
		},
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			if meta.IsList(r.Selection.Meta()) && !r.Selection.InsideList {
				if err := wtr.endList(); err != nil {
					return err
				}
			}
			if err := wtr.endContainer(); err != nil {
				return err
			}
			if err := wtr._out.Flush(); err != nil {
				return err
			}
			return nil
		},
	}
}

func (wtr *JSONWtr) container(lvl int) node.Node {
	first := true
	delim := func() (err error) {
		if !first {
			if _, err = wtr._out.WriteRune(','); err != nil {
				return
			}
		} else {
			first = false
		}
		if wtr.Pretty {
			wtr._out.WriteString("\n")
			wtr._out.WriteString(padding[0:(2 * lvl)])
		}
		return
	}
	s := &Basic{}
	s.OnChild = func(r node.ChildRequest) (child node.Node, err error) {
		if !r.New {
			return nil, nil
		}
		if err = delim(); err != nil {
			return nil, err
		}
		if meta.IsList(r.Meta) {
			if err = wtr.beginList(wtr.ident(r.Path)); err != nil {
				return nil, err
			}
			return wtr.container(lvl + 1), nil

		}
		if err = wtr.beginContainer(wtr.ident(r.Path), lvl); err != nil {
			return nil, err
		}
		return wtr.container(lvl + 1), nil
	}
	s.OnEndEdit = func(r node.NodeRequest) error {
		if !r.Selection.InsideList && meta.IsList(r.Selection.Meta()) {
			if err := wtr.endList(); err != nil {
				return err
			}
		} else {
			if err := wtr.endContainer(); err != nil {
				return err
			}
		}
		return nil
	}
	s.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		if !r.Write {
			panic("Not a reader")
		}
		if err = delim(); err != nil {
			return err
		}
		err = wtr.writeValue(r.Path, hnd.Val)
		return
	}
	s.OnNext = func(r node.ListRequest) (next node.Node, key []val.Value, err error) {
		if !r.New {
			return
		}
		if err = delim(); err != nil {
			return
		}
		if err = wtr.beginObject(); err != nil {
			return
		}
		return wtr.container(lvl + 1), r.Key, nil
	}
	return s
}

func (wtr *JSONWtr) ident(p *node.Path) string {
	var qualify bool
	s := p.Meta.(meta.Identifiable).Ident()
	thisMod := meta.OriginalModule(p.Meta)
	if p.Len() == 2 { // top-level
		qualify = true
	} else {
		parentMod := meta.OriginalModule(p.Parent.Meta)
		qualify = (parentMod != thisMod)
	}
	if qualify && wtr.QualifyNamespace {
		return fmt.Sprintf("%s:%s", thisMod.Ident(), s)
	}
	return s
}

func (wtr *JSONWtr) beginList(ident string) (err error) {
	if err = wtr.writeIdent(ident); err == nil {
		_, err = wtr._out.WriteRune('[')
	}
	return
}

func (wtr *JSONWtr) beginContainer(ident string, lvl int) (err error) {
	if err = wtr.writeIdent(ident); err != nil {
		return
	}
	if err = wtr.beginObject(); err != nil {
		return
	}
	return
}

func (wtr *JSONWtr) beginObject() (err error) {
	if err == nil {
		_, err = wtr._out.WriteRune('{')
	}
	return
}

func (wtr *JSONWtr) writeIdent(ident string) (err error) {
	if _, err = wtr._out.WriteRune(QUOTE); err != nil {
		return
	}
	if _, err = wtr._out.WriteString(ident); err != nil {
		return
	}
	if _, err = wtr._out.WriteRune(QUOTE); err != nil {
		return
	}
	_, err = wtr._out.WriteRune(':')
	return
}

func (wtr *JSONWtr) endList() (err error) {
	_, err = wtr._out.WriteRune(']')
	return
}

func (wtr *JSONWtr) endContainer() (err error) {
	_, err = wtr._out.WriteRune('}')
	return
}

func (wtr *JSONWtr) writeValue(p *node.Path, v val.Value) error {
	wtr.writeIdent(wtr.ident(p))
	if v.Format().IsList() {
		if _, err := wtr._out.WriteRune('['); err != nil {
			return err
		}
	}
	lerr := val.Reduce(v, nil, func(i int, item val.Value, ierr interface{}) interface{} {
		if ierr != nil {
			return ierr
		}
		if i > 0 {
			if _, err := wtr._out.WriteRune(','); err != nil {
				return err
			}
		}
		switch item.Format() {
		case val.FmtIdentityRef:
			idtyStr := item.String()
			leafMod := meta.OriginalModule(p.Meta)
			bases := p.Meta.(meta.HasType).Type().Base()
			idty := meta.FindIdentity(bases, idtyStr)
			if idty == nil {
				return fmt.Errorf("could not find ident '%s'", idtyStr)
			}
			idtyMod := meta.RootModule(idty)
			if idtyMod != leafMod {
				idtyStr = fmt.Sprint(idtyMod.Ident(), ":", idtyStr)
			}
			if err := wtr.writeString(idtyStr); err != nil {
				return err
			}
		case val.FmtString, val.FmtBinary:
			if err := wtr.writeString(item.String()); err != nil {
				return err
			}
		case val.FmtEnum:
			if wtr.EnumAsIds {
				id := strconv.Itoa(item.(val.Enum).Id)
				if _, err := wtr._out.WriteString(id); err != nil {
					return err
				}
			} else {
				if err := wtr.writeString(item.(val.Enum).Label); err != nil {
					return err
				}
			}
		case val.FmtDecimal64:
			f := item.Value().(float64)
			if _, err := wtr._out.WriteString(strconv.FormatFloat(f, 'f', -1, 64)); err != nil {
				return err
			}
		case val.FmtAny:
			var data []byte
			var err error
			x := item.Value()
			if sel, ok := x.(node.Selection); ok {
				wtr := &JSONWtr{Out: wtr._out, Pretty: wtr.Pretty}
				err = sel.InsertInto(wtr.Node()).LastErr
				if err != nil {
					return err
				}
			} else {
				data, err = json.Marshal(item.Value())
				if err != nil {
					return err
				}
				if _, err := wtr._out.Write(data); err != nil {
					return err
				}
			}
		default:
			if _, err := wtr._out.WriteString(item.String()); err != nil {
				return err
			}
		}
		return nil
	})
	if lerr != nil {
		return lerr.(error)
	}
	if v.Format().IsList() {
		if _, err := wtr._out.WriteRune(']'); err != nil {
			return err
		}
	}
	return nil
}

func (wtr *JSONWtr) writeString(s string) error {
	clean := bytes.NewBuffer(make([]byte, len(s)+2))
	clean.Reset()
	writeString(clean, s, true)
	_, ioErr := wtr._out.Write(clean.Bytes())
	return ioErr
}
