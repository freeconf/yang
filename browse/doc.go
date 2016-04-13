package browse

import (
	"io"
	"github.com/c2g/meta"
	"html/template"
	"fmt"
	"strings"
)

type Doc struct {
	LastErr error
	Title string
	Defs []*DocDef
	ModDefs []*DocModule
}

func (self *Doc) werr(n int, err error) {
	if self.LastErr != nil {
		self.LastErr = err
	}
}

type DocModule struct {
	Meta *meta.Module
}

type DocField struct {
	Meta meta.Meta
	Link string
	Title string
	IndentPx int
	Type string
	Expand []*DocField
	Details string
}

type DocAction struct {
	Meta *meta.Rpc
	Title string
	InputFields []*DocField
	OutputFields []*DocField
}

type DocEvent struct {
	Meta *meta.Notification
	Title string
	Fields []*DocField
}

type DocDef struct {
	Anchor string
	Descriptions []string
	ParentPath string
	LastPathSegment string
	Meta meta.MetaList
	Fields []*DocField
	Actions []*DocAction
	Events []*meta.Notification
}

func (self *Doc) Build(m *meta.Module) {
	if self.ModDefs == nil {
		self.ModDefs = make([]*DocModule, 0)
	}
	docMod := &DocModule{
		Meta : m,
	}
	self.ModDefs = append(self.ModDefs, docMod)
	if self.Defs == nil {
		self.Defs = make([]*DocDef, 0, 128)
	}
	self.AppendDef(m, "", 0)
}

func (self *Doc) Generate(out io.Writer) error {
	t := template.Must(template.New("c2doc").Parse(docHtml))
	err := t.Execute(out, self)
	return err
}

func (self *Doc) AppendDef(mdef meta.MetaList, parentPath string, level int) *DocDef {
	def := &DocDef{
		ParentPath : parentPath,
		Meta: mdef,
		Anchor: parentPath + "/" + mdef.GetIdent(),
	}
	var path string
	if _, isModule := mdef.(*meta.Module); !isModule {
		self.Defs = append(self.Defs, def)
		def.LastPathSegment = mdef.GetIdent()
		path = parentPath + "/" + def.LastPathSegment
		if mlist, isList := mdef.(*meta.List); isList {
			path = path + fmt.Sprintf("={%v}", strings.Join(mlist.Key, ","))
		}
		def.Descriptions = []string{mdef.(meta.Describable).GetDescription()}
	} else {
		description := fmt.Sprintf("%s - %s", mdef.GetIdent(),  mdef.(meta.Describable).GetDescription())
		if len(self.Defs) == 0 {
			self.Defs = append(self.Defs, def)
			def.Descriptions = []string{description}
		} else {
			// effectively merge module defs
			def = self.Defs[0]
			def.Descriptions = append(def.Descriptions, description)
		}
	}
	i := meta.NewMetaListIterator(mdef, true)
	for i.HasNextMeta() {
		m := i.NextMeta()
		if notif, isNotif := m.(*meta.Notification); isNotif {
			eventDef := &DocEvent{
				Meta: notif,
				Title: notif.Ident,
			}
			def.Events = append(def.Events, notif)
			eventDef.Fields = self.BuildFields(notif)
		} else if action, isAction := m.(*meta.Rpc); isAction {
			actionDef := &DocAction{
				Meta: action,
				Title: action.Ident,
			}
			def.Actions = append(def.Actions, actionDef)
			if action.Input != nil {
				actionDef.InputFields = self.BuildFields(action.Input)
			}
			if action.Output != nil {
				actionDef.OutputFields = self.BuildFields(action.Output)
			}
		} else {
			field := self.BuildField(m)
			def.Fields = append(def.Fields, field)
			if ! meta.IsLeaf(m) {
				childDef := self.AppendDef(m.(meta.MetaList), path, level + 1)
				field.Link = "#" + childDef.Anchor
			}
		}
	}
	return def
}

func (self *Doc) BuildField(m meta.Meta) *DocField {
	title := m.GetIdent()
	var fieldType string
	if ! meta.IsLeaf(m) {
		if meta.IsList(m) {
			title += "[\u2026]"
		} else {
			title += "{\u2026}"
		}
	} else {
		leafMeta := m.(meta.HasDataType)
		fieldType = leafMeta.GetDataType().Ident
	}
	f := &DocField{
		Title : title,
		Meta: m,
		Type: fieldType,
	}
	if mType, hasDataType := m.(meta.HasDataType); hasDataType {
		var details []string
		d := mType.GetDataType().Default()
		if len(d) > 0 {
			details = append(details, fmt.Sprintf("Default: %s", d))
		}
		e := mType.GetDataType().Enumeration()
		if len(e) > 0 {
			details = append(details, fmt.Sprintf("Allowed Values: (%s)", strings.Join(e, ",")))
		}
		if len(details) > 0 {
			f.Details = strings.Join(details, ", ")
		}
	}
	return f
}

