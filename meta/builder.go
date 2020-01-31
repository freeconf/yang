package meta

import (
	"fmt"
	"strings"
)

type Builder struct {
	LastErr error
	uid     int
	//module  *Module
}

func (b *Builder) setErr(err error) {
	if b.LastErr == nil {
		b.LastErr = err
	}
}

func (b *Builder) Module(ident string, fs FeatureSet) *Module {
	return &Module{
		ident:         ident,
		ver:           "1",
		featureSet:    fs,
		imports:       make(map[string]*Import),
		groupings:     make(map[string]*Grouping),
		typedefs:      make(map[string]*Typedef),
		identities:    make(map[string]*Identity),
		features:      make(map[string]*Feature),
		extensionDefs: make(map[string]*ExtensionDef),
	}
}

func (b *Builder) Description(o interface{}, desc string) {
	d, valid := o.(Describable)
	if !valid {
		b.setErr(fmt.Errorf("%T does not allow descriptions", o))
	} else {
		d.setDescription(desc)
	}
}

func (b *Builder) Presence(o interface{}, desc string) {
	d, valid := o.(*Container)
	if !valid {
		b.setErr(fmt.Errorf("%T does not allow presence. Only containers", o))
	} else {
		d.presence = desc
	}
}

func (b *Builder) Reference(o interface{}, r string) {
	d, valid := o.(Describable)
	if !valid {
		b.setErr(fmt.Errorf("%T does not allow reference", o))
	} else {
		d.setReference(r)
	}
}

func (b *Builder) Namespace(o interface{}, ns string) {
	m, valid := o.(*Module)
	if !valid {
		b.setErr(fmt.Errorf("%T does not allow namespace, only modules do", o))
	} else {
		m.namespace = ns
	}
}

func (b *Builder) Prefix(o interface{}, prefix string) {
	switch x := o.(type) {
	case *Module:
		x.prefix = prefix
	case *Import:
		x.prefix = prefix
	default:
		b.setErr(fmt.Errorf("%T does not allow prefix, only modules or imports do", o))
	}
}

func (b *Builder) YinElement(o interface{}, prop bool) {
	def, valid := o.(*ExtensionDefArg)
	if !valid {
		b.setErr(fmt.Errorf("%T does not allow yim elements, only extension arguments do", o))
	} else {
		def.yinElement = prop
	}
}

func (b *Builder) Revision(o interface{}, rev string) *Revision {
	r := Revision{
		ident: rev,
	}
	switch x := o.(type) {
	case *Module:
		r.parent = x
		r.scope = x
		x.rev = append(x.rev, &r)
	case *Import:
		r.parent = x
		r.scope = x
		x.rev = &r
	case *Include:
		r.parent = x
		r.scope = x
		x.rev = &r
	default:
		b.setErr(fmt.Errorf("%T does not allow revisions, only modules, imports or includes do", o))
	}
	return &r
}

func (b *Builder) Import(o interface{}, moduleName string, loader Loader) *Import {
	i := Import{
		moduleName: moduleName,
		loader:     loader,
	}
	parent, valid := o.(*Module)
	if !valid {
		b.setErr(fmt.Errorf("%T does not allow imports, only modules do", o))
	} else {
		i.parent = parent
		if parent.imports == nil {
			parent.imports = make(map[string]*Import)
		}
		parent.imports[i.moduleName] = &i
	}
	return &i
}

func (b *Builder) Include(o interface{}, subName string, loader Loader) *Include {
	i := Include{
		subName: subName,
		loader:  loader,
	}
	parent, valid := o.(*Module)
	if !valid {
		b.setErr(fmt.Errorf("%T does not allow imports, only modules do", o))
	} else {
		i.parent = parent
		parent.includes = append(parent.includes, &i)
	}
	return &i
}

func (b *Builder) YangVersion(o interface{}, ver string) {
	m, valid := o.(*Module)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support yang version, only modules do", o))
	} else {
		m.ver = ver
	}
}

func (b *Builder) Grouping(o interface{}, ident string) *Grouping {
	g := Grouping{
		ident: ident,
	}
	h, valid := o.(HasGroupings)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support groupings", o))
	} else {
		h.addGrouping(&g)
	}

	return &g
}

func (b *Builder) Organization(o interface{}, org string) {
	m, valid := o.(*Module)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support organization, only modules do", o))
	} else {
		m.org = org
	}
}

func (b *Builder) Contact(o interface{}, c string) {
	m, valid := o.(*Module)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support contact, only modules do", o))
	} else {
		m.contact = c
	}
}

func (b *Builder) Units(o interface{}, units string) {
	m, valid := o.(HasUnits)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support units", o))
	} else {
		m.setUnits(units)
	}
}

