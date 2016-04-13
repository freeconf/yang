package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/c2g/meta/yang"
	"github.com/c2g/browse"
	"github.com/c2g/meta"
	"strings"
)

var moduleNamePtr = flag.String("module", "", "Module to be documented.")
var titlePtr = flag.String("title", "RESTful API", "Title.")

func main() {
	flag.Parse()
	if moduleNamePtr == nil {
		fmt.Fprintf(os.Stderr, "Missing module name")
		os.Exit(-1)
	}

	moduleNames := strings.Split(*moduleNamePtr, ",")
	doc := &browse.Doc{Title:*titlePtr}
	for _, moduleName := range moduleNames {
		m, err := yang.LoadModule(meta.MultipleSources(yang.InternalYang(), yang.YangPath()), moduleName)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		doc.Build(m)
	}
	if err := doc.Generate(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
	os.Exit(0)
}

