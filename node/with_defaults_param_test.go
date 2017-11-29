package node_test

import (
	"testing"

	"github.com/freeconf/c2g/meta"
	"github.com/freeconf/c2g/node"
	"github.com/freeconf/c2g/val"
)

func TestWithDefaultsCheck(t *testing.T) {
	m := meta.NewModule("m", nil)
	leaf := meta.NewLeaf(m, "x")
	meta.Set(m, leaf)
	r := node.FieldRequest{
		Meta: leaf,
	}
	dt := meta.NewDataType(leaf, "string")
	meta.Set(leaf, dt)
	meta.Set(leaf, meta.SetDefault{Value: "a"})
	if err := meta.Validate(m); err != nil {
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
