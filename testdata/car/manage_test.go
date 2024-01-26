package car

import (
	"testing"
	"time"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"
)

// Test the car management logic in manage.go
func TestManage(t *testing.T) {

	// setup
	ypath := source.Path(".")
	mod := parser.RequireModule(ypath, "car")
	app := New()
	app.pollInterval = time.Millisecond

	// no web server needed, just your app and management function.
	brwsr := node.NewBrowser(mod, Manage(app))
	root := brwsr.Root()

	// read all config
	currCfg, err := nodeutil.WriteJSON(sel(root.Find("?content=config")))
	fc.AssertEqual(t, nil, err)
	expected := `{"speed":0,"oilLevel":10,"tire":[{"pos":0,"size":"H15"},{"pos":1,"size":"H15"},{"pos":2,"size":"H15"},{"pos":3,"size":"H15"}]}`
	fc.AssertEqual(t, expected, currCfg)

	// access car and verify w/API
	fc.AssertEqual(t, false, app.Running)

	// setup event listener, verify events later
	events := make(chan string)
	unsub, err := sel(root.Find("update")).Notifications(func(n node.Notification) {
		event, _ := nodeutil.WriteJSON(n.Event)
		events <- event
	})
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, 1, app.listeners.Len())

	// write config starts car
	err = root.UpdateFrom(nodeutil.ReadJSON(`{"speed":1000}`))
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, 1000, app.Speed)

	// start car
	fc.AssertEqual(t, nil, justErr(sel(root.Find("start")).Action(nil)))

	// should be first event
	fc.AssertEqual(t, `{"event":"carStarted"}`, <-events)
	fc.AssertEqual(t, true, app.Running)

	// unsubscribe
	unsub()
	fc.AssertEqual(t, 0, app.listeners.Len())

	// hit all the RPCs
	fc.AssertEqual(t, nil, justErr(sel(root.Find("rotateTires")).Action(nil)))
	fc.AssertEqual(t, nil, justErr(sel(root.Find("replaceTires")).Action(nil)))
	fc.AssertEqual(t, nil, justErr(sel(root.Find("reset")).Action(nil)))
	fc.AssertEqual(t, nil, justErr(sel(root.Find("tire=0/replace")).Action(nil)))
	fc.AssertEqual(t, nil, justErr(sel(root.Find("stop")).Action(nil)))
	fc.AssertEqual(t, false, app.Running)
	fc.AssertEqual(t, nil, justErr(sel(root.Find("stop")).Action(nil)))
	out, err := sel(root.Find("addOil")).Action(nodeutil.ReadJSON(`{"drainFirst": true, "amount": 10.4}`))
	fc.AssertEqual(t, nil, err)
	expectedOil, err := nodeutil.WriteJSON(out)
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, `{"oilLevel":10.4}`, expectedOil)

	fc.AssertEqual(t, nil, justErr(sel(root.Find("start")).Action(nil)))
	fc.AssertEqual(t, true, app.Running)
}

func sel(s *node.Selection, err error) *node.Selection {
	if err != nil {
		panic(err)
	}
	return s
}

func justErr(_ *node.Selection, err error) error {
	return err
}
