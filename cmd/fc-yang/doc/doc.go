package doc

import (
	"fmt"
	"strings"

	"github.com/freeconf/yang/meta"
)

type doc struct {
	LastErr error
	Title   string
	Def     *module
	Defs    []*def
	ModDefs []*module

	// Keep track of all meta to avoid repeating and handle recursive schemas
	History map[meta.Meta]*def
}

func (self *doc) werr(n int, err error) {
	if self.LastErr != nil {
		self.LastErr = err
	}
}

type module struct {
	Meta *meta.Module
}

type field struct {
	Meta    meta.Definition
	Case    *meta.ChoiceCase
	Def     *def
	Level   int
	Type    string
	Expand  []*field
	Details string
}

type action struct {
	Meta         *meta.Rpc
	Def          *def
	InputFields  []*field
	OutputFields []*field
}

type event struct {
	Def    *def
	Meta   *meta.Notification
	Fields []*field
}

type def struct {
	Parent  *def
	Meta    meta.HasDefinitions
	Fields  []*field
	Actions []*action
	Events  []*event
}

func (self *doc) build(m *meta.Module) error {
	if self.ModDefs == nil {
		self.ModDefs = make([]*module, 0)
	}
	self.History = make(map[meta.Meta]*def)
	docMod := &module{
		Meta: m,
		//LastPathSegment: m.GetIdent(),
	}
	self.Def = docMod
	self.ModDefs = append(self.ModDefs, docMod)
	if self.Defs == nil {
		self.Defs = make([]*def, 0, 128)
	}
	_, err := self.appendDef(m, nil, 0)
	return err
}

func (self *doc) appendDef(mdef meta.HasDefinitions, parent *def, level int) (*def, error) {
	d, isRepeat := self.History[mdef]
	if isRepeat {
		return d, nil
	}
	d = &def{
		Parent: parent,
		Meta:   mdef,
	}
	self.History[mdef] = d
	self.Defs = append(self.Defs, d)
	if x, ok := mdef.(meta.HasActions); ok {
		var err error
		d.Actions = make([]*action, 0, len(x.Actions()))
		for _, y := range x.Actions() {
			actionDef := &action{
				Meta: y,
				Def:  d,
			}
			d.Actions = append(d.Actions, actionDef)
			if y.Input() != nil {
				actionDef.InputFields, err = self.buildFields(y.Input())
				if err != nil {
					return nil, err
				}
			}
			if y.Output() != nil {
				actionDef.OutputFields, err = self.buildFields(y.Output())
				if err != nil {
					return nil, err
				}
			}
		}
	}
	if x, ok := mdef.(meta.HasNotifications); ok {
		var err error
		d.Events = make([]*event, 0, len(x.Notifications()))
		for _, y := range x.Notifications() {
			eventDef := &event{
				Meta: y,
				Def:  d,
			}
			d.Events = append(d.Events, eventDef)
			eventDef.Fields, err = self.buildFields(y)
			if err != nil {
				return nil, err
			}
		}
	}
	if x, ok := mdef.(meta.HasDataDefs); ok {
		d.Fields = make([]*field, 0, len(x.DataDefs()))
		for _, y := range x.DataDefs() {
			if choice, ok := y.(*meta.Choice); ok {
				for _, kase := range choice.DataDefs() {
					for _, kaseDef := range kase.(meta.HasDataDefs).DataDefs() {
						field, err := self.buildField(kaseDef)
						if err != nil {
							return nil, err
						}
						d.Fields = append(d.Fields, field)
						if !meta.IsLeaf(kaseDef) {
							// recurse
							childDef, err := self.appendDef(kaseDef.(meta.HasDefinitions), d, level+1)
							if err != nil {
								return nil, err
							}
							field.Def = childDef
						}
						field.Case = kase.(*meta.ChoiceCase)
						self.appendCaseDetails(field, choice, kase.(*meta.ChoiceCase))
					}
				}
			} else {
				field, err := self.buildField(y)
				if err != nil {
					return nil, err
				}
				d.Fields = append(d.Fields, field)
				if !meta.IsLeaf(y) {
					// recurse
					childDef, err := self.appendDef(y.(meta.HasDefinitions), d, level+1)
					if err != nil {
						return nil, err
					}
					field.Def = childDef
				}
			}
		}
	}

	return d, nil
}

func (self *doc) buildField(m meta.Definition) (*field, error) {
	f := &field{
		Meta: m,
	}
	if leafMeta, hasType := m.(meta.HasType); hasType {
		dt := leafMeta.Type()
		if meta.IsLeaf(m) {
			f.Type = dt.Ident()
			if dt.Format().IsList() {
				f.Type += "[]"
			}
		}
		var details []string
		if leafMeta.HasDefault() {
			details = append(details, fmt.Sprintf("Default: %v", leafMeta.Default()))
		}
		if len(dt.Enum()) > 0 {
			details = append(details, fmt.Sprintf("Allowed Values: %s", dt.Enum().String()))
		}
		if dets, valid := m.(meta.HasDetails); valid {
			if !dets.Config() {
				details = append(details, "r/o")
			}
			if dets.Mandatory() {
				details = append(details, "mandatory")
			}
		}
		if len(details) > 0 {
			f.Details = strings.Join(details, ", ")
		}
	}

	return f, nil
}

func (self *doc) appendCaseDetails(f *field, choice *meta.Choice, kase *meta.ChoiceCase) {
	details := fmt.Sprintf("choice: %s, case: %s", choice.Ident(), kase.Ident())
	if f.Details == "" {
		f.Details = details
	} else {
		f.Details = f.Details + ", " + details
	}
}

func (self *doc) buildFields(mlist meta.HasDataDefs) ([]*field, error) {
	fields := make([]*field, 0, len(mlist.DataDefs()))
	for _, ddef := range mlist.DataDefs() {
		field, err := self.buildField(ddef)
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)
		if !meta.IsLeaf(ddef) {
			self.appendExpandableFields(field, ddef.(meta.HasDataDefs), 0)
		}
	}
	return fields, nil
}

func (self *doc) appendExpandableFields(field *field, mlist meta.HasDataDefs, level int) error {
	for _, ddef := range mlist.DataDefs() {
		f, err := self.buildField(ddef)
		if err != nil {
			return err
		}
		f.Level = level + 1
		field.Expand = append(field.Expand, f)
		if !meta.IsLeaf(ddef) {
			self.appendExpandableFields(field, ddef.(meta.HasDataDefs), level+1)
		}
	}
	return nil
}
