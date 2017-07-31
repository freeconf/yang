package node

import (
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/val"
)

func TestWithDefaultsCheck(t *testing.T) {
	leaf := &meta.Leaf{Ident: "x"}
	r := FieldRequest{
		Meta: leaf,
	}
	dt := meta.NewDataType(leaf, "string")
	leaf.SetDataType(dt)
	dt.SetDefault("a")
	hnd := ValueHandle{
		Val: val.String("v"),
	}
	if proceed, err := WithDefaultsTrim.CheckFieldPostConstraints(r, &hnd); err != nil {
		t.Fatal(err)
	} else if !proceed {
		t.Error("expected to not proceed")
	}
	if hnd.Val == nil {
		t.Error("value was reset")
	}
	hnd.Val = val.String("a")
	if proceed, err := WithDefaultsTrim.CheckFieldPostConstraints(r, &hnd); err != nil {
		t.Fatal(err)
	} else if !proceed {
		t.Error("expected to not proceed")
	}
	if hnd.Val != nil {
		t.Error("value was not reset")
	}
}
