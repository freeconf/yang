package meta

import (
	"fmt"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/val"
)

///////////////////
// Implementation
//////////////////

// Module is top-most container of the information model. It's name
// does not appear in data model.
type Module struct {
	ident     string
	namespace string
	prefix    string
	desc      string
	contact   string
	org       string
	ref       string
	rev       []*Revision
	parent    Meta // non-null for submodules and imports
	defs      *defs
	typeDefs  map[string]*Typedef
	groupings map[string]*Grouping
	augments  []*Augment
	imports   map[string]*Import
	includes  []*Include
}

func NewModule(ident string) *Module {
	m := &Module{
		ident:     ident,
		imports:   make(map[string]*Import),
		groupings: make(map[string]*Grouping),
		typeDefs:  make(map[string]*Typedef),
		defs:      newDefs(),
	}
	return m
}

func (y *Module) Revision() *Revision {
	if len(y.rev) > 0 {
		return y.rev[0]
	}
	return nil
}

func (y *Module) RevisionHistory() []*Revision {
	return y.rev
}

func (y *Module) Includes() []*Include {
	return y.includes
}

func (y *Module) Imports() map[string]*Import {
	return y.imports
}

func (y *Module) Ident() string {
	return y.ident
}

func (y *Module) Description() string {
	return y.desc
}

func (y *Module) Reference() string {
	return y.ref
}

func (y *Module) Namespace() string {
	return y.namespace
}

func (y *Module) Prefix() string {
	return y.prefix
}

func (y *Module) Organization() string {
	return y.org
}

func (y *Module) Contact() string {
	return y.contact
}

func (y *Module) Parent() Meta {
	return y.parent
}

func (y *Module) Config() bool {
	return true
}

func (y *Module) Mandatory() bool {
	return false
}

func (y *Module) setParent(parent Meta) {
	y.parent = parent.(*Module)
}

func (y *Module) DataDefs() []DataDef {
	return y.defs.dataDefs
}

func (y *Module) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *Module) Groupings() map[string]*Grouping {
	return y.groupings
}

func (y *Module) Typedefs() map[string]*Typedef {
	return y.typeDefs
}

func (y *Module) Actions() map[string]*Rpc {
	return y.defs.actions
}

func (y *Module) Notifications() map[string]*Notification {
	return y.defs.notifications
}

func (y *Module) Augments() []*Augment {
	return y.augments
}

func (y *Module) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *Revision:
		x.parent = y
		y.rev = append(y.rev, x)
		return
	case *Include:
		x.parent = y
		y.includes = append(y.includes, x)
		return
	case *Import:
		x.parent = y
		y.imports[x.Ident()] = x
		return
	case SetPrefix:
		y.prefix = string(x)
		return
	case SetNamespace:
		y.namespace = string(x)
		return
	case SetContact:
		y.contact = string(x)
		return
	case SetOrganization:
		y.org = string(x)
		return
	case *Grouping:
		x.parent = y
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		x.parent = y
		y.typeDefs[x.Ident()] = x
		return
	case *Augment:
		x.parent = y
		y.augments = append(y.augments, x)
		return
	}
	y.defs.add(y, prop)
}

func (y *Module) clone(bool) DataDef {
	panic("not implemented")
}

func (y *Module) compile() error {
	for _, i := range y.includes {
		if err := i.compile(); err != nil {
			return err
		}
	}

	// imports independently resolve all their definitions
	if len(y.imports) > 0 {
		// imports were indexed by module name, but now that we know the
		// prefix, we need to reindex them
		byName := y.imports
		y.imports = make(map[string]*Import, len(byName))
		for _, i := range byName {
			if err := i.compile(); err != nil {
				return err
			}
			y.imports[i.Prefix()] = i
		}
	}
	return compile(y, y.defs)
}

////////////////////////////////////////////////////

type Import struct {
	prefix     string
	desc       string
	ref        string
	moduleName string
	rev        *Revision
	parent     *Module
	module     *Module
	loader     Loader
}

func NewImport(moduleName string, loader Loader) *Import {
	return &Import{
		moduleName: moduleName,
		loader:     loader,
	}
}

func (y *Import) Module() *Module {
	return y.module
}

func (y *Import) Revision() *Revision {
	return y.rev
}

func (y *Import) Prefix() string {
	return y.prefix
}

func (y *Import) Ident() string {
	return y.moduleName
}

func (y *Import) Description() string {
	return y.desc
}

func (y *Import) Reference() string {
	return y.ref
}

func (y *Import) Parent() Meta {
	return y.parent
}

func (y *Import) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *Revision:
		x.parent = y
		y.rev = x
		return
	case SetPrefix:
		y.prefix = string(x)
		return
	}
	panic(fmt.Sprintf("%T not supported in import", prop))
}

func (y *Import) compile() error {
	if y.loader == nil {
		return c2.NewErr("no module loader defined")
	}
	if y.prefix == "" {
		return c2.NewErr("prefix required on import")
	}
	var err error
	var rev string
	if y.rev != nil {
		rev = y.rev.Ident()
	}
	y.module, err = y.loader(nil, y.moduleName, rev)
	if err != nil {
		return c2.NewErr(y.moduleName + ":" + err.Error())
	}
	return y.module.compile()
}

////////////////////////////////////////////////////

type Include struct {
	subName string
	rev     *Revision
	desc    string
	ref     string
	parent  *Module
	loader  Loader
}

func NewInclude(subName string, loader Loader) *Include {
	return &Include{
		subName: subName,
		loader:  loader,
	}
}

func (y *Include) Revision() *Revision {
	return y.rev
}

func (y *Include) Ident() string {
	return y.subName
}

