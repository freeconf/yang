package nodes

import (
	"bufio"
	"encoding/json"
	"io"
	"strconv"

	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"

	"bytes"

	"github.com/c2stack/c2g/meta"
)

const QUOTE = '"'

type JSONWtr struct {
	Out    io.Writer
	Pretty bool
	_out   *bufio.Writer
}

func WriteJSON(s node.Selection) (string, error) {
	var buff bytes.Buffer
	wtr := &JSONWtr{Out: &buff}
	err := s.InsertInto(wtr.Node()).LastErr
	return buff.String(), err
}

func WritePrettyJSON(s node.Selection) (string, error) {
	var buff bytes.Buffer
	wtr := &JSONWtr{Out: &buff, Pretty: true}
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
				if err := self.beginList(r.Selection.Meta().GetIdent()); err != nil {
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
			if err = self.beginList(r.Meta.GetIdent()); err != nil {
				return nil, err
			}
			return self.container(lvl + 1), nil

		}
		if err = self.beginContainer(r.Meta.GetIdent(), lvl); err != nil {
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

func (self *JSONWtr) writeValue(m meta.Meta, v val.Value) error {
	self.writeIdent(m.GetIdent())
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
			if err := self.writeString(item.(val.Enum).Label); err != nil {
				return err
			}
		case val.FmtDecimal64:
			f := item.Value().(float64)
			if _, err := self._out.WriteString(strconv.FormatFloat(f, 'f', -1, 64)); err != nil {
				return err
			}
		case val.FmtAny:
			if data, marshalErr := json.Marshal(item.Value()); marshalErr != nil {
				return marshalErr
			} else {
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
	clean, cleanErr := json.Marshal(s)
	if cleanErr != nil {
		return cleanErr
	}
	_, ioErr := self._out.Write(clean)
	return ioErr
}
