package node

import (
	"github.com/freeconf/c2g/meta"
	"github.com/freeconf/c2g/xpath"
)

type CheckWhen struct {
}

func (y CheckWhen) CheckContainerPostConstraints(r ChildRequest, s Selection) (bool, error) {
	return y.check(s, r.Meta)
}

func (y CheckWhen) CheckFieldPreConstraints(r *FieldRequest, hnd *ValueHandle) (bool, error) {
	return y.check(r.Selection, r.Meta)
}

func (y CheckWhen) CheckListPreConstraints(r *ListRequest) (bool, error) {
	return y.check(r.Selection, r.Meta)
}

func (y CheckWhen) check(s Selection, m meta.Meta) (bool, error) {
	if s.IsNil() {
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