func (y *Include) Description() string {
	return y.desc
}

func (y *Include) Reference() string {
	return y.ref
}

func (y *Include) Parent() Meta {
	return y.parent
}

func (y *Include) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *Revision:
		x.parent = y
		y.rev = x
		return
	}
	panic(fmt.Sprintf("%T not supported in include", prop))
}

func (y *Include) compile() error {
	if y.loader == nil {
		return c2.NewErr("no module loader defined")
	}
	var err error
	var rev string
	if y.rev != nil {
		rev = y.rev.Ident()
	}
	_, err = y.loader(y.parent, y.subName, rev)
	if err != nil {
		return c2.NewErr(y.subName + ":" + err.Error())
	}
	// loader inject all definitions into this module
	// and the definitions will commpile as part of this module's
	// compile process so it's important include is run before all
	// other compile steps.  it also means all data definitions occur
	// from an include are ordered after non-included one's
	return nil
}

////////////////////////////////////////////////////

type Choice struct {
	parent     Meta
	ident      string
	desc       string
	ref        string
	conditions []Condition
	defs       *defs
	recursive  bool
}

func NewChoice(ident string) *Choice {
	return &Choice{
		ident: ident,
		defs:  newDefs(),
	}
}

func (y *Choice) Case(ident string) *ChoiceCase {
	kase, _ := y.defs.dataDefsIndex[ident].(*ChoiceCase)
	return kase
}

func (y *Choice) Ident() string {
	return y.ident
}

func (y *Choice) Description() string {
	return y.desc
}

func (y *Choice) Reference() string {
	return y.ref
}

func (y *Choice) Parent() Meta {
	return y.parent
}

func (y *Choice) setParent(parent Meta) {
	y.parent = parent
}

func (y *Choice) clone(deep bool) DataDef {
	copy := *y
	copy.defs = y.defs.clone(&copy, deep)
	return &copy
}

func (y *Choice) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *Choice) DataDefs() []DataDef {
	return y.defs.dataDefs
}

func (y *Choice) Conditions() []Condition {
	return y.conditions
}

func (y *Choice) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case Condition:
		y.conditions = append(y.conditions, x)
		return
	}
	y.defs.add(y, prop)
}

func (y *Choice) compile() error {
	return compile(y, y.defs)
}

////////////////////////////////////////////////////

type ChoiceCase struct {
	ident      string
	desc       string
	ref        string
	parent     Meta
	conditions []Condition
	defs       *defs
}

func NewChoiceCase(ident string) *ChoiceCase {
	return &ChoiceCase{
		ident: ident,
		defs:  newDefs(),
	}
}

func (y *ChoiceCase) Ident() string {
	return y.ident
}

func (y *ChoiceCase) Description() string {
	return y.desc
}

func (y *ChoiceCase) Reference() string {
	return y.ref
}

func (y *ChoiceCase) Parent() Meta {
	return y.parent
}

func (y *ChoiceCase) setParent(parent Meta) {
	y.parent = parent
}

func (y *ChoiceCase) DataDefs() []DataDef {
	return y.defs.dataDefs
}

func (y *ChoiceCase) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *ChoiceCase) clone(deep bool) DataDef {
	copy := *y
	copy.defs = y.defs.clone(&copy, deep)
	return &copy
}

func (y *ChoiceCase) Condition() []Condition {
	return y.conditions
}

func (y *ChoiceCase) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	}
	y.defs.add(y, prop)
}

func (y *ChoiceCase) compile() error {
	return compile(y, y.defs)
}

////////////////////////////////////////////////////

type Revision struct {
	parent Meta
	ident  string
	desc   string
	ref    string
}

func NewRevision(ident string) *Revision {
	return &Revision{
		ident: ident,
	}
}

func (y *Revision) Ident() string {
	return y.ident
}

func (y *Revision) Description() string {
	return y.desc
}

func (y *Revision) Reference() string {
	return y.ref
}

func (y *Revision) Parent() Meta {
	return y.parent
}

func (y *Revision) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	}
	panic(fmt.Sprintf("%T not supported in revision", prop))
}

func (y *Revision) compile() error {
	return nil
}

////////////////////////////////////////////////////

type Container struct {
	ident      string
	desc       string
	ref        string
	typeDefs   map[string]*Typedef
	groupings  map[string]*Grouping
	parent     Meta
	configPtr  *bool
	mandatory  bool
	conditions []Condition
	defs       *defs
	recursive  bool
}

func NewContainer(ident string) *Container {
	return &Container{
		ident:     ident,
		defs:      newDefs(),
		groupings: make(map[string]*Grouping),
		typeDefs:  make(map[string]*Typedef),
	}
}

func (y *Container) Ident() string {
	return y.ident
}

func (y *Container) Description() string {
	return y.desc
}

func (y *Container) Reference() string {
	return y.ref
}

func (y *Container) Parent() Meta {
	return y.parent
}

func (y *Container) setParent(parent Meta) {
	y.parent = parent
}

func (y *Container) DataDefs() []DataDef {
	return y.defs.dataDefs
}

func (y *Container) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *Container) Groupings() map[string]*Grouping {
	return y.groupings
}

func (y *Container) Typedefs() map[string]*Typedef {
	return y.typeDefs
}

func (y *Container) Actions() map[string]*Rpc {
	return y.defs.actions
}

func (y *Container) Notifications() map[string]*Notification {
	return y.defs.notifications
}

func (y *Container) Config() bool {
	return *y.configPtr
}

func (y *Container) Mandatory() bool {
	return y.mandatory
}

func (y *Container) Conditions() []Condition {
	return y.conditions
}

