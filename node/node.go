package node

import (
	"fmt"

	"context"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/val"
)

// Node is responsible for reading or writing leafs on a container or list
// getting nodes for other containers, or getting nodes for each items in a list.
// In general you do not want to keep a reference to a node as the data it would be
// pointing to might not be relevent anymore.
//
// You rarely implement this interface, but instead instantiate structs that implement
// this interface like MyNode or Extend
type Node interface {
	fmt.Stringer

	// Child is called to find or create other containers from this container. Request will
	// contain container you will need to create or return another node for
	Child(r ChildRequest) (child Node, err error)

	// Next is called to find or create items in a list.  Request will contain item in
	// list you will need to create or return another node for
	Next(r ListRequest) (next Node, key []val.Value, err error)

	// Field is called to read or write a leaf.
	Field(r FieldRequest, hnd *ValueHandle) error

	// Choose is called when model uses a 'choose' definition and walking logic
	// need to know which part of the model applies to give data.  Only reading
	// existing data models call this method. Writers do not need to implement this
	Choose(sel Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error)

	// Called when this node is begin deleted.  This is called Before Child() is called with
	// delete flag
	Delete(r NodeRequest) error

	// Called when this node is begin edited, or any of it's descendants were edited.  See r.Source for root
	// point of edit. This is called before edit has happened.  This is also called when a delete has happened
	// to any children
	BeginEdit(r NodeRequest) error

	// Called after a node has been edited, or any of it's descendants were edited.  See r.Source for root
	// point of edit. This is called after edit has happened.  This is also called when a delete has happened
	// to any children
	EndEdit(r NodeRequest) error

	// Called when caller wished to run a 'action' or 'rpc' definition.  Input can
	// be found in request if an input is defined.  Output only has to be returned for
	// definitions that declare an output.
	Action(r ActionRequest) (output Node, err error)

	// Called when caller wish to subscribe to events from a node.  Implementations
	// should be sure not to keep references to any other Node or Selection objects as
	// data may have changed.
	Notify(r NotifyRequest) (NotifyCloser, error)

	// Nodes abstract caller from real data, but this let's you peek at the single real object
	// behing this container.  It's up to the implementation to decide what the object is. Use
	// this call with caution.
	Peek(sel Selection, consumer interface{}) interface{}

	// Opportunity to add/change the context for all requests below this node
	Context(sel Selection) context.Context
}

// Used to pass values in/out of calls to Node.Field
type ValueHandle struct {

	// Readers do not set this, Writers will always have a valid value here
	Val val.Value
}

// Most common way to implement Node interface. Only supply the functions for operations your
// data node needs to support.  For example, if
type MyNode struct {

	// Only useful for debugging
	Label string

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

func (s *MyNode) String() string {
	return s.Label
}

func (s *MyNode) Child(r ChildRequest) (Node, error) {
	if s.OnChild == nil {
		return nil, c2.NewErrC(fmt.Sprint("Select not implemented on node ", r.Selection.String()), 501)
	}
	return s.OnChild(r)
}

func (s *MyNode) Next(r ListRequest) (Node, []val.Value, error) {
	if s.OnNext == nil {
		return nil, nil,
			c2.NewErrC(fmt.Sprint("Next not implemented on node ", r.Selection.String()), 501)
	}
	return s.OnNext(r)
}

func (s *MyNode) Field(r FieldRequest, hnd *ValueHandle) error {
	if s.OnField == nil {
		return c2.NewErrC(fmt.Sprint("Field not implemented on node ", r.Selection.String()), 501)
	}
	return s.OnField(r, hnd)
}

func (s *MyNode) Choose(sel Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
	if s.OnChoose == nil {
		return nil,
			c2.NewErrC(fmt.Sprint("Choose not implemented on node ", sel.String()), 501)
	}
	return s.OnChoose(sel, choice)
}

func (s *MyNode) Action(r ActionRequest) (output Node, err error) {
	if s.OnAction == nil {
		return nil,
			c2.NewErrC(fmt.Sprint("Action not implemented on node ", r.Selection.String()), 501)
	}
	return s.OnAction(r)
}

func (s *MyNode) Peek(sel Selection, consumer interface{}) interface{} {
	if s.OnPeek != nil {
		return s.OnPeek(sel, consumer)
	}
	return s.Peekable
}

func (s *MyNode) BeginEdit(r NodeRequest) error {
	if s.OnBeginEdit != nil {
		return s.OnBeginEdit(r)
	}
	return nil
}

func (s *MyNode) EndEdit(r NodeRequest) error {
	if s.OnEndEdit != nil {
		return s.OnEndEdit(r)
	}
	return nil
}

func (s *MyNode) Delete(r NodeRequest) error {
	if s.OnDelete != nil {
		return s.OnDelete(r)
	}
	return nil
}

func (s *MyNode) Context(sel Selection) context.Context {
	if s.OnContext != nil {
		return s.OnContext(sel)
	}
	return sel.Context
}

func (s *MyNode) Notify(r NotifyRequest) (NotifyCloser, error) {
	if s.OnNotify == nil {
		return nil, c2.NewErrC(fmt.Sprint("Notify not implemented on node ", r.Selection.String()), 501)
	}
	return s.OnNotify(r)
}

type NextFunc func(r ListRequest) (next Node, key []val.Value, err error)
type ChildFunc func(r ChildRequest) (child Node, err error)
type FieldFunc func(FieldRequest, *ValueHandle) error
type ChooseFunc func(sel Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error)
type ActionFunc func(ActionRequest) (output Node, err error)
type PeekFunc func(sel Selection, consumer interface{}) interface{}
type NotifyFunc func(r NotifyRequest) (NotifyCloser, error)
type DeleteFunc func(r NodeRequest) error
type BeginEditFunc func(r NodeRequest) error
type EndEditFunc func(r NodeRequest) error
type ContextFunc func(s Selection) context.Context
