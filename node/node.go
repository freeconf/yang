package node

import (
	"fmt"
	"github.com/blitter/meta"
	"github.com/blitter/blit"
)

type Node interface {
	fmt.Stringer
	Select(r ContainerRequest) (child Node, err error)
	Next(r ListRequest) (next Node, key []*Value, err error)
	Read(r FieldRequest) (*Value, error)
	Write(r FieldRequest, val *Value) error
	Choose(sel *Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error)
	Event(sel *Selection, e Event) error
	Action(r ActionRequest) (output Node, err error)
	Peek(sel *Selection, peekId string) interface{}
}

// A way to direct changes to another node to enable CopyOnWrite or other persistable options
type ChangeAwareNode interface {
	DirectChanges(config Node)
	Changes() Node
}

type MyNode struct {
	Label        string
	Peekables    map[string]interface{}
	ChangeAccess Node
	OnNext       NextFunc
	OnSelect     SelectFunc
	OnRead       ReadFunc
	OnWrite      WriteFunc
	OnChoose     ChooseFunc
	OnAction     ActionFunc
	OnEvent      EventFunc
	OnPeek       PeekFunc
	Resource     meta.Resource
}

func (n *MyNode) DirectChanges(changeNode Node) {
	n.ChangeAccess = changeNode
}

func (n *MyNode) Changes() Node {
	// If there's a change interceptor set, use it otherwise
	// changes go directly back to node
	if n.ChangeAccess != nil {
		return n.ChangeAccess
	}
	return n
}

func (s *MyNode) String() string {
	return s.Label
}

func (s *MyNode) Close() (err error) {
	if s.Resource != nil {
		err = s.Resource.Close()
		s.Resource = nil
	}
	return
}

func (s *MyNode) Select(r ContainerRequest) (Node, error) {
	if s.OnSelect == nil {
		return nil, blit.NewErrC(fmt.Sprint("Select not implemented on node ", r.Selection.String()), 501)
	}
	return s.OnSelect(r)
}

func (s *MyNode) Next(r ListRequest) (Node, []*Value, error) {
	if s.OnNext == nil {
		return nil, nil,
			blit.NewErrC(fmt.Sprint("Next not implemented on node ", r.Selection.String()), 501)
	}
	return s.OnNext(r)
}

func (s *MyNode) Read(r FieldRequest) (*Value, error) {
	if s.OnRead == nil {
		return nil,
		blit.NewErrC(fmt.Sprint("Read not implemented on node ", r.Selection.String()), 501)
	}
	return s.OnRead(r)
}

func (s *MyNode) Write(r FieldRequest, val *Value) error {
	if s.OnWrite == nil {
		return blit.NewErrC(fmt.Sprint("Write not implemented on node ", r.Selection.String()), 501)
	}
	return s.OnWrite(r, val)
}

func (s *MyNode) Choose(sel *Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
	if s.OnChoose == nil {
		return nil,
			blit.NewErrC(fmt.Sprint("Choose not implemented on node ", sel.String()), 501)
	}
	return s.OnChoose(sel, choice)
}

func (s *MyNode) Action(r ActionRequest) (output Node, err error) {
	if s.OnAction == nil {
		return nil,
		blit.NewErrC(fmt.Sprint("Action not implemented on node ", r.Selection.String()), 501)
	}
	return s.OnAction(r)
}

func (s *MyNode) Event(sel *Selection, e Event) (err error) {
	if s.OnEvent != nil {
		return s.OnEvent(sel, e)
	}
	return nil
}

func (s *MyNode) Peek(sel *Selection, peekId string) interface{} {
	if s.OnPeek != nil {
		return s.OnPeek(sel, peekId)
	}
	return s.Peekables[peekId]
}

// Useful when you want to return an error from Data.Node().  Any call to get data
// will return same error
//
// func (d *MyData) Node {
//    return ErrorNode(errors.New("bang"))
// }
type ErrorNode struct {
	Err error
}

func (e ErrorNode) Error() string {
	return e.Err.Error()
}

func (e ErrorNode) String() string {
	return e.Error()
}

func (e ErrorNode) Select(r ContainerRequest) (Node, error) {
	return nil, e.Err
}

func (e ErrorNode) Next(ListRequest) (Node, []*Value, error) {
	return nil, nil, e.Err
}

func (e ErrorNode) Read(FieldRequest) (*Value, error) {
	return nil, e.Err
}

func (e ErrorNode) Write(FieldRequest, *Value) error {
	return e.Err
}

func (e ErrorNode) Choose(*Selection, *meta.Choice) (*meta.ChoiceCase, error) {
	return nil, e.Err
}

func (e ErrorNode) Event(*Selection, Event) error {
	return e.Err
}

func (e ErrorNode) Action(ActionRequest) (Node, error) {
	return nil, e.Err
}

func (e ErrorNode) Peek(sel *Selection, peekId string) interface{} {
	return nil
}

type NextFunc func(r ListRequest) (next Node, key []*Value, err error)
type SelectFunc func(r ContainerRequest) (child Node, err error)
type ReadFunc func(FieldRequest) (*Value, error)
type WriteFunc func(FieldRequest, *Value) error
type ChooseFunc func(sel *Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error)
type ActionFunc func(ActionRequest) (output Node, err error)
type EventFunc func(sel *Selection, e Event) error
type PeekFunc func(sel *Selection, peekId string) interface{}
