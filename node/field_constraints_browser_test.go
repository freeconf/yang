package node_test

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
)

func TestFieldConstraintsRequests(t *testing.T) {
	y := `module m { prefix ""; namespace ""; revision 0;
		leaf p {
			type string {
				pattern "x.*";
			}
		}
}`
	m, err := parser.LoadModuleFromString(nil, y)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		JSON  string
		valid bool
	}{
		{`{"p":"xxxx"}`, true},
		{`{"p":"a"}`, false},
	}

	for _, test := range tests {
		data := make(map[string]interface{})
		b := node.NewBrowser(m, nodeutil.ReflectChild(data))
		root := b.Root()
		n, _ := nodeutil.ReadJSON(test.JSON)
		err = root.UpsertFrom(n)
		fc.AssertEqual(t, test.valid, err == nil, test.JSON)
	}
}
