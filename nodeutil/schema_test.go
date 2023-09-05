package nodeutil_test

import (
	"flag"
	"testing"

	"github.com/freeconf/yang"
	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"
)

var updateFlag = flag.Bool("update", false, "Update the golden files.")

func TestSchemaRead(t *testing.T) {
	ymod := parser.RequireModule(yang.InternalYPath, "fc-yang")
	tests := []string{
		// "json-test",
		// "choice",
		"ext",
	}
	for _, test := range tests {
		m := parser.RequireModule(source.Dir("./testdata"), test)
		sel := nodeutil.SchemaBrowser(ymod, m).Root()
		actual, err := nodeutil.WritePrettyJSON(sel)
		if err != nil {
			t.Error(err)
		}
		fc.Gold(t, *updateFlag, []byte(actual), "gold/"+test+".schema.json")
	}
}
