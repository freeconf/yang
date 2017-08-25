package nodes

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"
)

// Uses reflection to marshal data into go structs.  If you want to override
// then use:
//     data.Extend{
//         Node: Reflect{}.Node(obj),
//         OnChild:...
//     }

type reflect2 struct {
}

func Reflect(obj interface{}) node.Node {
	return reflect2{}.child(reflect.ValueOf(obj))
}

func ReflectList(obj interface{}) node.Node {
	return reflect2{}.list(reflect.ValueOf(obj), nil)
}

func (self reflect2) isEmpty(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}
	switch v.Type().Kind() {
	case reflect.Struct:
		return false
	default:
		return v.IsNil()
	}
	return false
}

func (self reflect2) child(v reflect.Value) node.Node {
	if self.isEmpty(v) {
		return nil
	}
	switch v.Kind() {
	case reflect.Map:
		return self.childMap(v)
	case reflect.Slice:
	case reflect.Array:
	case reflect.Interface:
		switch v.Elem().Type().Kind() {
		case reflect.Map:
			return self.childMap(v.Elem())
		}
		return self.strukt(v.Elem())
	case reflect.Ptr:
		return self.strukt(v)
	case reflect.Struct:
		return self.strukt(v.Addr())
	}
	panic("unsupported type for child container " + v.String())
}

func (self reflect2) list(v reflect.Value, onUpdate onListValueChange) node.Node {
	if self.isEmpty(v) {
		return nil
	}
	switch v.Kind() {
	case reflect.Map:
		return self.listMap(v)
	case reflect.Interface:
		switch v.Elem().Kind() {
		case reflect.Slice:
			return self.listSlice(v.Elem(), onUpdate)
		case reflect.Map:
			return self.listMap(v.Elem())
		}
	case reflect.Slice:
		return self.listSlice(v, onUpdate)
	}
	panic("unsupported type for listing " + v.String())
}

type sliceEntry struct {
	key []val.Value
	n   node.Node
	pos int
}

type sliceSorter []sliceEntry

func (self sliceSorter) Less(a, b int) bool {
	return val.CompareVals(self[a].key, self[b].key) < 0
}

func (self sliceSorter) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self sliceSorter) Len() int {
	return len(self)
}

func (self sliceSorter) findFunc(key []val.Value) func(int) bool {
	return func(i int) bool {
		return val.CompareVals(self[i].key, key) >= 0
	}
}

func (self sliceSorter) find(key []val.Value) (node.Node, int) {
	found := sort.Search(len(self), self.findFunc(key))
	if found >= 0 && found < len(self) {
		// Search find first match equal or greater, so we need to check for equality
		if val.EqualVals(self[found].key, key) {
			return self[found].n, self[found].pos
		}
	}
	return nil, -1
}

func (self reflect2) buildKeys(s node.Selection, keyMeta []meta.HasDataType, slce reflect.Value) (sliceSorter, error) {
	var err error
	entries := make(sliceSorter, slce.Len())
	for i := range entries {
		entries[i].pos = i
		entries[i].n = self.child(slce.Index(i))
		entries[i].key = make([]val.Value, len(keyMeta))
		for j, k := range keyMeta {
			r := node.FieldRequest{
				Meta: k.(meta.HasDataType),
			}
			var hnd node.ValueHandle
			if err = entries[i].n.Field(r, &hnd); err != nil {
				return nil, err
			}
			entries[i].key[j] = hnd.Val
		}
	}
	sort.Sort(entries)
	return entries, nil
}

func (self reflect2) listSlice(v reflect.Value, onChange onListValueChange) node.Node {
	var entries sliceSorter
	e := v.Type().Elem()
	return &Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			if entries == nil {
				var err error
				entries, err = self.buildKeys(r.Selection, r.Meta.KeyMeta(), v)
				if err != nil {
					return nil, nil, err
				}
			}
			key := r.Key
			if r.New {
				item := self.create(e)
				v = reflect.Append(v, item)
				onChange(v)
				return self.child(item), key, nil
			} else if key != nil {
				if found, i := entries.find(key); found != nil {
					if r.Delete {
						part1 := v.Slice(0, i)
						part2 := v.Slice(i+1, v.Len())
						v = reflect.AppendSlice(part1, part2)
						onChange(v)
						return nil, nil, nil
					}
					return found, key, nil
				}
			} else {
				if r.Row < v.Len() {
					e := entries[r.Row]
					return e.n, e.key, nil
				}
			}
			return nil, nil, nil
		},
	}
}

