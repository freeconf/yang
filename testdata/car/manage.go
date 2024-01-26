package car

import (
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
)

// ///////////////////////
// C A R    M A N A G E M E N T
//
// Manage your car application using FreeCONF library according to the car.yang
// model file.
//
// Manage is root handler from car.yang. i.e. module car { ... }
func Manage(car *Car) node.Node {

	// We're letting reflect do a lot of the work when the yang file matches
	// the field names and methods in the objects.  But we extend reflection
	// to add as much custom behavior as we want
	return &nodeutil.Node{

		// Initial object. Note: as the tree is traversed, new Node instances
		// will have different values in their Object reference
		Object: car,

		Options: nodeutil.NodeOptions{
			ActionInputExploded:  true,
			ActionOutputExploded: true,
		},

		// implement RPCs
		OnAction: func(n *nodeutil.Node, r node.ActionRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "reset":
				// here we implement custom handler for action just as an example
				// If there was a Reset() method on car then this switch/case would
				// not be nec.
				car.Miles = 0
			default:
				// all the actions like start, stop and rotateTire are called
				// thru reflecton here because their method names align with
				// the YANG.
				return n.DoAction(r)
			}
			return nil, nil
		},

		// implement yang notifications (which are really just event streams)
		OnNotify: func(p *nodeutil.Node, r node.NotifyRequest) (node.NotifyCloser, error) {
			switch r.Meta.Ident() {
			case "update":
				// can use an adhoc struct send event
				sub := car.OnUpdate(func(e UpdateEvent) {
					msg := struct {
						Event int
					}{
						Event: int(e),
					}
					// events are nodes too
					r.Send(nodeutil.ReflectChild(&msg))
				})

				// we return a close **function**, we are not actually closing here
				return sub.Close, nil
			}

			// there is no default implementation at this time, all notification streams
			// require custom handlers.
			return p.Notify(r)
		},

		// See nodeutil.Node for list of all possible handlers, but just some of the more
		// useful ones include:
		//
		//  OnField - custom leaf handling
		//  OnChild - custom container and list handling
		//  OnOptions - changing options for select parts of the yang tree
		//  OnRead - custom reflect value handling for reads
		//  OnWrite - custom reflect value handling for writes
		//  OnNewChild - custom data structure instantiation
		//  OnEndEdit - custom validation for making changes to data structures
		//  OnChoose - custom handling for choice statements
		//  OnContext - among other things, passing data down the tree
		//  ...more
	}
}
