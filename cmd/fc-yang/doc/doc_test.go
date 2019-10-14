package doc

import (
	"bytes"
	"flag"
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"
)

var update = flag.Bool("update", false, "update gold files")

func TestDocBuild(t *testing.T) {
	mstr := `module x-y {
	namespace "";
	prefix "";
	revision 0;
	container a-b {
		leaf z {
			type string;
		}
	}

	choice x {
		case y1 {
			leaf z1 {
				type string;
			}
		}
		case y2 {
			leaf z2 {
				type string;
			}
		}
	}
}`
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	d := &doc{}
	if d.build(m); d.LastErr != nil {
		t.Fatal(d.LastErr)
	}
	if !fc.AssertEqual(t, "x-y", d.Defs[0].Meta.Ident()) {
		t.Log(d.Defs[0])
	}
	if !fc.AssertEqual(t, "a-b", d.Defs[1].Meta.Ident()) {
		t.Log(d.Defs[1])
	}
	if fc.AssertEqual(t, 3, len(d.Defs[0].Fields)) {
		fc.AssertEqual(t, "y1", d.Defs[0].Fields[1].Case.Ident())
	}
}

func TestEscape(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{
			in:       "a",
			expected: "Bingo wxas his nxame-o",
		},
		{
			in:       "ao",
			expected: "Bingxo wxas his nxame-xo",
		},
	}
	for _, test := range tests {
		actual := escape(test.in, "x")("Bingo was his name-o")
		fc.AssertEqual(t, test.expected, actual)
	}
}

func TestDocBuiltIns(t *testing.T) {
	tests := []struct {
		builder builder
		ext     string
	}{
		{
			builder: &markdown{},
			ext:     "md",
		},
		{
			builder: &html{},
			ext:     "html",
		},
		{
			builder: &dot{},
			ext:     "dot",
		},
	}

	m := parser.RequireModule(source.Dir("testdata"), "doc-example")
	d := &doc{
		Title: "example",
	}
	d.build(m)
	for _, test := range tests {
		t.Log(test.ext)
		var buff bytes.Buffer
		tmpl := test.builder.builtinTemplate()
		if err := test.builder.generate(d, tmpl, &buff); err != nil {
			t.Error(err)
		}
		fc.Gold(t, *update, buff.Bytes(), "gold/doc-example."+test.ext)
	}
}
