package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/freeconf/c2g/c2"
	"github.com/freeconf/c2g/meta/render"
	"github.com/freeconf/c2g/meta/yang"
)

/*
	Utility that generates documentation, diagrams or anything from YANG module
	file.
*/
var moduleNamePtr = flag.String("module", "", "Module to be documented.")
var tmplPtr = flag.String("builtin", "", "use a built-in template for documentation generation : "+
	"html, md or dot.  Otherwise read template from stdin")
var titlePtr = flag.String("title", "RESTful API", "Title.")
var verbose = flag.Bool("verbose", false, "verbose")

func main() {
	flag.Parse()
	c2.DebugLog(*verbose)
	if moduleNamePtr == nil {
		chkErr(c2.NewErr("Missing module name"))
	}

	m, err := yang.LoadModule(yang.YangPath(), *moduleNamePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	doc := &render.Doc{Title: *titlePtr}
	chkErr(doc.Build(m))
	var builder render.DocDefBuilder
	switch *tmplPtr {
	case "html":
		builder = &render.DocHtml{}
	case "md":
		builder = &render.DocMarkdown{}
	case "dot":
		builder = &render.DocDot{}
	}
	chkErr(builder.Generate(doc, os.Stdout))
	os.Exit(0)
}

func chkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
}
