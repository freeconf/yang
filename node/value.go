package node

import (
	"fmt"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/val"
)

func EncodeKey(v []val.Value) string {
	var s string
	// TODO: read RFC and escape chars including commas
	for i, val := range v {
		if i > 0 {
			s += "," + val.String()
		}
		s += val.String()
	}
	return s
}

func NewValuesByString(m []meta.HasDataType, objs ...string) ([]val.Value, error) {
	var err error
	vals := make([]val.Value, len(m))
	for i, obj := range objs {
		vals[i], err = NewValue(m[i].DataType(), obj)
		if err != nil {
			return nil, err
		}
	}
	return vals, nil
}

func NewValues(m []meta.HasDataType, objs ...interface{}) ([]val.Value, error) {
	var err error
	vals := make([]val.Value, len(m))
	for i, obj := range objs {
		vals[i], err = NewValue(m[i].DataType(), obj)
		if err != nil {
			return nil, err
		}
	}
	return vals, nil
}

// Incoming value should be of appropriate type according to given data type format
func NewValue(typ *meta.DataType, v interface{}) (val.Value, error) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf("%s : %s", typ.TypeIdent(), r))
		}
	}()
	if v == nil {
		return nil, nil
	}
	switch typ.Format() {
	case val.FmtEnum:
		return toEnum(typ.Enum(), v)
	case val.FmtEnumList:
		return toEnumList(typ.Enum(), v)
	}
	return val.Conv(typ.Format(), v)
}

func toEnumList(src val.EnumList, v interface{}) (val.EnumList, error) {
	switch x := v.(type) {
	case []string:
		l := make([]val.Enum, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toEnum(src, x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	case []interface{}:
		l := make([]val.Enum, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toEnum(src, x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	case []int:
		l := make([]val.Enum, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toEnum(src, x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	default:
		if e, err := toEnum(src, v); err != nil {
			return val.EnumList([]val.Enum{e}), nil
		}
	}
	return nil, c2.NewErr(fmt.Sprintf("could not coerse %v into enum list", v))
}

func toEnum(src val.EnumList, v interface{}) (val.Enum, error) {
	id, isNum := val.Conv(val.FmtInt32, v)
	if isNum == nil {
		if e, found := src.ById(id.Value().(int)); found {
			return e, nil
		}
	} else {
		label, isLabel := val.Conv(val.FmtString, v)
		if isLabel == nil {
			if e, found := src.ByLabel(label.String()); found {
				return e, nil
			}
		}
	}
	return val.Enum{}, c2.NewErr(fmt.Sprintf("could not coerse %v into enum", v))
}
