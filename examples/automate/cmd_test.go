package automate

import (
	"bytes"
	"testing"
	"time"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
)

func Test_CmdStartup(t *testing.T) {
	sys := &CmdSystem{
		VarDir:       "./var",
		nextDeviceId: 9,
		nextPort:     8,
		proxyPort:    7,
	}
	e := sys.nextEntry("x")
	var buff bytes.Buffer
	sys.Startup(e, &buff)
	expected := `{
	"restconf" : {
		"web" : {
			"port" : ":8"
		},
		"callHome" : {
			"deviceId" : "x9",
			"localAddress" : "http://{REQUEST_ADDRESS}:8/restconf",
			"address" : "http://127.0.0.1:7",
			"endpoint" : "/restconf"
		}
	}	
}`
	c2.AssertEqual(t, buff.String(), expected)
}

// Fails on travis
//   dial tcp 127.0.0.1:9000: getsockopt: connection refused
func disabled_Test_CmdRun(t *testing.T) {
	c2.DebugLog(true)
	sys := &CmdSystem{
		VarDir:      "./var",
		ExamplesDir: "..",
		nextPort:    8000,
		proxyPort:   9000,
		YangPath:    &meta.FileStreamSource{Root: "../../yang"},
	}
	p, err := sys.New("proxy")
	if err != nil {
		t.Error(err)
	} else {
		defer p.Close()
	}
	c, err := sys.New("car")
	if err != nil {
		t.Error(err)
	} else {
		defer c.Close()
	}
	<-time.After(3 * time.Second)
}
