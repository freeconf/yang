package node

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"bytes"

	"github.com/c2stack/c2g/meta"
)

const QUOTE = '"'

type JsonWriter struct {
	out    *bufio.Writer
	pretty bool
}

type closerFunc func() error

func NewJsonWriter(out io.Writer) *JsonWriter {
	return &JsonWriter{
		out: bufio.NewWriter(out),
	}
}

func WriteJson(s Selection) (string, error) {
	var buff bytes.Buffer
	err := s.InsertInto(NewJsonWriter(&buff).Node()).LastErr
	return buff.String(), err
}

func NewJsonPretty(out io.Writer) *JsonWriter {
	return &JsonWriter{
		out:    bufio.NewWriter(out),
		pretty: true,
	}
}

func (self *JsonWriter) Node() Node {
	// JSON can begin at a container, inside a list or inside a container, each of these has
	// different results to make json legal
	return &Extend{
		Label: "JSON",
		Node:  self.container(0),
		OnBeginEdit: func(p Node, r NodeRequest) error {
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
		OnEndEdit: func(p Node, r NodeRequest) error {
			if meta.IsList(r.Selection.Meta()) && !r.Selection.InsideList {
				if err := self.endList(); err != nil {
					return err
				}
			}
			if err := self.endContainer(); err != nil {
				return err
			}
			if err := self.out.Flush(); err != nil {
				return err
			}
			return nil
		},
	}
}

func (self *JsonWriter) container(lvl int) Node {
	first := true
	delim := func() (err error) {
		if !first {
			if _, err = self.out.WriteRune(','); err != nil {
				return
			}
		} else {
			first = false
		}
		if self.pretty {
			self.out.WriteString("\n")
			self.out.WriteString(padding[0:(2 * lvl)])
		}
		return
	}
	s := &MyNode{Label: "JSON Write"}
	s.OnChild = func(r ChildRequest) (child Node, err error) {
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
	s.OnEndEdit = func(r NodeRequest) error {
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
	s.OnField = func(r FieldRequest, hnd *ValueHandle) (err error) {
		if !r.Write {
			panic("Not a reader")
		}
		if err = delim(); err != nil {
			return err
		}
		err = self.writeValue(r.Meta, hnd.Val)
		return
	}
	s.OnNext = func(r ListRequest) (next Node, key []*Value, err error) {
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

func (self *JsonWriter) beginList(ident string) (err error) {
	if err = self.writeIdent(ident); err == nil {
		_, err = self.out.WriteRune('[')
	}
	return
}

func (self *JsonWriter) beginContainer(ident string, lvl int) (err error) {
	if err = self.writeIdent(ident); err != nil {
		return
	}
	if err = self.beginObject(); err != nil {
		return
	}
	return
}

const padding = "                                                                          "

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

func (self *JsonWriter) writeValue(m meta.Meta, v *Value) error {
	self.writeIdent(m.GetIdent())
	// i, err := v.Type.Info()
	// if err != nil {
	// 	return err
	// }
	var err error
	if meta.IsListFormat(v.Format) {
		if _, err = self.out.WriteRune('['); err != nil {
			return err
		}
	}
	switch v.Format {
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
					return err
				}
			}
			if err = self.writeFloat(f); err != nil {
				return err
			}
		}
	case meta.FMT_STRING, meta.FMT_ENUMERATION:
		err = self.writeString(v.Str)
	case meta.FMT_BOOLEAN_LIST:
		for i, b := range v.Boollist {
			if i > 0 {
				if _, err = self.out.WriteRune(','); err != nil {
					return err
				}
			}
			if err = self.writeBool(b); err != nil {
				return err
			}
		}
	case meta.FMT_INT32_LIST:
		for i, n := range v.Intlist {
			if i > 0 {
				if _, err = self.out.WriteRune(','); err != nil {
					return err
				}
			}
			if err = self.writeInt(n); err != nil {
				return err
			}
		}
	case meta.FMT_INT64_LIST:
		for i, n := range v.Int64list {
			if i > 0 {
				if _, err = self.out.WriteRune(','); err != nil {
					return err
				}
			}
			if err = self.writeInt64(n); err != nil {
				return err
			}
		}
	case meta.FMT_STRING_LIST, meta.FMT_ENUMERATION_LIST:
		for i, s := range v.Strlist {
			if i > 0 {
				if _, err = self.out.WriteRune(','); err != nil {
					return err
				}
			}
			if err = self.writeString(s); err != nil {
				return err
			}
		}
	default:
		msg := fmt.Sprintf("JSON writing value type not implemented %s ", v.Format.String())
		return errors.New(msg)
	}
	if meta.IsListFormat(v.Format) {
		if _, err = self.out.WriteRune(']'); err != nil {
			return err
		}
	}
	return err
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

func (self *JsonWriter) writeString(s string) error {
	// PERFORMANCE: Using json.Marshal to encode json string, test if it's more
	// efficient to create and reuse a single encoder
	clean, cleanErr := json.Marshal(s)
	if cleanErr != nil {
		return cleanErr
	}
	_, ioErr := self.out.Write(clean)
	return ioErr
}
