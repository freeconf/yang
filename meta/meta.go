package meta

///////////////////
// Interfaces
//////////////////
// Definition represent nearly everythihng in YANG, more specifically, anything
// that can have an extention, which is nearly everything
type Meta interface {
	Parent() Meta
}

type compilable interface {
	compile() error
}

type buildable interface {
	add(o interface{})
}

type cloneable interface {
	scopedParent() Meta

	clone(parent Meta) Definition
}

type resolver interface {
	resolve(schemaPool) error
}

// Identifiable are things that have a unique identifier allowing it to be found
// in a list.
type Identifiable interface {
	Meta
	Ident() string
}

// Describable is anything that can have a description, oddly, most data definitions except
// 'case', 'input' or 'output'
type Describable interface {
	Meta
	Description() string
	Reference() string
}

// Definition data structure defining details. This includes data definitions like
// container and leaf, but also notofications and actions
type Definition interface {
	Identifiable
}

type HasStatus interface {
	Status() Status
}

type HasDefinitions interface {
	Definition

	Definition(ident string) Definition
}

type HasDataDefs interface {
	HasDefinitions

	DataDefs() []Definition
}

type HasNotifications interface {
	HasDataDefs
	Notifications() map[string]*Notification
}

type HasActions interface {
	HasDataDefs
	Actions() map[string]*Rpc
}

type HasGroupings interface {
	HasDataDefs
	Groupings() map[string]*Grouping
}

type HasAugments interface {
	HasDataDefs
	Augments() []*Augment
}

type HasTypedefs interface {
	Definition
	Typedefs() map[string]*Typedef
}

type HasIfFeatures interface {
	Meta
	IfFeatures() []*IfFeature
}

type HasWhen interface {
	Meta
	When() *When
}

type HasMusts interface {
	Meta
	Musts() []*Must
}

type HasDetails interface {
	Definition
	Config() bool
	Mandatory() bool
}

type HasListDetails interface {
	Definition
	MaxElements() int
	MinElements() int
	Unbounded() bool
}

type HasType interface {
	Definition
	HasDefault() bool
	Default() interface{}
	Type() *Type
	Units() string
}

type Status int

const (
	Current Status = iota
	Deprecated
	Obsolete
)

type Loader func(parent *Module, name string, rev string, features FeatureSet) (*Module, error)
