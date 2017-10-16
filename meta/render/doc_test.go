package render

import (
	"bytes"
	"flag"
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
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
	m, err := yang.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	doc := &Doc{}
	if doc.Build(m); doc.LastErr != nil {
		t.Fatal(doc.LastErr)
	}
	if !c2.AssertEqual(t, "x-y", doc.Defs[0].Meta.GetIdent()) {
		t.Log(doc.Defs[0])
	}
	if !c2.AssertEqual(t, "a-b", doc.Defs[1].Meta.GetIdent()) {
		t.Log(doc.Defs[1])
	}
	if c2.AssertEqual(t, 3, len(doc.Defs[0].Fields)) {
		c2.AssertEqual(t, "y1", doc.Defs[0].Fields[1].Case.GetIdent())
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
		c2.AssertEqual(t, test.expected, actual)
	}
}

func Test_DocBuiltIns(t *testing.T) {
	tests := []struct {
		Builder DocDefBuilder
		Ext     string
	}{
		{
			Builder: &DocMarkdown{},
			Ext:     "md",
		},
		{
			Builder: &DocHtml{},
			Ext:     "html",
		},
		{
			Builder: &DocDot{},
			Ext:     "dot",
		},
	}

	m := yang.RequireModule(&meta.FileStreamSource{Root: "testdata"}, "doc-example")
	d := &Doc{
		Title: "example",
	}
	d.Build(m)
	for _, test := range tests {
		t.Log(test.Ext)
		var buff bytes.Buffer
		if err := test.Builder.Generate(d, &buff); err != nil {
			t.Error(err)
		}
		c2.Gold(t, *update, buff.Bytes(), "gold/doc-example."+test.Ext)
	}
}
