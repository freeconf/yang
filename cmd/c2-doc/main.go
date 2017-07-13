package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/c2stack/c2g/browse"
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta/yang"
)

var moduleNamePtr = flag.String("module", "", "Module to be documented.")
var tmplPtr = flag.String("builtin", "", "use a built-in template for documentation generation : "+
	"html, md or dot.  Otherwise read template from stdin")
var titlePtr = flag.String("title", "RESTful API", "Title.")
var verbose = flag.Bool("verbose", false, "verbose")

func main() {
	flag.Parse()
	c2.DebugLog(*verbose)
	if moduleNamePtr == nil {
		fmt.Fprintf(os.Stderr, "Missing module name")
		os.Exit(-1)
	}

	m, err := yang.LoadModule(yang.YangPath(), *moduleNamePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	doc := &browse.Doc{Title: *titlePtr}
	doc.Build(m)
	var builder browse.DocDefBuilder
	switch *tmplPtr {
	case "html":
		builder = &browse.DocHtml{}
	case "md":
		builder = &browse.DocMarkdown{}
	case "dot":
		builder = &browse.DocDot{}
	}
	if err := builder.Generate(doc, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
	os.Exit(0)
}