func (y *Container) clone(deep bool) DataDef {
	copy := *y
	if deep && !y.recursive {
		copy.defs = y.defs.clone(&copy, deep)
	} else {
		copy.recursive = true
	}
	return &copy
}

func (y *Container) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case SetConfig:
		b := bool(x)
		y.configPtr = &b
		return
	case SetMandatory:
		y.mandatory = bool(x)
		return
	case *Grouping:
		x.parent = y
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		x.parent = y
		y.typeDefs[x.Ident()] = x
		return
	case Condition:
		x.setParent(y)
		y.conditions = append(y.conditions, x)
		return
	}
	y.defs.add(y, prop)
}

func (y *Container) compile() error {
	if y.configPtr == nil {
		b := inheritConfig(y.parent)
		y.configPtr = &b
	}

	return compile(y, y.defs)
}

func inheritConfig(m Meta) bool {
	if x, ok := m.(HasDetails); ok {
		return x.Config()
	}
	return true
}

////////////////////////////////////////////////////

type List struct {
	parent       Meta
	ident        string
	desc         string
	ref          string
	typeDefs     map[string]*Typedef
	groupings    map[string]*Grouping
	key          []string
	keyMeta      []HasDataType
	conditions   []Condition
	configPtr    *bool
	mandatory    bool
	minElements  int
	maxElements  int
	unboundedPtr *bool
	defs         *defs
	recursive    bool
}

func NewList(ident string) *List {
	return &List{
		ident:     ident,
		defs:      newDefs(),
		groupings: make(map[string]*Grouping),
		typeDefs:  make(map[string]*Typedef),
	}
}

func (y *List) KeyMeta() (keyMeta []HasDataType) {
	return y.keyMeta
}

func (y *List) Ident() string {
	return y.ident
}

func (y *List) Description() string {
	return y.desc
}

func (y *List) Reference() string {
	return y.ref
}

func (y *List) Parent() Meta {
	return y.parent
}

func (y *List) setParent(parent Meta) {
	y.parent = parent
}

func (y *List) Config() bool {
	return *y.configPtr
}

func (y *List) Mandatory() bool {
	return y.mandatory
}

func (y *List) MaxElements() int {
	return y.maxElements
}

func (y *List) MinElements() int {
	return y.minElements
}

func (y *List) Unbounded() bool {
	return *y.unboundedPtr
}

func (y *List) Conditions() []Condition {
	return y.conditions
}

func (y *List) DataDefs() []DataDef {
	return y.defs.dataDefs
}

func (y *List) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *List) Groupings() map[string]*Grouping {
	return y.groupings
}

func (y *List) Typedefs() map[string]*Typedef {
	return y.typeDefs
}

func (y *List) Actions() map[string]*Rpc {
	return y.defs.actions
}

func (y *List) Notifications() map[string]*Notification {
	return y.defs.notifications
}

func (y *List) clone(deep bool) DataDef {
	copy := *y
	if deep && !y.recursive {
		copy.defs = y.defs.clone(&copy, deep)
	} else {
		copy.recursive = true
	}
	return &copy
}

func (y *List) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case SetConfig:
		b := bool(x)
		y.configPtr = &b
		return
	case SetMandatory:
		y.mandatory = bool(x)
		return
	case SetUnbounded:
		b := bool(x)
		y.unboundedPtr = &b
		return
	case SetEncodedLength:
		x.decode(y)
		return
	case SetMaxElements:
		y.maxElements = int(x)
		return
	case SetMinElements:
		y.minElements = int(x)
		return
	case SetKey:
		y.key = x.decode()
		return
	case *Grouping:
		x.parent = y
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		x.parent = y
		y.typeDefs[x.Ident()] = x
		return
	case Condition:
		x.setParent(y)
		y.conditions = append(y.conditions, x)
		return
	}
	y.defs.add(y, prop)
}

func (y *List) compile() error {
	if y.configPtr == nil {
		b := inheritConfig(y.parent)
		y.configPtr = &b
	}

	if y.unboundedPtr == nil {
		b := (y.maxElements == 0)
		y.unboundedPtr = &b
	}

	c2.DebugLog(true)
	if err := compile(y, y.defs); err != nil {
		return err
	}
	c2.DebugLog(false)

	y.keyMeta = make([]HasDataType, len(y.key))
	for i, keyIdent := range y.key {
		// relies on res
		km, valid := y.defs.dataDefsIndex[keyIdent]
		if !valid {
			return c2.NewErr(keyIdent + " key not found for " + GetPath(y))
		}
		y.keyMeta[i], valid = km.(HasDataType)
		if !valid {
			return c2.NewErr(keyIdent + " expected key with data type")
		}
	}

	return nil
}

////////////////////////////////////////////////////

type Leaf struct {
	parent     Meta
	ident      string
	desc       string
	ref        string
	configPtr  *bool
	mandatory  bool
	defaultVal interface{}
	dtype      *DataType
	conditions []Condition
}

func NewLeaf(ident string) *Leaf {
	l := &Leaf{
		ident: ident,
	}
	return l
}

func NewLeafWithType(ident string, f val.Format) *Leaf {
	l := NewLeaf(ident)
	l.dtype = NewDataType(f.String())
	return l
}

func (y *Leaf) Ident() string {
	return y.ident
}

func (y *Leaf) Description() string {
	return y.desc
}

func (y *Leaf) Reference() string {
	return y.ref
}

func (y *Leaf) Parent() Meta {
	return y.parent
}

func (y *Leaf) setParent(parent Meta) {
	y.parent = parent
}

func (y *Leaf) DataType() *DataType {
	return y.dtype
}

func (y *Leaf) Conditions() []Condition {
	return y.conditions
}

func (y *Leaf) Config() bool {
	return *y.configPtr
}

