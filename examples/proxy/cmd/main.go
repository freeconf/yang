package main

import (
	"flag"

	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/restconf"
)

// Initialize and start our RESTCONF proxy service.
//
// To run:
//    cd ./src/vendor/github.com/c2stack/c2g/examples/proxy
//    go run ./main.go
//
// Then open web browser to
//   http://localhost:8080/restconf/ui/index.html
//

var startup = flag.String("startup", "startup.json", "startup configuration file.")

func main() {
	flag.Parse()

	// where UI files are stored
	uiPath := &meta.FileStreamSource{Root: "../web"}

	// where all yang files are stored just for the proxy
	// models for things that register are pulled automatically
	yangPath := &meta.FileStreamSource{Root: "../../../yang"}

	// Even though this is a server component, we still organize things thru a device
	// because this proxy will appear like a "Device" to application management systems
	// "northbound"" representing all the devices that are "southbound".
	d := device.NewWithUi(yangPath, uiPath)

	// Add RESTCONF service
	mgmt := restconf.NewServer(d)
	chkErr(d.Add("restconf", restconf.Node(mgmt, yangPath)))

	// Exposing your device manager means you can represent other devices
	dm := device.NewMap()
	client := restconf.NewClient(yangPath)
	chkErr(d.Add("device-manager", device.MapNode(dm, mgmt, client)))

	// bootstrap config for all local modules
	chkErr(d.ApplyStartupConfigFile(*startup))

	// Wait for cntrl-c...
	select {}
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}
