package c2

import "container/list"

type Subscription interface {
	Close() error
}

func NewSubscription(l *list.List, e *list.Element) Subscription {
	return &listSubscription{l, e}
}

type listSubscription struct {
	l *list.List
	e *list.Element
}

func (self *listSubscription) Close() error {
	self.l.Remove(self.e)
	return nil
}
