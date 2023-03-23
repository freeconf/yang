package node

import (
	"fmt"
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
	}
	return val.Conv(typ.Format(), v)
}

func toIdentRef(base *meta.Identity, v interface{}) (val.IdentRef, error) {
	var empty val.IdentRef
	x := fmt.Sprintf("%v", v)
	if colon := strings.IndexRune(x, ':'); colon > 0 {
		x = x[colon+1:]
	}
	ref, found := base.Derived()[x]
	if !found {
		return empty, fmt.Errorf("could not find identity ref for %T:'%s' in '%s'", v, x, base.Ident())
	}
	return val.IdentRef{Base: base.Ident(), Label: ref.Ident()}, nil
}

func toIdentRefList(base *meta.Identity, v interface{}) (val.IdentRefList, error) {
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
	return nil, fmt.Errorf("could not coerse '%v' into identref list", v)
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
	return nil, fmt.Errorf("could not coerse '%v' into enum list", v)
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
	return val.Enum{}, fmt.Errorf("could not coerse '%v' into enum", v)
}
