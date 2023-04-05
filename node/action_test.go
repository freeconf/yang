package node_test

import (
	"fmt"
	"testing"

	"github.com/freeconf/yang/fc"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/val"
)

func TestAction(t *testing.T) {
	y := `
module m { prefix ""; namespace ""; revision 0;
    rpc sayHello {
      input {
        leaf name {
          type string;
        }
      }
      output {
        leaf salutation {
          type string;
        }
      }
    }
}`
	m, err := parser.LoadModuleFromString(nil, y)
	if err != nil {
		t.Fatal(err)
	}
	// lazy trick, we stick all data, input, output into one bucket
	var yourName val.Value
	b := node.NewBrowser(m, &nodeutil.Basic{
		OnAction: func(r node.ActionRequest) (output node.Node, err error) {
			yourName, _ = r.Input.Find("name").Get()
			out := map[string]interface{}{
				"salutation": fmt.Sprint("Hello ", yourName.String()),
			}
			return nodeutil.ReflectChild(out), nil
		},
	})
	in := nodeutil.ReadJSON(`{"name":"joe"}`)
	sel := b.Root().Find("sayHello").Action(in)
	if sel.LastErr != nil {
		t.Fatal(sel.LastErr)
	}
	actual, err := nodeutil.WriteJSON(sel)
	if err != nil {
		t.Fatal(err)
	}
	fc.AssertEqual(t, "joe", yourName.String())
	fc.AssertEqual(t, `{"salutation":"Hello joe"}`, actual)
}
