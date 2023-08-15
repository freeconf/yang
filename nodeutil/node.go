package nodeutil

import (
	"context"
	"fmt"
	"reflect"
	"sort"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"
)

type NodeOptions struct {
	IdentitiesAsStrings bool
	EnumAsStrings       bool
	EnumAsInt           bool
	IgnoreEmpty         bool
	TryPluralOnLists    bool
	Ident               string
	GetterPrefix        string
	SetterPrefix        string
}

type Node struct {
	Object  any
	Options NodeOptions

	OnOptions func(ref *Node, m meta.Definition, o NodeOptions) NodeOptions

	OnChild       func(ref *Node, r node.ChildRequest) (node.Node, error)
	OnGetChild    func(ref *Node, r node.ChildRequest) (node.Node, error)
	OnNewChild    func(ref *Node, r node.ChildRequest) (node.Node, error)
	OnDeleteChild func(ref *Node, r node.ChildRequest) error

	OnField      func(ref *Node, r node.FieldRequest, hnd *node.ValueHandle) error
	OnGetField   func(ref *Node, r node.FieldRequest) (val.Value, error)
	OnSetField   func(ref *Node, r node.FieldRequest, v val.Value) error
	OnClearField func(ref *Node, r node.FieldRequest) error

	OnBeginEdit func(ref *Node, r node.NodeRequest) error
	OnEndEdit   func(ref *Node, r node.NodeRequest) error
	OnChoose    func(ref *Node, sel *node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error)

	OnGetListDef  func(ref *Node, r node.ListRequest) (reflectList, error)
	OnNewListItem func(ref *Node, r node.ListRequest) (node.Node, []val.Value, error)
	OnGetByKey    func(ref *Node, r node.ListRequest) (any, error)
	OnGetByRow    func(ref *Node, r node.ListRequest) (any, []val.Value, error)
	OnDeleteByKey func(ref *Node, r node.ListRequest) error

	c reflectContainer
	l reflectList
}

type NodeListUpdate func(update reflect.Value) error

func (ref *Node) Child(r node.ChildRequest) (node.Node, error) {
	if ref.OnChild != nil {
		return ref.OnChild(ref, r)
	}
	return ref.DoChild(r)
}

func (ref *Node) DoChild(r node.ChildRequest) (node.Node, error) {
	if r.Delete {
		if ref.OnDeleteChild != nil {
			return nil, ref.OnDeleteChild(ref, r)
		} else {
			return nil, ref.DoDeleteChild(r)
		}
	}
	if r.New {
		if ref.OnNewChild != nil {
			return ref.OnNewChild(ref, r)
		} else {
			return ref.DoNewChild(r)
		}
	}
	if ref.OnGetChild != nil {
		return ref.OnGetChild(ref, r)
	}
	return ref.DoGetChild(r)
}

func (ref *Node) Next(r node.ListRequest) (node.Node, []val.Value, error) {
	var found any
	var err error
	key := r.Key
	if r.New {
		if ref.OnNewListItem != nil {
			return ref.OnNewListItem(ref, r)
		}
		item, err := ref.DoNewListItem(r)
		if err != nil {
			return nil, nil, err
		}
		found = item
	} else if key != nil {
		if r.Delete {
			if ref.OnGetByKey != nil {
				err = ref.OnDeleteByKey(ref, r)
			} else {
				err = ref.DoDeleteByKey(r)
			}
		} else {
			if ref.OnGetByKey != nil {
				found, err = ref.OnGetByKey(ref, r)
			} else {
				found, err = ref.DoGetByKey(r)
			}
		}
		if err != nil {
			return nil, nil, err
		}
	} else {
		if ref.OnGetByRow != nil {
			found, key, err = ref.OnGetByRow(ref, r)
		} else {
			found, key, err = ref.DoGetByRow(r)
		}
		if err != nil {
			return nil, nil, err
		}
	}
	if found == nil {
		return nil, nil, nil
	}
	return ref.New(found), key, err
}

func (ref *Node) Field(r node.FieldRequest, hnd *node.ValueHandle) error {
	if ref.OnField != nil {
		return ref.OnField(ref, r, hnd)
	}
	return ref.DoField(r, hnd)
}

