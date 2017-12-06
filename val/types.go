package val

import (
	"fmt"
	"strconv"
	"strings"
)

type Value interface {
	Format() Format
	Value() interface{}
	String() string
}

type Comparable interface {
	Value
	Compare(Comparable) int
}

type Listable interface {
	Value
	Len() int
	Item(index int) Value
}

///////////////////////

type String string

func (String) Format() Format {
	return FmtString
}

func (x String) String() string {
	return string(x)
}

func (x String) Value() interface{} {
	return string(x)
}

func (x String) Compare(b Comparable) int {
	return strings.Compare(string(x), b.Value().(string))
}

///////////////////////

type StringList []string

func (StringList) Format() Format {
	return FmtStringList
}

func (x StringList) String() string {
	return fmt.Sprintf("%v", []string(x))
}

func (x StringList) Value() interface{} {
	return []string(x)
}

func (x StringList) Len() int {
	return len(x)
}

func (x StringList) Item(i int) Value {
	return String(x[i])
}

///////////////////////

type Bool bool

func (Bool) Format() Format {
	return FmtBool
}

func (x Bool) String() string {
	return fmt.Sprintf("%v", bool(x))
}

func (x Bool) Value() interface{} {
	return bool(x)
}

func (x Bool) Compare(y Comparable) int {
	yb := y.Value().(bool)
	xb := bool(x)
	if xb == yb {
		return 0
	}
	if xb {
		return 1
	}
	return -1
}

///////////////////////

type BoolList []bool

func (BoolList) Format() Format {
	return FmtBoolList
}

func (x BoolList) String() string {
	return fmt.Sprintf("%v", []bool(x))
}

func (x BoolList) Value() interface{} {
	return []bool(x)
}

func (x BoolList) Len() int {
	return len(x)
}

func (x BoolList) Item(i int) Value {
	return Bool(x[i])
}

///////////////////////

type Int32 int

func (Int32) Format() Format {
	return FmtInt32
}

func (x Int32) String() string {
	return strconv.Itoa(int(x))
}

func (x Int32) Value() interface{} {
	return int(x)
}

func (x Int32) Compare(y Comparable) int {
	return int(x) - y.Value().(int)
}

///////////////////////

type Int32List []int

func (Int32List) Format() Format {
	return FmtInt32List
}

func (x Int32List) String() string {
	return fmt.Sprintf("%v", x)
}

func (x Int32List) Value() interface{} {
	return []int(x)
}

func (x Int32List) Len() int {
	return len(x)
}

func (x Int32List) Item(i int) Value {
	return Int32(x[i])
}

///////////////////////

type Int64 int64

func (Int64) Format() Format {
	return FmtInt64
}

func (x Int64) String() string {
	return strconv.FormatInt(int64(x), 10)
}

func (x Int64) Value() interface{} {
	return int64(x)
}

func (x Int64) Compare(b Comparable) int {
	c := int64(x) - b.Value().(int64)
	if c < 0 {
		return -1
	} else if c > 0 {
		return 1
	}
	return 0
}

///////////////////////

type Int64List []int64

func (Int64List) Format() Format {
	return FmtInt64List
}

func (x Int64List) String() string {
	return fmt.Sprintf("%v", x)
}

func (x Int64List) Value() interface{} {
	return []int64(x)
}

func (x Int64List) Len() int {
	return len(x)
}

func (x Int64List) Item(i int) Value {
	return Int64(x[i])
}

///////////////////////

type Decimal64 float64

func (Decimal64) Format() Format {
	return FmtDecimal64
}

func (x Decimal64) String() string {
	return fmt.Sprintf("%d", float64(x))
}

func (x Decimal64) Value() interface{} {
	return float64(x)
}

func (x Decimal64) Compare(b Comparable) int {
	c := float64(x) - b.Value().(float64)
	if c < 0 {
		return -1
	} else if c > 0 {
		return 1
	}
	return 0
}

///////////////////////

type Decimal64List []float64

func (Decimal64List) Format() Format {
	return FmtDecimal64List
}

func (x Decimal64List) String() string {
	return fmt.Sprintf("%v", x)
}

func (x Decimal64List) Value() interface{} {
	return []float64(x)
}

func (x Decimal64List) Len() int {
	return len(x)
}

func (x Decimal64List) Item(i int) Value {
	return Decimal64(x[i])
}

///////////////////////

