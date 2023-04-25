package nodeutil

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"

	"bytes"

	"github.com/freeconf/yang/meta"
)

const QUOTE1 = '"'
const XML1 = '<'
const XML_CLOSE = '\\'
const XML2 = '>'

type XMLWtr struct {

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

func WriteXML(s node.Selection) (string, error) {
	buff := new(bytes.Buffer)
	wtr := &XMLWtr{Out: buff}
	err := s.InsertInto(wtr.Node()).LastErr
	return buff.String(), err
}

/*func WritePrettyJSON(s node.Selection) (string, error) {
	buff := new(bytes.Buffer)
	wtr := &XMLWtr{Out: buff, Pretty: true}
	err := s.InsertInto(wtr.Node()).LastErr
	return buff.String(), err
}*/

func (wtr XMLWtr) XML(s node.Selection) (string, error) {
	buff := new(bytes.Buffer)
	wtr.Out = buff
	err := s.InsertInto(wtr.Node()).LastErr
	return buff.String(), err
}

func NewXMLWtr(out io.Writer) *XMLWtr {
	return &XMLWtr{Out: out}
}

func (wtr *XMLWtr) Node() node.Node {
	wtr._out = bufio.NewWriter(wtr.Out)
	return wtr.container(0)
}

func (wtr *XMLWtr) container(lvl int) node.Node {
	first := true
	s := &Basic{}
	s.OnChild = func(r node.ChildRequest) (child node.Node, err error) {
		if !r.New {
			return nil, nil
		}
		if err = wtr.beginContainer(wtr.ident(r.Path)); err != nil {
			return nil, err
		}
		return wtr.container(lvl + 1), nil
	}
	s.OnBeginEdit = func(r node.NodeRequest) error {
		var ident string

		if first == true {
			first = false
			ident = wtr.ident(r.Selection.Path) + " xmlns=" + meta.OriginalModule(r.Selection.Path.Meta).Ident()
		} else {
			ident = wtr.ident(r.Selection.Path)
		}

		if err := wtr.beginContainer(ident); err != nil {
			return err
		}
		return nil
	}
	s.OnEndEdit = func(r node.NodeRequest) error {
		if err := wtr.endContainer(wtr.ident(r.Selection.Path)); err != nil {
			return err
		}
		if err := wtr._out.Flush(); err != nil {
			return err
		}
		return nil
	}
	s.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		if !r.Write {
			panic("Not a reader")
		}
		if err = wtr.writeOpenIdent(wtr.ident(r.Path)); err != nil {
			return err
		}
		if err = wtr.writeValue(r.Path, hnd.Val); err != nil {
			return err
		}
		if err = wtr.writeCloseIdent(wtr.ident(r.Path)); err != nil {
			return err
		}

		return nil
	}
	s.OnNext = func(r node.ListRequest) (next node.Node, key []val.Value, err error) {
		if !r.New {
			return
		}
		return wtr.container(lvl + 1), r.Key, nil
	}
	return s
}

func (wtr *XMLWtr) ident(p *node.Path) string {
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

func (wtr *XMLWtr) beginContainer(ident string) (err error) {
	if err = wtr.writeOpenIdent(ident); err != nil {
		return err
	}
	return
}

func (wtr *XMLWtr) writeOpenIdent(ident string) (err error) {
	if _, err = wtr._out.WriteRune(XML1); err != nil {
		return err
	}
	if _, err = wtr._out.WriteString(ident); err != nil {
		return err
	}
	if _, err = wtr._out.WriteRune(XML2); err != nil {
		return err
	}
	return nil
}

func (wtr *XMLWtr) writeCloseIdent(ident string) (err error) {
	if _, err = wtr._out.WriteRune(XML1); err != nil {
		return err
	}
	if _, err = wtr._out.WriteRune(XML_CLOSE); err != nil {
		return err
	}
	if _, err = wtr._out.WriteString(ident); err != nil {
		return err
	}
	if _, err = wtr._out.WriteRune(XML2); err != nil {
		return err
	}
	return nil
}

func (wtr *XMLWtr) endContainer(ident string) (err error) {
	if err = wtr.writeCloseIdent(ident); err != nil {
		return
	}
	return
}

func (wtr *XMLWtr) writeValue(p *node.Path, v val.Value) error {
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
			base := p.Meta.(meta.HasType).Type().Base()
			idty, found := base.Derived()[idtyStr]
			if !found {
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
			//TO DO
			return nil
			/*var data []byte
			var err error
			x := item.Value()
			if sel, ok := x.(node.Selection); ok {
				wtr := &XMLWtr{Out: wtr._out, Pretty: wtr.Pretty}
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
			}*/
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

func (wtr *XMLWtr) writeString(s string) error {
	clean := bytes.NewBuffer(make([]byte, len(s)+2))
	clean.Reset()
	writeString(clean, s, true)
	_, ioErr := wtr._out.Write(clean.Bytes())
	return ioErr
}
