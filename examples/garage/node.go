package garage

import (
	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/node"
)

func carStateNode(state *CarState) node.Node {
	return node.ReflectNode(state)
}

func ManageCars(sl conf.InterfaceLocator, g *Garage) (node.NotifyCloser, error) {
	drivers := make(map[string]CarHandle)
	onNew := func(id string, device conf.Device, module string) {
		b, err := device.Browser("car")
		if err != nil {
			panic(err)
		}
		driver := &carDriver{
			id: id,
			b:  b,
		}
		drivers[id] = g.AddCar(driver)
	}
	onRemove := func(id string, device conf.Device, module string) {
		if hnd, found := drivers[id]; found {
			g.RemoveCar(hnd)
			delete(drivers, id)
		}
	}
	return conf.FindBrowsers(sl, "car", onNew, onRemove)
}

func Node(g *Garage) node.Node {
	return &node.Extend{
		Node: node.ReflectNode(g),
		OnNotify: func(p node.Node, r node.NotifyRequest) (node.NotifyCloser, error) {
			switch r.Meta.GetIdent() {
			case "update":
				sub := g.OnUpdate(func(c Car, work workType) {
					r.Send(r.Context, carEventNode(c, work))
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
