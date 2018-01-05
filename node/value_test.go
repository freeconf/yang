package node_test

import (
	"testing"

	"github.com/freeconf/gconf/meta"
	"github.com/freeconf/gconf/node"
)

func TestCoerseValue(t *testing.T) {
	m := meta.NewModule("x", nil)
	l := meta.NewLeaf(m, "l")
	meta.Set(m, l)
	dt := meta.NewType("int32")
	meta.Set(l, dt)
	if err := meta.Validate(m); err != nil {
		t.Error(err)
	}
	v, err := node.NewValue(dt, 35)
	if err != nil {
		t.Error(err)
	} else if v.Value().(int) != 35 {
		t.Error("Coersion error")
	}
}
