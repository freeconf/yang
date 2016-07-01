package node

import (
	"bytes"
	"fmt"
	"github.com/c2g/meta/yang"
	"strings"
	"testing"
)

func TestAction(t *testing.T) {
	y := `
module m {
	prefix "";
	namespace "";
	revision 0000-00-00 {
	  description "";
    }
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
	store := NewBufferStore()
	b := NewStoreData(m, store)
	var yourName *Value
	store.Actions["sayHello"] = func(r ActionRequest) (output Node, err error) {
		if err = r.Input.Selector().InsertInto(b.Node()).LastErr; err != nil {
			return nil, err
		}
		yourName = store.Values["name"]
		store.Values["salutation"] = &Value{Str: fmt.Sprint("Hello ", yourName)}
		return b.Container(""), nil
	}
	in := NewJsonReader(strings.NewReader(`{"name":"joe"}`)).Node()
	var actual bytes.Buffer
	sel := b.Browser().Root().Selector().Find("sayHello").Action(in)
	if sel.LastErr != nil {
		t.Fatal(sel.LastErr)
	}
	if err = sel.InsertInto(NewJsonWriter(&actual).Node()).LastErr; err != nil {
		t.Fatal(err)
	}
	AssertStrEqual(t, "joe", yourName.Str)
	AssertStrEqual(t, `{"salutation":"Hello joe"}`, actual.String())
}
