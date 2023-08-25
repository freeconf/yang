package meta

import (
	"fmt"
)

// This is boilerplate functions generated from ./meta/gen/ package. Do not edit
// this file, instead edit ./gen/gen.in and run "cd gen && go generate"
// Ident is identity of Module
func (m *Module) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Module) Parent() Meta {
	return m.parent
}

// Description of Module
func (m *Module) Description() string {
	return m.desc
}

func (m *Module) setDescription(desc string) {
	m.desc = desc
}

func (m *Module) Reference() string {
	return m.ref
}

func (m *Module) setReference(ref string) {
	m.ref = ref
}

func (m *Module) Extensions() []*Extension {
	return m.extensions
}

func (m *Module) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Module) DataDefinitions() []Definition {
	return m.dataDefs
}

func (m *Module) DataDefinition(ident string) Definition {
	return m.dataDefsIndex[ident]
}

func (m *Module) addDataDefinition(d Definition) error {
	if c, isChoice := d.(*Choice); isChoice {
		for _, k := range c.Cases() {
			for _, kdef := range k.DataDefinitions() {
				// recurse in case it's another choice
				if err := m.indexDataDefinition(kdef); err != nil {
					return err
				}
			}
		}
 	}
	
	if err := m.indexDataDefinition(d); err != nil {
		return err
	}
	m.dataDefs = append(m.dataDefs, d)
	return nil
}

func (m *Module) indexDataDefinition(def Definition) error {
	if m.dataDefsIndex == nil {
		m.dataDefsIndex = make(map[string]Definition)
	} else if _, exists := m.dataDefsIndex[def.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", def.Ident(), m.Ident())
	}	
	m.dataDefsIndex[def.Ident()] = def
	return nil
}

func (m *Module) popDataDefinitions() []Definition {
	orig := m.dataDefs
	m.dataDefs = make([]Definition, 0, len(orig))
	for key := range m.dataDefsIndex {
		delete(m.dataDefsIndex, key)
	}
	return orig
}
func (m *Module) IsRecursive() bool {
	return false
}

func (m *Module) markRecursive() {
	panic("Cannot mark Module recursive")
}



func (m *Module) Augments() []*Augment {
	return m.augments
}

func (m *Module) addAugments(a *Augment) {
	a.parent = m
	m.augments = append(m.augments, a)
}

func (m *Module) Groupings() map[string]*Grouping {
	return m.groupings
}

func (m *Module) addGrouping(g *Grouping) error {
	g.parent = m
	if m.groupings == nil {
		m.groupings = make(map[string]*Grouping)
	} else if _, exists := m.groupings[g.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", g.Ident(), m.Ident())
	}
    m.groupings[g.Ident()] = g
	return nil
}

func (m *Module) Typedefs() map[string]*Typedef {
	return m.typedefs
}

func (m *Module) addTypedef(t *Typedef) error {
	t.parent = m
	if m.typedefs == nil {
		m.typedefs = make(map[string]*Typedef)
	} else if _, exists := m.typedefs[t.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", t.Ident(), m.Ident())
	}
    m.typedefs[t.Ident()] = t
	return nil
}

func (m *Module) Actions() map[string]*Rpc {
	return m.actions
}

func (m *Module) addAction(a *Rpc) error {
	a.parent = m
	if m.actions == nil {
		m.actions = make(map[string]*Rpc)
	} else if _, exists := m.actions[a.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", a.Ident(), m.Ident())
	}	
    m.actions[a.Ident()] = a
	return nil
}

func (m *Module) setActions(actions map[string]*Rpc) {
	m.actions = actions
}

func (m *Module) Notifications() map[string]*Notification {
	return m.notifications
}

func (m *Module) addNotification(n *Notification) error {
	n.parent = m
	if m.notifications == nil {
		m.notifications = make(map[string]*Notification)
	} else if _, exists := m.notifications[n.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", n.Ident(), m.Ident())
	}
    m.notifications[n.Ident()] = n
	return nil
}

func (m *Module) setNotifications(notifications map[string]*Notification) {
	m.notifications = notifications
}


// Definition can be a data defintion, action or notification
func (m *Module) Definition(ident string) Definition {
	if x, found := m.notifications[ident]; found {
		return x
	}
	
	if x, found := m.actions[ident]; found {
		return x
	}
	
	if x, found := m.dataDefsIndex[ident]; found {
		return x
	}
	
	return nil
}

func (m *Module) Config() bool {
	return *m.configPtr
}

func (m *Module) setConfig(c bool) {
	m.configPtr = &c
}

func (m *Module) IsConfigSet() bool {
	return m.configPtr != nil
}

func (m *Module) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.notifications != nil {
		copy.notifications = make(map[string]*Notification, len(m.notifications))
		for ident, notif := range m.notifications {
			copy.notifications[ident] = notif.clone(&copy).(*Notification)
		}
	}
	
	if m.actions != nil {
		copy.actions = make(map[string]*Rpc, len(m.actions))
		for ident, action := range m.actions {
			copy.actions[ident] = action.clone(&copy).(*Rpc)
		}
	}
	
	if m.dataDefs != nil {
		copy.dataDefs = make([]Definition, len(m.dataDefs))
		copy.dataDefsIndex = make(map[string]Definition, len(m.dataDefs))
		for i, def := range m.dataDefs {
			copyDef := def.(cloneable).clone(&copy).(Definition)
			copy.dataDefs[i] = copyDef
			copy.dataDefsIndex[def.Ident()] = copyDef
		}
	}
	

	return &copy
}





// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Import) Parent() Meta {
	return m.parent
}

// Description of Import
func (m *Import) Description() string {
	return m.desc
}

func (m *Import) setDescription(desc string) {
	m.desc = desc
}

func (m *Import) Reference() string {
	return m.ref
}

func (m *Import) setReference(ref string) {
	m.ref = ref
}

func (m *Import) Extensions() []*Extension {
	return m.extensions
}

func (m *Import) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}




// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Include) Parent() Meta {
	return m.parent
}

// Description of Include
func (m *Include) Description() string {
	return m.desc
}

func (m *Include) setDescription(desc string) {
	m.desc = desc
}

func (m *Include) Reference() string {
	return m.ref
}

func (m *Include) setReference(ref string) {
	m.ref = ref
}

func (m *Include) Extensions() []*Extension {
	return m.extensions
}

func (m *Include) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}




// Ident is identity of Choice
func (m *Choice) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Choice) Parent() Meta {
	return m.parent
}

// Description of Choice
func (m *Choice) Description() string {
	return m.desc
}

func (m *Choice) setDescription(desc string) {
	m.desc = desc
}

func (m *Choice) Reference() string {
	return m.ref
}

func (m *Choice) setReference(ref string) {
	m.ref = ref
}

func (m *Choice) Status() Status {
	return m.status
}

func (m *Choice) setStatus(status Status) {
	m.status = status
}

func (m *Choice) Extensions() []*Extension {
	return m.extensions
}

func (m *Choice) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Choice) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *Choice) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

func (m *Choice) When() *When {
	return m.when
}

func (m *Choice) setWhen(w *When) {
	w.parent = m
    m.when = w
}

func (m *Choice) Config() bool {
	return *m.configPtr
}

func (m *Choice) setConfig(c bool) {
	m.configPtr = &c
}

func (m *Choice) IsConfigSet() bool {
	return m.configPtr != nil
}

func (m *Choice) Mandatory() bool {
	return m.mandatoryPtr != nil && *m.mandatoryPtr
}

func (m *Choice) setMandatory(b bool) {
	m.mandatoryPtr = &b
}

func (m *Choice) IsMandatorySet() bool {
	return m.mandatoryPtr != nil
}

func (m *Choice) Default() string {
	if m.defaultVal == nil {
		return ""
	}
	return *m.defaultVal
}

func (m *Choice) HasDefault() bool {
	return m.defaultVal != nil
}

func (m *Choice) addDefault(d string) {
	if m.defaultVal != nil {
		panic("default already set")
	}
	m.defaultVal = &d
}

func (m *Choice) DefaultValue() interface{} {
	return m.Default()
}

func (m *Choice) setDefaultValue(d interface{}) {
	if s, valid := d.(string); valid {
		m.addDefault(s)
	} else {
		panic("expected string")
	}
}


func (m *Choice) setDefault(d string) {
	m.defaultVal = &d
}

func (m *Choice) clearDefault() {
    m.defaultVal = nil
}

func (m *Choice) getOriginalParent() Definition {
	return m.originalParent
}

func (m *Choice) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.cases != nil {
		copy.cases = make(map[string]*ChoiceCase, len(m.cases))
		for ident, kase := range m.cases {
			copy.cases[ident] = kase.clone(&copy).(*ChoiceCase)
		}
	}
	

	return &copy
}



// Ident is identity of ChoiceCase
func (m *ChoiceCase) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *ChoiceCase) Parent() Meta {
	return m.parent
}

// Description of ChoiceCase
func (m *ChoiceCase) Description() string {
	return m.desc
}

func (m *ChoiceCase) setDescription(desc string) {
	m.desc = desc
}

func (m *ChoiceCase) Reference() string {
	return m.ref
}

func (m *ChoiceCase) setReference(ref string) {
	m.ref = ref
}

func (m *ChoiceCase) Status() Status {
	return m.status
}

func (m *ChoiceCase) setStatus(status Status) {
	m.status = status
}

func (m *ChoiceCase) Extensions() []*Extension {
	return m.extensions
}

func (m *ChoiceCase) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *ChoiceCase) DataDefinitions() []Definition {
	return m.dataDefs
}

func (m *ChoiceCase) DataDefinition(ident string) Definition {
	return m.dataDefsIndex[ident]
}

func (m *ChoiceCase) addDataDefinition(d Definition) error {
	if c, isChoice := d.(*Choice); isChoice {
		for _, k := range c.Cases() {
			for _, kdef := range k.DataDefinitions() {
				// recurse in case it's another choice
				if err := m.indexDataDefinition(kdef); err != nil {
					return err
				}
			}
		}
 	}
	
	if err := m.indexDataDefinition(d); err != nil {
		return err
	}
	m.dataDefs = append(m.dataDefs, d)
	return nil
}

