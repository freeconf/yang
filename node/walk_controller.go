package node

import (
	"github.com/c2stack/c2g/meta"
)

type WalkController interface {
	ContainerIterator(sel Selection, m meta.MetaList) (meta.MetaIterator, error)
	VisitList(r *ListRequest) (next Selection, err error)
	VisitContainer(r *ContainerRequest) (child Selection, err error)
	VisitNotification(r *NotifyRequest) (Selection, error)
	VisitAction(r *ActionRequest) (Selection, error)
	VisitField(r *FieldRequest) error
}
