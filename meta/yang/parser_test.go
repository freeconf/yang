package yang_test

import (
	"testing"

	"github.com/c2stack/c2g/meta"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/nodes"
)

var yangTestFiles = []struct {
	dir   string
	fname string
}{
	{"/import", "x"},
	{"/include", "x"},
	{"/types", "anydata"},
	{"/types", "enum"},
	{"/types", "container"},
	{"/types", "leaf"},
	{"/types", "typedef"},
	{"/grouping", "x"},
	{"/grouping", "scope"},
	{"/grouping", "refine"},
	{"/extension", "x"},
	{"/extension", "y"},
	{"", "turing-machine"},
}

func TestParseSamples(t *testing.T) {
	//yyDebug = 4
	ylib := &meta.FileStreamSource{Root: "../../yang"}
	for _, test := range yangTestFiles {
		t.Log(test)
		ypath := &meta.FileStreamSource{Root: "testdata" + test.dir}
		m, err := yang.LoadModule(ypath, test.fname)
		if err != nil {
			t.Error(err)
			continue
		}
		b := nodes.SchemaWithYangPath(ylib, m, false)
		actual, err := nodes.WritePrettyJSON(b.Root())
		if err != nil {
			t.Error(err)
			continue
		}
		c2.Gold(t, *updateFlag, []byte(actual), "./testdata"+test.dir+"/gold/"+test.fname+".parse.json")

		b = nodes.SchemaWithYangPath(ylib, m, true)
		actual, err = nodes.WritePrettyJSON(b.Root())
		if err != nil {
			t.Error(err)
			continue
		}
		c2.Gold(t, *updateFlag, []byte(actual), "./testdata"+test.dir+"/gold/"+test.fname+".resolve.json")
	}
}
