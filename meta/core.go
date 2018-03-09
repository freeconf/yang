package meta

import (
	"errors"
	"fmt"
	"strings"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/val"
)

///////////////////
// Implementation
//////////////////

// Module is top-most container of the information model. It's name
// does not appear in data model.
type Module struct {
	ident      string
	namespace  string
	prefix     string
	desc       string
	contact    string
	org        string
	ref        string
	ver        string
	rev        []*Revision
	parent     Meta // non-null for submodules and imports
	defs       *defs
	typeDefs   map[string]*Typedef
	groupings  map[string]*Grouping
	augments   []*Augment
	imports    map[string]*Import
	includes   []*Include
	identities map[string]*Identity
	features   map[string]*Feature
	extensions map[string]*Extension
	featureSet FeatureSet
}

func NewModule(ident string, featureSet FeatureSet) *Module {
	m := &Module{
		ident:      ident,
		ver:        "1",
		imports:    make(map[string]*Import),
		groupings:  make(map[string]*Grouping),
		typeDefs:   make(map[string]*Typedef),
		identities: make(map[string]*Identity),
		features:   make(map[string]*Feature),
		extensions: make(map[string]*Extension),
		defs:       newDefs(),
		featureSet: featureSet,
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

func (y *Module) Identities() map[string]*Identity {
	return y.identities
}

func (y *Module) Features() map[string]*Feature {
	return y.features
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

func (y *Module) Version() string {
	return y.ver
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

func (y *Module) DataDefs() []Definition {
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

func (y *Module) FeatureSet() FeatureSet {
	return y.featureSet
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
		y.rev = append(y.rev, x)
		return
	case *Include:
		y.includes = append(y.includes, x)
		return
	case *Import:
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
	case SetYangVersion:
		y.ver = string(x)
		return
	case *Grouping:
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		y.typeDefs[x.Ident()] = x
		return
	case *Augment:
		y.augments = append(y.augments, x)
		return
	case *Identity:
		y.identities[x.Ident()] = x
		return
	case *Feature:
		y.features[x.Ident()] = x
		return
	case *Extension:
		y.extensions[x.Ident()] = x
		return
	}
	y.defs.add(y, prop.(Definition))
}

func (y *Module) resolve(pool schemaPool) error {
	return y.resolveInto(y, pool)
}

func (y *Module) resolveInto(top *Module, pool schemaPool) error {
	if y.featureSet != nil {
		if err := y.featureSet.Resolve(y); err != nil {
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
			if err := i.resolve(top, pool); err != nil {
				return err
			}
			y.imports[i.Prefix()] = i
		}
	}

	if err := y.defs.resolve(top, pool); err != nil {
		return err
	}

	for _, a := range y.augments {
		if err := a.resolve(pool); err != nil {
			return err
		}
	}

	return nil
}

func (y *Module) compile() error {
	for _, i := range y.includes {
		if err := i.compile(); err != nil {
			return err
		}
	}

	for _, i := range y.identities {
		if err := i.compile(); err != nil {
			return err
		}
	}

	if err := compile(y, y.defs); err != nil {
		return err
	}

	for _, a := range y.augments {
		if err := a.compile(); err != nil {
			return err
		}
		if err := a.expand(y); err != nil {
			return err
		}
	}
	return nil
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

func NewImport(parent *Module, moduleName string, loader Loader) *Import {
	return &Import{
		parent:     parent,
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
		y.rev = x
		return
	case SetPrefix:
		y.prefix = string(x)
		return
	}
	panic(fmt.Sprintf("%s:%T not supported in import", y.module.Ident(), prop))
}

func (y *Import) compile() error {
	return y.module.compile()
}

func (y *Import) resolve(parent *Module, pool schemaPool) error {
	if y.loader == nil {
		return c2.NewErr(y.moduleName + " - no module loader defined")
	}
	if y.prefix == "" {
		return c2.NewErr(y.moduleName + " - prefix required on import")
	}
	var err error
	var rev string
	if y.rev != nil {
		rev = y.rev.Ident()
	}
	y.module, err = y.loader(nil, y.moduleName, rev, y.parent.featureSet)
	if err != nil {
		return c2.NewErr(y.moduleName + " - " + err.Error())
	}
	return y.module.resolveInto(parent, pool)
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

func NewInclude(parent *Module, subName string, loader Loader) *Include {
	return &Include{
		parent:  parent,
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
	_, err = y.loader(y.parent, y.subName, rev, y.parent.featureSet)
	if err != nil {
		return c2.NewErr(y.subName + " - " + err.Error())
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
	parent    Meta
	scope     Meta
	ident     string
	desc      string
	ref       string
	when      *When
	defs      *defs
	cases     map[string]*ChoiceCase
	recursive bool
	ifs       []*IfFeature
}

func NewChoice(parent Meta, ident string) *Choice {
	return &Choice{
		parent: parent,
		scope:  parent,
		ident:  ident,
		defs:   newDefs(),
	}
}

func (y *Choice) Cases() map[string]*ChoiceCase {
	return y.cases
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

func (y *Choice) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *Choice) Parent() Meta {
	return y.parent
}

func (y *Choice) scopedParent() Meta {
	return y.scope
}

func (y *Choice) clone(parent Meta) Definition {
	copy := *y
	copy.parent = parent
	copy.defs = copy.defs.clone(&copy)
	return &copy
}

func (y *Choice) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *Choice) DataDefs() []Definition {
	return y.defs.dataDefs
}

func (y *Choice) When() *When {
	return y.when
}

func (y *Choice) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *When:
		y.when = x
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	}
	y.defs.add(y, prop.(Definition))
}

func (y *Choice) resolve(pool schemaPool) error {
	if err := y.defs.resolve(y, pool); err != nil {
		return err
	}
	y.buildCases()
	return nil
}

func (y *Choice) buildCases() {
	y.cases = make(map[string]*ChoiceCase)
	for _, ddef := range y.defs.dataDefs {
		y.cases[ddef.Ident()] = ddef.(*ChoiceCase)
	}
}

func (y *Choice) compile() error {
	return compile(y, y.defs)
}

////////////////////////////////////////////////////

type ChoiceCase struct {
	ident  string
	desc   string
	ref    string
	parent Meta
	scope  Meta
	when   *When
	defs   *defs
	ifs    []*IfFeature
}

func NewChoiceCase(parent Meta, ident string) *ChoiceCase {
	return &ChoiceCase{
		parent: parent,
		scope:  parent,
		ident:  ident,
		defs:   newDefs(),
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

func (y *ChoiceCase) DataDefs() []Definition {
	return y.defs.dataDefs
}

func (y *ChoiceCase) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *ChoiceCase) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *ChoiceCase) scopedParent() Meta {
	return y.scope
}

func (y *ChoiceCase) clone(parent Meta) Definition {
	copy := *y
	copy.parent = parent
	copy.defs = copy.defs.clone(&copy)
	return &copy
}

func (y *ChoiceCase) Condition() *When {
	return y.when
}

func (y *ChoiceCase) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	case *When:
		y.when = x
		return
	}
	y.defs.add(y, prop.(Definition))
}

func (y *ChoiceCase) resolve(pool schemaPool) error {
	return y.defs.resolve(y, pool)
}

func (y *ChoiceCase) compile() error {
	return compile(y, y.defs)
}

////////////////////////////////////////////////////

type Revision struct {
	parent Meta
	scope  Meta
	ident  string
	desc   string
	ref    string
}

func NewRevision(parent Meta, ident string) *Revision {
	return &Revision{
		parent: parent,
		scope:  parent,
		ident:  ident,
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

////////////////////////////////////////////////////

type Container struct {
	ident     string
	desc      string
	ref       string
	typeDefs  map[string]*Typedef
	groupings map[string]*Grouping
	parent    Meta
	scope     Meta
	status    Status
	configPtr *bool
	mandatory bool
	when      *When
	defs      *defs
	recursive bool
	ifs       []*IfFeature
	musts     []*Must
}

func NewContainer(parent Meta, ident string) *Container {
	return &Container{
		parent:    parent,
		scope:     parent,
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

func (y *Container) Status() Status {
	return y.status
}

func (y *Container) Parent() Meta {
	return y.parent
}

func (y *Container) DataDefs() []Definition {
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

func (y *Container) When() *When {
	return y.when
}

func (y *Container) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *Container) Musts() []*Must {
	return y.musts
}

func (y *Container) scopedParent() Meta {
	return y.scope
}

func (y *Container) clone(parent Meta) Definition {
	copy := *y
	copy.parent = parent
	copy.defs = copy.defs.clone(&copy)
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
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		y.typeDefs[x.Ident()] = x
		return
	case *When:
		y.when = x
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	case *Must:
		y.musts = append(y.musts, x)
		return
	case Status:
		y.status = x
	}
	y.defs.add(y, prop.(Definition))
}

func (y *Container) compile() error {
	if y.configPtr == nil {
		b := inheritConfig(y.parent)
		y.configPtr = &b
	}

	return compile(y, y.defs)
}

func (y *Container) resolve(pool schemaPool) error {
	return y.defs.resolve(y, pool)
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
	scope        Meta
	ident        string
	desc         string
	ref          string
	typeDefs     map[string]*Typedef
	groupings    map[string]*Grouping
	key          []string
	keyMeta      []HasType
	when         *When
	configPtr    *bool
	mandatory    bool
	minElements  int
	maxElements  int
	unboundedPtr *bool
	defs         *defs
	recursive    bool
	ifs          []*IfFeature
	musts        []*Must
}

func NewList(parent Meta, ident string) *List {
	return &List{
		parent:    parent,
		scope:     parent,
		ident:     ident,
		defs:      newDefs(),
		groupings: make(map[string]*Grouping),
		typeDefs:  make(map[string]*Typedef),
	}
}

func (y *List) KeyMeta() (keyMeta []HasType) {
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

func (y *List) When() *When {
	return y.when
}

func (y *List) DataDefs() []Definition {
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

func (y *List) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *List) Musts() []*Must {
	return y.musts
}

func (y *List) scopedParent() Meta {
	return y.scope
}

func (y *List) clone(parent Meta) Definition {
	copy := *y
	copy.parent = parent
	copy.defs = copy.defs.clone(&copy)
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
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		y.typeDefs[x.Ident()] = x
		return
	case *When:
		y.when = x
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	case *Must:
		y.musts = append(y.musts, x)
		return
	}
	y.defs.add(y, prop.(Definition))
}

func (y *List) resolve(pool schemaPool) error {
	return y.defs.resolve(y, pool)
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

	if err := compile(y, y.defs); err != nil {
		return err
	}

	y.keyMeta = make([]HasType, len(y.key))
	for i, keyIdent := range y.key {
		// relies on res
		km, valid := y.defs.dataDefsIndex[keyIdent]
		if !valid {
			return c2.NewErr(GetPath(y) + " - " + keyIdent + " key not found for " + GetPath(y))
		}
		y.keyMeta[i], valid = km.(HasType)
		if !valid {
			return c2.NewErr(GetPath(y) + " - " + keyIdent + " expected key with data type")
		}
	}

	return nil
}

////////////////////////////////////////////////////

type Leaf struct {
	parent     Meta
	scope      Meta
	ident      string
	desc       string
	ref        string
	units      string
	configPtr  *bool
	mandatory  bool
	defaultVal interface{}
	dtype      *Type
	when       *When
	ifs        []*IfFeature
	musts      []*Must
}

func NewLeaf(parent Meta, ident string) *Leaf {
	l := &Leaf{
		parent: parent,
		scope:  parent,
		ident:  ident,
	}
	return l
}

func NewLeafWithType(parent Meta, ident string, f val.Format) *Leaf {
	l := NewLeaf(parent, ident)
	l.dtype = NewType(f.String())
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

func (y *Leaf) Units() string {
	return y.units
}

func (y *Leaf) Parent() Meta {
	return y.parent
}

func (y *Leaf) Type() *Type {
	return y.dtype
}

func (y *Leaf) When() *When {
	return y.when
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

func (y *Leaf) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *Leaf) Musts() []*Must {
	return y.musts
}

func (y *Leaf) scopedParent() Meta {
	return y.scope
}

func (y *Leaf) clone(parent Meta) Definition {
	copy := *y
	copy.parent = parent
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
	case SetUnits:
		y.units = string(x)
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
	case *Type:
		y.dtype = x
		return
	case *When:
		y.when = x
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	case *Must:
		y.musts = append(y.musts, x)
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
	if y.units == "" {
		y.units = y.dtype.units
	}
	return nil
}

////////////////////////////////////////////////////

type LeafList struct {
	ident        string
	parent       Meta
	scope        Meta
	desc         string
	ref          string
	units        string
	configPtr    *bool
	mandatory    bool
	dtype        *Type
	minElements  int
	maxElements  int
	unboundedPtr *bool
	defaults     []interface{}
	when         *When
	ifs          []*IfFeature
	musts        []*Must
}

func NewLeafList(parent Meta, ident string) *LeafList {
	l := &LeafList{
		parent: parent,
		scope:  parent,
		ident:  ident,
	}
	return l
}

func NewLeafListWithType(parent Meta, ident string, f val.Format) *LeafList {
	l := NewLeafList(parent, ident)
	l.dtype = NewType(f.String())
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

func (y *LeafList) Units() string {
	return y.units
}

func (y *LeafList) Parent() Meta {
	return y.parent
}

func (y *LeafList) Type() *Type {
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

func (y *LeafList) When() *When {
	return y.when
}

func (y *LeafList) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *LeafList) Musts() []*Must {
	return y.musts
}

func (y *LeafList) scopedParent() Meta {
	return y.scope
}

func (y *LeafList) clone(parent Meta) Definition {
	copy := *y
	copy.parent = parent
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
	case SetUnits:
		y.units = string(x)
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
	case SetMaxElements:
		y.maxElements = int(x)
		return
	case SetMinElements:
		y.minElements = int(x)
		return
	case SetDefault:
		y.defaults = append(y.defaults, x.Value)
	case *Type:
		y.dtype = x
		return
	case *When:
		y.when = x
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	case *Must:
		y.musts = append(y.musts, x)
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
	ident     string
	desc      string
	ref       string
	parent    Meta
	scope     Meta
	configPtr *bool
	mandatory bool
	dtype     *Type
	when      *When
	ifs       []*IfFeature
	musts     []*Must
}

func NewAny(parent Meta, ident string) *Any {
	any := &Any{
		parent: parent,
		scope:  parent,
		ident:  ident,
		dtype:  NewType("any"),
	}
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

func (y *Any) Type() *Type {
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

func (y *Any) Units() string {
	return ""
}

func (y *Any) When() *When {
	return y.when
}

func (y *Any) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *Any) Musts() []*Must {
	return y.musts
}

func (y *Any) scopedParent() Meta {
	return y.scope
}

func (y *Any) clone(parent Meta) Definition {
	copy := *y
	copy.parent = parent
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
	case *When:
		y.when = x
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	case *Must:
		y.musts = append(y.musts, x)
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

	defs *defs
	// see RFC7950 Sec 14
	// no details (config, mandatory)
	// no when
}

func NewGrouping(parent Meta, ident string) *Grouping {
	return &Grouping{
		parent:    parent,
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

func (y *Grouping) Groupings() map[string]*Grouping {
	return y.groupings
}

func (y *Grouping) Typedefs() map[string]*Typedef {
	return y.typeDefs
}

func (y *Grouping) DataDefs() []Definition {
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
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		y.typeDefs[x.Ident()] = x
		return
	}
	y.defs.add(y, prop.(Definition))
}

////////////////////////////////////////////////////

type Uses struct {
	ident    string
	desc     string
	ref      string
	parent   Meta
	scope    Meta
	schemaId schemaId
	refines  []*Refine
	when     *When
	ifs      []*IfFeature
	augments []*Augment
}

func NewUses(parent Meta, ident string) *Uses {
	u := &Uses{
		parent: parent,
		scope:  parent,
		ident:  ident,
	}
	u.schemaId = schemaId(u) // unique identifier
	return u
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

func (y *Uses) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *Uses) Augments() []*Augment {
	return y.augments
}

func (y *Uses) Parent() Meta {
	return y.parent
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
		y.refines = append(y.refines, x)
		return
	case *Augment:
		y.augments = append(y.augments, x)
		return
	case *When:
		y.when = x
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	}
	panic(fmt.Sprintf("%T not supported in uses", prop))
}

func (y *Uses) scopedParent() Meta {
	return y.scope
}

func (y *Uses) clone(parent Meta) Definition {
	copy := *y
	copy.parent = parent
	return &copy
}

type schemaId interface{}
type schemaPool map[schemaId][]Definition

func (y *Uses) resolve(parent Meta, pool schemaPool, resolved resolvedListener) error {
	if on, err := checkFeature(y); !on || err != nil {
		return err
	}

	g := y.findScopedTarget()
	if g == nil {
		return c2.NewErr(GetPath(y) + " - " + y.ident + " group not found")
	}

	for _, a := range y.augments {
		if err := a.resolve(pool); err != nil {
			return err
		}
	}

	if ddefs, hasSource := pool[y.schemaId]; hasSource {
		for _, ddef := range ddefs {
			if err := resolved(ddef); err != nil {
				return err
			}
		}
	} else {
		var ddefs []Definition
		resolvedPassthru := func(gdef Definition) error {
			ddef := gdef.(cloneable).clone(parent)
			if on, err := checkFeature(ddef.(HasIfFeatures)); !on || err != nil {
				return err
			}
			if err := y.refine(ddef, pool); err != nil {
				return err
			}
			if y.when != nil {
				if err := Set(ddef, y.when); err != nil {
					return err
				}
			}
			ddefs = append(ddefs, ddef)
			return resolved(ddef)
		}
		if err := g.defs.resolveDefs(parent, pool, g.defs.unresolved, resolvedPassthru); err != nil {
			return err
		}
		for _, a := range y.augments {
			if err := a.expand(y.parent); err != nil {
				return err
			}
		}
		pool[y.schemaId] = ddefs
	}

	return nil
}

func (y *Uses) refine(d Definition, pool schemaPool) error {
	for _, r := range y.refines {
		if on, err := checkFeature(r); !on || err != nil {
			return err
		}
		ident, path := r.splitIdent()
		if ident == d.Ident() {
			if path == "" {
				return r.refine(d)
			}
			hasDefs, ok := d.(HasDefinitions)
			if !ok {
				return c2.NewErr(fmt.Sprintf("%s:cannot refine %s, %s has no children", GetPath(y), r.Ident(), d.Ident()))
			}
			// children are not resolved yet.
			if err := hasDefs.(resolver).resolve(pool); err != nil {
				return err
			}
			target := Find(hasDefs, path)
			if target == nil {
				return c2.NewErr(fmt.Sprintf("%s:could not find target for refine %s", GetPath(y), r.Ident()))
			}
			return r.refine(target)
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
		p := y.scopedParent()
		for p != nil {
			if hasGroups, ok := p.(HasGroupings); ok {
				if g, found := hasGroups.Groupings()[y.ident]; found {
					return g
				}
			}
			if hasScoped, ok := p.(cloneable); ok {
				p = hasScoped.scopedParent()
			} else {
				p = p.Parent()
			}
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
	ifs            []*IfFeature
	musts          []*Must
}

func NewRefine(parent *Uses, path string) *Refine {
	return &Refine{
		parent: parent,
		ident:  path,
	}
}

func (y *Refine) splitIdent() (string, string) {
	slash := strings.IndexRune(y.ident, '/')
	if slash < 0 {
		return y.ident, ""
	}
	return y.ident[:slash], y.ident[slash+1:]
}

func (y *Refine) refine(target Definition) error {
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
	for _, m := range y.Musts() {
		if err := Set(target, m); err != nil {
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

func (y *Refine) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *Refine) Parent() Meta {
	return y.parent
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

func (y *Refine) Musts() []*Must {
	return y.musts
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
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	case *Must:
		y.musts = append(y.musts, x)
		return
	}
	panic(fmt.Sprintf("%T not supported in refine", prop))
}

////////////////////////////////////////////////////

type RpcInput struct {
	parent    *Rpc
	scope     *Rpc
	desc      string
	ref       string
	typeDefs  map[string]*Typedef
	groupings map[string]*Grouping
	defs      *defs
	ifs       []*IfFeature
	musts     []*Must
}

func NewRpcInput(parent *Rpc) *RpcInput {
	return &RpcInput{
		parent:    parent,
		scope:     parent,
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

func (y *RpcInput) Groupings() map[string]*Grouping {
	return y.groupings
}

func (y *RpcInput) Typedefs() map[string]*Typedef {
	return y.typeDefs
}

func (y *RpcInput) DataDefs() []Definition {
	return y.defs.dataDefs
}

func (y *RpcInput) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *RpcInput) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *RpcInput) Musts() []*Must {
	return y.musts
}

func (y *RpcInput) scopedParent() Meta {
	return y.scope
}

func (y *RpcInput) clone(parent Meta) Definition {
	copy := *y
	copy.parent = parent.(*Rpc)
	copy.defs = y.defs.clone(&copy)
	return &copy
}

func (y *RpcInput) resolve(pool schemaPool) error {
	return y.defs.resolve(y, pool)
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
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		y.typeDefs[x.Ident()] = x
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	case *Must:
		y.musts = append(y.musts, x)
		return
	}
	y.defs.add(y, prop.(Definition))
}

func (y *RpcInput) compile() error {
	return compile(y, y.defs)
}

////////////////////////////////////////////////////

type RpcOutput struct {
	parent    *Rpc
	scope     *Rpc
	desc      string
	ref       string
	typeDefs  map[string]*Typedef
	groupings map[string]*Grouping
	defs      *defs
	ifs       []*IfFeature
	musts     []*Must
}

func NewRpcOutput(parent *Rpc) *RpcOutput {
	return &RpcOutput{
		parent:    parent,
		scope:     parent,
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

func (y *RpcOutput) Groupings() map[string]*Grouping {
	return y.groupings
}

func (y *RpcOutput) Typedefs() map[string]*Typedef {
	return y.typeDefs
}

func (y *RpcOutput) DataDefs() []Definition {
	return y.defs.dataDefs
}

func (y *RpcOutput) Mandatory() bool {
	return false
}

func (y *RpcOutput) Definition(ident string) Definition {
	return y.defs.definition(ident)
}

func (y *RpcOutput) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *RpcOutput) Musts() []*Must {
	return y.musts
}

func (y *RpcOutput) scopedParent() Meta {
	return y.scope
}

func (y *RpcOutput) clone(parent Meta) Definition {
	copy := *y
	copy.parent = parent.(*Rpc)
	copy.defs = y.defs.clone(&copy)
	return &copy
}

func (y *RpcOutput) resolve(pool schemaPool) error {
	return y.defs.resolve(y, pool)
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
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		y.typeDefs[x.Ident()] = x
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	case *Must:
		y.musts = append(y.musts, x)
		return
	}
	y.defs.add(y, prop.(Definition))
}

func (y *RpcOutput) compile() error {
	return compile(y, y.defs)
}

////////////////////////////////////////////////////

type Rpc struct {
	ident     string
	parent    Meta
	scope     Meta
	desc      string
	ref       string
	typeDefs  map[string]*Typedef
	groupings map[string]*Grouping
	input     *RpcInput
	output    *RpcOutput
	ifs       []*IfFeature
}

func NewRpc(parent Meta, ident string) *Rpc {
	return &Rpc{
		parent:    parent,
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

func (y *Rpc) Groupings() map[string]*Grouping {
	return y.groupings
}

func (y *Rpc) Typedefs() map[string]*Typedef {
	return y.typeDefs
}

func (y *Rpc) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *Rpc) resolve(pool schemaPool) error {
	if y.input != nil {
		if err := y.input.resolve(pool); err != nil {
			return err
		}
	}
	if y.output != nil {
		if err := y.output.resolve(pool); err != nil {
			return err
		}
	}
	return nil
}

func (y *Rpc) scopedParent() Meta {
	return y.scope
}

func (y *Rpc) clone(parent Meta) Definition {
	copy := *y
	copy.parent = parent
	if y.input != nil {
		copy.input = y.input.clone(&copy).(*RpcInput)
	}
	if y.output != nil {
		copy.output = y.output.clone(&copy).(*RpcOutput)
	}
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
		y.input = x
		return
	case *RpcOutput:
		y.output = x
		return
	case *Grouping:
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		y.typeDefs[x.Ident()] = x
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
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
	scope     Meta
	desc      string
	ref       string
	typeDefs  map[string]*Typedef
	groupings map[string]*Grouping
	defs      *defs
	ifs       []*IfFeature
}

func NewNotification(parent Meta, ident string) *Notification {
	return &Notification{
		parent:    parent,
		scope:     parent,
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

func (y *Notification) DataDefs() []Definition {
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

func (y *Notification) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *Notification) scopedParent() Meta {
	return y.scope
}

func (y *Notification) clone(parent Meta) Definition {
	copy := *y
	copy.parent = parent
	copy.defs = y.defs.clone(&copy)
	return &copy
}

func (y *Notification) resolve(pool schemaPool) error {
	return y.defs.resolve(y, pool)
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
		y.groupings[x.Ident()] = x
		return
	case *Typedef:
		y.typeDefs[x.Ident()] = x
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	}
	y.defs.add(y, prop.(Definition))
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
	units      string
	defaultVal interface{}
	dtype      *Type
}

func NewTypedef(parent Meta, ident string) *Typedef {
	return &Typedef{
		parent: parent,
		ident:  ident,
	}
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

func (y *Typedef) Units() string {
	return y.units
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

func (y *Typedef) Type() *Type {
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
	case SetUnits:
		y.units = string(x)
		return
	case *Type:
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
		c2.NewErr(GetPath(y) + " - " + y.ident + " type required")
	}

	return compile(y, nil)
}

////////////////////////////////////////////////

type Augment struct {
	ident  string
	parent Meta
	desc   string
	ref    string
	defs   *defs
	when   *When
	ifs    []*IfFeature
}

func NewAugment(parent Meta, path string) *Augment {
	return &Augment{
		parent: parent,
		ident:  path,
		defs:   newDefs(),
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

func (y *Augment) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *Augment) Parent() Meta {
	return y.parent
}

func (y *Augment) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *When:
		y.when = x
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	}
	y.defs.add(y, prop.(Definition))
}

func (y *Augment) resolve(pool schemaPool) error {
	if on, err := checkFeature(y); !on || err != nil {
		return err
	}
	return y.defs.resolve(y, pool)
}

func (y *Augment) compile() error {
	if err := compile(y, y.defs); err != nil {
		return err
	}
	return nil
}

func (y *Augment) expand(parent Meta) error {

	// RFC7950 Sec 7.17
	// "The target node MUST be either a container, list, choice, case, input,
	//   output, or notification node."
	target := Find(parent.(HasDefinitions), y.ident)
	if target == nil {
		return c2.NewErr(GetPath(y) + " - augment target is not found " + y.ident)
	}

	// expand
	for _, x := range y.defs.actions {
		if err := Set(target, x.clone(target)); err != nil {
			return err
		}
	}
	for _, x := range y.defs.notifications {
		if err := Set(target, x.clone(target)); err != nil {
			return err
		}
	}
	for _, x := range y.defs.dataDefs {
		if err := Set(target, x.(cloneable).clone(target)); err != nil {
			return err
		}
	}
	return nil
}

////////////////////////////////////////////////////

type Type struct {
	typeIdent      string
	desc           string
	ref            string
	format         val.Format
	enums          []*Enum
	enum           val.EnumList
	ranges         []Range
	lengths        []Range
	path           string
	units          string
	fractionDigits int
	patterns       []string
	defaultVal     interface{}
	delegate       *Type
	base           string
	identity       *Identity
	unionTypes     []*Type
}

func NewType(typeIdent string) *Type {
	return &Type{
		typeIdent: typeIdent,
	}
}

func (y *Type) Ident() string {
	return y.typeIdent
}

func (y *Type) Description() string {
	return y.desc
}

func (y *Type) Reference() string {
	return y.ref
}

func (y *Type) Range() []Range {
	return y.ranges
}

func (y *Type) Length() []Range {
	return y.lengths
}

func (y *Type) Patterns() []string {
	return y.patterns
}

func (y *Type) Format() val.Format {
	return y.format
}

func (y *Type) Path() string {
	return y.path
}

func (y *Type) Enum() val.EnumList {
	return y.enum
}

func (y *Type) Enums() []*Enum {
	return y.enums
}

func (y *Type) Base() *Identity {
	return y.identity
}

func (y *Type) Union() []*Type {
	return y.unionTypes
}

func (y *Type) FractionDigits() int {
	return y.fractionDigits
}

func (y *Type) UnionFormats() []val.Format {
	f := make([]val.Format, len(y.unionTypes))
	for i, u := range y.unionTypes {
		f[i] = u.Format()
	}
	return f
}

func (y *Type) HasDefault() bool {
	return y.defaultVal != nil
}

func (y *Type) DefaultValue() interface{} {
	return y.defaultVal
}

// Resolve is the effective datatype if this type points to a different
// dataType, which is the case for leafRefs.  Otherwise this just returns
// itself
func (y *Type) Resolve() *Type {
	if y.delegate == nil {
		panic("no delegate")
	}
	return y.delegate
}

func (y *Type) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case SetLenRange:
		y.lengths = append(y.lengths, Range(x))
		return
	case SetValueRange:
		y.ranges = append(y.ranges, Range(x))
		return
	case SetPattern:
		y.patterns = append(y.patterns, string(x))
		return
	case SetPath:
		y.path = string(x)
		return
	case *Enum:
		y.enums = append(y.enums, x)
		return
	case SetBase:
		y.base = string(x)
		return
	case *Type:
		y.unionTypes = append(y.unionTypes, x)
		return
	case SetFractionDigits:
		y.fractionDigits = int(x)
		return
	}
	panic(fmt.Sprintf("%T not supported in type", prop))
}

func (base *Type) mixin(derived *Type) {
	if len(derived.patterns) == 0 {
		derived.patterns = base.patterns
	}
	if base.path != "" && derived.path == "" {
		derived.path = base.path
	}
	if derived.ranges == nil {
		derived.ranges = base.ranges
	} else if base.ranges != nil {
		for _, r := range base.ranges {
			derived.ranges = append(derived.ranges, r)
		}
	}
	if derived.enums == nil {
		derived.enums = base.enums
	}
	if derived.base == "" {
		derived.base = base.base
	}
	if derived.unionTypes == nil {
		derived.unionTypes = base.unionTypes
	}
	if derived.lengths == nil {
		derived.lengths = base.lengths
	} else if base.lengths != nil {
		for _, r := range base.lengths {
			derived.lengths = append(derived.lengths, r)
		}
	}
	if derived.defaultVal == nil {
		derived.defaultVal = base.defaultVal
	}
	if derived.units == "" {
		derived.units = base.units
	}
	derived.format = base.format
}

func (y *Type) compile(parent Meta) error {
	if y == nil {
		return errors.New("no type set on " + GetPath(parent))
	}
	if int(y.format) != 0 {
		return nil
	}
	var hasTypedef bool
	y.format, hasTypedef = val.TypeAsFormat(y.typeIdent)
	if !hasTypedef {
		tdef, err := y.findScopedTypedef(parent, y.typeIdent)
		if err != nil {
			return err
		}

		// Don't use resolve here because if a typedef is a leafref, you want
		// the unreolved here and resolve it below
		tdef.dtype.mixin(y)

		// default and units are strangely not settable on type, only on leafs and
		// typedefs so we can blindy set values here
		y.defaultVal = tdef.defaultVal
		y.units = tdef.units
	}

	if y.format == val.FmtLeafRef || y.format == val.FmtLeafRefList {
		if y.path == "" {
			return c2.NewErr(GetPath(parent) + " - " + y.typeIdent + " path is required")
		}
		// parent is a leaf, so start with parent's parent which is a container-ish
		resolvedMeta := Find(parent, y.path)
		if resolvedMeta == nil {
			err := c2.NewErr(GetPath(parent) + " - " + y.typeIdent + " could not resolve leafref path " + y.path)
			fmt.Println(err.Error())
			y.delegate = y
		} else {
			y.delegate = resolvedMeta.(HasType).Type()
		}
	} else {
		y.delegate = y
	}

	if y.format == val.FmtIdentityRef {
		m, baseIdent, err := rootByIdent(parent, y.base)
		if err != nil {
			return err
		}
		identity, found := m.Identities()[baseIdent]
		if !found {
			return c2.NewErr(GetPath(parent) + " - " + y.base + " identity not found")
		}
		y.identity = identity
	}

	if _, isList := parent.(*LeafList); isList && !y.format.IsList() {
		y.format = val.Format(int(y.format) + 1024)
	}

	if y.format == val.FmtUnion {
		if len(y.unionTypes) == 0 {
			return c2.NewErr(GetPath(parent) + " - unions need at least one type")
		}
		for _, u := range y.unionTypes {
			if err := u.compile(parent); err != nil {
				return err
			}
		}
	} else if len(y.unionTypes) > 0 {
		return c2.NewErr(GetPath(parent) + " - embedded types are only for union types")
	}

	if y.format == val.FmtEnum || y.format == val.FmtEnumList {
		y.enum = make(val.EnumList, len(y.enums))
		nextId := 0
		for i, item := range y.enums {
			if item.val > 0 {
				nextId = item.val
			} else {
				item.val = nextId
			}
			y.enum[i] = val.Enum{
				Id:    nextId,
				Label: item.label,
			}
			nextId++
		}
	}

	return nil
}

func (y *Type) findScopedTypedef(parent Meta, ident string) (*Typedef, error) {
	// lazy load grouping
	var found *Typedef
	xMod, xIdent, err := externalModule(parent, ident)
	if err != nil {
		goto nomatch
	}
	if xMod != nil {
		found = xMod.Typedefs()[xIdent]
	} else {
		p := parent
		for p != nil {
			if ptd, ok := p.(HasTypedefs); ok {
				if found = ptd.Typedefs()[ident]; found != nil {
					break
				}
			}
			if hasScope, ok := p.(cloneable); ok {
				p = hasScope.scopedParent()
			} else {
				p = p.Parent()
			}
		}
	}
nomatch:
	if found == nil {
		return nil, c2.NewErr(GetPath(parent) + " - typedef " + y.typeIdent + " not found")
	}

	if err := found.compile(); err != nil {
		return nil, err
	}
	return found, nil
}

////////////////////////////////////////

type Identity struct {
	parent     *Module
	ident      string
	desc       string
	ref        string
	derivedIds []string
	derived    map[string]*Identity
	ifs        []*IfFeature
}

func NewIdentity(parent *Module, ident string) *Identity {
	return &Identity{
		parent: parent,
		ident:  ident,
	}
}

func (y *Identity) Description() string {
	return y.desc
}

func (y *Identity) Reference() string {
	return y.ref
}

func (y *Identity) Ident() string {
	return y.ident
}

func (y *Identity) BaseIds() []string {
	return y.derivedIds
}

func (y *Identity) Identities() map[string]*Identity {
	return y.derived
}

func (y *Identity) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *Identity) Parent() Meta {
	return y.parent
}

func (y *Identity) compile() error {
	if y.derived != nil {
		return nil
	}
	y.derived = make(map[string]*Identity)
	y.derived[y.ident] = y
	for _, baseId := range y.derivedIds {
		m, baseIdent, err := rootByIdent(y, baseId)
		if err != nil {
			return err
		}
		ident, found := m.Identities()[baseIdent]
		if !found {
			return c2.NewErr(GetPath(y) + " - " + baseId + " identity not found")
		}
		y.derived[baseId] = ident
		if err := ident.compile(); err != nil {
			return err
		}
		for subId, subIdent := range ident.Identities() {
			y.derived[subId] = subIdent
		}
	}
	return nil
}

func (y *Identity) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case SetBase:
		y.derivedIds = append(y.derivedIds, string(x))
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	}
	panic(fmt.Sprintf("%s : %T not supported in type", GetPath(y), prop))
}

////////////////////////////////////////

type Feature struct {
	parent *Module
	ident  string
	desc   string
	ref    string
	ifs    []*IfFeature
}

func NewFeature(parent *Module, ident string) *Feature {
	return &Feature{
		parent: parent,
		ident:  ident,
	}
}

func (y *Feature) Description() string {
	return y.desc
}

func (y *Feature) Reference() string {
	return y.ref
}

func (y *Feature) Ident() string {
	return y.ident
}

func (y *Feature) IfFeatures() []*IfFeature {
	return y.ifs
}

func (y *Feature) Parent() Meta {
	return y.parent
}

func (y *Feature) compile() error {
	return nil
}

func (y *Feature) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *IfFeature:
		y.ifs = append(y.ifs, x)
		return
	}
	panic(fmt.Sprintf("%T not supported in type", prop))
}

type IfFeature struct {
	expr string
}

func NewIfFeature(expr string) *IfFeature {
	return &IfFeature{
		expr: expr,
	}
}

func (y *IfFeature) Expression() string {
	return y.expr
}

func (y *IfFeature) Evaluate(enabled map[string]*Feature) (bool, error) {
	e := &ifFeatureEval{
		features: enabled,
		expr:     y.expr,
	}
	e.eval(false)
	b := e.pop()
	err := e.lastErr
	if err == nil && len(e.stack) != 0 {
		return false, c2.NewErr("syntax err in feature expression:" + y.expr)
	}
	return b, err
}

type ifFeatureEval struct {
	features map[string]*Feature
	expr     string
	stack    []bool
	pos      int
	lastErr  error
}

func (y *ifFeatureEval) eval(greedy bool) {
	for !y.end() {
		tok := y.next()
		switch tok {
		case "(":
			y.eval(false)
		case ")":
			return
		case "and":
			y.eval(true)
			a, b := y.pop(), y.pop()
			y.push(a && b)
		case "not":
			y.eval(true)
			y.push(!y.pop())
		case "or":
			y.eval(false)
			a, b := y.pop(), y.pop()
			y.push(a || b)
		default:
			_, found := y.features[tok]
			y.push(found)
		}
		if greedy {
			return
		}
	}
	return
}

func (y *ifFeatureEval) end() bool {
	return y.pos >= len(y.expr)
}

func (y *ifFeatureEval) eatws() {
	for !y.end() {
		if y.expr[y.pos] != ' ' {
			break
		}
		y.pos++
	}
}

func (y *ifFeatureEval) next() string {
	y.eatws()
	start := y.pos
	for !y.end() {
		switch y.expr[y.pos] {
		case ' ':
			goto brk
		case '(', ')':
			if y.pos == start {
				y.pos++
			}
			goto brk
		}
		y.pos++
	}
brk:
	tok := y.expr[start:y.pos]
	return tok
}

func (y *ifFeatureEval) pop() bool {
	if len(y.stack) == 0 {
		y.lastErr = c2.NewErr("syntax err in feature expression:" + y.expr)
		return false
	}
	last := len(y.stack) - 1
	b := y.stack[last]
	y.stack = y.stack[0:last]
	return b
}

func (y *ifFeatureEval) push(b bool) {
	y.stack = append(y.stack, b)
}

//////////////////////////////////
type When struct {
	expr string
	desc string
	ref  string
}

func NewWhen(expr string) *When {
	return &When{
		expr: expr,
	}
}

func (y *When) Expression() string {
	return y.expr
}

func (y *When) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	}
	panic(fmt.Sprintf("%T not supported in type", prop))
}

//////////////////////////////////
type Must struct {
	expr string
}

func NewMust(expr string) *Must {
	return &Must{
		expr: expr,
	}
}

func (y *Must) Expression() string {
	return y.expr
}

//////////////////////////////////
type Extension struct {
	parent *Module
	ident  string
	desc   string
	ref    string
	args   map[string]*ExtensionArg
}

func NewExtension(parent *Module, ident string) *Extension {
	return &Extension{
		parent: parent,
		ident:  ident,
		args:   make(map[string]*ExtensionArg),
	}
}

func (y *Extension) Ident() string {
	return y.ident
}

func (y *Extension) Description() string {
	return y.desc
}

func (y *Extension) Reference() string {
	return y.ref
}

func (y *Extension) Parent() Meta {
	return y.parent
}

func (y *Extension) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case *ExtensionArg:
		y.args[x.ident] = x
		return
	}
	panic(fmt.Sprintf("%T not supported in type", prop))
}

//////////////////////////////////
type ExtensionArg struct {
	parent     *Extension
	ident      string
	desc       string
	ref        string
	yinElement bool
}

func NewExtensionArg(parent *Extension, ident string) *ExtensionArg {
	return &ExtensionArg{
		parent: parent,
		ident:  ident,
	}
}

func (y *ExtensionArg) Ident() string {
	return y.ident
}

func (y *ExtensionArg) Description() string {
	return y.desc
}

func (y *ExtensionArg) Reference() string {
	return y.ref
}

func (y *ExtensionArg) Parent() Meta {
	return y.parent
}

func (y *ExtensionArg) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case SetYinElement:
		y.yinElement = bool(x)
		return
	}
	panic(fmt.Sprintf("%T not supported in type", prop))
}

//////////////////////////////////

type Enum struct {
	label string
	desc  string
	ref   string
	val   int
}

func NewEnum(label string) *Enum {
	return &Enum{
		label: label,
	}
}

func (y *Enum) Description() string {
	return y.desc
}

func (y *Enum) Reference() string {
	return y.ref
}

func (y *Enum) Ident() string {
	return y.label
}

func (y *Enum) Value() int {
	return y.val
}

func (y *Enum) add(prop interface{}) {
	switch x := prop.(type) {
	case SetDescription:
		y.desc = string(x)
		return
	case SetReference:
		y.ref = string(x)
		return
	case SetEnumValue:
		y.val = int(x)
		return
	}
	panic(fmt.Sprintf("%T not supported in type", prop))
}