func (y *Leaf) Mandatory() bool {
	return y.mandatory
}

func (y *Leaf) HasDefault() bool {
	return y.defaultVal != nil
}

func (y *Leaf) Default() interface{} {
	return y.defaultVal
}

func (y *Leaf) clone(deep bool) DataDef {
	copy := *y
	return &copy
}

func (y *Leaf) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case SetConfig:
		b := bool(x)
		y.configPtr = &b
		return
	case SetMandatory:
		y.mandatory = bool(x)
		return
	case SetDefault:
		y.defaultVal = x.Value
		return
	case *DataType:
		x.parent = y
		y.dtype = x
		return
	case Condition:
		x.setParent(y)
		y.conditions = append(y.conditions, x)
		return
	}
	panic(fmt.Sprintf("%T not supported in leaf", prop))
}

func (y *Leaf) compile() error {
	if y.configPtr == nil {
		b := inheritConfig(y.parent)
		y.configPtr = &b
	}
	if err := compile(y, nil); err != nil {
		return err
	}
	if y.defaultVal == nil {
		y.defaultVal = y.dtype.defaultVal
	}
	return nil
}

////////////////////////////////////////////////////

type LeafList struct {
	ident        string
	parent       Meta
	desc         string
	ref          string
	configPtr    *bool
	mandatory    bool
	dtype        *DataType
	minElements  int
	maxElements  int
	unboundedPtr *bool
	defaults     []interface{}
	conditions   []Condition
}

func NewLeafList(ident string) *LeafList {
	l := &LeafList{
		ident: ident,
	}
	return l
}

func NewLeafListWithType(ident string, f val.Format) *LeafList {
	l := NewLeafList(ident)
	l.dtype = NewDataType(f.String())
	return l
}

func (y *LeafList) Ident() string {
	return y.ident
}

func (y *LeafList) Description() string {
	return y.desc
}

func (y *LeafList) Reference() string {
	return y.ref
}

func (y *LeafList) Parent() Meta {
	return y.parent
}

func (y *LeafList) setParent(parent Meta) {
	y.parent = parent
}

func (y *LeafList) DataType() *DataType {
	return y.dtype
}

func (y *LeafList) Config() bool {
	return *y.configPtr
}

func (y *LeafList) Mandatory() bool {
	return y.mandatory
}

func (y *LeafList) MaxElements() int {
	return y.maxElements
}

func (y *LeafList) MinElements() int {
	return y.minElements
}

func (y *LeafList) Unbounded() bool {
	if y.unboundedPtr != nil {
		return *y.unboundedPtr
	}
	return y.maxElements == 0
}

func (y *LeafList) HasDefault() bool {
	return y.defaults != nil
}

func (y *LeafList) Default() interface{} {
	return y.defaults
}

func (y *LeafList) Conditions() []Condition {
	return y.conditions
}

func (y *LeafList) clone(bool) DataDef {
	copy := *y
	return &copy
}

func (y *LeafList) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case SetConfig:
		b := bool(x)
		y.configPtr = &b
		return
	case SetMandatory:
		y.mandatory = bool(x)
		return
	case SetUnbounded:
		b := bool(x)
		y.unboundedPtr = &b
		return
	case SetEncodedLength:
		x.decode(y)
		return
	case SetMaxElements:
		y.maxElements = int(x)
		return
	case SetMinElements:
		y.minElements = int(x)
		return
	case SetDefault:
		y.defaults = append(y.defaults, x.Value)
	case *DataType:
		x.parent = y
		y.dtype = x
		return
	case Condition:
		x.setParent(y)
		y.conditions = append(y.conditions, x)
		return
	}
	panic(fmt.Sprintf("%T not supported in leaf-list", prop))
}

func (y *LeafList) compile() error {
	if y.configPtr == nil {
		b := inheritConfig(y.parent)
		y.configPtr = &b
	}

	return compile(y, nil)
}

////////////////////////////////////////////////////

type Any struct {
	ident      string
	desc       string
	ref        string
	parent     Meta
	configPtr  *bool
	mandatory  bool
	dtype      *DataType
	conditions []Condition
}

func NewAny(ident string) *Any {
	any := &Any{
		ident: ident,
		dtype: NewDataType("any"),
	}
	any.dtype.parent = any
	return any
}

func (y *Any) Ident() string {
	return y.ident
}

func (y *Any) Description() string {
	return y.desc
}

func (y *Any) Reference() string {
	return y.ref
}

func (y *Any) Parent() Meta {
	return y.parent
}

func (y *Any) setParent(parent Meta) {
	y.parent = parent
}

func (y *Any) DataType() *DataType {
	return y.dtype
}

func (y *Any) Config() bool {
	return *y.configPtr
}

func (y *Any) Mandatory() bool {
	return y.mandatory
}

func (y *Any) HasDefault() bool {
	return false
}

func (y *Any) Default() interface{} {
	panic("anydata cannot have default value")
}

func (y *Any) Conditions() []Condition {
	return y.conditions
}

func (y *Any) clone(bool) DataDef {
	copy := *y
	return &copy
}

func (y *Any) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case SetConfig:
		b := bool(x)
		y.configPtr = &b
		return
	case SetMandatory:
		y.mandatory = bool(x)
		return
	case Condition:
		x.setParent(y)
		y.conditions = append(y.conditions, x)
		return
	}
	panic(fmt.Sprintf("%T not supported in any", prop))
}

func (y *Any) compile() error {
	if y.configPtr == nil {
		b := inheritConfig(y.parent)
		y.configPtr = &b
	}
	return compile(y, nil)
}

////////////////////////////////////////////////////

