package node

import (
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/xpath"
)

func Test_Find(t *testing.T) {
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
	}
	`
	m := yang.RequireModuleFromString(nil, mstr)
	b := NewBrowser(m, ReadJson(`{"a":{"b":10},"aa":{"bb":"hello"}}`))
	tests := []struct {
		path     string
		expected string
	}{
		{
			path:     `a/b<20`,
			expected: `{"b":10}`,
		},
		{
			path: `a/b<2`,
		},
		{
			path: `a/b!=10`,
		},
		{
			path:     `a/b=10`,
			expected: `{"b":10}`,
		},
		{
			path:     `aa/bb='hello'`,
			expected: `{"bb":"hello"}`,
		},
		{
			path: `aa/bb!='hello'`,
		},
	}
	for _, test := range tests {
		p, err := xpath.Parse(test.path)
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
				actual, _ := WriteJson(s)
				if notEqual := c2.CheckEqual(test.expected, actual); notEqual != nil {
					t.Error(notEqual)
				}
			}
		} else if !s.IsNil() {
			actual, _ := WriteJson(s)
			t.Error("expected not found but got ", actual)
		}
	}
}
