package node

type EventType int

type Event struct {
	Type EventType
	Src  Selection
}

func (self Event) String() string {
	return self.Type.String()
}

const (
	EDIT EventType = iota + 1
	DELETE
)

var eventNames = []string{
	"<invalid event id>",
	"EDIT",
	"DELETE",
}

func (e EventType) String() string {
	return eventNames[e]
}
