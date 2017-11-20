package gateway

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/freeconf/c2g/c2"
	"github.com/freeconf/c2g/device"
	"github.com/freeconf/c2g/nodes"
	"github.com/freeconf/c2g/testdata"
)

var update = flag.Bool("update", false, "update gold files")

func TestFileStoreOffline(t *testing.T) {
	reg := NewLocalRegistrar()
	fs := NewFileStore(reg, "./testdata/var")
	c2.AssertEqual(t, "[d1 d2]", fmt.Sprintf("%v", fs.deviceIds()))
	d1, err := fs.Device("d1")
	if err != nil {
		t.Fatal(err)
	}
	b1, err := d1.Browser("m1")
	if err != nil {
		t.Fatal(err)
	}
	actual, err := nodes.WritePrettyJSON(b1.Root())
	if err != nil {
		t.Fatal(err)
	}
	c2.Gold(t, *update, []byte(actual), "gold/m1.json")
}

func TestFileStoreOnline(t *testing.T) {
	reg := NewLocalRegistrar()
	dir := "./.var/file_map_test-tmp"
	os.RemoveAll(dir)
	reg.RegisterDevice("x", "foo")
	fs := NewFileStore(reg, dir)
	c2.AssertEqual(t, 0, len(fs.deviceIds()))
	birdDevice, birds := testdata.BirdDevice(`{
	}
	`)
	fs.AddProtocolHandler(func(string) (device.Device, error) {
		return birdDevice, nil
	})
	gwDevice, err := fs.Device("x")
	if err != nil {
		t.Fatal(err)
	}
	if gwDevice == nil {
		t.Fatal("no device returned")
	}
	b, err := gwDevice.Browser("bird")
	if err != nil {
		t.Fatal(err)
	}
	if b == nil {
		t.Fatal("no browser")
	}
	err = b.Root().InsertFrom(nodes.ReadJSON(`{
		"bird" : [{
			"name" : "bard owl"
		}]
	}
	`)).LastErr
	if err != nil {
		t.Fatal(err)
	}
	c2.AssertEqual(t, 1, len(birds))
	actual, err := ioutil.ReadFile(dir + "/config/x/bird.json")
	if err != nil {
		t.Fatal(err)
	}
	c2.Gold(t, *update, actual, "gold/online.json")
}
