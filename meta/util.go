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
	mod := RootModule(y)
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
	mod := RootModule(y)
	subName := ident[:i]
	sub, found := mod.imports[subName]
	if !found {
		return nil, "", errors.New("module not found in ident " + ident)
	}
	return sub.module, ident[i+1:], nil
}

// FindDefinition can return action, notification or any of the data definitions
// like container, leaf, list etc.
// func Def(parent interface{}, ident string) Definition {
// 	if x, ok := parent.(HasDataDefinitions); ok {
// 		if def := x.Definition(ident); def != nil {
// 			return def
// 		}
// 	}
// 	if x, ok := parent.(HasActions); ok {
// 		if def, found := x.Actions()[ident]; found {
// 			return def
// 		}
// 	}
// 	if x, ok := parent.(HasNotifications); ok {
// 		if def, found := x.Notifications()[ident]; found {
// 			return def
// 		}
// 	}
// 	return nil
// }
