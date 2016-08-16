package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/dhubler/c2g/meta/yang"
	"github.com/dhubler/c2g/browse"
	"strings"
)

var moduleNamePtr = flag.String("module", "", "Module to be documented.")
var appendNamesPtr = flag.String("append", "", "Append module to API doc.  Comma separated list.")
var tmplPtr = flag.String("tmpl", "html", "html or dot")
var titlePtr = flag.String("title", "RESTful API", "Title.")

func main() {
	flag.Parse()
	if moduleNamePtr == nil {
		fmt.Fprintf(os.Stderr, "Missing module name")
		os.Exit(-1)
	}

	m, err := yang.LoadModule(yang.YangPath(), *moduleNamePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	appendNames := strings.Split(*appendNamesPtr, ",")
	for _, appendName := range appendNames {
		if len(appendName) == 0 {
			continue
		}
		a, err := yang.LoadModule(yang.YangPath(), appendName)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		m.AddMeta(a)
	}
	doc := &browse.Doc{Title:*titlePtr}
	doc.Build(m, *tmplPtr)
	if err := doc.Generate(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
	os.Exit(0)
}

