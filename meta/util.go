package meta

import (
	"strings"

	"github.com/freeconf/c2g/c2"
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

func externalModule(y Meta, ident string) (*Module, string, error) {
	i := strings.IndexRune(ident, ':')
	if i < 0 {
		return nil, "", nil
	}
	mod := Root(y)
	subName := ident[:i]
	sub, found := mod.imports[subName]
	if !found {
		return nil, "", c2.NewErr("module not found in ident " + ident)
	}
	return sub.module, ident[i+1:], nil
}
