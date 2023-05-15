package nodeutil

import (
	"context"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"
)

// Extend let's you alter any Node behavior including the nodeutil it creates.
type Extend struct {
	Base        node.Node
	OnNext      func(parent node.Node, r node.ListRequest) (next node.Node, key []val.Value, err error)
	OnChild     func(parent node.Node, r node.ChildRequest) (child node.Node, err error)
	OnField     func(parent node.Node, r node.FieldRequest, hnd *node.ValueHandle) error
	OnChoose    func(parent node.Node, sel *node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error)
	OnAction    func(parent node.Node, r node.ActionRequest) (output node.Node, err error)
	OnNotify    func(parent node.Node, r node.NotifyRequest) (closer node.NotifyCloser, err error)
	OnExtend    func(e *Extend, sel *node.Selection, m meta.HasDefinitions, child node.Node) (node.Node, error)
	OnPeek      func(parent node.Node, sel *node.Selection, consumer interface{}) interface{}
	OnBeginEdit func(parent node.Node, r node.NodeRequest) error
	OnEndEdit   func(parent node.Node, r node.NodeRequest) error
	OnContext   func(parent node.Node, s *node.Selection) context.Context
}

func (e *Extend) Child(r node.ChildRequest) (node.Node, error) {
	var err error
	var child node.Node
	if e.OnChild == nil {
		child, err = e.Base.Child(r)
	} else {
		child, err = e.OnChild(e.Base, r)
	}
	if child == nil || err != nil {
		return child, err
	}
	if e.OnExtend != nil {
		child, err = e.OnExtend(e, r.Selection, r.Meta, child)
	}
	return child, err
}

func (e *Extend) Next(r node.ListRequest) (child node.Node, key []val.Value, err error) {
	if e.OnNext == nil {
		child, key, err = e.Base.Next(r)
	} else {
		child, key, err = e.OnNext(e.Base, r)
	}
	if child == nil || err != nil {
		return
	}
	if e.OnExtend != nil {
		child, err = e.OnExtend(e, r.Selection, r.Meta, child)
	}
	return
}

func (e *Extend) Extend(n node.Node) node.Node {
	extendedChild := *e
	extendedChild.Base = n
	return &extendedChild
}

func (e *Extend) Field(r node.FieldRequest, hnd *node.ValueHandle) error {
	if e.OnField == nil {
		return e.Base.Field(r, hnd)
	} else {
		return e.OnField(e.Base, r, hnd)
	}
}

func (e *Extend) Choose(sel *node.Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
	if e.OnChoose == nil {
		return e.Base.Choose(sel, choice)
	} else {
		return e.OnChoose(e.Base, sel, choice)
	}
}

func (e *Extend) Action(r node.ActionRequest) (output node.Node, err error) {
	if e.OnAction == nil {
		return e.Base.Action(r)
	} else {
		return e.OnAction(e.Base, r)
	}
}

func (e *Extend) Notify(r node.NotifyRequest) (closer node.NotifyCloser, err error) {
	if e.OnNotify == nil {
		return e.Base.Notify(r)
	} else {
		return e.OnNotify(e.Base, r)
	}
}

func (e *Extend) BeginEdit(r node.NodeRequest) error {
	if e.OnBeginEdit == nil {
		return e.Base.BeginEdit(r)
	}
	return e.OnBeginEdit(e.Base, r)
}

func (e *Extend) EndEdit(r node.NodeRequest) error {
	if e.OnEndEdit == nil {
		return e.Base.EndEdit(r)
	}
	return e.OnEndEdit(e.Base, r)
}

func (e *Extend) Context(sel *node.Selection) context.Context {
	if e.OnContext == nil {
		return e.Base.Context(sel)
	}
	return e.OnContext(e.Base, sel)
}

func (e *Extend) Peek(sel *node.Selection, consumer interface{}) interface{} {
	if e.OnPeek == nil {
		return e.Base.Peek(sel, consumer)
	}
	return e.OnPeek(e.Base, sel, consumer)
}
