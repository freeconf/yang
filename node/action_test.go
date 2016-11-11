package node

import (
	"bytes"
	"fmt"
	"github.com/c2stack/c2g/meta/yang"
	"strings"
	"testing"
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
	var yourName *Value
	b := NewBrowser(m, &MyNode{
		OnAction: func(r ActionRequest) (output Node, err error) {
			yourName, _ = r.Input.GetValue("name")
			out := map[string]interface{}{
				"salutation": fmt.Sprint("Hello ", yourName.Str),
			}
			return MapNode(out), nil
		},
	})
	in := NewJsonReader(strings.NewReader(`{"name":"joe"}`)).Node()
	var actual bytes.Buffer
	sel := b.Root().Find("sayHello").Action(in)
	if sel.LastErr != nil {
		t.Fatal(sel.LastErr)
	}
	if err = sel.InsertInto(NewJsonWriter(&actual).Node()).LastErr; err != nil {
		t.Fatal(err)
	}
	AssertStrEqual(t, "joe", yourName.Str)
	AssertStrEqual(t, `{"salutation":"Hello joe"}`, actual.String())
}
