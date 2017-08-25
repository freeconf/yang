package render

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

func (self *Doc) Build(m *meta.Module) error {
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
	_, err := self.AppendDef(m, nil, 0)
	return err
}

func (self *Doc) AppendDef(mdef meta.MetaList, parent *DocDef, level int) (*DocDef, error) {
	def, isRepeat := self.History[mdef]
	if isRepeat {
		return def, nil
	}
	def = &DocDef{
		Parent: parent,
		Meta:   mdef,
	}
	self.History[mdef] = def
	self.Defs = append(self.Defs, def)
	i := meta.Children(mdef, true)
	for i.HasNext() {
		m, err := i.Next()
		if err != nil {
			return nil, err
		}
		if notif, isNotif := m.(*meta.Notification); isNotif {
			eventDef := &DocEvent{
				Meta: notif,
				Def:  def,
			}
			def.Events = append(def.Events, eventDef)
			eventDef.Fields, err = self.BuildFields(notif)
			if err != nil {
				return nil, err
			}
		} else if action, isAction := m.(*meta.Rpc); isAction {
			actionDef := &DocAction{
				Meta: action,
				Def:  def,
			}
			def.Actions = append(def.Actions, actionDef)
			if action.Input != nil {
				actionDef.InputFields, err = self.BuildFields(action.Input)
				if err != nil {
					return nil, err
				}
			}
			if action.Output != nil {
				actionDef.OutputFields, err = self.BuildFields(action.Output)
				if err != nil {
					return nil, err
				}
			}
		} else if choice, isChoice := m.(*meta.Choice); isChoice {
			p := choice.FirstMeta
			for p != nil {
				cse := p.(*meta.ChoiceCase)
				csei := meta.Children(cse, true)
				for csei.HasNext() {
					fm, err := csei.Next()
					if err != nil {
						return nil, err
					}
					field, err := self.BuildField(fm)
					if err != nil {
						return nil, err
					}
					def.Fields = append(def.Fields, field)

					if !meta.IsLeaf(fm) {
						// recurse
						childDef, err := self.AppendDef(fm.(meta.MetaList), def, level+1)
						if err != nil {
							return nil, err
						}
						field.Def = childDef
					}
					field.Case = cse
				}
				p = p.GetSibling()
			}
		} else {
			field, err := self.BuildField(m)
			if err != nil {
				return nil, err
			}
			def.Fields = append(def.Fields, field)
			if !meta.IsLeaf(m) {
				// recurse
				childDef, err := self.AppendDef(m.(meta.MetaList), def, level+1)
				if err != nil {
					return nil, err
				}
				field.Def = childDef
			}
		}
	}
	return def, nil
}

func (self *Doc) BuildField(m meta.Meta) (*DocField, error) {
	leafMeta, hasDataType := m.(meta.HasDataType)
	f := &DocField{
		Meta: m,
	}
	if hasDataType {
		info, err := leafMeta.GetDataType().Info()
		if err != nil {
			return nil, err
		}
		if meta.IsLeaf(m) {
			f.Type = leafMeta.GetDataType().Ident
			if info.Format.IsList() {
				f.Type += "[]"
			}
		}
		var details []string
		if info.HasDefault {
			details = append(details, fmt.Sprintf("Default: %s", info.Default))
		}
		e := info.Enum
		if len(e) > 0 {
			details = append(details, fmt.Sprintf("Allowed Values: %s", e.String()))
		}
		if len(details) > 0 {
			f.Details = strings.Join(details, ", ")
		}
	}
	return f, nil
}

func (self *Doc) BuildFields(mlist meta.MetaList) ([]*DocField, error) {
	var fields []*DocField
	i := meta.Children(mlist, true)
	for i.HasNext() {
		m, err := i.Next()
		if err != nil {
			return nil, err
		}
		field, err := self.BuildField(m)
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)
		if !meta.IsLeaf(m) {
			self.AppendExpandableFields(field, m.(meta.MetaList), 0)
		}
	}
	return fields, nil
}

func (self *Doc) AppendExpandableFields(field *DocField, mlist meta.MetaList, level int) error {
	i := meta.Children(mlist, true)
	for i.HasNext() {
		m, err := i.Next()
		if err != nil {
			return err
		}
		f, err := self.BuildField(m)
		if err != nil {
			return err
		}
		f.Level = level + 1
		field.Expand = append(field.Expand, f)
		if !meta.IsLeaf(m) {
			self.AppendExpandableFields(field, m.(meta.MetaList), level+1)
		}
	}
	return nil
}
