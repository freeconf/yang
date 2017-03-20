package conf

import "github.com/c2stack/c2g/meta"

// handles:
//   register and unregister
//   download yang and cache it
//   delegates request to remote nodes as if they were local
//   works w/store so store can persist data
type Proxy struct {
	devices  map[string]Device
	factory  ProtocolHandler
	yangPath meta.StreamSource
}

type ProtocolHandler func(yangPath meta.StreamSource, address string, port string) (Device, error)

func NewProxy(yangPath meta.StreamSource, proto ProtocolHandler) *Proxy {
	return &Proxy{
		factory:  proto,
		yangPath: yangPath,
		devices:  make(map[string]Device),
	}
}

func (self *Proxy) Register(id string, address string, port string) error {
	d, err := self.factory(self.yangPath, address, port)
	if err != nil {
		return err
	}
	self.devices[id] = d
	return nil
}
