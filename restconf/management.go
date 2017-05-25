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

func (self *Management) SetVer(ver string) {
	self.DeviceHandler.Ver = ver
}

func NewManagement(d conf.Device) *Management {
	hndlr := NewDeviceHandler()
	m := &Management{
		Web:           stock.NewHttpServer(hndlr),
		DeviceHandler: hndlr,
	}
	hndlr.ServeDevice(d)
	return m
}
