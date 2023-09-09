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
	IdentitiesAsStrings  bool
	EnumAsStrings        bool
	EnumAsInt            bool
	IgnoreEmpty          bool
	TryPluralOnLists     bool
	Ident                string
	GetterPrefix         string
	SetterPrefix         string
	ActionOutputExploded bool
	ActionInputExploded  bool
}

// Node uses reflection to understand any object, map, slice and map requests to underlying
// go structures. In it's most basic form, just call:
//
//	var n node.Node
//	n = &nodeutil.Reflect{
//	    Object: yourObject,
//	}
//
// but there are a lot of ways to handle customizations all by setting additional fields
// on this struct
type Node struct {

	// Object initially you set this, but with each naviation into containers, lists and list
	// items, this takes on the respective struct, map, slice automaically.
	Object any

	// Options holds basic settings for all cloned *Node objects from this node. See docs
	// on NodeOptions for various options
	Options NodeOptions

	// OnOptions is a way to change options just for a single field
	OnOptions func(n *Node, m meta.Definition, o NodeOptions) NodeOptions

	// OnChild ignores all internal code for reading and writing a child structure
	// but if you call ref.DoChild(r) from inside this call it will get the default
	// behavior
	OnChild func(n *Node, r node.ChildRequest) (node.Node, error)

	// OnGetChild ignores all internal code for reading a child structure
	// but if you call ref.DoGetChild(r) from inside this call it will get the default
	// behavior
	OnGetChild func(n *Node, r node.ChildRequest) (node.Node, error)

	// OnNewChild ignores all internal code for creating a child structure
	// but if you call ref.DoNewChild(r) from inside this call it will get the default
	// behavior
	OnNewChild func(n *Node, r node.ChildRequest) (node.Node, error)

	// OnDeleteChild ignores all internal code for deleting a child structure
	// but if you call ref.DoDeleteChild(r) from inside this call it will get the default
	// behavior
	OnDeleteChild func(n *Node, r node.ChildRequest) error

	// OnField ignores all internal code for reading and writing a field
	// but if you call ref.DoField(r, hnd) from inside this call it will get the default
	// behavior
	OnField func(n *Node, r node.FieldRequest, hnd *node.ValueHandle) error

	// OnGetField ignores all internal code for reading a field
	// but if you call ref.DoGetField(r) from inside this call it will get the default
	// behavior
	OnGetField func(n *Node, r node.FieldRequest) (val.Value, error)

	// OnSetField ignores all internal code for writing a field
	// but if you call ref.DoSetField(r, v) from inside this call it will get the default
	// behavior
	OnSetField func(n *Node, r node.FieldRequest, v val.Value) error

	// OnClearField ignores all internal code for clearing a field
	// but if you call ref.DoClearField(r, v) from inside this call it will get the default
	// behavior
	OnClearField func(n *Node, r node.FieldRequest) error

	// OnRead allows you to transform a value immediately after reading it from your data
	// structure or map and convert it into a different reflect.Value.  If you do not wish
	// to convert a value, you must return the passed in `v` value
	OnRead func(n *Node, m meta.Definition, t reflect.Type, v reflect.Value) (reflect.Value, error)

	// OnWrite allows you to convert a value just before writing it to your data
	// structure or map and convert it into a different reflect.Value.  If you do not wish
	// to convert a value, you must return the passed in `v` value
	OnWrite func(n *Node, m meta.Definition, t reflect.Type, v reflect.Value) (reflect.Value, error)

	// OnBeginEdit is called just before this node is editing or before any of it's children
	// are edited according to docs: https://freeconf.org/docs/reference/node/edit-traversal/
	//
	// No need to call anything else if you do not need this call
	OnBeginEdit func(n *Node, r node.NodeRequest) error

	// OnEndEdit is called just after this node is editing and after any of it's children
	// are edited according to docs: https://freeconf.org/docs/reference/node/edit-traversal/
	//
	// No need to call anything else if you do not need this call
	OnEndEdit func(n *Node, r node.NodeRequest) error

	// OnChoose ignores all internal code for choosing which case of a yang choice case
	// is valid using reflection.  If you call ref.DoChoose(r, sel, choice) from inside
	// this call it will get the default behavior
	OnChoose func(n *Node, sel *node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error)

	// OnNewListItem ignores all internal code for creating AND adding a new items into a list
	// If a new slice instance is created in the process, the parent object reference
	// is updated.
	// If you call ref.DoNewListItem(r) from inside this call it will get the default behavior
	OnNewListItem func(n *Node, r node.ListRequest) (node.Node, error)

	// OnGetByKey ignores all internal code for getting an item out of a list.  For maps
	// the default assumes the map index is the key.  Slices need to iterate thru the list
	// until they find an item that has a key that matches the target key.
	// If you call ref.DoGetByKey(r) from inside this call it will get the default behavior
	OnGetByKey func(n *Node, r node.ListRequest) (node.Node, error)

	// OnGetByRow ignores all internal code for getting an item out of a list by row number.
	// For maps the default behavior uses the map index as the sort key to give items in a
	// predictable order.  Slices need to iterate thru the list in the order they appear in
	// the slice. If you call ref.DoGetByRow(r) from inside this call it will get the default
	// behavior
	OnGetByRow func(n *Node, r node.ListRequest) (node.Node, []val.Value, error)

	// OnDeleteByKey ignores all internal code for deleting an item out of a list by key.
	// For maps the default behavior uses the map index as the key.   Slices need to iterate
	// thru the list until they find an item that has a key that matches the target key to
	// delete. If a new slice instance is created in the process, the parent object reference
	// is updated.
	// If you call ref.OnDeleteByKey(r) from inside this call it will get the default
	// behavior
	OnDeleteByKey func(n *Node, r node.ListRequest) error

	OnAction func(n *Node, r node.ActionRequest) (node.Node, error)

	OnNotify func(n *Node, r node.NotifyRequest) (node.NotifyCloser, error)

	// OnNewObject is called when a new object needs to be created whether it is a child node
	// an rpc input.
	OnNewObject func(t reflect.Type, m meta.Definition, insideList bool) (reflect.Value, error)

	OnContext func(n *Node, s *node.Selection) context.Context

	c reflectContainer // internal handler based on object type to handle containers and leafs
	l reflectList      // internal handler based on object type created to handle lists
}

