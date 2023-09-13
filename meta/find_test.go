package meta

import (
	"testing"

	"github.com/freeconf/yang/fc"
)

func TestFindBadPath(t *testing.T) {
	b := &Builder{}
	m := b.Module("m", nil)
	l := b.Leaf(m, "l")
	b.Type(l, "int32")
	fc.AssertEqual(t, nil, Compile(m))
	fc.AssertEqual(t, nil, Find(m, "/../../l"))
}

func TestFindChoice(t *testing.T) {
	b := &Builder{}
	m := b.Module("m", nil)
	c := b.Choice(m, "c")
	c1 := b.Case(c, "one")
	l := b.Leaf(c1, "1")
	b.Type(l, "int32")
	fc.AssertEqual(t, nil, Compile(m))
	fc.AssertEqual(t, c, Find(m, "c"))
	fc.AssertEqual(t, c1, Find(m, "c/one"))
	fc.AssertEqual(t, nil, Find(m, "c/bogus"))
}
