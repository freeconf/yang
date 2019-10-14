package meta

import (
	"errors"
	"strings"
)

// GetPath as determined in the information model (e.g. YANG), not data model (e.g. RESTCONF)
func GetPath(m Meta) string {
	s := m.(Identifiable).Ident()
	if p := m.Parent(); p != nil {
		return GetPath(p) + "/" + s
	}
	return s
}

// Root finds root meta definition, which is the Module
func Root(m Meta) *Module {
	candidate := m
	for candidate.Parent() != nil {
		candidate = candidate.Parent()
	}
	return candidate.(*Module)
}

// Given ident
//   foo:bar
// return the module names foo and the string "bar"
// Given
//    bar
// return the local module and return back "bar"
func rootByIdent(y Meta, ident string) (*Module, string, error) {
	if c, ok := y.(cloneable); ok {
		y = c.scopedParent()
	}
	mod := Root(y)
	i := strings.IndexRune(ident, ':')
	if i < 0 {
		return mod, ident, nil
	}
	subName := ident[:i]
	sub, found := mod.imports[subName]
	if !found {
		return nil, "", errors.New("module not found in ident " + ident)
	}
	return sub.module, ident[i+1:], nil
}

// Given ident
//   foo:bar
// return the module names foo and the string "bar"
// Given
//    bar
// return nil as this is not an external module
func externalModule(y Meta, ident string) (*Module, string, error) {
	i := strings.IndexRune(ident, ':')
	if i < 0 {
		return nil, "", nil
	}
	if c, ok := y.(cloneable); ok {
		y = c.scopedParent()
	}
	mod := Root(y)
	subName := ident[:i]
	sub, found := mod.imports[subName]
	if !found {
		return nil, "", errors.New("module not found in ident " + ident)
	}
	return sub.module, ident[i+1:], nil
}
