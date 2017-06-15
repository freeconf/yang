package restconf

import (
	"fmt"
	"net/http"

	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/stock"
)

type Auth interface {
	ConstrainRoot(r *http.Request, constraints *node.Constraints) error
}

type Server struct {
	Web           *stock.HttpServer
	DeviceHandler *DeviceHandler
	CallHome      *device.CallHome
}

func NewServer(d *device.Local) *Server {
	hndlr := NewDeviceHandler()
	m := &Server{
		Web:           stock.NewHttpServer(hndlr),
		DeviceHandler: hndlr,
	}
	hndlr.ServeDevice(d)

	// Required by all devices according to RFC
	if err := d.Add("ietf-yang-library", device.LocalDeviceYangLibNode(m.ModuleAddress, d)); err != nil {
		panic(err)
	}
	return m
}

func (self *Server) ModuleAddress(m *meta.Module) string {
	return fmt.Sprint("/restconf/schema/", m.GetIdent(), ".yang")
}

func (self *Server) DeviceAddress(id string, d device.Device) string {
	return fmt.Sprint("/restconf=", id)
}
