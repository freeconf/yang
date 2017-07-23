package browse

import (
	"html/template"
	"io"
	"regexp"
	"strings"
)

type DocMarkdown struct {
}

func (self *DocMarkdown) Generate(doc *Doc, out io.Writer) error {
	funcMap := template.FuncMap{
		"repeat": strings.Repeat,
		"link":   docLink,
		"title":  docTitle,
		"type":   docFieldType,
		"path":   docPath,
		"desc":   mdCleanDescription,
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

func mdCleanDescription(d string) string {
	// Clear extra space at the beginning of a line otherwise it comes out
	// as code text.  While it would be nice to *not* strip these out, too many
	// unintential uses exist
	return regexp.MustCompile("(?m)^ *").ReplaceAllLiteralString(d, "")
}

const docMarkdown = `
{{$backtick := "\x60"}}
# {{.Doc.Title}}

{{range .Doc.Defs}}
## <a name="{{link .}}"></a>{{path .}}
{{desc .Meta.Description}}

{{range .Fields}}
  {{if .Def}}
* **[{{title .Meta}}](#{{link .Def}})**
  {{- else}}
* **{{title .Meta}}** {{$backtick}}{{type .}}{{$backtick}}
  {{- end}} - {{desc .Meta.Description}}. {{if .Details}} *{{.Details}}* {{end}}
{{end}}

{{if .Actions}}
### Actions:
{{range .Actions}}
* <a name="{{link .}}"></a>**{{path .Def}}{{title .Meta}}** - {{desc .Meta.Description}}
 
  {{if .InputFields}}
#### Input:

    {{- range .InputFields -}}
      {{- if .Expand}}
> * **{{title .Meta}}** - {{desc .Meta.Description}}
        {{- range .Expand}}
> {{repeat "   " .Level |noescape}} * **{{title .Meta}}** - {{desc .Meta.Description}} {{.Details}}
        {{- end -}}
		  {{- else}}
> * **{{title .Meta}}** {{$backtick}}{{type .}}{{$backtick}} - {{desc .Meta.Description}}
      {{- end -}}
    {{- end -}}
  {{- end}}


  {{if .OutputFields}}
#### Output:

    {{- range .OutputFields -}}
      {{- if .Expand}}
> * **{{title .Meta}}** - {{desc .Meta.Description}}
        {{- range .Expand}}
> {{repeat "   " .Level |noescape}} * **{{title .Meta}}** - {{desc .Meta.Description}} {{.Details}}
        {{- end -}}
		  {{- else}}
> * **{{title .Meta}}** {{$backtick}}{{type .}}{{$backtick}} - {{desc .Meta.Description}}
      {{- end -}}
    {{- end -}}
  {{- end}}

{{end}}
{{end}}

{{if .Events}}
### Events:
{{range .Events}}
* <a name="{{link .}}"></a>**{{path .Def}}{{title .Meta}}** - {{desc .Meta.Description}}

 {{if .Fields -}}
    {{- range .Fields -}}
      {{- if .Expand}}
> * **{{title .Meta}}** - {{desc .Meta.Description}}
        {{- range .Expand}}
> {{repeat "   " .Level |noescape}} * **{{title .Meta}}** - {{desc .Meta.Description}} {{.Details}}
        {{- end -}}
			{{- else}}	
> * **{{title .Meta}}** {{$backtick}}{{type .}}{{$backtick}} - {{desc .Meta.Description}}
      {{- end -}}
    {{- end -}}
  {{- end}}

{{end}}
{{end}}

{{end}}
`
