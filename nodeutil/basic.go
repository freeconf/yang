package nodeutil

import (
	"context"
	"fmt"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"
)

type NextFunc func(r node.ListRequest) (next node.Node, key []val.Value, err error)
type NextItemFunc func(r node.ListRequest) BasicNextItem
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

// Basic is the stubs every method of node.Node interface. Only supply the functions for operations your
// data node needs to support.
type Basic struct {

	// What to return on calls to Peek().  Doesn't have to be valid
	Peekable interface{}

	// Only if node is a list AND you don't implement OnNextItem
	OnNext NextFunc

	// Only if node is a list AND you don't implement OnNext
	OnNextItem NextItemFunc

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

// BasicNextItem is used to organize the function calls around individual list items
// and is returned from Basic.NextItem.  You must implement all functions except DeleteByKey
// and New if list is not editable
type BasicNextItem struct {

	// New when requested to create a new list item.  If you want to wait until all
	// fields are set before adding new item to list, then you can do that in OnEndEdit
	New func() error

	// Find item in list by it's key(s).  You will return found item in Node implementation
	GetByKey func() error

	// Find item in list by it's row position in list.  You will return found item in Node implementation
	GetByRow func() ([]val.Value, error)

	// Find item in list by it's row position in list.  You will return found item in Node implementation
	Node func() (node.Node, error)

	// Remove item. OnEndEdit will also be called if you want to finalize the delete
	DeleteByKey func() error
}

func (n *Basic) nextItem(r node.ListRequest) (BasicNextItem, error) {
	var err error
	item := n.OnNextItem(r)
	if item.GetByKey == nil {
		err = fmt.Errorf("func GetByKey not implemented on node %s ", r.Selection.Path)
	}
	if item.GetByRow == nil {
		err = fmt.Errorf("func GetByRow not implemented on node %s ", r.Selection.Path)
	}
	if item.Node == nil {
		err = fmt.Errorf("func Node not implemented on node %s ", r.Selection.Path)
	}
	return item, err
}

func (n *Basic) Next(r node.ListRequest) (node.Node, []val.Value, error) {
	if n.OnNext != nil {
		return n.OnNext(r)
	}
	if n.OnNextItem == nil {
		return nil, nil,
			fmt.Errorf("neither OnNext nor OnNextItem are implemented on node %s ", r.Selection.Path)
	}
	if len(r.Key) > 0 {
		if r.New {
			item, err := n.nextItem(r)
			if err != nil {
				return nil, nil, err
			}
			if item.New == nil {
				return nil, nil, fmt.Errorf("func New not implemented on node %s ", r.Selection.Path)
			}
			if err = item.New(); err != nil {
				return nil, nil, err
			}
			n, err := item.Node()
			return n, r.Key, err
		} else if r.Delete {
			item, err := n.nextItem(r)
			if err != nil {
				return nil, nil, err
			}
			if item.DeleteByKey == nil {
				return nil, nil, fmt.Errorf("func DeleteByKey not implemented on node %s ", r.Selection.Path)
			}
			return nil, nil, item.DeleteByKey()
		}
		item, err := n.nextItem(r)
		if err != nil {
			return nil, nil, err
		}
		if err = item.GetByKey(); err != nil {
			return nil, nil, err
		}
		n, err := item.Node()
		return n, r.Key, err
	}

	item, err := n.nextItem(r)
	if err != nil {
		return nil, nil, err
	}
	key, err := item.GetByRow()
	if err != nil {
		return nil, nil, err
	}
	child, err := item.Node()
	return child, key, err
}
