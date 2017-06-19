package orchestrator

import (
	"testing"
	"time"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

func Test_OrchestratorNode(t *testing.T) {
	config := `{
		"app" : [{
			"id": "c1",
			"type":"car",
			"startup":{
				"car" : {
					"speed" : 1
				}
			}
		},{
			"id": "c2",
			"type":"car",
			"startup":{
				"car" : {
					"speed" : 1
				}				
			}
		},{
			"id": "g1",
			"type":"garage"
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
	if assertNumDevices := c2.CheckEqual(3, dm.Len()); assertNumDevices != nil {
		t.Error(assertNumDevices)
	}
	gd, _ := dm.Device("g1")
	gb, _ := gd.Browser("garage")
	var updates int
	gb.Root().Find("maintenance").Notifications(func(msg node.Selection) {
		updates++
	})
	<-time.After(1 * time.Second)
	if updates < 5 {
		t.Errorf("expected at least 5 updates, got %d", updates)
	}
}
