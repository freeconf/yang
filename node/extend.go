package node

import (
	"fmt"

	"github.com/c2stack/c2g/meta"
)

// Used when you want to alter the response from a Node (and the nodes it creates)
// You can alters, reads, writes, event and event the child nodes it creates.

type Extend struct {
	Label       string
	Node        Node
	OnNext      ExtendNextFunc
	OnChild     ExtendChildFunc
	OnField     ExtendFieldFunc
	OnChoose    ExtendChooseFunc
	OnAction    ExtendActionFunc
	OnNotify    ExtendNotifyFunc
	OnExtend    ExtendFunc
	OnPeek      ExtendPeekFunc
	OnBeginEdit ExtendBeginEditFunc
	OnEndEdit   ExtendEndEditFunc
	OnDelete    ExtendDeleteFunc
}

func (e *Extend) String() string {
	return fmt.Sprintf("(%s) <- %s", e.Node.String(), e.Label)
}

func (e *Extend) Child(r ChildRequest) (Node, error) {
	var err error
	var child Node
	if e.OnChild == nil {
		child, err = e.Node.Child(r)
	} else {
		child, err = e.OnChild(e.Node, r)
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

func (e *Extend) Field(r FieldRequest, hnd *ValueHandle) error {
	if e.OnField == nil {
		return e.Node.Field(r, hnd)
	} else {
		return e.OnField(e.Node, r, hnd)
	}
}

func (e *Extend) Choose(sel Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
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

func (e *Extend) Notify(r NotifyRequest) (closer NotifyCloser, err error) {
	if e.OnNotify == nil {
		return e.Node.Notify(r)
	} else {
		return e.OnNotify(e.Node, r)
	}
}

func (e *Extend) Delete(r NodeRequest) error {
	if e.OnDelete == nil {
		return e.Node.Delete(r)
	}
	return e.OnDelete(e.Node, r)
}

func (e *Extend) BeginEdit(r NodeRequest) error {
	if e.OnBeginEdit == nil {
		return e.Node.BeginEdit(r)
	}
	return e.OnBeginEdit(e.Node, r)
}

func (e *Extend) EndEdit(r NodeRequest) error {
	if e.OnEndEdit == nil {
		return e.Node.EndEdit(r)
	}
	return e.OnEndEdit(e.Node, r)
}

func (e *Extend) Peek(sel Selection) interface{} {
	if e.OnPeek == nil {
		return e.Node.Peek(sel)
	}
	return e.OnPeek(e.Node, sel)
}

type ExtendNextFunc func(parent Node, r ListRequest) (next Node, key []*Value, err error)
type ExtendChildFunc func(parent Node, r ChildRequest) (child Node, err error)
type ExtendFieldFunc func(parent Node, r FieldRequest, hnd *ValueHandle) error
type ExtendChooseFunc func(parent Node, sel Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error)
type ExtendActionFunc func(parent Node, r ActionRequest) (output Node, err error)
type ExtendNotifyFunc func(parent Node, r NotifyRequest) (closer NotifyCloser, err error)
type ExtendFunc func(e *Extend, sel Selection, m meta.MetaList, child Node) (Node, error)
type ExtendPeekFunc func(parent Node, sel Selection) interface{}
type ExtendBeginEditFunc func(parent Node, r NodeRequest) error
type ExtendEndEditFunc func(parent Node, r NodeRequest) error
type ExtendDeleteFunc func(parent Node, r NodeRequest) error
