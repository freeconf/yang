package node_test

import (
	"fmt"
	"testing"

	"github.com/c2stack/c2g/c2"

	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
	"github.com/c2stack/c2g/val"
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
	m, err := yang.LoadModuleCustomImport(y, nil)
	if err != nil {
		t.Fatal(err)
	}
	// lazy trick, we stick all data, input, output into one bucket
	var yourName val.Value
	b := node.NewBrowser(m, &nodes.Basic{
		OnAction: func(r node.ActionRequest) (output node.Node, err error) {
			yourName, _ = r.Input.GetValue("name")
			out := map[string]interface{}{
				"salutation": fmt.Sprint("Hello ", yourName.String()),
			}
			return nodes.Reflect(out), nil
		},
	})
	in := nodes.ReadJSON(`{"name":"joe"}`)

	sel := b.Root().Find("sayHello").Action(in)
	if sel.LastErr != nil {
		t.Fatal(sel.LastErr)
	}
	actual, err := nodes.WriteJSON(sel)
	if err != nil {
		t.Fatal(err)
	}
	if err := c2.CheckEqual("joe", yourName.String()); err != nil {
		t.Error(err)
	}
	if err := c2.CheckEqual(`{"salutation":"Hello joe"}`, actual); err != nil {
		t.Error(err)
	}
}
