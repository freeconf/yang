package yang_test

import (
	"flag"
	"io/ioutil"
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/nodes"
)

var update = flag.Bool("update2", false, "update gold files instead of verify against them")

func TestParseExamples(t *testing.T) {
	//yyDebug = 4
	y, _ := ioutil.ReadFile("./testdata/turing-machine.yang")
	m, err := yang.LoadModuleCustomImport(string(y), nil)
	if err != nil {
		t.Error(err)
	} else {
		b := nodes.Schema(m, false)
		actual, err := nodes.WritePrettyJSON(b.Root())
		if err != nil {
			t.Error(err)
		} else {
			c2.Gold(t, *update, []byte(actual), "./gold/turing-machine.parse.json")
		}
	}
}
