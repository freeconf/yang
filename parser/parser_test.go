package parser_test

import (
	"fmt"
	"testing"

	"github.com/freeconf/yang/meta"

	"github.com/freeconf/yang/c2"
	"github.com/freeconf/yang/nodes"
	"github.com/freeconf/yang/parser"
)

func TestGroupCircular(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, `module x { revision 0;
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
		t.Error(err)
	}
	a := m.DataDefs()[0].(meta.HasDataDefs)
	c2.AssertEqual(t, "a", a.Ident())
	c2.AssertEqual(t, 2, len(a.DataDefs()))
	c2.AssertEqual(t, "c", a.DataDefs()[0].Ident())
	b := a.DataDefs()[1].(meta.HasDataDefs)
	c2.AssertEqual(t, "b", b.Ident())
	c2.AssertEqual(t, 2, len(b.DataDefs()))
	c2.AssertEqual(t, "d", b.DataDefs()[0].Ident())
	c2.AssertEqual(t, "a", b.DataDefs()[1].Ident())
}

func TestGroupInInput(t *testing.T) {
	_, err := parser.LoadModuleFromString(nil, `module x { revision 0;
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
	m, err := parser.LoadModuleFromString(nil, `module x { revision 0;
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
	c2.AssertEqual(t, "x", m.DataDefs()[0].Ident())
	y := m.DataDefs()[1].(meta.HasDataDefs)
	c2.AssertEqual(t, "y", y.Ident())
	c2.AssertEqual(t, "x", y.DataDefs()[0].Ident())
}

func TestEnum(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, `module x { revision 0;
		leaf l {
			type enumeration {
				enum a;
				enum b {
					value 100;
					description "d";
				}
			}
		}
	}`)
	if err != nil {
		t.Error(err)
	}
	l := m.DataDefs()[0].(meta.HasType)
	c2.AssertEqual(t, "a,b", l.Type().Enum().String())
	c2.AssertEqual(t, "d", l.Type().Enums()[1].Description())
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
		_, err := parser.LoadModuleFromString(nil, y)
		if err == nil {
			t.Error("expected error but didn't get one")
		} else {
			c2.AssertEqual(t, test.err, err.Error())
		}
	}
}

// list is used in lex_more_test.go as well
var yangTestFiles = []struct {
	dir   string
	fname string
}{
	{"/ddef", "container"},
	{"/import", "x"},
	{"/include", "x"},
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
	{"/extension", "x"},
	{"/extension", "y"},
	{"/augment", "x"},
	{"/identity", "x"},
	{"/feature", "x"},
	{"/when", "x"},
	{"/must", "x"},
	{"/choice", "no-case"},
	{"/general", "status"},
	{"", "turing-machine"},
}

func TestParseSamples(t *testing.T) {
	//yyDebug = 4
	ylib := &meta.FileStreamSource{Root: "../yang"}
	yangModule := parser.RequireModule(ylib, "yang")
	for _, test := range yangTestFiles {
		t.Log(test)
		ypath := &meta.FileStreamSource{Root: "testdata" + test.dir}
		features := meta.FeaturesOff([]string{"blacklisted"})
		m, err := parser.LoadModuleWithFeatures(ypath, test.fname, "", features)
		if err != nil {
			t.Error(err)
			continue
		}
		b := nodes.Schema(yangModule, m)
		actual, err := nodes.WritePrettyJSON(b.Root())
		if err != nil {
			t.Error(err)
			continue
		}
		c2.Gold(t, *updateFlag, []byte(actual), "./testdata"+test.dir+"/gold/"+test.fname+".json")
	}
}
