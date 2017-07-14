package device

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

/*
  Utility to read/write changes to a device to a storage system.  This is by no means
  the only way to implement this, but this should work for simple or possibly even some
  complicated use cases.
*/
type Db struct {
	Delegate ProtocolHandler
	IO       DbIO
}

type DbIO interface {
	DbRead(deviceId string, module string, b *node.Browser) error
	DbWrite(deviceId string, module string, b *node.Browser) error
}

func (self *Db) NewDevice(address string) (Device, error) {
	d, err := self.Delegate(address)
	if err != nil {
		return nil, err
	}
	return &dbDevice{
		delegate: d,
		io:       self.IO,
	}, nil
}

type dbDevice struct {
	deviceId string
	delegate Device
	io       DbIO
}

func (self *dbDevice) Id() string {
	return self.deviceId
}

func (self *dbDevice) SchemaSource() meta.StreamSource {
	return self.delegate.SchemaSource()
}

func (self *dbDevice) UiSource() meta.StreamSource {
	return self.delegate.UiSource()
}

func (self *dbDevice) Browser(module string) (*node.Browser, error) {
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
	return node.NewBrowser(b.Meta, DbNode{}.Node(aRoot.Node, bRoot.Node)), nil
}

func (self *dbDevice) Modules() map[string]*meta.Module {
	return self.delegate.Modules()
}

func (self *dbDevice) Close() {
	self.delegate.Close()
}

func (self *dbDevice) storeBrowser(meta meta.MetaList) (*node.Browser, error) {
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
				return self.io.DbWrite(self.deviceId, meta.GetIdent(), browser)
			}
			return nil
		},
	}
	browser = node.NewBrowser(meta, n)
	if err := self.io.DbRead(self.deviceId, meta.GetIdent(), browser); err != nil {
		return nil, err
	}
	loaded = true
	return browser, nil
}
