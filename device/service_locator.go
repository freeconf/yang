package device

import (
	"github.com/freeconf/c2g/c2"
)

type ServiceLocator interface {
	Device(id string) (Device, error)
	OnUpdate(l ChangeListener) c2.Subscription
	OnModuleUpdate(module string, l ChangeListener) c2.Subscription
}
