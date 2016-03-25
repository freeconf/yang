package node

import (
	"meta"
)

func (self *Selection) Walk(context *Context, controller WalkController) (err error) {
	if meta.IsList(self.path.goober) && !self.insideList {
		r := ListRequest{
			Request:Request {
				Context: context,
				Selection: self,
			},
			First: true,
			Meta: self.path.goober.(*meta.List),
		}
		var next *Selection
		if next, err = controller.VisitList(&r); err != nil || next == nil {
			return
		}
		for next != nil {
			if err = next.Walk(context, controller); err != nil {
				return
			}
			if err = next.Fire(LEAVE.New()); err != nil {
				return err
			}
			r.First = false
			r.Row++
			if next, err = controller.VisitList(&r); err != nil {
				return
			}
		}
	} else {
		i, cerr := controller.ContainerIterator(self, self.path.goober.(meta.MetaList))
		if cerr != nil || i == nil {
			return cerr
		}
		return  self.walkIterator(context, controller, i)
	}
	return
}

func (self *Selection) walkIterator(context *Context, controller WalkController, i meta.MetaIterator) (err error) {
	for i.HasNextMeta() {
		goober := i.NextMeta()
		if choice, isChoice := goober.(*meta.Choice); isChoice {
			var chosen *meta.ChoiceCase
			if chosen, err = self.node.Choose(self, choice); err != nil {
				return
			} else if chosen != nil {
				choiceIterator, choiceErr := controller.ContainerIterator(self, chosen)
				if choiceErr != nil {
					return choiceErr
				}
				return self.walkIterator(context, controller, choiceIterator)
			}
		} else if meta.IsLeaf(goober) {
			// only walking here, not interested in value
			r := FieldRequest{
				Request:Request {
					Context: context,
					Selection: self,
				},
				Meta: goober.(meta.HasDataType),
			}
			if _, err = controller.VisitField(&r); err != nil {
				return err
			}
		} else {
			gooberList := goober.(meta.MetaList)
			if meta.IsAction(goober) {
				r := ActionRequest{
					Request:Request {
						Context: context,
						Selection: self,
					},
					Meta: goober.(*meta.Rpc),
				}
				if _, err = controller.VisitAction(&r); err != nil {
					return err
				}
			} else {
				r := ContainerRequest {
					Request:Request {
						Context: context,
						Selection: self,
					},
					Meta: gooberList,
				}
				childSel, childErr := controller.VisitContainer(&r)
				if childErr != nil {
					return childErr
				} else if childSel == nil {
					continue
				}

				if err = childSel.Walk(context, controller); err != nil {
					return
				}
				if err = childSel.Fire(LEAVE.New()); err != nil {
					return err
				}
			}
		}
	}
	return
}
