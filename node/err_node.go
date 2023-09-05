package node

import (
	"context"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
)

// Useful when you want to return an error from Data.Node().  Any call to get data
// will return same error
//
//	func (d *MyData) Node {
//	   return ErrorNode(errors.New("bang"))
//	}
type ErrorNode struct {
	Err error
}

func (e ErrorNode) Error() string {
	return e.Err.Error()
}

func (e ErrorNode) String() string {
	return e.Error()
}

func (e ErrorNode) Child(r ChildRequest) (Node, error) {
	return nil, e.Err
}

func (e ErrorNode) Next(ListRequest) (Node, []val.Value, error) {
	return nil, nil, e.Err
}

func (e ErrorNode) Field(FieldRequest, *ValueHandle) error {
	return e.Err
}

func (e ErrorNode) Choose(*Selection, *meta.Choice) (*meta.ChoiceCase, error) {
	return nil, e.Err
}

func (e ErrorNode) Notify(NotifyRequest) (NotifyCloser, error) {
	return nil, e.Err
}

func (e ErrorNode) Action(ActionRequest) (Node, error) {
	return nil, e.Err
}

func (e ErrorNode) Peek(sel *Selection, consumer interface{}) interface{} {
	return nil
}

func (e ErrorNode) BeginEdit(r NodeRequest) error {
	return e.Err
}

func (e ErrorNode) EndEdit(r NodeRequest) error {
	return e.Err
}

func (e ErrorNode) Delete(r NodeRequest) error {
	return e.Err
}

func (ErrorNode) Context(s *Selection) context.Context {
	return s.Context
}

func (ErrorNode) Release(s *Selection) {}
