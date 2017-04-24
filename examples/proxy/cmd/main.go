package main

import (
	"flag"
	"strings"

	"github.com/c2stack/c2g/conf"
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
var defaultConfig = `
{
	"restconf" : {
		"web" : {
			"port" : ":8080"
		}
	}
}
`
var configFile = flag.String("config", "", "alternate configuration file.  Default config:"+defaultConfig)

func main() {
	flag.Parse()

	// where all yang files are stored
	yangPath := meta.MultipleSources(
		&meta.FileStreamSource{Root: "../../../yang"},
		&meta.FileStreamSource{Root: "../../yang"},
	)

	// where UI files are stored
	uiPath := &meta.FileStreamSource{Root: ".."}

	// Even though this is a server component, we still organize things thru a device
	// because this proxy will appear like a "Device" to application management systems
	// "northbound"" representing all the devices that are "southbound".
	d := conf.NewLocalDeviceWithUi(yangPath, uiPath)

	// Add RESTCONF service
	mgmt := restconf.NewManagement(d)
	chkErr(d.Add("restconf", restconf.Node(mgmt)))

	// RESTCONF Proxy is not an official part of RFCs but there is
	// a draft for NETCONF protocol.
	//  https://tools.ietf.org/id/draft-wangzheng-netconf-proxy-00.txt
	p := conf.NewProxy(yangPath, restconf.NewInsecureClientByHostAndPort, mgmt.DeviceHandler.MultiDevice)
	proxyDriver := conf.ProxyNode(p)
	chkErr(d.Add("proxy", proxyDriver))

	// Devices will be looking for this API on proxy.  Notice we give the same node
	// because call-home-register is a subset of the API for proxy.  This is a powerful
	// way to have the same code drive two similar APIs.
	chkErr(d.Add("call-home-register", proxyDriver))

	// bootstrap config for all local modules
	if *configFile == "" {
		chkErr(d.ApplyStartupConfig(strings.NewReader(defaultConfig)))
	} else {
		chkErr(d.ApplyStartupConfigFile(*configFile))
	}

	// Wait for cntrl-c...
	select {}
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}
