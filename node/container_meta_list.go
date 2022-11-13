package node

import (
	"fmt"

	"github.com/freeconf/yang/meta"
)

type containerMetaList struct {
	next       meta.Meta
	main       metaIterator
	choiceCase *containerMetaList
	s          Selection
}

type metaIterator interface {
	hasNextMeta() bool
	nextMeta() meta.Meta
}

func newContainerMetaList(s Selection) *containerMetaList {
	i := &containerMetaList{
		main: newMetaIterator(s.Path.Meta.(meta.HasDataDefinitions).DataDefinitions()),
		s:    s,
	}
	i.lookAhead()
	return i
}

type defIterator struct {
	i    int
	defs []meta.Definition
}

func newMetaIterator(defs []meta.Definition) metaIterator {
	return &defIterator{
		defs: defs,
	}
}

func (iter *defIterator) hasNextMeta() bool {
	return iter.i < len(iter.defs)
}

func (iter *defIterator) nextMeta() meta.Meta {
	d := iter.defs[iter.i]
	iter.i++
	return d
}

func newChoiceCaseIterator(s Selection, m *meta.ChoiceCase) *containerMetaList {
	i := &containerMetaList{
		main: newMetaIterator(m.DataDefinitions()),
		s:    s,
	}
	i.lookAhead()
	return i
}

func (self *containerMetaList) nextMeta() meta.Meta {
	var next = self.next
	self.lookAhead()
	return next
}

func (self *containerMetaList) lookAhead() {
	self.next = nil
	var m meta.Meta
	for {
		if self.choiceCase != nil {
			m = self.choiceCase.nextMeta()
			if m == nil {
				self.choiceCase = nil
				continue
			}
		} else if self.main != nil {
			if self.main.hasNextMeta() {
				m = self.main.nextMeta()
			} else {
				self.main = nil
				break
			}
		}
		if choice, isChoice := m.(*meta.Choice); isChoice {
			if chosen, err := self.s.Node.Choose(self.s, choice); err != nil {
				panic(fmt.Sprintf("%T - %s", self.s.Node, err))
			} else if chosen != nil {
				self.choiceCase = newChoiceCaseIterator(self.s, chosen)
				continue
			}
		} else {
			self.next = m
			break
		}
	}
}
