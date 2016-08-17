package node

import (
	"fmt"

	"github.com/dhubler/c2g/meta"

	"strings"
)

// Store key values in memory.  Useful for testing or moving temporary data

type BufferStore struct {
	Values  map[string]*Value
	Actions map[string]ActionFunc
	OnSave  func(*BufferStore) error
	OnLoad  func(*BufferStore) error
}

func NewBufferStore() *BufferStore {
	return &BufferStore{
		Values:  make(map[string]*Value, 10),
		Actions: make(map[string]ActionFunc, 10),
	}
}

func (self *BufferStore) Load() error {
	if self.OnLoad != nil {
		return self.OnLoad(self)
	}
	return nil
}

func (self *BufferStore) Clear() error {
	for k, _ := range self.Values {
		delete(self.Values, k)
	}
	return nil
}

func (self *BufferStore) HasValues(path string) bool {
	for k, _ := range self.Values {
		if strings.HasPrefix(k, path) {
			return true
		}
	}
	return false
}

func (self *BufferStore) Save() error {
	if self.OnSave != nil {
		return self.OnSave(self)
	}
	return nil
}

func (self *BufferStore) KeyList(key string, m *meta.List) ([]string, error) {
	builder := NewKeyListBuilder(key)
	for k, _ := range self.Values {
		builder.ParseKey(k)
	}
	return builder.List(), nil
}

func (self *BufferStore) Action(key string) (ActionFunc, error) {
	return self.Actions[key], nil
}

func (self *BufferStore) Value(key string, dataType *meta.DataType) *Value {
	if v, found := self.Values[key]; found {
		v.Type = dataType
		return v
	}
	return nil
}

func (self *BufferStore) SetValue(key string, v *Value) error {
	self.Values[key] = v
	return nil
}

func (self *BufferStore) RemoveAll(path string) error {
	for k, _ := range self.Values {
		if strings.HasPrefix(k, path) {
			delete(self.Values, k)
		}
	}
	for k, _ := range self.Actions {
		if strings.HasPrefix(k, path) {
			delete(self.Actions, k)
		}
	}
	return nil
}

func (self *BufferStore) RenameKey(oldPath string, newPath string) {
	for k, v := range self.Values {
		if strings.HasPrefix(k, oldPath) {
			newKey := fmt.Sprint(newPath, k[len(oldPath):])
			delete(self.Values, k)
			self.Values[newKey] = v
		}
	}
	for k, v := range self.Actions {
		if strings.HasPrefix(k, oldPath) {
			newKey := fmt.Sprint(newPath, k[len(oldPath):])
			delete(self.Actions, k)
			self.Actions[newKey] = v
		}
	}
}
