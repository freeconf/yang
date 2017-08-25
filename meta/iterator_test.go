package meta

import (
	"testing"
)

func TestEmptyIterator(t *testing.T) {
	i := empty(struct{}{})
	if i.HasNext() {
		t.Fail()
	}
}

func TestSingletonIterator(t *testing.T) {
	leaf := &Leaf{Ident: "L"}
	i := &single{leaf}
	if !i.HasNext() {
		t.Fail()
	}
	if l, _ := i.Next(); l != leaf {
		t.Fail()
	}
	if i.HasNext() {
		t.Fail()
	}
}

func TestEmptyContainerIterator(t *testing.T) {
	c := &Container{Ident: "C"}
	i := Children(c, true)
	if i.HasNext() {
		t.Fail()
	}
}

func TestContainerIterator(t *testing.T) {
	c := &Container{Ident: "C"}
	i := Children(c, true)
	if i.HasNext() {
		t.Fail()
	}
	leaf := &Leaf{Ident: "l"}
	c.AddMeta(leaf)
	i = Children(c, true)
	if !i.HasNext() {
		t.Fail()
	}
	if l, _ := i.Next(); l != leaf {
		t.Fail()
	}
	if i.HasNext() {
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
	i := Children(c, true)
	if i.HasNext() {
		t.Error("Container with uses pointing to empty group should have no items")
	}
	leaf := &Leaf{Ident: "l"}
	g.AddMeta(leaf)
	i = Children(c, true)
	if !i.HasNext() {
		t.Error("Container with uses pointing to group with one item should be found")
	}
	if l, _ := i.Next(); l != leaf {
		t.Error("Container with uses pointing to group with one item is not that item")
	}
	if i.HasNext() {
		t.Error("Container with uses pointing to group with one item did not end on time")
	}
}
