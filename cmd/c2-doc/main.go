package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/freeconf/gconf/node"

	"github.com/freeconf/gconf/nodes"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/meta/render"
	"github.com/freeconf/gconf/meta/yang"
)

/*
	Utility that generates documentation, diagrams or anything from YANG module
	file.
*/
var moduleNamePtr = flag.String("module", "", "Module to be documented.")
var tmplPtr = flag.String("builtin", "none", "use a built-in template for documentation generation : "+
	"html, md, json or dot.  Otherwise read template from stdin")
var titlePtr = flag.String("title", "RESTful API", "Title.")
var verbose = flag.Bool("verbose", false, "verbose")

func main() {
	flag.Parse()
	c2.DebugLog(*verbose)
	if moduleNamePtr == nil {
		chkErr(c2.NewErr("Missing module name"))
	}

	m, err := yang.LoadModule(yang.YangPath(), *moduleNamePtr)
	chkErr(err)

	if *tmplPtr == "none" {
		ymod := yang.RequireModule(yang.YangPath(), "yang")
		n := &nodes.JSONWtr{Out: os.Stdout, Pretty: true}
		chkErr(nodes.Schema(ymod, m).Root().InsertInto(n.Node()).LastErr)
	} else {
		doc := &render.Doc{Title: *titlePtr}
		chkErr(doc.Build(m))
		if *tmplPtr == "json" {
			ymod := yang.RequireModule(yang.YangPath(), "fc-doc")
			n := &nodes.JSONWtr{Out: os.Stdout, Pretty: true}
			b := node.NewBrowser(ymod, render.Api(doc))
			chkErr(b.Root().InsertInto(n.Node()).LastErr)
		} else {
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
		}
	}

	os.Exit(0)
}

func chkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}
}
