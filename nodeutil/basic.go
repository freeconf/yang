package nodeutil

import (
	"context"
	"fmt"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"
)

// Basic stubs every method of node.Node interface so you only have to supply the functions
// for operations your data node needs to support.
type Basic struct {

	// What to return on calls to Peek().  Peek let's you return underlying objects
	// behind a node and frankly breaks encapsulation so you need not set anything here
	Peekable interface{}

	// Only if node is a list. Return a second data structure that breaks down each request
	// that can be make on each item in a list.
	//
	// If you impement this, you should not implement OnNext
	OnNextItem func(r node.ListRequest) BasicNextItem

	// Only if node is a list AND you don't implement OnNextItem
	//
	// If you impement this, you should not implement OnNextItem
	OnNext func(r node.ListRequest) (next node.Node, key []val.Value, err error)

	// Only if there are other containers or lists defined
	OnChild func(r node.ChildRequest) (child node.Node, err error)

	// Only if you have leafs defined
	OnField func(node.FieldRequest, *node.ValueHandle) error

	// Only if there one or more 'choice' definitions on a list or container and data is used
	// on a reading mode
	OnChoose func(sel *node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error)

	// Only if there is one or more 'rpc' or 'action' defined in a model that could be
	// called.
	OnAction func(node.ActionRequest) (output node.Node, err error)

	// Only if there is one or more 'notification' defined in a model that could be subscribed to
	OnNotify func(r node.NotifyRequest) (node.NotifyCloser, error)

	// Peekable is often enough, but this always you to return an object dynamically
	OnPeek func(sel *node.Selection, consumer interface{}) interface{}

	// OnContext default implementation does nothing
	OnContext func(s *node.Selection) context.Context

	// OnBeginEdit default implementation does nothing
	OnBeginEdit func(r node.NodeRequest) error

	// OnEndEdit default implementation does nothing
	OnEndEdit func(r node.NodeRequest) error

	// OnRelease default implementation does nothing
	OnRelease func(s *node.Selection)
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

func (s *Basic) Choose(sel *node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
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

func (s *Basic) Peek(sel *node.Selection, consumer interface{}) interface{} {
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

func (s *Basic) Context(sel *node.Selection) context.Context {
	if s.OnContext != nil {
		return s.OnContext(sel)
	}
	return sel.Context
}

func (s *Basic) Release(sel *node.Selection) {
	if s.OnRelease != nil {
		s.OnRelease(sel)
	}
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
