package nodeutil

import (
	"fmt"
	"reflect"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"
)

type sliceAsList struct {
	ref    *Node
	src    reflect.Value
	update NodeListUpdate
}

func newSliceAsList(ref *Node, src reflect.Value, u NodeListUpdate) *sliceAsList {
	return &sliceAsList{
		ref:    ref,
		src:    src,
		update: u,
	}
}

func (def *sliceAsList) getByKey(r node.ListRequest) (reflect.Value, error) {
	row, v, err := def.findByKey(r.Meta, r.Key, r.Meta.KeyMeta())
	if row < 0 || err != nil {
		return v, err
	}
	return v, nil
}

func (def *sliceAsList) getByRow(r node.ListRequest) (reflect.Value, []reflect.Value, error) {
	var empty reflect.Value
	if r.Row >= def.src.Len() {
		return empty, nil, nil
	}
	v := def.src.Index(r.Row)
	if !v.IsValid() || v.IsZero() {
		return empty, nil, nil
	}
	key, _, err := def.getKey(v, r.Meta, r.Meta.KeyMeta())
	return v, key, err
}

func (def *sliceAsList) getKey(item reflect.Value, m meta.Meta, keyMeta []meta.Leafable) ([]reflect.Value, []val.Value, error) {
	if len(keyMeta) == 0 || reflectIsEmpty(item) {
		return nil, nil, nil
	}

	// construct mock field requests to get key so we ensure we consult the correct customizations
	ref2, err := def.ref.New(m, item.Interface())
	if err != nil {
		return nil, nil, fmt.Errorf("%w attempting to get key", err)
	}
	rvKey := make([]reflect.Value, len(keyMeta))
	nvKey := make([]val.Value, len(keyMeta))
	for i, kmeta := range keyMeta {
		r := node.FieldRequest{Meta: kmeta}
		var hnd node.ValueHandle
		err := ref2.Field(r, &hnd)
		if err != nil {
			return nil, nil, fmt.Errorf("%w when get key", err)
		}
		nvKey[i] = hnd.Val
		switch x := ref2.(type) {
		case *Node:
			// use opts to help coerse value to right
			rvKey[i] = reflect.ValueOf(x.getValue(hnd.Val))
		default:
			// not a *Node, so just use value directly and trust it's right type
			rvKey[i] = reflect.ValueOf(hnd.Val.Value())
		}
	}
	return rvKey, nvKey, nil
}

func (def *sliceAsList) findByKey(m meta.Meta, target []val.Value, keyMeta []meta.Leafable) (int, reflect.Value, error) {
	notfound := -1
	var empty reflect.Value
	// full table scan of items in list to find first item that matches key.  consider replacing
	// this with way for implementation to provide an index on a given list and allow this to
	// be the default implementation
	for row := 0; row < def.src.Len(); row++ {
		candidate := def.src.Index(row)
		if !candidate.IsValid() {
			return notfound, empty, fmt.Errorf("row %d of %T is invalid", row, def.src.Type())
		}
		_, candidateKey, err := def.getKey(candidate, m, keyMeta)
		if err != nil {
			return notfound, empty, err
		}
		for i, v := range candidateKey {
			if v.Value() != target[i].Value() {
				break
			}
			isLastKey := i == len(keyMeta)-1
			if isLastKey {
				return row, candidate, nil
			}
		}
	}
	return notfound, empty, nil
}

func (def *sliceAsList) deleteByKey(r node.ListRequest) error {
	row, _, err := def.findByKey(r.Meta, r.Key, r.Meta.KeyMeta())
	if row < 0 || err != nil {
		return err
	}
	part1 := def.src.Slice(0, row)
	part2 := def.src.Slice(row+1, def.src.Len())
	def.src = reflect.AppendSlice(part1, part2)
	if def.update != nil {
		return def.update(def.src)
	}
	return nil
}

func (def *sliceAsList) newListItem(r node.ListRequest) (reflect.Value, error) {
	var empty reflect.Value
	item, err := def.ref.NewObject(def.src.Type().Elem(), r.Meta, true)
	if err != nil {
		return empty, err
	}
	def.src = reflect.Append(def.src, item)
	if def.update != nil {
		err = def.update(def.src)
	}
	return item, err
}

func (def *sliceAsList) setComparator(c ReflectListComparator) {

}
