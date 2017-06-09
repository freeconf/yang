package main

import (
	"flag"

	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/examples/car"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/restconf"
)

// Initialize and start our Car micro-service application with C2Stack for
// RESTful based management
//
// To run:
//    export GOPATH=`pwd`
//    cd ./src/vendor/github.com/c2stack/c2g/examples/car/cmd
//    go run ./main.go -startup startup.json
//
// Then open web browser to
//   http://localhost:8080
//
var startup = flag.String("startup", "startup.json", "start-up configuration file.")

func main() {
	flag.Parse()

	// Any existing application
	app := car.New()

	// Where to looks for yang files, this tells library to use these
	// two relative paths.  StreamSource is an abstraction to data sources
	// that might be local or remote or combinations of all the above.
	uiPath := &meta.FileStreamSource{Root: "../web"}
	yangPath := meta.MultipleSources(
		&meta.FileStreamSource{Root: ".."},
		&meta.FileStreamSource{Root: "../../../yang"},
	)

	// Every management has a "device" container. A device can have many "modules"
	// installed which are really microservices.
	//   carPath - where UI files are located
	//   ypath - where *.yang files are located
	d := device.NewWithUi(yangPath, uiPath)

	// Here we are installing the "car" module which is our main application.
	//   "car" - the name of the module causing car.yang to load from yang path
	//   car.Node(app) - we are linking our car application with driver to handle
	//           car management requests.
	chkErr(d.Add("car", car.Node(app)))

	// Adding RESTCONF protocol support.  Should you want an alternate protocol,
	// you could
	mgmt := restconf.NewServer(d)
	chkErr(d.Add("restconf", restconf.Node(mgmt, yangPath)))

	// Even though the main configuration comes from the application management
	// system after call-home has registered this system it's often neccessary
	// to bootstrap config for some of the local modules
	chkErr(d.ApplyStartupConfigFile(*startup))

	// in our car app, we start off by running start.
	app.Start()

	// wait for cntrl-c...
	select {}
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}