type valSorter []reflect.Value

func (self valSorter) Len() int {
	return len(self)
}

func (self valSorter) Less(i, j int) bool {
	switch self[i].Type().Kind() {
	case reflect.String:
		return strings.Compare(self[i].String(), self[j].String()) < 0
	case reflect.Int:
		return self[i].Int() < self[j].Int()
	}
	panic("not supported")
}

func (self valSorter) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self reflect2) listMap(v reflect.Value) node.Node {
	var keys []reflect.Value
	e := v.Type().Elem()
	return &Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var item reflect.Value
			key := r.Key
			if r.New {
				item = self.create(e)
				keyVal := reflect.ValueOf(key[0].Value())
				v.SetMapIndex(keyVal, item)
			} else if key != nil {
				keyVal := reflect.ValueOf(key[0].Value())
				if r.Delete {
					v.SetMapIndex(keyVal, reflect.ValueOf(nil))
					return nil, nil, nil
				}
				item = v.MapIndex(keyVal)
			} else {
				if keys == nil {
					keys = v.MapKeys()
					sort.Sort(valSorter(keys))
				}
				if r.Row < len(keys) {
					keyVal := keys[r.Row]
					item = v.MapIndex(keyVal)
				}
			}
			if item.IsValid() {
				return self.child(item), key, nil
			}
			return nil, nil, nil
		},
	}
}

type onListValueChange func(update reflect.Value)

func (self reflect2) childMap(v reflect.Value) node.Node {
	k := v.Type().Key()
	e := v.Type().Elem()
	return &Basic{
		Peekable: v.Interface(),
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch k.Kind() {
			case reflect.String:
				mapKey := reflect.ValueOf(r.Meta.GetIdent())
				var childInstance reflect.Value
				if r.New {
					childInstance = self.create(e)
					v.SetMapIndex(mapKey, childInstance)
				} else if r.Delete {
					v.SetMapIndex(mapKey, reflect.Zero(e))
					return nil, nil
				} else {
					childInstance = v.MapIndex(mapKey)
				}
				if meta.IsList(r.Meta) {
					onUpdate := func(update reflect.Value) {
						v.SetMapIndex(mapKey, update)
					}
					return self.list(childInstance, onUpdate), nil
				} else {
					return self.child(childInstance), nil
				}
			default:
				return nil, c2.NewErr("key type not supported " + k.String())
			}
			return nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch k.Kind() {
			case reflect.String:
				mapKey := reflect.ValueOf(r.Meta.GetIdent())
				if r.Write {
					v.SetMapIndex(mapKey, reflect.ValueOf(hnd.Val.Value()))
				} else {
					fval := v.MapIndex(mapKey)
					if fval.IsValid() {
						var err error
						hnd.Val, err = node.NewValue(r.Meta.GetDataType(), fval.Interface())
						if err != nil {
							return err
						}
					}
				}
			default:
				return c2.NewErr("key type not supported " + k.String())
			}
			return nil
		},
	}
}

func (self reflect2) create(t reflect.Type) reflect.Value {
	switch t.Kind() {
	case reflect.Ptr:
		return reflect.New(t.Elem())
	case reflect.Interface:
		return reflect.ValueOf(make(map[string]interface{}))
	case reflect.Map:
		return reflect.MakeMap(t)
	case reflect.Slice:
		return reflect.MakeSlice(t, 0, 0)
	}
	panic(fmt.Sprintf("creating type not supported %v", t))
}

