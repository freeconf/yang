package car

import "github.com/c2stack/c2g/node"

/////////////////////////
// C A R  N O D E
//  Bridge from model to car app.

// carNode is root handler from car.yang
//    module car { ... }
func Node(c *Car) node.Node {
	return &node.Extend{
		Node: node.ReflectNode(c),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "tire":
				return tiresNode(c.Tire), nil
			default:
				return p.Child(r)
			}
			return nil, nil
		},
		OnAction: func(p node.Node, r node.ActionRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "rotateTires":
				c.rotateTires()
			case "replaceTires":
				c.replaceTires()
			}
			return nil, nil
		},
		OnNotify: func(p node.Node, r node.NotifyRequest) (node.NotifyCloser, error) {
			switch r.Meta.GetIdent() {
			case "update":
				sub := c.onUpdate(func(*Car) {
					r.Send(r.Context, Node(c))
				})
				return sub.Close, nil
			}
			return nil, nil
		},
	}
}

// tiresNode handles list of tires.
//     list tire { ... }
func tiresNode(tires []*tire) node.Node {
	return &node.MyNode{
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			var t *tire
			key := r.Key
			var pos int
			if key != nil {
				pos = key[0].Int
				if pos >= len(tires) {
					return nil, nil, nil
				}
			}
			if key != nil {
				t = tires[pos]
			} else {
				if r.Row < len(tires) {
					t = tires[r.Row]
					key = node.SetValues(r.Meta.KeyMeta(), r.Row)
				}
			}
			if t != nil {
				return tireNode(t), key, nil
			}
			return nil, nil, nil
		},
	}
}

// tireNode handles each tire node.  Everything *inside* list tire { ...}
func tireNode(t *tire) node.Node {
	return &node.Extend{
		Node: node.ReflectNode(t),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "worn":
				hnd.Val = &node.Value{Bool: t.Worn()}
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}
