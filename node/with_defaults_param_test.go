package node

import (
	"testing"
	"github.com/c2stack/c2g/meta"
)

func TestWithDefaultsCheck(t *testing.T) {
	leaf := &meta.Leaf{Ident:"x"}
	r := FieldRequest{
		Meta:leaf,
	}
	dt := meta.NewDataType(leaf, "string")
	leaf.SetDataType(dt)
	dt.SetDefault("a")
	hnd := ValueHandle{
		Val:&Value{Str:"v", Type:dt},
	}
	if proceed, err := WithDefaultsTrim.CheckFieldPostConstraints(r, &hnd, false); err != nil {
		t.Fatal(err)
	} else if !proceed {
		t.Error("expected to not proceed")
	}
	if hnd.Val == nil {
		t.Error("value was reset")
	}
	hnd.Val.Str = "a"
	if proceed, err := WithDefaultsTrim.CheckFieldPostConstraints(r, &hnd, false); err != nil {
		t.Fatal(err)
	} else if !proceed {
		t.Error("expected to not proceed")
	}
	if hnd.Val != nil {
		t.Error("value was not reset")
	}
}
