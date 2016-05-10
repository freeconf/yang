package node

import (
	"bufio"
	"bytes"
	"github.com/c2g/c2"
	"github.com/c2g/meta"
	"io"
	"runtime/debug"
	"strconv"
	"strings"
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

type NotifyChannel struct {
	channel      string
	Notification *meta.Notification
	send         chan<- *NotifyMessage
	closer       func()
}

func (self *NotifyChannel) Send(path *Path, n Node) {
	var buf bytes.Buffer
	json := NewJsonWriter(&buf).Node()
	c := NewContext()
	err := c.Select(self.Notification, n).InsertInto(json).LastErr
	if err != nil {
		panic(err.Error())
	}
	self.send <- &NotifyMessage{
		Channel: self.channel,
		Op:      'M',
		Path:    path.StringNoModule(),
		Payload: buf.Bytes(),
	}
}

func (self *NotifyChannel) RegisterCloser(closer func()) {
	self.closer = closer
}

func (self *NotifyChannel) Close() {
	if self.closer != nil {
		self.closer()
	}
}

type NotifyChannelFactory interface {
	NewChannel(channel *NotifyChannel, url string)
}

type NotifyMultiplexor struct {
	Channels     map[string]*NotifyChannel
	Factory      NotifyChannelFactory
	Send         chan *NotifyMessage
	Recv         io.ReadCloser
	Timeout      int
	LastErr      error
	SocketHandle io.Closer
	stopped      bool
}

func (self *NotifyMultiplexor) Start(r io.ReadCloser, w io.Writer) {
	self.Channels = make(map[string]*NotifyChannel)
	go self.decode(r)
	self.Recv = r
	self.Send = make(chan *NotifyMessage)
	self.encode(w)
}

func (self *NotifyMultiplexor) Close() error {
	self.stopped = true
	for _, ch := range self.Channels {
		ch.Close()
	}
	close(self.Send)
	return self.Recv.Close()
}

func (self *NotifyMultiplexor) checkErr(n int, err error) {
	if self.LastErr != nil {
		self.LastErr = err
	}
}

func (self *NotifyMultiplexor) encode(stream io.Writer) {
	w := bufio.NewWriter(stream)
	for !self.stopped {
		msg, open := <-self.Send
		if !open {
			return
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
			self.SocketHandle.Close()
			return
		}
	}
}

type NotifyMessage struct {
	Op      rune
	Path    string
	Channel string
	Payload []byte
}

func (self *NotifyMultiplexor) decode(r io.Reader) {
	var msg = make([]byte, 512)
	b := bufio.NewReader(r)
	for {
		n, err := b.Read(msg)
		if err != nil {
			if err == io.EOF {
				break
			}
			self.LastErr = err
			self.SocketHandle.Close()
			return
		}
		line := string(msg[:n])
		decoded := strings.Split(strings.Trim(line, "\n"), " ")
		op := decoded[0][0]
		var channelName string
		if len(decoded) > 1 {
			channelName = decoded[1]
		}
		switch op {
		case '+':
			channel := &NotifyChannel{
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
