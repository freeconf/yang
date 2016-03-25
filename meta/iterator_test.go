package meta

import (
	"testing"
)

func TestEmptyIterator(t *testing.T) {
	i := EmptyInterator(0)
	if i.HasNextMeta() {
		t.Fail()
	}
}

func TestSingletonIterator(t *testing.T) {
	leaf := &Leaf{Ident: "L"}
	i := &SingletonIterator{leaf}
	if !i.HasNextMeta() {
		t.Fail()
	}
	if i.NextMeta() != leaf {
		t.Fail()
	}
	if i.HasNextMeta() {
		t.Fail()
	}
}

func TestEmptyContainerIterator(t *testing.T) {
	c := &Container{Ident: "C"}
	i := NewMetaListIterator(c, true)
	if i.HasNextMeta() {
		t.Fail()
	}
}

func TestContainerIterator(t *testing.T) {
	c := &Container{Ident: "C"}
	i := NewMetaListIterator(c, true)
	if i.HasNextMeta() {
		t.Fail()
	}
	leaf := &Leaf{Ident: "l"}
	c.AddMeta(leaf)
	i = NewMetaListIterator(c, true)
	if !i.HasNextMeta() {
		t.Fail()
	}
	if i.NextMeta() != leaf {
		t.Fail()
	}
	if i.HasNextMeta() {
		t.Fail()
	}
}

func TestIteratorWithGrouping(t *testing.T) {
	p := &Container{Ident: "p"}
	c := &Container{Ident: "C"}
	p.AddMeta(c)
	c.AddMeta(&Uses{Ident: "g"})
	g := &Grouping{Ident: "g"}
	p.AddMeta(g)
	i := NewMetaListIterator(c, true)
	if i.HasNextMeta() {
		t.Error("Container with uses pointing to empty group should have no items")
	}
	leaf := &Leaf{Ident: "l"}
	g.AddMeta(leaf)
	i = NewMetaListIterator(c, true)
	if !i.HasNextMeta() {
		t.Error("Container with uses pointing to group with one item should be found")
	}
	if i.NextMeta() != leaf {
		t.Error("Container with uses pointing to group with one item is not that item")
	}
	if i.HasNextMeta() {
		t.Error("Container with uses pointing to group with one item did not end on time")
	}
}