func (self *Doc) BuildFields(mlist meta.MetaList) (fields []*DocField) {
	i := meta.NewMetaListIterator(mlist, true)
	for i.HasNextMeta() {
		m := i.NextMeta()
		field := self.BuildField(m)
		fields = append(fields, field)
		if ! meta.IsLeaf(m) {
			self.AppendExpandableFields(field, m.(meta.MetaList), 0)
		}
	}
	return
}


func (self *Doc) AppendExpandableFields(field *DocField, mlist meta.MetaList, level int) {
	i := meta.NewMetaListIterator(mlist, true)
	for i.HasNextMeta() {
		m := i.NextMeta()
		f := self.BuildField(m)
		f.IndentPx = 10 + (level * 10)
		field.Expand = append(field.Expand, f)
		if ! meta.IsLeaf(m) {
			self.AppendExpandableFields(field, m.(meta.MetaList), level + 1)
		}
	}
}

// Copyright disclaimer : Much of CSS and a portion of the HTML was adapted from Golang's godoc generated
// pages under the BSD License (implied 3-Clause)
//    Copyright (c) 2012 The Go Authors. All rights reserved.

// Known issue, summary/details does not work in IE, Edge, and Firefox, but degrades ok.
var docHtml = `
<!DOCTYPE html>
<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#375EAB">
  <title>{{.Title}}</title>
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
        background-color: #13b5ea;
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
<details>
  <summary>Index</summary>
<ul>
{{range .Defs}}
<li><a href="#{{.Anchor}}">{{.ParentPath}}/{{.LastPathSegment}}</a></li>
{{end}}
</ul>
</details>
<!-- BEGIN LOOP -->
{{range .Defs}}
	<h2><a name="{{.Anchor}}"></a>{{.ParentPath}}/<span class="metalist">{{.LastPathSegment}}</span></h2>
	{{range .Descriptions}}
	<p>{{.}}</p>
	{{end}}
{{range .Fields}}
        {{if .Link}}
	<code><b><a href="{{.Link}}">{{.Title}}</b></a> - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></code>
        {{else}}
	<code><b>{{.Title}}</b> {{.Type}} - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></code>
	{{end}}
{{end}}
{{if .Actions}}
<h3>Actions</h3>
{{range .Actions}}
<h4>{{.Title}}</h4>
        <p>{{.Meta.Description}}</p>
        {{if .InputFields}}
		{{range .InputFields}}
			{{if .Expand}}
	                    <code class="expandContainer"><details class="expandable"
	                    ><summary><b>{{.Title}}</b> {{.Type}} - {{.Meta.Description}}</summary>
	                    {{- range .Expand -}}
			       <div style="margin: 2px 0 0 {{.IndentPx}}px;"><b>{{.Title}}</b> {{.Type}} - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></div>
			    {{- end -}}</details></code>
			{{else}}
			    <code><b>{{.Title}}</b> {{.Type}} - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></code>
			{{end}}
		{{end}}
	{{end}}

        {{if .OutputFields}}
<h5>response</h5>
		{{range .OutputFields}}
			{{if .Expand}}
	                    <code class="expandContainer"><details class="expandable"
	                    ><summary><b>{{.Title}}</b> {{.Type}} - {{.Meta.Description}}</summary>
	                    {{- range .Expand -}}
			       <div style="margin: 2px 0 0 {{.IndentPx}}px;"><b>{{.Title}}</b> {{.Type}} - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></div>
			    {{- end -}}</details></code>
			{{else}}
			    <code><b>{{.Title}}</b> {{.Type}} - {{.Meta.Description}} <span class="fieldDetails">{{.Details}}</span></code>
			{{end}}
		{{end}}
	{{end}}
{{end}}
{{end}}
{{if .Events}}
<h3>Events</h3>
{{range .Events}}
<h4>{{.Ident}}</h4>
        <p>{{.Description}}</p>
	<code><a href=#><b>time</b> decimal64</a> - time to complete</code>
{{end}}
{{end}}

{{end}}

</div>
</div>
<!-- Generated using c2-doc -->
</body></html>

`