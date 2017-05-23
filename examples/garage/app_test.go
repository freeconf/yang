package garage

import (
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/examples/car"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

var devices = make(map[string]conf.Device)

func Test_App(t *testing.T) {

	ypath := meta.MultipleSources(
		&meta.FileStreamSource{Root: "../../yang"},
		&meta.FileStreamSource{Root: "../car"},
		&meta.FileStreamSource{Root: "."},
	)

	p := conf.NewProxy(ypath, localDevices, localServer)
	pNode := conf.ProxyNode(p)
	dev0 := conf.NewLocalDevice(ypath)
	chkErr(dev0.Add("proxy", pNode))
	devices["dev0"] = dev0
	chkErr(p.Mount("dev0", "local:", ""))

	car0 := car.New()
	dev1 := conf.NewLocalDevice(ypath)
	chkErr(dev1.Add("car", car.Node(car0)))
	car0.Start()
	devices["dev1"] = dev1
	chkErr(p.Mount("dev1", "local:", ""))
	car0.Speed = 10
	car0.Start()

	g := NewGarage()
	dev2 := conf.NewLocalDevice(ypath)
	chkErr(dev2.Add("garage", Node(g)))
	devices["dev2"] = dev2
	chkErr(p.Mount("dev2", "local:", ""))

	o := g.Options()
	o.TireRotateMiles = 100
	o.PollTimeMs = 100
	g.ApplyOptions(o)

	_, err := ManageCars(localServiceLocator{}, g)
	chkErr(err)

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

func localDevices(yangPath meta.StreamSource, address string, port string, deviceId string) (conf.Device, error) {
	return devices[deviceId], nil
}

func localServer(id string, d conf.Device) error {
	c2.Debug.Print("serving ", id)
	return nil
}

type localServiceLocator struct{}

func (localServiceLocator) FindDevice(deviceId string) conf.Device {
	return devices[deviceId]
}

func (localServiceLocator) FindBrowser(module string) *node.Browser {
	for _, d := range devices {
		hnds, _ := d.ModuleHandles()
		for name := range hnds {
			if name == module {
				b, _ := d.Browser(name)
				return b
			}
		}
	}
	return nil
}
