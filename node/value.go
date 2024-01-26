package node

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
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

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NewValuesByString(m []meta.Leafable, objs ...string) ([]val.Value, error) {
	var err error
	l := minInt(len(m), len(objs))
	vals := make([]val.Value, len(m))
	for i := 0; i < l; i++ {
		vals[i], err = NewValue(m[i].Type(), objs[i])
		if err != nil {
			return nil, err
		}
	}
	return vals, nil
}

func NewValues(m []meta.Leafable, objs ...interface{}) ([]val.Value, error) {
	var err error
	vals := make([]val.Value, len(m))
	for i, obj := range objs {
		vals[i], err = NewValue(m[i].Type(), obj)
		if err != nil {
			return nil, err
		}
	}
	return vals, nil
}

// Incoming value should be of appropriate type according to given data type format
func NewValue(typ *meta.Type, v interface{}) (val.Value, error) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf("%s : %s", typ.Ident(), r))
		}
	}()
	if v == nil {
		return nil, nil
	}
	switch typ.Format() {
	case val.FmtIdentityRef:
		return toIdentRef(typ.Base(), v)
	case val.FmtIdentityRefList:
		return toIdentRefList(typ.Base(), v)
	case val.FmtEnum:
		return toEnum(typ.Enum(), v)
	case val.FmtEnumList:
		return toEnumList(typ.Enum(), v)
	case val.FmtUnion:
		cvt, _, err := val.ConvOneOf(typ.UnionFormats(), v)
		return cvt, err
	case val.FmtUnionList:
		return toUnionList(typ, v)
	case val.FmtLeafRef, val.FmtLeafRefList:
		return NewValue(typ.Resolve(), v)
	case val.FmtBitsList:
		return toBitsList(typ.Bits(), v)
	case val.FmtBits:
		return toBits(typ.Bits(), v)
	}
	return val.Conv(typ.Format(), v)
}

func toIdentRef(bases []*meta.Identity, v interface{}) (val.IdentRef, error) {
	var empty val.IdentRef
	x := fmt.Sprintf("%v", v)
	if colon := strings.IndexRune(x, ':'); colon > 0 {
		x = x[colon+1:]
	}

	ref := meta.FindIdentity(bases, x)
	if ref == nil {
		return empty, fmt.Errorf("could not find identity ref for %T:'%s'", v, x)
	}
	return val.IdentRef{Label: ref.Ident()}, nil
}

func toIdentRefList(base []*meta.Identity, v interface{}) (val.IdentRefList, error) {
	switch x := v.(type) {
	case string:
		ref, err := toIdentRef(base, x)
		if err != nil {
			return nil, err
		}
		return val.IdentRefList([]val.IdentRef{ref}), err
	case []string:
		var refs []val.IdentRef
		for _, s := range x {
			ref, err := toIdentRef(base, s)
			if err != nil {
				return nil, err
			}
			refs = append(refs, ref)
		}
		return refs, nil
	}
	return nil, fmt.Errorf("could not coerce '%v' into identref list", v)
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
	return nil, fmt.Errorf("could not coerce '%v' into enum list", v)
}

func toEnum(src val.EnumList, v interface{}) (val.Enum, error) {
	if id, isNum := val.Conv(val.FmtInt32, v); isNum == nil {
		if e, found := src.ById(id.Value().(int)); found {
			return e, nil
		}
	} else if id, isNum := val.Conv(val.FmtUInt32, v); isNum == nil {
		if e, found := src.ById(int(id.Value().(uint))); found {
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
	return val.Enum{}, fmt.Errorf("could not coerce '%v' into enum %v", v, src.String())
}

func toUnionList(typ *meta.Type, v interface{}) (val.Value, error) {
	if v == nil {
		return nil, nil
	}
	sliceValue := reflect.ValueOf(v)
	if sliceValue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("could not coerce %v into UnionList", v)
	}
	if sliceValue.Len() == 0 {
		return nil, nil
	}
	for _, t := range typ.Union() {
		result, err := NewValue(t, v)
		if err == nil {
			return result, err
		}
	}
	return nil, fmt.Errorf("could not coerce %v into UnionList", v)
}

func toBitsListHandler[V uint64 | int | uint | int64 | string | []string | float64](bitDefintions []*meta.Bit, vList []V) (val.BitsList, error) {
	result := make([]val.Bits, len(vList))
	var err error
	for i, v := range vList {
		if result[i], err = toBits(bitDefintions, v); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func toBitsList(bitDefintions []*meta.Bit, v interface{}) (val.BitsList, error) {
	switch x := v.(type) {
	case []string: // treat string as list of bit identifiers separated by space
		return toBitsListHandler(bitDefintions, x)
	case [][]string:
		return toBitsListHandler(bitDefintions, x)
	case []uint64:
		return toBitsListHandler(bitDefintions, x)
	case []int:
		return toBitsListHandler(bitDefintions, x)
	case []float64: // default type for decimals from JSON parser
		return toBitsListHandler(bitDefintions, x)
	}
	return nil, fmt.Errorf("could not coerce %v into BitList", v)
}

func toBitsValueHandler[V int | uint | int64 | float64](bitDefintions []*meta.Bit, v V) (val.Bits, error) {
	return toBits(bitDefintions, uint64(v))
}

func toBits(bitDefintions []*meta.Bit, v interface{}) (val.Bits, error) {
	result := val.Bits{}
	switch x := v.(type) {
	case []string: // labels only
		for _, strBit := range x {
			for _, bitDef := range bitDefintions {
				if strBit == bitDef.Ident() {
					result.Labels = append(result.Labels, strBit)
					result.Positions = result.Positions | (1 << bitDef.Position)
				}
			}
		}
		return result, nil
	case uint64: // positions only
		for _, bitDef := range bitDefintions {
			if x&(1<<bitDef.Position) != 0 {
				result.Positions = result.Positions | (1 << bitDef.Position)
				result.Labels = append(result.Labels, bitDef.Ident())
			}
		}
		return result, nil
	case string: // treat string as list of bit identifiers separated by space
		return toBits(bitDefintions, strings.Split(x, " "))
	case int:
		return toBitsValueHandler(bitDefintions, x)
	case uint:
		return toBitsValueHandler(bitDefintions, x)
	case float64:
		return toBitsValueHandler(bitDefintions, x)
	case int64:
		return toBitsValueHandler(bitDefintions, x)
	default:
		return result, fmt.Errorf("could not coerce %v (of type %T) into UnionList", v, v)
	}
}
