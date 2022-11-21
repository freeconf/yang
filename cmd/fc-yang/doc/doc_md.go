package doc

import (
	"fmt"
	"html/template"
	"io"
	"regexp"
	"strings"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
)

type markdown struct {
}

func (self *markdown) generate(d *doc, tmpl string, out io.Writer) error {
	funcMap := template.FuncMap{
		"repeat": strings.Repeat,
		"link":   docLink,
		"title":  docTitle,
		"title2": docTitle2,
		"type":   docFieldType,
		"path":   docPath,
		"desc":   mdCleanDescription,
		"args":   args,
		"last":   last,
		"noescape": func(s string) template.HTML {
			return template.HTML(s)
		},
		"jsonExampleType": jsonExampleType,
	}
	t := template.Must(template.New("fc.oc").Funcs(funcMap).Parse(tmpl))
	err := t.Execute(out, struct {
		Doc *doc
	}{
		Doc: d,
	})
	return err
}

func jsonExampleType(d meta.Definition) string {
	var example string
	l := d.(meta.Leafable)
	if l.Type().Format().IsNumeric() {
		example = "n"
	} else if l.Type().Format().Single() == val.FmtBool {
		example = "true|false"
	} else {
		example = "\"\""
	}
	if l.Type().Format().IsList() {
		example = fmt.Sprintf("[%s,...]", example)
	}
	return example
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
{{ $backtick := "\x60" }}
{{- define "data" }}
{{- $nextindent := printf "   %s" .indent }}
{{- $def := .def }}
{{.indent}}"{{$def.Meta.Ident}}":{{- if $def.IsList -}}[{{ end }}
  {{- if $def.Leafable -}}
      {{- noescape (jsonExampleType $def.Meta) -}}
  {{- else -}}
{
        {{- range $index, $f := $def.Fields }}
            {{- template "data" args "def" $f "indent" $nextindent -}}{{if last $index $def.Fields | not }},{{end}}
        {{- end }}
{{.indent}}}{{- if $def.IsList -}}]{{ end }}
  {{- end -}}
{{ end }}

{{- define "props" -}}
{{- $path := .path }}
{{- $def := .def }}
  {{- if .def.Leafable }}
> | {{.path}}{{.def.Meta.Ident}} | {{.def.ScalarType}}  |  {{desc .def.Meta.Description}} | {{.def.Details}} |
  {{- else -}}
        {{- range .def.Fields -}}
{{ $nextpath := printf "%s%s." $path $def.Meta.Ident }}        
{{- template "props" args "def" . "path" $nextpath -}}
        {{- end -}}
  {{- end -}}
{{- end -}}
# {{.Doc.Title}}

{{- define "crud" -}}
{{ $backtick := "\x60" }}
{{ $def := .def }}
{{ $byKey := .byKey }}
{{ $showListIdent := (and $def.IsList (not $byKey)) }}
{{- $fields := $def.Fields}}
{{- $writeable := $def.WriteableFields}}
{{- $id := printf "%s%s" (path $def.Parent) ($def.Meta.Ident) -}}
{{- if $byKey }}{{- $id = printf "%s%s" (path $def.Parent) (title2 $def.Meta) -}}{{ end -}}
<details>
 <summary><a name="{{$id}}"></a><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:{{$id}}</b></code> {{desc $def.Meta.Description}}</summary>

{{ if $def.AllFieldsWritable -}}
#### GET Response Data / PUT, POST Request Data
{{- else }}
#### GET Response Data
{{- end }}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
{ {{- if $showListIdent -}}"{{$def.Meta.Ident}}":[{{ end }}
	{{- range $index, $f := $fields -}}
      {{- template "data" args "def" $f "indent" "  " -}}{{if last $index $fields | not }},{{end}}
  {{- end }}
{{- if $showListIdent -}}},...]{{ end }}}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}

{{ if not $def.AllFieldsWritable -}}
#### PUT, POST Request Data
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
{ {{- if $showListIdent -}}"{{$def.Meta.Ident}}":[{{ end }}
	{{- range $index, $f := $writeable -}}
      {{- template "data" args "def" $f "indent" "  " -}}{{if last $index $writeable | not }},{{end}}
  {{- end }}
{{- if $showListIdent -}}},...]{{ end }}}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
{{- end }}

