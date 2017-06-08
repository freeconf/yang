package car

import (
	"strings"
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

func Test_App(t *testing.T) {
	yangPath := &meta.FileStreamSource{Root: "."}
	app := New()

	m, err := yang.LoadModule(yangPath, "car")
	b := node.NewBrowser(m, Node(app))
	sel := b.Root()
	sel.Set("speed", 10)
	if err := sel.Find("replaceTires").Action(nil).LastErr; err != nil {
		t.Error(err)
	}
	wait := make(chan struct{})
	_, err = sel.Find("update").Notifications(func(update node.Selection) {
		tires, err := node.WriteJson(update.Find("tire?fields=worn%3Bflat"))
		t.Log("update", tires)
		if err != nil {
			t.Error(err)
		}
		t.Log(tires)

		// most likely worn, but could also be flat
		if strings.Contains(tires, "true") {
			wait <- struct{}{}
		}
	})
	t.Log("waiting for wear/flat...")
	<-wait
	if err != nil {
		t.Error(err)
	}
}
