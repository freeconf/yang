//go:generate go run core_gen_main.go

package meta

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/freeconf/yang/val"
)

// Module is top-most container of the information model. It's name
// does not appear in data model.
type Module struct {
	ident         string
	namespace     string
	prefix        string
	desc          string
	contact       string
	org           string
	ref           string
	ver           string
	rev           []*Revision
	parent        Meta // non-null for submodules and imports
	dataDefs      []Definition
	dataDefsIndex map[string]Definition
	notifications map[string]*Notification
	actions       map[string]*Rpc
	typedefs      map[string]*Typedef
	groupings     map[string]*Grouping
	augments      []*Augment
	imports       map[string]*Import
	includes      []*Include
	identities    map[string]*Identity
	features      map[string]*Feature
	extensionDefs map[string]*ExtensionDef
	featureSet    FeatureSet
	extensions    []*Extension
}

func NewModule(ident string, featureSet FeatureSet) *Module {
	return &Module{
		ident:         ident,
		ver:           "1",
		featureSet:    featureSet,
		imports:       make(map[string]*Import),
		groupings:     make(map[string]*Grouping),
		typedefs:      make(map[string]*Typedef),
		identities:    make(map[string]*Identity),
		features:      make(map[string]*Feature),
		extensionDefs: make(map[string]*ExtensionDef),
	}
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

func (y *Module) FeatureSet() FeatureSet {
	return y.featureSet
}

func (y *Module) ExtensionDefs() map[string]*ExtensionDef {
	return y.extensionDefs
}

func (y *Module) addExtensionDef(d *ExtensionDef) {
	y.extensionDefs[d.Ident()] = d
}

func (y *Module) ModuleByPrefix(prefix string) (*Module, error) {
	if y.Prefix() == prefix {
		return y, nil
	}
	i, found := y.imports[prefix]
	if !found {
		return nil, fmt.Errorf("cannot find module with prefix %s", prefix)
	}
	return i.module, nil
}

type Import struct {
	prefix     string
	desc       string
	ref        string
	moduleName string
	rev        *Revision
	parent     *Module
	module     *Module
	loader     Loader
	extensions []*Extension
}

func (y *Import) Module() *Module {
	return y.module
}

func (y *Import) Prefix() string {
	return y.prefix
}

type Include struct {
	subName    string
	rev        *Revision
	desc       string
	ref        string
	parent     *Module
	loader     Loader
	extensions []*Extension
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

type Choice struct {
	description string
	parent      Meta
	scope       Meta
	ident       string
	desc        string
	ref         string
	when        *When
	configPtr   *bool
	mandatory   bool
	defaultVal  interface{}
	status      Status
	cases       map[string]*ChoiceCase
	ifs         []*IfFeature
	extensions  []*Extension
}

func (y *Choice) Cases() map[string]*ChoiceCase {
	return y.cases
}

func (y *Choice) CaseIdents() []string {
	idents := make([]string, 0, len(y.cases))
	for ident := range y.cases {
		idents = append(idents, ident)

	}
	sort.Strings(idents)
	return idents
}

type ChoiceCase struct {
	ident         string
	desc          string
	ref           string
	parent        Meta
	scope         Meta
	when          *When
	dataDefs      []Definition
	dataDefsIndex map[string]Definition
	ifs           []*IfFeature
	extensions    []*Extension
	recursive     bool
}

// Revision is like a version for a module.  Format is YYYY-MM-DD and should match
// the name of the file on disk when multiple revisions of a file exisits.
type Revision struct {
	parent     Meta
	scope      Meta
	ident      string
	desc       string
	ref        string
	extensions []*Extension
}

type Container struct {
	ident         string
	desc          string
	ref           string
	typedefs      map[string]*Typedef
	groupings     map[string]*Grouping
	actions       map[string]*Rpc
	notifications map[string]*Notification
	dataDefs      []Definition
	dataDefsIndex map[string]Definition
	parent        Meta
	scope         Meta
	status        Status
	configPtr     *bool
	mandatory     bool
	when          *When
	ifs           []*IfFeature
	musts         []*Must
	extensions    []*Extension
	recursive     bool
}

type List struct {
	parent        Meta
	scope         Meta
	ident         string
	desc          string
	ref           string
	typedefs      map[string]*Typedef
	groupings     map[string]*Grouping
	key           []string
	keyMeta       []HasType
	when          *When
	configPtr     *bool
	mandatory     bool
	minElements   int
	maxElements   int
	unboundedPtr  *bool
	actions       map[string]*Rpc
	notifications map[string]*Notification
	dataDefs      []Definition
	dataDefsIndex map[string]Definition
	ifs           []*IfFeature
	musts         []*Must
	extensions    []*Extension
	recursive     bool
}

func (y *List) KeyMeta() (keyMeta []HasType) {
	return y.keyMeta
}

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
	extensions []*Extension
}

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
	defaultVal   interface{}
	when         *When
	ifs          []*IfFeature
	musts        []*Must
	extensions   []*Extension
}

