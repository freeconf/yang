package node

import (
	"time"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
)

// Request is base class for all other node requests.  There are two basic modes:
// 1. Navigation where NavTarget is set and 2.)Editing where WalkBase is set
type Request struct {
	Selection *Selection

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
	Event     *Selection
}

func NewNotification(msg *Selection) Notification {
	return NewNotificationWhen(msg, time.Now())
}

func NewNotificationWhen(msg *Selection, t time.Time) Notification {
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

func (r NotifyRequest) Send(n Node) {
	r.SendWhen(n, time.Now())
}

func (r NotifyRequest) SendWhen(n Node, t time.Time) {
	s := &Selection{
		parent:      r.Selection,
		Browser:     r.Selection.Browser,
		Path:        r.Selection.Path,
		Node:        n,
		Constraints: r.Selection.Constraints,
		Context:     r.Selection.Context,
	}
	r.Stream(NewNotificationWhen(s, t))
}

type ActionRequest struct {
	Request
	Meta  *meta.Rpc
	Input *Selection
}

type NodeRequest struct {
	Selection *Selection
	New       bool
	Delete    bool
	Source    *Selection
	EditRoot  bool
}

type ChildRequest struct {
	Request
	New    bool
	Delete bool
	Meta   meta.HasDataDefinitions
}

func (r *ChildRequest) IsNavigation() bool {
	return r.Target != nil
}

type ListRequest struct {
	Request
	From   *Selection
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

func (r *ListRequest) SetStartRow(row int64) {
	r.StartRow64 = row
	r.StartRow = int(row)
}

func (r *ListRequest) SetRow(row int64) {
	r.Row64 = row
	r.Row = int(row)
}

func (r *ListRequest) IncrementRow() {
	r.Row64++
	r.Row++
	r.First = false
}

func (r *ListRequest) IsNavigation() bool {
	return r.Target != nil
}

type FieldRequest struct {
	Request
	From  *Selection
	Meta  meta.Leafable
	Write bool
	Clear bool
}

func (r *FieldRequest) IsNavigation() bool {
	return r.Target != nil
}
