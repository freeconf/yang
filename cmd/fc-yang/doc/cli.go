package doc

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodes"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/render"
)

// Run "freeconf get ..." command
func Run() {
	var on featureParams
	var off featureParams

	moduleName := flag.String("module", "", "Module to be documented.")
	tmplPtr := flag.String("f", "none", "output format. available formats include "+
		"html, md, json or dot.")
	exportTemplatePtr := flag.Bool("x", false, "export the builting template to stdout. You can then edit "+
		"template and pass it back in using -t option.  Be sure to pick correct format.")
	useTemplatePtr := flag.String("t", "", "Use the template instead of the builtin template.")
	titlePtr := flag.String("title", "RESTful API", "Title.")
	imageLinkPtr := flag.String("img-link", "", "Link to image for HTML templates. Default is (module-name).svg.")
	flag.Var(&on, "on", "enable this feature.  You can specify -on multiple times to enable multiple features. You cannot specify both on and off however.")
	flag.Var(&off, "off", "disable this feature.  You can specify -off multiple times to disable multiple features. You cannot specify both on and off however.")

	flag.Parse()

	if *moduleName == "" {
		log.Fatal("missing module name")
	}

	imageLink := *imageLinkPtr
	if imageLink == "" {
		imageLink = *moduleName + ".svg"
	}

	var builder DocDefBuilder
	switch *tmplPtr {
	case "html":
		builder = &DocHtml{
			ImageLink: imageLink,
		}
	case "md":
		builder = &DocMarkdown{}
	case "dot":
		builder = &DocDot{}
	}
	if *exportTemplatePtr {
		if _, err := fmt.Print(builder.BuiltinTemplate()); err != nil {
			log.Fatal(err)
		}
		return
	}

	var err error
	var m *meta.Module
	var fs meta.FeatureSet
	if len(off) > 0 {
		if len(on) > 0 {
			log.Fatal("You cannot specify both on and off")
		}
		fs = meta.FeaturesOff(off)
	} else {
		fs = meta.FeaturesOn(on)
	}
	m, err = parser.LoadModuleWithFeatures(parser.YangPath(), *moduleName, "", fs)
	if err != nil {
		log.Fatal(err)
	}

	if *tmplPtr == "none" {
		ymod := parser.RequireModule(parser.YangPath(), "fc-yang")
		n := &nodes.JSONWtr{Out: os.Stdout, Pretty: true}
		if err = nodes.Schema(ymod, m).Root().InsertInto(n.Node()).LastErr; err != nil {
			log.Fatal(err)
		}
	} else {
		doc := &render.Doc{Title: *titlePtr}
		if err = doc.Build(m); err != nil {
			log.Fatal(err)
		}
		if *tmplPtr == "json" {
			ymod := parser.RequireModule(parser.YangPath(), "fc-doc")
			n := &nodes.JSONWtr{Out: os.Stdout, Pretty: true}
			b := node.NewBrowser(ymod, render.Api(doc))
			if err = b.Root().InsertInto(n.Node()).LastErr; err != nil {
				log.Fatal(err)
			}
		} else {
			var tmpl string
			if *useTemplatePtr != "" {
				f, err := os.Open(*useTemplatePtr)
				if err != nil {
					log.Fatal(err)
				}
				data, err := ioutil.ReadAll(f)
				if err != nil {
					log.Fatal(err)
				}
				tmpl = string(data)
			} else {
				tmpl = builder.BuiltinTemplate()
			}
			if err = builder.Generate(doc, tmpl, os.Stdout); err != nil {
				log.Fatal(err)
			}
		}
	}

	os.Exit(0)
}

type featureParams []string

func (f *featureParams) String() string {
	return strings.Join([]string(*f), ", ")
}

func (f *featureParams) Set(value string) error {
	*f = append(*f, value)
	return nil
}
