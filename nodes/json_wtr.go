package nodes

import (
	"bufio"
	"encoding/json"
	"io"
	"strconv"

	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/val"

	"bytes"

	"github.com/freeconf/gconf/meta"
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

func (self *JSONWtr) Node() node.Node {
	// JSON can begin at a container, inside a list or inside a container, each of these has
	// different results to make json legal
	self._out = bufio.NewWriter(self.Out)
	return &Extend{
		Base: self.container(0),
		OnBeginEdit: func(p node.Node, r node.NodeRequest) error {
			if err := self.beginObject(); err != nil {
				return err
			}
			if meta.IsList(r.Selection.Meta()) && !r.Selection.InsideList {
				if err := self.beginList(r.Selection.Meta().Ident()); err != nil {
					return err
				}
			}
			return nil
		},
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			if meta.IsList(r.Selection.Meta()) && !r.Selection.InsideList {
				if err := self.endList(); err != nil {
					return err
				}
			}
			if err := self.endContainer(); err != nil {
				return err
			}
			if err := self._out.Flush(); err != nil {
				return err
			}
			return nil
		},
	}
}

func (self *JSONWtr) container(lvl int) node.Node {
	first := true
	delim := func() (err error) {
		if !first {
			if _, err = self._out.WriteRune(','); err != nil {
				return
			}
		} else {
			first = false
		}
		if self.Pretty {
			self._out.WriteString("\n")
			self._out.WriteString(padding[0:(2 * lvl)])
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
			if err = self.beginList(r.Meta.Ident()); err != nil {
				return nil, err
			}
			return self.container(lvl + 1), nil

		}
		if err = self.beginContainer(r.Meta.Ident(), lvl); err != nil {
			return nil, err
		}
		return self.container(lvl + 1), nil
	}
	s.OnEndEdit = func(r node.NodeRequest) error {
		if !r.Selection.InsideList && meta.IsList(r.Selection.Meta()) {
			if err := self.endList(); err != nil {
				return err
			}
		} else {
			if err := self.endContainer(); err != nil {
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
		err = self.writeValue(r.Meta, hnd.Val)
		return
	}
	s.OnNext = func(r node.ListRequest) (next node.Node, key []val.Value, err error) {
		if !r.New {
			return
		}
		if err = delim(); err != nil {
			return
		}
		if err = self.beginObject(); err != nil {
			return
		}
		return self.container(lvl + 1), r.Key, nil
	}
	return s
}

func (self *JSONWtr) beginList(ident string) (err error) {
	if err = self.writeIdent(ident); err == nil {
		_, err = self._out.WriteRune('[')
	}
	return
}

func (self *JSONWtr) beginContainer(ident string, lvl int) (err error) {
	if err = self.writeIdent(ident); err != nil {
		return
	}
	if err = self.beginObject(); err != nil {
		return
	}
	return
}

func (self *JSONWtr) beginObject() (err error) {
	if err == nil {
		_, err = self._out.WriteRune('{')
	}
	return
}

func (self *JSONWtr) writeIdent(ident string) (err error) {
	if _, err = self._out.WriteRune(QUOTE); err != nil {
		return
	}
	if _, err = self._out.WriteString(ident); err != nil {
		return
	}
	if _, err = self._out.WriteRune(QUOTE); err != nil {
		return
	}
	_, err = self._out.WriteRune(':')
	return
}

func (self *JSONWtr) endList() (err error) {
	_, err = self._out.WriteRune(']')
	return
}

func (self *JSONWtr) endContainer() (err error) {
	_, err = self._out.WriteRune('}')
	return
}

func (self *JSONWtr) writeValue(m meta.Definition, v val.Value) error {
	self.writeIdent(m.Ident())
	if v.Format().IsList() {
		if _, err := self._out.WriteRune('['); err != nil {
			return err
		}
	}
	lerr := val.Reduce(v, nil, func(i int, item val.Value, ierr interface{}) interface{} {
		if ierr != nil {
			return ierr
		}
		if i > 0 {
			if _, err := self._out.WriteRune(','); err != nil {
				return err
			}
		}
		switch item.Format() {
		case val.FmtString:
			if err := self.writeString(item.String()); err != nil {
				return err
			}
		case val.FmtEnum:
			if self.EnumAsIds {
				id := strconv.Itoa(item.(val.Enum).Id)
				if _, err := self._out.WriteString(id); err != nil {
					return err
				}
			} else {
				if err := self.writeString(item.(val.Enum).Label); err != nil {
					return err
				}
			}
		case val.FmtDecimal64:
			f := item.Value().(float64)
			if _, err := self._out.WriteString(strconv.FormatFloat(f, 'f', -1, 64)); err != nil {
				return err
			}
		case val.FmtAny:
			var data []byte
			var err error
			x := item.Value()
			if sel, ok := x.(node.Selection); ok {
				wtr := &JSONWtr{Out: self._out, Pretty: self.Pretty}
				err = sel.InsertInto(wtr.Node()).LastErr
				if err != nil {
					return err
				}
			} else {
				data, err = json.Marshal(item.Value())
				if _, err := self._out.Write(data); err != nil {
					return err
				}
			}
		default:
			if _, err := self._out.WriteString(item.String()); err != nil {
				return err
			}
		}
		return nil
	})
	if lerr != nil {
		return lerr.(error)
	}
	if v.Format().IsList() {
		if _, err := self._out.WriteRune(']'); err != nil {
			return err
		}
	}
	return nil
}

func (self *JSONWtr) writeString(s string) error {
	// PERFORMANCE: Using json.Marshal to encode json string, test if it's more
	// efficient to create and reuse a single encoder
	clean := bytes.NewBuffer(make([]byte, len(s)+2))
	writeString(clean, s, true)
	_, ioErr := self._out.Write(clean.Bytes())
	return ioErr
}
