package secure

import (
	"testing"

	"github.com/c2stack/c2g/c2"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
)

func TestManage(t *testing.T) {
	a := NewRbac()
	ypath := &meta.FileStreamSource{Root: "../yang"}
	b := node.NewBrowser(yang.RequireModule(ypath, "secure"), Manage(a))
	err := b.Root().UpsertFrom(nodes.ReadJSON(`{
		"authorization" : {
			"role" : [{
				"id" : "sales",
				"access" : [{
					"path" : "m",
					"perm" : "read"
				},{
					"path" : "m/x",
					"perm" : "none"
				},{
					"path" : "m/z",
					"perm" : "full"				
				}]			
			}]	
		}
	}`)).LastErr
	if err != nil {
		t.Fatal(err)
	}
	c2.AssertEqual(t, 1, len(a.Roles))
	//c2.AssertEqual(t, 3, len(a.Roles["sales"].Access))
}
