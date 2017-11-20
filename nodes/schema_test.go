package nodes_test

import (
	"flag"
	"testing"

	"github.com/freeconf/c2g/c2"
	"github.com/freeconf/c2g/meta"
	"github.com/freeconf/c2g/meta/yang"
	"github.com/freeconf/c2g/nodes"
)

var updateFlag = flag.Bool("update", false, "Update the golden files.")

func TestSchemaRead(t *testing.T) {
	m := yang.RequireModule(&meta.FileStreamSource{Root: "./testdata"}, "json-test")
	ypath := &meta.FileStreamSource{Root: "../yang"}
	ymod := yang.RequireModule(ypath, "yang")
	sel := nodes.Schema(ymod, m).Root()
	actual, err := nodes.WritePrettyJSON(sel)
	if err != nil {
		t.Error(err)
	}
	c2.Gold(t, *updateFlag, []byte(actual), "gold/json-test.schema.json")
}
