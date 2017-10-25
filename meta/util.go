package meta

import (
	"strings"

	"github.com/c2stack/c2g/c2"
)

func HasChildren(parent MetaList) bool {
	i := Children(parent)
	return !i.HasNext()
}

// Len counts the items in an iterator
func Len(i Iterator) (len int) {
	for i.HasNext() {
		len++
		i.Next()
	}
	return
}

// GetPath as determined in the information model (e.g. YANG), not data model (e.g. RESTCONF)
func GetPath(m Meta) string {
	s := m.GetIdent()
	if p := m.GetParent(); p != nil {
		return GetPath(p) + "/" + s
	}
	return s
}

// Root finds root meta definition, which is the Module
func Root(m Meta) *Module {
	candidate := m
	for candidate.GetParent() != nil {
		candidate = candidate.GetParent()
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
	sub, found := mod.Imports[subName]
	if !found {
		return nil, "", c2.NewErr("module not found in ident " + ident)
	}
	return sub.Module, ident[i+1:], nil
}
