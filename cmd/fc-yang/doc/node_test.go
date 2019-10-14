package doc

import (
	"testing"

	"github.com/freeconf/yang/c2"
	"github.com/freeconf/yang/nodes"
	"github.com/freeconf/yang/source"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
)

func TestApi(t *testing.T) {
	mstr := `module x {
		namespace "";
		prefix "";
		revision 0;
		description "d";
		container a {
			leaf z {
				type string;
			}
		}
		rpc r {
			input {
				leaf x { 
					type string;
				}
			}
			output {
				leaf y {
					type int64;
				}
			}
		}
		notification n {
			container c {
				leaf l {
					type int32;
				}
			}
		}
	}`
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	d := &doc{}
	if d.build(m); d.LastErr != nil {
		t.Fatal(d.LastErr)
	}
	ypath := source.Dir("../../../yang")
	docM := parser.RequireModule(ypath, "fc-doc", "")
	b := node.NewBrowser(docM, api(d))
	actual, err := nodes.WritePrettyJSON(b.Root())
	if err != nil {
		t.Error(err)
	} else {
		c2.Gold(t, *update, []byte(actual), "gold/api.json")
	}
}
