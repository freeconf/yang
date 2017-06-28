package main

import (
	"flag"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/examples/launch"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/restconf"
)

// Initialize and start our RESTCONF proxy service.
//
// To run:
//    cd ./src/vendor/github.com/c2stack/c2g/examples/proxy
//    go run ./main.go
//
// Then open web browser to
//   http://localhost:8080/restconf/ui/index.html
//

var startup = flag.String("startup", "startup.json", "startup configuration file.")

func main() {
	c2.DebugLog(true)
	flag.Parse()

	yangPath := meta.PathStreamSource("..:../../car:../../garage:../../../yang")

	d := device.New(yangPath)
	p := launch.New()
	chkErr(d.Add("launch", launch.Node(p, yangPath)))

	// Add RESTCONF service
	restconf.NewServer(d)

	// bootstrap config for all local modules
	chkErr(d.ApplyStartupConfigFile(*startup))

	// Wait for cntrl-c...
	select {}
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}
