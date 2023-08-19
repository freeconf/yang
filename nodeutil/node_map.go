package nodeutil

import (
	"fmt"
	"reflect"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
)

type mapAsContainer struct {
	src reflect.Value
}

func newMapAsContainer(src reflect.Value) *mapAsContainer {
	return &mapAsContainer{
		src: src,
	}
}

func (def *mapAsContainer) clear(m meta.Definition) error {
	def.src.SetMapIndex(reflect.ValueOf(def.keyIdent(m)), reflect.Value{})
	return nil
}

func (def *mapAsContainer) keyIdent(m meta.Definition) string {
	return m.Ident()
}

func (def *mapAsContainer) get(m meta.Definition) (reflect.Value, error) {
	return def.src.MapIndex(reflect.ValueOf(def.keyIdent(m))), nil
}

func (def *mapAsContainer) set(m meta.Definition, v reflect.Value) error {
	def.src.SetMapIndex(reflect.ValueOf(def.keyIdent(m)), v)
	return nil
}

func (def *mapAsContainer) getType(m meta.Definition) (reflect.Type, error) {
	return def.src.Type().Elem(), nil
}

func (def *mapAsContainer) newChild(m meta.HasDataDefinitions) (reflect.Value, error) {
	return newObject(def.src.Type().Elem(), m)
}

func (def *mapAsContainer) exists(m meta.Definition) bool {
	fval := def.src.MapIndex(reflect.ValueOf(def.keyIdent(m)))
	return fval.IsValid() && !fval.IsZero()
}

type mapAsList struct {
	index *index
	src   reflect.Value
	c     func(a, b reflect.Value) bool
}

func newMapAsList(src reflect.Value) *mapAsList {
	return &mapAsList{
		src: src,
	}
}

func (def *mapAsList) setComparator(c ReflectListComparator) {
	def.c = c
}

func (def *mapAsList) getByKey(r node.ListRequest) (reflect.Value, error) {
	var empty reflect.Value
	if r.Key == nil || len(r.Key) == 0 {
		return empty, fmt.Errorf("no key specified for %s", r.Path.String())
	}
	keyVal := reflect.ValueOf(r.Key[0].Value())
	found := def.src.MapIndex(keyVal)
	if !found.IsValid() {
		return empty, nil
	}
	return found, nil
}

func (def *mapAsList) deleteByKey(r node.ListRequest) error {
	if r.Key == nil || len(r.Key) == 0 {
		return fmt.Errorf("no key specified for %s", r.Path.String())
	}
	keyVal := reflect.ValueOf(r.Key[0].Value())
	def.src.SetMapIndex(keyVal, reflect.ValueOf(nil))
	return nil
}

func (def *mapAsList) getByRow(r node.ListRequest) (reflect.Value, []reflect.Value, error) {
	var empty reflect.Value
	if def.index == nil {
		def.index = newIndex(def.src.MapKeys(), def.c)
	}
	if r.Row >= len(def.index.vals) {
		return empty, nil, nil
	}
	refKey := def.index.vals[r.Row]
	refVal := def.src.MapIndex(refKey)
	// assumes the map key is the yang key, otherwise the map would be inefficient at best
	return refVal, []reflect.Value{refKey}, nil
}

func (def *mapAsList) newListItem(r node.ListRequest) (reflect.Value, error) {
	t := def.src.Type().Elem()
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	itemVal := reflect.New(t)
	keyVal := reflect.ValueOf(r.Key[0].Value())
	def.src.SetMapIndex(keyVal, itemVal)
	return itemVal, nil
}
