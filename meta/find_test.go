package meta

import (
	"testing"

	"github.com/freeconf/yang/fc"
)

func TestFind(t *testing.T) {
	b := &Builder{}
	m := b.Module("m", nil)
	l := b.Leaf(m, "l")
	b.Type(l, "int32")
	fc.AssertEqual(t, nil, Compile(m))
	fc.AssertEqual(t, nil, Find(m, "/../../l"))
}
