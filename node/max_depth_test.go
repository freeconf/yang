package node_test

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/val"
)

var depthMstr = `module x {

	container one {
		container two {
			list three {
				container four {
					leaf fivea {
						type string;
					}
					container fiveb {
						leaf six {
							type string;
						}
					}
				}
			}
		}
	}
}
`

func TestMaxDepth(t *testing.T) {
	n := &nodeutil.Basic{
		OnField: func(fr node.FieldRequest, vh *node.ValueHandle) error {
			vh.Val = val.String("here")
			return nil
		},
	}
	n.OnChild = func(r node.ChildRequest) (child node.Node, err error) {
		return n, nil
	}
	n.OnNext = func(r node.ListRequest) (next node.Node, key []val.Value, err error) {
		if r.Row == 0 {
			return n, nil, nil
		}
		return nil, nil, nil
	}
	m, err := parser.LoadModuleFromString(nil, depthMstr)
	fc.AssertEqual(t, nil, err)
	b := node.NewBrowser(m, n)
	root := b.Root()
	tests := []struct {
		path     string
		expected string
	}{
		{
			path:     "?depth=1",
			expected: `{"one":{}}`,
		},
		{
			path:     "?depth=2",
			expected: `{"one":{"two":{}}}`,
		},
		{
			path:     "?depth=3",
			expected: `{"one":{"two":{"three":[{}]}}}`,
		},
		{
			path:     "?depth=4",
			expected: `{"one":{"two":{"three":[{"four":{}}]}}}`,
		},
		{
			path:     "?depth=5",
			expected: `{"one":{"two":{"three":[{"four":{"fivea":"here","fiveb":{}}}]}}}`,
		},
		{
			path:     "?depth=6",
			expected: `{"one":{"two":{"three":[{"four":{"fivea":"here","fiveb":{"six":"here"}}}]}}}`,
		},
	}
	for _, test := range tests {
		actual, err := nodeutil.WriteJSON(root.Find(test.path))
		fc.AssertEqual(t, nil, err)
		fc.AssertEqual(t, test.expected, actual, test.path)
	}
}
