package main

import (
	"log"
	"os"

	"context"

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
//  http://server:port/restconf/data/module:path
//  http://server:port/device=100/data/module:path
//
func main() {
	if len(os.Args) != 2 {
		usage()
	}
	address, module, path, err := restconf.SplitAddress(os.Args[1])
	if err != nil {
		panic(err)
	}
	d, err := restconf.NewDevice(yang.YangPath(), address)
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
	unsubscribe, err := b.Root().Find(path).NotificationsCntx(ctx, func(ctx2 context.Context, payload node.Selection) {
		if err = payload.InsertIntoCntx(ctx2, node.NewJsonWriter(os.Stdout).Node()).LastErr; err != nil {
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
