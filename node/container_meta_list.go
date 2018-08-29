package node

import (
	"fmt"

	"github.com/freeconf/gconf/meta"
)

type ContainerMetaList struct {
	next       meta.Meta
	main       meta.Iterator
	choiceCase *ContainerMetaList
	s          Selection
}

func NewContainerMetaList(s Selection) *ContainerMetaList {
	i := &ContainerMetaList{
		main: meta.Iterate(s.Path.meta.(meta.HasDataDefs).DataDefs()),
		s:    s,
	}
	i.lookAhead()
	return i
}

func newChoiceCaseIterator(s Selection, m *meta.ChoiceCase) *ContainerMetaList {
	i := &ContainerMetaList{
		main: meta.Iterate(m.DataDefs()),
		s:    s,
	}
	i.lookAhead()
	return i
}

func (self *ContainerMetaList) Next() meta.Meta {
	var next = self.next
	self.lookAhead()
	return next
}

func (self *ContainerMetaList) lookAhead() {
	self.next = nil
	var m meta.Meta
	for {
		if self.choiceCase != nil {
			m = self.choiceCase.Next()
			if m == nil {
				self.choiceCase = nil
				continue
			}
		} else if self.main != nil {
			if self.main.HasNext() {
				m = self.main.Next()
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
