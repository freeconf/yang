package restconf

import (
	"github.com/c2stack/c2g/node"
	"golang.org/x/net/websocket"
	"time"
	"github.com/c2stack/c2g/c2"
)

// Determined using default websocket settings and Chrome 49 and stop watch when it
// timesout out w/o any pings from server.
const PingRate = 30 * time.Second

// websocket library will kill connection after this time. This is mostly unnec.
// for our usage because we actively ping so this just has to be larger than ping rate
const serverSocketTimeout = 2 * PingRate

type WebSocketService struct {
	Timeout int
	Factory node.Subscriber
}

func (self *WebSocketService) Handle(ws *websocket.Conn) {
	var rate time.Duration
	if self.Timeout == 0 {
		rate = PingRate
	} else {
		rate = time.Duration(self.Timeout) * time.Millisecond
	}
	conn := &wsconn{
		pinger: time.NewTicker(rate),
		mgr:    node.NewSubscriptionManager(self.Factory, ws, ws),
	}
	defer conn.close()
	ws.Request().Body.Close()
	go conn.keepAlive(ws)
	if err := conn.mgr.Run(); err != nil {
		c2.Err.Print("Error handling subscription. ", err)
	}
	if err := ws.Close(); err != nil {
		c2.Err.Print("Error closing socket. ", err)
	}
}

type wsconn struct {
	pinger *time.Ticker
	mgr    *node.SubscriptionManager
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
		if _, running := <- self.pinger.C; !running {
			return
		}
	}
}

func (self *wsconn) close() {
	self.pinger.Stop()
	self.mgr.Close()
}
