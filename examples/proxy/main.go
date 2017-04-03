package main

import (
	"flag"

	"log"

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
var portParam = flag.String("port", "8080", "restconf port")

func main() {
	flag.Parse()

	// where all yang files are stored
	yangPath := &meta.FileStreamSource{Root: "../../yang"}

	// where UI files are stored
	uiPath := &meta.FileStreamSource{Root: "."}

	// Even though this is a server component, we still organize things thru a device
	// because this proxy will appear like a "Device" to application management systems
	// "northbound"" representing all the devices that are "southbound".
	d := conf.NewLocalDeviceWithUi(yangPath, uiPath)

	// Add RESTCONF service
	mgmt := restconf.NewManagement(d, ":"+*portParam)
	if err := d.Add("restconf", restconf.Node(mgmt)); err != nil {
		log.Fatal(err)
	}

	// RESTCONF Proxy is not an official part of RFCs but there is
	// a draft for NETCONF protocol.
	//  https://tools.ietf.org/id/draft-wangzheng-netconf-proxy-00.txt
	p := conf.NewProxy(yangPath, restconf.NewInsecureClientByHostAndPort, mgmt.DeviceHandler.MultiDevice)
	proxyDriver := conf.ProxyNode(p)
	d.Add("proxy", proxyDriver)

	// Devices will be looking for this API on proxy.  Notice we give the same node
	// because call-home-register is a subset of the API for proxy.  This is a powerful
	// way to have the same code drive two similar APIs.
	d.Add("call-home-register", proxyDriver)

	// Wait for cntrl-c...
	select {}
}
