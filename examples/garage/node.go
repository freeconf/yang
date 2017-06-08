package garage

import (
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/node"
)

func ManageCars(g *Garage, locator conf.ServiceLocator) c2.Subscription {
	cars := make(map[string]CarHandle)
	return locator.OnModuleUpdate("car", func(device conf.Device, id string, change conf.Change) {
		switch change {
		case conf.Added:
			b, err := device.Browser("car")
			if err != nil {
				panic(err)
			}
			car := &carDriver{
				id: id,
				b:  b,
			}
			cars[id] = g.AddCar(car)
		case conf.Removed:
			if car, found := cars[id]; found {
				g.RemoveCar(car)
				delete(cars, id)
			}
		}
	})
}

func Node(g *Garage) node.Node {
	return &node.Extend{
		Node: node.ReflectNode(g),
		OnNotify: func(p node.Node, r node.NotifyRequest) (node.NotifyCloser, error) {
			switch r.Meta.GetIdent() {
			case "update":
				sub := g.OnUpdate(func(c Car, work workType) {
					r.Send(carEventNode(c, work))
				})
				return sub.Close, nil
			}
			return nil, nil
		},
	}
}

func carEventNode(c Car, work workType) node.Node {
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