func (m *ChoiceCase) indexDataDefinition(def Definition) error {
	if m.dataDefsIndex == nil {
		m.dataDefsIndex = make(map[string]Definition)
	} else if _, exists := m.dataDefsIndex[def.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", def.Ident(), m.Ident())
	}	
	m.dataDefsIndex[def.Ident()] = def
	return nil
}

func (m *ChoiceCase) popDataDefinitions() []Definition {
	orig := m.dataDefs
	m.dataDefs = make([]Definition, 0, len(orig))
	for key := range m.dataDefsIndex {
		delete(m.dataDefsIndex, key)
	}
	return orig
}
func (m *ChoiceCase) IsRecursive() bool {
	return m.recursive
}

func (m *ChoiceCase) markRecursive() {
	m.recursive = true
}


func (m *ChoiceCase) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *ChoiceCase) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

func (m *ChoiceCase) When() *When {
	return m.when
}

func (m *ChoiceCase) setWhen(w *When) {
	w.parent = m
    m.when = w
}

// Definition can be a data defintion, action or notification
func (m *ChoiceCase) Definition(ident string) Definition {
	if x, found := m.dataDefsIndex[ident]; found {
		return x
	}
	
	return nil
}

func (m *ChoiceCase) Config() bool {
	return *m.configPtr
}

func (m *ChoiceCase) setConfig(c bool) {
	m.configPtr = &c
}

func (m *ChoiceCase) IsConfigSet() bool {
	return m.configPtr != nil
}

func (m *ChoiceCase) getOriginalParent() Definition {
	return m.originalParent
}

func (m *ChoiceCase) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.dataDefs != nil {
		copy.dataDefs = make([]Definition, len(m.dataDefs))
		copy.dataDefsIndex = make(map[string]Definition, len(m.dataDefs))
		for i, def := range m.dataDefs {
			copyDef := def.(cloneable).clone(&copy).(Definition)
			copy.dataDefs[i] = copyDef
			copy.dataDefsIndex[def.Ident()] = copyDef
		}
	}
	

	return &copy
}



// Ident is identity of Revision
func (m *Revision) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Revision) Parent() Meta {
	return m.parent
}

// Description of Revision
func (m *Revision) Description() string {
	return m.desc
}

func (m *Revision) setDescription(desc string) {
	m.desc = desc
}

func (m *Revision) Reference() string {
	return m.ref
}

func (m *Revision) setReference(ref string) {
	m.ref = ref
}

func (m *Revision) Extensions() []*Extension {
	return m.extensions
}

func (m *Revision) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}




// Ident is identity of Container
func (m *Container) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Container) Parent() Meta {
	return m.parent
}

// Description of Container
func (m *Container) Description() string {
	return m.desc
}

func (m *Container) setDescription(desc string) {
	m.desc = desc
}

func (m *Container) Reference() string {
	return m.ref
}

func (m *Container) setReference(ref string) {
	m.ref = ref
}

func (m *Container) Status() Status {
	return m.status
}

func (m *Container) setStatus(status Status) {
	m.status = status
}

func (m *Container) Extensions() []*Extension {
	return m.extensions
}

func (m *Container) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Container) DataDefinitions() []Definition {
	return m.dataDefs
}

func (m *Container) DataDefinition(ident string) Definition {
	return m.dataDefsIndex[ident]
}

func (m *Container) addDataDefinition(d Definition) error {
	if c, isChoice := d.(*Choice); isChoice {
		for _, k := range c.Cases() {
			for _, kdef := range k.DataDefinitions() {
				// recurse in case it's another choice
				if err := m.indexDataDefinition(kdef); err != nil {
					return err
				}
			}
		}
 	}
	
	if err := m.indexDataDefinition(d); err != nil {
		return err
	}
	m.dataDefs = append(m.dataDefs, d)
	return nil
}

func (m *Container) indexDataDefinition(def Definition) error {
	if m.dataDefsIndex == nil {
		m.dataDefsIndex = make(map[string]Definition)
	} else if _, exists := m.dataDefsIndex[def.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", def.Ident(), m.Ident())
	}	
	m.dataDefsIndex[def.Ident()] = def
	return nil
}

func (m *Container) popDataDefinitions() []Definition {
	orig := m.dataDefs
	m.dataDefs = make([]Definition, 0, len(orig))
	for key := range m.dataDefsIndex {
		delete(m.dataDefsIndex, key)
	}
	return orig
}
func (m *Container) IsRecursive() bool {
	return m.recursive
}

func (m *Container) markRecursive() {
	m.recursive = true
}


func (m *Container) Groupings() map[string]*Grouping {
	return m.groupings
}

func (m *Container) addGrouping(g *Grouping) error {
	g.parent = m
	if m.groupings == nil {
		m.groupings = make(map[string]*Grouping)
	} else if _, exists := m.groupings[g.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", g.Ident(), m.Ident())
	}
    m.groupings[g.Ident()] = g
	return nil
}

func (m *Container) Typedefs() map[string]*Typedef {
	return m.typedefs
}

func (m *Container) addTypedef(t *Typedef) error {
	t.parent = m
	if m.typedefs == nil {
		m.typedefs = make(map[string]*Typedef)
	} else if _, exists := m.typedefs[t.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", t.Ident(), m.Ident())
	}
    m.typedefs[t.Ident()] = t
	return nil
}

func (m *Container) Musts() []*Must {
	return m.musts
}

func (m *Container) addMust(x *Must) {
	x.parent = m
    m.musts = append(m.musts, x)
}

func (m *Container) setMusts(x []*Must) {
    m.musts = x
}

func (m *Container) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *Container) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

func (m *Container) When() *When {
	return m.when
}

func (m *Container) setWhen(w *When) {
	w.parent = m
    m.when = w
}

func (m *Container) Actions() map[string]*Rpc {
	return m.actions
}

func (m *Container) addAction(a *Rpc) error {
	a.parent = m
	if m.actions == nil {
		m.actions = make(map[string]*Rpc)
	} else if _, exists := m.actions[a.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", a.Ident(), m.Ident())
	}	
    m.actions[a.Ident()] = a
	return nil
}

func (m *Container) setActions(actions map[string]*Rpc) {
	m.actions = actions
}

func (m *Container) Notifications() map[string]*Notification {
	return m.notifications
}

func (m *Container) addNotification(n *Notification) error {
	n.parent = m
	if m.notifications == nil {
		m.notifications = make(map[string]*Notification)
	} else if _, exists := m.notifications[n.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", n.Ident(), m.Ident())
	}
    m.notifications[n.Ident()] = n
	return nil
}

func (m *Container) setNotifications(notifications map[string]*Notification) {
	m.notifications = notifications
}


// Definition can be a data defintion, action or notification
func (m *Container) Definition(ident string) Definition {
	if x, found := m.notifications[ident]; found {
		return x
	}
	
	if x, found := m.actions[ident]; found {
		return x
	}
	
	if x, found := m.dataDefsIndex[ident]; found {
		return x
	}
	
	return nil
}

func (m *Container) Config() bool {
	return *m.configPtr
}

func (m *Container) setConfig(c bool) {
	m.configPtr = &c
}

func (m *Container) IsConfigSet() bool {
	return m.configPtr != nil
}

func (m *Container) Mandatory() bool {
	return m.mandatoryPtr != nil && *m.mandatoryPtr
}

func (m *Container) setMandatory(b bool) {
	m.mandatoryPtr = &b
}

func (m *Container) IsMandatorySet() bool {
	return m.mandatoryPtr != nil
}

// Presence describes what the existance of this container in
// the data model means.
// https://tools.ietf.org/html/rfc7950#section-7.5.1
func (m *Container) Presence() string {
	return m.presence
}

func (m *Container) setPresence(p string) {
	m.presence = p
}

func (m *Container) getOriginalParent() Definition {
	return m.originalParent
}

func (m *Container) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.notifications != nil {
		copy.notifications = make(map[string]*Notification, len(m.notifications))
		for ident, notif := range m.notifications {
			copy.notifications[ident] = notif.clone(&copy).(*Notification)
		}
	}
	
	if m.actions != nil {
		copy.actions = make(map[string]*Rpc, len(m.actions))
		for ident, action := range m.actions {
			copy.actions[ident] = action.clone(&copy).(*Rpc)
		}
	}
	
	if m.dataDefs != nil {
		copy.dataDefs = make([]Definition, len(m.dataDefs))
		copy.dataDefsIndex = make(map[string]Definition, len(m.dataDefs))
		for i, def := range m.dataDefs {
			copyDef := def.(cloneable).clone(&copy).(Definition)
			copy.dataDefs[i] = copyDef
			copy.dataDefsIndex[def.Ident()] = copyDef
		}
	}
	
	if m.musts != nil {
		copy.musts = make([]*Must, len(m.musts))
		for i, must := range m.musts {
			copy.musts[i] = must.clone(&copy).(*Must)
		}
	}
	

	return &copy
}



// Ident is identity of List
func (m *List) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *List) Parent() Meta {
	return m.parent
}

// Description of List
func (m *List) Description() string {
	return m.desc
}

func (m *List) setDescription(desc string) {
	m.desc = desc
}

func (m *List) Reference() string {
	return m.ref
}

func (m *List) setReference(ref string) {
	m.ref = ref
}

func (m *List) Status() Status {
	return m.status
}

func (m *List) setStatus(status Status) {
	m.status = status
}

func (m *List) Extensions() []*Extension {
	return m.extensions
}

func (m *List) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *List) DataDefinitions() []Definition {
	return m.dataDefs
}

func (m *List) DataDefinition(ident string) Definition {
	return m.dataDefsIndex[ident]
}

