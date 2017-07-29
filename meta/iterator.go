package meta

type MetaIterator interface {
	NextMeta() (Meta, error)
	HasNextMeta() bool
}

type MetaListIterator struct {
	position       Meta
	next           Meta
	err            error
	currentProxy   MetaIterator
	resolveProxies bool
}

type EmptyInterator struct{}

func (EmptyInterator) HasNextMeta() bool {
	return false
}
func (EmptyInterator) NextMeta() (Meta, error) {
	return nil, nil
}

type SingletonIterator struct {
	Meta Meta
}

func (s *SingletonIterator) HasNextMeta() bool {
	return s.Meta != nil
}
func (s *SingletonIterator) NextMeta() (Meta, error) {
	m := s.Meta
	s.Meta = nil
	return m, nil
}

func NewMetaListIterator(m Meta, resolveProxies bool) MetaIterator {
	list, isMetaList := m.(MetaList)
	if !isMetaList {
		return EmptyInterator(struct{}{})
	}
	i := &MetaListIterator{position: list.GetFirstMeta(), resolveProxies: resolveProxies}
	i.next, i.err = i.lookAhead()
	return i
}

func (self *MetaListIterator) HasNextMeta() bool {
	return self.next != nil
}

func (self *MetaListIterator) NextMeta() (next Meta, err error) {
	next = self.next
	err = self.err
	self.next, self.err = self.lookAhead()
	return next, err
}

func (self *MetaListIterator) lookAhead() (Meta, error) {
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
					return next, nil
				} else {
					self.position = self.position.GetSibling()
					self.currentProxy = proxy.ResolveProxy()
				}
			} else {
				next := self.position
				self.position = self.position.GetSibling()
				return next, nil
			}
		}
	}
	return nil, nil
}
