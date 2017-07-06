package meta

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type DataType struct {
	Parent         HasDataType
	Ident          string
	FormatPtr      *DataFormat
	RangePtr       *string
	EnumerationRef Enum
	MinLengthPtr   *int
	MaxLengthPtr   *int
	PathPtr        *string
	PatternPtr     *string
	DefaultPtr     *string
	resolvedPtr    **DataType
	/*
		FractionDigits
		Bit
		Base
		RequireInstance
		Type?!  subtype?
	*/
}

type EnumRef struct {
	Label string
	Value int
}

func (self EnumRef) Nil() bool {
	return self.Label == ""
}

type Enum []EnumRef

func (self Enum) ByLabel(label string) EnumRef {
	for _, e := range self {
		if e.Label == label {
			return e
		}
	}
	return EnumRef{}
}

func (self Enum) String() string {
	var buffer bytes.Buffer
	for i, r := range self {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(r.Label)
		if r.Value != i {
			buffer.WriteString(fmt.Sprintf("(%d)", r.Value))
		}
	}
	return buffer.String()
}

func (self Enum) ByValue(v int) EnumRef {
	for _, e := range self {
		if e.Value == v {
			return e
		}
	}
	return EnumRef{}
}

func (self Enum) Update(ref EnumRef) {
	for _, e := range self {
		if e.Label == ref.Label {
			e = ref
			return
		}
	}
}

func NewDataType(Parent HasDataType, ident string) (t *DataType) {
	t = &DataType{Parent: Parent, Ident: ident}
	// if not found, then not internal type and Resolve should
	// determine type
	t.SetFormat(DataTypeImplicitFormat(ident))
	return
}

func (y *DataType) resolve() *DataType {
	if y.resolvedPtr == nil {
		var resolved *DataType
		y.resolvedPtr = &resolved
		if y.FormatPtr != nil && (*y.FormatPtr == FMT_LEAFREF || *y.FormatPtr == FMT_LEAFREF_LIST) {
			if y.PathPtr == nil {
				panic("Missing 'path' on leafref " + y.Ident)
			}
			resolvedMeta := FindByPath(y.Parent.GetParent(), *y.PathPtr)
			if resolvedMeta == nil {
				panic("Could not resolve 'path' on leafref " + y.Ident)
			}
			resolved = resolvedMeta.(HasDataType).GetDataType()
		} else if y.FormatPtr == nil {
			resolved = y.findTypedef(y.Parent)
		}
	}

	return *y.resolvedPtr
}

func (y *DataType) findTypedef(m Meta) *DataType {
	if tdefs, hasTds := m.(HasTypedefs); hasTds {
		if foundTd := FindByIdent2(tdefs.GetTypedefs(), y.Ident); foundTd != nil {
			return foundTd.(*Typedef).GetDataType()
		}
	}
	if m.GetParent() != nil {
		return y.findTypedef(m.GetParent())
	}
	return nil
}

func (y *DataType) SetFormat(format DataFormat) {
	if format > 0 {
		y.FormatPtr = &format
	}
}

func (y *DataType) Format() (format DataFormat) {
	if y.FormatPtr != nil && *y.FormatPtr != FMT_LEAFREF && *y.FormatPtr != FMT_LEAFREF_LIST {
		format = *y.FormatPtr
	} else if resolved := y.resolve(); resolved != nil {
		format = resolved.Format()
	}
	if _, isLeafList := y.Parent.(*LeafList); isLeafList && format < FMT_BINARY_LIST {
		format += 1024
	}
	return
}

func (y *DataType) SetPath(path string) {
	y.PathPtr = &path
}

func (y *DataType) Path() string {
	if y.PathPtr != nil {
		return *y.PathPtr
	}
	if resolved := y.resolve(); resolved != nil {
		return resolved.Path()
	}
	return ""
}

func (y *DataType) SetMinLength(len int) {
	y.MinLengthPtr = &len
}

func (y *DataType) MinLength() int {
	if y.MinLengthPtr != nil {
		return *y.MinLengthPtr
	}
	if resolved := y.resolve(); resolved != nil {
		return resolved.MinLength()
	}
	return 0
}

func (y *DataType) SetMaxLength(len int) {
	y.MaxLengthPtr = &len
}

func (y *DataType) MaxLength() int {
	if y.MaxLengthPtr != nil {
		return *y.MaxLengthPtr
	}
	if resolved := y.resolve(); resolved != nil {
		resolved.MaxLength()
	}
	return math.MaxInt32
}

func (y *DataType) DecodeLength(encoded string) error {
	/* TODO: Support multiple lengths using "|" */
	segments := strings.Split(encoded, "..")
	if len(segments) == 2 {
		if minLength, minErr := strconv.Atoi(segments[0]); minErr != nil {
			return minErr
		} else {
			y.MinLengthPtr = &minLength
		}
		if maxLength, maxErr := strconv.Atoi(segments[1]); maxErr != nil {
			return maxErr
		} else {
			y.MaxLengthPtr = &maxLength
		}
	} else {
		if maxLength, maxErr := strconv.Atoi(segments[0]); maxErr != nil {
			return maxErr
		} else {
			y.MaxLengthPtr = &maxLength
		}
	}
	return nil
}

func (y *DataType) HasDefault() bool {
	if y.DefaultPtr != nil {
		return true
	}
	if resolved := y.resolve(); resolved != nil {
		return resolved.HasDefault()
	}
	return false
}

func (y *DataType) SetDefault(def string) {
	y.DefaultPtr = &def
}

func (y *DataType) Default() string {
	if y.DefaultPtr != nil {
		return *y.DefaultPtr
	}
	if resolved := y.resolve(); resolved != nil {
		return resolved.Default()
	}
	return ""
}

func (y *DataType) AddEnumeration(e string) {
	eref := EnumRef{Label: e}
	if len(y.EnumerationRef) == 0 {
		y.EnumerationRef = Enum([]EnumRef{eref})
	} else {
		eref.Value = y.EnumerationRef[len(y.EnumerationRef)-1].Value + 1
		y.EnumerationRef = append(y.EnumerationRef, eref)
	}
}

func (y *DataType) AddEnumerationWithValue(e string, v int) {
	eref := EnumRef{Label: e, Value: v}
	if len(y.EnumerationRef) == 0 {
		y.EnumerationRef = Enum([]EnumRef{eref})
	} else {
		y.EnumerationRef = append(y.EnumerationRef, eref)
	}
}

func (y *DataType) SetEnumeration(en Enum) {
	y.EnumerationRef = en
}

func (y *DataType) Enumeration() Enum {
	if len(y.EnumerationRef) > 0 {
		return y.EnumerationRef
	}
	if resolved := y.resolve(); resolved != nil {
		return resolved.Enumeration()
	}
	return y.EnumerationRef
}