func (m *List) addDataDefinition(d Definition) error {
	if c, isChoice := d.(*Choice); isChoice {
		for _, k := range c.Cases() {
			for _, kdef := range k.DataDefinitions() {
				// recurse in case it's another choice
				if err := m.indexDataDefinition(kdef); err != nil {
					return err
				}
			}
		}
 	}
	
	if err := m.indexDataDefinition(d); err != nil {
		return err
	}
	m.dataDefs = append(m.dataDefs, d)
	return nil
}

func (m *List) indexDataDefinition(def Definition) error {
	if m.dataDefsIndex == nil {
		m.dataDefsIndex = make(map[string]Definition)
	} else if _, exists := m.dataDefsIndex[def.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", def.Ident(), m.Ident())
	}	
	m.dataDefsIndex[def.Ident()] = def
	return nil
}

func (m *List) popDataDefinitions() []Definition {
	orig := m.dataDefs
	m.dataDefs = make([]Definition, 0, len(orig))
	for key := range m.dataDefsIndex {
		delete(m.dataDefsIndex, key)
	}
	return orig
}
func (m *List) IsRecursive() bool {
	return m.recursive
}

func (m *List) markRecursive() {
	m.recursive = true
}


func (m *List) Groupings() map[string]*Grouping {
	return m.groupings
}

func (m *List) addGrouping(g *Grouping) error {
	g.parent = m
	if m.groupings == nil {
		m.groupings = make(map[string]*Grouping)
	} else if _, exists := m.groupings[g.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", g.Ident(), m.Ident())
	}
    m.groupings[g.Ident()] = g
	return nil
}

func (m *List) Typedefs() map[string]*Typedef {
	return m.typedefs
}

func (m *List) addTypedef(t *Typedef) error {
	t.parent = m
	if m.typedefs == nil {
		m.typedefs = make(map[string]*Typedef)
	} else if _, exists := m.typedefs[t.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", t.Ident(), m.Ident())
	}
    m.typedefs[t.Ident()] = t
	return nil
}

func (m *List) Musts() []*Must {
	return m.musts
}

func (m *List) addMust(x *Must) {
	x.parent = m
    m.musts = append(m.musts, x)
}

func (m *List) setMusts(x []*Must) {
    m.musts = x
}

func (m *List) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *List) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

func (m *List) When() *When {
	return m.when
}

func (m *List) setWhen(w *When) {
	w.parent = m
    m.when = w
}

func (m *List) Actions() map[string]*Rpc {
	return m.actions
}

func (m *List) addAction(a *Rpc) error {
	a.parent = m
	if m.actions == nil {
		m.actions = make(map[string]*Rpc)
	} else if _, exists := m.actions[a.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", a.Ident(), m.Ident())
	}	
    m.actions[a.Ident()] = a
	return nil
}

func (m *List) setActions(actions map[string]*Rpc) {
	m.actions = actions
}

func (m *List) Notifications() map[string]*Notification {
	return m.notifications
}

func (m *List) addNotification(n *Notification) error {
	n.parent = m
	if m.notifications == nil {
		m.notifications = make(map[string]*Notification)
	} else if _, exists := m.notifications[n.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", n.Ident(), m.Ident())
	}
    m.notifications[n.Ident()] = n
	return nil
}

func (m *List) setNotifications(notifications map[string]*Notification) {
	m.notifications = notifications
}


// Definition can be a data defintion, action or notification
func (m *List) Definition(ident string) Definition {
	if x, found := m.notifications[ident]; found {
		return x
	}
	
	if x, found := m.actions[ident]; found {
		return x
	}
	
	if x, found := m.dataDefsIndex[ident]; found {
		return x
	}
	
	return nil
}

func (m *List) Config() bool {
	return *m.configPtr
}

func (m *List) setConfig(c bool) {
	m.configPtr = &c
}

func (m *List) IsConfigSet() bool {
	return m.configPtr != nil
}

func (m *List) Mandatory() bool {
	return m.mandatoryPtr != nil && *m.mandatoryPtr
}

func (m *List) setMandatory(b bool) {
	m.mandatoryPtr = &b
}

func (m *List) IsMandatorySet() bool {
	return m.mandatoryPtr != nil
}

func (m *List) MinElements() int {
	if m.minElementsPtr != nil {
		return *m.minElementsPtr
	}
	return 0
}

func (m *List) setMinElements(i int) {
	m.minElementsPtr = &i
}

func (m *List) IsMinElementsSet() bool {
	return m.minElementsPtr != nil
}

// MaxElements return 0 when unbounded
func (m *List) MaxElements() int {
	if m.maxElementsPtr != nil {
		return *m.maxElementsPtr
	}
	return 0
}

func (m *List) setMaxElements(i int) {
	m.maxElementsPtr = &i
}

func (m *List) IsMaxElementsSet() bool {
	return m.maxElementsPtr != nil
}

func (m *List) Unbounded() bool {
	if m.unboundedPtr != nil {
		return *m.unboundedPtr
	}
	return m.maxElementsPtr == nil
}

func (m *List) setUnbounded(b bool) {
	m.unboundedPtr = &b
}

func (m *List) IsUnboundedSet() bool {
	return m.unboundedPtr != nil
}


func (m *List) OrderedBy() OrderedBy {
	return m.orderedBy
}

func (m *List) setOrderedBy(o OrderedBy) {
	m.orderedBy = o
}

// Unique is list of fields (or compound fields) that must be unque in the
// list of items. If there is a key listed, that is implicitly unique and would
// not be listed here.
func (m *List) Unique() [][]string {
	return m.unique
}

func (m *List) setUnique(unique [][]string) {
	m.unique = unique
}

func (m *List) getOriginalParent() Definition {
	return m.originalParent
}

func (m *List) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.notifications != nil {
		copy.notifications = make(map[string]*Notification, len(m.notifications))
		for ident, notif := range m.notifications {
			copy.notifications[ident] = notif.clone(&copy).(*Notification)
		}
	}
	
	if m.actions != nil {
		copy.actions = make(map[string]*Rpc, len(m.actions))
		for ident, action := range m.actions {
			copy.actions[ident] = action.clone(&copy).(*Rpc)
		}
	}
	
	if m.dataDefs != nil {
		copy.dataDefs = make([]Definition, len(m.dataDefs))
		copy.dataDefsIndex = make(map[string]Definition, len(m.dataDefs))
		for i, def := range m.dataDefs {
			copyDef := def.(cloneable).clone(&copy).(Definition)
			copy.dataDefs[i] = copyDef
			copy.dataDefsIndex[def.Ident()] = copyDef
		}
	}
	
	if m.musts != nil {
		copy.musts = make([]*Must, len(m.musts))
		for i, must := range m.musts {
			copy.musts[i] = must.clone(&copy).(*Must)
		}
	}
	

	return &copy
}



// Ident is identity of Leaf
func (m *Leaf) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Leaf) Parent() Meta {
	return m.parent
}

// Description of Leaf
func (m *Leaf) Description() string {
	return m.desc
}

func (m *Leaf) setDescription(desc string) {
	m.desc = desc
}

func (m *Leaf) Reference() string {
	return m.ref
}

func (m *Leaf) setReference(ref string) {
	m.ref = ref
}

func (m *Leaf) Status() Status {
	return m.status
}

func (m *Leaf) setStatus(status Status) {
	m.status = status
}

func (m *Leaf) Extensions() []*Extension {
	return m.extensions
}

func (m *Leaf) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Leaf) Musts() []*Must {
	return m.musts
}

func (m *Leaf) addMust(x *Must) {
	x.parent = m
    m.musts = append(m.musts, x)
}

func (m *Leaf) setMusts(x []*Must) {
    m.musts = x
}

func (m *Leaf) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *Leaf) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

func (m *Leaf) When() *When {
	return m.when
}

func (m *Leaf) setWhen(w *When) {
	w.parent = m
    m.when = w
}

func (m *Leaf) Config() bool {
	return *m.configPtr
}

func (m *Leaf) setConfig(c bool) {
	m.configPtr = &c
}

func (m *Leaf) IsConfigSet() bool {
	return m.configPtr != nil
}

func (m *Leaf) Mandatory() bool {
	return m.mandatoryPtr != nil && *m.mandatoryPtr
}

func (m *Leaf) setMandatory(b bool) {
	m.mandatoryPtr = &b
}

func (m *Leaf) IsMandatorySet() bool {
	return m.mandatoryPtr != nil
}

func (m *Leaf) Type() *Type {
	return m.dtype
}

func (m *Leaf) setType(t *Type) {
	m.dtype = t
}

func (m *Leaf) Units() string{
	return m.units
}

func (m *Leaf) setUnits(u string) {
    m.units = u
}

func (m *Leaf) Default() string {
	if m.defaultVal == nil {
		return ""
	}
	return *m.defaultVal
}

func (m *Leaf) HasDefault() bool {
	return m.defaultVal != nil
}

func (m *Leaf) addDefault(d string) {
	if m.defaultVal != nil {
		panic("default already set")
	}
	m.defaultVal = &d
}

func (m *Leaf) DefaultValue() interface{} {
	return m.Default()
}

func (m *Leaf) setDefaultValue(d interface{}) {
	if s, valid := d.(string); valid {
		m.addDefault(s)
	} else {
		panic("expected string")
	}
}


func (m *Leaf) setDefault(d string) {
	m.defaultVal = &d
}

func (m *Leaf) clearDefault() {
    m.defaultVal = nil
}

func (m *Leaf) getOriginalParent() Definition {
	return m.originalParent
}

func (m *Leaf) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.musts != nil {
		copy.musts = make([]*Must, len(m.musts))
		for i, must := range m.musts {
			copy.musts[i] = must.clone(&copy).(*Must)
		}
	}
	

	return &copy
}



// Ident is identity of LeafList
func (m *LeafList) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *LeafList) Parent() Meta {
	return m.parent
}

// Description of LeafList
func (m *LeafList) Description() string {
	return m.desc
}

func (m *LeafList) setDescription(desc string) {
	m.desc = desc
}

func (m *LeafList) Reference() string {
	return m.ref
}

func (m *LeafList) setReference(ref string) {
	m.ref = ref
}