func (self reflect2) strukt(ptrVal reflect.Value) node.Node {
	elemVal := ptrVal.Elem()
	return &Basic{
		Peekable: ptrVal.Interface(),
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			fieldName := node.MetaNameToFieldName(r.Meta.GetIdent())
			childVal := elemVal.FieldByName(fieldName)
			if r.New {
				childInstance := self.create(childVal.Type())
				childVal.Set(childInstance)
			}
			if meta.IsList(r.Meta) {
				onUpdate := func(update reflect.Value) {
					childVal.Set(update)
				}
				return self.list(childVal, onUpdate), nil
			}
			return self.child(childVal), nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
			if r.Write {
				err = WriteField(r.Meta, ptrVal, hnd.Val)
			} else {
				hnd.Val, err = ReadField(r.Meta, ptrVal)
			}
			return
		},
	}
}

/////////////////
func WriteField(m meta.HasDataType, ptrVal reflect.Value, v val.Value) error {
	return WriteFieldWithFieldName(node.MetaNameToFieldName(m.GetIdent()), m, ptrVal, v)
}

// Look for public fields that match fieldName.  Some attempt will be made to convert value to proper
// type.
//
// TODO: We only look for fields, but it would be useful to look for methods as well with pattern
// Set___(x) or the like
func WriteFieldWithFieldName(fieldName string, m meta.HasDataType, ptrVal reflect.Value, v val.Value) error {
	elemVal := ptrVal.Elem()
	if !elemVal.IsValid() {
		panic(fmt.Sprintf("Cannot find property \"%s\" on invalid or nil %s", fieldName, ptrVal))
	}

	fieldVal := elemVal.FieldByName(fieldName)
	if !fieldVal.IsValid() {
		panic(fmt.Sprintf("Invalid property \"%s\" on %s", fieldName, elemVal.Type()))
	}
	if v == nil {
		panic(fmt.Sprintf("No value given to set %s", m.GetIdent()))
	}
	switch v.Format() {
	case val.FmtEnum:
		e := v.(val.Enum)
		switch fieldVal.Kind() {
		case reflect.Int:
			fieldVal.SetInt(int64(e.Id))
		case reflect.String:
			fieldVal.SetString(e.Label)
		}
	case val.FmtEnumList:
		el := v.(val.EnumList)
		switch fieldVal.Elem().Kind() {
		case reflect.Int:
			fieldVal.Set(reflect.ValueOf(el.Ids()))
		case reflect.String:
			fieldVal.Set(reflect.ValueOf(el.Labels()))
		}
	default:
		fieldVal.Set(reflect.ValueOf(v.Value()))
	}
	return nil
}

func ReadField(m meta.HasDataType, ptrVal reflect.Value) (val.Value, error) {
	return ReadFieldWithFieldName(node.MetaNameToFieldName(m.GetIdent()), m, ptrVal)
}

func ReadFieldWithFieldName(fieldName string, m meta.HasDataType, ptrVal reflect.Value) (v val.Value, err error) {
	elemVal := ptrVal.Elem()
	if elemVal.Kind() == reflect.Ptr {
		panic(fmt.Sprintf("Pointer to a pointer not legal %s on %v ", m.GetIdent(), ptrVal))
	}
	fieldVal := elemVal.FieldByName(fieldName)

	if !fieldVal.IsValid() {
		panic(fmt.Sprintf("Field not found: %s on %v", m.GetIdent(), ptrVal))
	}

	// convert arrays to slices so casts work. this should not make a copy
	// of the array and therefore be efficient operation
	i, err := m.GetDataType().Info()
	if err != nil {
		return nil, err
	}

	// Turn arrays into slices to leverage more of val.Conv's ability to convert data
	if i.Format.IsList() && fieldVal.Kind() == reflect.Array {
		fieldVal = fieldVal.Slice(0, fieldVal.Len())
	}

	switch i.Format {
	case val.FmtString:
		s := fieldVal.String()
		if len(s) == 0 {
			return nil, nil
		}
		return val.String(s), nil
	case val.FmtAny:
		if fieldVal.IsNil() {
			return nil, nil
		}
	}
	return node.NewValue(m.GetDataType(), fieldVal.Interface())
}
