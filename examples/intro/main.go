package main

import (
	"os"
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
	d.ApplyStartupConfig(os.Stdin)

	// trick to sleep forever...
	select {}
}

// implement your mangement api
func manage(car *Car) node.Node {
	return &nodes.Extend{
		// use reflect when possible
		Base: nodes.ReflectChild(car),

		// handle action request
		OnAction: func(parent node.Node, req node.ActionRequest) (node.Node, error) {
			switch req.Meta.GetIdent() {
			case "start":
				go car.Start()
			}
			return nil, nil
		},
	}
}