func (b *Builder) AddExtension(o interface{}, keyword string, ext *Extension) {
	ext.keyword = keyword
	m, valid := o.(HasExtensions)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support extensions", o))
	} else {
		m.addExtension(ext)
	}
}

func (b *Builder) Extension(prefixAndIdent string, args []string) *Extension {
	ids := strings.Split(prefixAndIdent, ":")
	return &Extension{
		prefix: ids[0],
		ident:  ids[1],
		args:   args,
	}
}

func (b *Builder) ExtensionDef(o interface{}, ident string) *ExtensionDef {
	e := ExtensionDef{
		ident: ident,
	}
	m, valid := o.(*Module)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support extension definitions, only modules do", o))
	} else {
		e.parent = m
		m.extensionDefs[ident] = &e
	}
	return &e
}

func (b *Builder) ExtensionDefArg(o interface{}, ident string) *ExtensionDefArg {
	arg := ExtensionDefArg{
		ident: ident,
	}
	d, valid := o.(*ExtensionDef)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support extension definition arguments, only modules do", o))
	} else {
		arg.parent = d
		d.args = append(d.args, &arg)
	}
	return &arg
}

func (b *Builder) Feature(o interface{}, ident string) *Feature {
	f := Feature{
		ident: ident,
	}
	m, valid := o.(*Module)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support features, only modules do", o))
	} else {
		f.parent = m
		m.features[ident] = &f
	}
	return &f
}

func (b *Builder) Must(o interface{}, expression string) *Must {
	m := Must{
		expr: expression,
	}
	h, valid := o.(HasMusts)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support must", o))
	} else {
		h.addMust(&m)
		m.scopedParent = m.parent
	}
	return &m
}

func (b *Builder) IfFeature(o interface{}, expression string) *IfFeature {
	i := IfFeature{
		expr: expression,
	}
	h, valid := o.(HasIfFeatures)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support if-feature", o))
	} else {
		h.addIfFeature(&i)
	}
	return &i
}

func (b *Builder) When(o interface{}, expression string) *When {
	w := When{
		expr: expression,
	}
	h, valid := o.(HasWhen)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support when", o))
	} else {
		h.setWhen(&w)
	}
	return &w
}

func (b *Builder) Identity(o interface{}, ident string) *Identity {
	i := Identity{
		ident: ident,
	}
	m, valid := o.(*Module)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support identities, only modules do", o))
	} else {
		i.parent = m
		m.identities[i.ident] = &i
	}
	return &i
}

// Base  to set for identity or type objects
func (b *Builder) Base(o interface{}, base string) {
	i, valid := o.(*Identity)
	if valid {
		i.derivedIds = append(i.derivedIds, base)
	} else if t, isType := o.(*Type); isType {
		t.base = base
	} else {
		b.setErr(fmt.Errorf("%T does not support base, only identities do", o))
	}
}

func (b *Builder) Augment(o interface{}, path string) *Augment {
	a := Augment{
		ident: path,
	}
	if x, valid := o.(HasAugments); !valid {
		b.setErr(fmt.Errorf("%T does not allow augments", o))
	} else {
		x.addAugments(&a)
	}
	return &a
}

func (b *Builder) Refine(o interface{}, path string) *Refine {
	r := Refine{
		ident: path,
	}
	u, valid := o.(*Uses)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support refine, only uses does", o))
	} else {
		r.parent = u
		u.refines = append(u.refines, &r)
	}
	return &r
}

func (b *Builder) Deviation(o interface{}, ident string) *Deviation {
	m, valid := o.(*Module)
	d := Deviation{
		ident: ident,
	}
	if !valid {
		b.setErr(fmt.Errorf("%T does not allow deviation, only modules do", o))
	} else {
		d.parent = m
		m.deviations = append(m.deviations, &d)
	}
	return &d
}

func (b *Builder) NotSupported(o interface{}) {
	d, valid := o.(*Deviation)
	if !valid {
		b.setErr(fmt.Errorf("%T does not allow deviation, only modules do", o))
	} else {
		d.NotSupported = true
	}
}

func (b *Builder) AddDeviate(o interface{}) *AddDeviate {
	var add AddDeviate
	d, valid := o.(*Deviation)
	if !valid {
		b.setErr(fmt.Errorf("%T does not allow deviate, only deviations do", o))
	} else {
		d.Add = &add
	}
	return &add
}

func (b *Builder) ReplaceDeviate(o interface{}) *ReplaceDeviate {
	var x ReplaceDeviate
	d, valid := o.(*Deviation)
	if !valid {
		b.setErr(fmt.Errorf("%T does not allow deviate, only deviations do", o))
	} else {
		d.Replace = &x
	}
	return &x
}

func (b *Builder) DeleteDeviate(o interface{}) *DeleteDeviate {
	var x DeleteDeviate
	d, valid := o.(*Deviation)
	if !valid {
		b.setErr(fmt.Errorf("%T does not allow deviate, only deviations do", o))
	} else {
		d.Delete = &x
	}
	return &x
}

