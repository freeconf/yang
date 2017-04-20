package main

import (
	"flag"

	"strings"

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
//   http://localhost:8080
//
var defaultConfig = `
{
	"restconf" : {
		"web" : {
			"port" : ":8090"
		}
	},
	"call-home" : {
		"deviceId" : "car-advanced",
		"localPort" : "8090",
		"registrationPort" : "8080",
		"registrationAddress" : "127.0.0.1"
	}
}
`

var configFile = flag.String("config", "", "alternate configuration file.  Default config:"+defaultConfig)

func main() {
	flag.Parse()

	// Any existing application
	app := car.New()

	// Where to looks for yang files, this tells library to use these
	// two relative paths.  StreamSource is an abstraction to data sources
	// that might be local or remote or combinations of all the above.
	carPath := &meta.FileStreamSource{Root: "."}
	ypath := meta.MultipleSources(
		carPath,
		&meta.FileStreamSource{Root: "../../yang"},
	)

	// Every management has a "device" container. A device can have many "modules"
	// installed which are like mini-app.
	//   carPath - where UI files are located
	//   ypath - where *.yang files are located
	device := conf.NewLocalDeviceWithUi(ypath, carPath)

	// Here we are installing the "car" module which is our main application.
	//   "car" - the name of the module.  Will automatically try to find car.yang
	//           in yang path
	//   car.Node(app)  - we are linking our car application with driver to handle
	//           car management requests.
	chkErr(device.Add("car", car.Node(app)))

	// allows other end to discover local module information.  Required by all
	// devices according to RFC
	chkErr(device.Add("ietf-yang-library", conf.LocalDeviceYangLibNode(device)))

	// This optional module supports the Call-Home RFC draft.  It will register
	// this service with application management system.
	callHome := conf.NewCallHome(ypath, restconf.NewInsecureClientByHostAndPort)
	chkErr(device.Add("call-home", conf.CallHomeNode(callHome)))

	// Adding RESTCONF protocol support.  Should you want an alternate protocol,
	// you could
	mgmt := restconf.NewManagement(device)
	chkErr(device.Add("restconf", restconf.Node(mgmt)))

	// Even though the main configuration comes from the application management
	// system after call-home has registered this system it's often neccessary
	// to bootstrap config for some of the local modules
	if *configFile == "" {
		chkErr(device.ApplyStartupConfig(strings.NewReader(defaultConfig)))
	} else {
		chkErr(device.ApplyStartupConfigFile(*configFile))
	}

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
