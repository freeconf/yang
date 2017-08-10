package examples

import (
	"reflect"
	"testing"

	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
	"github.com/c2stack/c2g/val"
)

// MyNode is for complete custom handling.  While not used as much as nodes.Extend or nodes.Reflect
// it is the building block of many node handlers
func Example_01MyNode() {

	// auto management YANG fragment:
	// =========================
	//   module auto {
	//
	//      container engine {
	//          ...
	//      }
	//   }
	//   ...
	//

	// Manage auto
	_ = func(auto *auto) node.Node {

		return &nodes.Basic{

			// containers and list handlers
			OnChild: func(r node.ChildRequest) (node.Node, error) {
				switch r.Meta.GetIdent() {

				// app can have many containers and lists, here we implement the handler
				// just for "container engine {}"
				case "engine":
					if r.New {
						// request to create engine
						auto.Engine = &engine{}

					} else if r.Delete {
						// request to remove engine
						auto.Engine = nil
					}
					if auto.Engine != nil {
						// if still have an engine left, return engine wrapped in engine node handler
						// all remaining engine requests are delegated to engineNode
						return engineNode(auto.Engine), nil
					}
				}

				// this means item (engine, brakes, tires, etc) is not available
				return nil, nil
			},
		}
	}
}

// Extend lets your reuse another node handler for default implementation while leaving
// you a change to write your own handlers for whatever is needed
func Example_02Extend() {

	// individual tire management YANG
	// =========================
	//   ...
	//   container tire {
	//
	//      leaf size {
	//          type string;
	//      }
	//
	//      leaf pressure {
	//          type string;
	//          config "false";
	//      }
	//      ...
	//

	// Manage tire list
	_ = func(tire *tire) node.Node {

		return &nodes.Extend{

			// In this, we'll use reflection for anything we don't implement here.
			Base: nodes.Reflect(tire),

			// leafs, leaf-lists and anydata
			OnField: func(parent node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
				switch r.Meta.GetIdent() {

				//  marked as a "read-only"" field in YANG (config "false") so we only need to
				// to implement reader.
				case "pressure":
					hnd.Val = val.Int32(tire.pressure())

				// Delegate all else to reflection
				default:
					return parent.Field(r, hnd)

				}
				return nil
			},
		}
	}
}

// MapNode is used here to hold data in nested map.  This is useful when there's little reason
// to marshal model data into structs.  You can always decide to change into structs later.  MapNode
// can enable fast prototyping.
func Example_03MapNode(t *testing.T) {

	// engine management YANG:
	// =========================
	// ...
	//   container tire {
	//
	//      leaf size {
	//          type string;
	//      }
	//
	//      leaf pressure {
	//          type string;
	//          config "false";
	//      }
	//      ...
	//
	//

	// Manage engine
	_ = func(engine *engine) node.Node {

		return &nodes.Basic{
			OnChild: func(r node.ChildRequest) (node.Node, error) {
				switch r.Meta.GetIdent() {

				case "specs":
					if r.New {
						engine.Specs = make(map[string]interface{})
					} else if r.Delete {
						engine.Specs = nil
					}
					if engine.Specs != nil {
						return nodes.Reflect(engine.Specs), nil
					}
				}

				return nil, nil
			},
		}
	}
}

// MapNode is used here to hold data in nested map.  This is useful when there's little reason
// to marshal model data into structs.  You can always decide to change into structs later.  MapNode
// can enable fast prototyping.
func Example_03List(t *testing.T) {

	//  cup holder management YANG:
	// =========================
	// ...
	//   list cupHolders {
	//
	//      key "location";
	//
	//      leaf location {
	//          type string;
	//      }
	//
	//      ...

	// Manage cupHolders
	_ = func(holders map[string]*cupHolder) node.Node {

		// helper to navigate "list" requests from a map. This is mostly neccessary because you cannot
		// build an iterator over a map in Golang
		index := node.NewIndex(holders)

		// not strictly nec. to sort, but often a nice feature.  here we sort by key which is location
		index.Sort(func(a, b reflect.Value) bool {
			return b.Int() >= a.Int()
		})

		return &nodes.Basic{

			OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
				var found *cupHolder
				key := r.Key

				// If New or Delete is true and list defines a key, then key is
				// guaranteed to be set
				if r.New {
					location := key[0].String()
					found = &cupHolder{}
					holders[location] = found
				} else if r.Delete {

					// remove from list.  this happens AFTER cupHolderNode(found) was
					// called and that node got a change to handle deletion as well where
					// it might close resources or stop routines
					delete(holders, key[0].String())

				} else if key != nil {

					// lookup by key
					found = holders[key[0].String()]

				} else {

					// get next holder in next
					next := index.NextKey(r.Row)
					if next != node.NO_VALUE {
						location := next.String()
						found = holders[location]
						key = []val.Value{val.String(location)}
					}
				}
				if found != nil {
					return cupHolderNode(found), key, nil
				}

				// no item found or end of list
				return nil, nil, nil
			},
		}
	}
}

// TODO:
// slices
// find
// actions
// notifications
// json reading/writing
// restconf
// config v.s. operational
// constraints
// enumeration
// typedefs
// groups
// onDelete
// onBeginEdit
// onEndEdit
// schema browser
// triggers
// reusing internal and external node functions
// database
// peek
// diff
// tee
