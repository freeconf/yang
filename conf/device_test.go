package conf

import "testing"
import "github.com/c2stack/c2g/c2"

func Test_SplitAddress(t *testing.T) {
	tests := struct {
		url     string
		address string
		module  string
		path    string
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
	}
	for _, test := range tests {
		address, module, path := SplitAddress(test.url)
		if err := c2.CheckEqual(test.address, address); err != nil {
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
