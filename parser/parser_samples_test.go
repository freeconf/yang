package parser

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/source"
)

// list is used in lex_more_test.go as well
var yangTestFiles = []struct {
	dir   string
	fname string
}{
	{"/ddef", "container"},
	{"/ddef", "assort"},
	{"/ddef", "unique"},
	{"/import", "x"},
	{"/import", "example-barmod"},
	{"/include", "x"},
	{"/include", "top"},
	{"/types", "anydata"},
	{"/types", "enum"},
	{"/types", "container"},
	{"/types", "leaf"},
	{"/types", "union"},
	{"/types", "leafref"},
	{"/types", "leafref-i1"},
	{"/types", "union-units"},
	{"/types", "bits"},
	{"/typedef", "x"},
	{"/typedef", "typedef-x"},
	{"/typedef", "import"},
	{"/grouping", "x"},
	{"/grouping", "scope"},
	{"/grouping", "refine"},
	{"/grouping", "augment"},
	{"/grouping", "empty"},
	{"/grouping", "issue-46"},
	{"/grouping", "refine-default"},
	{"/grouping", "recurse-1"},
	{"/grouping", "recurse-2"},
	{"/grouping", "recurse-3"},
	{"/extension", "x"},
	{"/extension", "y"},
	{"/extension", "yin"},

	// not all the extensions are dumped but at least all extensions are
	// parsed.  lexer test does dump all tokens
	{"/extension", "extreme"},

	{"/augment", "x"},
	{"/augment", "aug-with-uses"},
	{"/augment", "aug-choice"},
	{"/identity", "x"},
	{"/feature", "x"},
	{"/when", "x"},
	{"/must", "x"},
	{"/choice", "no-case"},
	{"/choice", "choice-mandatory"},
	{"/choice", "choice-default"},
	{"/choice", "choice-x"},
	{"/general", "status"},
	{"/general", "rpc-groups"},
	{"/general", "notify-groups"},
	{"/general", "anydata"},

	{"/general", "rpc"},

	{"/deviate", "x"},

	{"", "turing-machine"},
	{"", "basic_config2"},
}

// recursive, we can parse it but dumping to json is infinite recursion
// not sure how to represent that yet.
// {"/grouping", "multiple"},

func TestParseSamples(t *testing.T) {
	//yyDebug = 4
	//fc.DebugLog(true)
	modules := make([]*meta.Module, len(yangTestFiles))
	var err error

	// parse then verify to gold files because we're using yang to
	// dump yang and we have to pass all parsing first
	failed := false
	for i, test := range yangTestFiles {
		t.Log("parse", test)
		ypath := source.Dir("testdata" + test.dir)
		features := meta.FeaturesOff([]string{"blacklisted"})
		modules[i], err = LoadModuleWithOptions(ypath, test.fname, Options{Features: features})
		if err != nil {
			t.Error(err)
			failed = true
			modules[i] = nil
		}
	}
	if failed {
		return
	}

	ylib := source.Dir("../yang")
	yangModule := RequireModule(ylib, "fc-yang")
	wtr := nodeutil.JSONWtr{Pretty: true}
	for i, test := range yangTestFiles {
		t.Log("diff", test)
		if modules[i] == nil {
			continue
		}
		b := nodeutil.SchemaBrowser(yangModule, modules[i])
		nodeutil.JSONWtr{Pretty: true}.JSON(b.Root())
		actual, err := wtr.JSON(b.Root())
		if err != nil {
			t.Error(err)
			continue
		}
		fc.Gold(t, *updateFlag, []byte(actual), "./testdata"+test.dir+"/gold/"+test.fname+".json")
	}
}
