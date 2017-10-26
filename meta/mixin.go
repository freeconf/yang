package meta

import "sort"

type mixinItem struct {
	m Meta
	n int
}

type mixinSet []*mixinItem

func (self mixinSet) Len() int {
	return len(self)
}

func (self mixinSet) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self mixinSet) Less(i, j int) bool {
	return self[i].n < self[j].n
}

func mixin(base MetaList, augment MetaList) error {
	metas := make(map[string]*mixinItem)
	n := 0
	for i := Children(base); i.HasNext(); n++ {
		meta, err := i.Next()
		if err != nil {
			return err
		}
		metas[meta.GetIdent()] = &mixinItem{n: n, m: meta}
	}
	for i := Children(augment); i.HasNext(); n++ {
		meta, err := i.Next()
		if err != nil {
			return err
		}
		if existing, exists := metas[meta.GetIdent()]; exists {
			existing.m = meta
		} else {
			metas[meta.GetIdent()] = &mixinItem{n: n, m: meta}
		}
	}
	set := make(mixinSet, len(metas))
	n = 0
	for _, item := range metas {
		set[n] = item
		n++
	}
	sort.Sort(set)
	base.Clear()
	for _, item := range set {
		base.AddMeta(item.m)
	}
	return nil
}
