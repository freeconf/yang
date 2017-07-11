package browse

import (
	"html/template"
	"io"
	"strings"
)

// self.tmpl = docMarkdown
// self.Delim = "/"
// self.ListKeyFmt = "={%v}"
// self.TitleFilter = func(s string) string {
// 	return s
// }
type DocMarkdown struct {
}

func (self *DocMarkdown) Generate(doc *Doc, out io.Writer) error {
	funcMap := template.FuncMap{
		"repeat": strings.Repeat,
		"link":   docLink,
		"title":  docTitle,
		"type":   docFieldType,
		"path":   docPath,
		"noescape": func(s string) template.HTML {
			return template.HTML(s)
		},
	}
	t := template.Must(template.New("c2doc").Funcs(funcMap).Parse(docMarkdown))
	err := t.Execute(out, struct {
		Doc *Doc
	}{
		Doc: doc,
	})
	return err
}

const docMarkdown = `
{{$backtick := "\x60"}}
# {{.Doc.Title}}

{{range .Doc.Defs}}
## <a name="{{link .}}"></a>{{path .}}
{{.Meta.Description}}

{{range .Fields}}
  {{if .Def}}
* **[{{title .Meta}}](#{{link .Def}})**
  {{- else}}
* **{{title .Meta}}** {{$backtick}}{{type .}}{{$backtick}}
  {{- end}} - {{.Meta.Description}}. {{if .Details}} *{{.Details}}* {{end}}
{{end}}

{{if .Actions}}
### Actions:
{{range .Actions}}
* <a name="{{link .}}"></a>**{{path .Def}}{{title .Meta}}** - {{.Meta.Description}}
 
  {{if .InputFields}}
#### Input:

    {{- range .InputFields -}}
      {{- if .Expand}}
> * **{{title .Meta}}** - {{.Meta.Description}}
        {{- range .Expand}}
> {{repeat "   " .Level |noescape}} * **{{title .Meta}}** - {{.Meta.Description}} {{.Details}}
        {{- end -}}
		  {{- else}}
> * **{{title .Meta}}** {{$backtick}}{{type .}}{{$backtick}} - {{.Meta.Description}}					
      {{- end -}}
    {{- end -}}
  {{- end}}


  {{if .OutputFields}}
#### Output:

    {{- range .OutputFields -}}
      {{- if .Expand}}
> * **{{title .Meta}}** - {{.Meta.Description}}
        {{- range .Expand}}
> {{repeat "   " .Level |noescape}} * **{{title .Meta}}** - {{.Meta.Description}} {{.Details}}
        {{- end -}}
		  {{- else}}
> * **{{title .Meta}}** {{$backtick}}{{type .}}{{$backtick}} - {{.Meta.Description}}					
      {{- end -}}
    {{- end -}}
  {{- end}}

{{end}}
{{end}}

{{if .Events}}
### Events:
{{range .Events}}
* <a name="{{link .}}"></a>**{{path .Def}}{{title .Meta}}** - {{.Meta.Description}}

 {{if .Fields -}}
    {{- range .Fields -}}
      {{- if .Expand}}
> * **{{title .Meta}}** - {{.Meta.Description}}
        {{- range .Expand}}
> {{repeat "   " .Level |noescape}} * **{{title .Meta}}** - {{.Meta.Description}} {{.Details}}
        {{- end -}}
			{{- else}}	
> * **{{title .Meta}}** {{$backtick}}{{type .}}{{$backtick}} - {{.Meta.Description}}
      {{- end -}}
    {{- end -}}
  {{- end}}

{{end}}
{{end}}

{{end}}
`
