package restconf

import (
	"testing"

	"github.com/c2stack/c2g/c2"
)

func Test_findDeviceIdInUrl(t *testing.T) {
	dev := findDeviceIdInUrl("http://server:port/restconf=abc/")
	c2.AssertEqual(t, "abc", dev)
	dev = findDeviceIdInUrl("http://server:port/restconf/")
	c2.AssertEqual(t, "", dev)
}
