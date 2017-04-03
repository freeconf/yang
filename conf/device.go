package conf

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

type Device interface {
	SchemaSource() meta.StreamSource
	UiSource() meta.StreamSource
	Browser(module string) (*node.Browser, error)
	ModuleHandles() (map[string]*ModuleHandle, error)
	Close()
}

type LocalDevice struct {
	browsers     map[string]*node.Browser
	schemaSource meta.StreamSource
	uiSource     meta.StreamSource
}

func NewLocalDevice(schemaSource meta.StreamSource) *LocalDevice {
	return &LocalDevice{
		schemaSource: schemaSource,
		browsers:     make(map[string]*node.Browser),
	}
}

func NewLocalDeviceWithUi(schemaSource meta.StreamSource, uiSource meta.StreamSource) *LocalDevice {
	return &LocalDevice{
		schemaSource: schemaSource,
		uiSource:     uiSource,
		browsers:     make(map[string]*node.Browser),
	}
}

func (self *LocalDevice) SchemaSource() meta.StreamSource {
	return self.schemaSource
}

func (self *LocalDevice) UiSource() meta.StreamSource {
	return self.uiSource
}

func (self *LocalDevice) ModuleHandles() (map[string]*ModuleHandle, error) {
	mods := make(map[string]*ModuleHandle)
	for _, b := range self.browsers {
		m := b.Meta.(*meta.Module)
		mods[m.GetIdent()] = &ModuleHandle{
			Name:      m.GetIdent(),
			Namespace: m.Namespace,
			Revision:  m.Revision.Ident,
			Schema:    m.GetIdent(),
		}
	}
	return mods, nil
}

func (self *LocalDevice) Browser(module string) (*node.Browser, error) {
	return self.browsers[module], nil
}

func (self *LocalDevice) Close() {
}

func (self *LocalDevice) Add(module string, n node.Node) error {
	m, err := yang.LoadModule(self.schemaSource, module)
	if err != nil {
		return err
	}
	self.browsers[module] = node.NewBrowser(m, n)
	return nil
}

func (self *LocalDevice) AddBrowser(b *node.Browser) {
	self.browsers[b.Meta.GetIdent()] = b
}
