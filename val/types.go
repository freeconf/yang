package val

import (
	"fmt"
	"strconv"
)

type Value interface {
	Format() Format
	Value() interface{}
	//Equal(Value) bool
	String() string
}

///////////////////////

type Bool bool

func (Bool) Format() Format {
	return FmtBool
}

func (b Bool) String() string {
	return fmt.Sprintf("%v", bool(b))
}

func (b Bool) Value() interface{} {
	return bool(b)
}

///////////////////////

type BoolList []bool

func (BoolList) Format() Format {
	return FmtBoolList
}

func (b BoolList) String() string {
	return fmt.Sprintf("%v", []bool(b))
}

func (b BoolList) Value() interface{} {
	return []bool(b)
}

///////////////////////

type Int32 int

func (Int32) Format() Format {
	return FmtInt32
}

func (i Int32) String() string {
	return strconv.Itoa(int(i))
}

func (i Int32) Value() interface{} {
	return int(i)
}

///////////////////////

type Int32List []int

func (Int32List) Format() Format {
	return FmtInt32List
}

func (i Int32List) String() string {
	return fmt.Sprintf("%v", i)
}

func (i Int32List) Value() interface{} {
	return []int(i)
}

///////////////////////

type Int64 int64

func (Int64) Format() Format {
	return FmtInt64
}

func (i Int64) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int64) Value() interface{} {
	return int64(i)
}

///////////////////////

type Enum EnumVal

type EnumVal struct {
	Id    int
	Label string
}

func (Enum) Format() Format {
	return FmtEnum
}

func (e Enum) String() string {
	return e.Label
}

///////////////////////

type EnumList []EnumVal

func (EnumList) Format() Format {
	return FmtEnum
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