func (m *LeafList) Status() Status {
	return m.status
}

func (m *LeafList) setStatus(status Status) {
	m.status = status
}

func (m *LeafList) Extensions() []*Extension {
	return m.extensions
}

func (m *LeafList) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *LeafList) Musts() []*Must {
	return m.musts
}

func (m *LeafList) addMust(x *Must) {
	x.parent = m
    m.musts = append(m.musts, x)
}

func (m *LeafList) setMusts(x []*Must) {
    m.musts = x
}

func (m *LeafList) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *LeafList) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

func (m *LeafList) When() *When {
	return m.when
}

func (m *LeafList) setWhen(w *When) {
	w.parent = m
    m.when = w
}

func (m *LeafList) Config() bool {
	return *m.configPtr
}

func (m *LeafList) setConfig(c bool) {
	m.configPtr = &c
}

func (m *LeafList) IsConfigSet() bool {
	return m.configPtr != nil
}

func (m *LeafList) Mandatory() bool {
	return m.mandatoryPtr != nil && *m.mandatoryPtr
}

func (m *LeafList) setMandatory(b bool) {
	m.mandatoryPtr = &b
}

func (m *LeafList) IsMandatorySet() bool {
	return m.mandatoryPtr != nil
}

func (m *LeafList) MinElements() int {
	if m.minElementsPtr != nil {
		return *m.minElementsPtr
	}
	return 0
}

func (m *LeafList) setMinElements(i int) {
	m.minElementsPtr = &i
}

func (m *LeafList) IsMinElementsSet() bool {
	return m.minElementsPtr != nil
}

// MaxElements return 0 when unbounded
func (m *LeafList) MaxElements() int {
	if m.maxElementsPtr != nil {
		return *m.maxElementsPtr
	}
	return 0
}

func (m *LeafList) setMaxElements(i int) {
	m.maxElementsPtr = &i
}

func (m *LeafList) IsMaxElementsSet() bool {
	return m.maxElementsPtr != nil
}

func (m *LeafList) Unbounded() bool {
	if m.unboundedPtr != nil {
		return *m.unboundedPtr
	}
	return m.maxElementsPtr == nil
}

func (m *LeafList) setUnbounded(b bool) {
	m.unboundedPtr = &b
}

func (m *LeafList) IsUnboundedSet() bool {
	return m.unboundedPtr != nil
}


func (m *LeafList) OrderedBy() OrderedBy {
	return m.orderedBy
}

func (m *LeafList) setOrderedBy(o OrderedBy) {
	m.orderedBy = o
}

func (m *LeafList) Type() *Type {
	return m.dtype
}

func (m *LeafList) setType(t *Type) {
	m.dtype = t
}

func (m *LeafList) Units() string{
	return m.units
}

func (m *LeafList) setUnits(u string) {
    m.units = u
}

func (m *LeafList) Default() []string {
	return m.defaultVals
}

func (m *LeafList) HasDefault() bool {
	return m.defaultVals != nil
}

func (m *LeafList) DefaultValue() interface{} {
	return m.Default()
}

func (m *LeafList) setDefaultValue(d interface{}) {
	if s, valid := d.([]string); valid {
		m.defaultVals = s
	} else {
		panic("expected []string")
	}
}

func (m *LeafList) addDefault(d string) {
	m.defaultVals = append(m.defaultVals, d)
}

func (m *LeafList) setDefault(d []string) {
	m.defaultVals = d
}

func (m *LeafList) clearDefault() {
    m.defaultVals = nil
}

func (m *LeafList) getOriginalParent() Definition {
	return m.originalParent
}

func (m *LeafList) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.musts != nil {
		copy.musts = make([]*Must, len(m.musts))
		for i, must := range m.musts {
			copy.musts[i] = must.clone(&copy).(*Must)
		}
	}
	

	return &copy
}



// Ident is identity of Any
func (m *Any) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Any) Parent() Meta {
	return m.parent
}

// Description of Any
func (m *Any) Description() string {
	return m.desc
}

func (m *Any) setDescription(desc string) {
	m.desc = desc
}

func (m *Any) Reference() string {
	return m.ref
}

func (m *Any) setReference(ref string) {
	m.ref = ref
}

func (m *Any) Status() Status {
	return m.status
}

func (m *Any) setStatus(status Status) {
	m.status = status
}

func (m *Any) Extensions() []*Extension {
	return m.extensions
}

func (m *Any) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Any) Musts() []*Must {
	return m.musts
}

func (m *Any) addMust(x *Must) {
	x.parent = m
    m.musts = append(m.musts, x)
}

func (m *Any) setMusts(x []*Must) {
    m.musts = x
}

func (m *Any) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *Any) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

func (m *Any) When() *When {
	return m.when
}

func (m *Any) setWhen(w *When) {
	w.parent = m
    m.when = w
}

func (m *Any) Config() bool {
	return *m.configPtr
}

func (m *Any) setConfig(c bool) {
	m.configPtr = &c
}

func (m *Any) IsConfigSet() bool {
	return m.configPtr != nil
}

func (m *Any) Mandatory() bool {
	return m.mandatoryPtr != nil && *m.mandatoryPtr
}

func (m *Any) setMandatory(b bool) {
	m.mandatoryPtr = &b
}

func (m *Any) IsMandatorySet() bool {
	return m.mandatoryPtr != nil
}

func (m *Any) getOriginalParent() Definition {
	return m.originalParent
}

func (m *Any) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.musts != nil {
		copy.musts = make([]*Must, len(m.musts))
		for i, must := range m.musts {
			copy.musts[i] = must.clone(&copy).(*Must)
		}
	}
	

	return &copy
}



// Ident is identity of Grouping
func (m *Grouping) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Grouping) Parent() Meta {
	return m.parent
}

// Description of Grouping
func (m *Grouping) Description() string {
	return m.desc
}

func (m *Grouping) setDescription(desc string) {
	m.desc = desc
}

func (m *Grouping) Reference() string {
	return m.ref
}

func (m *Grouping) setReference(ref string) {
	m.ref = ref
}

func (m *Grouping) Extensions() []*Extension {
	return m.extensions
}

func (m *Grouping) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Grouping) DataDefinitions() []Definition {
	return m.dataDefs
}

func (m *Grouping) DataDefinition(ident string) Definition {
	return m.dataDefsIndex[ident]
}

func (m *Grouping) addDataDefinition(d Definition) error {
	if c, isChoice := d.(*Choice); isChoice {
		for _, k := range c.Cases() {
			for _, kdef := range k.DataDefinitions() {
				// recurse in case it's another choice
				if err := m.indexDataDefinition(kdef); err != nil {
					return err
				}
			}
		}
 	}
	
	if err := m.indexDataDefinition(d); err != nil {
		return err
	}
	m.dataDefs = append(m.dataDefs, d)
	return nil
}

func (m *Grouping) indexDataDefinition(def Definition) error {
	if m.dataDefsIndex == nil {
		m.dataDefsIndex = make(map[string]Definition)
	} else if _, exists := m.dataDefsIndex[def.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", def.Ident(), m.Ident())
	}	
	m.dataDefsIndex[def.Ident()] = def
	return nil
}

func (m *Grouping) popDataDefinitions() []Definition {
	orig := m.dataDefs
	m.dataDefs = make([]Definition, 0, len(orig))
	for key := range m.dataDefsIndex {
		delete(m.dataDefsIndex, key)
	}
	return orig
}
func (m *Grouping) IsRecursive() bool {
	return false
}

func (m *Grouping) markRecursive() {
	panic("Cannot mark Grouping recursive")
}



func (m *Grouping) Groupings() map[string]*Grouping {
	return m.groupings
}

func (m *Grouping) addGrouping(g *Grouping) error {
	g.parent = m
	if m.groupings == nil {
		m.groupings = make(map[string]*Grouping)
	} else if _, exists := m.groupings[g.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", g.Ident(), m.Ident())
	}
    m.groupings[g.Ident()] = g
	return nil
}

func (m *Grouping) Typedefs() map[string]*Typedef {
	return m.typedefs
}

func (m *Grouping) addTypedef(t *Typedef) error {
	t.parent = m
	if m.typedefs == nil {
		m.typedefs = make(map[string]*Typedef)
	} else if _, exists := m.typedefs[t.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", t.Ident(), m.Ident())
	}
    m.typedefs[t.Ident()] = t
	return nil
}

func (m *Grouping) Actions() map[string]*Rpc {
	return m.actions
}

func (m *Grouping) addAction(a *Rpc) error {
	a.parent = m
	if m.actions == nil {
		m.actions = make(map[string]*Rpc)
	} else if _, exists := m.actions[a.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", a.Ident(), m.Ident())
	}	
    m.actions[a.Ident()] = a
	return nil
}

func (m *Grouping) setActions(actions map[string]*Rpc) {
	m.actions = actions
}

func (m *Grouping) Notifications() map[string]*Notification {
	return m.notifications
}

func (m *Grouping) addNotification(n *Notification) error {
	n.parent = m
	if m.notifications == nil {
		m.notifications = make(map[string]*Notification)
	} else if _, exists := m.notifications[n.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", n.Ident(), m.Ident())
	}
    m.notifications[n.Ident()] = n
	return nil
}

func (m *Grouping) setNotifications(notifications map[string]*Notification) {
	m.notifications = notifications
}


// Definition can be a data defintion, action or notification
func (m *Grouping) Definition(ident string) Definition {
	if x, found := m.notifications[ident]; found {
		return x
	}
	
	if x, found := m.actions[ident]; found {
		return x
	}
	
	if x, found := m.dataDefsIndex[ident]; found {
		return x
	}
	
	return nil
}

func (m *Grouping) getOriginalParent() Definition {
	return m.originalParent
}

