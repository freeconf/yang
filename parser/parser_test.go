package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/source"

	"github.com/freeconf/yang/fc"
)

func TestParseBasic(t *testing.T) {
	_, err := LoadModuleFromString(nil, `
	  module x { 
		  revision 0;
		  namespace "";
		  prefix "";
	}`)
	if err != nil {
		t.Error(err)
	}
}

func TestTokenString(t *testing.T) {
	fc.AssertEqual(t, `whitespace  surrounded`, tokenString(` whitespace  surrounded  `))
	fc.AssertEqual(t, `whitespace  surrounded`, tokenString(` "whitespace  surrounded"  `))
	fc.AssertEqual(t, `whitespace  'surrounded'`, tokenString(` "whitespace  'surrounded'"  `))
}

func TestParseEnum(t *testing.T) {
	m, err := LoadModuleFromString(nil, `module x { revision 0;
		leaf l {
			type enumeration {
				enum a;
				enum b {
					value 100;
					description "b";
				}
				enum "c" {
					description "c";
				}
			}
		}
	}`)
	if err != nil {
		t.Error(err)
	}
	l := m.DataDefinitions()[0].(meta.HasType)
	fc.AssertEqual(t, "a,b,c", l.Type().Enum().String())
	fc.AssertEqual(t, "b", l.Type().Enums()[1].Description())
	fc.AssertEqual(t, "c", l.Type().Enums()[2].Description())
}

func TestParseErr(t *testing.T) {
	tests := []struct {
		y   string
		err string
	}{
		{
			y:   `uses g1;`,
			err: "x/g1 - g1 group not found",
		},
		{
			y:   `container x { uses g1; }`,
			err: "x/x/g1 - g1 group not found",
		},
		{
			y:   `container x { choice z { case q { uses g1; } } }`,
			err: "x/x/z/q/g1 - g1 group not found",
		},
	}
	for _, test := range tests {
		t.Log(test.y)
		y := fmt.Sprintf(`module x { revision 0; %s }`, test.y)
		_, err := LoadModuleFromString(nil, y)
		if err == nil {
			t.Error("expected error but didn't get one")
		} else {
			fc.AssertEqual(t, test.err, err.Error())
		}
	}
}

func TestInvalid(t *testing.T) {
	tests := []struct {
		dir   string
		fname string
		err   string
	}{
		{"/ddef", "config", "config cannot be true when parent config is false"},
		{"/types", "leafref-bad", "path cannot be resolved"},
		{"/types", "leafref-invalid-path", "path cannot be resolved"},
		{"/import", "missing-import", "module not found imp"},
		{"/general", "incomplete", "syntax error"},
	}
	for _, test := range tests {
		ypath := source.Dir("testdata" + test.dir)
		_, err := LoadModule(ypath, test.fname)

		// we verify contents of error because we want to make sure it is failing for the right reason.
		if err == nil {
			t.Error("no error. expected ", test.err)
		} else {
			msg := fmt.Sprintf("got error but unexpected content:\nexpected string: '%s'\n full string: '%s'\n", err.Error(), test.err)
			fc.AssertEqual(t, true, strings.Contains(err.Error(), test.err), msg)
		}
	}
}

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
	{"/typedef", "x"},
	{"/typedef", "import"},
	{"/grouping", "x"},
	{"/grouping", "scope"},
	{"/grouping", "refine"},
	{"/grouping", "augment"},
	{"/grouping", "empty"},
	{"/grouping", "issue-46"},
	{"/extension", "x"},
	{"/extension", "y"},

	// not all the extensions are dumped but at least all extensions are
	// parsed.  lexer test does dump all tokens
	{"/extension", "extreme"},

	// BROKEN!
	// {"/extension", "yin"},

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
	{"/general", "status"},
	{"/general", "rpc-groups"},
	{"/general", "notify-groups"},
	{"/general", "anydata"},

	{"/general", "rpc"},

	{"/deviate", "x"},
	{"", "turing-machine"},
}

// recursive, we can parse it but dumping to json is infinite recursion
// not sure how to represent that yet.
// {"/grouping", "multiple"},

