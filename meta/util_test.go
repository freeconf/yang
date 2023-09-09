package meta

import (
	"testing"

	"github.com/freeconf/yang/fc"
)

func TestPath(t *testing.T) {
	b := &Builder{}
	m := b.Module("m", nil)
	c := b.Container(m, "c")
	l := b.List(c, "l")
	fc.AssertEqual(t, "m/c/l", SchemaPath(l))
	fc.AssertEqual(t, "c/l", SchemaPathNoModule(l))
}
