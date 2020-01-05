package doc

import (
	"html/template"
	"io"
	"regexp"
	"strings"
)

type markdown struct {
}

func (self *markdown) generate(d *doc, tmpl string, out io.Writer) error {
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
	t := template.Must(template.New("fc.oc").Funcs(funcMap).Parse(tmpl))
	err := t.Execute(out, struct {
		Doc *doc
	}{
		Doc: d,
	})
	return err
}

func (self *markdown) builtinTemplate() string {
	return docMarkdown
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


## <a name="{{link .Doc.Module}}"></a>{{path .Doc.Module}}
{{desc .Doc.Module.Meta.Description}}

{{range .Doc.Module.Expand}}
* **[{{title .Meta}}]
  {{- if .Leafable -}}
     ** {{$backtick}}{{type .}}{{$backtick}}
  {{- else -}}
     (#{{link .}})**
  {{- end -}} 
  - {{desc .Meta.Description}}. {{if .Details}} *{{.Details}}* {{end}}
{{end}}

{{if .Doc.Actions}}
### Actions:
{{range .Doc.Actions}}
* <a name="{{link .}}"></a>**{{path .Parent}}{{title .Meta}}** - {{desc .Meta.Description}}
 
  {{if .Input}}
#### Input:

	{{ range .Input.Expand }}
> {{repeat "   " .Level |noescape}} * **{{title .Meta}}**	
		{{- if .Leafable -}}
		   {{$backtick}}{{type .}}{{$backtick}} - {{desc .Meta.Description}}
		{{- else -}}
		- {{desc .Meta.Description}}
      	{{- end -}}
	{{ end }}	
  {{- end}}


  {{if .Output}}
#### Output:

	{{ range .Output.Expand }}
> {{repeat "   " .Level |noescape}} * **{{title .Meta}}**	
		{{- if .Leafable -}}
		   {{$backtick}}{{type .}}{{$backtick}} - {{desc .Meta.Description}}
		{{- else -}}
		- {{desc .Meta.Description}}
      	{{- end -}}
	{{ end }}	
  {{- end}}

{{end}}

{{if .Doc.Events}}
### Events:
{{range .Doc.Events}}
* <a name="{{link .}}"></a>**{{path .}}{{title .Meta}}** - {{desc .Meta.Description}}

  {{ range .Expand }}
> {{repeat "   " .Level |noescape}} * **{{title .Meta}}**	
		{{- if .Leafable -}}
		   {{$backtick}}{{type .}}{{$backtick}} - {{desc .Meta.Description}}
		{{- else -}}
		- {{desc .Meta.Description}}
      	{{- end -}}
	{{ end }}	
  {{- end}}

{{end}}
{{end}}

`
