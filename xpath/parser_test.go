package xpath

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
)

func TestXPathToString(t *testing.T) {
	yyDebug = 10
	yyErrorVerbose = true
	tests := []struct {
		expr string
	}{
		{
			expr: "a/b",
		},
		{
			expr: "a/b<3",
		},
		{
			expr: "a/b!='x'",
		},
	}
	for _, test := range tests {
		actual, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
		}
		fc.AssertEqual(t, test.expr, actual.String())
	}
}

func TestXpathNs(t *testing.T) {
	b := meta.Builder{}
	m := b.Module("module-x", nil)
	lookup := func(string) (*meta.Module, error) {
		return m, nil
	}
	actual, err := Parse2(lookup, "x:a")
	if err != nil {
		t.Error(err)
	}
	fc.AssertEqual(t, "module-x:a", actual.String())
}
