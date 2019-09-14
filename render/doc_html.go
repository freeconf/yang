package render

import (
	"html/template"
	"io"
	"strings"
)

type DocHtml struct {
	ImageLink string
}

func (self *DocHtml) Generate(doc *Doc, tmpl string, out io.Writer) error {
	funcMap := template.FuncMap{
		"repeat":   strings.Repeat,
		"link":     docLink,
		"title":    docTitle,
		"title2":   docTitle2,
		"type":     docFieldType,
		"path":     docPath,
		"indentPx": htmlIndentPx,
		"noescape": func(s string) template.HTML {
			return template.HTML(s)
		},
	}
	t := template.Must(template.New("c2doc").Funcs(funcMap).Parse(tmpl))
	err := t.Execute(out, struct {
		Doc       *Doc
		ImageLink string
	}{
		Doc:       doc,
		ImageLink: self.ImageLink,
	})
	return err
}

func (self *DocHtml) BuiltinTemplate() string {
	return docHtml
}

func htmlIndentPx(level int) int {
	return 10 + (level * 10)
}

// Copyright disclaimer : Much of CSS and a portion of the HTML was adapted from Golang's godoc generated
// pages under the BSD License (implied 3-Clause)
//    Copyright (c) 2012 The Go Authors. All rights reserved.
//
// Known issue, summary/details does not work in IE, Edge, and Firefox, but degrades ok.
const docHtml = `
<!DOCTYPE html>
<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#375EAB">
  <title>{{.Doc.Title}}</title>
  <style>
body {
	margin: 0;
	font-family: Arial, sans-serif;
	font-size: 16px;
	background-color: #fff;
	line-height: 1.3em;
	color: #222;
	text-align: center;
}

code {
	display: block;
	font-family: Menlo, monospace;
	font-size: 14px;
	line-height: 1.4em;
	overflow-x: auto;
	background: #EFEFEF;
	padding: 10px;

	-webkit-border-radius: 5px;
	-moz-border-radius: 5px;
	border-radius: 5px;
}

a {
	color: #375EAB;
	text-decoration: none;
}

a:hover {
	text-decoration: underline;
}
p, li {
	max-width: 800px;
	word-wrap: break-word;
	list-style: none;
}

li:before {
    content: "•";
    font-size: 28px;
    padding-right: 7px;
}

li.def:before {
    color:#13b5ea;
}

li.action:before {
    color:#b64ff7;
}

li.notification:before {
    color:#4fb32e;
}

p, code {
	margin: 20px;
}

h1,
h2,
h3,
h4 {
	margin: 20px 0 20px;
	padding: 0;
	font-weight: bold;
}
h1 {
	font-size: 28px;
	line-height: 1;
}
h2 {
	font-size: 20px;
	color: white;
	padding: 8px;
	line-height: 1.25;
	font-weight: normal;
	clear: right;
}
h3 {
	font-size: 20px;
}
h3,
h4 {
	margin: 20px 5px;
}
h4 {
	font-size: 16px;
}
h5 {
    margin-left: 30px;
}

h2.def {
        background-color: #13b5ea;
}

h2.action::after {
	content: " (Action) ";
	font-style: italic;
	font-size: smaller;
}

h2.notification::after {
	content: " (Notification) ";
	font-style: italic;
	font-size: smaller;
}

h2.action {
	background-color: #b64ff7;
}

h2.notification {
	background-color: #4fb32e;
}

div#page {
	width: 100%;
}

div#page > .container {
	text-align: left;
	margin-left: auto;
	margin-right: auto;
	padding: 0 20px;
	max-width: 950px;
}

.metalist {
    font-size: larger;
    font-weight: bolder;
}

.output {
    background-color: beige;
}

div#page.wide > .container {
	max-width: none;
}

.fieldDetails {
    font-size: smaller;
    font-style: italic;
}

hr {
    border-style: none;
    border-top: 1px solid black;
}

@media (max-width: 930px) {
	#heading-wide {
		display: none;
	}
	#heading-narrow {
		display: block;
	}
}


@media (max-width: 760px) {
	.container .left,
	.container .right {
		width: auto;
		float: none;
	}
}

@media (max-width: 700px) {
	body {
		font-size: 15px;
	}
	code {
		font-size: 13px;
	}

	div#page > .container {
		padding: 0 10px;
	}

	#heading-wide {
		display: block;
	}
	#heading-narrow {
		display: none;
	}

	p,
	code,
	ul,
	ol {
		margin: 10px;
	}
}

@media (max-width: 480px) {
	#heading-wide {
		display: none;
	}
	#heading-narrow {
		display: block;
	}
}

@media print {
	code {
		background: #FFF;
		border: 1px solid #BBB;
		white-space: pre-wrap;
	}
}
  </style>
</head>
<body>

<div id="page" class="wide" tabindex="-1" style="outline: 0px;">
<div class="container">
<h1>{{.Doc.Title}}</h1>
<img src="{{index .ImageLink}}" onerror="this.style.display='none'">
<details>
  <summary>Index</summary>
<ul>
{{range .Doc.Defs}}
<li class="def"><a href="#{{link .}}">{{path .}}</a></li>

{{range .Actions}}
<li class="action"><a href="#{{link .}}">{{path .Def}}{{title2 .Meta}}</a></li>
{{end}}

{{range .Events}}
<li class="notification"><a href="#{{link .}}">{{path .Def}}{{title2 .Meta}}</a></li>
{{end}}

{{end}}
</ul>
</details>
<!-- BEGIN LOOP -->
{{range .Doc.Defs}}
	<h2 class="def"><a name="{{link .}}"></a>{{path .Parent}}<span class="metalist">{{title2 .Meta}}</span></h2>
	<p>{{.Meta.Description}}</p>
{{range .Fields}}
        {{if .Def}}
	<code><strong><a href="#{{link .Def}}">{{title .Meta}}</strong></a> - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></code>
        {{else}}
	<code><strong>{{title .Meta}}</strong> {{type .}} - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></code>
	{{end}}
{{end}}
{{if .Actions}}
{{range .Actions}}
	<h2 class="action"><a name="{{link .}}"></a>{{path .Def}}<span class="metalist">{{title2 .Meta}}</span></h2>
	<p>{{.Meta.Description}}</p>
        {{if .InputFields}}
		{{range .InputFields}}
			{{if .Expand}}
	                    <code class="expandContainer"><details class="expandable"
	                    ><summary><strong>{{title .Meta}}</strong> {{type .}} - {{.Meta.Description}}</summary>
	                    {{- range .Expand -}}
			       <div style="margin: 2px 0 0 {{indentPx .Level}}px;"><strong>{{title .Meta}}</strong> {{type .}} - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></div>
			    {{- end -}}</details></code>
			{{else}}
			    <code><strong>{{title .Meta}}</strong> {{type .}} - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></code>
			{{end}}
		{{end}}
	{{end}}

        {{if .OutputFields}}
<h5>response</h5>
		{{range .OutputFields}}
			{{if .Expand}}
	                    <code class="expandContainer"><details class="expandable"
	                    ><summary><strong>{{title .Meta}}</strong> {{type .}} - {{.Meta.Description}}</summary>
	                    {{- range .Expand -}}
			       <div style="margin: 2px 0 0 {{indentPx .Level}}px;"><strong>{{title .Meta}}</strong> {{type .}} - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></div>
			    {{- end -}}</details></code>
			{{else}}
			    <code><strong>{{title .Meta}}</strong> {{type .}} - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></code>
			{{end}}
		{{end}}
	{{end}}
{{end}}
{{end}}
{{if .Events}}
{{range .Events}}
	<h2 class="notification"><a name="{{link .}}"></a>{{path .Def}}<span class="metalist">{{title2 .Meta}}</span></h2>
	<p>{{.Meta.Description}}</p>

        {{if .Fields}}
		{{range .Fields}}
			{{if .Expand}}
	                    <code class="expandContainer"><details class="expandable"
	                    ><summary><strong>{{title .Meta}}</strong> {{type .}} - {{.Meta.Description}}</summary>
	                    {{- range .Expand -}}
			       <div style="margin: 2px 0 0 {{indentPx .Level}}px;"><strong>{{title .Meta}}</strong> {{type .}} - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></div>
			    {{- end -}}</details></code>
			{{else}}
			    <code><strong>{{title .Meta}}</strong> {{type .}} - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></code>
			{{end}}
		{{end}}
	{{end}}
	
{{end}}
{{end}}

{{end}}

</div>
</div>
<!-- Generated using c2-doc -->
</body></html>
`
