package main

import (
	"flag"

	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/examples/car"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/restconf"
)

// Initialize and start our Car micro-service application with C2Stack for
// RESTful based management
//
// To run:
//    cd ./src/vendor/github.com/c2stack/c2g/examples/car/car-advanced
//    go run ./main.go
//
// Then open web browser to
//   http://localhost:8080/
//

var portParam = flag.String("port", "8090", "restconf port")
var deviceIdParam = flag.String("id", "car-advanced", "device id")

func main() {
	flag.Parse()

	// Your application
	app := car.New()

	// Where to looks for yang files, this tells library to use these
	// two relative paths
	localYpath := &meta.FileStreamSource{Root: ".."}
	ypath := meta.MultipleSources(
		localYpath,
		&meta.FileStreamSource{Root: "../../../yang"},
	)

	// Every management has a "device" container for.  The second
	// argument is path to UI files and that's only found in one of the paths
	device := conf.NewLocalDeviceWithUi(ypath, localYpath)

	// Browser is the management handle to your entire application
	device.Add("car", car.Node(app))

	// allows other end to discover local module information
	device.Add("ietf-yang-library", conf.LocalDeviceYangLibNode(device))

	// will handle registering to application management system
	callHome := conf.NewCallHome(ypath, restconf.NewInsecureDeviceByHostAndPort)
	device.Add("call-home", conf.CallHomeNode(callHome))

	// in our car app, we start off by running start.
	app.Start()

	// RESTful management.  Only required for RESTful based management.
	restconf.NewManagement(device, ":"+*portParam)

	// Normally loaded from local config file, but we'll configure it
	// manually here to a local application server we presume it running
	options := callHome.Options()
	options.DeviceId = *deviceIdParam
	options.LocalPort = *portParam
	options.RegistrationPort = "8080"
	options.RegistrationAddress = "127.0.0.1"
	if err := callHome.ApplyOptions(options); err != nil {
		panic(err)
	}

	select {}
}
