package xpath

import "testing"
import "github.com/c2stack/c2g/c2"

func Test_XPathToString(t *testing.T) {
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
		if err := c2.CheckEqual(test.expr, actual.String()); err != nil {
			t.Error(err)
		}
	}
}
