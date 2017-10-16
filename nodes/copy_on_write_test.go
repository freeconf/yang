package nodes_test

import (
	"testing"

	"github.com/c2stack/c2g/c2"

	"github.com/c2stack/c2g/node"

	"github.com/c2stack/c2g/nodes"
	"github.com/c2stack/c2g/testdata"
)

func TestCopyOnWrite(t *testing.T) {
	setup := `{
		"bird" : [{
			"name" : "robin",
			"species" : {
				"name" : "robin red breast"
			}
		}]
	}`
	tests := []struct {
		desc   string
		sel    string
		change string
		gold   string
	}{
		{
			desc:   "list add",
			sel:    "bird",
			change: `{"bird":[{"name":"owl"}]}`,
			gold:   "gold/cow-list-add.json",
		},
		{
			desc:   "list edit",
			sel:    "bird",
			change: `{"bird":[{"name":"robin", "wingspan":11}]}`,
			gold:   "gold/cow-list-edit.json",
		},
		{
			desc:   "list edit from root",
			sel:    "",
			change: `{"bird":[{"name":"robin", "wingspan":11}]}`,
			gold:   "gold/cow-list-edit-2.json",
		},
	}
	for _, test := range tests {
		t.Log(test.desc)
		a, aBirds := testdata.BirdBrowser(setup)
		b, _ := testdata.BirdBrowser(setup)
		c := node.NewBrowser(a.Meta, nodes.CopyOnWrite{}.Node(a.Root(), a.Root().Node, b.Root().Node))
		sel := c.Root()
		if test.sel != "" {
			sel = sel.Find(test.sel)
		}
		if err := sel.UpsertFrom(nodes.ReadJSON(test.change)).LastErr; err != nil {
			t.Fatal(err)
		}
		c2.AssertEqual(t, 1, len(aBirds))
		actual, _ := nodes.WritePrettyJSON(b.Root())
		c2.Gold(t, *updateFlag, []byte(actual), test.gold)
	}
}