#### Data Details

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
{{- range $def.Fields}}
{{- template "props" args "def" . "path" ""}}
{{- end}}

#### Responses
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | {{$backtick}}POST{{$backtick}}       |  *JSON data*   | - none -      |
> | {{$backtick}}PUT{{$backtick}}       |  *JSON data*   | - none -      |
> | {{$backtick}}GET{{$backtick}}       |  - none -      | *JSON data*   |
> | {{$backtick}}DELETE{{$backtick}}     |  - none -      | - none -      |

#### HTTP response codes
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

#### Examples
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
# retrieve data
curl https://server/restconf/data/acc:{{path $def.Parent}}{{$def.Meta.Ident}}

# update existing data
curl -X PUT -d @data.json https://server/restconf/data/acc:{{path $def.Parent}}{{$def.Meta.Ident}}

# create new data
curl -X POST -d @data.json https://server/restconf/data/acc:{{path $def.Parent}}{{$def.Meta.Ident}}

# delete current data
curl -X DELETE https://server/restconf/data/acc:{{path $def.Parent}}{{$def.Meta.Ident}}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
</details>
{{- end -}}

## <a name="{{link .Doc.Module}}"></a>{{path .Doc.Module}}
{{desc .Doc.Module.Meta.Description}}

{{range .Doc.DataDefs}}
{{template "crud" args "def" . "byKey" false}}
{{ if .IsList }}
{{template "crud" args "def" . "byKey" true }}
{{ end }}
{{end}}

{{if .Doc.Actions}}
  {{range .Doc.Actions}}
  {{- $id := printf "%s%s" (path .Parent) (.Meta.Ident) -}}
<details>
 <summary><a name="{{$id}}"></a><code>[POST]</code> <code><b>restconf/data/acc:{{path .Parent}}{{.Meta.Ident}}</b></code> {{desc .Meta.Description}}</summary>

##### Request Body
    {{if .Input}}
      {{ $fields := .Input.Expand }}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
{
	    {{- range $index, $def := $fields -}}
        {{- template "data" args "def" $def "indent" "  " -}}{{if last $index $fields | not }},{{end}}
      {{- end }}
}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}

#### Request Body Details

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
{{- range $fields }}
{{- template "props" args "def" . "path" ""}}
{{- end}}
    {{else}}
  *none*
    {{end}}

##### Response Body
    {{if .Output}}
      {{ $fields := .Output.Expand }}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
{
	    {{- range $index, $def := $fields -}}
        {{- template "data" args "def" $def "indent" "  " -}}{{if last $index $fields | not }},{{end}}
      {{- end }}
}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}

#### Response Body Details

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
{{- range $fields }}
{{- template "props" args "def" . "path" ""}}
{{- end}}
    {{else}}
  *none*
    {{end}}

  <details><summary>more</summary>

#### HTTP response codes
> | http code |  reason for code |
> |-----------|------------------|
> | 200       | success          |
> | 401       | not authorized   |
> | 400       | invalid request  |
> | 404       | data does not exist |
> | 500       | internal error   |

#### Examples
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
# create new data
curl -X POST -d @request.json https://server/restconf/data/acc:{{path .Parent}}{{title .Meta}}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
  </details>

</details>
  {{end}}
{{end}}

{{if .Doc.Events}}
  {{range .Doc.Events}}
{{- $id := printf "%s%s" (path .Parent) (title2 .Meta) -}}
<details>
 <summary><a name="{{$id}}"></a><code>[GET]</code> <code><b>restconf/data/acc:{{path .Parent}}{{.Meta.Ident}}</b></code> {{desc .Meta.Description}}</summary>

##### Response Stream [SSE Format](https://html.spec.whatwg.org/multipage/server-sent-events.html#server-sent-events)
{{ $fields := .Expand }}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
data: {
	    {{- range $index, $def := $fields -}}
        {{- template "data" args "def" $def "indent" "  " -}}{{if last $index $fields | not }},{{end}}
      {{- end -}}
}\n
\n
data: {{$backtick}}{  ... next message with same format ... }{{$backtick}}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}

#### Response Body Details

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
{{- range $fields }}
{{- template "props" args "def" . "path" ""}}
{{- end}}

</details>
  {{end}}
{{end}}
`
