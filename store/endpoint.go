// +build ignore

package store

import (
	"bytes"

	"github.com/c2stack/c2g/browse"
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

type Endpoint struct {
	YangPath        meta.StreamSource
	Id              string
	Module          string
	Meta            *meta.Module
	EndpointAddress string
	TxSource        ConfigStoreSource
	ClientSource    browse.ClientSource
	basePath        *node.Path
}

func (self *Endpoint) Schema() (*meta.Module, error) {
	if self.Meta != nil {
		return self.Meta, nil
	}
	in, err := self.getRequest("meta/?noexpand")
	if err != nil {
		return nil, err
	}
	m := &meta.Module{}
	if err = node.SelectModule(self.YangPath, m, false).Root().UpsertFrom(in).LastErr; err != nil {
		return nil, err
	}
	self.Meta = m
	return self.Meta, err
}

func (self *Endpoint) handleRequest(target node.PathSlice) (node.Node, error) {
	self.basePath = target.Head
	var path string
	if !target.Empty() {
		path = target.String()
	}
	operational, err := self.getRequest("restconf/" + path + "?content=nonconfig")
	if err != nil && !c2.IsNotFoundErr(err) {
		return nil, err
	}
	tx, createTxErr := self.TxSource.ConfigStore(self)
	if createTxErr != nil {
		return nil, createTxErr
	}
	config, beginTxErr := tx.ConfigNode(path)
	if beginTxErr != nil && !c2.IsNotFoundErr(beginTxErr) {
		return nil, beginTxErr
	}
	proxy := &proxy{
		stripPathPrefix: target.Head.String(),
		onCommit:        tx.SaveConfig,
		onRequest:       self.request,
	}
	prxy := proxy.proxy(config, operational)
	if len(path) == 0 {
		return prxy, nil
	}
	return navigate(target.Tail.StringNoModule(), prxy), nil
}

func (self *Endpoint) pushConfig() error {
	tx, createTxErr := self.TxSource.ConfigStore(self)
	if createTxErr != nil {
		return createTxErr
	}
	localConfig, beginTxErr := tx.ConfigNode("")
	if beginTxErr != nil && !c2.IsNotFoundErr(beginTxErr) {
		return beginTxErr
	}
	var payload bytes.Buffer
	payloadNode := node.NewJsonWriter(&payload).Node()
	if err := node.NewBrowser(self.Meta, localConfig).Root().InsertInto(payloadNode).LastErr; err != nil {
		return err
	}
	if _, err := self.request("PUT", "restconf/", &payload); err != nil {
		return err
	}
	return nil
}

func (self *Endpoint) pullConfig() error {
	remoteConfig, err := self.getRequest("restconf/?content=config")
	if err != nil {
		return err
	}
	tx, createTxErr := self.TxSource.ConfigStore(self)
	if createTxErr != nil {
		return createTxErr
	}
	localConfig, beginTxErr := tx.ConfigNode("")
	if beginTxErr != nil && !c2.IsNotFoundErr(beginTxErr) {
		return beginTxErr
	}
	if err := node.NewBrowser(self.Meta, localConfig).Root().UpsertFrom(remoteConfig).LastErr; err != nil {
		return err
	}
	return tx.SaveConfig()
}
