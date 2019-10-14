package xpath

import (
	"testing"

	"github.com/freeconf/yang/fc"
)

func TestXPathToString(t *testing.T) {
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
