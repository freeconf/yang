package meta

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/c2stack/c2g/c2"
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

func (y *DataType) resolve() (*DataType, error) {
	if y.resolvedPtr == nil {
		var resolved *DataType
		y.resolvedPtr = &resolved
		if y.FormatPtr != nil && (*y.FormatPtr == FMT_LEAFREF || *y.FormatPtr == FMT_LEAFREF_LIST) {
			if y.PathPtr == nil {
				return nil, c2.NewErr("Missing 'path' on leafref " + y.Ident)
			}
			resolvedMeta, err := FindByPath(y.Parent.GetParent(), *y.PathPtr)
			if err != nil {
				return nil, err
			} else if resolvedMeta == nil {
				return nil, c2.NewErr("Could not resolve 'path' on leafref " + y.Ident)
			}
			resolved = resolvedMeta.(HasDataType).GetDataType()
		} else if y.FormatPtr == nil {
			var err error
			if resolved, err = y.findTypedef(y.Parent); err != nil {
				return nil, err
			}
		}
	}

	return *y.resolvedPtr, nil
}

type TypeInfo struct {
	Format     DataFormat
	MinLength  int
	MaxLength  int
	Path       string
	HasDefault bool
	Default    string
	Enum       Enum
}

func (y *DataType) Info() (info TypeInfo, err error) {
	var r *DataType
	r, err = y.resolve()
	if err != nil {
		return
	}
	if r != nil {
		if info, err = r.Info(); err != nil {
			return
		}
	}
	if y.FormatPtr != nil && *y.FormatPtr != FMT_LEAFREF && *y.FormatPtr != FMT_LEAFREF_LIST {
		info.Format = *y.FormatPtr
		if _, isLeafList := y.Parent.(*LeafList); isLeafList && info.Format < FMT_BINARY_LIST {
			info.Format += 1024
		}
	}
	if y.PathPtr != nil {
		info.Path = *y.PathPtr
	}
	if y.MinLengthPtr != nil {
		info.MinLength = *y.MinLengthPtr
	}
	if y.MaxLengthPtr != nil {
		info.MaxLength = *y.MaxLengthPtr
	}
	if y.DefaultPtr != nil {
		info.HasDefault = true
		info.Default = *y.DefaultPtr
	}
	if y.EnumerationRef != nil {
		info.Enum = y.EnumerationRef
	}
	return
}

func (y *DataType) findTypedef(m Meta) (*DataType, error) {
	if tdefs, hasTds := m.(HasTypedefs); hasTds {
		if foundTd, err := FindByIdent2(tdefs.GetTypedefs(), y.Ident); err != nil {
			return nil, err
		} else if foundTd != nil {
			return foundTd.(*Typedef).GetDataType(), nil
		}
	}
	if m.GetParent() != nil {
		return y.findTypedef(m.GetParent())
	}
	return nil, nil
}

func (y *DataType) SetFormat(format DataFormat) {
	if format > 0 {
		y.FormatPtr = &format
	}
}

func (y *DataType) SetPath(path string) {
	y.PathPtr = &path
}

func (y *DataType) SetMinLength(len int) {
	y.MinLengthPtr = &len
}

func (y *DataType) SetMaxLength(len int) {
	y.MaxLengthPtr = &len
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

func (y *DataType) SetDefault(def string) {
	y.DefaultPtr = &def
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