func (ref *Node) DoField(r node.FieldRequest, hnd *node.ValueHandle) error {
	if r.Clear {
		if ref.OnClearField != nil {
			return ref.OnClearField(ref, r)
		} else {
			return ref.DoClearField(r)
		}
	}
	if r.Write {
		if ref.OnSetField != nil {
			return ref.OnSetField(ref, r, hnd.Val)
		} else {
			return ref.DoSetField(r, hnd.Val)
		}
	}
	var err error
	if ref.OnGetField != nil {
		hnd.Val, err = ref.OnGetField(ref, r)
	} else {
		hnd.Val, err = ref.DoGetField(r)
	}
	return err
}

func (ref *Node) BeginEdit(r node.NodeRequest) error {
	if ref.OnBeginEdit != nil {
		return ref.OnBeginEdit(ref, r)
	}
	return nil
}

func (ref *Node) EndEdit(r node.NodeRequest) error {
	if ref.OnEndEdit != nil {
		return ref.OnEndEdit(ref, r)
	}
	return nil
}

func (ref *Node) Action(r node.ActionRequest) (node.Node, error) {
	// possible but not now
	return nil, fc.NotImplementedError
}

func (ref *Node) Notify(r node.NotifyRequest) (node.NotifyCloser, error) {
	return nil, fc.NotImplementedError
}

func (ref *Node) Peek(sel *node.Selection, consumer interface{}) interface{} {
	return ref.Object
}

func (ref *Node) Choose(sel *node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
	if ref.OnChoose != nil {
		return ref.OnChoose(ref, sel, choice)
	}
	return ref.DoChoose(sel, choice)
}

func (ref *Node) Context(sel *node.Selection) context.Context {
	return sel.Context
}

// Release If you need to implement this, use Extend
func (ref *Node) Release(sel *node.Selection) {}

func (ref *Node) DoChoose(sel *node.Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
	c, err := ref.container()
	if err != nil {
		return nil, err
	}
	for _, caseId := range choice.CaseIdents() {
		cs := choice.Cases()[caseId] // by iterating thru case ids and not cases we get a predictable order
		for _, ddef := range cs.DataDefinitions() {
			if c.exists(ddef) {
				return cs, nil
			}
		}
	}
	return nil, nil
}

func (ref *Node) DoGetField(r node.FieldRequest) (val.Value, error) {
	c, err := ref.container()
	if err != nil {
		return nil, err
	}
	v, err := c.get(r.Meta)
	if err != nil || !v.IsValid() {
		return nil, err
	}
	return node.NewValue(r.Meta.Type(), v.Interface())
}

func reflectIsEmpty(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}
	if v.IsZero() {
		return true
	}
	if canNil(v.Kind()) && v.IsNil() {
		return true
	}
	src := v
	if src.Kind() == reflect.Pointer {
		src = src.Elem()
	}
	switch src.Kind() {
	case reflect.Slice, reflect.Map:
		if src.Len() == 0 {
			return true
		}
	}
	return false
}

func (ref *Node) DoClearField(r node.FieldRequest) error {
	c, err := ref.container()
	if err != nil {
		return err
	}
	return c.clear(r.Meta)
}

func (ref *Node) DoSetField(r node.FieldRequest, v val.Value) error {
	c, err := ref.container()
	if err != nil {
		return err
	}
	return c.set(r.Meta, reflect.ValueOf(ref.getValue(v)))
}

func (ref *Node) DoDeleteChild(r node.ChildRequest) error {
	c, err := ref.container()
	if err != nil {
		return err
	}
	return c.clear(r.Meta)
}

func (ref *Node) DoNewChild(r node.ChildRequest) (node.Node, error) {
	c, err := ref.container()
	if err != nil {
		return nil, err
	}
	obj, err := c.newChild(r.Meta)
	if err != nil {
		return nil, err
	}
	if obj.IsNil() {
		return nil, fmt.Errorf("could not create new item at %s", r.Path)
	}
	if err = c.set(r.Meta, obj); err != nil {
		return nil, err
	}
	if meta.IsList(r.Meta) && r.Selection.Path.Meta != r.Meta {
		return ref.NewList(obj.Interface(), ref.onListUpdate(r.Meta.(*meta.List)))
	}

	return ref.New(obj.Interface()), nil
}

