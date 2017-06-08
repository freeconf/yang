package conf

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

type Store struct {
	YangPath meta.StreamSource
	Delegate Client
	Support  StoreSupport
}

type StoreSupport interface {
	LoadStore(deviceId string, module string, b *node.Browser) error
	SaveStore(deviceId string, module string, b *node.Browser) error
}

func (self *Store) NewDevice(address string) (Device, error) {
	d, err := self.Delegate.NewDevice(address)
	if err != nil {
		return nil, err
	}
	return &storeDevice{
		delegate: d,
		support:  self.Support,
	}, nil
}

type storeDevice struct {
	deviceId string
	delegate Device
	support  StoreSupport
}

func (self *storeDevice) Id() string {
	return self.deviceId
}

func (self *storeDevice) SchemaSource() meta.StreamSource {
	return self.delegate.SchemaSource()
}

func (self *storeDevice) UiSource() meta.StreamSource {
	return self.delegate.UiSource()
}

func (self *storeDevice) Browser(module string) (*node.Browser, error) {
	b, err := self.delegate.Browser(module)
	if err != nil {
		return nil, err
	}
	bRoot := b.Root()
	a, err := self.storeBrowser(bRoot.Meta().(meta.MetaList))
	if err != nil {
		return nil, err
	}
	aRoot := a.Root()
	return node.NewBrowser(b.Meta, StoreNode{}.Node(aRoot.Node, bRoot.Node)), nil
}

func (self *storeDevice) Modules() map[string]*meta.Module {
	return self.delegate.Modules()
}

func (self *storeDevice) Close() {
	self.delegate.Close()
}

func (self *storeDevice) storeBrowser(meta meta.MetaList) (*node.Browser, error) {
	// avoid recursive read -> save -> read ... by not saving on first load
	var loaded bool

	var browser *node.Browser

	// where all data is stored, we actually never r/w to this directly, only thru
	// return browser object
	data := make(map[string]interface{})

	n := &node.Extend{
		Node: node.MapNode(data),
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			// trigger save on edit
			if err := p.BeginEdit(r); err != nil {
				return err
			}
			if loaded {
				return self.support.SaveStore(self.deviceId, meta.GetIdent(), browser)
			}
			return nil
		},
	}
	browser = node.NewBrowser(meta, n)
	if err := self.support.LoadStore(self.deviceId, meta.GetIdent(), browser); err != nil {
		return nil, err
	}
	loaded = true
	return browser, nil
}
