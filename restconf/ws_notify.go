package restconf

import (
	"time"

	"context"

	"github.com/freeconf/yang/c2"
	"golang.org/x/net/websocket"
)

// Determined using default websocket settings and Chrome 49 and stop watch when it
// timesout out w/o any pings from server.
const PingRate = 30 * time.Second

// websocket library will kill connection after this time. This is mostly unnec.
// for our usage because we actively ping so this just has to be larger than ping rate
const serverSocketTimeout = 2 * PingRate

type wsNotifyService struct {
	timeout int
	factory Subscriber
	cancel  context.CancelFunc
	conn    *wsconn
}

func (self *wsNotifyService) Handle(ws *websocket.Conn) {
	eatPanics := false
	defer func() {
		// we want to ignore errors when other-side disappears w/o warning which is
		// outside our control.  We should still panic otherwise because it would
		// point out programtic errors which probably should be surfaced and fixed.
		if eatPanics {
			if r := recover(); r != nil {
				c2.Debug.Printf("recovering from panic from unclear browser disconnect. %s", r)
			}
		}
	}()
	defer ws.Close()
	defer self.cancel()

	var rate time.Duration
	if self.timeout == 0 {
		rate = PingRate
	} else {
		rate = time.Duration(self.timeout) * time.Millisecond
	}
	self.conn = &wsconn{
		pinger: time.NewTicker(rate),
		mgr:    NewSubscriptionManager(self.factory, ws, ws),
	}
	defer self.conn.close()
	ws.Request().Body.Close()
	go self.conn.keepAlive(ws)
	if err := self.conn.mgr.Run(); err != nil {
		eatPanics = true
		c2.Debug.Printf("unclean terminination of web socket: (%s). other side may have close browser. closing socket.", err)
	}
}

type wsconn struct {
	pinger *time.Ticker
	mgr    *SubscriptionManager
}

func (self *wsconn) keepAlive(ws *websocket.Conn) {
	for {
		ws.SetDeadline(time.Now().Add(serverSocketTimeout))
		if fw, err := ws.NewFrameWriter(websocket.PingFrame); err != nil {
			//self.Close()
			return
		} else if _, err = fw.Write([]byte{}); err != nil {
			//self.Close()
			return
		}
		if _, running := <-self.pinger.C; !running {
			return
		}
	}
}

func (self *wsconn) close() {
	self.pinger.Stop()
	self.mgr.Close()
}
