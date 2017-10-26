package meta

import "testing"

func TestMixin(t *testing.T) {
	x := &Container{Ident: "x"}
	a1 := NewLeaf("a", "int32")
	b1 := NewLeaf("b", "int32")
	c1 := NewLeaf("c", "int32")
	x.AddMeta(a1)
	x.AddMeta(b1)
	x.AddMeta(c1)

	y := &Container{Ident: "y"}
	a2 := NewLeaf("a", "int32")
	f2 := NewLeaf("f", "int32")
	c2 := NewLeaf("c", "int32")
	y.AddMeta(c2)
	y.AddMeta(a2)
	y.AddMeta(f2)
	err := mixin(x, y)
	if err != nil {
		t.Fatal(err)
	}
	actual := Children(x)
	expected := []Meta{
		a2, b1, c2, f2,
	}
	for _, e := range expected {
		a, err := actual.Next()
		if err != nil {
			t.Fatal(err)
		}
		if a != e {
			t.Errorf("expected %s(%p) but got %s(%p)", e.GetIdent(), e, a.GetIdent(), a)
		}
	}
	if actual.HasNext() {
		t.Error("extra items")
	}
}
