package meta

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/val"
)

func TestMetaLeafList(t *testing.T) {
	b := &Builder{}
	m := b.Module("m", nil)
	l1 := b.Leaf(m, "x")
	dt1 := b.Type(l1, "string")
	l2 := b.LeafList(m, "y")
	dt2 := b.Type(l2, "string")
	if err := Compile(m); err != nil {
		t.Error(err)
	}
	fc.AssertEqual(t, val.FmtString, dt1.Format())
	fc.AssertEqual(t, val.FmtStringList, dt2.Format())
}

func TestMetaIsConfig(t *testing.T) {
	b := &Builder{}
	m := b.Module("m", nil)
	c := b.Container(m, "c")
	l := b.List(c, "l")
	if err := Compile(m); err != nil {
		t.Error(err)
	}
	if !l.Config() {
		t.Error("Should be config")
	}
}

func TestMetaUses(t *testing.T) {
	b := &Builder{}
	fc.DebugLog(true)
	m := b.Module("m", nil)
	g := b.Grouping(m, "g")
	b.List(g, "l")
	b.Uses(m, "g")
	if err := Compile(m); err != nil {
		t.Error(err)
	}
	fc.AssertEqual(t, "l", m.DataDefinitions()[0].Ident())
}

func TestChoice(t *testing.T) {
	b := &Builder{}
	m := b.Module("m", nil)
	c := b.Choice(m, "c")

	cc1 := b.Case(c, "cc1")
	b.Type(b.Leaf(cc1, "l1"), val.FmtString.String())

	cc2 := b.Case(c, "cc2")
	b.Type(b.Leaf(cc2, "l2"), val.FmtString.String())

	if err := Compile(m); err != nil {
		t.Error(err)
	}
	t.Logf("%v", m.DataDefinitions())
	actual := c.Cases()["cc2"]
	if actual.Ident() != "cc2" {
		t.Error("GetCase failed")
	}
}

func TestRefine(t *testing.T) {
	b := &Builder{}
	m := b.Module("m", nil)
	g := b.Grouping(m, "x")
	u := b.Uses(m, "x")
	b.Type(b.Leaf(g, "l"), val.FmtString.String())

	r := b.Refine(u, "l")
	b.Config(r, false)
	if err := Compile(m); err != nil {
		t.Error(err)
	}
	ddef := m.DataDefinitions()[0]
	if ddef.(HasDetails).Config() {
		t.Fail()
	}
}

func TestIfFeature(t *testing.T) {
	features := map[string]*Feature{
		"foo": &Feature{ident: "foo"},
		"bar": &Feature{ident: "bar"},
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
		iff := &IfFeature{expr: test.expr}
		actual, err := iff.Evaluate(features)
		if err != nil {
			if !test.err {
				t.Error(err)
			}
		} else {
			fc.AssertEqual(t, test.expected, actual)
		}
	}
}

func TestRefineSplit(t *testing.T) {
	r := &Refine{ident: "a/b/c"}
	ident, path := r.splitIdent()
	fc.AssertEqual(t, "a", ident)
	fc.AssertEqual(t, path, "b/c")

	r = &Refine{ident: "a"}
	ident, path = r.splitIdent()
	fc.AssertEqual(t, "a", ident)
	fc.AssertEqual(t, path, "")
}

func TestAugment(t *testing.T) {
	b := &Builder{}
	m := b.Module("m", nil)
	x := b.Container(m, "x")
	b.Type(b.Leaf(x, "a"), val.FmtInt32.String())
	b.Type(b.Leaf(x, "b"), val.FmtInt32.String())

	y := b.Augment(m, "x")
	b.Type(b.Leaf(y, "c"), val.FmtString.String())
	b.Type(b.Leaf(y, "f"), val.FmtString.String())

	if err := Compile(m); err != nil {
		t.Error(err)
	}

	expected := []struct {
		ident  string
		format val.Format
	}{
		{
			"a", val.FmtInt32,
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
	actual := x.DataDefinitions()
	for i, e := range expected {
		if e.ident != actual[i].Ident() {
			t.Errorf("expected %s but got %s", e.ident, actual[i].Ident())
		}
		f := actual[i].(HasType).Type().Format()
		if e.format != f {
			t.Errorf("%s : expected format %s but got %s", e.ident, e.format, f)
		}
	}

}
