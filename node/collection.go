package node

import (
	"meta"
)

// Decode and encode data from Go's map and slices of interface{}.  Will assume collections recursively
// but On{Hook} functions provide option to alter this at any place.

type Collection struct {
	OnKey        MetaIdentFunc
	OnExtendMap  ExtendMapFunc
	OnExtendList ExtendListFunc
}

func MapNode(container map[string]interface{}) Node {
	return (&Collection{}).Node(container)
}

func ListNode(entry MappedListHandler) Node {
	return (&Collection{}).List(entry)
}

type MetaIdentFunc func(sel *Selection, goober meta.Meta) (key string)
type ExtendMapFunc func(sel *Selection, goober meta.MetaList, container map[string]interface{}) (Node, error)
type ExtendListFunc func(sel *Selection, goober *meta.List, entry MappedListHandler) (Node, error)

type MapEntry struct {
	Key     string
	Parent  map[string]interface{}
	ListHnd []map[string]interface{}
}

type MappedListHandler interface {
	Append(item map[string]interface{})
	List() []map[string]interface{}
}

func (self *MapEntry) Append(item map[string]interface{}) {
	self.ListHnd = append(self.ListHnd, item)
	self.Parent[self.Key] = self.ListHnd
}

func (self *MapEntry) List() []map[string]interface{} {
	return self.ListHnd
}

func (self *Collection) Node(container map[string]interface{}) (Node) {
	s := &MyNode{}
	s.OnSelect = func(r ContainerRequest) (Node, error) {
		var data interface{}
		keyIdent := self.MetaIdent(r.Selection, r.Meta)
		if r.New {
			if meta.IsList(r.Meta) {
				data = make([]map[string]interface{}, 0, 10)
			} else {
				data = make(map[string]interface{})
			}
			container[keyIdent] = data
		} else {
			data = container[keyIdent]
		}
		if data != nil {
			if meta.IsList(r.Meta) {
				return self.ExtendList(r.Selection, r.Meta.(*meta.List), &MapEntry{
					Key: keyIdent,
					Parent: container,
					ListHnd: data.([]map[string]interface{}),
				})
			}
			return self.ExtendContainer(r.Selection, r.Meta, data)
		}
		return nil, nil
	}
	s.OnWrite = func(r FieldRequest, val *Value) error {
		return self.UpdateLeaf(r.Selection, container, r.Meta, val)
	}
	s.OnRead = func(r FieldRequest) (*Value, error) {
		return self.ReadLeaf(r.Selection, container, r.Meta)
	}
	return s
}

func (self *Collection) ExtendContainer(sel *Selection, goober meta.MetaList, data interface{}) (Node, error) {
	// TODO: Silently ignoring unexpected format. We should be *less*
	// tolerant and fail here otherwise we silently ignore bad data.
	c, found := data.(map[string]interface{})
	if !found {
		return nil, nil
	}
	if self.OnExtendMap != nil {
		if n, err := self.OnExtendMap(sel, goober, data.(map[string]interface{})); n != nil || err != nil {
			return n, err
		}
	}
	return self.Node(c), nil
}

func (self *Collection) ExtendList(sel *Selection, goober *meta.List, entry *MapEntry) (Node, error) {
	if self.OnExtendList != nil {
		if n, err := self.OnExtendList(sel, goober, entry); n != nil || err != nil {
			return n, err
		}
	}
	return self.List(entry), nil
}

func (self *Collection) MetaIdent(sel *Selection, goober meta.Meta) string {
	if self.OnKey != nil {
		return self.OnKey(sel, goober)
	}

	return goober.GetIdent()
}

func (self *Collection) ReadKey(sel *Selection, container map[string]interface{}, goober *meta.List) (key []*Value, err error) {
	keyMeta := goober.KeyMeta()
	key = make([]*Value, len(keyMeta))
	for i, m := range keyMeta {
		if key[i], err = self.ReadLeaf(sel, container, m); err != nil {
			return nil, err
		}
	}
	return
}

func (self *Collection) List(entry MappedListHandler) (Node) {
	s := &MyNode{}
	s.OnNext = func(r ListRequest) (Node, []*Value, error) {
		var selected map[string]interface{}
		if r.New {
			selection := make(map[string]interface{})
			entry.Append(selection)
			n, err := self.ExtendContainer(r.Selection, r.Meta, selection)
			return n, r.Key, err
		} else if len(entry.List()) > 0 {
			if len(r.Key) > 0 {
				if !r.First {
					return nil, nil, nil
				}
				// looping not very efficient, but we do not have an index
				for _, candidate := range entry.List() {
					// TODO: Support compound keys
					candidateKey := SetValues(r.Meta.KeyMeta(), candidate[self.MetaIdent(r.Selection, r.Meta.KeyMeta()[0])])
					if  r.Key[0].Equal(candidateKey[0]) {
						selected = candidate
						break
					}
				}
			} else {
				if int(r.Row) < len(entry.List()) {
					selected = entry.List()[r.Row]
				}
				var err error
				if r.Key, err = self.ReadKey(r.Selection, selected, r.Meta); err != nil {
					return nil, nil, err
				}
			}
		}
		if selected != nil {
			n, err := self.ExtendContainer(r.Selection, r.Meta, selected)
			return n, r.Key, err
		}
		return nil, nil, nil
	}
	return s
}

func (self *Collection) ReadLeaf(sel *Selection, container map[string]interface{}, m meta.HasDataType) (*Value, error) {
	return SetValue(m.GetDataType(), container[self.MetaIdent(sel, m)])
}

func (self *Collection) UpdateLeaf(sel *Selection, container map[string]interface{}, m meta.HasDataType, v *Value) error {
	container[self.MetaIdent(sel, m)] = v.Value()
	return nil
}
