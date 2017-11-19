package meta

import (
	"testing"

	"github.com/c2stack/c2g/c2"
)

func TestContainerIterator(t *testing.T) {
	ddefs := []DataDef{
		NewContainer("A"),
		NewLeaf("B"),
	}
	i := Iterate(ddefs)
	c2.AssertEqual(t, "A", i.Next().Ident())
	c2.AssertEqual(t, "B", i.Next().Ident())
	if i.HasNext() {
		t.Fail()
	}
}
