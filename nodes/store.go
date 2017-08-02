package nodes

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/val"
)

type Store interface {
	Load() error
	Save() error
	HasValues(path string) bool
	Value(path string, typ *meta.DataType) val.Value
	SetValue(path string, v val.Value) error
	KeyList(path string, m *meta.List) ([]string, error)
	RenameKey(oldPath string, newPath string)
	Action(path string) (ActionFunc, error)
	RemoveAll(path string) error
}
