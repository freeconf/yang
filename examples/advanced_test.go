package examples

import (
	"fmt"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"testing"
)

/*
TestExampleReflection uses reflection to look for containers, lists and leafs on
data. While reflection implementation does handle a variety of data coersion, for
anything too advanced, you'll want to look at node.Extend to write your own custom
data handling.

Output:
==============
Hello
*/
func TestExampleReflection(t *testing.T) {
	// Model
	model, _ := yang.LoadModuleFromString(nil,
		`module x {
			namespace "";
			prefix "";
			revision 0;

			container suggestionBox {
			   leaf message { type string; }
			}
		}`)

	// Data
	msg := &ExampleApp{
		SuggestionBox: &ExampleMessage{
			Id:      "123",
			Message: "Hello",
		},
	}
	data := node.MarshalContainer(msg)

	// Browser = Model + Data
	brwsr := node.NewBrowser(model, data)

	// Read
	out, _ := brwsr.Root().Find("suggestionBox").Get("message")
	fmt.Println(out)
}

/*
TestExampleReflectExtend uses reflection to look for containers, lists and leafs on
data.  While reflection if very handy, often handling small deviations when model
differs from source code can be difficult. For this, you'll want to use the reflection
combined with other methods details in other examples to handle these cases, most notably
node.Extend

Output:
==============
5
*/
func TestExampleReflectExtend(t *testing.T) {
	// Model
	model, _ := yang.LoadModuleFromString(nil,
		`module x {
			namespace "";
			prefix "";
			revision 0;

			container suggestionBox {
			   leaf message { type string; }
			   leaf length { config "false"; type int32; }
			}
		}`)

	// Data
	app := &ExampleApp{
		SuggestionBox: &ExampleMessage{
			Id:      "123",
			Message: "Hello",
		},
	}
	boxData := func(msg *ExampleMessage) node.Node {
		return &node.Extend{
			Node: node.MarshalContainer(msg),
			OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
				switch r.Meta.GetIdent() {
				case "length":
					hnd.Val = &node.Value{Int: len(msg.Message)}
					return nil
				}
				return p.Field(r, hnd)
			},
		}
	}
	data := &node.MyNode{
		OnSelect: func(r node.ContainerRequest) (node.Node, error) {
			return boxData(app.SuggestionBox), nil
		},
	}

	// Browser = Model + Data
	brwsr := node.NewBrowser(model, data)

	// Read
	out, _ := brwsr.Root().Find("suggestionBox").Get("length")
	fmt.Println(out)
}

type ExampleApp struct {
	SuggestionBox *ExampleMessage
}

type ExampleMessage struct {
	Id      string
	Message string
}
