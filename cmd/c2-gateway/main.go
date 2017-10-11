package main

import (
	"flag"

	"github.com/c2stack/c2g/gateway"
	"github.com/c2stack/c2g/restconf"

	"github.com/c2stack/c2g/meta"

	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/meta/yang"

	"github.com/c2stack/c2g/c2"
)

var startup = flag.String("startup", "startup.json", "startup configuration file.")
var verbose = flag.Bool("verbose", false, "verbose")
var web = flag.String("web", "", "web directory")
var varDir = flag.String("var", "var", "directory to store files")

func main() {
	flag.Parse()
	c2.DebugLog(*verbose)

	ypath := yang.YangPath()
	var d *device.Local
	if *web == "" {
		d = device.New(ypath)
	} else {
		d = device.NewWithUi(ypath, &meta.FileStreamSource{Root: *web})
	}

	reg := gateway.NewLocalRegistrar()
	m := gateway.NewFileStore(reg, "var")
	gateway.NewService(d, m, reg)

	mgmt := restconf.NewServer(d)

	mgmt.ServeDevices(m)

	chkErr(d.ApplyStartupConfigFile(*startup))

	// Wait for cntrl-c...
	select {}
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}
