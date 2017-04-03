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

type Management struct {
	Web           *stock.HttpServer
	DeviceHandler *DeviceHandler
}

func NewManagement(d conf.Device, port string) *Management {
	hndlr := NewDeviceHandler()
	m := &Management{
		Web:           stock.NewHttpServer(hndlr),
		DeviceHandler: hndlr,
	}
	hndlr.ServeDevice(d)
	options := m.Web.Options()
	options.Port = port
	m.Web.ApplyOptions(options)
	return m
}
