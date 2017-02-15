package main

import (
	"flag"
	"os"

	"io/ioutil"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

func main() {
	flag.Parse()
	data, err := ioutil.ReadAll(os.Stdin)
	m, err := yang.LoadModuleCustomImport(string(data), func(m *meta.Module, name string) error {
		m.Includes = append(m.Includes, name)
		return nil
	})
	if err != nil {
		panic(err)
	}
	b := node.SelectModule(m, false)
	if err := b.Root().InsertInto(node.NewJsonWriter(os.Stdout).Node()).LastErr; err != nil {
		panic(err)
	}
}
