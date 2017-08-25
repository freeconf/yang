package meta

import "testing"

func TestMetaDeepCopy(t *testing.T) {
	p := &Container{Ident: "p"}
	c := &Container{Ident: "C"}
	p.AddMeta(c)
	c.AddMeta(&Uses{Ident: "g"})
	g := &Grouping{Ident: "g"}
	p.AddMeta(g)

	copy := DeepCopy(p).(*Container)
	if copy == p {
		t.Error("did not copy")
	}
	if ListLen(p) != ListLen(copy) {
		t.Error("not same size")
	}
}
