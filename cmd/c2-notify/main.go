package main

import (
	"log"
	"os"

	"context"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/restconf"
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
		if err = payload.InsertInto(node.NewJsonWriter(os.Stdout).Node()).LastErr; err != nil {
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
