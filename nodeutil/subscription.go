package nodeutil

import "container/list"

// Subscription is handle into a list.List that when closed
// will automatically remove item from list.  Useful for maintaining
// a set of listeners that can easily remove themselves.
type Subscription interface {
	Close() error
}

// NewSubscription is used by subscription managers to give a token
// to caller the can close to unsubscribe to events
func NewSubscription(l *list.List, e *list.Element) Subscription {
	return &listSubscription{l, e}
}

type listSubscription struct {
	l *list.List
	e *list.Element
}

// Close will unsubscribe to events.
func (self *listSubscription) Close() error {
	self.l.Remove(self.e)
	return nil
}
