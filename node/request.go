package node

import (
	"time"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
)

// Request is base class for all other node requests.  There are two basic modes:
// 1. Navigation where NavTarget is set and 2.)Editing where WalkBase is set
type Request struct {
	Selection Selection

	// Path to meta item requested, including leaf requests
	Path *Path

	Target *Path
	Base   *Path
}

// NotifyCloser callback when caller is not interested in events anymore. Typically
// this is where you remove listeners
type NotifyCloser func() error

type Notification struct {
	EventTime time.Time
	Event     Selection
}

func NewNotification(msg Selection) Notification {
	return NewNotificationWhen(msg, time.Now())
}

func NewNotificationWhen(msg Selection, t time.Time) Notification {
	return Notification{
		EventTime: t,
		Event:     msg,
	}
}

// NotifyStream is pipe back to subscriber.
type NotifyStream func(n Notification)

type NotifyRequest struct {
	Request
	Meta   *meta.Notification
	Stream NotifyStream
}

func (self NotifyRequest) Send(n Node) {
	self.SendWhen(n, time.Now())
}

func (self NotifyRequest) SendWhen(n Node, t time.Time) {
	s := Selection{
		Parent:      &self.Selection,
		Browser:     self.Selection.Browser,
		Path:        &Path{Meta: self.Meta},
		Node:        n,
		Constraints: self.Selection.Constraints,
		Context:     self.Selection.Context,
	}
	self.Stream(NewNotificationWhen(s, t))
}

type ActionRequest struct {
	Request
	Meta  *meta.Rpc
	Input Selection
}

type NodeRequest struct {
	Selection Selection
	New       bool
	Delete    bool
	Source    Selection
	EditRoot  bool
}

type ChildRequest struct {
	Request
	From   Selection
	New    bool
	Delete bool
	Meta   meta.HasDataDefinitions
}

func (self *ChildRequest) IsNavigation() bool {
	return self.Target != nil
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
	Key        []val.Value
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
	self.First = false
}

func (self *ListRequest) IsNavigation() bool {
	return self.Target != nil
}

type FieldRequest struct {
	Request
	Meta  meta.Leafable
	Write bool
	Clear bool
}

func (self *FieldRequest) IsNavigation() bool {
	return self.Target != nil
}
