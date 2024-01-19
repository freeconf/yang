package nodeutil_test

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/nodeutil"

	"github.com/freeconf/yang/node"

	"github.com/freeconf/yang/testdata"
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
	var err error
	for _, test := range tests {
		t.Log(test.desc)
		a, aBirds := testdata.BirdBrowser(setup)
		b, _ := testdata.BirdBrowser(setup)
		c := node.NewBrowser(a.Meta, nodeutil.CopyOnWrite{}.Node(a.Root(), a.Root().Node, b.Root().Node))
		sel := c.Root()
		if test.sel != "" {
			sel, err = sel.Find(test.sel)
			fc.RequireEqual(t, nil, err)
		}
		n, _ := nodeutil.ReadJSON(test.change)
		fc.RequireEqual(t, nil, sel.UpsertFrom(n))
		fc.AssertEqual(t, 1, len(aBirds))
		actual, _ := nodeutil.WritePrettyJSON(b.Root())
		fc.Gold(t, *updateFlag, []byte(actual), test.gold)
	}
}
