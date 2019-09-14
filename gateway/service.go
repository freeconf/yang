package gateway

import (
	"github.com/freeconf/yang/device"
)

type Service struct {
	log []logEntry
}

type logEntry struct {
	deviceId string
	module   string
	err      error
}

func NewLocalService(d *device.Local) *Service {
	reg := NewLocalRegistrar()
	m := NewFileStore(reg, "var")
	return NewService(d, m, reg)
}

func NewService(d *device.Local, m device.Map, reg Registrar) *Service {
	d.Add("map", device.MapNode(m))
	d.Add("registrar", RegistrarNode(reg))
	return &Service{}
}

func (self *Service) logErr(id string, moduleName string, err error) {
	self.log = append(self.log, logEntry{
		deviceId: id,
		module:   moduleName,
		err:      err,
	})
}
