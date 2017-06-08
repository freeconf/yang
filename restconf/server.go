package restconf

import (
	"net/http"

	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/stock"
)

type Auth interface {
	ConstrainRoot(r *http.Request, constraints *node.Constraints) error
}

type Server struct {
	Web           *stock.HttpServer
	DeviceHandler *DeviceHandler
	CallHome      *conf.CallHome
}

func NewServer(d *conf.LocalDevice) *Server {
	hndlr := NewDeviceHandler()
	m := &Server{
		Web:           stock.NewHttpServer(hndlr),
		DeviceHandler: hndlr,
	}
	hndlr.ServeDevice(d)

	// Required by all devices according to RFC
	if err := d.Add("ietf-yang-library", conf.LocalDeviceYangLibNode(d)); err != nil {
		panic(err)
	}
	return m
}

func (self *Server) DeviceAddress(id string, d conf.Device) string {
	if id == "" {
		return self.DeviceHandler.BaseAddress
	}
	return self.DeviceHandler.BaseAddress + "=" + id
}
