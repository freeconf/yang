package nodes_test

import (
	"flag"
	"testing"

	"github.com/freeconf/yang/c2"
	"github.com/freeconf/yang/nodes"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"
)

var updateFlag = flag.Bool("update", false, "Update the golden files.")

func TestSchemaRead(t *testing.T) {
	ypath := source.Dir("../yang")
	ymod := parser.RequireModule(ypath, "fc-yang", "")
	tests := []string{
		"json-test",
		"choice",
	}
	for _, test := range tests {
		m := parser.RequireModule(source.Dir("./testdata"), test, "")
		sel := nodes.Schema(ymod, m).Root()
		actual, err := nodes.WritePrettyJSON(sel)
		if err != nil {
			t.Error(err)
		}
		c2.Gold(t, *updateFlag, []byte(actual), "gold/"+test+".schema.json")
	}
}
