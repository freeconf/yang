package node

import (
	"bufio"
	"io"
	"github.com/c2g/meta"
	"strconv"
	"encoding/json"
	"fmt"
	"errors"
)

const QUOTE = '"'

type JsonWriter struct {
	out                *bufio.Writer
}

type closerFunc func() error

func NewJsonWriter(out io.Writer) *JsonWriter {
	return &JsonWriter{
		out:              bufio.NewWriter(out),
	}
}

func (self *JsonWriter) Node() Node {
	var closer closerFunc
	// JSON can begin at a container, inside a list or inside a container, each of these has
	// different results to make json legal
	return &Extend{
		Label: "JSON",
		Node: self.Container(self.endContainer),
		OnSelect: func(p Node, r ContainerRequest) (child Node, err error) {
			if closer == nil {
				self.beginObject()
				closer = self.endContainer
			}
			return p.Select(r)
		},
		OnNext: func(p Node, r ListRequest) (next Node, key []*Value, err error) {
			if closer == nil {
				self.beginObject()
				self.beginList(r.Meta.GetIdent())
				closer = func() (closeErr error) {
					if closeErr = self.endList(); closeErr == nil {
						closeErr = self.endContainer()
					}
					return closeErr
				}
			}
			return p.Next(r)
		},
		OnWrite: func(p Node, r FieldRequest, v *Value) (err error) {
			if closer == nil {
				self.beginObject()
				closer = self.endContainer
			}
			return p.Write(r, v)
		},
		OnEvent:func(p Node, s *Selection, e Event) error {
			var err error
			switch e.Type {
			case LEAVE, END_TREE_EDIT:
				if closer != nil {
					if err = closer(); err != nil {
						return err
					}
				}
				err = self.out.Flush()
			default:
				err = p.Event(s, e)
			}
			return err
		},
	}
}

func (self *JsonWriter) Container(closer closerFunc) Node {
	first := true
	delim := func() (err error) {
		if ! first {
			if _, err = self.out.WriteRune(','); err != nil {
				return
			}
		} else {
			first = false
		}
		return
	}
	s := &MyNode{Label: "JSON Write"}
	s.OnSelect = func(r ContainerRequest) (child Node, err error) {
		if ! r.New {
			return nil, nil
		}
		if err = delim(); err != nil {
			return nil, err
		}
		if meta.IsList(r.Meta) {
			if err = self.beginList(r.Meta.GetIdent()); err != nil {
				return nil, err
			}
			return self.Container(self.endList), nil

		}
		if err = self.beginContainer(r.Meta.GetIdent()); err != nil {
			return nil, err
		}
		return self.Container(self.endContainer), nil
	}
	s.OnEvent = func(sel *Selection, e Event) (err error) {
		switch e.Type {
		case LEAVE:
			err = closer()
		}
		return
	}
	s.OnWrite = func(r FieldRequest, v *Value) (err error) {
		if err = delim(); err != nil {
			return err
		}
		err = self.writeValue(r.Meta, v)
		return
	}
	s.OnNext = func(r ListRequest) (next Node, key []*Value, err error) {
		if ! r.New {
			return
		}
		if err = delim(); err != nil {
			return
		}
		if err = self.beginObject(); err != nil {
			return
		}
		return self.Container(self.endContainer), r.Key, nil
	}
	return s
}

func (self *JsonWriter) beginList(ident string) (err error) {
	if err = self.writeIdent(ident); err == nil {
		_, err = self.out.WriteRune('[')
	}
	return
}

func (self *JsonWriter) beginContainer(ident string) (err error) {
	if err = self.writeIdent(ident); err != nil {
		return
	}
	if err = self.beginObject(); err != nil {
		return
	}
	return
}

func (self *JsonWriter) beginObject() (err error) {
	if err == nil {
		_, err = self.out.WriteRune('{')
	}
	return
}

func (self *JsonWriter) writeIdent(ident string) (err error) {
	if _, err = self.out.WriteRune(QUOTE); err != nil {
		return
	}
	if _, err = self.out.WriteString(ident); err != nil {
		return
	}
	if _, err = self.out.WriteRune(QUOTE); err != nil {
		return
	}
	_, err = self.out.WriteRune(':')
	return
}

func (self *JsonWriter) endList() (err error) {
	_, err = self.out.WriteRune(']')
	return
}

