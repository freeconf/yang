package conf

import (
	"context"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

// every server is required to support ietf-yang-library module
//

type ModuleHandle struct {
	Name      string
	Namespace string
	Revision  string
	Schema    string
	src       meta.StreamSource
	module    *meta.Module
	Submodule map[string]*ModuleHandle
}

func (self *ModuleHandle) Module() (*meta.Module, error) {
	if self.module == nil {
		m, err := yang.LoadModule(self.src, self.Name)
		if err != nil {
			return nil, err
		}
		self.module = m
	}
	return self.module, nil
}

func LoadModules(yangPath meta.StreamSource, remoteYangPath meta.StreamSource, driver node.Node) (map[string]*ModuleHandle, error) {
	yanglib := yang.RequireModule(yangPath, "ietf-yang-library")
	entries := make(map[string]*ModuleHandle)
	n := YangLibModuleList(entries, remoteYangPath)
	b := node.NewBrowser(yanglib, driver)
	if err := b.Root().Find("module-state/module").InsertInto(context.Background(), n).LastErr; err != nil {
		return nil, err
	}
	return entries, nil
}
