package doc

import (
	"html/template"
	"io"
	"strings"

	"github.com/freeconf/yang/meta"
)

type dot struct {
}

func (self *dot) generate(d *doc, tmpl string, out io.Writer) error {
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
	t := template.Must(template.New("fc.oc").Funcs(funcMap).Parse(tmpl))
	err := t.Execute(out, struct {
		Doc *doc
	}{
		Doc: d,
	})
	return err
}

func (self *dot) builtinTemplate() string {
	return docDot
}

func dotId(d *def) string {
	p := d.Parent
	id := d.Meta.Ident()
	for p != nil {
		id = p.Meta.Ident() + "_" + id
		p = p.Parent
	}
	return strings.Replace(id, "-", "_", -1)
}

func dotTitle(m meta.Identifiable) string {
	return escape("{}", "\\")(docTitle(m))
}

func dotDetails(f *def) string {
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

{{range .Doc.DataDefs}}
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
           {{- if .Input}}Input|
		{{- range .Input.Fields}}&#32;&#32;{{title .Meta}} : {{type .}}{{details .}}\l{{end -}}|
           {{- end -}}
           {{- if .Output}}Output|
		{{- range .Output.Fields}}&#32;&#32;{{title .Meta}} : {{type .}}{{details .}}\l{{end -}}
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
		     {{- range .Fields}}{{title .Meta}} : {{type .}}{{details .}}\l{{end -}}
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

{{range .Doc.DataDefs}}
  {{$x := id .}}
  {{- range .Fields}}
    {{if not .Leafable -}}
       {{$x}} -> {{id .}} [
         label="{{details .}}"
       ]
    {{- end}}
  {{- end}}
{{end}}

}
`
