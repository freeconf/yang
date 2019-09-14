package nodes_test

import (
	"testing"

	"github.com/freeconf/yang/c2"
	"github.com/freeconf/yang/node"

	"github.com/freeconf/yang/nodes"
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
		nodes.ConfigProxy{}.Node(local.Root().Node, remote.Root().Node),
	)

	t.Run("read", func(t *testing.T) {
		actual, err := nodes.WritePrettyJSON(proxy.Root())
		if err != nil {
			t.Fatal(err)
		}
		c2.Gold(t, *updateFlag, []byte(actual), "gold/config_proxy.json")
	})

	t.Run("editContainer", func(t *testing.T) {
		edit := nodes.ReadJSON(`{"class":"thrush"}`)
		if err := proxy.Root().Find("bird=robin/species").InsertFrom(edit).LastErr; err != nil {
			t.Fatal(err)
		}
		c2.AssertEqual(t, "thrush", localData["robin"].Species.Class)
	})

	t.Run("editList", func(t *testing.T) {
		edit := nodes.ReadJSON(`{"wingspan":10}`)
		if err := proxy.Root().Find("bird=robin").UpsertFrom(edit).LastErr; err != nil {
			t.Fatal(err)
		}
		c2.AssertEqual(t, 10, localData["robin"].Wingspan)
	})

	t.Run("addListItem", func(t *testing.T) {
		edit := nodes.ReadJSON(`{"bird":[{"name":"owl"}]}`)
		if err := proxy.Root().Find("bird").InsertFrom(edit).LastErr; err != nil {
			t.Fatal(err)
		}
		c2.AssertEqual(t, "owl", localData["owl"].Name)
	})
}