func (self *JsonWriter) endContainer() (err error) {
	_, err = self.out.WriteRune('}')
	return
}

func (self *JsonWriter) writeValue(m meta.Meta, v *Value) (err error) {
	self.writeIdent(m.GetIdent())
	if meta.IsListFormat(v.Type.Format()) {
		if _, err = self.out.WriteRune('['); err != nil {
			return
		}
	}
	format := v.Type.Format()
	switch format {
	case meta.FMT_BOOLEAN:
		err = self.writeBool(v.Bool)
	case meta.FMT_ANYDATA:
		if data, marshalErr := json.Marshal(v.AnyData); marshalErr != nil {
			return marshalErr
		} else {
			self.out.Write(data)
		}
	case meta.FMT_INT64:
		err = self.writeInt64(v.Int64)
	case meta.FMT_UINT64:
		err = self.writeUInt64(v.UInt64)
	case meta.FMT_INT32:
		err = self.writeInt(v.Int)
	case meta.FMT_UINT32:
		err = self.writeUInt(v.UInt)
	case meta.FMT_DECIMAL64:
		err = self.writeFloat(v.Float)
	case meta.FMT_DECIMAL64_LIST:
		for i, f := range v.Floatlist {
			if i > 0 {
				if _, err = self.out.WriteRune(','); err != nil {
					return
				}
			}
			if err = self.writeFloat(f); err != nil {
				return
			}
		}
	case meta.FMT_STRING, meta.FMT_ENUMERATION:
		err = self.writeString(v.Str)
	case meta.FMT_BOOLEAN_LIST:
		for i, b := range v.Boollist {
			if i > 0 {
				if _, err = self.out.WriteRune(','); err != nil {
					return
				}
			}
			if err = self.writeBool(b); err != nil {
				return
			}
		}
	case meta.FMT_INT32_LIST:
		for i, n := range v.Intlist {
			if i > 0 {
				if _, err = self.out.WriteRune(','); err != nil {
					return
				}
			}
			if err = self.writeInt(n); err != nil {
				return
			}
		}
	case meta.FMT_INT64_LIST:
		for i, n := range v.Int64list {
			if i > 0 {
				if _, err = self.out.WriteRune(','); err != nil {
					return
				}
			}
			if err = self.writeInt64(n); err != nil {
				return
			}
		}
	case meta.FMT_STRING_LIST, meta.FMT_ENUMERATION_LIST:
		for i, s := range v.Strlist {
			if i > 0 {
				if _, err = self.out.WriteRune(','); err != nil {
					return
				}
			}
			if err = self.writeString(s); err != nil {
				return
			}
		}
	default:
		msg := fmt.Sprintf("JSON writing value type not implemented %s ", format.String())
		return errors.New(msg)
	}
	if meta.IsListFormat(v.Type.Format()) {
		if _, err = self.out.WriteRune(']'); err != nil {
			return
		}
	}
	return
}

func (self *JsonWriter) writeBool(b bool) (err error) {
	if b {
		_, err = self.out.WriteString("true")
	} else {
		_, err = self.out.WriteString("false")
	}
	return
}

func (self *JsonWriter) writeFloat(f float64) (err error) {
	_, err = self.out.WriteString(strconv.FormatFloat(f, 'f', -1, 64))
	return
}

func (self *JsonWriter) writeInt(i int) (err error) {
	_, err = self.out.WriteString(strconv.Itoa(i))
	return
}

func (self *JsonWriter) writeUInt(i uint) (err error) {
	_, err = self.out.WriteString(strconv.FormatUint(uint64(i), 10))
	return
}

func (self *JsonWriter) writeUInt64(i uint64) (err error) {
	_, err = self.out.WriteString(strconv.FormatUint(i, 10))
	return
}

func (self *JsonWriter) writeInt64(i int64) (err error) {
	_, err = self.out.WriteString(strconv.FormatInt(i, 10))
	return
}

func (self *JsonWriter) writeString(s string) (error) {
	// PERFORMANCE: Using json.Marshal to encode json string, test if it's more
	// efficient to create and reuse a single encoder
	clean, cleanErr := json.Marshal(s)
	if cleanErr != nil {
		return cleanErr
	}
	 _, ioErr := self.out.Write(clean)
	return ioErr
}

