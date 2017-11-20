package node_test

import (
	"testing"

	"github.com/freeconf/c2g/c2"
	"github.com/freeconf/c2g/meta/yang"
	"github.com/freeconf/c2g/node"
	"github.com/freeconf/c2g/nodes"
	"github.com/freeconf/c2g/xpath"
)

func TestXFind(t *testing.T) {
	c2.DebugLog(true)
	mstr := ` module m { namespace ""; prefix ""; revision 0; 
		container a {
			leaf b {
				type int32;
			}
		}
		container aa {
			leaf bb {
				type string;
			}
		}
		list list {
			leaf leaf {
				type int32;
			} 
		}
	}
	`
	m := yang.RequireModuleFromString(nil, mstr)
	b := node.NewBrowser(m, nodes.ReadJSON(`{
		"a":{"b":10},
		"aa":{"bb":"hello"},
		"list":[{"leaf":99},{"leaf":100}]
	}`))
	tests := []struct {
		xpath    string
		expected string
	}{
		{
			xpath:    `a/b<20`,
			expected: `{"b":10}`,
		},
		{
			xpath: `a/b<2`,
		},
		{
			xpath: `a/b!=10`,
		},
		{
			xpath:    `a/b=10`,
			expected: `{"b":10}`,
		},
		{
			xpath:    `aa/bb='hello'`,
			expected: `{"bb":"hello"}`,
		},
		{
			xpath:    `list/leaf=99`,
			expected: `{"leaf":99}`,
		},
		{
			xpath: `aa/bb!='hello'`,
		},
	}
	for _, test := range tests {
		p, err := xpath.Parse(test.xpath)
		if err != nil {
			t.Error(err)
		}
		s := b.Root().XFind(p)
		if s.LastErr != nil {
			t.Error(s.LastErr)
		} else if test.expected != "" {
			if s.IsNil() {
				t.Error("not found but expected to find ", test.expected)
			} else {
				actual, _ := nodes.WriteJSON(s)
				c2.AssertEqual(t, test.expected, actual)
			}
		} else if !s.IsNil() {
			actual, _ := nodes.WriteJSON(s)
			t.Errorf("expected no results from %s but found %s", test.xpath, actual)
		}
	}
}
