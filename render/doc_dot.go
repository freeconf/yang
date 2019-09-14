package render

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/freeconf/yang/meta"
)

type DocDot struct {
}

func (self *DocDot) Generate(doc *Doc, tmpl string, out io.Writer) error {
	funcMap := template.FuncMap{
		"repeat":  strings.Repeat,
		"id":      dotId,
		"title":   dotTitle,
		"type":    docFieldType,
		"details": dotDetails,
		"noescape": func(s string) template.HTML {
			return template.HTML(s)
		},
	}
	t := template.Must(template.New("c2doc").Funcs(funcMap).Parse(tmpl))
	err := t.Execute(out, struct {
		Doc *Doc
	}{
		Doc: doc,
	})
	return err
}

func (self *DocDot) BuiltinTemplate() string {
	return docDot
}

func dotId(o interface{}) string {
	switch x := o.(type) {
	case meta.Identifiable:
		return strings.Replace(x.Ident(), "-", "_", -1)
	case *DocDef:
		if x.Parent == nil {
			return dotId(x.Meta)
		}
		return dotId(x.Parent) + "_" + dotId(x.Meta)
	case *DocAction:
		return dotId(x.Def) + "_" + dotId(x.Meta)
	case *DocEvent:
		return dotId(x.Def) + "_" + dotId(x.Meta)
	}
	panic(fmt.Sprintf("not supported %T", o))
}

func dotTitle(m meta.Identifiable) string {
	return escape("{}", "\\")(docTitle(m))
}

func dotDetails(f *DocField) string {
	var details []string
	if hasDets, ok := f.Meta.(meta.HasDetails); ok {
		if !hasDets.Config() {
			details = append(details, "r/o")
		}
		if hasDets.Mandatory() {
			details = append(details, "m")
		}
	}
	if f.Case != nil {
		details = append(details, f.Case.Ident()+"?")
	}
	if len(details) == 0 {
		return ""
	}
	return " (" + strings.Join(details, ",") + ")"
}

const docDot = `digraph G {
        fontname = "Bitstream Vera Sans"
        fontsize = 8

        node [
                fontname = "Bitstream Vera Sans"
                fontsize = 8
                shape = "record"
        ]

        edge [
                fontname = "Bitstream Vera Sans"
                fontsize = 8
        ]

{{range .Doc.Defs}}
       {{id .}} [
         label = "{
           {{- title .Meta}}|
           {{- range .Fields}}
             {{- if type . -}}
               {{- title .Meta}} : {{type .}}{{details .}}\l
             {{- end -}}
           {{- end -}}
         }"
       ]
{{$x := id .}}

{{range .Actions}}
       {{id .}} [
         label = "{
           {{- title .Meta}} (action)|
           {{- if .InputFields}}Input|
		{{- range .InputFields}}&#32;&#32;{{title .Meta}} : {{type .}}\l{{end -}}|
           {{- end -}}
           {{- if .OutputFields}}Output|
		{{- range .OutputFields}}&#32;&#32;{{title .Meta}} : {{type .}}\l{{end -}}
           {{- end -}}
         }"
         color = "#b64ff7"
       ]
       {{$x}} -> {{id .}} [
         style = "dashed"
         color = "#b64ff7"
       ]
{{end}}

{{range .Events}}
       {{id .}} [
         label = "{
           {{- title .Meta}} (notification)|
           {{- if .Fields}}
		{{- range .Fields}}{{title .Meta}} : {{type .}}\l{{end -}}
           {{- end -}}
         }"
         color = "#4fb32e"
       ]
       {{$x}} -> {{id .}} [
         style = "dashed"
         color = "#4fb32e"
       ]
{{end}}

{{end}}

{{range .Doc.Defs}}
  {{$x := id .}}
  {{- range .Fields}}
    {{if .Def -}}
       {{$x}} -> {{id .Def}} [
         label="{{details .}}"
       ]
    {{- end}}
  {{- end}}
{{end}}

}
`
