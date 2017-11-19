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
	Meta    meta.HasDefinitions
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

func (self *Doc) AppendDef(mdef meta.HasDefinitions, parent *DocDef, level int) (*DocDef, error) {
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
	if x, ok := mdef.(meta.HasActions); ok {
		var err error
		def.Actions = make([]*DocAction, 0, len(x.Actions()))
		for _, y := range x.Actions() {
			actionDef := &DocAction{
				Meta: y,
				Def:  def,
			}
			def.Actions = append(def.Actions, actionDef)
			if y.Input() != nil {
				actionDef.InputFields, err = self.BuildFields(y.Input())
				if err != nil {
					return nil, err
				}
			}
			if y.Output() != nil {
				actionDef.OutputFields, err = self.BuildFields(y.Output())
				if err != nil {
					return nil, err
				}
			}
		}
	}
	if x, ok := mdef.(meta.HasNotifications); ok {
		var err error
		def.Events = make([]*DocEvent, 0, len(x.Notifications()))
		for _, y := range x.Notifications() {
			eventDef := &DocEvent{
				Meta: y,
				Def:  def,
			}
			def.Events = append(def.Events, eventDef)
			eventDef.Fields, err = self.BuildFields(y)
			if err != nil {
				return nil, err
			}
		}
	}
	if x, ok := mdef.(meta.HasDataDefs); ok {
		def.Fields = make([]*DocField, 0, len(x.DataDefs()))
		for _, y := range x.DataDefs() {
			if choice, ok := y.(*meta.Choice); ok {
				for _, kase := range choice.DataDefs() {
					for _, kaseDef := range kase.(meta.HasDataDefs).DataDefs() {
						field, err := self.BuildField(kaseDef)
						if err != nil {
							return nil, err
						}
						def.Fields = append(def.Fields, field)
						if !meta.IsLeaf(kaseDef) {
							// recurse
							childDef, err := self.AppendDef(kaseDef.(meta.HasDefinitions), def, level+1)
							if err != nil {
								return nil, err
							}
							field.Def = childDef
						}
						field.Case = kase.(*meta.ChoiceCase)
					}
				}
			} else {
				field, err := self.BuildField(y)
				if err != nil {
					return nil, err
				}
				def.Fields = append(def.Fields, field)
				if !meta.IsLeaf(y) {
					// recurse
					childDef, err := self.AppendDef(y.(meta.HasDefinitions), def, level+1)
					if err != nil {
						return nil, err
					}
					field.Def = childDef
				}
			}
		}
	}

	return def, nil
}

func (self *Doc) BuildField(m meta.Meta) (*DocField, error) {
	f := &DocField{
		Meta: m,
	}
	if leafMeta, hasDataType := m.(meta.HasDataType); hasDataType {
		dt := leafMeta.DataType()
		if meta.IsLeaf(m) {
			f.Type = dt.TypeIdent()
			if dt.Format().IsList() {
				f.Type += "[]"
			}
		}
		var details []string
		if leafMeta.HasDefault() {
			details = append(details, fmt.Sprintf("Default: %v", leafMeta.Default()))
		}
		e := dt.Enum()
		if len(e) > 0 {
			details = append(details, fmt.Sprintf("Allowed Values: %s", e.String()))
		}
		if len(details) > 0 {
			f.Details = strings.Join(details, ", ")
		}
	}

	return f, nil
}

func (self *Doc) BuildFields(mlist meta.HasDataDefs) ([]*DocField, error) {
	fields := make([]*DocField, 0, len(mlist.DataDefs()))
	for _, ddef := range mlist.DataDefs() {
		field, err := self.BuildField(ddef)
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)
		if !meta.IsLeaf(ddef) {
			self.AppendExpandableFields(field, ddef.(meta.HasDataDefs), 0)
		}
	}
	return fields, nil
}

func (self *Doc) AppendExpandableFields(field *DocField, mlist meta.HasDataDefs, level int) error {
	for _, ddef := range mlist.DataDefs() {
		f, err := self.BuildField(ddef)
		if err != nil {
			return err
		}
		f.Level = level + 1
		field.Expand = append(field.Expand, f)
		if !meta.IsLeaf(ddef) {
			self.AppendExpandableFields(field, ddef.(meta.HasDataDefs), level+1)
		}
	}
	return nil
}