// If you create a node to handle a list, you might want to receive an update if the list
// address pointer has changed. This is called when a slice is expanded or contracted.  It is
// never called on maps as their address is never changed
type NodeListUpdate func(update reflect.Value) error

// Child node.Node implementation
func (n *Node) Child(r node.ChildRequest) (node.Node, error) {
	if n.OnChild != nil {
		return n.OnChild(n, r)
	}
	return n.DoChild(r)
}

func (n *Node) DoChild(r node.ChildRequest) (node.Node, error) {
	if r.Delete {
		if n.OnDeleteChild != nil {
			return nil, n.OnDeleteChild(n, r)
		} else {
			return nil, n.DoDeleteChild(r)
		}
	}
	if r.New {
		if n.OnNewChild != nil {
			return n.OnNewChild(n, r)
		} else {
			return n.DoNewChild(r)
		}
	}
	if n.OnGetChild != nil {
		return n.OnGetChild(n, r)
	}
	return n.DoGetChild(r)
}

// Child node.Node implementation
func (n *Node) Next(r node.ListRequest) (node.Node, []val.Value, error) {
	var found node.Node
	var err error
	key := r.Key
	if r.New {
		if n.OnNewListItem != nil {
			found, err = n.OnNewListItem(n, r)
		} else {
			found, err = n.DoNewListItem(r)
		}
	} else if key != nil {
		if r.Delete {
			if n.OnGetByKey != nil {
				err = n.OnDeleteByKey(n, r)
			} else {
				err = n.DoDeleteByKey(r)
			}
		} else {
			if n.OnGetByKey != nil {
				found, err = n.OnGetByKey(n, r)
			} else {
				found, err = n.DoGetByKey(r)
			}
		}
	} else {
		if n.OnGetByRow != nil {
			found, key, err = n.OnGetByRow(n, r)
		} else {
			found, key, err = n.DoGetByRow(r)
		}
	}
	if found == nil || err != nil {
		return nil, nil, err
	}
	return found, key, nil
}

// Child node.Node implementation
func (n *Node) Field(r node.FieldRequest, hnd *node.ValueHandle) error {
	if n.OnField != nil {
		return n.OnField(n, r, hnd)
	}
	return n.DoField(r, hnd)
}

func (n *Node) DoField(r node.FieldRequest, hnd *node.ValueHandle) error {
	if r.Clear {
		if n.OnClearField != nil {
			return n.OnClearField(n, r)
		} else {
			return n.DoClearField(r)
		}
	}
	if r.Write {
		if n.OnSetField != nil {
			return n.OnSetField(n, r, hnd.Val)
		} else {
			return n.DoSetField(r, hnd.Val)
		}
	}
	var err error
	if n.OnGetField != nil {
		hnd.Val, err = n.OnGetField(n, r)
	} else {
		hnd.Val, err = n.DoGetField(r)
	}
	return err
}

// Child node.Node implementation
func (n *Node) BeginEdit(r node.NodeRequest) error {
	if n.OnBeginEdit != nil {
		return n.OnBeginEdit(n, r)
	}
	return nil
}

