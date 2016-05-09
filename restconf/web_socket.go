package restconf

import (
	"golang.org/x/net/websocket"
	"time"
	"github.com/c2g/node"
)

// Determined using default websocket settingsa and Chrome 49 and stop watch when it
// timesout out w/o any pings from server.
const PingRate = 30 * time.Second

// websocket library will kill connection after this time. This is mostly unnec.
// for our usage because we actively ping so this just has to be larger than ping rate
const serverSocketTimeout = 2 * PingRate

type WebSocket struct {
	Timeout int
	stop    bool
	mux     *node.NotifyMultiplexor
}

func (self *WebSocket) ConnectionHandler(ws *websocket.Conn, factory node.NotifyChannelFactory) {
	defer ws.Close()
	self.mux = &node.NotifyMultiplexor{
		Factory: factory,
		Timeout: self.Timeout,
		SocketHandle: self,
	}
	ws.Request().Body.Close()
	go self.keepAlive(ws)
	self.mux.Start(ws, ws)
}

func (self *WebSocket) keepAlive(ws *websocket.Conn) {
	for ! self.stop {
		ws.SetDeadline(time.Now().Add(serverSocketTimeout))
		<- time.After(PingRate)
		if fw, err := ws.NewFrameWriter(websocket.PingFrame); err != nil {
			self.Close()
		} else if _, err = fw.Write([]byte{}); err != nil {
			self.Close()
		}
	}
}

func (self *WebSocket) Close() error {
	self.stop = true
	if self.mux != nil {
		m := self.mux
		self.mux = nil
		return m.Close()
	}
	return nil
}
