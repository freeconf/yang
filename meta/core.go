//go:generate go run core_gen_main.go

package meta

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
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
	deviations    []*Deviation
	imports       map[string]*Import
	includes      []*Include
	identities    map[string]*Identity
	features      map[string]*Feature
	extensionDefs map[string]*ExtensionDef
	featureSet    FeatureSet
	extensions    []*Extension
	belongsTo     *BelongsTo
	configPtr     *bool
}

type BelongsTo struct {
	prefix     string
	moduleName string
	Module     *Module
}

func (b *BelongsTo) Prefix() string {
	return b.prefix
}

func (y *Module) Revision() *Revision {
	if len(y.rev) > 0 {
		return y.rev[0]
	}
	return nil
}

func (y *Module) getOriginalParent() Definition {
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

func (y *Module) Revisions() []*Revision {
	return y.rev
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

func (y *Module) Deviations() []*Deviation {
	return y.deviations
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
	prefix       string
	desc         string
	ref          string
	moduleName   string
	rev          *Revision
	revisionDate string
	parent       *Module
	module       *Module
	loader       Loader
	extensions   []*Extension
}

func (y *Import) Module() *Module {
	return y.module
}

func (y *Import) Prefix() string {
	return y.prefix
}

func (y *Import) RevisionDate() string {
	return y.revisionDate
}

type Include struct {
	subName      string
	rev          *Revision
	revisionDate string
	desc         string
	ref          string
	parent       *Module
	loader       Loader
	extensions   []*Extension
}

func (y *Include) Revision() *Revision {
	return y.rev
}

func (y *Include) RevisionDate() string {
	return y.revisionDate
}

type Choice struct {
	desc           string
	parent         Meta
	originalParent Definition
	ident          string
	ref            string
	when           *When
	configPtr      *bool
	mandatoryPtr   *bool
	defaultVal     *string
	status         Status
	cases          map[string]*ChoiceCase
	ifs            []*IfFeature
	extensions     []*Extension
}

func (y *Choice) addCase(c *ChoiceCase) error {
	if _, exists := y.cases[c.Ident()]; exists {
		return fmt.Errorf("conflict adding add %s to %s. ", c.Ident(), y.Ident())
	}
	y.cases[c.ident] = c
	return nil
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
	ident          string
	desc           string
	ref            string
	parent         Meta
	originalParent Definition
	when           *When
	status         Status
	configPtr      *bool
	dataDefs       []Definition
	dataDefsIndex  map[string]Definition
	ifs            []*IfFeature
	extensions     []*Extension
	recursive      bool
}

// Revision is like a version for a module.  Format is YYYY-MM-DD and should match
// the name of the file on disk when multiple revisions of a file exisits.
type Revision struct {
	parent     Meta
	ident      string
	desc       string
	ref        string
	extensions []*Extension
}

type Container struct {
	parent         Meta
	originalParent Definition
	ident          string
	desc           string
	ref            string
	presence       string
	typedefs       map[string]*Typedef
	groupings      map[string]*Grouping
	actions        map[string]*Rpc
	notifications  map[string]*Notification
	dataDefs       []Definition
	dataDefsIndex  map[string]Definition
	configPtr      *bool
	mandatoryPtr   *bool
	when           *When
	status         Status
	ifs            []*IfFeature
	musts          []*Must
	extensions     []*Extension
	recursive      bool
}

type OrderedBy int

const (
	OrderedBySystem = iota
	OrderedByUser
)

type List struct {
	parent         Meta
	originalParent Definition
	ident          string
	desc           string
	ref            string
	typedefs       map[string]*Typedef
	groupings      map[string]*Grouping
	key            []string
	keyMeta        []Leafable
	orderedBy      OrderedBy
	when           *When
	status         Status
	configPtr      *bool
	mandatoryPtr   *bool
	minElementsPtr *int
	maxElementsPtr *int
	unboundedPtr   *bool
	actions        map[string]*Rpc
	notifications  map[string]*Notification
	dataDefs       []Definition
	dataDefsIndex  map[string]Definition
	ifs            []*IfFeature
	musts          []*Must
	extensions     []*Extension
	unique         [][]string
	recursive      bool
}

func (y *List) KeyMeta() (keyMeta []Leafable) {
	return y.keyMeta
}

type Leaf struct {
	parent         Meta
	originalParent Definition
	ident          string
	desc           string
	ref            string
	units          string
	configPtr      *bool
	mandatoryPtr   *bool
	defaultVal     *string
	dtype          *Type
	when           *When
	status         Status
	ifs            []*IfFeature
	musts          []*Must
	extensions     []*Extension
}

type LeafList struct {
	ident          string
	parent         Meta
	originalParent Definition
	desc           string
	ref            string
	units          string
	configPtr      *bool
	mandatoryPtr   *bool
	dtype          *Type
	minElementsPtr *int
	maxElementsPtr *int
	unboundedPtr   *bool
	orderedBy      OrderedBy
	defaultVals    []string
	when           *When
	status         Status
	ifs            []*IfFeature
	musts          []*Must
	extensions     []*Extension
}

var anyType = newType("any")

type Any struct {
	ident          string
	desc           string
	ref            string
	parent         Meta
	originalParent Definition
	configPtr      *bool
	mandatoryPtr   *bool
	when           *When
	status         Status
	ifs            []*IfFeature
	musts          []*Must
	extensions     []*Extension
}

func (y *Any) HasDefault() bool {
	return false
}

func (y *Any) DefaultValue() interface{} {
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

func (y *Any) addDefault(string) {
	panic("anydata cannot have default value")
}

func (y *Any) setDefaultValue(interface{}) {
	panic("anydata cannot have default value")
}

func (y *Any) clearDefault() {
	panic("anydata cannot have default value")
}

func (y *Any) setUnits(string) {
	panic("anydata cannot have units")
}

/*
*

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
	ident          string
	parent         Meta
	originalParent Definition
	desc           string
	ref            string
	typedefs       map[string]*Typedef
	groupings      map[string]*Grouping
	dataDefs       []Definition
	dataDefsIndex  map[string]Definition
	actions        map[string]*Rpc
	notifications  map[string]*Notification

	// see RFC7950 Sec 14
	// no details (config, mandatory)
	// no when

	extensions []*Extension
}

type Uses struct {
	ident          string
	desc           string
	ref            string
	parent         Meta
	originalParent Definition
	schemaId       interface{}
	refines        []*Refine
	when           *When
	ifs            []*IfFeature
	augments       []*Augment
	extensions     []*Extension
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
	presence       string
	defaultVals    []string
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

type RpcInput struct {
	parent         Meta
	originalParent Definition
	desc           string
	ref            string
	typedefs       map[string]*Typedef
	groupings      map[string]*Grouping
	dataDefs       []Definition
	dataDefsIndex  map[string]Definition
	ifs            []*IfFeature
	musts          []*Must
	extensions     []*Extension
}

func (y *RpcInput) Ident() string {
	return "input"
}

type RpcOutput struct {
	parent         Meta
	originalParent Definition
	desc           string
	ref            string
	typedefs       map[string]*Typedef
	groupings      map[string]*Grouping
	dataDefs       []Definition
	dataDefsIndex  map[string]Definition
	ifs            []*IfFeature
	musts          []*Must
	extensions     []*Extension
}

func (y *RpcOutput) Ident() string {
	return "output"
}

type Rpc struct {
	ident          string
	parent         Meta
	originalParent Definition
	desc           string
	ref            string
	status         Status
	typedefs       map[string]*Typedef
	groupings      map[string]*Grouping
	input          *RpcInput
	output         *RpcOutput
	ifs            []*IfFeature
	extensions     []*Extension
}

func (y *Rpc) Input() *RpcInput {
	return y.input
}

func (y *Rpc) Output() *RpcOutput {
	return y.output
}

type Notification struct {
	ident          string
	parent         Meta
	originalParent Definition
	desc           string
	ref            string
	status         Status
	typedefs       map[string]*Typedef
	groupings      map[string]*Grouping
	dataDefs       []Definition
	dataDefsIndex  map[string]Definition
	ifs            []*IfFeature
	extensions     []*Extension
}

type Typedef struct {
	ident          string
	parent         Meta
	originalParent Definition
	desc           string
	ref            string
	units          string
	defaultVal     *string
	dtype          *Type
	extensions     []*Extension
}

type Augment struct {
	ident          string
	parent         Meta
	originalParent Definition
	desc           string
	ref            string
	actions        map[string]*Rpc
	notifications  map[string]*Notification
	dataDefs       []Definition
	dataDefsIndex  map[string]Definition
	when           *When
	ifs            []*IfFeature
	extensions     []*Extension
}

func (a *Augment) addCase(c *ChoiceCase) error {
	return a.addDataDefinition(c)
}

type AddDeviate struct {
	parent         *Deviation
	configPtr      *bool
	mandatoryPtr   *bool
	maxElementsPtr *int
	minElementsPtr *int
	musts          []*Must
	units          string
	unique         [][]string
	defaultVals    []string
	extensions     []*Extension
}

type ReplaceDeviate struct {
	parent         *Deviation
	dtype          *Type
	units          string
	defaultVals    []string
	configPtr      *bool
	mandatoryPtr   *bool
	minElementsPtr *int
	maxElementsPtr *int
	extensions     []*Extension
}

type DeleteDeviate struct {
	parent      *Deviation
	units       string
	musts       []*Must
	unique      [][]string
	defaultVals []string
	extensions  []*Extension
}

// Deviation is a lot like refine but can be used without
// the "uses" statement and instead right on node tree. This
// is typically used for vendors attempting to implement a
// industry standard YANG model but have some changes
type Deviation struct {
	ident  string // target path
	parent Meta
	desc   string
	ref    string

	// NotSupported is true, then target node can be removed
	NotSupported bool

	// Add properties to target
	Add *AddDeviate

	// Replace properties from target
	Replace *ReplaceDeviate

	// Delete properties from target
	Delete *DeleteDeviate

	extensions []*Extension
}

type Type struct {
	ident           string
	desc            string
	ref             string
	format          val.Format
	enums           []*Enum
	enum            val.EnumList
	bits            []*Bit
	ranges          []*Range
	lengths         []*Range
	patterns        []*Pattern
	path            string
	fractionDigits  int
	delegate        *Type
	base            []string
	identities      []*Identity
	requireInstance bool
	unionTypes      []*Type
	extensions      []*Extension
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

func (y *Type) Patterns() []*Pattern {
	return y.patterns
}

func (y *Type) Format() val.Format {
	return y.format
}

func (y *Type) Path() string {
	return y.path
}

func (y *Type) Bits() []*Bit {
	return y.bits
}

func (y *Type) Enum() val.EnumList {
	return y.enum
}

func (y *Type) Enums() []*Enum {
	return y.enums
}

func (y *Type) Base() []*Identity {
	return y.identities
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

func (y *Type) RequireInstance() bool {
	return y.requireInstance
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

	// merge ranges
	if derived.ranges == nil {
		derived.ranges = base.ranges
	} else if base.ranges != nil {
		derived.ranges = append(derived.ranges, base.ranges...)
	}

	if derived.enums == nil {
		derived.enums = base.enums
	}
	if len(derived.base) == 0 {
		derived.base = base.base
	}
	if derived.unionTypes == nil {
		derived.unionTypes = base.unionTypes
	}

	// merge lengths
	if derived.lengths == nil {
		derived.lengths = base.lengths
	} else if base.lengths != nil {
		derived.lengths = append(derived.lengths, base.lengths...)
	}

	if len(derived.identities) == 0 {
		derived.identities = base.identities
	}
	if derived.fractionDigits == 0 {
		derived.fractionDigits = base.fractionDigits
	}

	// merge bits
	if derived.bits == nil {
		derived.bits = base.bits
	} else if base.bits != nil {
		derived.bits = append(derived.bits, base.bits...)
	}

	derived.format = base.format
}

type Identity struct {
	parent     *Module
	ident      string
	desc       string
	ref        string
	baseIds    []string
	status     Status
	base       []*Identity // normally 1 base, but multiple allowed
	derived    []*Identity
	ifs        []*IfFeature
	extensions []*Extension
}

func (y *Identity) BaseIds() []string {
	return y.baseIds
}

func (y *Identity) Base() []*Identity {
	return y.base
}

func (y *Identity) DerivedDirect() []*Identity {
	return y.derived
}

func FindIdentity(candidates []*Identity, target string) *Identity {
	for _, candidate := range candidates {
		if candidate.ident == target {
			return candidate
		}
		if derived := FindIdentity(candidate.derived, target); derived != nil {
			return derived
		}
	}
	return nil
}

type Feature struct {
	parent     *Module
	ident      string
	desc       string
	ref        string
	status     Status
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
	desc         string
	ref          string
	errorMessage string
	errorAppTag  string
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
//
//	https://tools.ietf.org/html/rfc7950#section-6.3.1
//
// YANG lets you extend everything, including simple statements like
// description.  e.g.
//
//	    container x {
//		       description "X" {
//	           my-ext:this-is-secondary-on-container-keyword-description;
//	        }
//	    }
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
//
//	container x {
//	    foo:bar;
//	}
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

// Definition is define the schema for this extension instance
func (y *Extension) Definition() *ExtensionDef {
	return y.def
}

type Bit struct {
	ident      string
	desc       string
	ref        string
	Position   int
	extensions []*Extension
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

type Range struct {
	desc         string
	ref          string
	errorMessage string
	errorAppTag  string
	extensions   []*Extension
	Entries      []*RangeEntry
}

type RangeEntry struct {
	Min   RangeNumber
	Max   RangeNumber
	Exact RangeNumber
}

func (r *RangeEntry) String() string {
	if !r.Exact.Empty() {
		return r.Exact.String()
	}
	return fmt.Sprintf("%s..%s", r.Min, r.Max)
}

var errNotExpectedValue = errors.New("not expected value")

func (r *RangeEntry) CheckValue(v val.Value) error {
	if !r.Exact.Empty() {
		if cmp, err := r.Exact.Compare(v); err != nil {
			return err
		} else if cmp != 0 {
			return errNotExpectedValue
		}
	}
	if !r.Min.Empty() {
		if cmp, err := r.Min.Compare(v); err != nil {
			return err
		} else if cmp > 0 {
			return errOutsideRange
		}
	}
	if !r.Max.Empty() {
		if cmp, err := r.Max.Compare(v); err != nil {
			return err
		} else if cmp < 0 {
			return errOutsideRange
		}
	}
	return nil
}

type RangeNumber struct {
	str      string
	isMax    bool
	isMin    bool
	integer  *int64
	unsigned *uint64
	float    *float64
}

func (n RangeNumber) Empty() bool {
	return n.str == ""
}

func (n RangeNumber) String() string {
	return n.str
}

func (n RangeNumber) getUnit64() uint64 {
	if n.unsigned != nil {
		return *n.unsigned
	}
	if n.integer != nil && *n.integer >= 0 {
		return uint64(*n.integer)
	}
	if n.float != nil && *n.float >= 0 {
		return uint64(*n.float)
	}
	panic("invalid number range comparison")
}

func (n RangeNumber) getInt64() int64 {
	if n.integer != nil {
		return *n.integer
	}
	if n.float != nil {
		return int64(*n.float)
	}
	panic("invalid number range comparison")
}

func (n RangeNumber) getFloat64() float64 {
	if n.float != nil {
		return *n.float
	}
	if n.integer != nil {
		return float64(*n.integer)
	}
	if n.unsigned != nil {
		return float64(*n.unsigned)
	}
	panic("invalid number range comparison")
}

func (n RangeNumber) Compare(v val.Value) (int64, error) {
	if v.Format().IsList() {
		var cmp0 int64
		var err0 error
		val.ForEach(v, func(index int, item val.Value) {
			cmp, err := n.Compare(item)
			if err != nil {
				err0 = err
			}
			if index == 0 {
				cmp0 = cmp
			} else if cmp != cmp0 && err0 == nil {
				// when comparing and not all values are identical, then result
				// is mixed and therefore invalid. something cannot be both higher
				// and lower than an value.
				err0 = errListItemsRangeVaries
			}
		})
		return cmp0, err0
	} else {
		switch v.Format() {
		case val.FmtDecimal64:
			a := n.getFloat64()
			b := v.Value().(float64)
			if a < b {
				return -1, nil
			}
			if a > b {
				return 1, nil
			}
			return 0, nil
		case val.FmtUInt64:
			a := n.getUnit64()
			b := v.Value().(uint64)
			if a < b {
				return -1, nil
			}
			if a > b {
				return 1, nil
			}
			return 0, nil
		default:
			if i, ok := v.(val.Int64able); ok {
				a := n.getInt64()
				b := i.Int64()
				if a < b {
					return -1, nil
				}
				if a > b {
					return 1, nil
				}
				return 0, nil
			}
		}
	}
	// hard to imagine YANG would allow range on non-numerical but error if here
	return 0, fmt.Errorf("cannot do a numerical comparison on type %s", v.Format().String())
}

func newRangeNumber(s string) (RangeNumber, error) {
	s = strings.TrimSpace(s)
	n := RangeNumber{str: s}
	if s == "max" {
		n.isMax = true
	} else if s == "min" {
		n.isMin = true
	} else if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		n.integer = &i
	} else if u, err := strconv.ParseUint(s, 10, 64); err == nil {
		n.unsigned = &u
	} else if f, err := strconv.ParseFloat(s, 64); err == nil {
		n.float = &f
	} else {
		return n, fmt.Errorf("unrecognized number '%s'", s)
	}
	return n, nil
}

func (r *Range) Empty() bool {
	return len(r.Entries) == 0
}

func newRange(encoded string) (*Range, error) {
	var err error
	r := &Range{}
	sranges := strings.Split(encoded, "|")
	for _, srange := range sranges {
		e := &RangeEntry{}
		segments := strings.Split(string(srange), "..")
		if len(segments) == 2 {
			if e.Min, err = newRangeNumber(segments[0]); err != nil {
				return nil, err
			}
			if e.Max, err = newRangeNumber(segments[1]); err != nil {
				return nil, err
			}
		} else {
			if e.Exact, err = newRangeNumber(srange); err != nil {
				return nil, err
			}
		}
		r.Entries = append(r.Entries, e)
	}
	return r, nil
}

var errOutsideRange = errors.New("value is outside all allowed ranges")
var errListItemsRangeVaries = errors.New("values in list vary on both inside and outside of comparison")

func (r *Range) CheckValue(v val.Value) error {
	if len(r.Entries) == 0 {
		return nil
	}
	for _, e := range r.Entries {
		if err := e.CheckValue(v); err == nil {
			return nil
		}
	}
	return errOutsideRange
}

func (r *Range) String() string {
	var s string
	for i, e := range r.Entries {
		if i != 0 {
			s = s + "|"
		}
		s = s + e.String()
	}
	return s
}

// This is a start but I think the ideal solution collapses a list of
// ranges by looking at overlapping areas while validating each range
// is more restrictive as per RFC7950 Sec. 9.2.4 the "range" statement
//
// func MergeRanges(l []Range) []Range {
// 	return append(l, r)
// }
//
// func (a Range) Merge(b Range) Range {
// 	r := Range{
// 		notNil: true,
// 	}
// 	if a.Min == "" || a.min < b.min {
// 		r.min = b.min
// 		r.Min = b.Min
// 	} else {
// 		r.min = a.min
// 		r.Min = a.Min
// 	}
// 	if a.Max == "" || a.max > b.max {
// 		r.max = b.max
// 		r.Max = b.Max
// 	} else {
// 		r.max = a.max
// 		r.Max = a.Max
// 	}
// 	return r
// }

// Pattern is used to confine string types to specific regular expression
// values
type Pattern struct {
	desc         string
	ref          string
	Pattern      string
	errorMessage string
	errorAppTag  string
	inverted     bool
	extensions   []*Extension
	regex        *regexp.Regexp
}

func newPattern(pattern string) (*Pattern, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &Pattern{
		Pattern: pattern,
		regex:   r,
	}, nil
}

func (p *Pattern) CheckValue(s string) bool {
	return p.regex.MatchString(s) != p.inverted
}

func (p *Pattern) Inverted() bool {
	return p.inverted
}
