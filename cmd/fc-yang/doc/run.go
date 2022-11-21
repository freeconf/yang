package doc

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"
)

// Run "freeconf get ..." command
func Run() {
	var on featureParams
	var off featureParams
	fc.DebugLog(true)

	moduleName := flag.String("module", "", "Module to be documented.")
	tmplPtr := flag.String("f", "none", "output format. available formats include "+
		"html, md, json or dot.")
	exportTemplatePtr := flag.Bool("x", false, "export the builting template to stdout. You can then edit "+
		"template and pass it back in using -t option.  Be sure to pick correct format.")
	useTemplatePtr := flag.String("t", "", "Use the template instead of the builtin template.")
	titlePtr := flag.String("title", "RESTful API", "Title.")
	ypathArg := flag.String("ypath", os.Getenv("YANGPATH"), "Path to YANG files")
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

	var builder builder
	switch *tmplPtr {
	case "html":
		builder = &html{
			ImageLink: imageLink,
		}
	case "md":
		builder = &markdown{}
	case "dot":
		builder = &dot{}
	}
	if *exportTemplatePtr {
		if _, err := fmt.Print(builder.builtinTemplate()); err != nil {
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
	options := parser.Options{
		Features: fs,
	}
	ypath := source.Path(*ypathArg)
	m, err = parser.LoadModuleWithOptions(ypath, *moduleName, options)
	if err != nil {
		log.Fatalf("could not load %s. %s", *moduleName, err)
	}

	if *tmplPtr == "none" {
		ymod := parser.RequireModule(ypath, "fc-yang")
		n := &nodeutil.JSONWtr{Out: os.Stdout, Pretty: true}
		if err = nodeutil.Schema(ymod, m).Root().InsertInto(n.Node()).LastErr; err != nil {
			log.Fatal(err)
		}
	} else {
		d := &doc{Title: *titlePtr}
		if err = d.build(m); err != nil {
			log.Fatal(err)
		}
		if *tmplPtr == "json" {
			ymod := parser.RequireModule(ypath, "fc-doc")
			n := &nodeutil.JSONWtr{Out: os.Stdout, Pretty: true}
			b := node.NewBrowser(ymod, api(d))
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
				tmpl = builder.builtinTemplate()
			}
			if err = builder.generate(d, tmpl, os.Stdout); err != nil {
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
