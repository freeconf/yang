package browse

import (
	"os"
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
)

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
	if err := c2.CheckEqual("x-y", doc.Defs[0].Meta.GetIdent()); err != nil {
		t.Error(err)
		t.Log(doc.Defs[0])
	}
	if err := c2.CheckEqual("a-b", doc.Defs[1].Meta.GetIdent()); err != nil {
		t.Error(err)
		t.Log(doc.Defs[1])
	}
	if err := c2.CheckEqual(3, len(doc.Defs[0].Fields)); err != nil {
		t.Error(err)
	} else {
		if err2 := c2.CheckEqual("y1", doc.Defs[0].Fields[1].Case.GetIdent()); err2 != nil {
			t.Error(err2)
		}
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
		f := escape(test.in, "x")
		actual := f("Bingo was his name-o")
		if actual != test.expected {
			t.Error(actual, test.expected)
		}
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
		buff, err := os.Create("testdata/.doc-example." + test.Ext)
		if err != nil {
			t.Fatal(err)
		}
		if err := test.Builder.Generate(d, buff); err != nil {
			t.Error(err)
		}
		buff.Close()
	}
}
