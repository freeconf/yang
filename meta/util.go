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
	s := m.(Identifiable).Ident()
	if p := m.Parent(); p != nil {
		return SchemaPath(p) + "/" + s
	}
	return s
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
func originalModule(m Definition) *Module {
	for {
		if mod, isMod := m.(*Module); isMod {
			return mod
		}
		m = m.(Definition).getOriginalParent()
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
	m := originalModule(y)
	if prefix == "" || m.Prefix() == prefix {
		return m, false, nil
	}
	sub, found := m.imports[prefix]
	if !found {
		return nil, true, errors.New("module not found " + prefix)
	}
	return sub.module, true, nil
}
