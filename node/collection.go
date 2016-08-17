package node

import (
	"github.com/c2g/meta"
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

type MetaIdentFunc func(sel *Selection, m meta.Meta) (key string)
type ExtendMapFunc func(sel *Selection, m meta.MetaList, container map[string]interface{}) (Node, error)
type ExtendListFunc func(sel *Selection, m *meta.List, entry MappedListHandler) (Node, error)

type MapEntry struct {
	Key         string
	Parent      map[string]interface{}

	// data can come in both these formats, so support either
	listOption1 []interface{}

	listOption2 []map[string]interface{}
}

type MappedListHandler interface {
	Append(item map[string]interface{})
	Len() int
	Item(index int) map[string]interface{}
}

func (self *MapEntry) Append(item map[string]interface{}) {
	if self.listOption1 != nil {
		self.listOption1 = append(self.listOption1, item)
		self.Parent[self.Key] = self.listOption1
	} else {
		self.listOption2 = append(self.listOption2, item)
		self.Parent[self.Key] = self.listOption2
	}
}

func (self *MapEntry) Len() int {
	if self.listOption1 != nil {
		return len(self.listOption1)
	}
	return len(self.listOption2)
}

func (self *MapEntry) Item(i int) map[string]interface{} {
	if self.listOption1 != nil {
		return self.listOption1[i].(map[string]interface{})
	}
	return self.listOption2[i]
}

func (self *Collection) Node(container map[string]interface{}) Node {
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
				me :=  &MapEntry{
					Key:       keyIdent,
					Parent:    container,
				}
				if option1, isOption1 := data.([]interface{}); isOption1 {
					me.listOption1 = option1
				} else {
					me.listOption2 = data.([]map[string]interface{})
				}
				return self.ExtendList(r.Selection, r.Meta.(*meta.List), me)
			}
			return self.ExtendContainer(r.Selection, r.Meta, data)
		}
		return nil, nil
	}
	s.OnField = func(r FieldRequest, hnd *ValueHandle) (err error) {
		if r.Write {
			err = self.UpdateLeaf(r.Selection, container, r.Meta, hnd.Val)
		} else {
			hnd.Val, err = self.ReadLeaf(r.Selection, container, r.Meta)
		}
		return
	}
	return s
}

func (self *Collection) ExtendContainer(sel *Selection, m meta.MetaList, data interface{}) (Node, error) {
	// TODO: Silently ignoring unexpected format. We should be *less*
	// tolerant and fail here otherwise we silently ignore bad data.
	c, found := data.(map[string]interface{})
	if !found {
		return nil, nil
	}
	if self.OnExtendMap != nil {
		if n, err := self.OnExtendMap(sel, m, data.(map[string]interface{})); n != nil || err != nil {
			return n, err
		}
	}
	return self.Node(c), nil
}

func (self *Collection) ExtendList(sel *Selection, m *meta.List, entry *MapEntry) (Node, error) {
	if self.OnExtendList != nil {
		if n, err := self.OnExtendList(sel, m, entry); n != nil || err != nil {
			return n, err
		}
	}
	return self.List(entry), nil
}

func (self *Collection) MetaIdent(sel *Selection, m meta.Meta) string {
	if self.OnKey != nil {
		return self.OnKey(sel, m)
	}

	return m.GetIdent()
}

func (self *Collection) ReadKey(sel *Selection, container map[string]interface{}, m *meta.List) (key []*Value, err error) {
	keyMeta := m.KeyMeta()
	key = make([]*Value, len(keyMeta))
	for i, m := range keyMeta {
		if key[i], err = self.ReadLeaf(sel, container, m); err != nil {
			return nil, err
		}
	}
	return
}

func (self *Collection) List(entry MappedListHandler) Node {
	s := &MyNode{}
	s.OnNext = func(r ListRequest) (Node, []*Value, error) {
		var selected map[string]interface{}
		if r.New {
			selection := make(map[string]interface{})
			entry.Append(selection)
			n, err := self.ExtendContainer(r.Selection, r.Meta, selection)
			return n, r.Key, err
		} else if entry.Len() > 0 {
			if len(r.Key) > 0 {
				if !r.First {
					return nil, nil, nil
				}
				// looping not very efficient, but we do not have an index
				for i := 0; i < entry.Len(); i++ {
					// TODO: Support compound keys
					candidate := entry.Item(i)
					candidateKey := SetValues(r.Meta.KeyMeta(), candidate[self.MetaIdent(r.Selection, r.Meta.KeyMeta()[0])])
					if r.Key[0].Equal(candidateKey[0]) {
						selected = candidate
						break
					}
				}
			} else {
				if int(r.Row) < entry.Len() {
					selected = entry.Item(r.Row)
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