var anyType = newType("any")

type Any struct {
	ident      string
	desc       string
	ref        string
	parent     Meta
	scope      Meta
	configPtr  *bool
	mandatory  bool
	when       *When
	ifs        []*IfFeature
	musts      []*Must
	extensions []*Extension
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

func (y *Any) setType(*Type) {
	panic("cannot set type on an any type")
}

func (y *Any) Type() *Type {
	return anyType
}

func (y *Any) setDefault(interface{}) {
	panic("anydata cannot have default value")
}

func (y *Any) setUnits(string) {
	panic("anydata cannot have units")
}

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
	ident         string
	parent        Meta
	desc          string
	ref           string
	typedefs      map[string]*Typedef
	groupings     map[string]*Grouping
	dataDefs      []Definition
	dataDefsIndex map[string]Definition
	actions       map[string]*Rpc
	notifications map[string]*Notification

	// see RFC7950 Sec 14
	// no details (config, mandatory)
	// no when

	extensions []*Extension
}

type Uses struct {
	ident      string
	desc       string
	ref        string
	parent     Meta
	scope      Meta
	schemaId   interface{}
	refines    []*Refine
	when       *When
	ifs        []*IfFeature
	augments   []*Augment
	extensions []*Extension
}

func (y *Uses) Refinements() []*Refine {
	return y.refines
}

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
	extensions     []*Extension
}

