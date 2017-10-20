package nodes_test

import (
	"flag"
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/nodes"
)

var updateFlag = flag.Bool("update", false, "Update the golden files.")

func TestSchemaRead(t *testing.T) {
	m := yang.RequireModule(&meta.FileStreamSource{Root: "./testdata"}, "json-test")
	ypath := &meta.FileStreamSource{Root: "../yang"}
	sel := nodes.SchemaWithYangPath(ypath, m, false).Root()
	actual, err := nodes.WritePrettyJSON(sel)
	if err != nil {
		t.Error(err)
	}
	c2.Gold(t, *updateFlag, []byte(actual), "gold/json-test.schema.json")
}
