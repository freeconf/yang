package node

import (
	"fmt"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
)

type fieldConstraints struct {
}

func (check fieldConstraints) CheckFieldPreConstraints(r *FieldRequest, hnd *ValueHandle) (bool, error) {
	t := r.Meta.Type()
	if hnd.Val == nil {
		return true, nil
	}

	switch t.Format() {
	case val.FmtString:
		if err := check.checkString(hnd.Val.String(), t); err != nil {
			return false, err
		}
	case val.FmtStringList:
		strs := hnd.Val.Value().([]string)
		for _, s := range strs {
			if err := check.checkString(s, t); err != nil {
				return false, err
			}
		}
	}
	if t.Format().IsNumeric() {
		if err := check.checkRange(hnd.Val, t); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (check fieldConstraints) checkString(s string, t *meta.Type) error {
	if err := check.patternCheck(s, t.Patterns()); err != nil {
		return err
	}
	if err := check.lenCheck(s, t.Length()); err != nil {
		return err
	}
	return nil
}

func (check fieldConstraints) checkRange(v val.Value, t *meta.Type) error {
	if len(t.Range()) == 0 {
		return nil
	}
	for _, r := range t.Range() {
		if err := r.CheckValue(v); err == nil {
			return nil
		}
	}
	return fmt.Errorf("'%s' did not match any of the required ranges", v)
}

func (fieldConstraints) patternCheck(s string, patterns []*meta.Pattern) error {
	if len(patterns) == 0 {
		return nil
	}
	for _, p := range patterns {
		if p.CheckValue(s) {
			return nil
		}
	}
	return fmt.Errorf("'%s' did not match any of the required patterns", s)
}

func (fieldConstraints) lenCheck(s string, lengths []*meta.Range) error {
	if len(lengths) == 0 {
		return nil
	}
	for _, length := range lengths {
		if err := length.CheckValue(val.Int32(len(s))); err == nil {
			return nil
		}
	}
	return fmt.Errorf("string length outside allowed ranges. %s", s)
}
