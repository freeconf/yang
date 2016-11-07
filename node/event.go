package node

type EventType int

type Event struct {
	Type    EventType
	Src     Selection
}

func (self Event) String() string {
	return self.Type.String()
}

const (
	BEGIN_EDIT EventType = iota + 1

	// TODO: Consider making these event propagate up tree until handler cancel's
	// propagation (like w3c DOM mouse click events)
	END_EDIT
	DELETE
)

var eventNames = []string{
	"<invalid event id>",
	"BEGIN_EDIT",
	"END_EDIT",
	"DELETE",
}

func (e EventType) String() string {
	return eventNames[e]
}
