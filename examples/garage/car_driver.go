package garage

import (
	"context"

	"github.com/c2stack/c2g/node"
)

type carDriver struct {
	id   string
	b    *node.Browser
	sub  node.NotifyCloser
	errs chan error
}

func (self *carDriver) Close() {
	self.sub()
}

func (self *carDriver) Id() string {
	return self.id
}

func (self *carDriver) OnChange(l CarChangeListener) error {
	notify := self.b.Root().Find("update")
	if notify.LastErr != nil {
		return notify.LastErr
	}
	var err error
	self.sub, err = notify.Notifications(func(c context.Context, msg node.Selection) {
		var state CarState
		msg.InsertIntoCntx(c, carStateNode(&state))
		l(self, state)
	})
	return err
}

func (self *carDriver) State() (CarState, error) {
	var state CarState
	err := self.b.Root().InsertInto(carStateNode(&state)).LastErr
	return state, err
}

func (self *carDriver) ReplaceTires(state CarState) error {
	return self.b.Root().Find("replaceTires").Action(nil).LastErr
}

func (self *carDriver) RotateTires(state CarState) error {
	return self.b.Root().Find("rotateTires").Action(nil).LastErr
}
