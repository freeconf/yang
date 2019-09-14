package restconf

import (
	"testing"

	"github.com/freeconf/yang/c2"
)

func Test_findDeviceIdInUrl(t *testing.T) {
	dev := findDeviceIdInUrl("http://server:port/restconf=abc/")
	c2.AssertEqual(t, "abc", dev)
	dev = findDeviceIdInUrl("http://server:port/restconf/")
	c2.AssertEqual(t, "", dev)
}
