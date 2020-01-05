package node_test

import (
	"testing"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"
)

func TestWithDefaultsCheck(t *testing.T) {
	b := &meta.Builder{}
	m := b.Module("m", nil)
	leaf := b.Leaf(m, "x")
	r := node.FieldRequest{
		Meta: leaf,
	}
	b.Type(leaf, "string")
	b.Default(leaf, "a")
	if b.LastErr != nil {
		t.Error(b.LastErr)
	}
	if err := meta.Compile(m); err != nil {
		t.Error(err)
	}
	hnd := node.ValueHandle{
		Val: val.String("v"),
	}
	if proceed, err := node.WithDefaultsTrim.CheckFieldPostConstraints(r, &hnd); err != nil {
		t.Fatal(err)
	} else if !proceed {
		t.Error("expected to not proceed")
	}
	if hnd.Val == nil {
		t.Error("value was reset")
	}
	hnd.Val = val.String("a")
	if proceed, err := node.WithDefaultsTrim.CheckFieldPostConstraints(r, &hnd); err != nil {
		t.Fatal(err)
	} else if !proceed {
		t.Error("expected to not proceed")
	}
	if hnd.Val != nil {
		t.Error("value was not reset")
	}
}
