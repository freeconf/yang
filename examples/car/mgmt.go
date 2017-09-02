package car

import (
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
	"github.com/c2stack/c2g/val"
)

/////////////////////////
// C A R  N O D E
//  Bridge from model to car app.

// carNode is root handler from car.yang
//    module car { ... }
func Manage(c *Car) node.Node {

	// Powerful combination, we're letting reflect do a lot of the CRUD
	// when the yang file matches the field names.  But we extend reflection
	// to add as much custom behavior as we want
	return &nodes.Extend{
		Base: nodes.ReflectChild(c),

		// drilling into child objects defined by yang file
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "tire":
				return tiresNode(c.Tire), nil
			case "specs":
				// knows how to r/w config from a map
				return nodes.ReflectChild(c.Specs), nil
			default:
				// return control back to handler we're extending, in this case
				// it's reflection
				return p.Child(r)
			}
			return nil, nil
		},

		// RPCs
		OnAction: func(p node.Node, r node.ActionRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "rotateTires":
				c.rotateTires()
			case "replaceTires":
				c.replaceTires()
			}
			return nil, nil
		},

		// Events
		OnNotify: func(p node.Node, r node.NotifyRequest) (node.NotifyCloser, error) {
			switch r.Meta.GetIdent() {
			case "update":
				// very easy bridging from
				sub := c.OnUpdate(func(*Car) {

					// cleverly reuses node handler to send car data
					r.Send(Manage(c))

				})

				// NOTE: we return a close function, we are not actually closing
				// here
				return sub.Close, nil
			}
			return nil, nil
		},

		// override OnEndEdit just to just to know when car has been creates and
		// fully initialized so we can start the car running
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			// allow reflection node handler to finish, this is where defaults
			// get set.
			if err := p.EndEdit(r); err != nil {
				return err
			}
			c.Start()
			return nil
		},
	}
}

// tiresNode handles list of tires.
//     list tire { ... }
func tiresNode(tires []*tire) node.Node {
	return &nodes.Basic{
		// Handling lists are
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var t *tire
			key := r.Key
			var pos int

			// request for specific item in list
			if key != nil {
				pos = key[0].Value().(int)
				if pos >= len(tires) {
					return nil, nil, nil
				}
			}
			if key != nil {
				t = tires[pos]
			} else {
				// request for nth item in list
				if r.Row < len(tires) {
					t = tires[r.Row]
					key = []val.Value{val.Int32(r.Row)}
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

	// Again, let reflection do a lot of the work
	return &nodes.Extend{
		Base: nodes.ReflectChild(t),

		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {

			case "worn":
				// worn is a method call, so our current reflection handler doesn't
				// check for that.  Maybe you reflection handler would.
				hnd.Val = val.Bool(t.Worn())

			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}
