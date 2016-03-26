package browse

import (
	"bufio"
	"encoding/binary"
	"meta"
	"io"
	"node"
	"blit"
	"fmt"
)

const (
	BinBeginContainer rune = '{'
	BinBeginList = '['
	BinLeaf = ':'
	BinEof = '.'
)

type BinaryWriter struct {
	Out *bufio.Writer
	LastErr error
}

func NewBinaryWriter(out io.Writer) *BinaryWriter {
	return &BinaryWriter{
		Out: bufio.NewWriter(out),
	}
}

/**
  Format:

  container : '{'
     string         // meta
     | container[..]
     | list[...]
     | leaf[..]
     "!"


  list : '['
    value[0...n]     // key
    | container[..]
    | list[...]
    | leaf[..]

  leaf : ':'
      string     // meta
      value

  value :
      | string
      | int
      | bool
      | ...
      | array

  string :
      int     // string len in bytes
      bytes  byte[len]

  array :
      int      // number of values in array
      value[len]

 */

func (self *BinaryWriter) Node() node.Node {
	n := &node.MyNode{}
	n.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		if ! r.New {
			return nil, nil
		}
		self.WriteOp(BinBeginContainer)
		self.WriteString(r.Meta.GetIdent())
		return n, self.LastErr
	}
	n.OnWrite = func(r node.FieldRequest, v *node.Value) error {
		self.WriteOp(BinLeaf)
		self.WriteString(r.Meta.GetIdent())
		self.WriteValue(v)
		return self.LastErr
	}
	n.OnNext = func(r node.ListRequest) (node.Node, []*node.Value, error) {
		self.WriteOp(BinBeginList)
		keyMeta := r.Meta.KeyMeta()
		if len(keyMeta) > 0 {
			if len(keyMeta) != len(r.Key) {
				panic("no key given : " + r.Meta.GetIdent())
			}
			for _, k := range r.Key {
				self.WriteValue(k)
			}
		}
		return n, r.Key, self.LastErr
	}
	n.OnEvent = func(sel *node.Selection, e node.Event) error {
		switch e.Type {
		case node.END_TREE_EDIT, node.LEAVE:
			self.Out.Flush()
		}
		return self.LastErr
	}
	return n
}

func (self *BinaryWriter) WriteOp(op rune) {
	if self.LastErr == nil {
		_, self.LastErr = self.Out.WriteRune(op)
	}
}

func (self *BinaryWriter) WriteValue(v *node.Value) {
	format := v.Type.Format()
	switch format {
	case meta.FMT_INT32, meta.FMT_ENUMERATION:
		self.WriteInt(v.Int)
	case meta.FMT_INT64:
		self.checkErr(binary.Write(self.Out, binary.BigEndian, v.Int64))
	case meta.FMT_INT64_LIST:
		self.WriteInt(len(v.Int64list))
		for _, i := range v.Int64list {
			self.checkErr(binary.Write(self.Out, binary.BigEndian, i))
		}
	case meta.FMT_BOOLEAN:
		self.WriteBool(v.Bool)
	case meta.FMT_STRING:
		self.WriteString(v.Str)
	case meta.FMT_DECIMAL64:
		self.checkErr(binary.Write(self.Out, binary.BigEndian, v.Float))
	case meta.FMT_DECIMAL64_LIST:
		self.WriteInt(len(v.Floatlist))
		for _, f := range v.Floatlist {
			self.checkErr(binary.Write(self.Out, binary.BigEndian, f))
		}
	case meta.FMT_INT32_LIST, meta.FMT_ENUMERATION_LIST:
		self.WriteInt(len(v.Intlist))
		for _, i := range v.Intlist {
			self.WriteInt(i)
		}
	case meta.FMT_STRING_LIST:
		self.WriteInt(len(v.Strlist))
		for _, s := range v.Strlist {
			self.WriteString(s)
		}
	case meta.FMT_BOOLEAN_LIST:
		self.WriteInt(len(v.Boollist))
		for _, b := range v.Boollist {
			self.WriteBool(b)
		}
	default:
		panic(fmt.Sprintf("format not implemented %s(%d)", format.String(), format))
	}
}

func (self *BinaryWriter) WriteBool(b bool) {
	var i8 int8
	if b {
		i8 = 1
	}
	self.checkErr(binary.Write(self.Out, binary.BigEndian, i8))
}

func (self *BinaryWriter) WriteInt(i int) {
	self.checkErr(binary.Write(self.Out, binary.BigEndian, int32(i)))
}

func (self *BinaryWriter) WriteString(s string) {
	self.WriteInt(len(s))
	self.Out.WriteString(s)
}


type BinaryReader struct {
	In        *bufio.Reader
	op        rune
	nextIdent string
	LastErr   error
}

func NewBinaryReader(in io.Reader) *BinaryReader {
	return &BinaryReader{
		In: bufio.NewReader(in),
	}
}

func (self *BinaryWriter) checkErr(e error) {
	if self.LastErr == nil {
		self.LastErr = e
	}
}