func (ref *Node) onListUpdate(m *meta.List) NodeListUpdate {
	return func(update reflect.Value) error {
		c, err := ref.container()
		if err != nil {
			return err
		}
		return c.set(m, update)
	}
}

func (ref *Node) container() (reflectContainer, error) {
	if ref.c == nil {
		var err error
		if ref.c, err = ref.newContainerHandler(); err != nil {
			return nil, err
		}
	}
	return ref.c, nil
}

func (ref *Node) options(m meta.Definition) NodeOptions {
	if ref.OnOptions != nil {
		return ref.OnOptions(ref, m, ref.Options)
	}
	return ref.Options
}

func (ref *Node) DoGetChild(r node.ChildRequest) (node.Node, error) {
	c, err := ref.container()
	if err != nil {
		return nil, err
	}
	obj, err := c.get(r.Meta)
	if err != nil || !obj.IsValid() || obj.IsNil() {
		return nil, err
	}
	if ref.Options.IgnoreEmpty && !r.New && reflectIsEmpty(obj) {
		return nil, nil
	}
	if meta.IsList(r.Meta) && r.Selection.Path.Meta != r.Meta {
		return ref.NewList(obj.Interface(), ref.onListUpdate(r.Meta.(*meta.List)))
	}
	return ref.New(obj.Interface()), nil
}

func (ref *Node) NewList(obj any, u NodeListUpdate) (*Node, error) {
	copy := ref.New(obj)
	var err error
	copy.l, err = ref.newListHandler(reflect.ValueOf(obj), u)
	if err != nil {
		return nil, err
	}
	return copy, nil
}

func (ref *Node) New(obj any) *Node {
	if _, isVal := obj.(reflect.Value); isVal {
		panic("passing in reflect.Value and not true obj")
	}
	copy := *ref
	copy.Object = obj
	copy.l = nil
	copy.c = nil
	return &copy
}

type reflectContainer interface {
	get(meta.Definition) (reflect.Value, error)
	set(meta.Definition, reflect.Value) error
	exists(meta.Definition) bool
	clear(meta.Definition) error
	newChild(meta.HasDataDefinitions) (reflect.Value, error)
}

type reflectList interface {
	getByKey(r node.ListRequest) (reflect.Value, error)
	getByRow(r node.ListRequest) (reflect.Value, []reflect.Value, error)
	deleteByKey(r node.ListRequest) error
	newListItem(r node.ListRequest) (reflect.Value, error)
	setComparator(c ReflectListComparator)
}

func (ref *Node) newListHandler(src reflect.Value, u NodeListUpdate) (reflectList, error) {
	if src.Kind() == reflect.Interface {
		src = src.Elem()
	}
	if src.Kind() == reflect.Map {
		return newMapAsList(src), nil
	}
	if src.Kind() == reflect.Slice {
		return newSliceAsList(ref, src, u), nil
	}
	return nil, fmt.Errorf("could not find list handler for '%s'", src.Type())
}

func (ref *Node) newContainerHandler() (reflectContainer, error) {
	src := reflect.ValueOf(ref.Object)
	k := src.Kind()
	if k == reflect.Map {
		return &mapAsContainer{
			src: src,
		}, nil
	}
	if k == reflect.Struct || (k == reflect.Pointer && src.Elem().Kind() == reflect.Struct) {
		return newStructAsContainer(ref, src), nil
	}

	// if you get here, default handlers do not handle your case and you need to either override
	// or submit merge request to FreeCONF project

	return nil, fmt.Errorf("could not use type '%s' for a container definition", src.Type())
}

func (ref *Node) DoGetByKey(r node.ListRequest) (any, error) {
	r.Selection.Find(r.Meta.Ident())
	item, err := ref.l.getByKey(r)
	if err != nil || !item.IsValid() || item.IsNil() {
		return nil, err
	}
	return item.Interface(), nil
}

