package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"io"
)

type Subscriber interface {
	Subscribe(sub *Subscription) error
}

type SubscriptionManager struct {
	subscriptions map[string]*Subscription
	factory       Subscriber
	Send          chan *SubscriptionMessage
	rdr           io.Reader
	wtr           io.Writer
	decoder       func(r io.Reader, conn *SubscriptionManager) error
	encoder       func(w io.Writer, msg *SubscriptionMessage) error
	lastErr       error
}

func NewSubscriptionManager(factory Subscriber, r io.Reader, w io.Writer) *SubscriptionManager {
	return &SubscriptionManager{
		rdr:           r,
		wtr:           w,
		factory:       factory,
		subscriptions: make(map[string]*Subscription),
		Send:          make(chan *SubscriptionMessage),
		decoder:       DecodeSubscriptionStream,
		encoder:       EncodeSubscriptionStream,
	}
}

func (self *SubscriptionManager) Run() error {
	go func() {
		self.checkErr(self.decoder(self.rdr, self))
		//self.Close()
	}()
	for {
		msg, open := <-self.Send
		if !open {
			break
		}
		if err := self.encoder(self.wtr, msg); err != nil {
			// This is fatal error sending, not err in application layer so
			// we should abort all further communication
			return err
		}
	}
	return self.lastErr
}

func (self *SubscriptionManager) checkErr(err error) {
	if err != nil && self.lastErr == nil {
		self.lastErr = err
	}
}

func (self *SubscriptionManager) Close() error {
	self.lastErr = nil
	for _, sub := range self.subscriptions {
		self.checkErr(sub.Close())
	}
	self.subscriptions = nil
	close(self.Send)
	return self.lastErr
}

// Subscribe to notifcation event
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
//
// Unsubscribe to notification event
// SEND
// =============
// - foo job/states
//
// *No response expected*
//
// Unsubscribe to all notification events for a context
// SEND
// =============
// - foo
//
// *No response expected*

type Subscription struct {
	Path         string
	Notification *meta.Notification
	Closer       NotifyCloser
	id           string
	group        string
	send         chan<- *SubscriptionMessage
}

func (self *Subscription) Notify(notification *meta.Notification, path *Path, n Node) {
	var payload []byte
	if n != nil {
		var buf bytes.Buffer
		json := NewJsonWriter(&buf).Node()
		sel := NewBrowser2(self.Notification, n).Root()
		err := sel.InsertInto(json).LastErr
		if err != nil {
			panic(err.Error())
		}
		payload = buf.Bytes()
	}
	self.send <- &SubscriptionMessage{
		Group:   self.group,
		Path:    path.StringNoModule(),
		Type:    "notify",
		Payload: payload,
	}
}

func (self *Subscription) Close() error {
	if self.Closer != nil {
		return self.Closer()
	}
	return nil
}

func EncodeSubscriptionStream(w io.Writer, msg *SubscriptionMessage) (err error) {
	return json.NewEncoder(w).Encode(msg)
}

type SubscriptionMessage struct {
	Path    string `json:"path"`
	Type    string `json:"type"`
	Group   string `json:"group,omitempty"`
	Payload []byte `json:"payload,omitempty"`
}

// This assumes each message will
//{
//  "op" : "+|-",
//  "context" : "x",
//  "path" : "y"
//}
func DecodeSubscriptionStream(r io.Reader, conn *SubscriptionManager) error {
	jsonDecoder := json.NewDecoder(r)
	msg := make(map[string]string)
	for {
		for k, _ := range msg {
			delete(msg, k)
		}

		// This only works if each message fits in one read buffer and a read
		// buffer contains one message.  This might only be true for web sockets
		// and i'm not 100% positive it's true for web sockets spec, but it does
		// appear to be the standard practice I've observed
		if err := jsonDecoder.Decode(&msg); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		op, hasOp := msg["op"]
		if !hasOp {
			conn.Send <- &SubscriptionMessage{
				Type:    "error",
				Payload: []byte("Missing operation value"),
			}
			continue
		}
		group := msg["group"]
		path, hasPath := msg["path"]
		recoverableBlock := func() error {
			defer func() {
				if r := recover(); r != nil {
					err, isErr := r.(error)
					if !isErr {
						err = c2.NewErr(fmt.Sprintf("%v", r))
					}
					conn.Send <- &SubscriptionMessage{
						Path:    path,
						Group:   group,
						Type:    "error",
						Payload: []byte(err.Error()),
					}
				}
			}()
			switch op {
			case "+":
				if !hasPath {
					return c2.NewErr("Missing path value in subscription")
				}
				return conn.newSubscription(group, path)
			case "-":
				// TODO: unlisten
				if hasPath {
					id := fmt.Sprint(group + "|" + path)
					conn.removeSubscription(id)
				} else {
					conn.removeGroup(group)
				}
			default:
				return c2.NewErr("Unrecognized notify operation: " + op)
			}
			return nil
		}
		if err := recoverableBlock(); err != nil {
			conn.Send <- &SubscriptionMessage{
				Group:   group,
				Path:    path,
				Type:    "error",
				Payload: []byte(err.Error()),
			}
		}
	}
}

func (self *SubscriptionManager) newSubscription(group string, path string) error {
	id := fmt.Sprint(group + "|" + path)
	sub := &Subscription{
		id:    id,
		group: group,
		Path:  path,
		send:  self.Send,
	}
	if err := self.factory.Subscribe(sub); err != nil {
		return err
	}

	self.subscriptions[id] = sub
	return nil
}

func (self *SubscriptionManager) removeSubscription(id string) {
	if sub, found := self.subscriptions[id]; found {
		sub.Close()
		delete(self.subscriptions, id)
	}
}

func (self *SubscriptionManager) removeGroup(group string) {
	for id, sub := range self.subscriptions {
		if sub.group == group {
			self.removeSubscription(id)
		}
	}
}

func (self *SubscriptionManager) removeAllSubscriptions() {
	for id, _ := range self.subscriptions {
		self.removeSubscription(id)
	}
}