func (m *Grouping) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.notifications != nil {
		copy.notifications = make(map[string]*Notification, len(m.notifications))
		for ident, notif := range m.notifications {
			copy.notifications[ident] = notif.clone(&copy).(*Notification)
		}
	}
	
	if m.actions != nil {
		copy.actions = make(map[string]*Rpc, len(m.actions))
		for ident, action := range m.actions {
			copy.actions[ident] = action.clone(&copy).(*Rpc)
		}
	}
	
	if m.dataDefs != nil {
		copy.dataDefs = make([]Definition, len(m.dataDefs))
		copy.dataDefsIndex = make(map[string]Definition, len(m.dataDefs))
		for i, def := range m.dataDefs {
			copyDef := def.(cloneable).clone(&copy).(Definition)
			copy.dataDefs[i] = copyDef
			copy.dataDefsIndex[def.Ident()] = copyDef
		}
	}
	

	return &copy
}



// Ident is identity of Uses
func (m *Uses) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Uses) Parent() Meta {
	return m.parent
}

// Description of Uses
func (m *Uses) Description() string {
	return m.desc
}

func (m *Uses) setDescription(desc string) {
	m.desc = desc
}

func (m *Uses) Reference() string {
	return m.ref
}

func (m *Uses) setReference(ref string) {
	m.ref = ref
}

func (m *Uses) Extensions() []*Extension {
	return m.extensions
}

func (m *Uses) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Uses) Augments() []*Augment {
	return m.augments
}

func (m *Uses) addAugments(a *Augment) {
	a.parent = m
	m.augments = append(m.augments, a)
}

func (m *Uses) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *Uses) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

func (m *Uses) When() *When {
	return m.when
}

func (m *Uses) setWhen(w *When) {
	w.parent = m
    m.when = w
}

func (m *Uses) getOriginalParent() Definition {
	return m.originalParent
}

func (m *Uses) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent

	return &copy
}



// Ident is identity of Refine
func (m *Refine) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Refine) Parent() Meta {
	return m.parent
}

// Description of Refine
func (m *Refine) Description() string {
	return m.desc
}

func (m *Refine) setDescription(desc string) {
	m.desc = desc
}

func (m *Refine) Reference() string {
	return m.ref
}

func (m *Refine) setReference(ref string) {
	m.ref = ref
}

func (m *Refine) Extensions() []*Extension {
	return m.extensions
}

func (m *Refine) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Refine) Musts() []*Must {
	return m.musts
}

func (m *Refine) addMust(x *Must) {
	x.parent = m
    m.musts = append(m.musts, x)
}

func (m *Refine) setMusts(x []*Must) {
    m.musts = x
}

func (m *Refine) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *Refine) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

func (m *Refine) Config() bool {
	return *m.configPtr
}

func (m *Refine) setConfig(c bool) {
	m.configPtr = &c
}

func (m *Refine) IsConfigSet() bool {
	return m.configPtr != nil
}

func (m *Refine) Mandatory() bool {
	return m.mandatoryPtr != nil && *m.mandatoryPtr
}

func (m *Refine) setMandatory(b bool) {
	m.mandatoryPtr = &b
}

func (m *Refine) IsMandatorySet() bool {
	return m.mandatoryPtr != nil
}

func (m *Refine) MinElements() int {
	if m.minElementsPtr != nil {
		return *m.minElementsPtr
	}
	return 0
}

func (m *Refine) setMinElements(i int) {
	m.minElementsPtr = &i
}

func (m *Refine) IsMinElementsSet() bool {
	return m.minElementsPtr != nil
}

// MaxElements return 0 when unbounded
func (m *Refine) MaxElements() int {
	if m.maxElementsPtr != nil {
		return *m.maxElementsPtr
	}
	return 0
}

func (m *Refine) setMaxElements(i int) {
	m.maxElementsPtr = &i
}

func (m *Refine) IsMaxElementsSet() bool {
	return m.maxElementsPtr != nil
}

func (m *Refine) Unbounded() bool {
	if m.unboundedPtr != nil {
		return *m.unboundedPtr
	}
	return m.maxElementsPtr == nil
}

func (m *Refine) setUnbounded(b bool) {
	m.unboundedPtr = &b
}

func (m *Refine) IsUnboundedSet() bool {
	return m.unboundedPtr != nil
}


// Presence describes what the existance of this container in
// the data model means.
// https://tools.ietf.org/html/rfc7950#section-7.5.1
func (m *Refine) Presence() string {
	return m.presence
}

func (m *Refine) setPresence(p string) {
	m.presence = p
}

func (m *Refine) Default() string {
	if m.defaultVal == nil {
		return ""
	}
	return *m.defaultVal
}

func (m *Refine) HasDefault() bool {
	return m.defaultVal != nil
}

func (m *Refine) addDefault(d string) {
	if m.defaultVal != nil {
		panic("default already set")
	}
	m.defaultVal = &d
}

func (m *Refine) DefaultValue() interface{} {
	return m.Default()
}

func (m *Refine) setDefaultValue(d interface{}) {
	if s, valid := d.(string); valid {
		m.addDefault(s)
	} else {
		panic("expected string")
	}
}


func (m *Refine) setDefault(d string) {
	m.defaultVal = &d
}

func (m *Refine) clearDefault() {
    m.defaultVal = nil
}



// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *RpcInput) Parent() Meta {
	return m.parent
}

// Description of RpcInput
func (m *RpcInput) Description() string {
	return m.desc
}

func (m *RpcInput) setDescription(desc string) {
	m.desc = desc
}

func (m *RpcInput) Reference() string {
	return m.ref
}

func (m *RpcInput) setReference(ref string) {
	m.ref = ref
}

func (m *RpcInput) Extensions() []*Extension {
	return m.extensions
}

func (m *RpcInput) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *RpcInput) DataDefinitions() []Definition {
	return m.dataDefs
}

func (m *RpcInput) DataDefinition(ident string) Definition {
	return m.dataDefsIndex[ident]
}

func (m *RpcInput) addDataDefinition(d Definition) error {
	if c, isChoice := d.(*Choice); isChoice {
		for _, k := range c.Cases() {
			for _, kdef := range k.DataDefinitions() {
				// recurse in case it's another choice
				if err := m.indexDataDefinition(kdef); err != nil {
					return err
				}
			}
		}
 	}
	
	if err := m.indexDataDefinition(d); err != nil {
		return err
	}
	m.dataDefs = append(m.dataDefs, d)
	return nil
}

func (m *RpcInput) indexDataDefinition(def Definition) error {
	if m.dataDefsIndex == nil {
		m.dataDefsIndex = make(map[string]Definition)
	} else if _, exists := m.dataDefsIndex[def.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", def.Ident(), m.Ident())
	}	
	m.dataDefsIndex[def.Ident()] = def
	return nil
}

func (m *RpcInput) popDataDefinitions() []Definition {
	orig := m.dataDefs
	m.dataDefs = make([]Definition, 0, len(orig))
	for key := range m.dataDefsIndex {
		delete(m.dataDefsIndex, key)
	}
	return orig
}
func (m *RpcInput) IsRecursive() bool {
	return false
}

func (m *RpcInput) markRecursive() {
	panic("Cannot mark RpcInput recursive")
}



func (m *RpcInput) Groupings() map[string]*Grouping {
	return m.groupings
}

func (m *RpcInput) addGrouping(g *Grouping) error {
	g.parent = m
	if m.groupings == nil {
		m.groupings = make(map[string]*Grouping)
	} else if _, exists := m.groupings[g.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", g.Ident(), m.Ident())
	}
    m.groupings[g.Ident()] = g
	return nil
}

func (m *RpcInput) Typedefs() map[string]*Typedef {
	return m.typedefs
}

func (m *RpcInput) addTypedef(t *Typedef) error {
	t.parent = m
	if m.typedefs == nil {
		m.typedefs = make(map[string]*Typedef)
	} else if _, exists := m.typedefs[t.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", t.Ident(), m.Ident())
	}
    m.typedefs[t.Ident()] = t
	return nil
}

func (m *RpcInput) Musts() []*Must {
	return m.musts
}

func (m *RpcInput) addMust(x *Must) {
	x.parent = m
    m.musts = append(m.musts, x)
}

func (m *RpcInput) setMusts(x []*Must) {
    m.musts = x
}

func (m *RpcInput) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *RpcInput) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

// Definition can be a data defintion, action or notification
func (m *RpcInput) Definition(ident string) Definition {
	if x, found := m.dataDefsIndex[ident]; found {
		return x
	}
	
	return nil
}

func (m *RpcInput) getOriginalParent() Definition {
	return m.originalParent
}

func (m *RpcInput) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.dataDefs != nil {
		copy.dataDefs = make([]Definition, len(m.dataDefs))
		copy.dataDefsIndex = make(map[string]Definition, len(m.dataDefs))
		for i, def := range m.dataDefs {
			copyDef := def.(cloneable).clone(&copy).(Definition)
			copy.dataDefs[i] = copyDef
			copy.dataDefsIndex[def.Ident()] = copyDef
		}
	}
	
	if m.musts != nil {
		copy.musts = make([]*Must, len(m.musts))
		for i, must := range m.musts {
			copy.musts[i] = must.clone(&copy).(*Must)
		}
	}
	

	return &copy
}



// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *RpcOutput) Parent() Meta {
	return m.parent
}

// Description of RpcOutput
func (m *RpcOutput) Description() string {
	return m.desc
}

func (m *RpcOutput) setDescription(desc string) {
	m.desc = desc
}

func (m *RpcOutput) Reference() string {
	return m.ref
}

func (m *RpcOutput) setReference(ref string) {
	m.ref = ref
}

func (m *RpcOutput) Extensions() []*Extension {
	return m.extensions
}

func (m *RpcOutput) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *RpcOutput) DataDefinitions() []Definition {
	return m.dataDefs
}

func (m *RpcOutput) DataDefinition(ident string) Definition {
	return m.dataDefsIndex[ident]
}

