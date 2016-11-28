package node

import "github.com/c2stack/c2g/meta"

type ControlledWalk struct {
	Constraints *Constraints
	Handler     *ConstraintHandler
}

func (self *ControlledWalk) VisitField(r *FieldRequest) (err error) {
	var hnd ValueHandle

	// Constraints are now handled in Selection

	if err = r.Selection.Node.Field(*r, &hnd); err != nil {
		return err
	}

	return nil
}

func (self *ControlledWalk) VisitAction(r *ActionRequest) (Selection, error) {
	// Not sure what a full walk would do when hitting an action, so do nothing
	return Selection{}, nil
}

func (self *ControlledWalk) VisitNotification(r *NotifyRequest) (Selection, error) {
	// Not sure what a full walk would do when hitting an action, so do nothing
	return Selection{}, nil
}

func (self *ControlledWalk) VisitContainer(r *ContainerRequest) (Selection, error) {
	if self.Constraints != nil {
		r.Constraints = self.Constraints
		r.ConstraintsHandler = self.Handler
		if proceed, constraintErr := self.Constraints.CheckContainerPreConstraints(r, false); !proceed || constraintErr != nil {
			return Selection{}, constraintErr
		}
	}
	childNode, err := r.Selection.Node.Select(*r)
	if err != nil {
		return Selection{}, err
	}
	var child Selection
	if childNode != nil {
		if child, err = r.Selection.selectChild(r.Meta, childNode), nil; err != nil {
			return Selection{}, err
		}
	}
	if self.Constraints != nil {
		if proceed, constraintErr := self.Constraints.CheckContainerPostConstraints(*r, child, false); !proceed || constraintErr != nil {
			return Selection{}, constraintErr
		}
	}
	return child, err
}

func (self *ControlledWalk) VisitList(r *ListRequest) (next Selection, err error) {
	var proceed bool
	if self.Constraints != nil {
		r.Constraints = self.Constraints
		r.ConstraintsHandler = self.Handler
		if proceed, err = self.Constraints.CheckListPreConstraints(r, false); !proceed || err != nil {
			return
		}
	}
	var listNode Node
	listNode, r.Selection.Path.key, err = r.Selection.Node.Next(*r)
	if listNode == nil || err != nil {
		return
	}
	next = r.Selection.selectListItem(listNode, r.Selection.Path.key)

	if self.Constraints != nil {
		if proceed, err = self.Constraints.CheckListPostConstraints(*r, next, r.Selection.Path.key, false); !proceed || err != nil {
			return
		}
	}
	return
}

func (self *ControlledWalk) ContainerIterator(sel Selection, m meta.MetaList) (meta.MetaIterator, error) {
	return meta.NewMetaListIterator(m, true), nil
}
