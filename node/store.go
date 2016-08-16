package node

import (
	"github.com/dhubler/c2g/meta"
)

type Store interface {
	Load() error
	Save() error
	HasValues(path string) bool
	Value(path string, typ *meta.DataType) *Value
	SetValue(path string, v *Value) error
	KeyList(path string, m *meta.List) ([]string, error)
	RenameKey(oldPath string, newPath string)
	Action(path string) (ActionFunc, error)
	RemoveAll(path string) error
}
