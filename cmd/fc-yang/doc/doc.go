package doc

import (
	"fmt"

	"github.com/freeconf/yang/meta"
)

type doc struct {
	LastErr error
	Title   string
	Module  *def

	// Flattened list of all containers, lists, module that are for config and
	// metrics definitions in order
	DataDefs []*def

	// Flattened list of all actions
	Actions []*def

	// Flattened list of all events
	Events []*def

	// Keep track of all meta to avoid repeating and handle recursive schemas
	History map[meta.Definition]*def
}

func (self *doc) werr(n int, err error) {
	if self.LastErr != nil {
		self.LastErr = err
	}
}

// fields are public so templates can access them

// superset of all types of definitions
type def struct {
	Parent *def
	Meta   meta.Definition
	Level  int

	// leaf, leaf-list
	ScalarType string
	Details    string

	// container, list, module, choice, notification
	Fields []*def
	Case   *meta.ChoiceCase

	Events  []*def
	Actions []*def

	// rpc
	Input  *def
	Output *def
}

func (d *def) Type() string {
	// strip meta.
	return fmt.Sprintf("%T", d.Meta)[4:]
}

func (d *def) Leafable() bool {
	return meta.IsLeaf(d.Meta)
}

func (d *def) appendDetail(s string) {
	if d.Details == "" {
		d.Details = s
	} else {
		d.Details = d.Details + ", " + s
	}
}

// Expand returns every child, and their child and so on. Use level to show
// what def is under what def
func (d *def) Expand() []*def {
	var x []*def
	for _, child := range d.Fields {
		x = child.add(x)
	}
	return x
}

func (d *def) add(in []*def) []*def {
	out := append(in, d)
	for _, child := range d.Fields {
		out = child.add(out)
	}
	return out
}

func (self *doc) build(m *meta.Module) error {
	self.History = make(map[meta.Definition]*def)
	var err error
	self.Module, err = self.appendDef(nil, m, 0)
	self.DataDefs = []*def{self.Module}
	for _, d := range self.Module.Expand() {
		if !d.Leafable() {
			self.DataDefs = append(self.DataDefs, d)
		}
	}
	return err
}

func (self *doc) appendDefs(parent *def, mdefs []meta.Definition, level int) ([]*def, error) {
	defs := make([]*def, 0, len(mdefs))
	for _, y := range mdefs {
		if choice, ok := y.(*meta.Choice); ok {
			for _, kaseIdent := range choice.CaseIdents() {
				kase := choice.Cases()[kaseIdent]
				caseDefs, err := self.appendDefs(parent, kase.DataDefinitions(), level)
				if err != nil {
					return nil, err
				}
				for _, cdef := range caseDefs {
					cdef.Case = kase
					cdef.appendDetail(fmt.Sprintf("choice: %s, case: %s", choice.Ident(), kase.Ident()))
				}
				defs = append(defs, caseDefs...)
			}
		} else {
			d, err := self.appendDef(parent, y, level)
			if err != nil {
				return nil, err
			}
			defs = append(defs, d)
		}
	}
	return defs, nil
}

func (self *doc) appendDef(parent *def, m meta.Definition, level int) (*def, error) {
	d, isRepeat := self.History[m]
	if isRepeat {
		// handle recursive definitions
		return d, nil
	}
	d = &def{
		Parent: parent,
		Meta:   m,
		Level:  level,
	}
	self.History[m] = d
	if leafMeta, hasType := m.(meta.Leafable); hasType {
		dt := leafMeta.Type()
		if meta.IsLeaf(m) {
			d.ScalarType = dt.Ident()
			if dt.Format().IsList() {
				d.ScalarType += "[]"
			}
		}
		if leafMeta.HasDefault() {
			d.appendDetail(fmt.Sprintf("Default: %v", leafMeta.Default()))
		}
		if len(dt.Enum()) > 0 {
			d.appendDetail(fmt.Sprintf("Allowed Values: %s", dt.Enum().String()))
		}
		if dets, valid := m.(meta.HasDetails); valid {
			if !dets.Config() {
				d.appendDetail("r/o")
			}
			if dets.Mandatory() {
				d.appendDetail("mandatory")
			}
		}
	}
	if x, ok := m.(meta.HasActions); ok {
		for _, y := range x.Actions() {
			actionDef := &def{
				Meta:   y,
				Parent: d,
			}
			self.Actions = append(self.Actions, actionDef)
			d.Actions = append(d.Actions, actionDef)
			if y.Input() != nil {
				inputDef := &def{
					Meta:   y,
					Parent: d,
				}
				fields, err := self.appendDefs(inputDef, y.Input().DataDefinitions(), 0)
				if err != nil {
					return nil, err
				}
				inputDef.Fields = fields
				actionDef.Input = inputDef
			}
			if y.Output() != nil {
				outputDef := &def{
					Meta:   y,
					Parent: d,
				}
				fields, err := self.appendDefs(outputDef, y.Output().DataDefinitions(), 0)
				if err != nil {
					return nil, err
				}
				outputDef.Fields = fields
				actionDef.Output = outputDef
			}
		}
	}
	if x, ok := m.(meta.HasNotifications); ok {
		for _, y := range x.Notifications() {
			eventDef := &def{
				Meta:   y,
				Parent: d,
			}
			self.Events = append(self.Events, eventDef)
			d.Events = append(d.Events, eventDef)
			fields, err := self.appendDefs(eventDef, y.DataDefinitions(), 0)
			if err != nil {
				return nil, err
			}
			eventDef.Fields = fields
		}
	}
	if x, ok := m.(meta.HasDataDefinitions); ok {
		fields, err := self.appendDefs(d, x.DataDefinitions(), level+1)
		if err != nil {
			return nil, err
		}
		d.Fields = fields
	}

	return d, nil
}
