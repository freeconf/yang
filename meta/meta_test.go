package meta

import (
	"testing"

	"github.com/freeconf/c2g/c2"
	"github.com/freeconf/c2g/val"
)

func TestMetaLeafList(t *testing.T) {
	m := NewModule("m", nil)
	l1 := NewLeaf(m, "x")
	addMeta(t, m, l1)
	dt1 := NewDataType(l1, "string")
	addMeta(t, l1, dt1)
	l2 := NewLeafList(m, "y")
	addMeta(t, m, l2)
	dt2 := NewDataType(l2, "string")
	addMeta(t, l2, dt2)
	if err := Validate(m); err != nil {
		t.Error(err)
	}
	c2.AssertEqual(t, val.FmtString, dt1.Format())
	c2.AssertEqual(t, val.FmtStringList, dt2.Format())
}

func TestMetaIsConfig(t *testing.T) {
	m := NewModule("m", nil)
	c := NewContainer(m, "c")
	addMeta(t, m, c)
	l := NewList(c, "l")
	addMeta(t, c, l)
	if err := Validate(m); err != nil {
		t.Error(err)
	}
	if !l.Config() {
		t.Error("Should be config")
	}
}

func TestMetaUses(t *testing.T) {
	c2.DebugLog(true)
	m := NewModule("m", nil)
	g := NewGrouping(m, "g")
	addMeta(t, m, g)
	addMeta(t, g, NewList(g, "l"))
	addMeta(t, m, NewUses(m, "g"))
	if err := Validate(m); err != nil {
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
	m := NewModule("m", nil)
	c := NewChoice(m, "c")
	addMeta(t, m, c)
	cc1 := NewChoiceCase(c, "cc1")
	addMeta(t, c, cc1)
	addMeta(t, cc1, NewLeafWithType(cc1, "l1", val.FmtString))

	cc2 := NewChoiceCase(c, "cc2")
	addMeta(t, c, cc2)
	addMeta(t, cc2, NewLeafWithType(cc2, "l2", val.FmtString))
	if err := Validate(m); err != nil {
		t.Error(err)
	}
	t.Logf("%v", m.DataDefs())
	actual := c.Cases()["cc2"]
	if actual.Ident() != "cc2" {
		t.Error("GetCase failed")
	}
}

func TestRefine(t *testing.T) {
	m := NewModule("m", nil)
	g := NewGrouping(m, "x")
	addMeta(t, m, g)
	u := NewUses(m, "x")
	addMeta(t, m, u)
	l := NewLeafWithType(g, "l", val.FmtString)
	addMeta(t, g, l)
	r := NewRefine(u, "l")
	if err := Set(r, SetConfig(false)); err != nil {
		t.Error(err)
	}
	addMeta(t, u, r)
	if err := Validate(m); err != nil {
		t.Error(err)
	}
	ddef := m.DataDefs()[0]
	if ddef.(HasDetails).Config() {
		t.Fail()
	}
}

func TestIfFeature(t *testing.T) {
	features := map[string]*Feature{
		"foo": NewFeature(nil, "foo"),
		"bar": NewFeature(nil, "bar"),
	}
	tests := []struct {
		expr     string
		expected bool
		err      bool
	}{
		{
			expr:     "foo",
			expected: true,
		},
		{
			expr:     "not foo",
			expected: false,
		},
		{
			expr:     "not ( foo )",
			expected: false,
		},
		{
			expr:     "goo",
			expected: false,
		},
		{
			expr:     "foo and goo",
			expected: false,
		},
		{
			expr:     "foo or goo",
			expected: true,
		},
		{
			expr:     "not foo or goo",
			expected: false,
		},
		{
			expr:     "not (foo and goo)",
			expected: true,
		},
		{
			expr:     "not foo or bar and baz",
			expected: false,
		},
		{
			expr:     "(not foo) or (bar and baz)",
			expected: false,
		},
		{
			expr:     "not not foo",
			expected: true,
		},
		{
			expr: "foo bar",
			err:  true,
		},
		{
			expr: "and foo",
			err:  true,
		},
	}
	for _, test := range tests {
		t.Log(test.expr)
		iff := NewIfFeature(test.expr)
		actual, err := iff.Evaluate(features)
		if err != nil {
			if !test.err {
				t.Error(err)
			}
		} else {
			c2.AssertEqual(t, test.expected, actual)
		}
	}
}

func TestRefineSplit(t *testing.T) {
	r := NewRefine(nil, "a/b/c")
	ident, path := r.splitIdent()
	c2.AssertEqual(t, "a", ident)
	c2.AssertEqual(t, path, "b/c")

	r = NewRefine(nil, "a")
	ident, path = r.splitIdent()
	c2.AssertEqual(t, "a", ident)
	c2.AssertEqual(t, path, "")
}

func TestAugment(t *testing.T) {
	m := NewModule("m", nil)
	x := NewContainer(m, "x")
	addMeta(t, m, x)
	a1 := NewLeafWithType(x, "a", val.FmtInt32)
	b1 := NewLeafWithType(x, "b", val.FmtInt32)
	c1 := NewLeafWithType(x, "c", val.FmtInt32)
	addMeta(t, x, a1)
	addMeta(t, x, b1)
	addMeta(t, x, c1)

	y := NewAugment(m, "x")
	addMeta(t, m, y)
	a2 := NewLeafWithType(y, "a", val.FmtString)
	f2 := NewLeafWithType(y, "f", val.FmtString)
	c2 := NewLeafWithType(y, "c", val.FmtString)
	addMeta(t, y, c2)
	addMeta(t, y, a2)
	addMeta(t, y, f2)

	if err := Validate(m); err != nil {
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