/**
  RFC7950 Sec 7.12 The "grouping" Statement

  Identifiers appearing inside
  the grouping are resolved relative to the scope in which the grouping
  is defined, not where it is used.  Prefix mappings, type names,
  grouping names, and extension usage are evaluated in the hierarchy
  where the "grouping" statement appears.  For extensions, this means
  that extensions defined as direct children to a "grouping" statement
  are applied to the grouping itself.
*/
type Grouping struct {
	ident     string
	parent    Meta
	desc      string
	ref       string
	typeDefs  map[string]*Typedef
	groupings map[string]*Grouping

	compiling   bool
	compiled    bool
	postCompile []func() error

	defs *defs
	// see RFC7950 Sec 14
	// no details (config, mandatory)
	// no conditions
}

func NewGrouping(ident string) *Grouping {
	return &Grouping{
		ident:     ident,
		defs:      newDefs(),
		groupings: make(map[string]*Grouping),
		typeDefs:  make(map[string]*Typedef),
	}
}

func (y *Grouping) Ident() string {
	return y.ident
}

func (y *Grouping) Description() string {
	return y.desc
}

func (y *Grouping) Reference() string {
	return y.ref
}

func (y *Grouping) Parent() Meta {
	return y.parent
}

func (y *Grouping) setParent(parent Meta) {
	y.parent = parent
}

func (y *Grouping) Groupings() map[string]*Grouping {
	return y.groupings
}

func (y *Grouping) Typedefs() map[string]*Typedef {
	return y.typeDefs
}

func (y *Grouping) DataDefs() []DataDef {
	return y.defs.dataDefs
}

func (y *Grouping) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *Grouping) Actions() map[string]*Rpc {
	return y.defs.actions
}

func (y *Grouping) Notifications() map[string]*Notification {
	return y.defs.notifications
}

func (y *Grouping) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *Grouping:
		x.parent = y
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		x.parent = y
		y.typeDefs[x.Ident()] = x
		return
	}
	y.defs.add(y, prop)
}

func (y *Grouping) copyInto(parent Meta) error {
	for _, x := range y.defs.notifications {
		if err := Set(parent, x); err != nil {
			return err
		}
	}
	for _, x := range y.defs.actions {
		if err := Set(parent, x); err != nil {
			return err
		}
	}
	for _, x := range y.defs.dataDefs {
		if err := Set(parent, x); err != nil {
			return err
		}
	}
	return nil
}

func (y *Grouping) clone(deep bool) *Grouping {
	copy := *y
	copy.defs = y.defs.clone(&copy, deep)
	return &copy
}

func (y *Grouping) compile() error {
	if y.compiled {
		return nil
	}
	y.compiling = true
	if err := compile(y, y.defs); err != nil {
		return err
	}
	for _, f := range y.postCompile {
		if err := f(); err != nil {
			return err
		}
	}
	y.compiling = false
	y.compiled = true
	return nil
}

func (y *Grouping) onPostCompile(f func() error) {
	y.postCompile = append(y.postCompile, f)
}

////////////////////////////////////////////////////

type Uses struct {
	ident     string
	desc      string
	ref       string
	parent    Meta
	refines   []*Refine
	condition []Condition
}

func NewUses(ident string) *Uses {
	return &Uses{
		ident: ident,
	}
}

func (y *Uses) Refinements() []*Refine {
	return y.refines
}

func (y *Uses) Ident() string {
	return y.ident
}

func (y *Uses) Description() string {
	return y.desc
}

func (y *Uses) Reference() string {
	return y.ref
}

func (y *Uses) Parent() Meta {
	return y.parent
}

func (y *Uses) setParent(parent Meta) {
	y.parent = parent
}

func (y *Uses) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *Refine:
		x.parent = y
		y.refines = append(y.refines, x)
		return
	}
	panic(fmt.Sprintf("%T not supported in uses", prop))
}

func (y *Uses) compile() error {
	target := y.findScopedTarget()
	if target == nil {
		return c2.NewErr(y.ident + " group not found")
	}
	if target.compiling {
		// circular reference, we copy in group content
		if len(y.refines) > 0 {
			return c2.NewErr(y.ident + " cannot refine on a circular group")
		}
		target.onPostCompile(func() error {
			copy := target.clone(false)
			return copy.copyInto(y.parent)
		})
	} else {
		// harmless is already compiled
		if err := target.compile(); err != nil {
			return err
		}
		// clone then copy into parent
		copy := target.clone(true)
		for _, r := range y.refines {
			r.refine(copy)
		}
		if err := copy.copyInto(y.parent); err != nil {
			return err
		}
	}
	return nil
}

func (y *Uses) findScopedTarget() *Grouping {
	// lazy load grouping
	if xMod, xIdent, err := externalModule(y, y.ident); err != nil {
		return nil
	} else if xMod != nil {
		return xMod.Groupings()[xIdent]
	} else {
		p := y.Parent()
		for p != nil {
			if hasGroups, ok := p.(HasGroupings); ok {
				if g, found := hasGroups.Groupings()[y.ident]; found {
					return g
				}
			}
			p = p.Parent()
		}
	}
	return nil
}

/////////////////////

type Refine struct {
	ident          string
	desc           string
	ref            string
	parent         *Uses
	configPtr      *bool
	mandatoryPtr   *bool
	maxElementsPtr *int
	minElementsPtr *int
	unboundedPtr   *bool
	defaultVal     interface{}
}

func NewRefine(path string) *Refine {
	return &Refine{
		ident: path,
	}
}