// Child node.Node implementation
func (n *Node) EndEdit(r node.NodeRequest) error {
	if n.OnEndEdit != nil {
		return n.OnEndEdit(n, r)
	}
	return nil
}

type inputParms map[string]interface{}

func (n *Node) Action(r node.ActionRequest) (node.Node, error) {
	if n.OnAction != nil {
		return n.OnAction(n, r)
	}
	return n.DoAction(r)
}

func (n *Node) DoAction(r node.ActionRequest) (node.Node, error) {
	a, err := newActionHandler(reflect.ValueOf(n.Object), r.Meta, n.options(r.Meta))
	if err != nil {
		return nil, err
	}
	return a.do(n, r.Input)
}

func (n *Node) Notify(r node.NotifyRequest) (node.NotifyCloser, error) {
	if n.OnNotify != nil {
		return n.OnNotify(n, r)
	}
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
	if ref.OnContext != nil {
		return ref.OnContext(ref, sel)
	}
	return sel.Context
}

// Release If you need to implement this, use Extend
func (ref *Node) Release(sel *node.Selection) {}

func (ref *Node) DoChoose(sel *node.Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
	for _, caseId := range choice.CaseIdents() {
		cs := choice.Cases()[caseId] // by iterating thru case ids and not cases we get a predictable order
		for _, ddef := range cs.DataDefinitions() {
			if ref.exists(ddef) {
				return cs, nil
			}
		}
	}
	return nil, nil
}

func (ref *Node) exists(m meta.Definition) bool {
	// we make requests and go thru node.Node API so that custom hooks
	// are called
	if meta.IsList(m) || meta.IsContainer(m) {
		r := node.ChildRequest{Meta: m.(meta.HasDataDefinitions)}
		found, cerr := ref.Child(r)
		if found != nil && cerr == nil {
			return true
		}
	} else {
		r := node.FieldRequest{Meta: m.(meta.Leafable)}
		var hnd node.ValueHandle
		ferr := ref.Field(r, &hnd)
		if hnd.Val != nil && ferr == nil {
			return true
		}
	}
	return false
}

func (ref *Node) writeValue(m meta.Definition, v reflect.Value) error {
	c, err := ref.container()
	if err != nil {
		return err
	}

	if ref.OnWrite != nil {
		t, err := c.getType(m)
		if err != nil {
			return err
		}
		v, err = ref.OnWrite(ref, m, t, v)
		if err != nil {
			return err
		}
	}

	err = c.set(m, v)
	if err != nil {
		return err
	}

	return nil
}

func (ref *Node) readValue(m meta.Definition) (reflect.Value, error) {
	var empty reflect.Value
	c, err := ref.container()
	if err != nil {
		return empty, err
	}
	v, err := c.get(m)
	if err != nil {
		return empty, err
	}
	if ref.OnRead != nil {
		return ref.OnRead(ref, m, v.Type(), v)
	}
	return v, nil
}

func (ref *Node) DoGetField(r node.FieldRequest) (val.Value, error) {
	v, err := ref.readValue(r.Meta)
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
	return ref.writeValue(r.Meta, reflect.ValueOf(ref.getValue(v)))
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
	if err = ref.writeValue(r.Meta, obj); err != nil {
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
	obj, err := ref.readValue(r.Meta)
	if err != nil || !obj.IsValid() || (canNil(obj.Kind()) && obj.IsNil()) {
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
	getType(meta.Definition) (reflect.Type, error)
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
		return newMapAsList(ref, src), nil
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

func (ref *Node) DoGetByKey(r node.ListRequest) (node.Node, error) {
	//r.Selection.Find(r.Meta.Ident())
	item, err := ref.l.getByKey(r)
	if err != nil || !item.IsValid() || item.IsNil() {
		return nil, err
	}
	return ref.New(item.Interface()), nil
}

func (ref *Node) DoGetByRow(r node.ListRequest) (node.Node, []val.Value, error) {
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
	return ref.New(item.Interface()), key, nil

}

func (ref *Node) DoDeleteByKey(r node.ListRequest) error {
	return ref.l.deleteByKey(r)
}

func (ref *Node) DoNewListItem(r node.ListRequest) (node.Node, error) {
	item, err := ref.l.newListItem(r)
	if err != nil || !item.IsValid() || item.IsNil() {
		return nil, err
	}
	return ref.New(item.Interface()), nil
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

func (ref *Node) NewObject(t reflect.Type, m meta.Definition, insideList bool) (reflect.Value, error) {
	if ref.OnNewObject != nil {
		return ref.OnNewObject(t, m, insideList)
	}
	return ref.DoNewObject(t, m, insideList)
}

func (ref *Node) DoNewObject(t reflect.Type, m meta.Definition, insideList bool) (reflect.Value, error) {
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
