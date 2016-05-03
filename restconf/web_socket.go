package restconf

import (
	"golang.org/x/net/websocket"
	"time"
)

type WebSocket struct {
	Timeout int
}

func (self *WebSocket) ConnectionHandler(ws *websocket.Conn, factory ChannelFactory) {
	defer ws.Close()
	ws.SetDeadline(time.Now().Add(5 * time.Minute))
	multi := &SocketMultiplexor{
		Factory: factory,
		Timeout: self.Timeout,
	}
	ws.Request().Body.Close()
	multi.Start(ws, ws)
}
