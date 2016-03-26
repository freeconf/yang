package main

import (
	"flag"
	"fmt"
	"github.com/blitter/meta/yang"
	"github.com/blitter/node"
	"os"
)

var moduleName = flag.String("module", "", "Module name (w/o *.yang extension)")

func main() {
	flag.Parse()

	// TODO: Change this to a file-persistent store.
	if len(*moduleName) == 0 {
		fmt.Fprint(os.Stderr, "Required 'module' parameter missing\n")
		os.Exit(-1)
	}

	m, err := yang.LoadModule(yang.YangPath(), *moduleName)
	if err != nil {
		panic(err)
	}
	config := node.NewJsonWriter(os.Stdout).Node()
	c := node.NewContext()
	if err != nil {
		panic(err)
	}
	if err = c.Selector(node.SelectModule(m, false)).InsertInto(config).LastErr; err != nil {
		panic(err)
	}
}
