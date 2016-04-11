package restconf

import (
	"bufio"
	"bytes"
	"github.com/c2g/meta"
	"github.com/c2g/node"
	"io"
	"strconv"
	"strings"
	"time"
)

type SocketChannel struct {
	channel string
	notif   *meta.Notification
	send    chan<- *SocketMessage
	closer  func()
}

func (self *SocketChannel) Send(n node.Node) {
	var buf bytes.Buffer
	json := node.NewJsonWriter(&buf).Node()
	c := node.NewContext()
	err := c.Select(self.notif, n).InsertInto(json).LastErr
	if err != nil {
		panic(err.Error())
	}
	self.send <- &SocketMessage{
		Channel: self.channel,
		Op:      'M',
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
	Send     chan<- *SocketMessage
	Timeout  int
	WriteErr error
}

func (self *SocketMultiplexor) Start(r io.Reader, w io.Writer) {
	self.Channels = make(map[string]*SocketChannel)
	go self.decode(r)
	c := make(chan *SocketMessage)
	self.Send = c
	go self.encode(w, c)
	self.ping()
}

func (self *SocketMultiplexor) ping() {
	msg := &SocketMessage{
		Op:      '.',
		Channel: "ping",
	}
	for {
		<-time.After(time.Second * time.Duration(self.Timeout / 2))
		self.Send <- msg
	}
}

func (self *SocketMultiplexor) checkErr(n int, err error) {
	if self.WriteErr != nil {
		self.WriteErr = err
	}
}

func (self *SocketMultiplexor) encode(stream io.Writer, c <-chan *SocketMessage) {
	w := bufio.NewWriter(stream)
	for {
		msg := <-c
		if msg == nil {
			break
		}
		if msg.Op == '?' {
			self.checkErr(w.WriteString("?\n"))
		} else {
			self.checkErr(w.WriteRune(msg.Op))
			self.checkErr(w.WriteRune(' '))
			self.checkErr(w.WriteString(msg.Channel))
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
		if self.WriteErr == nil {
			self.WriteErr = w.Flush()
		}
		if self.WriteErr != nil {
			panic(self.WriteErr.Error())
		}
	}
}

type SocketMessage struct {
	Op      rune
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
			panic(err.Error())
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
			self.Factory.NewChannel(channel, url)
			self.Channels[channelName] = channel
		case '-':
			// TODO: unlisten
			if channel, found := self.Channels[channelName]; found {
				channel.Close()
				delete(self.Channels, channelName)
			}
		}
	}
}