func (y *Refine) refine(g *Grouping) error {
	target := Find(g, y.ident)
	if target == nil {
		return c2.NewErr(y.ident + " refine target not found in " + g.Ident())
	}
	if y.desc != "" {
		if err := Set(target, SetDescription(y.desc)); err != nil {
			return err
		}
	}
	if y.ref != "" {
		if err := Set(target, SetReference(y.ref)); err != nil {
			return err
		}
	}
	if y.defaultVal != nil {
		if err := Set(target, SetDefault{Value: y.defaultVal}); err != nil {
			return err
		}
	}
	if y.configPtr != nil {
		if err := Set(target, SetConfig(*y.configPtr)); err != nil {
			return err
		}
	}
	if y.mandatoryPtr != nil {
		if err := Set(target, SetMandatory(*y.mandatoryPtr)); err != nil {
			return err
		}
	}
	if y.maxElementsPtr != nil {
		if err := Set(target, SetMaxElements(*y.maxElementsPtr)); err != nil {
			return err
		}
	}
	if y.minElementsPtr != nil {
		if err := Set(target, SetMinElements(*y.minElementsPtr)); err != nil {
			return err
		}
	}
	if y.unboundedPtr != nil {
		if err := Set(target, SetUnbounded(*y.unboundedPtr)); err != nil {
			return err
		}
	}
	return nil
}

func (y *Refine) Ident() string {
	return y.ident
}

func (y *Refine) Description() string {
	return y.desc
}

func (y *Refine) Reference() string {
	return y.ref
}

func (y *Refine) Parent() Meta {
	return y.parent
}

func (y *Refine) setParent(parent Meta) {
	y.parent = parent.(*Uses)
}

func (y *Refine) ConfigPtr() *bool {
	return y.configPtr
}

func (y *Refine) MandatoryPtr() *bool {
	return y.mandatoryPtr
}

func (y *Refine) DefaultPtr() interface{} {
	return y.defaultVal
}

func (y *Refine) MaxElementsPtr() *int {
	return y.maxElementsPtr
}

func (y *Refine) MinElementsPtr() *int {
	return y.minElementsPtr
}

func (y *Refine) UnboundedPtr() *bool {
	return y.unboundedPtr
}

func (y *Refine) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case SetConfig:
		b := bool(x)
		y.configPtr = &b
		return
	case SetMandatory:
		b := bool(x)
		y.mandatoryPtr = &b
		return
	case SetUnbounded:
		b := bool(x)
		y.unboundedPtr = &b
		return
	case SetEncodedLength:
		x.decode(y)
		return
	case SetMaxElements:
		i := int(x)
		y.maxElementsPtr = &i
		return
	case SetMinElements:
		i := int(x)
		y.minElementsPtr = &i
		return
	case SetDefault:
		y.defaultVal = x.Value
		return
	}
	panic(fmt.Sprintf("%T not supported in refine", prop))
}

func (y *Refine) compile() error {
	return nil
}

////////////////////////////////////////////////////

type RpcInput struct {
	parent    *Rpc
	desc      string
	ref       string
	typeDefs  map[string]*Typedef
	groupings map[string]*Grouping
	defs      *defs
}

func NewRpcInput() *RpcInput {
	return &RpcInput{
		defs:      newDefs(),
		groupings: make(map[string]*Grouping),
		typeDefs:  make(map[string]*Typedef),
	}
}

func (y *RpcInput) Ident() string {
	return "input"
}

func (y *RpcInput) Description() string {
	return y.desc
}

func (y *RpcInput) Reference() string {
	return y.ref
}

func (y *RpcInput) Parent() Meta {
	return y.parent
}

func (y *RpcInput) setParent(parent Meta) {
	y.parent = (parent).(*Rpc)
}

func (y *RpcInput) Groupings() map[string]*Grouping {
	return y.groupings
}

func (y *RpcInput) Typedefs() map[string]*Typedef {
	return y.typeDefs
}

func (y *RpcInput) DataDefs() []DataDef {
	return y.defs.dataDefs
}

func (y *RpcInput) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *RpcInput) clone(deep bool) *RpcInput {
	copy := *y
	copy.defs = y.defs.clone(&copy, deep)
	return &copy
}

func (y *RpcInput) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *Grouping:
		x.parent = y
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		x.parent = y
		y.typeDefs[x.Ident()] = x
		return
	}
	y.defs.add(y, prop)
}

func (y *RpcInput) compile() error {
	return compile(y, y.defs)
}

////////////////////////////////////////////////////

type RpcOutput struct {
	parent    *Rpc
	desc      string
	ref       string
	typeDefs  map[string]*Typedef
	groupings map[string]*Grouping
	defs      *defs
}

func NewRpcOutput() *RpcOutput {
	return &RpcOutput{
		defs:      newDefs(),
		groupings: make(map[string]*Grouping),
		typeDefs:  make(map[string]*Typedef),
	}
}

func (y *RpcOutput) Ident() string {
	return "output"
}

func (y *RpcOutput) Description() string {
	return y.desc
}

func (y *RpcOutput) Reference() string {
	return y.ref
}

func (y *RpcOutput) Parent() Meta {
	return y.parent
}

func (y *RpcOutput) setParent(parent Meta) {
	y.parent = (parent).(*Rpc)
}

func (y *RpcOutput) Groupings() map[string]*Grouping {
	return y.groupings
}

func (y *RpcOutput) Typedefs() map[string]*Typedef {
	return y.typeDefs
}

func (y *RpcOutput) DataDefs() []DataDef {
	return y.defs.dataDefs
}

func (y *RpcOutput) Config() bool {
	return false
}

func (y *RpcOutput) Mandatory() bool {
	return false
}

func (y *RpcOutput) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *RpcOutput) clone(deep bool) *RpcOutput {
	copy := *y
	copy.defs = y.defs.clone(&copy, deep)
	return &copy
}

