package browse

import (
	"fmt"
	"io"
	"strings"

	"github.com/c2stack/c2g/meta"
)

type Doc struct {
	LastErr error
	Title   string
	Defs    []*DocDef
	ModDefs []*DocModule

	// Keep track of all meta to avoid repeating and handle recursive schemas
	History map[meta.Meta]*DocDef
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
	Meta    meta.Meta
	Case    *meta.ChoiceCase
	Def     *DocDef
	Level   int
	Type    string
	Expand  []*DocField
	Details string
}

type DocAction struct {
	Meta         *meta.Rpc
	Def          *DocDef
	InputFields  []*DocField
	OutputFields []*DocField
}

type DocEvent struct {
	Def    *DocDef
	Meta   *meta.Notification
	Fields []*DocField
}

type DocDef struct {
	Parent  *DocDef
	Meta    meta.MetaList
	Fields  []*DocField
	Actions []*DocAction
	Events  []*DocEvent
}

type DocDefBuilder interface {
	Generate(doc *Doc, out io.Writer) error
}

func (self *Doc) Build(m *meta.Module) {
	if self.ModDefs == nil {
		self.ModDefs = make([]*DocModule, 0)
	}
	self.History = make(map[meta.Meta]*DocDef)
	docMod := &DocModule{
		Meta: m,
		//LastPathSegment: m.GetIdent(),
	}
	self.ModDefs = append(self.ModDefs, docMod)
	if self.Defs == nil {
		self.Defs = make([]*DocDef, 0, 128)
	}
	self.AppendDef(m, nil, 0)
}

func (self *Doc) AppendDef(mdef meta.MetaList, parent *DocDef, level int) *DocDef {
	def, isRepeat := self.History[mdef]
	if isRepeat {
		return def
	}
	def = &DocDef{
		Parent: parent,
		Meta:   mdef,
	}
	self.History[mdef] = def
	self.Defs = append(self.Defs, def)
	i := meta.NewMetaListIterator(mdef, true)
	for i.HasNextMeta() {
		m := i.NextMeta()
		if notif, isNotif := m.(*meta.Notification); isNotif {
			eventDef := &DocEvent{
				Meta: notif,
				Def:  def,
			}
			def.Events = append(def.Events, eventDef)
			eventDef.Fields = self.BuildFields(notif)
		} else if action, isAction := m.(*meta.Rpc); isAction {
			actionDef := &DocAction{
				Meta: action,
				Def:  def,
			}
			def.Actions = append(def.Actions, actionDef)
			if action.Input != nil {
				actionDef.InputFields = self.BuildFields(action.Input)
			}
			if action.Output != nil {
				actionDef.OutputFields = self.BuildFields(action.Output)
			}
		} else if choice, isChoice := m.(*meta.Choice); isChoice {
			p := choice.FirstMeta
			for p != nil {
				cse := p.(*meta.ChoiceCase)
				csei := meta.NewMetaListIterator(cse, true)
				for csei.HasNextMeta() {
					fm := csei.NextMeta()
					field := self.BuildField(fm)
					def.Fields = append(def.Fields, field)

					if !meta.IsLeaf(fm) {
						// recurse
						childDef := self.AppendDef(fm.(meta.MetaList), def, level+1)
						field.Def = childDef
					}
					field.Case = cse
				}
				p = p.GetSibling()
			}
		} else {
			field := self.BuildField(m)
			def.Fields = append(def.Fields, field)
			if !meta.IsLeaf(m) {
				// recurse
				childDef := self.AppendDef(m.(meta.MetaList), def, level+1)
				field.Def = childDef
			}
		}
	}
	return def
}

func (self *Doc) BuildField(m meta.Meta) *DocField {
	f := &DocField{
		Meta: m,
	}
	if meta.IsLeaf(m) {
		leafMeta := m.(meta.HasDataType)
		f.Type = leafMeta.GetDataType().Ident
		if meta.IsListFormat(leafMeta.GetDataType().Format()) {
			f.Type += "[]"
		}
	}
	if mType, hasDataType := m.(meta.HasDataType); hasDataType {
		var details []string
		d := mType.GetDataType().Default()
		if len(d) > 0 {
			details = append(details, fmt.Sprintf("Default: %s", d))
		}
		e := mType.GetDataType().Enumeration()
		if len(e) > 0 {
			details = append(details, fmt.Sprintf("Allowed Values: %s", e.String()))
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
		if !meta.IsLeaf(m) {
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
		f.Level = level + 1
		field.Expand = append(field.Expand, f)
		if !meta.IsLeaf(m) {
			self.AppendExpandableFields(field, m.(meta.MetaList), level+1)
		}
	}
}
