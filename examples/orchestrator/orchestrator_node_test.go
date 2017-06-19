package orchestrator

import (
	"testing"

	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

func Test_OrchestratorNode(t *testing.T) {
	config := `{
		"app" : [{
			"id": "c1",
			"type":"car",
			"startup":{}
		},{
			"id": "c2",
			"type":"car",
			"startup":{}
		},{
			"id": "g1",
			"type":"garage",
			"startup":{}				
		}]
	}`
	dm := device.NewMap()
	f := &LocalFactory{
		Ypath: testYPath,
		Map:   dm,
	}
	o := New(f)
	m := yang.RequireModule(testYPath, "orchestrator")
	b := node.NewBrowser(m, Node(o))
	if err := b.Root().UpsertFrom(node.ReadJson(config)).LastErr; err != nil {
		t.Error(err)
	}
	t.Logf("%d", dm.Len())
}
