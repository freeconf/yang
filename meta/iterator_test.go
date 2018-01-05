package meta

import (
	"testing"

	"github.com/freeconf/gconf/c2"
)

func TestContainerIterator(t *testing.T) {
	ddefs := []Definition{
		NewContainer(nil, "A"),
		NewLeaf(nil, "B"),
	}
	i := Iterate(ddefs)
	c2.AssertEqual(t, "A", i.Next().Ident())
	c2.AssertEqual(t, "B", i.Next().Ident())
	if i.HasNext() {
		t.Fail()
	}
}