func (self *BinaryReader) Node() node.Node {
	n := &node.MyNode{}
	self.NextOp()
	n.OnChoose = func(sel *node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
		cases := meta.NewMetaListIterator(choice, false)
		for cases.HasNextMeta() {
			kase := cases.NextMeta().(*meta.ChoiceCase)
			props := meta.NewMetaListIterator(kase, true)
			for props.HasNextMeta() {
				prop := props.NextMeta()
				if self.nextIdent == prop.GetIdent() {
					return kase, nil
				}
			}
		}
		return nil, nil
	}
	n.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		if r.New {
			return nil, blit.NewErr("Not a writer")
		}
		if self.op != BinBeginContainer || r.Meta.GetIdent() != self.nextIdent {
			return nil, self.LastErr
		}
		self.NextOp()
		return n, self.LastErr
	}
	n.OnRead = func(r node.FieldRequest) (*node.Value, error) {
		if self.op != BinLeaf || r.Meta.GetIdent() != self.nextIdent {
			return nil, self.LastErr
		}
		if r.Meta.GetIdent() != self.nextIdent {
			return nil, nil
		}
		v := self.ReadValue(r.Meta)
		self.NextOp()
		return v, self.LastErr
	}
	n.OnNext = func(r node.ListRequest) (node.Node, []*node.Value, error) {
		if r.New {
			return nil, nil, blit.NewErr("Not a writer")
		}
		if self.op != BinBeginList {
			return nil, nil, self.LastErr
		}
		var key []*node.Value
		keyMeta := r.Meta.KeyMeta()
		if len(keyMeta) > 0 {
			key = make([]*node.Value, len(keyMeta))
			for i, k := range keyMeta {
				key[i] = self.ReadValue(k)
			}
		}
		self.NextOp()

		return n, key, self.LastErr
	}
	return n
}

func (self *BinaryReader) NextOp() {
	if self.LastErr == nil {
		self.op, _, self.LastErr = self.In.ReadRune()
		if self.LastErr == io.EOF {
			self.op = BinEof
			self.LastErr = nil
		}
	}
	if self.op == BinBeginContainer || self.op == BinLeaf {
		self.nextIdent = self.ReadString()
	} else {
		self.nextIdent = ""
	}
}

func (self *BinaryReader) ReadValue(m meta.HasDataType) *node.Value {
	v := node.Value{
		Type:m.GetDataType(),
	}
	format := m.GetDataType().Format()
	switch format {
	case meta.FMT_INT32:
		v.Int = self.ReadInt()
	case meta.FMT_ENUMERATION:
		v.SetEnum(self.ReadInt())
	case meta.FMT_BOOLEAN:
		v.Bool = self.ReadBool()
	case meta.FMT_STRING:
		v.Str = self.ReadString()
	case meta.FMT_DECIMAL64:
		self.checkErr(binary.Read(self.In, binary.BigEndian, &v.Float))
	case meta.FMT_DECIMAL64_LIST:
		len := self.ReadInt()
		v.Floatlist = make([]float64, len)
		for i, _ := range v.Floatlist {
			self.checkErr(binary.Read(self.In, binary.BigEndian, &v.Floatlist[i]))
		}
	case meta.FMT_INT64:
		self.checkErr(binary.Read(self.In, binary.BigEndian, &v.Int64))
	case meta.FMT_INT64_LIST:
		len := self.ReadInt()
		v.Int64list = make([]int64, len)
		for i, _ := range v.Int64list {
			self.checkErr(binary.Read(self.In, binary.BigEndian, &v.Int64list[i]))
		}
	case meta.FMT_INT32_LIST, meta.FMT_ENUMERATION_LIST:
		len := self.ReadInt()
		v.Intlist = make([]int, len)
		for i, _ := range v.Intlist {
			v.Intlist[i] = self.ReadInt()
		}
		if format == meta.FMT_ENUMERATION_LIST {
			v.SetEnumList(v.Intlist)
		}
	case meta.FMT_STRING_LIST:
		len := self.ReadInt()
		v.Strlist = make([]string, len)
		for i, _ := range v.Strlist {
			v.Strlist[i] = self.ReadString()
		}
	case meta.FMT_BOOLEAN_LIST:
		len := self.ReadInt()
		v.Boollist = make([]bool, len)
		for i, _ := range v.Boollist {
			v.Boollist[i] = self.ReadBool()
		}
	default:
		panic("format not supported " + format.String())
	}
	return &v
}

func (self *BinaryReader) ReadInt() int {
	var i int32
	self.checkErr(binary.Read(self.In, binary.BigEndian, &i))
	return int(i)
}

func (self *BinaryReader) ReadBool() bool {
	var i int8
	self.checkErr(binary.Read(self.In, binary.BigEndian, &i))
	return i != 0
}

func (self *BinaryReader) ReadString() string {
	// TODO: performance - probably could read in []bytes and make string from
	// same copy
	l := self.ReadInt()
	b := make([]byte, l)
	for bytesRead := 0; bytesRead < l && self.LastErr == nil; {
		var n int
		n, self.LastErr = self.In.Read(b[bytesRead:])
		bytesRead += n
	}
	s := string(b)
	return s
}

func (self *BinaryReader) checkErr(e error) {
	if self.LastErr == nil {
		self.LastErr = e
	}
}
