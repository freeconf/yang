package automate

import "github.com/c2stack/c2g/device"

type System interface {
	New(role string) (*Handle, error)
}

type Handle struct {
	Id     string
	Device device.Device
	Close  func()
}