type UInt32 uint

func (UInt32) Format() Format {
	return FmtUInt32
}

func (x UInt32) String() string {
	return strconv.FormatUint(uint64(x), 10)
}

func (x UInt32) Value() interface{} {
	return int64(x)
}

func (x UInt32) Compare(b Comparable) int {
	c := uint(x) - b.Value().(uint)
	if c < 0 {
		return -1
	} else if c > 0 {
		return 1
	}
	return 0
}

///////////////////////

type UInt32List []uint

func (UInt32List) Format() Format {
	return FmtUInt32List
}

func (x UInt32List) String() string {
	return fmt.Sprintf("%v", x)
}

func (x UInt32List) Value() interface{} {
	return []uint(x)
}

func (x UInt32List) Len() int {
	return len(x)
}

func (x UInt32List) Item(i int) Value {
	return UInt32(x[i])
}

///////////////////////

type UInt64 uint

func (UInt64) Format() Format {
	return FmtUInt32
}

func (x UInt64) String() string {
	return strconv.FormatUint(uint64(x), 10)
}

func (x UInt64) Value() interface{} {
	return int64(x)
}

func (x UInt64) Compare(b Comparable) int {
	c := uint64(x) - b.Value().(uint64)
	if c < 0 {
		return -1
	} else if c > 0 {
		return 1
	}
	return 0
}

///////////////////////

type UInt64List []uint64

func (UInt64List) Format() Format {
	return FmtUInt64List
}

func (x UInt64List) String() string {
	return fmt.Sprintf("%v", x)
}

func (x UInt64List) Value() interface{} {
	return []uint64(x)
}

func (x UInt64List) Len() int {
	return len(x)
}

func (x UInt64List) Item(i int) Value {
	return UInt64(x[i])
}

///////////////////////

type Enum struct {
	Id    int
	Label string
}

func (Enum) Format() Format {
	return FmtEnum
}

func (x Enum) String() string {
	return x.Label
}

func (x Enum) Compare(b Comparable) int {
	return x.Id - b.Value().(Enum).Id
}

func (x Enum) Value() interface{} {
	return x
}

func (x Enum) Empty() bool {
	return x == Enum{}
}

///////////////////////

type EnumList []Enum

func (EnumList) Format() Format {
	return FmtEnumList
}

func (e EnumList) String() string {
	var s string
	for i, x := range e {
		if i != 0 {
			s += ","
		}
		s += x.Label
	}
	return s
}

func (x EnumList) Value() interface{} {
	return x
}

func (e EnumList) NextId() int {
	if len(e) == 0 {
		return 0
	}
	return e[len(e)-1].Id + 1
}

func (e EnumList) Labels() []string {
	l := make([]string, len(e))
	for i := range e {
		l[i] = e[i].Label
	}
	return l
}

func (e EnumList) Ids() []int {
	l := make([]int, len(e))
	for i := range e {
		l[i] = e[i].Id
	}
	return l
}

func (e EnumList) ById(id int) (Enum, bool) {
	for _, i := range e {
		if i.Id == id {
			return i, true
		}
	}
	return Enum{}, false
}

func (e EnumList) ByLabel(label string) (Enum, bool) {
	for _, i := range e {
		if i.Label == label {
			return i, true
		}
	}
	return Enum{}, false
}

func (x EnumList) Len() int {
	return len(x)
}

func (x EnumList) Item(i int) Value {
	return x[i]
}

func (x EnumList) Add(e string) EnumList {
	var id int
	if len(x) == 0 {
		id = 0
	} else {
		id = x[len(x)-1].Id + 1
	}
	return append(x, Enum{
		Label: e,
		Id:    id,
	})
}

///////////////////////

type Any struct {
	Thing interface{}
}

func (Any) Format() Format {
	return FmtAny
}

func (x Any) String() string {
	return fmt.Sprintf("%v", x.Thing)
}

func (x Any) Value() interface{} {
	return x.Thing
}

///////////////////////
type Union struct {
	Format     Format
	Bool       bool
	Int        int
	UInt       uint
	Int64      int64
	UInt64     uint64
	Int64list  []int64
	UInt64list []uint64
	Str        string
	Float      float64
	Intlist    []int
	UIntlist   []uint
	Strlist    []string
	Boollist   []bool
	Floatlist  []float64
	Keys       []string
	AnyData    interface{}
}
