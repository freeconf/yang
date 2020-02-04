package meta

import (
	"errors"
	"fmt"

	"github.com/freeconf/yang/val"
)

// responsiblities: ensuring all the definitions are valid when considered
// all together.

func Compile(root *Module) error {
	c := &compiler{
		root: root,
	}
	// resolve uses with groupings
	if err := resolve(root); err != nil {
		return err
	}
	return c.module(root)
}

type compiler struct {
	root *Module
}

func (c *compiler) module(y *Module) error {
	if y.featureSet != nil {
		if err := y.featureSet.Initialize(y); err != nil {
			return err
		}
	}
	for _, i := range y.identities {
		if err := c.compile(i); err != nil {
			return err
		}
	}

	for _, r := range y.rev {
		if err := c.compile(r); err != nil {
			return err
		}
	}

	return c.compile(y)
}

func (c *compiler) compile(o interface{}) error {

	if x, ok := o.(HasTypedefs); ok {
		for _, y := range x.Typedefs() {
			if err := c.compile(y); err != nil {
				return err
			}
		}
	}
	if x, ok := o.(HasType); ok {
		if err := c.compileType(x.Type(), x.(Leafable)); err != nil {
			return err
		}
		if err := c.compile(x.Type()); err != nil {
			return err
		}
	}

	if x, ok := o.(HasConfig); ok {
		if !x.IsConfigSet() {
			x.setConfig(c.inheritConfig(x.(Meta).Parent()))
		}
	}

	switch x := o.(type) {
	case *Extension:
		if err := c.extension(x); err != nil {
			return err
		}
	case *Typedef:
		if err := c.typedef(x); err != nil {
			return err
		}
	case *Identity:
		if err := c.identity(x); err != nil {
			return err
		}
	case *Rpc:
		if x.input != nil {
			if err := c.compile(x.input); err != nil {
				return err
			}
		}
		if x.output != nil {
			if err := c.compile(x.output); err != nil {
				return err
			}
		}
	case *List:
		if err := c.list(x); err != nil {
			return err
		}
	case *Choice:
		for _, k := range x.Cases() {
			if err := c.compile(k); err != nil {
				return err
			}
		}
	}

	if x, ok := o.(HasDataDefinitions); ok {
		if !x.IsRecursive() {
			for _, y := range x.DataDefinitions() {
				if err := c.compile(y); err != nil {
					return err
				}
			}
		}
	}
	if x, ok := o.(HasActions); ok {
		for _, y := range x.Actions() {
			if err := c.compile(y); err != nil {
				return err
			}
		}
	}
	if x, ok := o.(HasNotifications); ok {
		for _, y := range x.Notifications() {
			if err := c.compile(y); err != nil {
				return err
			}
		}
	}
	if x, ok := o.(Meta); ok {
		for _, y := range x.Extensions() {
			if err := c.compile(y); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *compiler) inheritConfig(m Meta) bool {
	if x, ok := m.(HasDetails); ok {
		if !x.IsConfigSet() {
			x.setConfig(c.inheritConfig(x.(Meta).Parent()))
			//panic(fmt.Sprintf("%s (%T)", SchemaPath(m), x))
		}
		return x.Config()
	}
	return true
}

func (c *compiler) list(y *List) error {
	y.keyMeta = make([]Leafable, len(y.key))
	for i, keyIdent := range y.key {
		// relies on resolver happening first
		km, valid := y.dataDefsIndex[keyIdent]
		if !valid {
			return fmt.Errorf("%s - %s key not found", SchemaPath(y), keyIdent)
		}
		y.keyMeta[i], valid = km.(Leafable)
		if !valid {
			return fmt.Errorf("%s - %s expected key with data type", SchemaPath(y), keyIdent)
		}
	}
	return nil
}

func (c *compiler) extension(e *Extension) error {
	target, err := c.root.ModuleByPrefix(e.Prefix())
	if err != nil {
		return err
	}
	var found bool
	e.def, found = target.extensionDefs[e.Ident()]
	if !found {
		return fmt.Errorf("could not find extension definition for extension %s:%s", e.Prefix(), e.Ident())
	}

	// TODO: check args of extension match the allowed args of the definition
	return nil
}

func (c *compiler) identity(y *Identity) error {
	if y.derived != nil {
		// already done
		return nil
	}
	y.derived = make(map[string]*Identity)

	// add yourself to list
	y.derived[y.ident] = y

	// find all the derived identities
	for _, baseId := range y.derivedIds {
		m := y.parent
		prefix, ident := splitIdent(baseId)
		m, _, err := findModuleAndIsExternal(m, prefix)
		if err != nil {
			return err
		}
		identity, found := m.Identities()[ident]
		if !found {
			return errors.New(SchemaPath(y) + " - " + baseId + " identity not found")
		}
		y.derived[baseId] = identity
		if err := c.compile(identity); err != nil {
			return err
		}
		// copy in all the other identities because they all become part of
		// the same identity set
		for subId, subIdent := range identity.Identities() {
			y.derived[subId] = subIdent
		}
	}
	return nil
}

func (c *compiler) compileType(y *Type, parent Leafable) error {
	if y == nil {
		return errors.New("no type set on " + SchemaPath(parent))
	}
	if int(y.format) != 0 {
		return nil
	}
	var hasTypedef bool
	y.format, hasTypedef = val.TypeAsFormat(y.ident)
	if !hasTypedef {
		tdef, err := c.findTypedef(y, parent, y.ident)
		if err != nil {
			return err
		}

		// Don't use resolve here because if a typedef is a leafref, you want
		// the unresolved here and resolve it below
		tdef.dtype.mixin(y)

		if !parent.HasDefault() {
			parent.setDefault(tdef.Default())
		}
		if parent.Units() == "" {
			parent.setUnits(tdef.Units())
		}
	}

	if y.format == val.FmtLeafRef || y.format == val.FmtLeafRefList {
		if y.path == "" {
			return errors.New(SchemaPath(parent) + " - " + y.ident + " path is required")
		}
		// parent is a leaf, so start with parent's parent which is a container-ish
		resolvedMeta := Find(parent, y.path)
		if resolvedMeta == nil {
			// eat err as this will be rather common until leafref parsing improves
			// err := errors.New(SchemaPath(parent) + " - " + y.typeIdent + " could not resolve leafref path " + y.path)
			y.delegate = y
		} else {
			y.delegate = resolvedMeta.(HasType).Type()
		}
	} else {
		y.delegate = y
	}

	if y.format == val.FmtIdentityRef {
		prefix, ident := splitIdent(y.base)
		m, _, err := findModuleAndIsExternal(parent, prefix)
		if err != nil {
			return err
		}
		identity, found := m.Identities()[ident]
		if !found {
			return errors.New(SchemaPath(parent) + " - " + y.base + " identity not found")
		}
		y.identity = identity
	}

	if _, isList := parent.(*LeafList); isList && !y.format.IsList() {
		y.format = val.Format(int(y.format) + 1024)
	}

	if y.format == val.FmtUnion {
		if len(y.unionTypes) == 0 {
			return errors.New(SchemaPath(parent) + " - unions need at least one type")
		}
		for _, u := range y.unionTypes {
			if err := c.compileType(u, parent); err != nil {
				return err
			}
		}
	} else if len(y.unionTypes) > 0 {
		return errors.New(SchemaPath(parent) + " - embedded types are only for union types")
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
				Label: item.ident,
			}
			nextId++
		}
	}

	if y.format == val.FmtBits || y.format == val.FmtBitsList {
		nextPos := 0
		for _, item := range y.bits {
			if item.Position > 0 {
				nextPos = item.Position
			} else {
				item.Position = nextPos
			}
			nextPos++
		}
	}

	return nil
}

