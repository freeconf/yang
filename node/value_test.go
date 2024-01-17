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
	l2 := b.Leaf(m, "l2")
	dt2 := b.Type(l2, "empty")

	if err := meta.Compile(m); err != nil {
		t.Error(err)
	}
	v, err := NewValue(dt, 35)
	if err != nil {
		t.Error(err)
	} else if v.Value().(int) != 35 {
		t.Error("Coersion error")
	}

	v2, err := NewValue(dt2, "anything")
	if err != nil {
		t.Error(err)
	} else if v2.Value() == nil {
		t.Error("Coersion error")
	}
}

func TestToIdentRef(t *testing.T) {
	b := &meta.Builder{}
	m := b.Module("x", nil)
	i0 := []*meta.Identity{b.Identity(m, "i0")}
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
}

func TestToUnion(t *testing.T) {
	b := &meta.Builder{}
	m := b.Module("x", nil)
	l := b.Leaf(m, "l")
	u := b.Type(l, "union")
	b.Type(u, "int32")
	b.Type(u, "string")
	fc.RequireEqual(t, nil, meta.Compile(m))
	v, err := NewValue(u, "32")
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, 32, v.Value())

	v, err = NewValue(u, "thirty-two")
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, "thirty-two", v.Value())
}

func TestToUnionList(t *testing.T) {
	b := &meta.Builder{}
	m := b.Module("x", nil)
	l := b.LeafList(m, "l")
	u := b.Type(l, "union")
	b.Type(u, "int32")
	b.Type(u, "string")
	fc.RequireEqual(t, nil, meta.Compile(m))
	v, err := NewValue(u, []string{"32"})
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, []int{32}, v.Value())

	v, err = NewValue(u, []string{"thirty-two", "thirty-three"})
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, []string{"thirty-two", "thirty-three"}, v.Value())
}

func TestToBits(t *testing.T) {
	b := &meta.Builder{}
	m := b.Module("x", nil)
	l := b.Leaf(m, "l")
	dt := b.Type(l, "bits")
	b0 := b.Bit(dt, "b0")
	b.Position(b0, 0)
	b1 := b.Bit(dt, "b1")
	b.Position(b1, 1)
	fc.RequireEqual(t, nil, meta.Compile(m))

	// cast from int (empty)
	v, err := NewValue(dt, 0)
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, uint64(0), v.Value())
	fc.AssertEqual(t, "", v.String())

	// cast from int (non empty)
	v, err = NewValue(dt, 0b11)
	fc.AssertEqual(t, uint64(0b11), v.Value())
	fc.AssertEqual(t, "b0 b1", v.String())

	// cast from string
	v, err = NewValue(dt, "b1")
	fc.AssertEqual(t, uint64(0b10), v.Value())
	fc.AssertEqual(t, "b1", v.String())

	// cast from []string (wrong order)
	v, err = NewValue(dt, []string{"b1", "b0"})
	fc.AssertEqual(t, uint64(0b11), v.Value())
	// side effect: keeping order so input and output data are equal
	fc.AssertEqual(t, "b1 b0", v.String())
}
