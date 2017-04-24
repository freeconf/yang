package main

import (
	"flag"
	"strings"

	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/examples/garage"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/restconf"
)

var defaultConfig = `
{
	"restconf" : {
		"web" : {
			"port" : ":8091"
		}
	},
	"call-home" : {
		"deviceId" : "g1",
		"localPort" : "8091",
		"registrationPort" : "8080",
		"registrationAddress" : "127.0.0.1"
	}
}
`
var configFile = flag.String("config", "", "alternate configuration file.  Default config:"+defaultConfig)

func main() {
	flag.Parse()

	app := garage.NewGarage()

	// Where to looks for yang files, this tells library to use these
	// two relative paths.  StreamSource is an abstraction to data sources
	// that might be local or remote or combinations of all the above.
	garagePath := &meta.FileStreamSource{Root: ".."}
	ypath := meta.MultipleSources(
		garagePath,
		&meta.FileStreamSource{Root: "../../yang"},
	)

	device := conf.NewLocalDeviceWithUi(ypath, garagePath)
	chkErr(device.Add("garage", garage.Node(app)))

	// Standard management modules
	chkErr(device.Add("ietf-yang-library", conf.LocalDeviceYangLibNode(device)))
	callHome := conf.NewCallHome(ypath, restconf.NewInsecureClientByHostAndPort)
	chkErr(device.Add("call-home", conf.CallHomeNode(callHome)))
	mgmt := restconf.NewManagement(device)
	chkErr(device.Add("restconf", restconf.Node(mgmt)))
	if *configFile == "" {
		chkErr(device.ApplyStartupConfig(strings.NewReader(defaultConfig)))
	} else {
		chkErr(device.ApplyStartupConfigFile(*configFile))
	}

	select {}
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}
