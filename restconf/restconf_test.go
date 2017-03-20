package restconf

import "testing"
import "github.com/c2stack/c2g/c2"

func Test_SplitAddress(t *testing.T) {
	tests := []struct {
		url     string
		address string
		port    string
		module  string
		path    string
		hasErr  bool
	}{
		{
			url:     "http://server:port/restconf/module:path/some=x/where",
			address: "http://server:port/restconf/",
			module:  "module",
			path:    "path/some=x/where",
		},
		{
			url:     "http://server/device=100/some-mod:path=z?p=1&z=x",
			address: "http://server/device=100/",
			module:  "some-mod",
			path:    "path=z?p=1&z=x",
		},
		{
			url:    "no-protocol",
			hasErr: true,
		},
		{
			url:    "foo://no-module-or-path",
			hasErr: true,
		},
		{
			url:    "foo://server/no-mount",
			hasErr: true,
		},
		{
			url:    "foo://server/mount/no-module",
			hasErr: true,
		},
		{
			url:     "foo://server/mount/module:",
			address: "foo://server/mount/",
			module:  "module",
			path:    "",
		},
	}
	for _, test := range tests {
		address, module, path, err := SplitAddress(test.url)
		if test.hasErr && err == nil {
			t.Error("Expected parse error ", test.url)
			continue
		}
		if !test.hasErr && err != nil {
			t.Error(err)
			continue
		}
		if err := c2.CheckEqual(test.address, address); err != nil {
			t.Error(err)
		}
		if err := c2.CheckEqual(test.port, port); err != nil {
			t.Error(err)
		}
		if err := c2.CheckEqual(test.module, module); err != nil {
			t.Error(err)
		}
		if err := c2.CheckEqual(test.path, path); err != nil {
			t.Error(err)
		}
	}
}
