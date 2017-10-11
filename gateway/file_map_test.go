package gateway

import (
	"flag"
	"fmt"
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/nodes"
)

var update = flag.Bool("update", false, "update gold files")

func TestFileStore(t *testing.T) {
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
	c2.Gold(t, *update, []byte(actual), "m1.json")
}
