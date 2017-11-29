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

type HasConditions interface {
	Meta
	Conditions() []Condition
}

// 'when' conditions
type Condition interface {
	Meta
	compilable
	Evaluate() (bool, error)
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

type HasDataType interface {
	Definition
	HasDefault() bool
	Default() interface{}
	DataType() *DataType
}

type Loader func(parent *Module, name string, rev string) (*Module, error)

type FeatureSet interface {
	FeatureOn(*IfFeature) bool
}

type SupportedFeatures struct {
	features map[string]*Feature
	cache    map[string]bool
}

func Whitelist(m *Module, features []string) *SupportedFeatures {
	enabled := make(map[string]*Feature)
	for _, id := range features {
		if f, found := m.Features()[id]; found {
			enabled[id] = f
		}
	}
	return NewSupportedFeatures(enabled)
}

func Backlist(m *Module, features []string) *SupportedFeatures {
	enabled := make(map[string]*Feature)
	for id, f := range m.Features() {
		enabled[id] = f
	}
	for _, j := range features {
		delete(enabled, j)
	}
	return NewSupportedFeatures(enabled)
}

func NewSupportedFeatures(features map[string]*Feature) *SupportedFeatures {
	return &SupportedFeatures{
		features: features,
		cache:    make(map[string]bool),
	}
}

func (self *SupportedFeatures) FeatureOn(f *IfFeature) (bool, error) {
	if on, found := self.cache[f.Expression()]; found {
		return on, nil
	}
	on, err := f.Evaluate(self.features)
	if err != nil {
		return false, err
	}
	self.cache[f.Expression()] = on
	return on, err
}
