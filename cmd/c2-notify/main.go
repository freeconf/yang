package main

import (
	"log"
	"os"

	"context"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/meta/yang"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/nodes"
	"github.com/freeconf/gconf/restconf"
)

// Subscribes to a notification and exits on first message
//
// this can be expanded to repeat indefinitely as an option
// or supply an alternate value for 'origin' should the default
// not be valid for some reason
//
//  http://server:port/restconf/streams/module:path?c2-device=car-advanced
//  http://server:port/restconf=device/streams/module:path
//
func main() {
	c2.DebugLog(true)
	if len(os.Args) != 2 {
		usage()
	}
	address, module, path, err := restconf.SplitAddress(os.Args[1])
	c2.Info.Printf("%s %s %s %v", address, module, path, err)
	if err != nil {
		panic(err)
	}
	d, err := restconf.ProtocolHandler(yang.YangPath())(address)
	if err != nil {
		panic(err)
	}
	defer d.Close()
	b, err := d.Browser(module)
	if err != nil {
		panic(err)
	}
	wait := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	unsubscribe, err := b.RootWithContext(ctx).Find(path).Notifications(func(payload node.Selection) {
		wtr := &nodes.JSONWtr{Out: os.Stdout}
		if err = payload.InsertInto(wtr.Node()).LastErr; err != nil {
			log.Fatal(err)
		}
		wait <- true
	})
	defer unsubscribe()
	if err != nil {
		log.Fatal(err)
	}
	<-wait
}

func usage() {
	log.Fatalf(`usage : %s http://server:port/restconf/module:path/some=x/where`, os.Args[0])
}
