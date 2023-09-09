//go:build ignore
// +build ignore

// core_gen generates boilerplate functions for meta structs looking
// at specfic field names and generating functions based on that.

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	elems := buildElements()
	tmplFile, err := os.Open("core_gen.in")
	if err != nil {
		panic(err)
	}
	tmplSrc, err := ioutil.ReadAll(tmplFile)
	if err != nil {
		panic(err)
	}
	t, err := template.New("code_gen").Parse(string(tmplSrc))
	if err != nil {
		panic(err)
	}
	codeGen, err := os.Create("./core_gen.go")
	if err != nil {
		panic(err)
	}
	defer codeGen.Close()
	if err := t.Execute(codeGen, elems); err != nil {
		panic(err)
	}
}

type elem struct {
	Name                string
	Parent              bool
	Description         bool
	Ident               bool
	Extensions          bool
	SecondaryExtensions bool
	Status              bool
	DataDefinitions     bool
	Groupings           bool
	Typedefs            bool
	Musts               bool
	When                bool
	IfFeatures          bool
	Actions             bool
	Notifications       bool
	Config              bool
	Mandatory           bool
	Units               bool
	MinMax              bool
	Unbounded           bool
	Default             bool
	Defaults            bool
	Type                bool
	OriginalParent      bool
	Clone               bool
	Augments            bool
	Presence            bool
	Unique              bool
	OrderedBy           bool
	ErrorMessage        bool
}

func buildElements() []*elem {
	fset := token.NewFileSet()
	src, err := parser.ParseFile(fset, "./core.go", nil, 0)
	if err != nil {
		panic(err)
	}
	var v visitor
	ast.Walk(&v, src)
	return v.elems
}

type visitor struct {
	ident string
	elem  *elem
	elems []*elem
}

func (v *visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	switch d := n.(type) {
	case *ast.Ident:
		v.ident = d.Name
	case *ast.StructType:
		v.elem = &elem{
			Name: v.ident,
		}
		v.elems = append(v.elems, v.elem)
		switch v.elem.Name {
		case "Rpc", "Must", "Leaf", "LeafList", "Any", "Uses", "Choice", "ChoiceCase":
			v.elem.Clone = true
		}
	case *ast.Field:
		// if specfic field names are found on structure, we flag
		// this struct as such which activates parts of the template
		// this assumes a specific field name convention is followed
		if len(d.Names) > 0 {
			switch d.Names[0].Name {
			case "parent":
				v.elem.Parent = true
			case "ident":
				v.elem.Ident = true
			case "desc":
				v.elem.Description = true
			case "extensions":
				v.elem.Extensions = true
			case "secondaryExtensions":
				v.elem.SecondaryExtensions = true
			case "status":
				v.elem.Status = true
			case "dataDefs":
				v.elem.DataDefinitions = true
				v.elem.Clone = true
			case "groupings":
				v.elem.Groupings = true
			case "typedefs":
				v.elem.Typedefs = true
			case "musts":
				v.elem.Musts = true
			case "when":
				v.elem.When = true
			case "ifs":
				v.elem.IfFeatures = true
			case "actions":
				v.elem.Actions = true
			case "notifications":
				v.elem.Notifications = true
			// assumes config too
			case "mandatoryPtr":
				v.elem.Mandatory = true
			case "configPtr":
				v.elem.Config = true
			// assumes max and min too
			case "minElementsPtr":
				v.elem.MinMax = true
			case "unboundedPtr":
				v.elem.Unbounded = true
			case "dtype":
				v.elem.Type = true
			case "originalParent":
				v.elem.OriginalParent = true
			case "augments":
				v.elem.Augments = true
			case "defaultVal":
				v.elem.Default = true
			case "defaultVals":
				v.elem.Defaults = true
			case "units":
				v.elem.Units = true
			case "presence":
				v.elem.Presence = true
			case "unique":
				v.elem.Unique = true
			case "orderedBy":
				v.elem.OrderedBy = true
			case "errorMessage":
				// assumes error-app-tag too
				v.elem.ErrorMessage = true
			}
		}
	}
	return v
}
