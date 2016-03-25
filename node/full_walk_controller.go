package node

import (
	"meta"
)

type ControlledWalk struct {
	Constraints *Constraints
}

func (self *ControlledWalk) VisitField(r *FieldRequest) (v *Value, err error) {
	if self.Constraints != nil {
		r.Constraints = self.Constraints
		if proceed, constraintErr := self.Constraints.CheckFieldPreConstraints(r); !proceed || constraintErr != nil {
			return nil, constraintErr
		}
	}

	if v, err = r.Selection.node.Read(*r); err != nil {
		return nil, err
	}

	if self.Constraints != nil {
		if proceed, constraintErr := self.Constraints.CheckFieldPostConstraints(*r, v); !proceed || constraintErr != nil {
			return nil, constraintErr
		}
	}

	return v, nil
}

func (self *ControlledWalk) VisitAction(r *ActionRequest) (*Selection, error) {
	// Not sure what a full walk would do when hitting an action, so do nothing
	return nil, nil
}

func (self *ControlledWalk) VisitContainer(r *ContainerRequest) (*Selection, error) {
	if self.Constraints != nil {
		r.Constraints = self.Constraints
		if proceed, constraintErr := self.Constraints.CheckContainerPreConstraints(r); !proceed || constraintErr != nil {
			return nil, constraintErr
		}
	}
	childNode, err := r.Selection.node.Select(*r)
	if err != nil {
		return nil, err
	}
	var child *Selection
	if childNode != nil {
		if child, err = r.Selection.SelectChild(r.Meta, childNode), nil; err != nil {
			return nil, err
		}
	}
	if self.Constraints != nil {
		if proceed, constraintErr := self.Constraints.CheckContainerPostConstraints(*r, child); !proceed || constraintErr != nil {
			return nil, constraintErr
		}
	}
	return child, err
}

func (self *ControlledWalk) VisitList(r *ListRequest) (next *Selection, err error) {
	if self.Constraints != nil {
		r.Constraints = self.Constraints
		if proceed, constraintErr := self.Constraints.CheckListPreConstraints(r); !proceed || constraintErr != nil {
			return nil, constraintErr
		}
	}
	var listNode Node
	listNode, r.Selection.path.key, err = r.Selection.node.Next(*r)
	if listNode == nil || err != nil {
		return nil, err
	}
	next = r.Selection.SelectListItem(listNode, r.Selection.path.key)

	if self.Constraints != nil {
		if proceed, constraintErr := self.Constraints.CheckListPostConstraints(*r, next, r.Selection.path.key); !proceed || constraintErr != nil {
			return nil, constraintErr
		}
	}
	return
}

func (self *ControlledWalk) ContainerIterator(sel *Selection, goober meta.MetaList) (meta.MetaIterator, error) {
	return meta.NewMetaListIterator(goober, true), nil
}
