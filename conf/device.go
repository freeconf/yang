package conf

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

// Create device from address string associated with protocol
// often referred to south/east/west bound
type Client interface {
	NewDevice(address string) (Device, error)
}

// Export device by it's address so protocol server can serve a device
// often referred to northbound
type Server interface {
	DeviceAddress(id string, d Device) string
}

type Device interface {
	SchemaSource() meta.StreamSource
	UiSource() meta.StreamSource
	Browser(module string) (*node.Browser, error)
	Modules() map[string]*meta.Module
	Close()
}
