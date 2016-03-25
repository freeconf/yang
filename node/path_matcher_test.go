package node

import (
	"testing"
	"meta/yang"
	"meta"
)

func TestPathMatcherLex(t *testing.T) {
	l := lex{selector:"aaa(bbb;ccc)"}
	expected := []string {
		"aaa", "(", "bbb", ";", "ccc", ")",
	}
	for i, e := range expected {
		actual := l.next()
		if actual != e {
			t.Error(i, e, "!=", actual)
		}
	}
	if ! l.done() {
		t.Error("!done")
	}
}

func TestPathMatcherParse(t *testing.T) {
	tests := []struct{
		expression   string
		nExpressions int
		expected string
	}{
		{
			"aa",
			1,
			"[aa]",
		},
		{
			"aa/bbb",
			1,
			"[aa,bbb]",
		},
		{
			"aa/(bbb;ccc)",
			2,
			"[aa,bbb],[aa,ccc]",
		},
		{
			"aa/(bbb;ccc)ddd",
			2,
			"[aa,bbb,ddd],[aa,ccc,ddd]",
		},
		{
			"aa/(bbb;ccc)(ddd;eee)",
			4,
			"[aa,bbb,ddd],[aa,bbb,eee],[aa,ccc,ddd],[aa,ccc,eee]",
		},
		{
			"aa/bbb/ccc;ddd",
			2,
			"[aa,bbb,ccc],[ddd]",
		},
	}
	for i, test := range tests {
		expr, err := ParsePathExpression(nil, test.expression)
		if err != nil {
			t.Errorf("#%d %s", i, err.Error())
		}
		if len(expr.slices) != test.nExpressions {
			t.Errorf("#%d %d", i, len(expr.slices))
		}
		actual := expr.String()
		if actual != test.expected {
			t.Errorf("\nExpected:%s\n  Actual:%s", test.expected, actual)
		}
	}
}

func TestPathMatcherMatch(t *testing.T) {
	moduleStr := `
module m {
	prefix "";
	namespace "";
	revision 0;
	container aaa {
		list bbb {
			leaf ccc {
				type string;
			}
		}
	}
	container ddd {
		leaf eee {
			type string;
		}
	}
}
	`
	module, err := yang.LoadModuleCustomImport(moduleStr, nil)
	if err != nil {
		t.Fatal(err)
	}
	expr, err := ParsePathExpression(NewRootPath(module), "aaa")
	p1, _ := ParsePath("aaa/bbb", module)
	if ! expr.PathMatches(p1.Tail) {
		t.Fail()
	}
	p2, _ := ParsePath("ddd", module)
	if expr.PathMatches(p2.Tail) {
		t.Fail()
	}
	p3, _ := ParsePath("aaa/bbb/ccc", module)
	if ! expr.FieldMatches(p3.Tail.parent, p3.Tail.goober.(meta.HasDataType)) {
		t.Fail()
	}
	p4, _ := ParsePath("ddd/eee", module)
	if expr.FieldMatches(p4.Tail.parent, p4.Tail.goober.(meta.HasDataType)) {
		t.Fail()
	}
}