func TestParseSamples(t *testing.T) {
	//yyDebug = 4
	modules := make([]*meta.Module, len(yangTestFiles))
	var err error

	// parse then verify to gold files because we're using yang to
	// dump yang and we have to pass all parsing first
	for i, test := range yangTestFiles {
		t.Log("parse", test)
		ypath := source.Dir("testdata" + test.dir)
		features := meta.FeaturesOff([]string{"blacklisted"})
		modules[i], err = LoadModuleWithOptions(ypath, test.fname, Options{Features: features})
		if err != nil {
			t.Error(err)
			modules[i] = nil
		}
	}

	ylib := source.Dir("../yang")
	yangModule := RequireModule(ylib, "fc-yang")
	wtr := nodeutil.JSONWtr{Pretty: true}
	for i, test := range yangTestFiles {
		t.Log("diff", test)
		if modules[i] == nil {
			continue
		}
		b := nodeutil.Schema(yangModule, modules[i])
		nodeutil.JSONWtr{Pretty: true}.JSON(b.Root())
		actual, err := wtr.JSON(b.Root())
		if err != nil {
			t.Error(err)
			continue
		}
		fc.Gold(t, *updateFlag, []byte(actual), "./testdata"+test.dir+"/gold/"+test.fname+".json")
	}
}

func TestFcYangParse(t *testing.T) {
	// this is a complicated schema and parsing this w/o crashing
	// or going into infinited recursion is worthy test
	ylib := source.Dir("../yang")
	RequireModule(ylib, "fc-yang")
}

// While not allowed as part of RFC, it has major benefits and
// hopefully will be allowed in upcoming YANG specs
func TestGroupCircular(t *testing.T) {
	m, err := LoadModuleFromString(nil, `module x { revision 0;
		grouping g1 {
			container a {
				leaf c {
					type string;
				}
				uses g2;	
			}
		}

		grouping g2 {
			container b {
				leaf d {
					type string;
				}
				uses g1;
			}
		}

		uses g1;
	}`)
	if err != nil {
		t.Fatal(err)
	}
	a := m.DataDefinitions()[0].(meta.HasDataDefinitions)
	fc.AssertEqual(t, "a", a.Ident())
	fc.AssertEqual(t, 2, len(a.DataDefinitions()))
	fc.AssertEqual(t, "c", a.DataDefinitions()[0].Ident())
	b := a.DataDefinitions()[1].(meta.HasDataDefinitions)
	fc.AssertEqual(t, "b", b.Ident())
	fc.AssertEqual(t, 2, len(b.DataDefinitions()))
	fc.AssertEqual(t, "d", b.DataDefinitions()[0].Ident())
	fc.AssertEqual(t, "a", b.DataDefinitions()[1].Ident())
}

func TestGroupInInput(t *testing.T) {
	_, err := LoadModuleFromString(nil, `module x { revision 0;
		grouping g1 {
			leaf x {
				type string;
			}
		}

		rpc y {
			input {
				uses g1;
			}
		}
	}`)
	if err != nil {
		t.Error(err)
	}
}

func TestGroupMultiple(t *testing.T) {
	m, err := LoadModuleFromString(nil, `module x { revision 0;
		grouping g1 {
			leaf x {
				type string;
			}
		}

		uses g1;

		container y {
			uses g1;
		}
	}`)
	if err != nil {
		t.Error(err)
	}
	fc.AssertEqual(t, "x", m.DataDefinitions()[0].Ident())
	y := m.DataDefinitions()[1].(meta.HasDataDefinitions)
	fc.AssertEqual(t, "y", y.Ident())
	fc.AssertEqual(t, "x", y.DataDefinitions()[0].Ident())
}

func TestSymanticallyBadYang(t *testing.T) {
	tests := []struct {
		bad  string
		good string
	}{
		{ // unbalanced regex
			`leaf l {
				type string {
					pattern "x[";
				}
			}`,
			`leaf l {
				type string {
					pattern "x[x]";
				}
			}`,
		},
	}
	for _, test := range tests {
		// we test both good and bad so if ever there was an unrelated systematic error, the good tests
		// would start to fail and we'd catch it here
		bad := fmt.Sprintf(`module y { prefix ""; namespace ""; revision 0; %s }`, test.bad)
		_, err := LoadModuleFromString(nil, bad)
		fc.AssertEqual(t, true, err != nil, test.bad)

		good := fmt.Sprintf(`module y { prefix ""; namespace ""; revision 0; %s }`, test.good)
		_, err = LoadModuleFromString(nil, good)
		fc.AssertEqual(t, true, err == nil, test.good)
	}
}

func TestIdentityDerived(t *testing.T) {
	ypath := source.Path("./testdata/identity")
	m := RequireModule(ypath, "derived-a")
	l := meta.Find(m, "l").(*meta.Leaf)
	i := l.Type().Base()[0]
	fc.AssertEqual(t, 2, len(i.DerivedDirect()))
}
