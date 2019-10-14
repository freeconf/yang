package meta

import (
	"testing"

	"github.com/freeconf/yang/fc"
)

func TestContainerIterator(t *testing.T) {
	ddefs := []Definition{
		NewContainer(nil, "A"),
		NewLeaf(nil, "B"),
	}
	i := Iterate(ddefs)
	fc.AssertEqual(t, "A", i.Next().Ident())
	fc.AssertEqual(t, "B", i.Next().Ident())
	if i.HasNext() {
		t.Fail()
	}
}
