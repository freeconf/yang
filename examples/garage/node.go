package garage

import (
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/node"
)

func ManageCars(g *Garage, locator device.ServiceLocator) c2.Subscription {
	cars := make(map[string]CarHandle)
	return locator.OnModuleUpdate("car", func(d device.Device, id string, change device.Change) {
		switch change {
		case device.Added:
			b, err := d.Browser("car")
			if err != nil {
				panic(err)
			}
			car := &carDriver{
				id: id,
				b:  b,
			}
			cars[id] = g.AddCar(car)
		case device.Removed:
			if car, found := cars[id]; found {
				g.RemoveCar(car)
				delete(cars, id)
			}
		}
	})
}

func Node(g *Garage) node.Node {
	o := g.Options()
	return &node.Extend{
		Node: node.ReflectNode(&o),
		OnNotify: func(p node.Node, r node.NotifyRequest) (node.NotifyCloser, error) {
			switch r.Meta.GetIdent() {
			case "maintenance":
				sub := g.OnUpdate(func(c Car, work WorkType) {
					r.Send(carEventNode(c, work))
				})
				return sub.Close, nil
			}
			return nil, nil
		},
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "carCount":
				hnd.Val = &node.Value{Int: g.CarCount()}
			case "carsServiced":
				hnd.Val = &node.Value{Int: g.CarsServiced()}
			case "tireRotations":
				hnd.Val = &node.Value{Int: g.TireRotations}
			case "tireReplacements":
				hnd.Val = &node.Value{Int: g.TireReplacements}
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			if err := p.EndEdit(r); err != nil {
				return err
			}
			return g.ApplyOptions(o)
		},
	}
}

func carEventNode(c Car, work WorkType) node.Node {
	return &node.MyNode{
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "car":
				hnd.Val = &node.Value{Str: c.Id()}
			case "work":
				hnd.Val = &node.Value{Int: int(work)}
			}
			return nil
		},
	}
}

func carStateNode(state *CarState) node.Node {
	return node.ReflectNode(state)
}
