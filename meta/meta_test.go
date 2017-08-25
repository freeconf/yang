package meta

import "testing"
import "github.com/c2stack/c2g/val"

func TestLeafListFormatSetting(t *testing.T) {
	leafList := &LeafList{}
	leafList.SetDataType(NewDataType(leafList, "string"))

	if i, _ := leafList.DataType.Info(); i.Format != val.FmtStringList {
		t.Errorf("Not converted to list : %d", i.Format)
	}
}

func TestMetaIsConfig(t *testing.T) {
	m := &Module{Ident: "m"}
	c := &Container{Ident: "c"}
	m.AddMeta(c)
	l := &List{Ident: "l"}
	c.AddMeta(l)
	path := &MetaPath{
		parent: &MetaPath{
			meta: m,
		},
		meta: c,
	}
	if !l.Details().Config(path) {
		t.Error("Should be config")
	}
	c.details.SetConfig(false)
	if l.Details().Config(path) {
		t.Errorf(" %s should not be config", path.String())
	}
}

func TestMetaList(t *testing.T) {
	g1 := &Grouping{Ident: "G1"}
	g2 := &Grouping{Ident: "G2"}
	c := MetaContainer{}
	c.AddMeta(g1)
	c.AddMeta(g2)
	if c.FirstMeta != g1 {
		t.Error("g1 is first child of container")
	}
	if c.LastMeta != g2 {
		t.Error("g2 is last child of container")
	}
	if g1.GetParent() != &c {
		t.Error("g1 parent is not container")
	}
	if g2.GetParent() != &c {
		t.Error("g2 parent is not container")
	}
	if g1.Sibling != g2 {
		t.Error("g1 is not linked to g2")
	}
	if g2.Sibling != nil {
		t.Error("g2 sibling should be nil")
	}
}

func TestMetaProxy(t *testing.T) {
	g1 := &Grouping{Ident: "G1"}
	g1a := &Leaf{Ident: "G1A"}
	g1.AddMeta(g1a)
	u1 := &Uses{Ident: "G1"}
	groupings := MetaContainer{}
	groupings.AddMeta(g1)
	u1.grouping = g1
	i := u1.ResolveProxy()
	nextMeta, err := i.Next()
	if err != nil {
		t.Error(err)
	} else if nextMeta == nil {
		t.Error("resolved registrar is nil")
	} else if nextMeta != g1a {
		t.Error("expected G1A and got ", nextMeta)
	}

	uparent := MetaContainer{}
	uparent.AddMeta(u1)
	i2 := Children(&uparent, true)
	nextResolvedMeta, err := i2.Next()
	if err != nil {
		t.Error(err)
	} else if nextResolvedMeta != g1a {
		t.Error("resolved in iterator didn't work")
	}
}

func TestChoiceGetCase(t *testing.T) {
	c1 := Choice{Ident: "c1"}
	cc1 := ChoiceCase{Ident: "cc1"}
	l1 := Leaf{Ident: "l1"}
	cc1.AddMeta(&l1)
	cc2 := ChoiceCase{Ident: "cc2"}
	l2 := Leaf{Ident: "l2"}
	cc2.AddMeta(&l2)
	c1.AddMeta(&cc1)
	c1.AddMeta(&cc2)
	actual, _ := c1.GetCase("cc2")
	if actual.GetIdent() != "cc2" {
		t.Error("GetCase failed")
	}
}