func (b *Builder) Uses(o interface{}, ident string) *Uses {
	x := Uses{
		ident: ident,
	}
	if h, valid := b.parentDataDefinition(o, ident); valid {
		x.parent = h
		x.scope = h
		x.schemaId = b.uid
		b.uid++
		h.addDataDefinition(&x)
	}
	// anything unique
	//x.schemaId = &x
	return &x
}

func (b *Builder) Container(o interface{}, ident string) *Container {
	x := Container{
		ident: ident,
	}
	if h, valid := b.parentDataDefinition(o, ident); valid {
		x.parent = h
		x.scope = h
		h.addDataDefinition(&x)
	}
	return &x
}

func (b *Builder) List(o interface{}, ident string) *List {
	x := List{
		ident: ident,
	}
	if h, valid := b.parentDataDefinition(o, ident); valid {
		x.parent = h
		x.scope = h
		h.addDataDefinition(&x)
	}
	return &x
}

func (b *Builder) Key(o interface{}, keys string) {
	i, valid := o.(*List)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support key, only lists do", o))
	} else {
		i.key = strings.Split(keys, " ")
	}
}

func (b *Builder) Leaf(o interface{}, ident string) *Leaf {
	x := Leaf{
		ident: ident,
	}
	if h, valid := b.parentDataDefinition(o, ident); valid {
		x.parent = h
		x.scope = h
		h.addDataDefinition(&x)
	}
	return &x
}

func (b *Builder) LeafList(o interface{}, ident string) *LeafList {
	x := LeafList{
		ident: ident,
	}
	if h, valid := b.parentDataDefinition(o, ident); valid {
		x.parent = h
		x.scope = h
		h.addDataDefinition(&x)
	}
	return &x
}

func (b *Builder) Any(o interface{}, ident string) *Any {
	x := Any{
		ident: ident,
	}
	if h, valid := b.parentDataDefinition(o, ident); valid {
		x.parent = h
		x.scope = h
		h.addDataDefinition(&x)
	}
	return &x
}

func (b *Builder) Choice(o interface{}, ident string) *Choice {
	x := Choice{
		ident: ident,
		cases: make(map[string]*ChoiceCase),
	}
	if h, valid := b.parentDataDefinition(o, ident); valid {
		x.parent = h
		x.scope = h
		h.addDataDefinition(&x)
	}
	return &x
}

func (b *Builder) Case(o interface{}, ident string) *ChoiceCase {
	x := ChoiceCase{
		ident: ident,
	}
	h, valid := o.(*Choice)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support case definitions, only choice does", o))
	} else {
		x.parent = h
		x.scope = h
		h.cases[ident] = &x
	}
	return &x
}

func (b *Builder) parentDataDefinition(o interface{}, ident string) (HasDataDefinitions, bool) {
	if _, parentIsChoice := o.(*Choice); parentIsChoice {
		// YANG 1.1 allows implicit "case" statements so we detect and inject a case
		// statement here if relevant
		return b.Case(o, ident), true
	}
	h, valid := o.(HasDataDefinitions)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support data definitions", o))
	}
	return h, valid
}

func (b *Builder) Action(o interface{}, ident string) *Rpc {
	x := Rpc{
		ident: ident,
	}
	h, valid := o.(HasActions)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support actions", o))
	} else {
		x.parent = h
		x.scope = h
		h.addAction(&x)
	}
	return &x
}

func (b *Builder) ActionInput(o interface{}) *RpcInput {
	x := RpcInput{}
	h, valid := o.(*Rpc)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support action input, only rpc or action does", o))
	} else {
		x.parent = h
		x.scope = h
		h.input = &x
	}
	return &x
}

func (b *Builder) ActionOutput(o interface{}) *RpcOutput {
	x := RpcOutput{}
	h, valid := o.(*Rpc)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support action output, only rpc or action does", o))
	} else {
		x.parent = h
		x.scope = h
		h.output = &x
	}
	return &x
}

func (b *Builder) Notification(o interface{}, ident string) *Notification {
	x := Notification{
		ident: ident,
	}
	h, valid := o.(HasNotifications)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support action output, only rpc or action does", o))
	} else {
		x.parent = h
		x.scope = h
		h.addNotification(&x)
	}
	return &x
}

func (b *Builder) Config(o interface{}, config bool) {
	i, valid := o.(HasConfig)
	if valid {
		i.setConfig(config)
	} else {
		b.setErr(fmt.Errorf("%T does not support config", o))
	}
}

func (b *Builder) Mandatory(o interface{}, m bool) {
	i, valid := o.(HasMandatory)
	if valid {
		i.setMandatory(m)
	} else {
		b.setErr(fmt.Errorf("%T does not support mandatory", o))
	}
}

