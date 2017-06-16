package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/restconf"
)

// Initialize and start our RESTCONF data store service.
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
	chkErr(d.Add("restconf", restconf.Node(mgmt, yangPath)))

	// We "wrap" each device with a device that splits CRUD operations
	// to local store AND the original device.  This gives of transparent
	// persistance of device data w/o altering the device API.
	store := &device.Store{
		YangPath: yangPath,
		Delegate: restconf.ProtocolHandler(yangPath),

		// Supplying your own code to read/write configuration is surprisingly
		// easy.  You might store config in mongo, etcd, redis, git or any
		// other hierarchical data store.
		Support: fileStore{},
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

// Store all data in simple files.  Normally you would save this to a highly
// available service like a database.
type fileStore struct{}

func (fileStore) fname(deviceId string, module string) string {
	return fmt.Sprintf("%s:%s.json", deviceId, module)
}

// LoadStore implements device.StoreSupport interface to load data
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

// SaveStore implements device.StoreSupport interface to save data
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

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}
