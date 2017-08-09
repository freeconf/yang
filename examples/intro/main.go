package main

import (
	"strings"
	"time"

	"github.com/c2stack/c2g/node"

	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/nodes"
	"github.com/c2stack/c2g/restconf"
)

type Car struct {
	Speed int
	Miles int64
}

func (c *Car) Start() {
	for {
		<-time.After(time.Duration(c.Speed) * time.Millisecond)
		c.Miles += 1
	}
}

func main() {

	// Your app
	car := &Car{}

	// Add management
	d := device.New(yang.YangPath())
	d.Add("car", manage(car))
	restconf.NewServer(d)
	d.ApplyStartupConfig(strings.NewReader(`{"restconf":{"web":{"port":":8080"}},"car":{}}`))

	// trick to sleep forever...
	select {}
}

// implement your mangement api
func manage(c *Car) node.Node {
	return &nodes.Extend{
		Node: nodes.Reflect(c),
		OnAction: func(p node.Node, r node.ActionRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "start":
				go c.Start()
			}
			return nil, nil
		},
	}
}
