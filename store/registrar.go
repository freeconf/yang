// +build ignore

package store

import (
	"fmt"
	"time"

	"github.com/c2stack/c2g/browse"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

type ConfigStoreSource interface {
	ConfigStore(endpoint *Endpoint) (ConfigStore, error)
}

type ConfigStore interface {
	ConfigNode(path string) (node.Node, error)
	SaveConfig() error
}

type Registrar struct {
	YangPath             meta.StreamSource
	Endpoints            map[string]*Endpoint
	SchemaInsertionPoint *meta.Choice
	StoreDir             string
	TxSource             ConfigStoreSource
	ClientSource         browse.ClientSource
}

func NewRegistrar(ypath meta.StreamSource, schemaInsertionPoint *meta.Choice) *Registrar {
	return &Registrar{
		YangPath: ypath,
	}
}

func (self *Registrar) RegisterEndpoint(endpoint *Endpoint) error {
	if len(endpoint.Id) == 0 {

		// TODO: use real UUID
		uuid := fmt.Sprintf("%s.%d", endpoint.Module, time.Now().UnixNano())

		endpoint.Id = uuid
	}

	if module, err := endpoint.Schema(); err != nil {
		return err
	} else {
		kase := &meta.ChoiceCase{Ident: module.GetIdent()}
		self.SchemaInsertionPoint.AddMeta(kase)
		kase.AddMeta(module)
		endpoint.TxSource = self.TxSource
	}
	if self.Endpoints == nil {
		self.Endpoints = make(map[string]*Endpoint)
	}
	self.Endpoints[endpoint.Id] = endpoint

	return nil
}
