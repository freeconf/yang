package restconf

import (
	"testing"

	"net/url"

	"github.com/freeconf/c2g/c2"
)

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
			url:     "http://server:port/restconf/data/module:path/some=x/where",
			address: "http://server:port/restconf/data/",
			module:  "module",
			path:    "path/some=x/where",
		},
		{
			url:     "http://server/restconf=100/streams/module:path=z?p=1&z=x",
			address: "http://server/restconf=100/streams/",
			module:  "module",
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
		c2.AssertEqual(t, test.address, address)
		c2.AssertEqual(t, test.module, module)
		c2.AssertEqual(t, test.path, path)
	}
}

func Test_AppendUrlSegment(t *testing.T) {
	tests := [][]string{
		{
			"a", "b", "a/b",
		},
		{
			"a/", "b", "a/b",
		},
		{
			"a/", "/b", "a/b",
		},
		{
			"a", "/b", "a/b",
		},
		{
			"", "", "",
		},
		{
			"a/", "", "a/",
		},
		{
			"", "/b", "/b",
		},
	}
	for _, test := range tests {
		actual := appendUrlSegment(test[0], test[1])
		c2.AssertEqual(t, test[2], actual)
	}
}

func Test_ipAddrSplitHostPort(t *testing.T) {
	tests := [][]string{
		{"127.0.0.1:1000", "127.0.0.1", "1000"},
		{"127.0.0.1", "127.0.0.1", ""},
		{"[::1]:1000", "[::1]", "1000"},
		{"::1", "::1", ""},
		{"[0:0:0:0:0:0:0]:1000", "[0:0:0:0:0:0:0]", "1000"},
	}
	for _, test := range tests {
		host, port := ipAddrSplitHostPort(test[0])
		c2.AssertEqual(t, test[1], host)
		c2.AssertEqual(t, test[2], port)
	}
}

func Test_shift(t *testing.T) {
	tests := []struct {
		in              string
		expectedSegment string
		expectedPath    string
	}{
		{
			in:              "http://server:port/some/path/here",
			expectedSegment: "some",
			expectedPath:    "path/here",
		},
		{
			in:              "http://server:port/some/path/here?p=1&z=x",
			expectedSegment: "some",
			expectedPath:    "path/here",
		},
		{
			in:              "some/path/here",
			expectedSegment: "some",
			expectedPath:    "path/here",
		},
		{
			in:              "some",
			expectedSegment: "some",
			expectedPath:    "",
		},
		{
			in:              "some/",
			expectedSegment: "some",
			expectedPath:    "",
		},
	}
	for _, test := range tests {
		orig, err := url.Parse(test.in)
		if err != nil {
			panic(err)
		}
		actualSeg, actualPath := shift(orig, '/')
		c2.AssertEqual(t, test.expectedSegment, actualSeg)
		c2.AssertEqual(t, test.expectedPath, actualPath.Path)
	}
}

func Test_shiftOptionalParamWithinSegment(t *testing.T) {
	tests := []struct {
		in    string
		seg   string
		param string
		path  string
	}{
		{
			in:   "http://server:port/some/path/here",
			seg:  "some",
			path: "path/here",
		},
		{
			in:  "some/",
			seg: "some",
		},
		{
			in:  "some=/",
			seg: "some",
		},
		{
			in:    "some=x/",
			param: "x",
			seg:   "some",
		},
		{
			in:   "data/call-home-register:",
			seg:  "data",
			path: "call-home-register:",
		},
		{
			in:   "/some",
			seg:  "some",
			path: "",
		},
	}
	for _, test := range tests {
		orig, err := url.Parse(test.in)
		if err != nil {
			panic(err)
		}
		seg, param, path := shiftOptionalParamWithinSegment(orig, '=', '/')
		c2.AssertEqual(t, test.seg, seg)
		c2.AssertEqual(t, test.param, param)
		c2.AssertEqual(t, test.path, path.Path)
	}
}