func (m *RpcOutput) addDataDefinition(d Definition) error {
	if c, isChoice := d.(*Choice); isChoice {
		for _, k := range c.Cases() {
			for _, kdef := range k.DataDefinitions() {
				// recurse in case it's another choice
				if err := m.indexDataDefinition(kdef); err != nil {
					return err
				}
			}
		}
 	}
	
	if err := m.indexDataDefinition(d); err != nil {
		return err
	}
	m.dataDefs = append(m.dataDefs, d)
	return nil
}

func (m *RpcOutput) indexDataDefinition(def Definition) error {
	if m.dataDefsIndex == nil {
		m.dataDefsIndex = make(map[string]Definition)
	} else if _, exists := m.dataDefsIndex[def.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", def.Ident(), m.Ident())
	}	
	m.dataDefsIndex[def.Ident()] = def
	return nil
}

func (m *RpcOutput) popDataDefinitions() []Definition {
	orig := m.dataDefs
	m.dataDefs = make([]Definition, 0, len(orig))
	for key := range m.dataDefsIndex {
		delete(m.dataDefsIndex, key)
	}
	return orig
}
func (m *RpcOutput) IsRecursive() bool {
	return false
}

func (m *RpcOutput) markRecursive() {
	panic("Cannot mark RpcOutput recursive")
}



func (m *RpcOutput) Groupings() map[string]*Grouping {
	return m.groupings
}

func (m *RpcOutput) addGrouping(g *Grouping) error {
	g.parent = m
	if m.groupings == nil {
		m.groupings = make(map[string]*Grouping)
	} else if _, exists := m.groupings[g.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", g.Ident(), m.Ident())
	}
    m.groupings[g.Ident()] = g
	return nil
}

func (m *RpcOutput) Typedefs() map[string]*Typedef {
	return m.typedefs
}

func (m *RpcOutput) addTypedef(t *Typedef) error {
	t.parent = m
	if m.typedefs == nil {
		m.typedefs = make(map[string]*Typedef)
	} else if _, exists := m.typedefs[t.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", t.Ident(), m.Ident())
	}
    m.typedefs[t.Ident()] = t
	return nil
}

func (m *RpcOutput) Musts() []*Must {
	return m.musts
}

func (m *RpcOutput) addMust(x *Must) {
	x.parent = m
    m.musts = append(m.musts, x)
}

func (m *RpcOutput) setMusts(x []*Must) {
    m.musts = x
}

func (m *RpcOutput) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *RpcOutput) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

// Definition can be a data defintion, action or notification
func (m *RpcOutput) Definition(ident string) Definition {
	if x, found := m.dataDefsIndex[ident]; found {
		return x
	}
	
	return nil
}

func (m *RpcOutput) getOriginalParent() Definition {
	return m.originalParent
}

func (m *RpcOutput) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.dataDefs != nil {
		copy.dataDefs = make([]Definition, len(m.dataDefs))
		copy.dataDefsIndex = make(map[string]Definition, len(m.dataDefs))
		for i, def := range m.dataDefs {
			copyDef := def.(cloneable).clone(&copy).(Definition)
			copy.dataDefs[i] = copyDef
			copy.dataDefsIndex[def.Ident()] = copyDef
		}
	}
	
	if m.musts != nil {
		copy.musts = make([]*Must, len(m.musts))
		for i, must := range m.musts {
			copy.musts[i] = must.clone(&copy).(*Must)
		}
	}
	

	return &copy
}



// Ident is identity of Rpc
func (m *Rpc) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Rpc) Parent() Meta {
	return m.parent
}

// Description of Rpc
func (m *Rpc) Description() string {
	return m.desc
}

func (m *Rpc) setDescription(desc string) {
	m.desc = desc
}

func (m *Rpc) Reference() string {
	return m.ref
}

func (m *Rpc) setReference(ref string) {
	m.ref = ref
}

func (m *Rpc) Status() Status {
	return m.status
}

func (m *Rpc) setStatus(status Status) {
	m.status = status
}

func (m *Rpc) Extensions() []*Extension {
	return m.extensions
}

func (m *Rpc) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Rpc) Groupings() map[string]*Grouping {
	return m.groupings
}

func (m *Rpc) addGrouping(g *Grouping) error {
	g.parent = m
	if m.groupings == nil {
		m.groupings = make(map[string]*Grouping)
	} else if _, exists := m.groupings[g.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", g.Ident(), m.Ident())
	}
    m.groupings[g.Ident()] = g
	return nil
}

func (m *Rpc) Typedefs() map[string]*Typedef {
	return m.typedefs
}

func (m *Rpc) addTypedef(t *Typedef) error {
	t.parent = m
	if m.typedefs == nil {
		m.typedefs = make(map[string]*Typedef)
	} else if _, exists := m.typedefs[t.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", t.Ident(), m.Ident())
	}
    m.typedefs[t.Ident()] = t
	return nil
}

func (m *Rpc) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *Rpc) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

func (m *Rpc) getOriginalParent() Definition {
	return m.originalParent
}

func (m *Rpc) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.input != nil {
		copy.input = m.input.clone(&copy).(*RpcInput)
	}
	if m.output != nil {
		copy.output = m.output.clone(&copy).(*RpcOutput)
	}
	

	return &copy
}



// Ident is identity of Notification
func (m *Notification) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Notification) Parent() Meta {
	return m.parent
}

// Description of Notification
func (m *Notification) Description() string {
	return m.desc
}

func (m *Notification) setDescription(desc string) {
	m.desc = desc
}

func (m *Notification) Reference() string {
	return m.ref
}

func (m *Notification) setReference(ref string) {
	m.ref = ref
}

func (m *Notification) Status() Status {
	return m.status
}

func (m *Notification) setStatus(status Status) {
	m.status = status
}

func (m *Notification) Extensions() []*Extension {
	return m.extensions
}

func (m *Notification) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Notification) DataDefinitions() []Definition {
	return m.dataDefs
}

func (m *Notification) DataDefinition(ident string) Definition {
	return m.dataDefsIndex[ident]
}

func (m *Notification) addDataDefinition(d Definition) error {
	if c, isChoice := d.(*Choice); isChoice {
		for _, k := range c.Cases() {
			for _, kdef := range k.DataDefinitions() {
				// recurse in case it's another choice
				if err := m.indexDataDefinition(kdef); err != nil {
					return err
				}
			}
		}
 	}
	
	if err := m.indexDataDefinition(d); err != nil {
		return err
	}
	m.dataDefs = append(m.dataDefs, d)
	return nil
}

func (m *Notification) indexDataDefinition(def Definition) error {
	if m.dataDefsIndex == nil {
		m.dataDefsIndex = make(map[string]Definition)
	} else if _, exists := m.dataDefsIndex[def.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", def.Ident(), m.Ident())
	}	
	m.dataDefsIndex[def.Ident()] = def
	return nil
}

func (m *Notification) popDataDefinitions() []Definition {
	orig := m.dataDefs
	m.dataDefs = make([]Definition, 0, len(orig))
	for key := range m.dataDefsIndex {
		delete(m.dataDefsIndex, key)
	}
	return orig
}
func (m *Notification) IsRecursive() bool {
	return false
}

func (m *Notification) markRecursive() {
	panic("Cannot mark Notification recursive")
}



func (m *Notification) Groupings() map[string]*Grouping {
	return m.groupings
}

func (m *Notification) addGrouping(g *Grouping) error {
	g.parent = m
	if m.groupings == nil {
		m.groupings = make(map[string]*Grouping)
	} else if _, exists := m.groupings[g.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", g.Ident(), m.Ident())
	}
    m.groupings[g.Ident()] = g
	return nil
}

func (m *Notification) Typedefs() map[string]*Typedef {
	return m.typedefs
}

func (m *Notification) addTypedef(t *Typedef) error {
	t.parent = m
	if m.typedefs == nil {
		m.typedefs = make(map[string]*Typedef)
	} else if _, exists := m.typedefs[t.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", t.Ident(), m.Ident())
	}
    m.typedefs[t.Ident()] = t
	return nil
}

func (m *Notification) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *Notification) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

// Definition can be a data defintion, action or notification
func (m *Notification) Definition(ident string) Definition {
	if x, found := m.dataDefsIndex[ident]; found {
		return x
	}
	
	return nil
}

func (m *Notification) getOriginalParent() Definition {
	return m.originalParent
}

func (m *Notification) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.dataDefs != nil {
		copy.dataDefs = make([]Definition, len(m.dataDefs))
		copy.dataDefsIndex = make(map[string]Definition, len(m.dataDefs))
		for i, def := range m.dataDefs {
			copyDef := def.(cloneable).clone(&copy).(Definition)
			copy.dataDefs[i] = copyDef
			copy.dataDefsIndex[def.Ident()] = copyDef
		}
	}
	

	return &copy
}



// Ident is identity of Typedef
func (m *Typedef) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Typedef) Parent() Meta {
	return m.parent
}

// Description of Typedef
func (m *Typedef) Description() string {
	return m.desc
}

func (m *Typedef) setDescription(desc string) {
	m.desc = desc
}

func (m *Typedef) Reference() string {
	return m.ref
}

func (m *Typedef) setReference(ref string) {
	m.ref = ref
}

func (m *Typedef) Extensions() []*Extension {
	return m.extensions
}

func (m *Typedef) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Typedef) Type() *Type {
	return m.dtype
}

func (m *Typedef) setType(t *Type) {
	m.dtype = t
}

func (m *Typedef) Units() string{
	return m.units
}

func (m *Typedef) setUnits(u string) {
    m.units = u
}

func (m *Typedef) Default() string {
	if m.defaultVal == nil {
		return ""
	}
	return *m.defaultVal
}

func (m *Typedef) HasDefault() bool {
	return m.defaultVal != nil
}

func (m *Typedef) addDefault(d string) {
	if m.defaultVal != nil {
		panic("default already set")
	}
	m.defaultVal = &d
}

func (m *Typedef) DefaultValue() interface{} {
	return m.Default()
}

func (m *Typedef) setDefaultValue(d interface{}) {
	if s, valid := d.(string); valid {
		m.addDefault(s)
	} else {
		panic("expected string")
	}
}


