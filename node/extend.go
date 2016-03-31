package node
import (
	"github.com/c2g/meta"
	"fmt"
)

// Used when you want to alter the response from a Node (and the nodes it creates)
// You can alters, reads, writes, event and event the child nodes it creates.

type Extend struct {
	Label    string
	Node     Node
	OnNext   ExtendNextFunc
	OnSelect ExtendSelectFunc
	OnRead   ExtendReadFunc
	OnWrite  ExtendWriteFunc
	OnChoose ExtendChooseFunc
	OnAction ExtendActionFunc
	OnEvent  ExtendEventFunc
	OnExtend ExtendFunc
	OnPeek ExtendPeekFunc
}

func (e *Extend) String() string {
	return fmt.Sprintf("(%s) <- %s", e.Node.String(), e.Label)
}

func (e *Extend) Select(r ContainerRequest) (Node, error) {
	var err error
	var child Node
	if e.OnSelect == nil {
		child, err = e.Node.Select(r)
	} else {
		child, err = e.OnSelect(e.Node, r)
	}
	if child == nil || err != nil {
		return child, err
	}
	if e.OnExtend != nil {
		child, err = e.OnExtend(e, r.Selection, r.Meta, child)
	}
	return child, err
}

func (e *Extend) Next(r ListRequest) (child Node, key []*Value, err error) {
	if e.OnNext == nil {
		child, key, err = e.Node.Next(r)
	} else {
		child, key, err = e.OnNext(e.Node, r)
	}
	if child == nil || err != nil {
		return
	}
	if e.OnExtend != nil {
		child, err = e.OnExtend(e, r.Selection, r.Meta, child)
	}
	return
}

func (e *Extend) Extend(n Node) Node {
	extendedChild := *e
	extendedChild.Node = n
	return &extendedChild
}

func (e *Extend) Read(r FieldRequest) (*Value, error) {
	if e.OnRead == nil {
		return e.Node.Read(r)
	} else {
		return e.OnRead(e.Node, r)
	}
}

func (e *Extend) Write(r FieldRequest, v *Value) (error) {
	if e.OnWrite == nil {
		return e.Node.Write(r, v)
	} else {
		return e.OnWrite(e.Node, r, v)
	}
}

func (e *Extend) Choose(sel *Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
	if e.OnChoose == nil {
		return e.Node.Choose(sel, choice)
	} else {
		return e.OnChoose(e.Node, sel, choice)
	}
}

func (e *Extend) Action(r ActionRequest) (output Node, err error) {
	if e.OnAction == nil {
		return e.Node.Action(r)
	} else {
		return e.OnAction(e.Node, r)
	}
}

func (e *Extend) Event(sel *Selection, event Event) (err error) {
	if e.OnEvent == nil {
		return e.Node.Event(sel, event)
	} else {
		return e.OnEvent(e.Node, sel, event)
	}
}

func (e *Extend) Peek(sel *Selection, peekId string) interface{} {
	if e.OnPeek == nil {
		if found := e.Node.Peek(sel, peekId); found != nil {
			return found
		}
	}
	return e.OnPeek(e.Node, sel, peekId)
}

type ExtendNextFunc func(parent Node, r ListRequest) (next Node, key []*Value, err error)
type ExtendSelectFunc func(parent Node, r ContainerRequest) (child Node, err error)
type ExtendReadFunc func(parent Node, r FieldRequest) (*Value, error)
type ExtendWriteFunc func(parent Node, r FieldRequest, val *Value) error
type ExtendChooseFunc func(parent Node, sel *Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error)
type ExtendActionFunc func(parent Node, r ActionRequest) (output Node, err error)
type ExtendEventFunc func(parent Node, sel *Selection, e Event) error
type ExtendFunc func(e *Extend, sel *Selection, m meta.MetaList, child Node) (Node, error)
type ExtendPeekFunc func(parent Node, sel *Selection, peekId string) interface{}
