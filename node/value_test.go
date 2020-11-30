package node

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
)

func TestCoerseValue(t *testing.T) {
	b := &meta.Builder{}
	m := b.Module("x", nil)
	l := b.Leaf(m, "l")
	dt := b.Type(l, "int32")
	if err := meta.Compile(m); err != nil {
		t.Error(err)
	}
	v, err := NewValue(dt, 35)
	if err != nil {
		t.Error(err)
	} else if v.Value().(int) != 35 {
		t.Error("Coersion error")
	}
}

func TestToIdentRef(t *testing.T) {
	b := &meta.Builder{}
	m := b.Module("x", nil)
	i0 := b.Identity(m, "i0")
	i00 := b.Identity(m, "i00")
	b.Base(i00, "i0")
	if err := meta.Compile(m); err != nil {
		t.Error(err)
	}

	ref, err := toIdentRef(i0, "i00")
	if err != nil {
		t.Error(err)
	}
	fc.AssertEqual(t, "i00", ref.Label)
	fc.AssertEqual(t, "i0", ref.Base)
}
