package node

import (
	"testing"
)

func TestExtend(t *testing.T) {
	child := &MyNode{Label: "Bloop"}
	n := &MyNode{
		Label: "Blop",
		OnField: func(r FieldRequest, hnd *ValueHandle) error {
			hnd.Val = &Value{Str: "Hello"}
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
			hnd.Val = &Value{Str: hnd.Val.Str + " World"}
			return nil
		},
	}
	var actualValueHnd ValueHandle
	x.Field(FieldRequest{}, &actualValueHnd)
	if actualValueHnd.Val.Str != "Hello World" {
		t.Error(actualValueHnd.Val.Str)
	}
	if x.String() != "(Blop) <- Bleep" {
		t.Error(x.String())
	}
	actualChild, _ := x.Child(ChildRequest{})
	if actualChild.String() != "Bloop" {
		t.Error(actualChild.String())
	}
}
