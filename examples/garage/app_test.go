package garage

import (
	"testing"

	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/examples/car"
	"github.com/c2stack/c2g/meta"
)

var devices = make(map[string]device.Device)

func Test_App(t *testing.T) {

	ypath := meta.MultipleSources(
		&meta.FileStreamSource{Root: "../../yang"},
		&meta.FileStreamSource{Root: "../car"},
		&meta.FileStreamSource{Root: "."},
	)

	dm := device.NewMap()

	car0 := car.New()
	dev0 := device.New(ypath)
	chkErr(dev0.Add("car", car.Node(car0)))
	dm.Add("dev0", dev0)
	car0.Speed = 10
	car0.Start()

	g := NewGarage()
	dev1 := device.New(ypath)
	chkErr(dev1.Add("garage", Node(g)))
	dm.Add("dev1", dev1)
	o := g.Options()
	o.TireRotateMiles = 100
	o.PollTimeMs = 100
	g.ApplyOptions(o)

	ManageCars(g, dm)

	wait := make(chan struct{})
	g.OnUpdate(func(c Car, w workType) {
		t.Logf("car %s did %s", c.Id(), w)
		wait <- struct{}{}
	})
	t.Log("waiting for maintenance...")
	<-wait
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}