func (y *Refine) splitIdent() (string, string) {
	slash := strings.IndexRune(y.ident, '/')
	if slash < 0 {
		return y.ident, ""
	}
	return y.ident[:slash], y.ident[slash+1:]
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

type RpcInput struct {
	parent        Meta
	scope         Meta
	desc          string
	ref           string
	typedefs      map[string]*Typedef
	groupings     map[string]*Grouping
	dataDefs      []Definition
	dataDefsIndex map[string]Definition
	ifs           []*IfFeature
	musts         []*Must
	extensions    []*Extension
}

func (y *RpcInput) Ident() string {
	return "input"
}

type RpcOutput struct {
	parent        Meta
	scope         Meta
	desc          string
	ref           string
	typedefs      map[string]*Typedef
	groupings     map[string]*Grouping
	dataDefs      []Definition
	dataDefsIndex map[string]Definition
	ifs           []*IfFeature
	musts         []*Must
	extensions    []*Extension
}

func (y *RpcOutput) Ident() string {
	return "output"
}

type Rpc struct {
	ident      string
	parent     Meta
	scope      Meta
	desc       string
	ref        string
	typedefs   map[string]*Typedef
	groupings  map[string]*Grouping
	input      *RpcInput
	output     *RpcOutput
	ifs        []*IfFeature
	extensions []*Extension
}

func (y *Rpc) Input() *RpcInput {
	return y.input
}

func (y *Rpc) Output() *RpcOutput {
	return y.output
}

type Notification struct {
	ident         string
	parent        Meta
	scope         Meta
	desc          string
	ref           string
	typedefs      map[string]*Typedef
	groupings     map[string]*Grouping
	dataDefs      []Definition
	dataDefsIndex map[string]Definition
	ifs           []*IfFeature
	extensions    []*Extension
}

type Typedef struct {
	ident      string
	parent     Meta
	desc       string
	ref        string
	units      string
	defaultVal interface{}
	dtype      *Type
	extensions []*Extension
}

type Augment struct {
	ident         string
	parent        Meta
	desc          string
	ref           string
	actions       map[string]*Rpc
	notifications map[string]*Notification
	dataDefs      []Definition
	dataDefsIndex map[string]Definition
	when          *When
	ifs           []*IfFeature
	extensions    []*Extension
}

type Type struct {
	ident   string
	desc    string
	ref     string
	format  val.Format
	enums   []*Enum
	enum    val.EnumList
	ranges  []*Range
	lengths []*Range
	path    string
	//units               string
	fractionDigits int
	patterns       []string
	//defaultVal          interface{}
	delegate   *Type
	base       string
	identity   *Identity
	unionTypes []*Type
	extensions []*Extension
}

func newType(ident string) *Type {
	return &Type{
		ident: ident,
	}
}

func (y *Type) Range() []*Range {
	return y.ranges
}

func (y *Type) Length() []*Range {
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

// Resolve is the effective datatype if this type points to a different
// dataType, which is the case for leafRefs.  Otherwise this just returns
// itself
func (y *Type) Resolve() *Type {
	if y.delegate == nil {
		panic("no delegate")
	}
	return y.delegate
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
	derived.format = base.format
}

type Identity struct {
	parent     *Module
	ident      string
	desc       string
	ref        string
	derivedIds []string
	derived    map[string]*Identity
	ifs        []*IfFeature
	extensions []*Extension
}

func (y *Identity) BaseIds() []string {
	return y.derivedIds
}

func (y *Identity) Identities() map[string]*Identity {
	return y.derived
}

type Feature struct {
	parent     *Module
	ident      string
	desc       string
	ref        string
	ifs        []*IfFeature
	extensions []*Extension
}

type IfFeature struct {
	parent     Meta
	expr       string
	extensions []*Extension
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
		return false, errors.New("syntax err in feature expression:" + y.expr)
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
		y.lastErr = errors.New("syntax err in feature expression:" + y.expr)
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

type When struct {
	parent     Meta
	expr       string
	desc       string
	ref        string
	extensions []*Extension
}

func (y *When) Expression() string {
	return y.expr
}

type Must struct {
	parent       Meta
	scopedParent Meta
	expr         string
	extensions   []*Extension
}

func (y *Must) Expression() string {
	return y.expr
}

type ExtensionDef struct {
	parent *Module
	ident  string
	desc   string
	ref    string
	status Status
	args   []*ExtensionDefArg

	// yes, even extension dataDefsIndex can have extensions
	extensions []*Extension
}

func (y *ExtensionDef) Arguments() []*ExtensionDefArg {
	return y.args
}

type ExtensionDefArg struct {
	parent     *ExtensionDef
	ident      string
	desc       string
	ref        string
	yinElement bool
	extensions []*Extension
}

func (y *ExtensionDefArg) YinElement() bool {
	return y.yinElement
}

// Extension is a very powerful concept in YANG.  It let's you extend YANG
// language to have defintions for whatever you wish it had. It's like a meta language
// inside the YANG (which is already a meta language). Can extensions have extensions?
// you bet.  See YANG RFC on extensions for more information.
//    https://tools.ietf.org/html/rfc7950#section-6.3.1
//
// YANG lets you extend everything, including simple statements like
// description.  e.g.
//     container x {
//	       description "X" {
//            my-ext:this-is-secondary-on-container-keyword-description;
//         }
//     }
//
// This extension would be listed in the extensions for the *Container object
// but would have OnKeyword of "description" to distinguish it from extensions
// extension of container itself.
type Extension struct {
	parent  Meta
	prefix  string
	ident   string
	keyword string
	def     *ExtensionDef
	args    []string

	// yes even extensions can have extensions
	extensions []*Extension
}

// Prefix name of extention which according to YANG spec is ALWAYS required even
// when the extension definition is local.  In this example it is "foo"
//  container x {
//      foo:bar;
//  }
func (y *Extension) Prefix() string {
	return y.prefix
}

// Keyword is set when there are extensions of things that do not
// have a data structure to store them and a likely just Go scalar
// values.  Examples: description, reference, units, max-length
// etc.
func (y *Extension) Keyword() string {
	return y.keyword
}

// Arguments are optional argumes to extension.  The extension definition will
// define what arguments are allowed if any.
func (y *Extension) Arguments() []string {
	return y.args
}

type Enum struct {
	ident      string
	desc       string
	ref        string
	val        int
	extensions []*Extension
}

func (y *Enum) Value() int {
	return y.val
}
