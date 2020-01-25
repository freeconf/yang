package meta

///////////////////
// Interfaces
//////////////////
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

	addExtension(x *Extension)
}

type cloneable interface {
	scopedParent() Meta

	clone(parent Meta) interface{}
}

type recursable interface {
	HasDataDefinitions
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
}

type HasStatus interface {

	// Status is useful to mark things deprecated
	Status() Status

	setStatus(s Status)
}

type HasDefinitions interface {
	Definition

	// Definition returns DataDefinition, Action or Notification by name
	Definition(ident string) Definition

	// rare chance this is part of a recursive schema.  If so, care should
	// be taken navigating the schema tree (information model).  Navigating
	// the actual config/metrics (data model) should not be a problem
	IsRecursive() bool

	markRecursive()
}

// HasDefinitions holds container, leaf, list, etc definitions which
// often (but not always) also hold notifications and actions
type HasDataDefinitions interface {
	HasDefinitions

	DataDefinitions() []Definition

	addDataDefinition(Definition)

	popDataDefinitions() []Definition
}

type HasUnits interface {
	Units() string

	setUnits(units string)
}

type HasNotifications interface {
	HasDataDefinitions
	Notifications() map[string]*Notification

	addNotification(*Notification)
	setNotifications(map[string]*Notification)
}

type HasActions interface {
	HasDataDefinitions
	Actions() map[string]*Rpc

	addAction(a *Rpc)
	setActions(map[string]*Rpc)
}

type HasGroupings interface {
	HasDataDefinitions
	Groupings() map[string]*Grouping
	addGrouping(g *Grouping)
}

type HasAugments interface {
	Augments() []*Augment
	addAugments(*Augment)
}

type HasTypedefs interface {
	Definition
	Typedefs() map[string]*Typedef
	addTypedef(t *Typedef)
}

type HasIfFeatures interface {
	Meta
	IfFeatures() []*IfFeature
	addIfFeature(*IfFeature)
}

type HasWhen interface {
	Meta
	When() *When
	setWhen(*When)
}

type HasMusts interface {
	Meta
	Musts() []*Must
	addMust(*Must)
}

type HasDetails interface {
	Definition
	Config() bool
	Mandatory() bool

	setConfig(bool)
	isConfigSet() bool
	setMandatory(bool)
}

type HasListDetails interface {
	Definition
	MaxElements() int
	MinElements() int
	Unbounded() bool

	setUnbounded(bool)
	setMinElements(int)
	setMaxElements(int)
}

// TODO: rename to Leafable because there's more than just Type
type HasType interface {
	Definition
	HasDefault() bool
	Default() interface{}
	Type() *Type
	Units() string

	setType(*Type)
	setUnits(string)
	setDefault(interface{})
}

type Status int

const (
	Current Status = iota
	Deprecated
	Obsolete
)

// Loader abstracts yang modules are loaded from file parsers.
type Loader func(parent *Module, name string, rev string, features FeatureSet, loader Loader) (*Module, error)
