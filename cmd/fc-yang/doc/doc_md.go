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
		example = "0"
	} else if l.Type().Format().Single() == val.FmtBool {
		example = "false"
	} else {
		example = "\"\""
	}
	if l.Type().Format().IsList() {
		example = fmt.Sprintf("[%s, \"...\"]", example)
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

{{- define "crud" -}}
{{ $backtick := "\x60" }}
{{ $def := .def }}
{{ $byKey := .byKey }}
{{ $showListIdent := (and $def.IsList (not $byKey)) }}
{{- $fields := $def.Fields}}
{{- $writeable := $def.WriteableFields}}
{{- $path := printf "%s%s" (path $def.Parent) ($def.Meta.Ident) -}}
{{- if $byKey }}{{- $path = printf "%s%s" (path $def.Parent) (title2 $def.Meta) -}}{{ end -}}
<details>
 <summary><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:{{$path}}</b></code> {{desc $def.Meta.Description}}</summary>

#### {{$path}}

{{ if $def.AllFieldsWritable -}}
**GET Response Data / PUT, POST Request Data**
{{- else }}
**GET Response Data**
{{- end }}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}json
{ {{- if $showListIdent -}}"{{$def.Meta.Ident}}":[{{ end }}
	{{- range $index, $f := $fields -}}
      {{- template "data" args "def" $f "indent" "  " -}}{{if last $index $fields | not }},{{end}}
  {{- end }}
{{- if $showListIdent -}}}, {"..."}]{{ end }}}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}

{{ if and (not $def.AllFieldsWritable) (gt (len $writeable) 0) -}}
**PUT, POST Request Data**
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}json
{ {{- if $showListIdent -}}"{{$def.Meta.Ident}}":[{{ end }}
	{{- range $index, $f := $writeable -}}
      {{- template "data" args "def" $f "indent" "  " -}}{{if last $index $writeable | not }},{{end}}
  {{- end }}
{{- if $showListIdent -}}},{"..."}]{{ end }}}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
{{- end }}

**Data Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
{{- range $def.Fields}}
{{- template "props" args "def" . "path" ""}}
{{- end}}

**Responses**
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | {{$backtick}}POST{{$backtick}}       |  *JSON data*   | - none -      |
> | {{$backtick}}PUT{{$backtick}}       |  *JSON data*   | - none -      |
> | {{$backtick}}GET{{$backtick}}       |  - none -      | *JSON data*   |
> | {{$backtick}}DELETE{{$backtick}}     |  - none -      | - none -      |

**HTTP response codes**
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

**Examples**
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}bash
# retrieve data
curl https://server/restconf/data/acc:{{$path}}

# update existing data
curl -X PUT -d @data.json https://server/restconf/data/acc:{{$path}}

# create new data
curl -X POST -d @data.json https://server/restconf/data/acc:{{$path}}

# delete current data
curl -X DELETE https://server/restconf/data/acc:{{$path}}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
</details>
{{- end -}}

# {{.Doc.Title}}

{{desc .Doc.Module.Meta.Description}}

<details><summary>API Usage Notes:</summary>

#### General API Usage Notes
* {{$backtick}}DELETE{{$backtick}} implementation may be disallowed or ignored depending on the context
* Lists use {{$backtick}}../path={key}/...{{$backtick}} instead of {{$backtick}}.../path/key/...{{$backtick}} to avoid API name collision

#### {{$backtick}}GET{{$backtick}} Query Parameters

These parameters can be combined.

> | param                            | description | example |
> |----------------------------------|-------------|---------|
> | {{$backtick}}content=[non-config\|config]{{$backtick}} | Show only read-only fields or only read/write fields |   {{$backtick}}.../path?content=config{{$backtick}}|
> | {{$backtick}}fields=field1;field2{{$backtick}} | Return a portion of the data limited to fields listed | {{$backtick}}.../path?fields=user%2faddress{{$backtick}} |
> | {{$backtick}}depth=n{{$backtick}} | Return a portion of the data limited to depth of the hierarchy | {{$backtick}}.../path?depth=1{{$backtick}}
> | {{$backtick}}fc.xfields=field1;fields{{$backtick}} | Return a portion of the data excluding the fields listed | {{$backtick}}.../path?fc.xfields=user%2faddress{{$backtick}} |
> | {{$backtick}}fc.range=field!{startRow}-[{endRow}]{{$backtick}} | For lists, return only limited number of rows | {{$backtick}}.../path?fc.range=user!10-20{{$backtick}} 

</details>

{{range .Doc.DataDefs}}
{{template "crud" args "def" . "byKey" false}}
{{ if .IsList }}
{{template "crud" args "def" . "byKey" true }}
{{ end }}
{{end}}

{{if .Doc.Actions}}
  {{range .Doc.Actions}}
  {{- $path := printf "%s%s" (path .Parent) (.Meta.Ident) -}}
<details>
 <summary><code>[POST]</code> <code><b>restconf/data/acc:{{path .Parent}}{{.Meta.Ident}}</b></code> {{desc .Meta.Description}}</summary>
 
#### {{$path}}

 **Request Body**
    {{if .Input}}
      {{ $fields := .Input.Expand }}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}json
{
	  {{- range $index, $def := $fields -}}
        {{- template "data" args "def" $def "indent" "  " -}}{{if last $index $fields | not }},{{end}}
      {{- end }}
}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}

**Request Body Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
{{- range $fields }}
{{- template "props" args "def" . "path" ""}}
{{- end}}
    {{else}}
  *none*
    {{end}}

**Response Body**
    {{if .Output}}
      {{ $fields := .Output.Expand }}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}json
{
	  {{- range $index, $def := $fields -}}
        {{- template "data" args "def" $def "indent" "  " -}}{{if last $index $fields | not }},{{end}}
      {{- end }}
}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}

**Response Body Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
{{- range $fields }}
{{- template "props" args "def" . "path" ""}}
{{- end}}
    {{else}}
  *none*
    {{end}}

**HTTP response codes**

> | http code |  reason for code |
> |-----------|------------------|
> | 200       | success          |
> | 401       | not authorized   |
> | 400       | invalid request  |
> | 404       | data does not exist |
> | 500       | internal error   |

**Examples**
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}bash
# call function
curl -X POST {{if and .Input (gt (len .Input.Expand) 0)}}-d @request.json]{{- end}} https://server/restconf/data/acc:{{$path}}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
  </details>

  {{end}}
{{end}}

{{if .Doc.Events}}
  {{range .Doc.Events}}
{{- $path := printf "%s%s" (path .Parent) (title2 .Meta) -}}
<details>
 <summary><code>[GET]</code> <code><b>restconf/data/acc:{{path .Parent}}{{.Meta.Ident}}</b></code> {{desc .Meta.Description}}</summary>

#### {{$path}}

**Response Stream** [SSE Format](https://html.spec.whatwg.org/multipage/server-sent-events.html#server-sent-events)
{{ $fields := .Expand }}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}
data: {first JSON message all on one line followed by 2 CRLFs}

data: {next JSON message with same format all on one line ...}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}

Each JSON message would have following data
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}json
{
	{{- range $index, $def := $fields -}}
		{{- template "data" args "def" $def "indent" "  " -}}{{if last $index $fields | not }},{{end}}
	{{- end}}
}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}

**Response Body Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
{{- range $fields }}
{{- template "props" args "def" . "path" ""}}
{{- end}}

**Example**
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}bash
# retrieve data stream, adjust timeout for slower streams
curl -N https://server/restconf/data/acc:{{$path}}
{{$backtick}}{{$backtick}}{{$backtick}}{{$backtick}}

</details>
  {{end}}
{{end}}
`
