package node

import (
	"errors"
	"fmt"
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
)

type FindTarget struct {
	Path                   PathSlice
	Target                 Selection
	WalkConstraints        *Constraints
	WalkConstraintsHandler *ConstraintHandler
}

func (self *FindTarget) VisitList(r *ListRequest) (next Selection, err error) {
	if !r.First {
		// when we're finding targets, we never iterate more than one item in a list
		return
	}
	if (self.Path.Empty()) && len(self.Path.Head.Key()) == 0 {
		self.Target = r.Selection
		return
	}
	if len(self.Path.Head.Key()) == 0 {
		return next, errors.New("Key required when navigating lists")
	}
	var nextNode Node
	r.Target = self.Path
	r.Key = self.Path.Head.Key()
	nextNode, r.Selection.Path.key, err = r.Selection.Node.Next(*r)
	if err != nil {
		return
	}
	if nextNode == nil {
		return next, c2.NewErrC("List item not found", 404)
	}
	next = r.Selection.selectListItem(nextNode, self.Path.Head.Key())
	if self.Path.Empty() {
		self.Target = r.Selection
	}
	return
}

func (self *FindTarget) VisitNotification(r *NotifyRequest) (Selection, error) {
	sel := r.Selection.selectChild(r.Meta, r.Selection.Node)
	self.Target = sel
	return sel, nil
}

func (self *FindTarget) VisitAction(r *ActionRequest) (Selection, error) {
	actionSel := r.Selection.selectChild(r.Meta, r.Selection.Node)
	self.Target = actionSel
	return actionSel, nil
}

func (self *FindTarget) VisitField(*FieldRequest) (error) {
	// N/A
	return nil
}

func (self *FindTarget) VisitContainer(r *ContainerRequest) (Selection, error) {
	r.Target = self.Path
	childNode, err := r.Selection.Node.Select(*r)
	if err != nil {
		return Selection{}, err
	}
	if childNode == nil {
		msg := fmt.Sprintf("Container not found %s/%s", r.Selection.Path.String(), r.Meta.GetIdent())
		return Selection{}, c2.NewErrC(msg, 404)
	}
	return r.Selection.selectChild(r.Meta, childNode), nil
}

func (self *FindTarget) ContainerIterator(sel Selection, m meta.MetaList) (meta.MetaIterator, error) {
	if _, isChoiceCase := m.(*meta.ChoiceCase); isChoiceCase {
		panic("find target into choice case not expected")
	}
	if self.Path.Empty() {
		self.Target = sel
		return meta.EmptyInterator(0), nil
	}

	self.Path = self.Path.PopHead()
	i := &meta.SingletonIterator{Meta: self.Path.Head.Meta()}
	return i, nil
}
