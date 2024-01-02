package meta

// /////////////////
// Interfaces
// ////////////////
// Definition represent nearly everythihng in YANG, more specifically, anything
// that can have an extention, which is nearly everything
type Meta interface {
	HasExtensions

	// Parent in the YANG schema tree
	Parent() Meta
}

// HasExtensions is support by almost every structure. See YANG
// language extensions for more information
type HasExtensions interface {
	// User customized YANG found in the body
	Extensions() []*Extension
}

type hasAddExtension interface {
	HasExtensions

	addExtension(x *Extension)
}

type cloneable interface {
	clone(parent Meta) interface{}
}

// Identifiable are things that have a unique identifier allowing it to be found
// in a list.
type Identifiable interface {

	// Ident is short for identifier or name of item.  Example: 'leaf foo {...' then 'foo' is ident
	Ident() string
}

// Describable is anything that can have a description, oddly, most data definitions except
// 'case', 'input' or 'output'
type Describable interface {

	// Description of meta item
	Description() string

	// Reference is a human-readable, cross-reference to some external source.  Example: Item #89 of foo catalog"
	Reference() string

	setDescription(desc string)
	setReference(ref string)
}

// Definition data structure defining details. This includes data definitions like
// container and leaf, but also notifications and actions
type Definition interface {
	Meta
	Identifiable
	setParent(p Meta)
	getOriginalParent() Definition
}

type HasPresence interface {
	Meta
	Presence() string
	setPresence(string)
}

type HasUnique interface {
	Meta
	Unique() [][]string
	setUnique([][]string)
}

type HasStatus interface {
	Meta

	// Status is useful to mark things deprecated
	Status() Status

	setStatus(s Status)
}

type HasDefinitions interface {
	Definition

	// Definition returns DataDefinition, Action or Notification by name
	Definition(ident string) Definition
}

// HasDefinitions holds container, leaf, list, etc definitions which
// often (but not always) also hold notifications and actions
type HasDataDefinitions interface {
	HasDefinitions

	DataDefinitions() []Definition

	addDataDefinition(Definition) error

	// addDataDefinition will change parent, this won't and useful when moving around
	// potentially recursive defintions where parent's children aren't always the child's
	// parent.
	addDataDefinitionWithoutOwning(Definition) error

	popDataDefinitions() []Definition
}

type HasUnits interface {
	Units() string
	setUnits(units string)
}

type HasNotifications interface {
	HasDataDefinitions
	Notifications() map[string]*Notification

	addNotification(*Notification) error
	setNotifications(map[string]*Notification)
}

type HasActions interface {
	HasDataDefinitions
	Actions() map[string]*Rpc

	addAction(a *Rpc) error
	setActions(map[string]*Rpc)
}

type HasGroupings interface {
	HasDataDefinitions
	Groupings() map[string]*Grouping
	addGrouping(g *Grouping) error
}

type HasAugments interface {
	Augments() []*Augment
	addAugments(*Augment)
}

type HasTypedefs interface {
	Typedefs() map[string]*Typedef
	addTypedef(t *Typedef) error
}

type HasIfFeatures interface {
	IfFeatures() []*IfFeature
	addIfFeature(*IfFeature)
}

type HasWhen interface {
	When() *When
	setWhen(*When)
}

type HasMusts interface {
	Musts() []*Must
	addMust(*Must)
	setMusts([]*Must)
}

type HasConfig interface {
	Config() bool
	IsConfigSet() bool
	setConfig(bool)
}

type HasMandatory interface {
	Mandatory() bool
	IsMandatorySet() bool
	setMandatory(bool)
}

type HasDetails interface {
	Definition
	HasMandatory
	HasConfig
}

type HasCases interface {
	Definition
	addCase(*ChoiceCase) error
}

type HasOrderedBy interface {
	OrderedBy() OrderedBy
	setOrderedBy(order OrderedBy)
}

type HasErrorMessage interface {
	ErrorMessage() string
	setErrorMessage(string)

	ErrorAppTag() string
	setErrorAppTag(string)
}

type HasMinMax interface {
	MaxElements() int
	IsMaxElementsSet() bool
	setMaxElements(int)

	MinElements() int
	IsMinElementsSet() bool
	setMinElements(int)
}

type HasUnbounded interface {
	Unbounded() bool
	IsUnboundedSet() bool
	setUnbounded(bool)
}

type HasListDetails interface {
	Definition
	HasMinMax
	HasUnbounded
	HasOrderedBy
}

type HasDefault interface {
	HasDefault() bool
	addDefault(string)
	DefaultValue() interface{}
	setDefaultValue(interface{})
	clearDefault()
}

type HasDefaultValue interface {
	HasDefault
	Default() string
	setDefault(string)
}

type HasDefaultValues interface {
	HasDefault
	Default() []string
	setDefault(d []string)
}

type Leafable interface {
	Definition
	HasDefault
	HasUnits
	HasType
}

type HasType interface {
	Type() *Type
	setType(*Type)
}

// Status is indication of definition obsolense
type Status int

const (
	Current Status = iota
	Deprecated
	Obsolete
)

// Loader abstracts yang modules are loaded from file parsers.
type Loader func(parent *Module, name string, rev string, features FeatureSet, loader Loader) (*Module, error)