func (c *compiler) findTypedef(y *Type, parent Definition, qualifiedIdent string) (*Typedef, error) {
	prefix, ident := splitIdent(qualifiedIdent)

	// From RFC
	//   A reference to an unprefixed type or grouping, or one that uses the
	//   prefix of the current module, is resolved by locating the matching
	//   "typedef" or "grouping" statement among the immediate substatements
	//   of each ancestor statement.
	// this means if prefix is local module, then ignore it and follow chain
	searchHeirarcy := (prefix == "")
	var module *Module
	if !searchHeirarcy {
		m, isExternal, err := findModuleAndIsExternal(parent, prefix)
		if err != nil {
			return nil, err
		}
		if !isExternal {
			searchHeirarcy = true
		} else {
			module = m
		}
	}

	var found *Typedef
	if searchHeirarcy {
		p := parent
		for p != nil {
			if ptd, ok := p.(HasTypedefs); ok {
				if found = ptd.Typedefs()[ident]; found != nil {
					break
				}
			}
			p = p.getOriginalParent()
		}
	} else {
		found = module.Typedefs()[ident]
	}

	if found == nil {
		return nil, errors.New(SchemaPath(parent) + " - typedef " + y.ident + " not found")
	}

	// this will recurse if typedef references another typedef
	if err := c.compile(found); err != nil {
		return nil, err
	}

	return found, nil
}

func (c *compiler) typedef(t *Typedef) error {
	if t.dtype == nil {
		return fmt.Errorf("%s - %s type required", SchemaPath(t), t.ident)
	}
	return nil
}
