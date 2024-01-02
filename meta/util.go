package meta

import (
	"errors"
	"strings"
)

/*
SchemaPath as determined in the information model (e.g. YANG), not data model (e.g. RESTCONF).

Example:

	    module x  {
		  list foo {
			  leaf bar {
				  ...

		Path on Meta structure for leaf would be 'x/foo/bar'
*/
func SchemaPath(m Meta) string {
	s := ""
	for ; m != nil; m = m.Parent() {
		s = stringPrepend(s, m.(Identifiable).Ident(), "/")
	}
	return s
}

/*
SchemaPathNoModule is like SchemaPath except root module name is not added
*/
func SchemaPathNoModule(m Meta) string {
	s := ""
	for ; m.Parent() != nil; m = m.Parent() {
		s = stringPrepend(s, m.(Identifiable).Ident(), "/")
	}
	return s
}

func stringPrepend(target string, seg string, sep string) string {
	if target == "" {
		return seg
	}
	return seg + sep + target
}

// FindExtension simply finds an extension by name in a list of extensions
func FindExtension(name string, candidates []*Extension) *Extension {
	for _, e := range candidates {
		if e.Ident() == name {
			return e
		}
	}
	return nil
}

// RootModule finds root meta definition, which is the Module
func RootModule(m Meta) *Module {
	candidate := m
	for candidate.Parent() != nil {
		candidate = candidate.Parent()
	}
	return candidate.(*Module)
}

// Module a definition was defined in, not the module it ended up in.
// this is useful for resolving typedefs and uses
func OriginalModule(m Meta) *Module {
	for {
		if mod, isMod := m.(*Module); isMod {
			return mod
		}
		if d, valid := m.(Definition); valid {
			m = d.getOriginalParent()
		} else {
			m = m.Parent()
		}
	}
}

func splitIdent(ident string) (string, string) {
	i := strings.IndexRune(ident, ':')
	if i < 0 {
		return "", ident
	}
	return ident[:i], ident[i+1:]
}

func findModuleAndIsExternal(y Definition, prefix string) (*Module, bool, error) {
	m := OriginalModule(y)
	if prefix == "" || m.Prefix() == prefix {
		return m, false, nil
	}
	sub, found := m.imports[prefix]
	if !found {
		if m.belongsTo != nil && m.belongsTo.prefix == prefix {
			return m.parent.(*Module), true, nil
		}
		return nil, true, errors.New("module not found " + prefix)
	}
	return sub.module, true, nil
}
