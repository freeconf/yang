package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/freeconf/gconf/meta"
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
var tmplPtr = flag.String("f", "none", "output format. available formats include "+
	"html, md, json or dot.")
var exportTemplatePtr = flag.Bool("x", false, "export the builting template to stdout. You can then edit "+
	"template and pass it back in using -t option.  Be sure to pick correct format.")
var useTemplatePtr = flag.String("t", "", "Use the template instead of the builtin template.")
var titlePtr = flag.String("title", "RESTful API", "Title.")
var verbose = flag.Bool("verbose", false, "verbose")
var on featureParams
var off featureParams

func init() {
	flag.Var(&on, "on", "enable this feature.  You can specify -on multiple times to enable multiple features. You cannot specify both on and off however.")
	flag.Var(&off, "off", "disable this feature.  You can specify -off multiple times to disable multiple features. You cannot specify both on and off however.")
}

type featureParams []string

func (f *featureParams) String() string {
	return strings.Join([]string(*f), ", ")
}

func (f *featureParams) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func main() {
	flag.Parse()
	c2.DebugLog(*verbose)

	var builder render.DocDefBuilder
	switch *tmplPtr {
	case "html":
		builder = &render.DocHtml{}
	case "md":
		builder = &render.DocMarkdown{}
	case "dot":
		builder = &render.DocDot{}
	}
	if *exportTemplatePtr {
		_, err := fmt.Print(builder.BuiltinTemplate())
		chkErr(err)
		return
	}

	if moduleNamePtr == nil {
		chkErr(c2.NewErr("Missing module name"))
	}

	var err error
	var m *meta.Module
	var fs meta.FeatureSet
	if len(off) > 0 {
		if len(on) > 0 {
			chkErr(errors.New("You cannot specify both on and off"))
		}
		fs = meta.FeaturesOff(off)
	} else {
		fs = meta.FeaturesOn(on)
	}
	m, err = yang.LoadModuleWithFeatures(yang.YangPath(), *moduleNamePtr, "", fs)
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
			var tmpl string
			if *useTemplatePtr != "" {
				f, err := os.Open(*useTemplatePtr)
				chkErr(err)
				data, err := ioutil.ReadAll(f)
				chkErr(err)
				tmpl = string(data)
			} else {
				tmpl = builder.BuiltinTemplate()
			}
			chkErr(builder.Generate(doc, tmpl, os.Stdout))
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
