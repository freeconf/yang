package car

import (
	"strings"
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
)

func TestManagement(t *testing.T) {
	yangPath := &meta.FileStreamSource{Root: "."}

	// parse YANG
	m := yang.RequireModule(yangPath, "car")

	// attach YANG to management code to create management browser object
	// use browser object to test interface.  Notice that http server is unnec.
	b := node.NewBrowser(m, Manage(New()))

	// Configuration
	if err := b.Root().Set("speed", 1); err != nil {
		t.Error(err)
	}

	// Operations
	if err := b.Root().Find("replaceTires").Action(nil).LastErr; err != nil {
		t.Error(err)
	}

	// Alerts
	wait := make(chan struct{})
	_, err := b.Root().Find("update").Notifications(func(update node.Selection) {
		tires, err := nodes.WriteJSON(update.Find("tire?fields=worn%3Bflat"))
		if err != nil {
			t.Error(err)
		}

		// most likely worn, but could also be flat
		if strings.Contains(tires, "true") {
			wait <- struct{}{}
		}
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for wear/flat...")
	<-wait

	// Metrics
	if v, err := b.Root().GetValue("running"); err != nil {
		t.Error(err)
	} else {
		if v.Value().(bool) != true {
			t.Error("not running")
		}
	}
}
