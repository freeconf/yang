package node

import (
	"github.com/c2g/meta"
)

type Request struct {
	Selection   *Selection
	Target      PathSlice
	Context     Context
}

type Stream interface {
	Send(Node)
	RegisterCloser(f func())
}

type NotifyRequest struct {
	Request
	Meta *meta.Notification
	Stream Stream
}

type ActionRequest struct {
	Request
	Meta  *meta.Rpc
	Input *Selection
}

type ContainerRequest struct {
	Request
	New  bool
	Meta meta.MetaList
}

type ListRequest struct {
	Request
	New      bool
	StartRow int64
	Row      int64
	First    bool
	Meta     *meta.List
	Key      []*Value
}

type FieldRequest struct {
	Request
	Meta meta.HasDataType
}

type WalkController interface {
	ContainerIterator(sel *Selection, m meta.MetaList) (meta.MetaIterator, error)
	VisitList(r *ListRequest) (next *Selection, err error)
	VisitContainer(r *ContainerRequest) (child *Selection, err error)
	VisitNotification(r *NotifyRequest) (*Selection, error)
	VisitAction(r *ActionRequest) (*Selection, error)
	VisitField(r *FieldRequest) (*Value, error)
}
