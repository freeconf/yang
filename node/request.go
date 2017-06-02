package node

import "github.com/c2stack/c2g/meta"

// Request is base class for all other node requests.  There are two basic modes:
// 1. Navigation where NavTarget is set and 2.)Editing where WalkBase is set
type Request struct {
	Selection          Selection
	Path               *Path
	Target             *Path
	Base               *Path
	Constraints        *Constraints
	ConstraintsHandler *ConstraintHandler
}

// NotifyCloser callback when caller is not interested in events anymore. Typically
// this is where you remove listeners
type NotifyCloser func() error

// NotifyStream is pipe back to subscriber.
type NotifyStream func(msg Selection)

type NotifyRequest struct {
	Request
	Meta   *meta.Notification
	Stream NotifyStream
}

func (self NotifyRequest) Send(n Node) {
	s := Selection{
		Parent:      &self.Selection,
		Browser:     self.Selection.Browser,
		Path:        NewRootPath(self.Meta),
		Node:        n,
		Constraints: self.Constraints,
		Context:     self.Selection.Context,
	}
	self.Stream(s)
}

type ActionRequest struct {
	Request
	Meta  *meta.Rpc
	Input Selection
}

type NodeRequest struct {
	Selection Selection
	New       bool
	Source    Selection
	EditRoot  bool
}

type ChildRequest struct {
	Request
	From   Selection
	New    bool
	Delete bool
	Meta   meta.MetaList
}

type ListRequest struct {
	Request
	From   Selection
	New    bool
	Delete bool

	StartRow int

	// We make row available as a 32bit value for convenience but in theory
	// could be 64bit.  If you know you're list could not exceed 2 billion then
	// it's safe to use this value
	Row int

	StartRow64 int64
	Row64      int64
	First      bool
	Meta       *meta.List
	Key        []*Value
}

func (self *ListRequest) SetStartRow(row int64) {
	self.StartRow64 = row
	self.StartRow = int(row)
}

func (self *ListRequest) SetRow(row int64) {
	self.Row64 = row
	self.Row = int(row)
}

func (self *ListRequest) IncrementRow() {
	self.Row64++
	self.Row++
}

type FieldRequest struct {
	Request
	Meta  meta.HasDataType
	Write bool
}

func (self *ChildRequest) IsNavigation() bool {
	return self.Target != nil
}

func (self *ListRequest) IsNavigation() bool {
	return self.Target != nil
}

func (self *FieldRequest) IsNavigation() bool {
	return self.Target != nil
}
