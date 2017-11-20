package device

import (
	"github.com/freeconf/c2g/meta"
	"github.com/freeconf/c2g/node"
)

// Create device from address string associated with protocol
// often referred to south/east/west bound
type ProtocolHandler func(addr string) (Device, error)

type Device interface {
	SchemaSource() meta.StreamSource
	UiSource() meta.StreamSource
	Browser(module string) (*node.Browser, error)
	Modules() map[string]*meta.Module
	Close()
}
