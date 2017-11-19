package node

import (
	"context"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/val"
)

// Node is responsible for reading or writing leafs on a container or list
// getting nodes for other containers, or getting nodes for each items in a list.
// In general you do not want to keep a reference to a node as the data it would be
// pointing to might not be relevent anymore.
//
// You rarely implement this interface, but instead instantiate structs that implement
// this interface like Basic or Extend
type Node interface {

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