func (b *Builder) MinElements(o interface{}, i int) {
	h, valid := o.(HasMinMax)
	if valid {
		h.setMinElements(i)
	} else {
		b.setErr(fmt.Errorf("%T does not support list details", o))
	}
}

func (b *Builder) OrderedBy(o interface{}, order OrderedBy) {
	h, valid := o.(HasOrderedBy)
	if valid {
		h.setOrderedBy(order)
	} else {
		b.setErr(fmt.Errorf("%T does not support ordered-by", o))
	}
}

func (b *Builder) MaxElements(o interface{}, i int) {
	h, valid := o.(HasMinMax)
	if valid {
		h.setMaxElements(i)
	} else {
		b.setErr(fmt.Errorf("%T does not support list details", o))
	}
}

func (b *Builder) UnBounded(o interface{}, x bool) {
	i, valid := o.(HasUnbounded)
	if valid {
		i.setUnbounded(x)
	} else {
		b.setErr(fmt.Errorf("%T does not support list details", o))
	}
}

func (b *Builder) Default(o interface{}, defaultVal interface{}) {
	h, valid := o.(HasDefault)
	if valid {
		h.setDefault(defaultVal)
	} else {
		b.setErr(fmt.Errorf("%T does not support list details", o))
	}
}

func (b *Builder) Typedef(o interface{}, ident string) *Typedef {
	t := Typedef{
		ident: ident,
	}
	h, valid := o.(HasTypedefs)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support typedefs", o))
	} else {
		t.parent = h.(Meta)
		h.addTypedef(&t)
	}
	return &t
}

func (b *Builder) RequireInstance(o interface{}, require bool) {
	i, valid := o.(*Type)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support path, only type does", o))
	} else {
		i.requireInstance = require
	}
}

func (b *Builder) Path(o interface{}, p string) {
	i, valid := o.(*Type)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support path, only type does", o))
	} else {
		i.path = p
	}
}

func (b *Builder) Pattern(o interface{}, pattern string) *Pattern {
	p := Pattern{
		Pattern: pattern,
	}
	i, valid := o.(*Type)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support pattern, only type does", o))
	} else {
		i.patterns = append(i.patterns, &p)
	}
	return &p
}

func (b *Builder) Enum(o interface{}, label string) *Enum {
	e := Enum{
		ident: label,
	}
	i, valid := o.(*Type)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support pattern, only type does", o))
	} else {
		i.enums = append(i.enums, &e)
	}
	return &e
}

func (b *Builder) Bit(o interface{}, ident string) *Bit {
	e := Bit{
		ident: ident,
	}
	i, valid := o.(*Type)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support bit, only type does", o))
	} else {
		i.bits = append(i.bits, &e)
	}
	return &e
}

func (b *Builder) Position(o interface{}, x int) {
	i, valid := o.(*Bit)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support position, only type bit", o))
	} else {
		i.Position = x
	}
}

func (b *Builder) ErrorMessage(o interface{}, msg string) {
	i, valid := o.(HasErrorMessage)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support error-messages", o))
	} else {
		i.setErrorMessage(msg)
	}
}

func (b *Builder) ErrorAppTag(o interface{}, tag string) {
	i, valid := o.(HasErrorMessage)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support error-app-tag", o))
	} else {
		i.setErrorAppTag(tag)
	}
}

func (b *Builder) EnumValue(o interface{}, x int) {
	i, valid := o.(*Enum)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support value, only type enum", o))
	} else {
		i.val = x
	}
}

func (b *Builder) FractionDigits(o interface{}, x int) {
	i, valid := o.(*Type)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support fraction digits, only type does", o))
	} else {
		i.fractionDigits = x
	}
}

func (b *Builder) Type(o interface{}, ident string) *Type {
	t := newType(ident)
	if x, valid := o.(HasType); valid {
		x.setType(t)
	} else if x, valid := o.(*Type); valid {
		x.unionTypes = append(x.unionTypes, t)
	} else {
		b.setErr(fmt.Errorf("%T does not support types", o))
	}
	return t
}

func (b *Builder) LengthRange(o interface{}, arg string) *Range {
	r, err := newRange(arg)
	if err != nil {
		b.setErr(err)
		return r
	}
	i, valid := o.(*Type)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support length range, only type does", o))
	} else {
		i.lengths = append(i.lengths, r)
	}
	return r
}

func (b *Builder) ValueRange(o interface{}, arg string) *Range {
	r, err := newRange(arg)
	if err != nil {
		b.setErr(err)
		return r
	}
	i, valid := o.(*Type)
	if !valid {
		b.setErr(fmt.Errorf("%T does not support value range, only type does", o))
	} else {
		i.ranges = append(i.ranges, r)
	}
	return r
}
