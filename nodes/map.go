package nodes

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"
)

// Decode and encode data from Go's map and slices of interface{}.  Will assume collections recursively
// but On{Hook} functions provide option to alter this at any place.

type Collection struct {
	OnKey        MetaIdentFunc
	OnExtendMap  ExtendMapFunc
	OnExtendList ExtendListFunc
}

func MapNode(container map[string]interface{}) node.Node {
	return (&Collection{}).Node(container)
}

func ListNode(entry MappedListHandler) node.Node {
	return (&Collection{}).List(entry)
}

type MetaIdentFunc func(sel node.Selection, m meta.Meta) (key string)
type ExtendMapFunc func(sel node.Selection, m meta.MetaList, container map[string]interface{}) (node.Node, error)
type ExtendListFunc func(sel node.Selection, m *meta.List, entry MappedListHandler) (node.Node, error)

type MapEntry struct {
	Key    string
	Parent map[string]interface{}

	// data can come in both these formats, so support either
	listOption1 []interface{}

	listOption2 []map[string]interface{}
}

type MappedListHandler interface {
	Append(item map[string]interface{})
	Remove(i int)
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

func (self *MapEntry) Remove(i int) {
	if self.listOption1 != nil {
		copy(self.listOption1[i:], self.listOption1[i+1:])
		self.listOption1[len(self.listOption1)-1] = nil
		self.listOption1 = self.listOption1[:len(self.listOption1)-1]
		self.Parent[self.Key] = self.listOption1
	} else {
		copy(self.listOption2[i:], self.listOption2[i+1:])
		self.listOption2[len(self.listOption2)-1] = nil
		self.listOption2 = self.listOption2[:len(self.listOption2)-1]
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

func (self *Collection) Node(container map[string]interface{}) node.Node {
	s := &Basic{}
	s.OnChild = func(r node.ChildRequest) (node.Node, error) {
		var data interface{}
		keyIdent := self.MetaIdent(r.Selection, r.Meta)
		if r.New {
			if meta.IsList(r.Meta) {
				data = make([]map[string]interface{}, 0, 10)
			} else {
				data = make(map[string]interface{})
			}
			container[keyIdent] = data
		} else if r.Delete {
			delete(container, r.Meta.GetIdent())
		} else {
			data = container[keyIdent]
		}
		if data != nil {
			if meta.IsList(r.Meta) {
				me := &MapEntry{
					Key:    keyIdent,
					Parent: container,
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
	s.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		if r.Write {
			err = self.UpdateLeaf(r.Selection, container, r.Meta, hnd.Val)
		} else {
			hnd.Val, err = self.ReadLeaf(r.Selection, container, r.Meta)
		}
		return
	}
	s.OnChoose = func(sel node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
		cases := meta.NewMetaListIterator(choice, false)
		for cases.HasNextMeta() {
			m, err := cases.NextMeta()
			if err != nil {
				return nil, err
			}
			kase := m.(*meta.ChoiceCase)
			props := meta.NewMetaListIterator(kase, true)
			for props.HasNextMeta() {
				prop, err := props.NextMeta()
				if err != nil {
					return nil, err
				}
				if _, found := container[prop.GetIdent()]; found {
					return kase, nil
				}
				// just because you didn't find a property doesnt
				// mean it's invalid, it's only if you don't find any
				// of the properties of a case
			}
		}
		return nil, nil
	}
	return s
}

func (self *Collection) ExtendContainer(sel node.Selection, m meta.MetaList, data interface{}) (node.Node, error) {
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

func (self *Collection) ExtendList(sel node.Selection, m *meta.List, entry *MapEntry) (node.Node, error) {
	if self.OnExtendList != nil {
		if n, err := self.OnExtendList(sel, m, entry); n != nil || err != nil {
			return n, err
		}
	}
	return self.List(entry), nil
}

func (self *Collection) MetaIdent(sel node.Selection, m meta.Meta) string {
	if self.OnKey != nil {
		return self.OnKey(sel, m)
	}

	return m.GetIdent()
}

func (self *Collection) ReadKey(sel node.Selection, container map[string]interface{}, m *meta.List) (key []val.Value, err error) {
	keyMeta := m.KeyMeta()
	key = make([]val.Value, len(keyMeta))
	for i, m := range keyMeta {
		if key[i], err = self.ReadLeaf(sel, container, m); err != nil {
			return nil, err
		}
	}
	return
}

func (self *Collection) findByKeyValue(key []val.Value, meta *meta.List, entry MappedListHandler) int {
	keyMeta := meta.KeyMeta()
	// looping not very efficient, but we do not have an index
	for i := 0; i < entry.Len(); i++ {
		candidate := entry.Item(i)

		// TODO : Support compound keys
		candidateKey := candidate[meta.Key[0]]
		candidateKeyValue, _ := node.NewValues(keyMeta, candidateKey)
		if val.Equal(key[0], candidateKeyValue[0]) {
			return i
		}
	}
	return -1
}

func (self *Collection) List(entry MappedListHandler) node.Node {
	s := &Basic{}
	s.OnNext = func(r node.ListRequest) (node.Node, []val.Value, error) {
		var selected map[string]interface{}
		if r.New {
			selection := make(map[string]interface{})
			entry.Append(selection)
			n, err := self.ExtendContainer(r.Selection, r.Meta, selection)
			return n, r.Key, err
		} else if r.Delete {
			if found := self.findByKeyValue(r.Key, r.Meta, entry); found >= 0 {
				entry.Remove(found)
			}
		} else if entry.Len() > 0 {
			if len(r.Key) > 0 {
				if !r.First {
					return nil, nil, nil
				}
				if found := self.findByKeyValue(r.Key, r.Meta, entry); found >= 0 {
					selected = entry.Item(found)
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

func (self *Collection) ReadLeaf(sel node.Selection, container map[string]interface{}, m meta.HasDataType) (val.Value, error) {
	return node.NewValue(m.GetDataType(), container[self.MetaIdent(sel, m)])
}

func (self *Collection) UpdateLeaf(sel node.Selection, container map[string]interface{}, m meta.HasDataType, v val.Value) error {
	container[self.MetaIdent(sel, m)] = v.Value()
	return nil
}
