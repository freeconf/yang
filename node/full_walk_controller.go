package node

import (
	"github.com/c2g/meta"
)

type ControlledWalk struct {
}

func (self *ControlledWalk) VisitField(r *FieldRequest) (v *Value, err error) {
	if r.Context.Constraints() != nil {
		if proceed, constraintErr := r.Context.Constraints().CheckFieldPreConstraints(r, false); !proceed || constraintErr != nil {
			return nil, constraintErr
		}
	}

	if v, err = r.Selection.node.Read(*r); err != nil {
		return nil, err
	}

	if r.Context.Constraints != nil {
		if proceed, constraintErr := r.Context.Constraints().CheckFieldPostConstraints(*r, v, false); !proceed || constraintErr != nil {
			return nil, constraintErr
		}
	}

	return v, nil
}

func (self *ControlledWalk) VisitAction(r *ActionRequest) (*Selection, error) {
	// Not sure what a full walk would do when hitting an action, so do nothing
	return nil, nil
}

func (self *ControlledWalk) VisitNotification(r *NotifyRequest) (*Selection, error) {
	// Not sure what a full walk would do when hitting an action, so do nothing
	return nil, nil
}

func (self *ControlledWalk) VisitContainer(r *ContainerRequest) (*Selection, error) {
	if r.Context.Constraints != nil {
		if proceed, constraintErr := r.Context.Constraints().CheckContainerPreConstraints(r, false); !proceed || constraintErr != nil {
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
	if r.Context.Constraints != nil {
		if proceed, constraintErr := r.Context.Constraints().CheckContainerPostConstraints(*r, child, false); !proceed || constraintErr != nil {
			return nil, constraintErr
		}
	}
	return child, err
}

func (self *ControlledWalk) VisitList(r *ListRequest) (next *Selection, err error) {
	if r.Context.Constraints != nil {
		if proceed, constraintErr := r.Context.Constraints().CheckListPreConstraints(r, false); !proceed || constraintErr != nil {
			return nil, constraintErr
		}
	}
	var listNode Node
	listNode, r.Selection.path.key, err = r.Selection.node.Next(*r)
	if listNode == nil || err != nil {
		return nil, err
	}
	next = r.Selection.SelectListItem(listNode, r.Selection.path.key)

	if r.Context.Constraints != nil {
		if proceed, constraintErr := r.Context.Constraints().CheckListPostConstraints(*r, next, r.Selection.path.key, false); !proceed || constraintErr != nil {
			return nil, constraintErr
		}
	}
	return
}

func (self *ControlledWalk) ContainerIterator(sel *Selection, m meta.MetaList) (meta.MetaIterator, error) {
	return meta.NewMetaListIterator(m, true), nil
}
