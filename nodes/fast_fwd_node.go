package nodes

import (
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"
)

// FastForwardingNode will always claim to have children for requests that follow
// a specific target path. Kinds like fast forwarding to a specific spot.  Useful
// in a variety of cases including proxying.
func FastForwardingNode() node.Node {
	e := &Basic{}
	e.OnChild = func(r node.ChildRequest) (node.Node, error) {
		return e, nil
	}
	e.OnNext = func(r node.ListRequest) (node.Node, []val.Value, error) {
		return e, r.Key, nil
	}
	return e
}
