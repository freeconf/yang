package restconf

import (
	"testing"

	"github.com/c2stack/c2g/c2"
)

func Test_findDeviceIdInUrl(t *testing.T) {
	dev := findDeviceIdInUrl("http://server:port/restconf=abc/")
	if err := c2.CheckEqual("abc", dev); err != nil {
		t.Error(err)
	}
	dev = findDeviceIdInUrl("http://server:port/restconf/")
	if err := c2.CheckEqual("", dev); err != nil {
		t.Error(err)
	}
}
