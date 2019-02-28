package nodes

import (
	"context"
	"fmt"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/meta"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/val"
)

// Most common way to implement Node interface. Only supply the functions for operations your
// data node needs to support.  For example, if
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

	OnContext ContextFunc

	OnDelete    DeleteFunc
	OnBeginEdit BeginEditFunc
	OnEndEdit   EndEditFunc
}

func (s *Basic) Child(r node.ChildRequest) (node.Node, error) {
	if s.OnChild == nil {
		return nil, c2.NewErrC(fmt.Sprintf("OnChild not implemented for %s.%s", r.Selection.Path.String(), r.Meta.Ident()), 501)
	}
	return s.OnChild(r)
}

func (s *Basic) Next(r node.ListRequest) (node.Node, []val.Value, error) {
	if s.OnNext == nil {
		return nil, nil,
			c2.NewErrC(fmt.Sprint("OnNext not implemented on node ", r.Selection.Path.String()), 501)
	}
	return s.OnNext(r)
}

func (s *Basic) Field(r node.FieldRequest, hnd *node.ValueHandle) error {
	if s.OnField == nil {
		return c2.NewErrC(fmt.Sprintf("OnField not implemented on node for %s.%s", r.Selection.Path.String(), r.Meta.Ident()), 501)
	}
	return s.OnField(r, hnd)
}

func (s *Basic) Choose(sel node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
	if s.OnChoose == nil {
		return nil,
			c2.NewErrC(fmt.Sprintf("OnChoose not implemented for %s.%s", sel.Path.String(), choice.Ident()), 501)
	}
	return s.OnChoose(sel, choice)
}

func (s *Basic) Action(r node.ActionRequest) (output node.Node, err error) {
	if s.OnAction == nil {
		return nil,
			c2.NewErrC(fmt.Sprintf("OnAction not implemented for %s.%s", r.Selection.Path.String(), r.Meta.Ident()), 501)
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
		return nil, c2.NewErrC(fmt.Sprint("Notify not implemented on node ", r.Selection.Path.String()), 501)
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
