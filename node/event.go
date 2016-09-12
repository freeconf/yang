package node

type EventType int

type eventState struct {
	propagationStopped bool
}

type Event struct {
	Type    EventType
	Src     Selection
	Details interface{}
	state   *eventState
}

func (self Event) String() string {
	return self.Type.String()
}

func (self EventType) New(src Selection) Event {
	return Event{Type: self, Src: src, state: &eventState{}}
}

func (self EventType) NewWithDetails(src Selection, details interface{}) Event {
	return Event{Type: self, Src: src, Details: details, state: &eventState{}}
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
	REMOVE_LIST_ITEM
	REMOVE_CONTAINER
	ADD_CONTAINER
)

var eventNames = []string{
	"<invalid event id>",
	"NEW",
	"START_TREE_EDIT",
	"END_TREE_EDIT",

	"LEAVE_EDIT",
	"LEAVE",
	"DELETE",
	"REMOVE_LIST_ITEM",
	"REMOVE_CONTAINER",
	"ADD_CONTAINER",
}

func (e EventType) String() string {
	return eventNames[e]
}