func (y *RpcOutput) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *Grouping:
		x.parent = y
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		x.parent = y
		y.typeDefs[x.Ident()] = x
		return
	}
	y.defs.add(y, prop)
}

func (y *RpcOutput) compile() error {
	return compile(y, y.defs)
}

////////////////////////////////////////////////////

type Rpc struct {
	ident     string
	parent    Meta
	desc      string
	ref       string
	typeDefs  map[string]*Typedef
	groupings map[string]*Grouping
	input     *RpcInput
	output    *RpcOutput
}

func NewRpc(ident string) *Rpc {
	return &Rpc{
		ident:     ident,
		groupings: make(map[string]*Grouping),
		typeDefs:  make(map[string]*Typedef),
	}
}

func (y *Rpc) Input() *RpcInput {
	return y.input
}

func (y *Rpc) Output() *RpcOutput {
	return y.output
}

func (y *Rpc) Ident() string {
	return y.ident
}

func (y *Rpc) Description() string {
	return y.desc
}

func (y *Rpc) Reference() string {
	return y.ref
}

func (y *Rpc) Parent() Meta {
	return y.parent
}

func (y *Rpc) setParent(parent Meta) {
	y.parent = parent
}

func (y *Rpc) Groupings() map[string]*Grouping {
	return y.groupings
}

func (y *Rpc) Typedefs() map[string]*Typedef {
	return y.typeDefs
}

func (y *Rpc) clone(deep bool) *Rpc {
	copy := *y
	copy.input = y.input.clone(deep)
	copy.input.parent = &copy
	copy.output = y.output.clone(deep)
	copy.output.parent = &copy
	return &copy
}

func (y *Rpc) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *RpcInput:
		x.parent = y
		y.input = x
		return
	case *RpcOutput:
		x.parent = y
		y.output = x
		return
	case *Grouping:
		x.parent = y
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		x.parent = y
		y.typeDefs[x.Ident()] = x
		return
	}
	panic(fmt.Sprintf("%T not supported in action", prop))
}

func (y *Rpc) compile() error {
	if err := compile(y, nil); err != nil {
		return err
	}

	if y.input != nil {
		if err := y.input.compile(); err != nil {
			return err
		}
	}
	if y.output != nil {
		if err := y.output.compile(); err != nil {
			return err
		}
	}

	return nil
}

////////////////////////////////////////////////////

type Notification struct {
	ident     string
	parent    Meta
	desc      string
	ref       string
	typeDefs  map[string]*Typedef
	groupings map[string]*Grouping
	defs      *defs
}

func NewNotification(ident string) *Notification {
	return &Notification{
		ident:     ident,
		typeDefs:  make(map[string]*Typedef),
		groupings: make(map[string]*Grouping),
		defs:      newDefs(),
	}
}

func (y *Notification) Ident() string {
	return y.ident
}

func (y *Notification) Description() string {
	return y.desc
}

func (y *Notification) Reference() string {
	return y.ref
}

func (y *Notification) Parent() Meta {
	return y.parent
}

func (y *Notification) setParent(parent Meta) {
	y.parent = parent
}

func (y *Notification) DataDefs() []DataDef {
	return y.defs.dataDefs
}

func (y *Notification) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *Notification) Groupings() map[string]*Grouping {
	return y.groupings
}

func (y *Notification) Typedefs() map[string]*Typedef {
	return y.typeDefs
}

func (y *Notification) clone(deep bool) *Notification {
	copy := *y
	copy.defs = y.defs.clone(&copy, deep)
	return &copy
}

func (y *Notification) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *Grouping:
		x.parent = y
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		x.parent = y
		y.typeDefs[x.Ident()] = x
		return
	}
	y.defs.add(y, prop)
}

func (y *Notification) compile() error {
	return compile(y, y.defs)
}

////////////////////////////////////////////////////

type Typedef struct {
	ident      string
	parent     Meta
	desc       string
	ref        string
	defaultVal interface{}
	dtype      *DataType
}

func NewTypedef(ident string) *Typedef {
	return &Typedef{ident: ident}
}

func (y *Typedef) Ident() string {
	return y.ident
}

func (y *Typedef) Description() string {
	return y.desc
}

func (y *Typedef) Reference() string {
	return y.ref
}

func (y *Typedef) Parent() Meta {
	return y.parent
}

func (y *Typedef) HasDefault() bool {
	return y.defaultVal != nil
}

func (y *Typedef) Default() interface{} {
	return y.defaultVal
}

func (y *Typedef) setParent(parent Meta) {
	y.parent = parent
}

func (y *Typedef) DataType() *DataType {
	return y.dtype
}

func (y *Typedef) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *DataType:
		x.parent = y
		y.dtype = x
		return
	case SetDefault:
		y.defaultVal = x.Value
		return
	}
	panic(fmt.Sprintf("%T not supported in typedef", prop))
}

func (y *Typedef) compile() error {
	if y.dtype == nil {
		c2.NewErr(y.ident + " type required")
	}

	return compile(y, nil)
}

////////////////////////////////////////////////

type Augment struct {
	ident      string
	parent     Meta
	desc       string
	ref        string
	defs       *defs
	conditions []Condition
}

func NewAugment(path string) *Augment {
	return &Augment{
		ident: path,
		defs:  newDefs(),
	}
}

func (y *Augment) Ident() string {
	return y.ident
}

func (y *Augment) Description() string {
	return y.desc
}

func (y *Augment) Reference() string {
	return y.ref
}

func (y *Augment) Parent() Meta {
	return y.parent
}

func (y *Augment) setParent(parent Meta) {
	y.parent = parent
}

