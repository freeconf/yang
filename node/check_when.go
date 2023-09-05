package node

import (
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
	"github.com/freeconf/yang/xpath"
)

type CheckWhen struct {
}

func (y CheckWhen) CheckContainerPostConstraints(r ChildRequest, s *Selection) (bool, error) {
	return y.check(s, r.Meta)
}

func (y CheckWhen) CheckFieldPreConstraints(r *FieldRequest, hnd *ValueHandle) (bool, error) {
	return y.check(r.Selection, r.Meta)
}

func (y CheckWhen) CheckListPostConstraints(r ListRequest, child *Selection, key []val.Value) (bool, error) {
	return y.check(child, r.Meta)
}

func (y CheckWhen) check(s *Selection, m meta.Meta) (bool, error) {
	if s == nil {
		return true, nil
	}
	if hw, ok := m.(meta.HasWhen); ok {
		if hw.When() != nil {
			xp, err := xpath.Parse(hw.When().Expression())
			if err != nil {
				return false, err
			}
			proceed, err := s.XPredicate(xp)
			return proceed, err
		}
	}
	return true, nil
}
