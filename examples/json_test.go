package examples

/*
If you already have a browser, reading and writing to JSON is trivial. c2g does
not use annotated source tags.
 */
import (
	"testing"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"strings"
	"fmt"
	"bytes"
)


/*
  Normally you do not need to work with directly json because that happens
  in protocol handlers like RESTCONF.  But there are times it comes in handy.

  Here we're going to read in json according to a model

  Output
  ===============
  map[suggestionBox:map[message:Hello]]
 */
func TestExampleJsonRead(t *testing.T) {
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

	rdr := node.NewJsonReader(strings.NewReader(data)).Node()

	// You can insert, upsert or update json into any other node, but here
	// we'll insert it into a map
	result := make(map[string]interface{})
	b := node.NewBrowser2(model, node.MapNode(result))

	b.Root().InsertFrom(rdr)
	fmt.Print(result)
}

/*
Writing to JSON is as simple as reading.  You can use constraints to control what gets
marshalled to JSON.

Ouput
===========
{"suggestionBox":[{"id":"123","message":"Hello"}]}
 */
func TestExampleJsonWrite(t *testing.T) {
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
		"suggestionBox" : []map[string]interface{} {
			{
				"id" : "123",
				"message" : "Hello",
			},
		},
	}
	var result bytes.Buffer
	wtr := node.NewJsonWriter(&result).Node()

	// You can pull from any other node, you will always want to insert
	// into a json writer node

	b := node.NewBrowser2(model, node.MapNode(data))

	t.Log(b.Root().InsertInto(wtr).LastErr)
	fmt.Print(result.String())
}