func (y *Augment) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case Condition:
		y.conditions = append(y.conditions, x)
		return
	}
	y.defs.add(y, prop)
}

func (y *Augment) compile() error {
	if err := compile(y, y.defs); err != nil {
		return err
	}

	// RFC7950 Sec 7.17
	// "The target node MUST be either a container, list, choice, case, input,
	//   output, or notification node."
	target := Find(y.parent.(HasDefinitions), y.ident)
	if target == nil {
		return c2.NewErr("augment target is not found " + y.ident)
	}

	// expand
	for _, x := range y.defs.actions {
		if err := Set(target, x); err != nil {
			return err
		}
	}
	for _, x := range y.defs.notifications {
		if err := Set(target, x); err != nil {
			return err
		}
	}
	for _, x := range y.defs.dataDefs {
		if err := Set(target, x); err != nil {
			return err
		}
	}
	return nil
}

////////////////////////////////////////////////////

type DataType struct {
	parent    HasDataType
	typeIdent string
	desc      string
	ref       string
	format    val.Format
	rangeVal  string
	enum      val.EnumList
	// because minLength of 0 is legit value, we store pointer so we know if it's
	// been explicitly set
	minLengthPtr *int
	maxLength    int
	path         string
	pattern      string
	defaultVal   interface{}
	delegate     *DataType
	/*
		FractionDigits
		Bit
		Base
		RequireInstance
		Type?!  subtype?
	*/
}

func NewDataType(typeIdent string) *DataType {
	return &DataType{
		typeIdent: typeIdent,
	}
}

func (y *DataType) TypeIdent() string {
	return y.typeIdent
}

func (y *DataType) Parent() Meta {
	return y.parent
}

func (y *DataType) Description() string {
	return y.desc
}

func (y *DataType) Reference() string {
	return y.ref
}

func (y *DataType) setParent(parent Meta) {
	y.parent = parent.(HasDataType)
}

func (y *DataType) MaxLength() int {
	return y.maxLength
}

func (y *DataType) MinLength() int {
	if y.minLengthPtr != nil {
		return *y.minLengthPtr
	}
	return 0
}

func (y *DataType) Pattern() string {
	return y.pattern
}

// TODO: This has to expand to be slice of min/max numbers
func (y *DataType) Range() string {
	return y.rangeVal
}

func (y *DataType) Format() val.Format {
	return y.format
}

func (y *DataType) Path() string {
	return y.path
}

func (y *DataType) Enum() val.EnumList {
	return y.enum
}

func (y *DataType) HasDefault() bool {
	return y.defaultVal != nil
}

func (y *DataType) DefaultValue() interface{} {
	return y.defaultVal
}

// Resolve is the effective datatype if this type points to a different
// dataType, which is the case for leafRefs.  Otherwise this just returns
// itself
func (y *DataType) Resolve() *DataType {
	if y.delegate == nil {

		// TODO: this should be on "post compile" phase when all paths are
		// resolved but before usage.  Putting this here for now.
		if y.format == val.FmtLeafRef || y.format == val.FmtLeafRefList {
			if y.path == "" {
				panic(y.typeIdent + " path is required")
			}
			// parent is a leaf, so start with parent's parent which is a container-ish
			c2.DebugLog(true)
			resolvedMeta := Find(y.parent.Parent().(HasDefinitions), y.path)
			c2.DebugLog(false)
			if resolvedMeta == nil {
				panic(y.typeIdent + " could not resolve 'path' on leafref " + y.path)
			}
			y.delegate = resolvedMeta.(HasDataType).DataType()
		} else {
			y.delegate = y
		}
	}
	return y.delegate
}

func (y *DataType) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case SetRange:
		y.rangeVal = string(x)
		return
	case SetMaxLength:
		y.maxLength = int(x)
		return
	case SetMinLength:
		i := int(x)
		y.minLengthPtr = &i
		return
	case SetEncodedLength:
		x.decode(y)
		return
	case SetPattern:
		y.pattern = string(x)
		return
	case SetPath:
		y.path = string(x)
		return
	case val.Enum:
		y.enum = append(y.enum, x)
		return
	case SetEnumLabel:
		y.enum = y.enum.Add(string(x))
		return
	}
	panic(fmt.Sprintf("%T not supported in type", prop))
}

func (base *DataType) mixin(derived *DataType) {
	if base.pattern != "" && derived.pattern == "" {
		derived.pattern = base.pattern
	}
	if base.path != "" && derived.path == "" {
		derived.path = base.path
	}
	if base.minLengthPtr != nil && derived.minLengthPtr != nil {
		derived.minLengthPtr = base.minLengthPtr
	}
	if base.maxLength != 0 && derived.maxLength == 0 {
		derived.maxLength = base.maxLength
	}
	derived.format = base.format
}

func (y *DataType) compile() error {
	var hasTypedef bool
	y.format, hasTypedef = val.TypeAsFormat(y.typeIdent)
	var tdef *Typedef
	if !hasTypedef {
		if tdef = y.findScopedTypedef(y.typeIdent); tdef == nil {
			return c2.NewErr(y.typeIdent + " not found")
		}
		tdef.dtype.mixin(y)
		y.defaultVal = tdef.defaultVal
	}

	if _, isList := y.parent.(*LeafList); isList && !y.format.IsList() {
		y.format = val.Format(int(y.format) + 1024)
	}

	return nil
}

// TODO: support namesapce
func (y *DataType) findScopedTypedef(ident string) *Typedef {
	p := y.Parent()
	for p != nil {
		if ptd, ok := p.(HasTypedefs); ok {
			if td, found := ptd.Typedefs()[ident]; found {
				return td
			}
		}
		p = p.Parent()
	}
	return nil
}
