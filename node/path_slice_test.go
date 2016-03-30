package node

import (
	"fmt"
	"testing"
	"github.com/blitter/meta"
	"github.com/blitter/meta/yang"
)

func TestPathEmpty(t *testing.T) {
	p, e := ParsePath("", &meta.Container{})
	if e != nil {
		t.Error(e)
	}
	if p.Len() != 0 {
		t.Errorf("expected no segments, got %d", p.Len())
	}
	if !p.Empty() {
		t.Errorf("expected empty path")
	}

}

func TestPathSliceSplit(t *testing.T) {
	m, err := yang.LoadModuleCustomImport(pathTestModule, nil)
	if err != nil {
		t.Fatal(err)
	}
	p, e := ParsePath("a/b=y/e", m)
	if e != nil {
		t.Fatal(e)
	}
	frag := p.SplitAfter(p.Tail.Parent().Parent())
	actual := frag.String()
	if actual != "b=y/e" {
		t.Error(actual)
	}
}

func TestPathPopHead(t *testing.T) {
	m, err := yang.LoadModuleCustomImport(pathTestModule, nil)
	if err != nil {
		t.Fatal(err)
	}
	p, e := ParsePath("a/b=y/e", m)
	if e != nil {
		t.Fatal(e)
	}
	b := p.PopHead().PopHead()
	if b.Head.meta.GetIdent() != "b" {
		t.Error(b.Head.meta.GetIdent())
	}
}

func TestPathStringAndEqual(t *testing.T) {
	m, err := yang.LoadModuleCustomImport(pathTestModule, nil)
	if err != nil {
		t.Fatal(err)
	}
	tests := []string {
		"",
		"a/b",
		"a/b=x",
		"a/b=y/e",
		"x=9",
	}
	for _, test := range tests {
		p, e := ParsePath(test, m)
		if e != nil {
			t.Error(e)
		}
		actual := p.String()
		if test != actual {
			t.Errorf("\nExpected: '%s'\n  Actual:'%s'", test, actual)
		}

		// Test equals
		p2, _ := ParsePath(test, m)
		if ! p.Equal(p2) {
			t.Errorf("%s does not equal itself", test)
		}
	}
}

var pathTestModule = `
module m {
	prefix "";
	namespace "";
	revision 0;
	container a {
		list b {
			key "d";
			leaf d {
				type string;
			}
			container e {
				leaf g {
					type string;
				}
			}
			leaf f {
				type string;
			}
		}
	}
	list x {
		key "y";
		leaf y {
			type int32;
		}
	}
}
`
func TestPathSegment(t *testing.T) {
	m, err := yang.LoadModuleCustomImport(pathTestModule, nil)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		in       string
		expected []string
	}{
		{"a/b", []string{"a", "b"}},
		{"a/b=x", []string{"a", "b"}},
		{"a/b=y/e", []string{"a", "b", "e"}},
		{"a/b?foo=1", []string{"a", "b"}},
	}
	for _, test := range tests {
		p, e := ParsePath(test.in, m)
		if e != nil {
			t.Errorf("Error parsing %s : %s", test.in, e)
		}
		if len(test.expected) != p.Len() {
			t.Errorf("Expected %d segments for %s but got %d", len(test.expected), test.in, p.Len())
		}
		segments := p.Segments()
		for i, e := range test.expected {
			if e != segments[i].Meta().GetIdent() {
				msg := "expected to find \"%s\" as segment number %d in \"%s\" but got \"%s\""
				t.Error(fmt.Sprintf(msg, e, i, test.in, segments[i].Meta().GetIdent()))
			}
		}
	}
}

func TestPathSegmentKeys(t *testing.T) {
	m, err := yang.LoadModuleCustomImport(pathTestModule, nil)
	if err != nil {
		t.Fatal(err)
	}
	none := []interface{}{}
	tests := []struct {
		in       string
		expected [][]interface{}
	}{
		{"a/b", [][]interface{}{ none, none}},
		{"a/b=c/e", [][]interface{}{ none, []interface{}{"c"}, none}},
		{"x=9", [][]interface{}{[]interface{}{9}}},
	}
	for _, test := range tests {
		p, e := ParsePath(test.in, m)
		if e != nil {
			t.Errorf("Error parsing %s : %s", test.in, e)
		}
		if len(test.expected) != p.Len() {
			t.Error("wrong number of expected segments for", test.in)
		}
		segments := p.Segments()
		for i, expected := range test.expected {
			for j, key := range expected {
				if segments[i].Key()[j].Value() != key {
					desc := fmt.Sprintf("\"%s\" segment \"%s\" - expected \"%s\" - got \"%s\"",
						test.in, segments[i].Meta().GetIdent(), key, segments[i].Key()[j])
					t.Error(desc)
				}
			}
		}
	}
}

