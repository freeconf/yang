package node

import (
	"blit"
	"fmt"
	"regexp"
)

type EventType int

type eventState struct {
	propagationStopped bool
}

type Event struct {
	Type  EventType
	Details interface{}
	state *eventState
}

type FetchDetails struct {
	Path *Path
}

func (self Event) String() string {
	return self.Type.String()
}

func (self EventType) New() Event {
	return Event{Type: self, state : &eventState{}}
}

func (self EventType) NewWithDetails(details interface{}) Event {
	return Event{Type: self, Details: details, state : &eventState{}}
}

func (self EventType) Bubbles() bool {
	switch self {
	case START_TREE_EDIT, END_TREE_EDIT:
		return true
	}
	return false
}

func (self Event) StopPropagation() {
	self.state.propagationStopped = true
}

const (
	NEW EventType = iota + 1

	// TODO: Consider making these event propagate up tree until handler cancel's
	// propagation (like w3c DOM mouse click events)
	START_TREE_EDIT
	END_TREE_EDIT

	LEAVE_EDIT
	LEAVE
	DELETE
	UNSELECT
)

var eventNames = []string{
	"<invalid event id>",
	"NEW",
	"START_TREE_EDIT",
	"END_TREE_EDIT",

	"LEAVE_EDIT",
	"LEAVE",
	"DELETE",
	"UNSELECT",
}

type Events interface {
	AddListener(*Listener)
	RemoveListener(*Listener)
	Fire(path *Path, e Event) error
}

type EventsImpl struct {
	Parent    Events
	listeners []*Listener
}

type ListenFunc func() error

func (e EventType) String() string {
	return eventNames[e]
}

type Listener struct {
	path    string
	regex   *regexp.Regexp
	event   EventType
	handler ListenFunc
}

func (l *Listener) String() string {
	if len(l.path) > 0 {
		return fmt.Sprintf("%s:%s=>%p", l.event, l.path, l.handler)
	}
	return fmt.Sprintf("%s:%v=>%p", l.event, l.regex, l.handler)
}

func (impl *EventsImpl) dump() {
	for _, l := range impl.listeners {
		blit.Debug.Print(l.String())
	}
}

func (impl *EventsImpl) AddListener(l *Listener) {
	impl.listeners = append(impl.listeners, l)
}

func (impl *EventsImpl) RemoveListener(l *Listener) {
	for i, candidate := range impl.listeners {
		if l == candidate {
			impl.listeners = append(impl.listeners[:i], impl.listeners[i+1:]...)
			break
		}
	}
}

func (impl *EventsImpl) Fire(path *Path, e Event) (err error) {
	if len(impl.listeners) > 0 {
		pathStr := path.String()
		for _, l := range impl.listeners {
			if l.event == e.Type {
				if len(l.path) > 0 {
					if l.path != pathStr {
						continue
					}
				} else if l.regex != nil {
					if !l.regex.MatchString(pathStr) {
						continue
					}
				}
				if err = l.handler(); err != nil {
					return err
				}
			}
		}
	}
	if impl.Parent != nil {
		if err = impl.Parent.Fire(path, e); err != nil {
			return err
		}
	}
	return nil
}
