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
	add(o interface{})
}

type movable interface {
	compilable
	setParent(Meta)
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
	compilable
}

type HasDefinitions interface {
	Definition

	Definition(ident string) Definition
}

// Everything in Definitions except Action and Notification
type DataDef interface {
	Definition

	setParent(Meta)
	clone(deep bool) DataDef
}

type HasDataDefs interface {
	HasDefinitions
	DataDefs() []DataDef
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

type HasConditions interface {
	Meta
	Conditions() []Condition
}

// 'when' conditions
type Condition interface {
	Meta
	compilable
	setParent(Meta)
	Evaluate() (bool, error)
}

type HasDetails interface {
	DataDef
	Config() bool
	Mandatory() bool
}

type HasListDetails interface {
	DataDef
	MaxElements() int
	MinElements() int
	Unbounded() bool
}

type HasDataType interface {
	Definition
	HasDefault() bool
	Default() interface{}
	DataType() *DataType
}

type Loader func(parent *Module, name string, rev string) (*Module, error)
