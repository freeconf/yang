package main

import (
	"flag"
	"fmt"
	"os"

	"log"

	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
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
	s := &conf.Store{
		Delegate: restconf.NewInsecureClientByHostAndPort,
		Support:  fileStore{},
	}
	p := conf.NewProxy(yangPath, s.StoreDevice, mgmt.DeviceHandler.MultiDevice)
	proxyDriver := conf.ProxyNode(p)
	d.Add("proxy", proxyDriver)

	// Devices will be looking for this API on proxy.  Notice we give the same node
	// because call-home-register is a subset of the API for proxy.  This is a powerful
	// way to have the same code drive two similar APIs.
	d.Add("call-home-register", proxyDriver)

	// Wait for cntrl-c...
	select {}
}

// Store all data in simple files.  Normally you would save this to a highly
// available service like a database.
type fileStore struct{}

func (fileStore) fname(deviceId string, module string) string {
	return fmt.Sprintf("%s:%s.json", deviceId, module)
}

// LoadStore implements conf.StoreSupport interface to load data
func (self fileStore) LoadStore(deviceId string, module string, b *node.Browser) error {
	fname := self.fname(deviceId, module)
	_, err := os.Stat(fname)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	rdr, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer rdr.Close()
	// this walks data for device's data for this module (a device might have multiple
	// modules) and sends it to json
	if err := b.Root().InsertFrom(node.NewJsonReader(rdr).Node()).LastErr; err != nil {
		return err
	}
	return nil
}

// SaveStore implements conf.StoreSupport interface to save data
func (self fileStore) SaveStore(deviceId string, module string, b *node.Browser) error {
	fname := self.fname(deviceId, module)
	wtr, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer wtr.Close()

	// We only want to look at config and only config that isn't set to the default values
	params := "content=config&with-defaults=trim"

	// this walks data for device's data for this module (a device might have multiple
	// modules) and sends it to json
	if err := b.Root().Constrain(params).InsertInto(node.NewJsonPretty(wtr).Node()).LastErr; err != nil {
		return err
	}
	return nil
}
