package nodes_test

import (
	"fmt"

	"github.com/freeconf/gconf/nodes"
)

func ExampleReadJSON() {
	model := `
		list bird {

			key "name";

			leaf name {
				type string;
			}

			leaf wingSpan {
				type int32;
			}
		}

		container location {
			leaf continent {
				type enumeration {
					enum northAmerica;
					enum southAmerica;
					enum africa;
					enum antartica;
					enum europe;
					enum austrailia;
					enum asia;
				}
			}
		}
	`

	myApp := make(map[string]interface{})
	sel := exampleSelection(model, nodes.ReflectChild(myApp))
	data := `{
		"bird" : [{
			"name": "swallow",
			"wingSpan" : 10
		}],
		"ignored" : "because it's not in model",
		"location" : {
			"continent" : "africa"
		}
	}
	`
	if err := sel.InsertFrom(nodes.ReadJSON(data)).LastErr; err != nil {
		fmt.Print(err.Error())
	}
	out, _ := nodes.WriteJSON(sel)
	fmt.Printf(out)
	// Output:
	// {"bird":[{"name":"swallow","wingSpan":10}],"location":{"continent":"africa"}}
}
