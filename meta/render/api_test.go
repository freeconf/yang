package render

import (
	"testing"

	"github.com/freeconf/gconf/meta"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/nodes"

	"github.com/freeconf/gconf/meta/yang"
	"github.com/freeconf/gconf/node"
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
	m, err := yang.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	doc := &Doc{}
	if doc.Build(m); doc.LastErr != nil {
		t.Fatal(doc.LastErr)
	}
	ypath := &meta.FileStreamSource{Root: "../../yang"}
	docM := yang.RequireModule(ypath, "fc-doc")
	b := node.NewBrowser(docM, Api(doc))
	actual, err := nodes.WritePrettyJSON(b.Root())
	if err != nil {
		t.Error(err)
	} else {
		c2.Gold(t, *update, []byte(actual), "gold/api.json")
	}
}
