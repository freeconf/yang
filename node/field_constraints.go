package node

import (
	"fmt"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
)

type fieldConstraints struct {
}

func newFieldConstraints() fieldConstraints {
	return fieldConstraints{}
}

func (check fieldConstraints) CheckFieldPreConstraints(r *FieldRequest, hnd *ValueHandle) (bool, error) {
	t := r.Meta.Type()

	switch t.Format() {
	case val.FmtString:
		if valid, err := check.checkString(hnd.Val.String(), t); !valid || err != nil {
			return false, err
		}
	case val.FmtStringList:
		strs := hnd.Val.Value().([]string)
		for _, s := range strs {
			if valid, err := check.checkString(s, t); !valid || err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

func (check fieldConstraints) checkString(s string, t *meta.Type) (bool, error) {
	if valid, err := check.patternCheck(s, t.Patterns()); !valid || err != nil {
		return false, err
	}
	if valid, err := check.lenCheck(s, t.Length()); !valid || err != nil {
		return false, err
	}
	return true, nil
}

func (fieldConstraints) patternCheck(s string, patterns []*meta.Pattern) (bool, error) {
	if len(patterns) == 0 {
		return true, nil
	}
	for _, p := range patterns {
		if p.CheckValue(s) {
			return true, nil
		}
	}
	return false, fmt.Errorf("'%s' did not match any of the required patterns", s)
}

func (fieldConstraints) lenCheck(s string, lengths []*meta.Range) (bool, error) {
	if len(lengths) == 0 {
		return true, nil
	}
	for _, length := range lengths {
		if cmp, err := length.CheckValue(val.Int32(len(s))); cmp || err != nil {
			return cmp, err
		}
	}
	return false, fmt.Errorf("string length outside allowed ranges. %s", s)
}
