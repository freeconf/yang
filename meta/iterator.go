package meta

// Iterator iterates over meta. Use meta.Children for most common way to
// iterate.
type Iterator interface {
	Next() (Meta, error)
	HasNext() bool
}

type iterator struct {
	position       Meta
	next           Meta
	err            error
	currentProxy   Iterator
	resolveProxies bool
}

type empty struct{}

func (empty) HasNext() bool {
	return false
}
func (empty) Next() (Meta, error) {
	return nil, nil
}

type single struct {
	Meta Meta
}

func (s *single) HasNext() bool {
	return s.Meta != nil
}

func (s *single) Next() (Meta, error) {
	m := s.Meta
	s.Meta = nil
	return m, nil
}

// Children of a meta list returning only containers, lists and leafs
func Children(m MetaList) Iterator {
	return children(m, true)
}

// Children of a meta list returning items as they are written in YANG file
// including groupings, uses, and choices
func ChildrenNoResolve(m MetaList) Iterator {
	return children(m, false)
}

func children(m MetaList, resolveProxies bool) Iterator {
	i := &iterator{position: m.GetFirstMeta(), resolveProxies: resolveProxies}
	i.next, i.err = i.lookAhead()
	return i
}

func (self *iterator) HasNext() bool {
	return self.next != nil
}

func (self *iterator) Next() (next Meta, err error) {
	next = self.next
	err = self.err
	self.next, self.err = self.lookAhead()
	return next, err
}

func (self *iterator) lookAhead() (Meta, error) {
	for self.position != nil || self.currentProxy != nil {
		if self.currentProxy != nil {
			if self.currentProxy.HasNext() {
				return self.currentProxy.Next()
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
