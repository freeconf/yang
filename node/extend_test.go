package node

import (
	"testing"

	"github.com/c2stack/c2g/val"
)

func TestExtend(t *testing.T) {
	child := &MyNode{Label: "Bloop"}
	n := &MyNode{
		Label: "Blop",
		OnField: func(r FieldRequest, hnd *ValueHandle) error {
			hnd.Val = val.String("Hello")
			return nil
		},
		OnChild: func(r ChildRequest) (Node, error) {
			return child, nil
		},
	}
	x := Extend{
		Label: "Bleep",
		Node:  n,
		OnField: func(p Node, r FieldRequest, hnd *ValueHandle) error {
			p.Field(r, hnd)
			hnd.Val = val.String(hnd.Val.String() + " World")
			return nil
		},
	}
	var actualValueHnd ValueHandle
	x.Field(FieldRequest{}, &actualValueHnd)
	if actualValueHnd.Val.String() != "Hello World" {
		t.Error(actualValueHnd.Val.String())
	}
	if x.String() != "(Blop) <- Bleep" {
		t.Error(x.String())
	}
	actualChild, _ := x.Child(ChildRequest{})
	if actualChild.String() != "Bloop" {
		t.Error(actualChild.String())
	}
}
