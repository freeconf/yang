package main

import (
	"flag"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/examples/db"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/restconf"
)

// Initialize and start our RESTCONF data store service.
//
// To run:
//    cd ./src/vendor/github.com/c2stack/c2g/examples/db/cmd
//    go run ./main.go
//
// Then open web browser to
//   http://localhost:8080/restconf/ui/index.html
//

var startup = flag.String("startup", "startup.json", "startup configuration file.")
var verbose = flag.Bool("verbose", false, "verbose")

func main() {
	flag.Parse()
	c2.DebugLog(*verbose)

	// where all yang files are stored
	yangPath := &meta.FileStreamSource{Root: "../../../yang"}

	// where UI files are stored
	uiPath := &meta.FileStreamSource{Root: ".."}

	// Even though this is a server component, we still organize things thru a device
	// because this proxy will appear like a "Device" to application management systems
	// "northbound"" representing all the devices that are "southbound".
	d := device.NewWithUi(yangPath, uiPath)

	// Add RESTCONF service
	mgmt := restconf.NewServer(d)

	// We "wrap" each device with a device that splits CRUD operations
	// to local store AND the original device.  This gives us transparent
	// persistance of device data w/o altering the device API.
	store := &device.Db{
		Delegate: restconf.ProtocolHandler(yangPath),

		// Supplying your own code to read/write configuration is surprisingly
		// easy.  You might store config in mongo, etcd, redis, git or any
		// other hierarchical data store.
		IO: db.FileStore{VarDir: "./var"},
	}

	// Devices will be looking for this API on proxy.  Notice we give the same node
	// because call-home-register is a subset of the API for proxy.  This is a powerful
	// way to have the same code drive two similar APIs.
	dm := device.NewMap()
	chkErr(d.Add("map", device.MapNode(dm, mgmt.DeviceAddress, store.NewDevice)))

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
