package node

import (
	"reflect"
	"sort"
)

// Index help you implement Node.OnNext when you have a map to interate over
type Index struct {
	Keys       []reflect.Value
	comparator func(a, b reflect.Value) bool
}

func NewIndex(mmap interface{}) *Index {
	mapVal := reflect.ValueOf(mmap)
	index := &Index{
		Keys: mapVal.MapKeys(),
	}
	return index
}

func (self *Index) Sort(comparator func(a, b reflect.Value) bool) {
	self.comparator = comparator
	sort.Sort(self)
}

func (self *Index) Len() int {
	return len(self.Keys)
}

func (self *Index) Swap(i, j int) {
	self.Keys[i], self.Keys[j] = self.Keys[j], self.Keys[i]
}

func (self *Index) Less(i, j int) bool {
	return self.comparator(self.Keys[i], self.Keys[j])
}

var NO_VALUE reflect.Value

func (self *Index) NextKey(row int) reflect.Value {
	if row < len(self.Keys) {
		return self.Keys[row]
	}

	return NO_VALUE
}
