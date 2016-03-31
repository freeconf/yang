package main

import (
	"github.com/c2g/node"
	"flag"
	"fmt"
	"os"
	"github.com/c2g/restconf"
	"github.com/c2g/meta/yang"
	"github.com/c2g/meta"
)

var configFileName = flag.String("config", "", "Configuration file")

func main() {
	flag.Parse()

	// TODO: Change this to a file-persistent store.
	if len(*configFileName) == 0 {
		fmt.Fprint(os.Stderr, "Required 'config' parameter missing\n")
		os.Exit(-1)
	}

	configFile, err := os.Open(*configFileName)
	if err != nil {
		panic(err)
	}
	config := node.NewJsonReader(configFile).Node()
	app := &app{}
	c := node.NewContext()
	if err != nil {
		panic(err)
	}
	if err = c.Selector(app.Select()).InsertFrom(config).LastErr; err != nil {
		panic(err)
	}
	// sleep "forever"
	app.Management.Listen()
}

type app struct {
	Management *restconf.Service
	c1 map[string]interface{}
	m *meta.Module
}

func (self *app) Select() *node.Selection {
	if self.m == nil {
		var err error
		if self.m, err = yang.LoadModule(yang.YangPath(), "c2-compliance.yang"); err != nil {
			panic(err)
		}
	}
	return node.Select(self.m, self.Node())
}

func (self *app) Node() node.Node {
	n := &node.MyNode{}
	n.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "management":
			if r.New {
				self.Management = restconf.NewService(self)
			}
			if self.Management != nil {
				return self.Management.Manage(), nil
			}
		case "c1":
			if r.New {
				self.c1 = make(map[string]interface{})
			}
			if self.c1 != nil {
				return node.MapNode(self.c1), nil
			}
		}
		return nil, nil
	}
	n.OnAction = func(r node.ActionRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "reset":
			self.c1 = nil
		}
		return nil, nil
	}
	return n
}