func (m *Typedef) setDefault(d string) {
	m.defaultVal = &d
}

func (m *Typedef) clearDefault() {
    m.defaultVal = nil
}

func (m *Typedef) getOriginalParent() Definition {
	return m.originalParent
}



// Ident is identity of Augment
func (m *Augment) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Augment) Parent() Meta {
	return m.parent
}

// Description of Augment
func (m *Augment) Description() string {
	return m.desc
}

func (m *Augment) setDescription(desc string) {
	m.desc = desc
}

func (m *Augment) Reference() string {
	return m.ref
}

func (m *Augment) setReference(ref string) {
	m.ref = ref
}

func (m *Augment) Extensions() []*Extension {
	return m.extensions
}

func (m *Augment) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Augment) DataDefinitions() []Definition {
	return m.dataDefs
}

func (m *Augment) DataDefinition(ident string) Definition {
	return m.dataDefsIndex[ident]
}

func (m *Augment) addDataDefinition(d Definition) error {
	if c, isChoice := d.(*Choice); isChoice {
		for _, k := range c.Cases() {
			for _, kdef := range k.DataDefinitions() {
				// recurse in case it's another choice
				if err := m.indexDataDefinition(kdef); err != nil {
					return err
				}
			}
		}
 	}
	
	if err := m.indexDataDefinition(d); err != nil {
		return err
	}
	m.dataDefs = append(m.dataDefs, d)
	return nil
}

func (m *Augment) indexDataDefinition(def Definition) error {
	if m.dataDefsIndex == nil {
		m.dataDefsIndex = make(map[string]Definition)
	} else if _, exists := m.dataDefsIndex[def.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", def.Ident(), m.Ident())
	}	
	m.dataDefsIndex[def.Ident()] = def
	return nil
}

func (m *Augment) popDataDefinitions() []Definition {
	orig := m.dataDefs
	m.dataDefs = make([]Definition, 0, len(orig))
	for key := range m.dataDefsIndex {
		delete(m.dataDefsIndex, key)
	}
	return orig
}
func (m *Augment) IsRecursive() bool {
	return false
}

func (m *Augment) markRecursive() {
	panic("Cannot mark Augment recursive")
}



func (m *Augment) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *Augment) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}

func (m *Augment) When() *When {
	return m.when
}

func (m *Augment) setWhen(w *When) {
	w.parent = m
    m.when = w
}

func (m *Augment) Actions() map[string]*Rpc {
	return m.actions
}

func (m *Augment) addAction(a *Rpc) error {
	a.parent = m
	if m.actions == nil {
		m.actions = make(map[string]*Rpc)
	} else if _, exists := m.actions[a.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", a.Ident(), m.Ident())
	}	
    m.actions[a.Ident()] = a
	return nil
}

func (m *Augment) setActions(actions map[string]*Rpc) {
	m.actions = actions
}

func (m *Augment) Notifications() map[string]*Notification {
	return m.notifications
}

func (m *Augment) addNotification(n *Notification) error {
	n.parent = m
	if m.notifications == nil {
		m.notifications = make(map[string]*Notification)
	} else if _, exists := m.notifications[n.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s", n.Ident(), m.Ident())
	}
    m.notifications[n.Ident()] = n
	return nil
}

func (m *Augment) setNotifications(notifications map[string]*Notification) {
	m.notifications = notifications
}


// Definition can be a data defintion, action or notification
func (m *Augment) Definition(ident string) Definition {
	if x, found := m.notifications[ident]; found {
		return x
	}
	
	if x, found := m.actions[ident]; found {
		return x
	}
	
	if x, found := m.dataDefsIndex[ident]; found {
		return x
	}
	
	return nil
}

func (m *Augment) getOriginalParent() Definition {
	return m.originalParent
}

func (m *Augment) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent
	if m.notifications != nil {
		copy.notifications = make(map[string]*Notification, len(m.notifications))
		for ident, notif := range m.notifications {
			copy.notifications[ident] = notif.clone(&copy).(*Notification)
		}
	}
	
	if m.actions != nil {
		copy.actions = make(map[string]*Rpc, len(m.actions))
		for ident, action := range m.actions {
			copy.actions[ident] = action.clone(&copy).(*Rpc)
		}
	}
	
	if m.dataDefs != nil {
		copy.dataDefs = make([]Definition, len(m.dataDefs))
		copy.dataDefsIndex = make(map[string]Definition, len(m.dataDefs))
		for i, def := range m.dataDefs {
			copyDef := def.(cloneable).clone(&copy).(Definition)
			copy.dataDefs[i] = copyDef
			copy.dataDefsIndex[def.Ident()] = copyDef
		}
	}
	

	return &copy
}



// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *AddDeviate) Parent() Meta {
	return m.parent
}

func (m *AddDeviate) Extensions() []*Extension {
	return m.extensions
}

func (m *AddDeviate) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *AddDeviate) Musts() []*Must {
	return m.musts
}

func (m *AddDeviate) addMust(x *Must) {
	x.parent = m
    m.musts = append(m.musts, x)
}

func (m *AddDeviate) setMusts(x []*Must) {
    m.musts = x
}

func (m *AddDeviate) Config() bool {
	return *m.configPtr
}

func (m *AddDeviate) setConfig(c bool) {
	m.configPtr = &c
}

func (m *AddDeviate) IsConfigSet() bool {
	return m.configPtr != nil
}

func (m *AddDeviate) Mandatory() bool {
	return m.mandatoryPtr != nil && *m.mandatoryPtr
}

func (m *AddDeviate) setMandatory(b bool) {
	m.mandatoryPtr = &b
}

func (m *AddDeviate) IsMandatorySet() bool {
	return m.mandatoryPtr != nil
}

func (m *AddDeviate) MinElements() int {
	if m.minElementsPtr != nil {
		return *m.minElementsPtr
	}
	return 0
}

func (m *AddDeviate) setMinElements(i int) {
	m.minElementsPtr = &i
}

func (m *AddDeviate) IsMinElementsSet() bool {
	return m.minElementsPtr != nil
}

// MaxElements return 0 when unbounded
func (m *AddDeviate) MaxElements() int {
	if m.maxElementsPtr != nil {
		return *m.maxElementsPtr
	}
	return 0
}

func (m *AddDeviate) setMaxElements(i int) {
	m.maxElementsPtr = &i
}

func (m *AddDeviate) IsMaxElementsSet() bool {
	return m.maxElementsPtr != nil
}

// Unique is list of fields (or compound fields) that must be unque in the
// list of items. If there is a key listed, that is implicitly unique and would
// not be listed here.
func (m *AddDeviate) Unique() [][]string {
	return m.unique
}

func (m *AddDeviate) setUnique(unique [][]string) {
	m.unique = unique
}

func (m *AddDeviate) Units() string{
	return m.units
}

func (m *AddDeviate) setUnits(u string) {
    m.units = u
}

func (m *AddDeviate) Default() []string {
	return m.defaultVals
}

func (m *AddDeviate) HasDefault() bool {
	return m.defaultVals != nil
}

func (m *AddDeviate) DefaultValue() interface{} {
	return m.Default()
}

func (m *AddDeviate) setDefaultValue(d interface{}) {
	if s, valid := d.([]string); valid {
		m.defaultVals = s
	} else {
		panic("expected []string")
	}
}

func (m *AddDeviate) addDefault(d string) {
	m.defaultVals = append(m.defaultVals, d)
}

func (m *AddDeviate) setDefault(d []string) {
	m.defaultVals = d
}

func (m *AddDeviate) clearDefault() {
    m.defaultVals = nil
}



// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *ReplaceDeviate) Parent() Meta {
	return m.parent
}

func (m *ReplaceDeviate) Extensions() []*Extension {
	return m.extensions
}

func (m *ReplaceDeviate) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *ReplaceDeviate) Config() bool {
	return *m.configPtr
}

func (m *ReplaceDeviate) setConfig(c bool) {
	m.configPtr = &c
}

func (m *ReplaceDeviate) IsConfigSet() bool {
	return m.configPtr != nil
}

func (m *ReplaceDeviate) Mandatory() bool {
	return m.mandatoryPtr != nil && *m.mandatoryPtr
}

func (m *ReplaceDeviate) setMandatory(b bool) {
	m.mandatoryPtr = &b
}

func (m *ReplaceDeviate) IsMandatorySet() bool {
	return m.mandatoryPtr != nil
}

func (m *ReplaceDeviate) MinElements() int {
	if m.minElementsPtr != nil {
		return *m.minElementsPtr
	}
	return 0
}

func (m *ReplaceDeviate) setMinElements(i int) {
	m.minElementsPtr = &i
}

func (m *ReplaceDeviate) IsMinElementsSet() bool {
	return m.minElementsPtr != nil
}

// MaxElements return 0 when unbounded
func (m *ReplaceDeviate) MaxElements() int {
	if m.maxElementsPtr != nil {
		return *m.maxElementsPtr
	}
	return 0
}

func (m *ReplaceDeviate) setMaxElements(i int) {
	m.maxElementsPtr = &i
}

func (m *ReplaceDeviate) IsMaxElementsSet() bool {
	return m.maxElementsPtr != nil
}

func (m *ReplaceDeviate) Type() *Type {
	return m.dtype
}

func (m *ReplaceDeviate) setType(t *Type) {
	m.dtype = t
}

func (m *ReplaceDeviate) Units() string{
	return m.units
}

func (m *ReplaceDeviate) setUnits(u string) {
    m.units = u
}

func (m *ReplaceDeviate) Default() []string {
	return m.defaultVals
}

func (m *ReplaceDeviate) HasDefault() bool {
	return m.defaultVals != nil
}

func (m *ReplaceDeviate) DefaultValue() interface{} {
	return m.Default()
}

func (m *ReplaceDeviate) setDefaultValue(d interface{}) {
	if s, valid := d.([]string); valid {
		m.defaultVals = s
	} else {
		panic("expected []string")
	}
}

func (m *ReplaceDeviate) addDefault(d string) {
	m.defaultVals = append(m.defaultVals, d)
}

