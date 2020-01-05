package node_test

import (
	"testing"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
)

func TestCoerseValue(t *testing.T) {
	b := &meta.Builder{}
	m := b.Module("x", nil)
	l := b.Leaf(m, "l")
	dt := b.Type(l, "int32")
	if err := meta.Compile(m); err != nil {
		t.Error(err)
	}
	v, err := node.NewValue(dt, 35)
	if err != nil {
		t.Error(err)
	} else if v.Value().(int) != 35 {
		t.Error("Coersion error")
	}
}
