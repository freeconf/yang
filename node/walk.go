package node

//import (
//	"github.com/c2stack/c2g/meta"
//)
//
//type WalkState struct {
//	stack    []Selection
//	stackPos int
//}
//
//func NewWalkState() *WalkState {
//	return &WalkState{
//		stack : make([]Selection, 32),
//	}
//}
//
//func (self *WalkState) Push(s Selection) {
//	self.stackPos++
//	self.stack = append(self.stack)
//	self.stack[self.stackPos] = s
//}
//
//func (self *WalkState) Pop() {
//	self.stackPos--
//}
//
//
//// Walk is at the root of almost all operations that need to find, read or write a data source
//// using a given model.  Controller navigates the operation and potentially gathers whatever data
//// it's looking for.
//func (self Selection) Walk(controller WalkController) (err error) {
//	if meta.IsList(self.Path.meta) && !self.InsideList {
//		r := ListRequest{
//			Request:Request {
//				Selection: self,
//			},
//			First: true,
//			Meta: self.Path.meta.(*meta.List),
//		}
//		var next Selection
//		if next, err = controller.VisitList(&r); err != nil || next.IsNil() {
//			return
//		}
//		for ! next.IsNil() {
//			if err = next.Walk(controller); err != nil {
//				return
//			}
//			if err = next.Fire(LEAVE.New(next)); err != nil {
//				return err
//			}
//			r.First = false
//			r.Row++
//			if next, err = controller.VisitList(&r); err != nil {
//				return
//			}
//		}
//	} else {
//		i, cerr := controller.ContainerIterator(self, self.Path.meta.(meta.MetaList))
//		if cerr != nil || i == nil {
//			return cerr
//		}
//		return self.walkIterator(controller, i)
//	}
//	return
//}
//
//func (self Selection) walkIterator(controller WalkController, i meta.MetaIterator) (err error) {
//	for i.HasNextMeta() {
//		m := i.NextMeta()
//		if choice, isChoice := m.(*meta.Choice); isChoice {
//			var chosen *meta.ChoiceCase
//			if chosen, err = self.Node.Choose(self, choice); err != nil {
//				return
//			} else if chosen != nil {
//				choiceIterator, choiceErr := controller.ContainerIterator(self, chosen)
//				if choiceErr != nil {
//					return choiceErr
//				}
//				return self.walkIterator(controller, choiceIterator)
//			}
//		} else if meta.IsLeaf(m) {
//			// only walking here, not interested in value
//			r := FieldRequest{
//				Request:Request {
//					Selection: self,
//				},
//				Meta: m.(meta.HasDataType),
//			}
//			if err = controller.VisitField(&r); err != nil {
//				return err
//			}
//		} else {
//			mList := m.(meta.MetaList)
//			if meta.IsAction(m) {
//				r := ActionRequest{
//					Request:Request {
//						Selection: self,
//					},
//					Meta: m.(*meta.Rpc),
//				}
//				if _, err = controller.VisitAction(&r); err != nil {
//					return err
//				}
//			} else if notif, isNotification := m.(*meta.Notification); isNotification {
//				r := NotifyRequest {
//					Request:Request {
//						Selection: self,
//					},
//					Meta: notif,
//				}
//				if _, err = controller.VisitNotification(&r); err != nil {
//					return err
//				}
//			} else {
//				r := ContainerRequest {
//					Request:Request {
//						Selection: self,
//					},
//					Meta: mList,
//				}
//				childSel, childErr := controller.VisitContainer(&r)
//				if childErr != nil {
//					return childErr
//				} else if childSel.IsNil() {
//					continue
//				}
//
//				if err = childSel.Walk(controller); err != nil {
//					return
//				}
//				if err = childSel.Fire(LEAVE.New(childSel)); err != nil {
//					return err
//				}
//			}
//		}
//	}
//	return
//}
