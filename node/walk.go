package node

import (
	"github.com/c2stack/c2g/meta"
)

func (self *Selection) Walk(controller WalkController) (err error) {
	if meta.IsList(self.path.meta) && !self.insideList {
		r := ListRequest{
			Request:Request {
				Selection: self,
			},
			First: true,
			Meta: self.path.meta.(*meta.List),
		}
		var next *Selection
		if next, err = controller.VisitList(&r); err != nil || next == nil {
			return
		}
		for next != nil {
			if err = next.Walk(controller); err != nil {
				return
			}
			if err = next.Fire(LEAVE.New(next)); err != nil {
				return err
			}
			r.First = false
			r.Row++
			if next, err = controller.VisitList(&r); err != nil {
				return
			}
		}
	} else {
		i, cerr := controller.ContainerIterator(self, self.path.meta.(meta.MetaList))
		if cerr != nil || i == nil {
			return cerr
		}
		return  self.walkIterator(controller, i)
	}
	return
}

func (self *Selection) walkIterator(controller WalkController, i meta.MetaIterator) (err error) {
	for i.HasNextMeta() {
		m := i.NextMeta()
		if choice, isChoice := m.(*meta.Choice); isChoice {
			var chosen *meta.ChoiceCase
			if chosen, err = self.node.Choose(self, choice); err != nil {
				return
			} else if chosen != nil {
				choiceIterator, choiceErr := controller.ContainerIterator(self, chosen)
				if choiceErr != nil {
					return choiceErr
				}
				return self.walkIterator(controller, choiceIterator)
			}
		} else if meta.IsLeaf(m) {
			// only walking here, not interested in value
			r := FieldRequest{
				Request:Request {
					Selection: self,
				},
				Meta: m.(meta.HasDataType),
			}
			if err = controller.VisitField(&r); err != nil {
				return err
			}
		} else {
			mList := m.(meta.MetaList)
			if meta.IsAction(m) {
				r := ActionRequest{
					Request:Request {
						Selection: self,
					},
					Meta: m.(*meta.Rpc),
				}
				if _, err = controller.VisitAction(&r); err != nil {
					return err
				}
			} else if notif, isNotification := m.(*meta.Notification); isNotification {
				r := NotifyRequest {
					Request:Request {
						Selection: self,
					},
					Meta: notif,
				}
				if _, err = controller.VisitNotification(&r); err != nil {
					return err
				}
			} else {
				r := ContainerRequest {
					Request:Request {
						Selection: self,
					},
					Meta: mList,
				}
				childSel, childErr := controller.VisitContainer(&r)
				if childErr != nil {
					return childErr
				} else if childSel == nil {
					continue
				}

				if err = childSel.Walk(controller); err != nil {
					return
				}
				if err = childSel.Fire(LEAVE.New(childSel)); err != nil {
					return err
				}
			}
		}
	}
	return
}
