package nodes

import (
	"testing"

	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"
)

func TestExtend(t *testing.T) {
	child := &Basic{Label: "Bloop"}
	n := &Basic{
		Label: "Blop",
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			hnd.Val = val.String("Hello")
			return nil
		},
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			return child, nil
		},
	}
	x := Extend{
		Label: "Bleep",
		Node:  n,
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			p.Field(r, hnd)
			hnd.Val = val.String(hnd.Val.String() + " World")
			return nil
		},
	}
	var actualValueHnd node.ValueHandle
	x.Field(node.FieldRequest{}, &actualValueHnd)
	if actualValueHnd.Val.String() != "Hello World" {
		t.Error(actualValueHnd.Val.String())
	}
	if x.String() != "(Blop) <- Bleep" {
		t.Error(x.String())
	}
	actualChild, _ := x.Child(node.ChildRequest{})
	if actualChild.String() != "Bloop" {
		t.Error(actualChild.String())
	}
}
