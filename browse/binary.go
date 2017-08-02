package browse

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
	"github.com/c2stack/c2g/val"
)

const (
	BinFormatV1 = 1
)

const (
	BinBeginContainer     rune = '{'
	BinEndListOrContainer      = '!'
	BinBeginList               = '['
	BinKey                     = '#'
	BinLeaf                    = ':'
	BinEof                     = '.'
)

type BinaryWriter struct {
	Out     *bufio.Writer
	LastErr error
}

func NewBinaryWriter(out io.Writer) *BinaryWriter {
	w := &BinaryWriter{
		Out: bufio.NewWriter(out),
	}
	w.WriteInt(BinFormatV1)
	return w
}

/**
  Format:

  struct : '{'
     string         // meta
     key           // optional
     | struct[..]
     | struct_array[...]
     | leaf[..]
     '}'

  key : '#'
     len int
     leafs leaf[len]

  struct_array : '['
    string         // meta
    | struct[..]
    ']'

  leaf : ':'
      string     // meta
      value

  value :
      | string
      | int
      | bool
      | ...
      | value_array

  string :
      len int            // string len in bytes
      data bytes[len]

  value_array :
      len     int        // number of values in array
      values value[len]

*/

func (self *BinaryWriter) Node() node.Node {
	n := &nodes.Basic{}
	var level int
	n.OnChild = func(r node.ChildRequest) (node.Node, error) {
		if !r.New {
			return nil, nil
		}
		self.WriteOp(BinBeginContainer)
		level++
		self.WriteString(r.Meta.GetIdent())
		return n, self.LastErr
	}
	n.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) error {
		if !r.Write {
			return nil
		}
		self.WriteOp(BinLeaf)
		self.WriteString(r.Meta.GetIdent())
		self.WriteValue(hnd.Val)
		return self.LastErr
	}
	n.OnNext = func(r node.ListRequest) (node.Node, []val.Value, error) {
		self.WriteOp(BinBeginList)
		keyMeta := r.Meta.KeyMeta()
		if len(keyMeta) > 0 {
			self.WriteOp(BinKey)
			if len(keyMeta) != len(r.Key) {
				panic("no key given : " + r.Meta.GetIdent())
			}
			self.WriteInt(len(r.Key))
			for i, k := range r.Key {
				self.WriteString(r.Meta.KeyMeta()[i].GetIdent())
				self.WriteValue(k)
			}
		}
		return n, r.Key, self.LastErr
	}
	n.OnEndEdit = func(r node.NodeRequest) error {
		if level > 0 {
			self.WriteOp(BinEndListOrContainer)
		}
		if level == 0 {
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

func (self *BinaryWriter) WriteValue(v val.Value) {
	if v.Format().IsList() {
		self.WriteInt(v.(val.Listable).Len())
	}
	val.ForEach(v, func(i int, item val.Value) {
		switch x := item.(type) {
		case val.String:
			self.WriteString(v.String())
		case val.Enum:
			self.WriteInt(x.Id)
		case val.Int32:
			self.WriteInt(int(x))
		case val.Int64:
			self.checkErr(binary.Write(self.Out, binary.BigEndian, v.Value().(int64)))
		case val.Decimal64:
			self.checkErr(binary.Write(self.Out, binary.BigEndian, v.Value().(float64)))
		case val.Bool:
			self.WriteBool(v.Value().(bool))
		default:
			panic(fmt.Sprintf("format not implemented %s", item.Format()))
		}
	})
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
	r := &BinaryReader{
		In: bufio.NewReader(in),
	}
	ver := r.ReadInt()
	if ver != BinFormatV1 {
		panic(fmt.Sprintf("unknown binary file format version : %d", ver))
	}
	return r
}

func (self *BinaryWriter) checkErr(e error) {
	if self.LastErr == nil {
		self.LastErr = e
	}
}

func (self *BinaryReader) Node() node.Node {
	n := &nodes.Basic{}
	self.NextOp()
	n.OnChoose = func(sel node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
		cases := meta.NewMetaListIterator(choice, false)
		for cases.HasNextMeta() {
			kase, err := cases.NextMeta()
			if err != nil {
				return nil, err
			}
			props := meta.NewMetaListIterator(kase, true)
			for props.HasNextMeta() {
				prop, err := props.NextMeta()
				if err != nil {
					return nil, err
				}
				if self.nextIdent == prop.GetIdent() {
					return kase.(*meta.ChoiceCase), nil
				}
			}
		}
		return nil, nil
	}
	n.OnChild = func(r node.ChildRequest) (node.Node, error) {
		if r.New {
			return nil, c2.NewErr("Not a writer")
		}
		if self.op != BinBeginContainer || r.Meta.GetIdent() != self.nextIdent {
			return nil, self.LastErr
		}
		self.NextOp()
		return n, self.LastErr
	}
	n.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) error {
		if r.Write || self.op != BinLeaf || r.Meta.GetIdent() != self.nextIdent {
			return self.LastErr
		}
		if r.Meta.GetIdent() != self.nextIdent {
			return nil
		}
		var err error
		hnd.Val, err = self.ReadValue(r.Meta)
		if err != nil {
			return err
		}
		self.NextOp()
		return self.LastErr
	}
	n.OnNext = func(r node.ListRequest) (node.Node, []val.Value, error) {
		if r.New {
			return nil, nil, c2.NewErr("Not a writer")
		}
		if self.op != BinBeginList {
			return nil, nil, self.LastErr
		}
		var key []val.Value
		var err error
		self.NextOp()
		if self.op == BinKey {
			if key, err = self.readKey(r.Meta); err != nil {
				return nil, nil, err
			}
			self.NextOp()
		}
		return n, key, self.LastErr
	}
	n.OnEndEdit = func(r node.NodeRequest) error {
		if self.op != BinEndListOrContainer {
			tmpl := "bad file format or error in parser. Expected '!' but got '%c'"
			panic(fmt.Sprintf(tmpl, self.op))
		}
		self.NextOp()
		return nil
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

func (self *BinaryReader) readKey(m *meta.List) ([]val.Value, error) {
	givenKeySegments := self.ReadInt()
	expectedKeySegments := len(m.KeyMeta())
	var key []val.Value

	// It's ok if we don't expect any keys yet keys are given, but if expect keys
	// they better be in data stream
	if expectedKeySegments > 0 && givenKeySegments != expectedKeySegments {
		return nil, c2.NewErr("Expected keys in binary format for list " + m.GetIdent())
	}
	key = make([]val.Value, givenKeySegments)
	for i := 0; i < givenKeySegments; i++ {
		segIdent := self.ReadString()
		segMeta, err := meta.FindByIdent2(m, segIdent)
		if err != nil {
			return nil, err
		}
		key[i], err = self.ReadValue(segMeta.(meta.HasDataType))
		if err != nil {
			return nil, err
		}
	}
	return key, nil
}

func (self *BinaryReader) ReadValue(m meta.HasDataType) (val.Value, error) {
	info, err := m.GetDataType().Info()
	if err != nil {
		return nil, err
	}
	var len int
	if info.Format.IsList() {
		len = self.ReadInt()
	}
	switch info.Format {
	case val.FmtInt32:
		return val.Int32(self.ReadInt()), nil
	case val.FmtInt32List:
		list := make([]int, len)
		for i := 0; i < len; i++ {
			list[i] = self.ReadInt()
		}
		return val.Int32List(list), nil
	case val.FmtString:
		return val.String(self.ReadString()), nil
	case val.FmtStringList:
		list := make([]string, len)
		for i := 0; i < len; i++ {
			list[i] = self.ReadString()
		}
		return val.StringList(list), nil
	case val.FmtDecimal64:
		var f float64
		binary.Read(self.In, binary.BigEndian, &f)
		return val.Decimal64(f), nil
	case val.FmtDecimal64List:
		list := make([]float64, len)
		for i := 0; i < len; i++ {
			var f float64
			binary.Read(self.In, binary.BigEndian, &f)
			list[i] = f
		}
		return val.Decimal64List(list), nil
	case val.FmtBool:
		return val.Bool(self.ReadBool()), nil
	case val.FmtBoolList:
		list := make([]bool, len)
		for i := 0; i < len; i++ {
			list[i] = self.ReadBool()
		}
		return val.BoolList(list), nil
	case val.FmtEnum:
		id := self.ReadInt()
		e, found := info.Enum.ById(id)
		if !found {
			return nil, c2.NewErr(fmt.Sprintf("Enumeraton not found with value %d", id))
		}
		return e, nil
	case val.FmtEnumList:
		list := make([]val.Enum, len)
		for i := 0; i < len; i++ {
			id := self.ReadInt()
			e, found := info.Enum.ById(id)
			if !found {
				return nil, c2.NewErr(fmt.Sprintf("Enumeraton not found with value %d", id))
			}
			list[i] = e
		}
		return val.EnumList(list), nil
	}
	panic("format not supported " + info.Format.String())
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
