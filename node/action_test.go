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
			s, _ := r.Input.Find("name")
			yourName, _ = s.Get()
			out := map[string]interface{}{
				"salutation": fmt.Sprint("Hello ", yourName.String()),
			}
			return nodeutil.ReflectChild(out), nil
		},
	})
	in, _ := nodeutil.ReadJSON(`{"name":"joe"}`)
	sel, err := b.Root().Find("sayHello")
	fc.RequireEqual(t, nil, err)
	out, err := sel.Action(in)
	fc.RequireEqual(t, nil, err)
	fc.RequireEqual(t, true, out != nil)
	actual, err := nodeutil.WriteJSON(out)
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, "joe", yourName.String())
	fc.AssertEqual(t, `{"salutation":"Hello joe"}`, actual)
}
