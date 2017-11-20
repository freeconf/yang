package meta

import "testing"
import "github.com/freeconf/c2g/val"
import "github.com/freeconf/c2g/c2"

func TestMetaLeafList(t *testing.T) {
	dt := NewDataType("string")
	dt.setParent(NewLeaf("x"))
	if err := dt.compile(); err != nil {
		t.Error(err)
	}
	c2.AssertEqual(t, val.FmtString, dt.Format())

	dt.setParent(NewLeafList("x"))
	if err := dt.compile(); err != nil {
		t.Error(err)
	}
	c2.AssertEqual(t, val.FmtStringList, dt.Format())
}

func TestMetaIsConfig(t *testing.T) {
	m := NewModule("m")
	c := NewContainer("c")
	addMeta(t, m, c)
	l := NewList("l")
	addMeta(t, c, l)
	if err := m.compile(); err != nil {
		t.Error(err)
	}
	if !l.Config() {
		t.Error("Should be config")
	}
}

func TestMetaUses(t *testing.T) {
	c2.DebugLog(true)
	m := NewModule("m")
	g := NewGrouping("g")
	addMeta(t, m, g)
	addMeta(t, g, NewList("l"))
	addMeta(t, m, NewUses("g"))
	if err := m.compile(); err != nil {
		t.Error(err)
	}
	c2.AssertEqual(t, "l", m.DataDefs()[0].Ident())
}

func addMeta(t *testing.T, parent Meta, child Meta) {
	t.Helper()
	if err := Set(parent, child); err != nil {
		t.Error(err)
	}
}
func TestChoice(t *testing.T) {
	m := NewModule("m")
	c := NewChoice("c")
	addMeta(t, m, c)
	cc1 := NewChoiceCase("cc1")
	addMeta(t, c, cc1)
	addMeta(t, cc1, NewLeafWithType("l1", val.FmtString))

	cc2 := NewChoiceCase("cc2")
	addMeta(t, c, cc2)
	addMeta(t, cc2, NewLeafWithType("l2", val.FmtString))
	if err := m.compile(); err != nil {
		t.Error(err)
	}
	t.Logf("%v", m.DataDefs())
	actual := c.Case("cc2")
	if actual.Ident() != "cc2" {
		t.Error("GetCase failed")
	}
}

func TestRefine(t *testing.T) {
	m := NewModule("m")
	g := NewGrouping("x")
	addMeta(t, m, g)
	u := NewUses("x")
	addMeta(t, m, u)
	l := NewLeafWithType("l", val.FmtString)
	addMeta(t, g, l)
	r := NewRefine("l")
	if err := Set(r, SetConfig(false)); err != nil {
		t.Error(err)
	}
	addMeta(t, u, r)
	if err := m.compile(); err != nil {
		t.Error(err)
	}
	ddef := m.DataDefs()[0]
	if ddef.(HasDetails).Config() {
		t.Fail()
	}
}

func TestAugment(t *testing.T) {
	m := NewModule("m")
	x := NewContainer("x")
	addMeta(t, m, x)
	a1 := NewLeafWithType("a", val.FmtInt32)
	b1 := NewLeafWithType("b", val.FmtInt32)
	c1 := NewLeafWithType("c", val.FmtInt32)
	addMeta(t, x, a1)
	addMeta(t, x, b1)
	addMeta(t, x, c1)

	y := NewAugment("x")
	addMeta(t, m, y)
	a2 := NewLeafWithType("a", val.FmtString)
	f2 := NewLeafWithType("f", val.FmtString)
	c2 := NewLeafWithType("c", val.FmtString)
	addMeta(t, y, c2)
	addMeta(t, y, a2)
	addMeta(t, y, f2)

	if err := m.compile(); err != nil {
		t.Error(err)
	}

	expected := []struct {
		ident  string
		format val.Format
	}{
		{
			"a", val.FmtString,
		},
		{
			"b", val.FmtInt32,
		},
		{
			"c", val.FmtString,
		},
		{
			"f", val.FmtString,
		},
	}
	actual := x.DataDefs()
	for i, e := range expected {
		if e.ident != actual[i].Ident() {
			t.Errorf("expected %s but got %s", e.ident, actual[i].Ident())
		}
		f := actual[i].(HasDataType).DataType().Format()
		if e.format != f {
			t.Errorf("%s : expected format %s but got %s", e.ident, e.format, f)
		}
	}
}