func (ref *Node) DoGetByRow(r node.ListRequest) (any, []val.Value, error) {
	item, keyVals, err := ref.l.getByRow(r)
	if err != nil || reflectIsEmpty(item) {
		return nil, nil, err
	}
	var key []val.Value
	if len(keyVals) > 0 {
		key = make([]val.Value, len(keyVals))
		var err error
		for i := 0; i < len(keyVals); i++ {
			if key[i], err = node.NewValue(r.Meta.KeyMeta()[0].Type(), keyVals[i].Interface()); err != nil {
				return nil, nil, err
			}
		}
	}
	return item.Interface(), key, nil

}

func (ref *Node) DoDeleteByKey(r node.ListRequest) error {
	return ref.l.deleteByKey(r)
}

func (ref *Node) DoNewListItem(r node.ListRequest) (any, error) {
	item, err := ref.l.newListItem(r)
	if err != nil || !item.IsValid() || item.IsNil() {
		return nil, err
	}
	return item.Interface(), nil
}

func (ref *Node) getValue(v val.Value) any {
	switch x := v.(type) {
	case val.IdentRef:
		if ref.Options.IdentitiesAsStrings {
			return x.Label
		}
	case val.IdentRefList:
		if ref.Options.IdentitiesAsStrings {
			return x.Labels()
		}
	case val.Enum:
		if ref.Options.EnumAsStrings {
			return x.Label
		}
		if ref.Options.EnumAsInt {
			return int64(x.Id)
		}
	case val.EnumList:
		if ref.Options.EnumAsStrings {
			return x.Labels()
		}
		if ref.Options.EnumAsInt {
			return x.Ids()
		}
	}
	return v.Value()
}

func newObject(t reflect.Type, m meta.Definition) (reflect.Value, error) {
	switch t.Kind() {
	case reflect.Ptr:
		return reflect.New(t.Elem()), nil
	case reflect.Interface:
		switch x := m.(type) {
		case *meta.List:
			keyMeta := x.KeyMeta()
			if len(keyMeta) == 1 {
				// support some common key types, but anything too unusual should have
				// custom implementation and would default to map[interface{}]interface{}
				// which is likely fine
				switch keyMeta[0].Type().Format() {
				case val.FmtString:
					return reflect.ValueOf(make(map[string]interface{})), nil
				case val.FmtInt32:
					return reflect.ValueOf(make(map[int]interface{})), nil
				case val.FmtInt64:
					return reflect.ValueOf(make(map[int64]interface{})), nil
				case val.FmtDecimal64:
					return reflect.ValueOf(make(map[float64]interface{})), nil
				}
			}
		}
		return reflect.ValueOf(make(map[interface{}]interface{})), nil
	case reflect.Map:
		return reflect.MakeMap(t), nil
	case reflect.Slice:
		return reflect.MakeSlice(t, 0, 0), nil
	}
	panic(fmt.Sprintf("creating type not supported %v", t))
}

type ReflectListComparator func(a, b reflect.Value) bool

// Index help you implement Node.OnNext when you have a map to interate over
type index struct {
	vals       []reflect.Value
	comparator ReflectListComparator
}

func newIndex(vals []reflect.Value, c ReflectListComparator) *index {
	index := &index{
		vals:       vals,
		comparator: c,
	}

	if index.comparator == nil {
		index.comparator = reflectCompare
	}
	sort.Sort(index)
	return index
}

func reflectCompare(a, b reflect.Value) bool {
	if a.CanInt() {
		return a.Int() < b.Int()
	}
	if a.CanFloat() {
		return a.Float() < b.Float()
	}
	if a.Kind() == reflect.String {
		return a.String() < b.String()
	}
	panic(fmt.Sprintf("cannot compare %s. you must set comparator or implement your own list handler", a.Type()))
}

func (ndx *index) Len() int {
	return len(ndx.vals)
}

func (ndx *index) Swap(i, j int) {
	ndx.vals[i], ndx.vals[j] = ndx.vals[j], ndx.vals[i]
}

func (ndx *index) Less(i, j int) bool {
	return ndx.comparator(ndx.vals[i], ndx.vals[j])
}
