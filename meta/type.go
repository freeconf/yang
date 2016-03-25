package meta

import (
	"math"
	"strconv"
	"strings"
)

type DataType struct {
	Parent         HasDataType
	Ident          string
	FormatPtr      *DataFormat
	RangePtr       *string
	EnumerationRef []string
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
		}
		// TODO: else resolve typedefs
	}

	return *y.resolvedPtr
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
	if len(y.EnumerationRef) == 0 {
		y.EnumerationRef = []string{e}
	} else {
		y.EnumerationRef = append(y.EnumerationRef, e)
	}
}

func (y *DataType) SetEnumeration(en []string) {
	y.EnumerationRef = en
}

func (y *DataType) Enumeration() []string {
	if len(y.EnumerationRef) > 0 {
		return y.EnumerationRef
	}
	if resolved := y.resolve(); resolved != nil {
		return resolved.Enumeration()
	}
	return y.EnumerationRef
}
