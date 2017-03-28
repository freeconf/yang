package conf

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

// every server is required to support ietf-yang-library module
//

type ModuleHandle struct {
	Name            string
	Namespace       string
	Revision        string
	Schema          string
	Feature         []string
	ConformanceType string
	Submodule       map[string]*ModuleHandle
}

func LoadModules(yangPath meta.StreamSource, driver node.Node) (map[string]*ModuleHandle, error) {
	yanglib := yang.RequireModule(yangPath, "ietf-yang-library")
	entries := make(map[string]*ModuleHandle)
	n := YangLibModuleList(entries)
	b := node.NewBrowser(yanglib, driver)
	if err := b.Root().Find("modules-state/module").InsertInto(n).LastErr; err != nil {
		return nil, err
	}
	return entries, nil
}
