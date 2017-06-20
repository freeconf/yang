package launch

import (
	"testing"
	"time"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

func Test_NodeInMemory(t *testing.T) {
	config := `{
		"launcher" : {
			"inMemory" :{}
		},
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
	o := New()
	m := yang.RequireModule(testYPath, "launch")
	b := node.NewBrowser(m, Node(o, testYPath))
	sel := b.Root()
	if err := sel.UpsertFrom(node.ReadJson(config)).LastErr; err != nil {
		t.Error(err)
	}
	if v, err := sel.GetValue("count"); err != nil {
		t.Error(err)
	} else {
		if assertNumDevices := c2.CheckEqual(3, v.Int); assertNumDevices != nil {
			t.Error(assertNumDevices)
		}
	}
	dm := o.Launcher.(*InMemory).Map
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

func Test_NodeExec(t *testing.T) {
	config := `{
		"launcher" : {
			"exec" :{}
		},
		"app" : [{
			"id": "c1",
			"type":"car",
			"startup":{
				"car" : {
					"speed" : 1
				},
				"restconf" : {
					"web" : {
						"port" : ":8090"
					},
					"debug" : true,
					"callHome" : {
						"deviceId" : "c1",
						"address" : "http://127.0.0.1:8080/restconf",
						"localAddress" : "http://{REQUEST_ADDRESS}:8090/restconf"
					}
				}				
			}
		},{
			"id": "c2",
			"type":"car",
			"startup":{
				"car" : {
					"speed" : 1
				},				
				"restconf" : {
					"web" : {
						"port" : ":8091"
					},
					"debug" : true,
					"callHome" : {
						"deviceId" : "c1",
						"address" : "http://127.0.0.1:8080/restconf",
						"localAddress" : "http://{REQUEST_ADDRESS}:8091/restconf"
					}
				}				
			}
		},{
			"id": "g1",
			"type":"garage",
			"startup" : {
				"restconf" : {
					"web" : {
						"port" : ":8092"
					},
					"callHome" : {
						"deviceId" : "g1",
						"localAddress" : "http://{REQUEST_ADDRESS}:8092/restconf",
						"address" : "http://127.0.0.1:8080/restconf"
					}
				}				
			}
		},{
			"id": "p1",
			"type":"proxy",
			"startup" : {
				"restconf" : {
					"web" : {
						"port" : ":8080"
					}
				}				
			}
		}]
	}`
	o := New()
	m := yang.RequireModule(testYPath, "launch")
	b := node.NewBrowser(m, Node(o, testYPath))
	sel := b.Root()
	if err := sel.UpsertFrom(node.ReadJson(config)).LastErr; err != nil {
		t.Error(err)
	}
	if v, err := sel.GetValue("count"); err != nil {
		t.Error(err)
	} else {
		if assertNumDevices := c2.CheckEqual(4, v.Int); assertNumDevices != nil {
			t.Error(assertNumDevices)
		}
	}
	// dm := o.Builder.(*InMemory).Map
	// gd, _ := dm.Device("g1")
	// gb, _ := gd.Browser("garage")
	// var updates int
	// gb.Root().Find("maintenance").Notifications(func(msg node.Selection) {
	// 	updates++
	// })
	//<-time.After(30 * time.Second)
	//select {}
	// if updates < 5 {
	// 	t.Errorf("expected at least 5 updates, got %d", updates)
	// }
}
