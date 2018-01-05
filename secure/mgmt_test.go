package secure

import (
	"testing"

	"github.com/freeconf/gconf/c2"

	"github.com/freeconf/gconf/meta"
	"github.com/freeconf/gconf/meta/yang"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/nodes"
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
