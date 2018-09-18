package restconf

import (
	"strings"
	"testing"
	"time"

	"github.com/freeconf/gconf/device"
	"github.com/freeconf/gconf/meta"
	"github.com/freeconf/gconf/meta/yang"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/nodes"
)

func TestClientReconnect(t *testing.T) {
	t.Skip("Fails until we figure out how to get WS connections to autoconnect")
	ypath := meta.PathStreamSource("./testdata:../yang")
	m := yang.RequireModule(ypath, "x")
	var msgs chan string
	var s *Server
	connect := func() {
		msgs = make(chan string)
		n := &nodes.Basic{
			OnNotify: func(r node.NotifyRequest) (node.NotifyCloser, error) {
				go func() {
					for {
						select {
						case s := <-msgs:
							r.Send(nodes.ReflectChild(map[string]interface{}{
								"z": s,
							}))
						default:
							return
						}
					}
				}()
				return func() error {
					return nil
				}, nil
			},
		}
		b := node.NewBrowser(m, n)
		d := device.New(ypath)
		d.AddBrowser(b)
		s = NewServer(d)
		err := d.ApplyStartupConfig(strings.NewReader(`
		{
			"restconf" : {
				"web": {
					"port" : ":9080"
				},
				"debug" : true
			}
		}`))
		if err != nil {
			t.Fatal(err)
		}
	}
	connect()
	<-time.After(2 * time.Second)
	factory := Client{YangPath: ypath}
	c, err := factory.NewDevice("http://localhost:9080/restconf")
	if err != nil {
		t.Fatal(err)
	}
	b, err := c.Browser("x")
	if err != nil {
		t.Fatal(err)
	}
	_, err = b.Root().Find("y").Notifications(func(sel node.Selection) {
		t.Log("got here")
	})
	if err != nil {
		t.Fatal(err)
	}
	msgs <- "original session "
	s.Close()
	connect()
	<-time.After(2 * time.Second)
	msgs <- "new session"
}
