package nodeutil

import (
	"context"
	"fmt"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"
)

// Basic is the stubs every method of node.Node interface. Only supply the functions for operations your
// data node needs to support.
type Basic struct {

	// What to return on calls to Peek().  Doesn't have to be valid
	Peekable interface{}

	// Only if node is a list
	OnNext NextFunc

	// Only if there are other containers or lists defined
	OnChild ChildFunc

	// Only if you have leafs defined
	OnField FieldFunc

	// Only if there one or more 'choice' definitions on a list or container and data is used
	// on a reading mode
	OnChoose ChooseFunc

	// Only if there is one or more 'rpc' or 'action' defined in a model that could be
	// called.
	OnAction ActionFunc

	// Only if there is one or more 'notification' defined in a model that could be subscribed to
	OnNotify NotifyFunc

	// Peekable is often enough, but this always you to return an object dynamically
	OnPeek PeekFunc

	// OnContext default implementation does nothing
	OnContext ContextFunc

	// OnDelete default implementation does nothing
	OnDelete DeleteFunc

	// OnBeginEdit default implementation does nothing
	OnBeginEdit BeginEditFunc

	// OnEndEdit default implementation does nothing
	OnEndEdit EndEditFunc
}

func (s *Basic) Child(r node.ChildRequest) (node.Node, error) {
	if s.OnChild == nil {
		return nil, fmt.Errorf("OnChild not implemented for %s.%s", r.Selection.Path, r.Meta.Ident())
	}
	return s.OnChild(r)
}

func (s *Basic) Next(r node.ListRequest) (node.Node, []val.Value, error) {
	if s.OnNext == nil {
		return nil, nil,
			fmt.Errorf("OnNext not implemented on node %s ", r.Selection.Path)
	}
	return s.OnNext(r)
}

func (s *Basic) Field(r node.FieldRequest, hnd *node.ValueHandle) error {
	if s.OnField == nil {
		return fmt.Errorf("OnField not implemented on node for %s.%s", r.Selection.Path, r.Meta.Ident())
	}
	return s.OnField(r, hnd)
}

func (s *Basic) Choose(sel node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
	if s.OnChoose == nil {
		return nil, fmt.Errorf("OnChoose not implemented for %s.%s", sel.Path, choice.Ident())
	}
	return s.OnChoose(sel, choice)
}

func (s *Basic) Action(r node.ActionRequest) (output node.Node, err error) {
	if s.OnAction == nil {
		return nil, fmt.Errorf("OnAction not implemented for %s.%s", r.Selection.Path, r.Meta.Ident())
	}
	return s.OnAction(r)
}

func (s *Basic) Peek(sel node.Selection, consumer interface{}) interface{} {
	if s.OnPeek != nil {
		return s.OnPeek(sel, consumer)
	}
	return s.Peekable
}

func (s *Basic) BeginEdit(r node.NodeRequest) error {
	if s.OnBeginEdit != nil {
		return s.OnBeginEdit(r)
	}
	return nil
}

func (s *Basic) EndEdit(r node.NodeRequest) error {
	if s.OnEndEdit != nil {
		return s.OnEndEdit(r)
	}
	return nil
}

func (s *Basic) Delete(r node.NodeRequest) error {
	if s.OnDelete != nil {
		return s.OnDelete(r)
	}
	return nil
}

func (s *Basic) Context(sel node.Selection) context.Context {
	if s.OnContext != nil {
		return s.OnContext(sel)
	}
	return sel.Context
}

func (s *Basic) Notify(r node.NotifyRequest) (node.NotifyCloser, error) {
	if s.OnNotify == nil {
		return nil, fmt.Errorf("Notify not implemented on node %s", r.Selection.Path)
	}
	return s.OnNotify(r)
}

type NextFunc func(r node.ListRequest) (next node.Node, key []val.Value, err error)
type ChildFunc func(r node.ChildRequest) (child node.Node, err error)
type FieldFunc func(node.FieldRequest, *node.ValueHandle) error
type ChooseFunc func(sel node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error)
type ActionFunc func(node.ActionRequest) (output node.Node, err error)
type PeekFunc func(sel node.Selection, consumer interface{}) interface{}
type NotifyFunc func(r node.NotifyRequest) (node.NotifyCloser, error)
type DeleteFunc func(r node.NodeRequest) error
type BeginEditFunc func(r node.NodeRequest) error
type EndEditFunc func(r node.NodeRequest) error
type ContextFunc func(s node.Selection) context.Context
