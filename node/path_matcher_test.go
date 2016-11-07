package node

import (
	"testing"

	"github.com/c2stack/c2g/meta/yang"
)

func TestPathMatcherLex(t *testing.T) {
	tests := []struct {
		expression string
		expected   []string
	}{
		{
			"aaa(bbb;ccc)",
			[]string{"aaa", "(", "bbb", ";", "ccc", ")"},
		},
		{
			"aaa;bbb",
			[]string{"aaa", ";", "bbb"},
		},
	}
	for i, test := range tests {
		l := lex{selector: test.expression}
		for j, e := range test.expected {
			actual := l.next()
			if actual != e {
				t.Errorf("test=%d, segment=%d '%s' != '%s'", i, j, e, actual)
			}
		}
		if !l.done() {
			t.Errorf("%d !done", i)
		}
	}
}

func TestPathMatcherParse(t *testing.T) {
	tests := []struct {
		expression   string
		nExpressions int
		expected     string
	}{
		{
			"aa",
			1,
			"[aa]",
		},
		{
			"aa;bb",
			2,
			"[aa],[bb]",
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
		expr, err := ParsePathExpression(test.expression)
		if err != nil {
			t.Errorf("#%d error parsing expression: %s", i, err.Error())
		}
		if len(expr.slices) != test.nExpressions {
			t.Errorf("#%d wrong number of expected expressions: %d", i, len(expr.slices))
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
	expr, err := ParsePathExpression("aaa")
	p1, _ := ParsePath("aaa/bbb", module)
	if !expr.PathMatches(p1.Head, p1.Tail) {
		t.Fail()
	}
	p2, _ := ParsePath("ddd", module)
	if expr.PathMatches(p2.Head, p2.Tail) {
		t.Fail()
	}
	p3, _ := ParsePath("aaa/bbb/ccc", module)
	if !expr.PathMatches(p3.Head, p3.Tail) {
		t.Fail()
	}
	p4, _ := ParsePath("ddd/eee", module)
	if expr.PathMatches(p4.Head, p4.Tail) {
		t.Fail()
	}
}
