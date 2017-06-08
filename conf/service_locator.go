package conf

import (
	"github.com/c2stack/c2g/c2"
)

type ServiceLocator interface {
	Device(id string) (Device, error)
	OnUpdate(l DeviceChangeListener) c2.Subscription
	OnModuleUpdate(module string, l DeviceChangeListener) c2.Subscription
}
