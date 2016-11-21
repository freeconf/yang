package browse

import (
	"testing"

	"github.com/c2stack/c2g/meta/yang"
)

func TestDocBuild(t *testing.T) {
	mstr := `module x {
	namespace "";
	prefix "";
	revision 0;
	container x {
		leaf z {
			type string;
		}
	}
}`
	m, err := yang.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	doc := &Doc{}
	if doc.Build(m, "html"); doc.LastErr != nil {
		t.Fatal(doc.LastErr)
	}
	// TODO: Compare to golden file
}

func TestEscape(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{
			in:       "a",
			expected: "Bingo wxas his nxame-o",
		},
		{
			in:       "ao",
			expected: "Bingxo wxas his nxame-xo",
		},
	}
	for _, test := range tests {
		f := escape(test.in, "x")
		actual := f("Bingo was his name-o")
		if actual != test.expected {
			t.Error(actual, test.expected)
		}
	}
}
