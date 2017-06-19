package orchestrator

import (
	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/examples/car"
	"github.com/c2stack/c2g/examples/garage"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

type LocalFactory struct {
	Ypath meta.StreamSource
	Map   *device.Map
}

func (self *LocalFactory) NewApp(app *App) error {
	var n node.Node
	switch app.Type {
	case "car":
		c := car.New()
		n = car.Node(c)
		defer func() {
			c.Start()
		}()
	case "garage":
		g := garage.NewGarage()
		n = garage.Node(g)
		defer func() {
			garage.ManageCars(g, self.Map)
		}()
	default:
		panic("unknown type : " + app.Type)
	}
	d := device.New(self.Ypath)
	if err := d.Add(app.Type, n); err != nil {
		return err
	}
	if err := d.ApplyStartupConfigData(app.Startup); err != nil {
		return err
	}
	self.Map.Add(app.Id, d)
	return nil
}
