package node

import (
	"fmt"
	"github.com/c2stack/c2g/meta"
	"strings"
)

type StoreData struct {
	Meta  meta.MetaList
	Store Store
}

func StoreNode(store Store) Node {
	return (&StoreData{Store:store}).Node()
}

func NewStoreData(meta meta.MetaList, store Store) *StoreData {
	return &StoreData{
		Meta: meta,
		Store:  store,
	}
}

func (kv *StoreData) Browser() *Browser {
	return NewBrowser(kv.Meta, func() Node { return kv.Node() })
}

func (kv *StoreData) Node() (Node) {
	var err error
	if err = kv.Store.Load(); err != nil {
		return ErrorNode{Err:err}
	}
	return kv.Container("")
}

func (kv *StoreData) OnEvent(sel Selection, e Event) error {
	switch e.Type {
	case END_TREE_EDIT:
		return kv.Store.Save()
	}
	return nil
}

func (kv *StoreData) List(parentPath string) Node {
	s := &MyNode{Label:"StoreData List"}
	var keyList []string
	s.OnNext = func(r ListRequest) (Node, []*Value, error) {
		key := r.Key
		if r.New {
			var childPath string
			if len(key) > 0 {
				childPath = kv.listPath(parentPath, key)
			} else {
				childPath = parentPath + "=unknown"
			}
			return kv.Container(childPath), key, nil
		}
		if len(key) > 0 {
			if r.First {
				path := kv.listPath(parentPath, key)
				if hasMore := kv.Store.HasValues(path); hasMore {
					return kv.Container(path), key, nil
				}
			} else {
				return nil, nil, nil
			}
		} else {
			var err error
			if r.First {
				if keyList, err = kv.Store.KeyList(parentPath, r.Meta); err != nil {
					return nil, nil, err
				}
			}
			if hasMore := r.Row < len(keyList); hasMore {
				if key, err = CoerseKeys(r.Meta, []string{keyList[r.Row]}); err != nil {
					return nil, nil, err
				}
				path := kv.listPath(parentPath, key)
				return kv.Container(path), key, nil
			}
		}
		return nil, nil, nil
	}
	s.OnEvent = func(sel Selection, e Event) error {
		switch e.Type {
		case DELETE:
			return kv.Store.RemoveAll(parentPath)
		}
		return kv.OnEvent(sel, e)
	}
	s.OnAction = func(r ActionRequest) (output Node, err error) {
		path := kv.listPath(parentPath, r.Selection.Path.key)
		var action ActionFunc
		if action, err = kv.Store.Action(path); err != nil {
			return
		}
		return action(r)
	}
	return s
}

func (kv *StoreData) containerPath(parentPath string, m meta.Meta) string {
	if len(parentPath) == 0 {
		return m.GetIdent()
	}
	return fmt.Sprint(parentPath, "/", m.GetIdent())
}

func (kv *StoreData) listPath(parentPath string, key []*Value) string {
	// TODO: support compound keys
	return fmt.Sprint(parentPath, "=", key[0].String())
}

func (kv *StoreData) listPathWithNewKey(parentPath string, key []*Value) string {
	eq := strings.LastIndex(parentPath, "=")
	return kv.listPath(parentPath[:eq], key)
}

func (kv *StoreData) Container(copy string) Node {
	s := &MyNode{Label:"StoreData Container"}
	//path := storePath{parent:parentPath}
	s.OnChoose = func(sel Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
		// go thru each case and if there are any properties in the data that are not
		// part of the meta, that disqualifies that case and we move onto next case
		// until one case aligns with data.  If no cases align then input in inconclusive
		// i.e. non-discriminating and we should error out.
		cases := meta.NewMetaListIterator(choice, false)
		for cases.HasNextMeta() {
			kase := cases.NextMeta().(*meta.ChoiceCase)
			props := meta.NewMetaListIterator(kase, true)
			for props.HasNextMeta() {
				prop := props.NextMeta()
				candidatePath := kv.containerPath(copy, prop)
				if found := kv.Store.HasValues(candidatePath); found {
					return kase, nil
				}
			}
		}
		return nil, nil
	}
	s.OnSelect = func(r ContainerRequest) (child Node, err error) {
		if r.New {
			if meta.IsList(r.Meta) {
				childPath := kv.containerPath(copy, r.Meta)
				return kv.List(childPath), nil
			} else {
				childPath := kv.containerPath(copy, r.Meta)
				return kv.Container(childPath), nil
			}
		}
		childPath := kv.containerPath(copy, r.Meta)
		if kv.Store.HasValues(childPath) {
		if meta.IsList(r.Meta) {
				return kv.List(childPath), nil
			} else {
				return kv.Container(childPath), nil
			}
		}
		return
	}
	s.OnField = func(r FieldRequest, hnd *ValueHandle) (err error) {
		if r.Write {
			propPath := kv.containerPath(copy, r.Meta)
			if err = kv.Store.SetValue(propPath, hnd.Val); err != nil {
				return err
			}
			if meta.IsKeyLeaf(r.Selection.Path.meta.(meta.MetaList), r.Meta) {
				oldPath := copy
				// TODO: Support compound keys
				newKey := []*Value{hnd.Val}
				newPath := kv.listPathWithNewKey(copy, newKey)
				kv.Store.RenameKey(oldPath, newPath)
			}
		} else {
			hnd.Val = kv.Store.Value(kv.containerPath(copy, r.Meta), r.Meta.GetDataType())
		}
		return
	}
	s.OnEvent = func(sel Selection, e Event) error {
		switch e.Type {
		case DELETE:
			return kv.Store.RemoveAll(copy)
		}
		return kv.OnEvent(sel, e)
	}
	s.OnAction = func(r ActionRequest) (output Node, err error) {
		path := kv.containerPath(copy, r.Meta)
		var action ActionFunc
		if action, err = kv.Store.Action(path); err != nil {
			return
		}
		return action(r)
	}
	return s
}
