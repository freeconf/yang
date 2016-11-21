package node

import (
	"reflect"

	"github.com/c2stack/c2g/meta"
)

// Uses reflection to marshal data into go structs.  If you want to override
// then use:
//     data.Extend{
//         Node: ReflectNode(obj),
//         OnSelect:...
//     }
func ReflectNode(Obj interface{}) Node {
	s := &MyNode{
		Label:    "MarshalContainer " + reflect.TypeOf(Obj).Name(),
		Peekable: Obj,
	}
	s.OnSelect = func(r ContainerRequest) (Node, error) {
		objVal := reflect.ValueOf(Obj)
		if objVal.Kind() == reflect.Interface || objVal.Kind() == reflect.Ptr {
			objVal = objVal.Elem()
		}
		fieldName := meta.MetaNameToFieldName(r.Meta.GetIdent())
		value := objVal.FieldByName(fieldName)
		if meta.IsList(r.Meta) {
			if value.Kind() == reflect.Map {
				marshal := &MarshalMap{
					Map: value.Interface(),
				}
				return marshal.Node(), nil
			} else {
				marshal := &MarshalArray{
					ArrayValue: &value,
				}
				return marshal.Node(), nil
			}
		} else {
			if value.Kind() == reflect.Struct {
				return ReflectNode(value.Addr().Interface()), nil
			} else if value.CanAddr() {
				var child interface{}
				if r.New {
					childValue := reflect.New(value.Type().Elem())
					value.Set(childValue)
					child = childValue.Interface()
				} else {
					if value.IsNil() {
						return nil, nil
					}
					child = value.Interface()
				}
				if child != nil {
					return ReflectNode(child), nil
				}
				return nil, nil
			}
		}
		return nil, nil
	}
	s.OnField = func(r FieldRequest, hnd *ValueHandle) (err error) {
		if r.Write {
			err = WriteField(r.Meta, Obj, hnd.Val)
		} else {
			hnd.Val, err = ReadField(r.Meta, Obj)
		}
		return
	}
	return s
}

type MarshalArray struct {
	ArrayValue   *reflect.Value
	OnNewItem    func() interface{}
	OnSelectItem func(item interface{}, index int) Node
}

func (self *MarshalArray) Node() Node {
	n := &MyNode{
		Label:    "MarshalArray " + self.ArrayValue.Type().Name(),
		Peekable: self.ArrayValue.Interface(),
	}
	n.OnNext = func(r ListRequest) (next Node, key []*Value, err error) {
		var item interface{}
		var index int
		if r.New {
			var itemValue reflect.Value
			if self.OnNewItem != nil {
				item = self.OnNewItem()
				itemValue = reflect.ValueOf(item)
			} else {
				itemValue = reflect.New(self.ArrayValue.Type().Elem().Elem())
				item = itemValue.Interface()
			}
			self.ArrayValue.Set(reflect.Append(*self.ArrayValue, itemValue))
			index = self.ArrayValue.Len() - 1
		} else if len(r.Key) > 0 {
			// TODO: Should we hook into OnKey ?  Can default assume key in
			// array index which is, i feel, 90% of use cases?
			panic("Keys only implemented on MarshalMap, not MarshalArray")
		} else {
			if r.Row < self.ArrayValue.Len() {
				itemValue := self.ArrayValue.Index(r.Row)
				if itemValue.CanAddr() {
					// If we don't pass pointer, we will make edits on a copy
					item = itemValue.Addr().Interface()
				} else {
					item = itemValue.Interface()
				}
				index = r.Row
			}
		}
		if item != nil {
			if self.OnSelectItem != nil {
				return self.OnSelectItem(item, index), nil, nil
			}
			return ReflectNode(item), nil, nil
		}
		return nil, nil, nil
	}
	return n
}

type MarshalMap struct {
	Map          interface{}
	OnNewItem    func(r ListRequest) interface{}
	OnSelectItem func(item interface{}) Node
}

func (self *MarshalMap) Node() Node {
	mapReflect := reflect.ValueOf(self.Map)
	n := &MyNode{
		Label:    "MarshalMap " + mapReflect.Type().Name(),
		Peekable: self.Map,
	}
	index := NewIndex(self.Map)
	n.OnNext = func(r ListRequest) (next Node, key []*Value, err error) {
		var item interface{}
		key = r.Key
		if r.New {
			item = self.OnNewItem(r)
			mapKey := reflect.ValueOf(r.Key[0].Value())
			mapReflect.SetMapIndex(mapKey, reflect.ValueOf(item))
		} else if len(r.Key) > 0 {
			mapKey := reflect.ValueOf(r.Key[0].Value())
			itemVal := mapReflect.MapIndex(mapKey)
			if itemVal.IsValid() {
				item = itemVal.Interface()
			}
		} else {
			nextKey := index.NextKey(r.Row)
			if nextKey != NO_VALUE {
				key = SetValues(r.Meta.KeyMeta(), nextKey.Interface())
				itemVal := mapReflect.MapIndex(nextKey)
				item = itemVal.Interface()
			}
		}
		if item != nil {
			if self.OnSelectItem != nil {
				return self.OnSelectItem(item), key, nil
			}
			return ReflectNode(item), key, nil
		}
		return nil, nil, nil
	}
	return n
}
