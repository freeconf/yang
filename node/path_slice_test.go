package node

import (
	"fmt"
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/parser"
)

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
	m, err := parser.LoadModuleFromString(nil, pathTestModule)
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
	}
	for _, test := range tests {
		p, err := parseUrlPath(test.in, m)
		fc.AssertEqual(t, nil, err, test.in)
		fc.AssertEqual(t, len(test.expected), len(p), test.in)
		for i, e := range test.expected {
			if e != p[i].Meta.Ident() {
				msg := "expected to find \"%s\" as segment number %d in \"%s\" but got \"%s\""
				t.Errorf(msg, e, i, test.in, p[i].Meta.Ident())
			}
		}
	}
}

func TestPathSegmentKeys(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, pathTestModule)
	if err != nil {
		t.Fatal(err)
	}
	none := []interface{}{}
	tests := []struct {
		in       string
		expected [][]interface{}
	}{
		{"a/b", [][]interface{}{none, none}},
		{"a/b=c/e", [][]interface{}{none, []interface{}{"c"}, none}},
		{"x=9", [][]interface{}{[]interface{}{9}}},
		{"a/b=c%2fc/e", [][]interface{}{none, []interface{}{"c/c"}, none}},
	}
	for _, test := range tests {
		segments, e := parseUrlPath(test.in, m)
		if e != nil {
			t.Errorf("Error parsing %s : %s", test.in, e)
		}
		if len(test.expected) != len(segments) {
			t.Error("wrong number of expected segments for", test.in)
		}
		for i, expected := range test.expected {
			for j, key := range expected {
				if segments[i].Key[j].Value() != key {
					desc := fmt.Sprintf("\"%s\" segment \"%s\" - expected \"%s\" - got \"%s\"",
						test.in, segments[i].Meta.Ident(), key, segments[i].Key[j])
					t.Error(desc)
				}
			}
		}
	}
}

func TestPathEmpty(t *testing.T) {
	p, err := parseUrlPath("", &meta.Container{})
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, 0, len(p))
}
