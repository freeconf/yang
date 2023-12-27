package node

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/parser"
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
	for _, test := range tests {
		expr, err := ParsePathExpression(test.expression)
		fc.AssertEqual(t, nil, err, test.expression)
		fc.AssertEqual(t, test.nExpressions, len(expr.paths), test.expression)
		actual := expr.String()
		fc.AssertEqual(t, test.expected, actual, test.expression)
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
	module, err := parser.LoadModuleFromString(nil, moduleStr)
	if err != nil {
		t.Fatal(err)
	}
	expr, err := ParsePathExpression("aaa")
	fc.AssertEqual(t, nil, err)

	p1, _ := parseUrlPath("aaa/bbb", module)
	fc.AssertEqual(t, true, expr.PathMatches(p1[0].Parent, p1[len(p1)-1]))

	p2, _ := parseUrlPath("ddd", module)
	fc.AssertEqual(t, false, expr.PathMatches(p2[0].Parent, p2[len(p2)-1]))

	p3, _ := parseUrlPath("aaa/bbb/ccc", module)
	fc.AssertEqual(t, true, expr.PathMatches(p3[0].Parent, p3[len(p3)-1]))

	p4, _ := parseUrlPath("ddd/eee", module)
	fc.AssertEqual(t, false, expr.PathMatches(p4[0].Parent, p4[len(p4)-1]))
}
