package meta

type MetaIterator interface {
	NextMeta() Meta
	HasNextMeta() bool
}

type MetaListIterator struct {
	position       Meta
	next           Meta
	currentProxy   MetaIterator
	resolveProxies bool
}

type EmptyInterator int

func (e EmptyInterator) HasNextMeta() bool {
	return false
}
func (e EmptyInterator) NextMeta() Meta {
	return nil
}

type SingletonIterator struct {
	Meta Meta
}

func (s *SingletonIterator) HasNextMeta() bool {
	return s.Meta != nil
}
func (s *SingletonIterator) NextMeta() Meta {
	m := s.Meta
	s.Meta = nil
	return m
}

func NewMetaListIterator(m Meta, resolveProxies bool) MetaIterator {
	list, isMetaList := m.(MetaList)
	if !isMetaList {
		return nil
	}
	i := &MetaListIterator{position: list.GetFirstMeta(), resolveProxies: resolveProxies}
	i.next = i.lookAhead()
	return i
}

func (self *MetaListIterator) HasNextMeta() bool {
	return self.next != nil
}

func (self *MetaListIterator) NextMeta() (next Meta) {
	next, self.next = self.next, self.lookAhead()
	return next
}

func (self *MetaListIterator) lookAhead() Meta {
	for self.position != nil || self.currentProxy != nil {
		if self.currentProxy != nil {
			if self.currentProxy.HasNextMeta() {
				return self.currentProxy.NextMeta()
			}
			self.currentProxy = nil
		} else {
			if self.resolveProxies {
				proxy, isProxy := self.position.(MetaProxy)
				if !isProxy {
					next := self.position
					self.position = self.position.GetSibling()
					return next
				} else {
					self.position = self.position.GetSibling()
					self.currentProxy = proxy.ResolveProxy()
				}
			} else {
				next := self.position
				self.position = self.position.GetSibling()
				return next
			}
		}
	}
	return nil
}
