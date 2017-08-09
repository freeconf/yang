package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"io/ioutil"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/render"
	"github.com/c2stack/c2g/meta/yang"
)

var moduleNamePtr = flag.String("module", "", "Module to be loaded.")
var yangPath = flag.String("yangPath", "", "Path to all yang files e.g.  foo:bar:bleep ")

// stdin/out can be problematic w/go generate integration so provide optional ways
var inPtr = flag.String("in", "", "template file, otherwise stdin is used")
var outPtr = flag.String("out", "", "output file otherwise stdout is used")

var verbose = flag.Bool("verbose", false, "verbose")

func main() {
	flag.Parse()
	c2.DebugLog(*verbose)
	if *moduleNamePtr == "" {
		fmt.Fprintf(os.Stderr, "module required")
		os.Exit(-1)
	}
	if *yangPath == "" {
		fmt.Fprintf(os.Stderr, "yangPath required")
		os.Exit(-1)
	}

	m, err := yang.LoadModule(meta.PathStreamSource(*yangPath), *moduleNamePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	doc := &render.Doc{}
	doc.Build(m)
	builder := &codeGen{}

	var in io.Reader
	if *inPtr == "" {
		in = os.Stdin
	} else {
		if fIn, err := os.Open(*inPtr); err != nil {
			panic(err)
		} else {
			defer fIn.Close()
			in = fIn
		}
	}

	var out io.Writer
	if *outPtr == "" {
		out = os.Stdout
	} else {
		if fOut, err := os.Create(*outPtr); err != nil {
			panic(err)
		} else {
			defer fOut.Close()
			out = fOut
		}
	}

	if err := builder.Generate(doc, in, out); err != nil {
		panic(err)
	}
}

type codeGen struct {
}

func (self *codeGen) Generate(doc *render.Doc, in io.Reader, out io.Writer) error {
	funcMap := template.FuncMap{
		"repeat":      strings.Repeat,
		"publicIdent": goPublic,
		"publicType":  goPublicType,
	}
	buff, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	t := template.Must(template.New("gen").Funcs(funcMap).Parse(string(buff)))
	return t.Execute(out, struct {
		Doc *render.Doc
	}{
		Doc: doc,
	})
}

var goTypes map[string]string

func init() {
	goTypes = map[string]string{
		"int32":     "int",
		"int64":     "int64",
		"boolean":   "bool",
		"decimal64": "float64",
	}
}

func goPublic(name string) string {
	return strings.ToUpper(name[0:1]) + name[1:]
}

func goPublicType(f *render.DocField) string {
	var t string
	if meta.IsLeaf(f.Meta) {
		if alias, hasAlias := goTypes[f.Type]; hasAlias {
			t = alias
		} else {
			t = f.Type
		}
		if _, isLeafList := f.Meta.(*meta.LeafList); isLeafList {
			t = "[]" + t
		}
	} else {
		t = "*" + goPublic(f.Meta.GetIdent())
		if meta.IsList(f.Meta) {
			t = "[]" + t
		}
	}
	return t
}
