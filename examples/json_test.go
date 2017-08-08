package examples

/*
If you already have a browser, reading and writing to JSON is trivial. c2g does
not use annotated source tags.
*/
import (
	"fmt"
	"testing"

	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
)

/*
  Normally you do not need to work with directly json because that happens
  in protocol handlers like RESTCONF.  But there are times it comes in handy.

  Here we're going to read in json according to a model

  Output
  ===============
  map[suggestionBox:map[message:Hello]]
*/
func Example_jsonRead(t *testing.T) {
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
	data := `
	{
		"suggestionBox" : {
			"message" : "Hello"
		}
	}
	`

	// You can insert, upsert or update json into any other node, but here
	// we'll insert it into a map
	result := make(map[string]interface{})
	b := node.NewBrowser(model, nodes.Reflect(result))

	b.Root().InsertFrom(nodes.ReadJSON(data))
	fmt.Print(result)
}

/*
Writing to JSON is as simple as reading.  You can use constraints to control what gets
marshalled to JSON.

Ouput
===========
{"suggestionBox":[{"id":"123","message":"Hello"}]}
*/
func Example_jsonWrite(t *testing.T) {
	// Model
	model, _ := yang.LoadModuleFromString(nil,
		`module x {
			namespace "";
			prefix "";
			revision 0;

			list suggestionBox {
			   key "id";
			   leaf id { type string; }
			   leaf message { type string; }
			}
		}`)

	// Data
	data := map[string]interface{}{
		"suggestionBox": []map[string]interface{}{
			{
				"id":      "123",
				"message": "Hello",
			},
		},
	}

	b := node.NewBrowser(model, nodes.Reflect(data))
	result, err := nodes.WriteJSON(b.Root())
	if err != nil {
		panic(err)
	}
	fmt.Print(result)
}