func (m *ReplaceDeviate) setDefault(d []string) {
	m.defaultVals = d
}

func (m *ReplaceDeviate) clearDefault() {
    m.defaultVals = nil
}



// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *DeleteDeviate) Parent() Meta {
	return m.parent
}

func (m *DeleteDeviate) Extensions() []*Extension {
	return m.extensions
}

func (m *DeleteDeviate) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *DeleteDeviate) Musts() []*Must {
	return m.musts
}

func (m *DeleteDeviate) addMust(x *Must) {
	x.parent = m
    m.musts = append(m.musts, x)
}

func (m *DeleteDeviate) setMusts(x []*Must) {
    m.musts = x
}

// Unique is list of fields (or compound fields) that must be unque in the
// list of items. If there is a key listed, that is implicitly unique and would
// not be listed here.
func (m *DeleteDeviate) Unique() [][]string {
	return m.unique
}

func (m *DeleteDeviate) setUnique(unique [][]string) {
	m.unique = unique
}

func (m *DeleteDeviate) Units() string{
	return m.units
}

func (m *DeleteDeviate) setUnits(u string) {
    m.units = u
}

func (m *DeleteDeviate) Default() []string {
	return m.defaultVals
}

func (m *DeleteDeviate) HasDefault() bool {
	return m.defaultVals != nil
}

func (m *DeleteDeviate) DefaultValue() interface{} {
	return m.Default()
}

func (m *DeleteDeviate) setDefaultValue(d interface{}) {
	if s, valid := d.([]string); valid {
		m.defaultVals = s
	} else {
		panic("expected []string")
	}
}

func (m *DeleteDeviate) addDefault(d string) {
	m.defaultVals = append(m.defaultVals, d)
}

func (m *DeleteDeviate) setDefault(d []string) {
	m.defaultVals = d
}

func (m *DeleteDeviate) clearDefault() {
    m.defaultVals = nil
}



// Ident is identity of Deviation
func (m *Deviation) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Deviation) Parent() Meta {
	return m.parent
}

// Description of Deviation
func (m *Deviation) Description() string {
	return m.desc
}

func (m *Deviation) setDescription(desc string) {
	m.desc = desc
}

func (m *Deviation) Reference() string {
	return m.ref
}

func (m *Deviation) setReference(ref string) {
	m.ref = ref
}

func (m *Deviation) Extensions() []*Extension {
	return m.extensions
}

func (m *Deviation) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}




// Ident is identity of Type
func (m *Type) Ident() string {
	return m.ident
}

// Description of Type
func (m *Type) Description() string {
	return m.desc
}

func (m *Type) setDescription(desc string) {
	m.desc = desc
}

func (m *Type) Reference() string {
	return m.ref
}

func (m *Type) setReference(ref string) {
	m.ref = ref
}

func (m *Type) Extensions() []*Extension {
	return m.extensions
}

func (m *Type) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}




// Ident is identity of Identity
func (m *Identity) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Identity) Parent() Meta {
	return m.parent
}

// Description of Identity
func (m *Identity) Description() string {
	return m.desc
}

func (m *Identity) setDescription(desc string) {
	m.desc = desc
}

func (m *Identity) Reference() string {
	return m.ref
}

func (m *Identity) setReference(ref string) {
	m.ref = ref
}

func (m *Identity) Status() Status {
	return m.status
}

func (m *Identity) setStatus(status Status) {
	m.status = status
}

func (m *Identity) Extensions() []*Extension {
	return m.extensions
}

func (m *Identity) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Identity) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *Identity) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}



// Ident is identity of Feature
func (m *Feature) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Feature) Parent() Meta {
	return m.parent
}

// Description of Feature
func (m *Feature) Description() string {
	return m.desc
}

func (m *Feature) setDescription(desc string) {
	m.desc = desc
}

func (m *Feature) Reference() string {
	return m.ref
}

func (m *Feature) setReference(ref string) {
	m.ref = ref
}

func (m *Feature) Status() Status {
	return m.status
}

func (m *Feature) setStatus(status Status) {
	m.status = status
}

func (m *Feature) Extensions() []*Extension {
	return m.extensions
}

func (m *Feature) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Feature) IfFeatures() []*IfFeature {
	return m.ifs
}

func (m *Feature) addIfFeature(i *IfFeature) {
	i.parent = m
    m.ifs = append(m.ifs, i)
}



// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *IfFeature) Parent() Meta {
	return m.parent
}

func (m *IfFeature) Extensions() []*Extension {
	return m.extensions
}

func (m *IfFeature) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}






// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *When) Parent() Meta {
	return m.parent
}

// Description of When
func (m *When) Description() string {
	return m.desc
}

func (m *When) setDescription(desc string) {
	m.desc = desc
}

func (m *When) Reference() string {
	return m.ref
}

func (m *When) setReference(ref string) {
	m.ref = ref
}

func (m *When) Extensions() []*Extension {
	return m.extensions
}

func (m *When) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}




// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Must) Parent() Meta {
	return m.parent
}

// Description of Must
func (m *Must) Description() string {
	return m.desc
}

func (m *Must) setDescription(desc string) {
	m.desc = desc
}

func (m *Must) Reference() string {
	return m.ref
}

func (m *Must) setReference(ref string) {
	m.ref = ref
}

func (m *Must) Extensions() []*Extension {
	return m.extensions
}

func (m *Must) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Must) ErrorMessage() string {
	return m.errorMessage
}

func (m *Must) setErrorMessage(msg string) {
	m.errorMessage = msg
}

func (m *Must) ErrorAppTag() string {
	return m.errorAppTag
}

func (m *Must) setErrorAppTag(tag string) {
	m.errorAppTag = tag
}

func (m *Must) clone(parent Meta) interface{} {
	copy := *m
	copy.parent = parent

	return &copy
}



// Ident is identity of ExtensionDef
func (m *ExtensionDef) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *ExtensionDef) Parent() Meta {
	return m.parent
}

// Description of ExtensionDef
func (m *ExtensionDef) Description() string {
	return m.desc
}

func (m *ExtensionDef) setDescription(desc string) {
	m.desc = desc
}

func (m *ExtensionDef) Reference() string {
	return m.ref
}

func (m *ExtensionDef) setReference(ref string) {
	m.ref = ref
}

func (m *ExtensionDef) Status() Status {
	return m.status
}

func (m *ExtensionDef) setStatus(status Status) {
	m.status = status
}

func (m *ExtensionDef) Extensions() []*Extension {
	return m.extensions
}

func (m *ExtensionDef) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}




// Ident is identity of ExtensionDefArg
func (m *ExtensionDefArg) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *ExtensionDefArg) Parent() Meta {
	return m.parent
}

// Description of ExtensionDefArg
func (m *ExtensionDefArg) Description() string {
	return m.desc
}

func (m *ExtensionDefArg) setDescription(desc string) {
	m.desc = desc
}

func (m *ExtensionDefArg) Reference() string {
	return m.ref
}

func (m *ExtensionDefArg) setReference(ref string) {
	m.ref = ref
}

func (m *ExtensionDefArg) Extensions() []*Extension {
	return m.extensions
}

func (m *ExtensionDefArg) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}




// Ident is identity of Extension
func (m *Extension) Ident() string {
	return m.ident
}

// Parent is where this extension is define unless the extension is a
// secondary extension like a description and then this is the parent
// of that description
func (m *Extension) Parent() Meta {
	return m.parent
}

func (m *Extension) Extensions() []*Extension {
	return m.extensions
}

func (m *Extension) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}




// Ident is identity of Bit
func (m *Bit) Ident() string {
	return m.ident
}

// Description of Bit
func (m *Bit) Description() string {
	return m.desc
}

func (m *Bit) setDescription(desc string) {
	m.desc = desc
}

func (m *Bit) Reference() string {
	return m.ref
}

func (m *Bit) setReference(ref string) {
	m.ref = ref
}

func (m *Bit) Extensions() []*Extension {
	return m.extensions
}

func (m *Bit) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}




// Ident is identity of Enum
func (m *Enum) Ident() string {
	return m.ident
}

// Description of Enum
func (m *Enum) Description() string {
	return m.desc
}

func (m *Enum) setDescription(desc string) {
	m.desc = desc
}

func (m *Enum) Reference() string {
	return m.ref
}

func (m *Enum) setReference(ref string) {
	m.ref = ref
}

func (m *Enum) Extensions() []*Extension {
	return m.extensions
}

func (m *Enum) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}




// Description of Range
func (m *Range) Description() string {
	return m.desc
}

func (m *Range) setDescription(desc string) {
	m.desc = desc
}

func (m *Range) Reference() string {
	return m.ref
}

func (m *Range) setReference(ref string) {
	m.ref = ref
}

func (m *Range) Extensions() []*Extension {
	return m.extensions
}

func (m *Range) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Range) ErrorMessage() string {
	return m.errorMessage
}

func (m *Range) setErrorMessage(msg string) {
	m.errorMessage = msg
}

func (m *Range) ErrorAppTag() string {
	return m.errorAppTag
}

func (m *Range) setErrorAppTag(tag string) {
	m.errorAppTag = tag
}







// Description of Pattern
func (m *Pattern) Description() string {
	return m.desc
}

func (m *Pattern) setDescription(desc string) {
	m.desc = desc
}

func (m *Pattern) Reference() string {
	return m.ref
}

func (m *Pattern) setReference(ref string) {
	m.ref = ref
}

func (m *Pattern) Extensions() []*Extension {
	return m.extensions
}

func (m *Pattern) addExtension(extension *Extension) {
	m.extensions = append(m.extensions, extension)
}


func (m *Pattern) ErrorMessage() string {
	return m.errorMessage
}

func (m *Pattern) setErrorMessage(msg string) {
	m.errorMessage = msg
}

func (m *Pattern) ErrorAppTag() string {
	return m.errorAppTag
}

func (m *Pattern) setErrorAppTag(tag string) {
	m.errorAppTag = tag
}




