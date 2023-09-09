package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/freeconf/yang/meta"
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
		{
			y:   `container c { leaf l1 { type int32; } leaf l1 { type int32; } }`,
			err: "conflict adding",
		},
	}
	for _, test := range tests {
		t.Log(test.y)
		y := fmt.Sprintf(`module x { revision 0; %s }`, test.y)
		_, err := LoadModuleFromString(nil, y)
		if err == nil {
			t.Error("expected error but didn't get one")
		} else {
			fc.AssertEqual(t, true, strings.Contains(err.Error(), test.err), err.Error())
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
		{"/types", "leaf-dup", "conflict adding add leaf-root to root-container"},
		{"/choice", "choice-conflict", "conflict adding add leaf-root to root-container"},
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

func TestFcYangParse(t *testing.T) {
	// this is a complicated schema and parsing this w/o crashing
	// or going into infinited recursion is worthy test
	ylib := source.Dir("../yang")
	RequireModule(ylib, "fc-yang")
}

// While not allowed as part of RFC, it has major benefits and
// hopefully will be allowed in upcoming YANG specs
//
//	     a
//	c        b
//	       d   a  <- recursive ...
func TestRecurse(t *testing.T) {
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

	ab := a.DataDefinitions()[1].(meta.HasDataDefinitions)
	fc.AssertEqual(t, "b", ab.Ident())
	fc.AssertEqual(t, 2, len(ab.DataDefinitions()))
	fc.AssertEqual(t, "d", ab.DataDefinitions()[0].Ident())

	aba := ab.DataDefinitions()[1].(meta.HasDataDefinitions)
	fc.AssertEqual(t, "a", aba.Ident())
	fc.AssertEqual(t, a, aba)

	abab := aba.Definition("b")
	fc.AssertEqual(t, ab, abab)
}

func TestFcYang(t *testing.T) {
	// this is a complicated schema and parsing this w/o crashing
	// or going into infinited recursion is worthy test
	ylib := source.Dir("../yang")
	RequireModule(ylib, "fc-yang")
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
