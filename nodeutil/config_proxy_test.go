package nodeutil_test

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"

	"github.com/freeconf/yang/testdata"
)

func TestConfigProxy(t *testing.T) {
	setup := `{
		"bird" : [{
			"name" : "robin",
			"species" : {
				"name" : "robin red breast"
			}
		}]
	}`
	local, localData := testdata.BirdBrowser(setup)
	remote, _ := testdata.BirdBrowser(setup)
	proxy := node.NewBrowser(
		testdata.BirdModule(),
		nodeutil.ConfigProxy{}.Node(local.Root().Node, remote.Root().Node),
	)

	t.Run("read", func(t *testing.T) {
		actual, err := nodeutil.WritePrettyJSON(proxy.Root())
		if err != nil {
			t.Fatal(err)
		}
		fc.Gold(t, *updateFlag, []byte(actual), "gold/config_proxy.json")
	})

	t.Run("editContainer", func(t *testing.T) {
		edit := nodeutil.ReadJSON(`{"class":"thrush"}`)
		sel, err := proxy.Root().Find("bird=robin/species")
		fc.RequireEqual(t, nil, err)
		fc.RequireEqual(t, nil, sel.InsertFrom(edit))
		fc.AssertEqual(t, "thrush", localData["robin"].Species.Class)
	})

	t.Run("editList", func(t *testing.T) {
		edit := nodeutil.ReadJSON(`{"wingspan":10}`)
		sel, err := proxy.Root().Find("bird=robin")
		fc.RequireEqual(t, nil, err)
		fc.RequireEqual(t, nil, sel.UpsertFrom(edit))
		fc.AssertEqual(t, 10, localData["robin"].Wingspan)
	})

	t.Run("addListItem", func(t *testing.T) {
		edit := nodeutil.ReadJSON(`{"bird":[{"name":"owl"}]}`)
		sel, err := proxy.Root().Find("bird")
		fc.RequireEqual(t, nil, err)
		fc.RequireEqual(t, nil, sel.InsertFrom(edit))
		fc.AssertEqual(t, "owl", localData["owl"].Name)
	})
}
