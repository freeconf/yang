package restconf

import (
	"bufio"
	"bytes"
	"github.com/c2g/c2"
	"github.com/c2g/meta"
	"github.com/c2g/node"
	"io"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// SEND
// =============
// + foo job/states
//
// RECEIVE
// =============
// M foo some/path 51
// {"id":"5723adaa3684664e2b1936e9","state":"running"}
// M foo some/path 53
// {"id":"5723adaa3684664e2b1936e9","state":"succeeded"}
// . ping 0

type SocketChannel struct {
	channel      string
	Notification *meta.Notification
	send         chan<- *SocketMessage
	closer       func()
}

func (self *SocketChannel) Send(path *node.Path, n node.Node) {
	var buf bytes.Buffer
	json := node.NewJsonWriter(&buf).Node()
	c := node.NewContext()
	err := c.Select(self.Notification, n).InsertInto(json).LastErr
	if err != nil {
		panic(err.Error())
	}
	self.send <- &SocketMessage{
		Channel: self.channel,
		Op:      'M',
		Path:    path.StringNoModule(),
		Payload: buf.Bytes(),
	}
}

func (self *SocketChannel) RegisterCloser(closer func()) {
	self.closer = closer
}

func (self *SocketChannel) Close() {
	if self.closer != nil {
		self.closer()
	}
}

type ChannelFactory interface {
	NewChannel(channel *SocketChannel, url string)
}

type SocketMultiplexor struct {
	Channels map[string]*SocketChannel
	Factory  ChannelFactory
	Send     chan *SocketMessage
	Recv     io.ReadCloser
	Timeout  int
	LastErr  error
	stopped bool
}

func (self *SocketMultiplexor) Start(r io.ReadCloser, w io.Writer) {
	self.Channels = make(map[string]*SocketChannel)
	go self.decode(r)
	self.Recv = r
	self.Send = make(chan *SocketMessage)
	self.encode(w)
	//self.ping()
}

func (self *SocketMultiplexor) ping() {
	msg := &SocketMessage{
		Op:      '.',
		Channel: "ping",
	}
	for !self.stopped {
		<-time.After(time.Second * time.Duration(self.Timeout/2))
		if !self.stopped {
c2.Debug.Printf("sending ping")
			self.Send <- msg
		}
	}
}

func (self *SocketMultiplexor) close() {
	if self.LastErr != nil {
		c2.Err.Printf(self.LastErr.Error())
	}
c2.Debug.Printf("closing socket")
	self.stopped = true
	for _, ch := range self.Channels {
		ch.Close()
	}
	close(self.Send)
	self.Recv.Close()
}

func (self *SocketMultiplexor) checkErr(n int, err error) {
	if self.LastErr != nil {
		self.LastErr = err
	}
}

func (self *SocketMultiplexor) encode(stream io.Writer) {
	w := bufio.NewWriter(stream)
	for {
		msg := <-self.Send
		if msg == nil {
			break
		}
		if msg.Op == '.' {
			self.checkErr(w.WriteString(".\n"))
		} else {
			self.checkErr(w.WriteRune(msg.Op))
			self.checkErr(w.WriteRune(' '))
			self.checkErr(w.WriteString(msg.Channel))
			self.checkErr(w.WriteRune(' '))
			self.checkErr(w.WriteString(msg.Path))
			if msg.Payload != nil {
				l := len(msg.Payload)
				self.checkErr(w.WriteRune(' '))
				self.checkErr(w.WriteString(strconv.Itoa(l)))
				self.checkErr(w.WriteRune('\n'))
				self.checkErr(w.Write(msg.Payload))
			} else {
				self.checkErr(w.WriteString(" 0\n"))
			}
		}
		if self.LastErr == nil {
			self.LastErr = w.Flush()
		}
		if self.LastErr != nil {
			self.close()
			return
		}
	}
}

type SocketMessage struct {
	Op      rune
	Path    string
	Channel string
	Payload []byte
}

func (self *SocketMultiplexor) decode(r io.Reader) {
	b := bufio.NewReader(r)
	for {
		line, err := b.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			self.LastErr = err
			self.close()
			return
		}
		decoded := strings.Split(strings.Trim(line, "\n"), " ")
		op := decoded[0][0]
		var channelName string
		if len(decoded) > 1 {
			channelName = decoded[1]
		}
		switch op {
		case '?':
			// ping
		case '+':
			channel := &SocketChannel{
				channel: channelName,
				send:    self.Send,
			}
			url := decoded[2]
			func() {
				defer func() {
					if err := recover(); err != nil {
						c2.Err.Printf("%s\n%s", err, string(debug.Stack()))
						// TODO: send err back to client
					}
				}()
				self.Factory.NewChannel(channel, url)
				self.Channels[channelName] = channel
			}()
		case '-':
			// TODO: unlisten
			if channel, found := self.Channels[channelName]; found {
				channel.Close()
				delete(self.Channels, channelName)
			}
		}
	}
}
