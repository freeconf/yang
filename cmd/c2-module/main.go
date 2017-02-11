package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

var moduleName = flag.String("module", "", "Module name (w/o *.yang extension)")

func main() {
	flag.Parse()

	// TODO: Change this to a file-persistent store.
	if len(*moduleName) == 0 {
		fmt.Fprint(os.Stderr, "Required 'module' parameter missing\n")
		os.Exit(-1)
	}
	yangPath := yang.YangPath()
	m, err := yang.LoadModule(yangPath, *moduleName)
	if err != nil {
		panic(err)
	}
	b := node.SelectModule(yangPath, m, false)
	if err := b.Root().InsertInto(node.NewJsonWriter(os.Stdout).Node()).LastErr; err != nil {
		panic(err)
	}
}
