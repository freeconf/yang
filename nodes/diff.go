package nodes

import (
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"
)

// Diff compares two nodes and returns only the difference. Use Selection constraints
// to control what is compared and how deep.
func Diff(a, b node.Node) node.Node {
	return &Basic{
		OnChild: func(r node.ChildRequest) (n node.Node, err error) {
			var aNode, bNode node.Node
			r.New = false
			if aNode, err = a.Child(r); err != nil {
				return nil, err
			}
			if bNode, err = b.Child(r); err != nil {
				return nil, err
			}
			if aNode == nil {
				return nil, nil
			}
			if bNode == nil {
				return aNode, nil
			}
			return Diff(aNode, bNode), nil
		},
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var err error
			var aNode, bNode node.Node
			var aKey []val.Value
			r.New = false
			if aNode, aKey, err = a.Next(r); err != nil {
				return nil, nil, err
			}
			if bNode, _, err = b.Next(r); err != nil {
				return nil, nil, err
			}
			if aNode == nil {
				return nil, nil, nil
			}
			if bNode == nil {
				return aNode, aKey, nil
			}
			return Diff(aNode, bNode), aKey, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
			if err = a.Field(r, hnd); err != nil {
				return err
			}
			aVal := hnd.Val
			if err = b.Field(r, hnd); err != nil {
				return err
			}
			bVal := hnd.Val
			if aVal == nil {
				if bVal == nil {
					return nil
				}
				hnd.Val = bVal
				return nil
			}
			if val.Equal(aVal, bVal) {
				hnd.Val = nil
				return nil
			}
			hnd.Val = aVal
			return nil
		},
	}
}
