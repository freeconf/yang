package main

import (
	"flag"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/examples/garage"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/restconf"
)

var startup = flag.String("startup", "startup.json", "startup configuration file.")

func main() {
	flag.Parse()

	app := garage.NewGarage()

	// Where to looks for yang files, this tells library to use these
	// two relative paths.  StreamSource is an abstraction to data sources
	// that might be local or remote or combinations of all the above.
	uiPath := &meta.FileStreamSource{Root: "../web"}

	// notice the garage doesn't need yang for car.  it will get
	// that from proxy, that will in turn get it from car node, having
	// said that, if it does find yang locally, it will use it
	yangPath := meta.PathStreamSource("..:../../../yang")

	d := device.NewWithUi(yangPath, uiPath)

	mgmt := restconf.NewServer(d)

	chkErr(d.Add("garage", garage.Node(app)))

	// apply start-up config, just enough to initialize connection to
	// services that will finishing configuration
	chkErr(d.ApplyStartupConfigFile(*startup))

	var sub c2.Subscription
	mgmt.CallHome.OnRegister(func(d device.Device, u device.RegisterUpdate) {
		if sub != nil {
			sub.Close()
		}
		if u == device.Register {
			baseAddress := mgmt.CallHome.Options().Address
			dm := device.NewMapClient(d, baseAddress, restconf.ProtocolHandler(yangPath))
			sub = garage.ManageCars(app, dm)
		}
	})

	// wait for cntrl-c...
	select {}
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}
