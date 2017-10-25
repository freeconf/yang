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

func ErrIterator(err error) Iterator {
	return &errIter{err: err}
}

type errIter struct {
	err error
}

func (self *errIter) HasNext() bool {
	return self.err != nil
}

func (self *errIter) Next() (Meta, error) {
	e := self.err
	self.err = nil
	return nil, e
}

func SingleIterator(m Meta) Iterator {
	return &singleIter{m: m}
}

type singleIter struct {
	m Meta
}

func (s *singleIter) HasNext() bool {
	return s.m != nil
}

func (s *singleIter) Next() (Meta, error) {
	m := s.m
	s.m = nil
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
