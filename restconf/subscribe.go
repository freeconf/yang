package restconf

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/freeconf/c2g/c2"
	"github.com/freeconf/c2g/meta"
	"github.com/freeconf/c2g/node"
	"github.com/freeconf/c2g/nodes"
)

type Subscriber interface {
	Subscribe(sub *Subscription) error
}

type SubscriptionManager struct {
	subscriptions map[string]*Subscription
	factory       Subscriber
	Send          chan *SubscriptionOutgoing
	rdr           io.Reader
	wtr           io.Writer
	decoder       func(r io.Reader, conn *SubscriptionManager) error
	encoder       func(w io.Writer, msg *SubscriptionOutgoing) error
	lastErr       error
}

func NewSubscriptionManager(factory Subscriber, r io.Reader, w io.Writer) *SubscriptionManager {
	return &SubscriptionManager{
		rdr:           r,
		wtr:           w,
		factory:       factory,
		subscriptions: make(map[string]*Subscription),
		Send:          make(chan *SubscriptionOutgoing),
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
		sub.Close()
	}
	self.subscriptions = nil
	close(self.Send)
	return self.lastErr
}

// Subscribe to notifcation event
// SEND
// =============
// {op:+, id:foo, path:birdSpotted, group:birding, ...}
//
// RECEIVE
// =============
// message with actual data in json format is base64 encoded into payload
// field
//
//   {id:foo,payload:Aq235Tg12B2}
// or
//   {id:foo,type:error,payload:GWe234haD}
//
// Unsubscribe to notification event
// SEND
// =============
// {op:-, id:foo}
//
// *No response expected*
//
// Unsubscribe to all notification events for a group
// SEND
// =============
// {op:-, group:bar}
//
// *No response expected*

type Subscription struct {
	Path         string
	DeviceId     string
	Module       string
	Notification *meta.Notification
	Closer       node.NotifyCloser
	Context      context.Context
	Id           string
	Group        string
	send         chan<- *SubscriptionOutgoing
}

func (self *Subscription) Notify(message node.Selection) {
	defer func() {
		// we want to ignore errors when other-side disappears w/o warning which is
		// outside our control.  We should still panic otherwise because it would
		// point out programtic errors which probably should be surfaced and fixed.
		if r := recover(); r != nil {
			c2.Debug.Printf("recovering from panic disconnecting. %s", r)
		}
	}()
	var payload []byte
	if message.Node != nil {
		buf, err := nodes.WriteJSON(message)
		if err != nil {
			panic(err.Error())
		}
		if c2.DebugLogEnabled() {
			c2.Debug.Printf("NOTIFY %s %s", self.Path, buf)
		}
		payload = []byte(buf)
	}
	self.send <- &SubscriptionOutgoing{
		Id:      self.Id,
		Type:    "notify",
		Payload: payload,
	}
}

func (self *Subscription) Close() error {
	c2.Debug.Printf("notify close %s", self.Path)
	if self.Closer != nil {
		return self.Closer()
	}
	return nil
}

func EncodeSubscriptionStream(w io.Writer, msg *SubscriptionOutgoing) (err error) {
	return json.NewEncoder(w).Encode(msg)
}

type SubscriptionOutgoing struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Payload []byte `json:"payload,omitempty"`
}

type SubscriptionIncoming struct {
	Op     string `json:"op"`
	Id     string `json:"id"`
	Path   string `json:"path"`
	Module string `json:"module"`
	Group  string `json:"group,omitempty"`
	Device string `json:"device,omitempty"`
}

// This assumes each message will
//{
//  "op" : "+|-",
//  "id" : "x",
//  ...
//}
func DecodeSubscriptionStream(r io.Reader, conn *SubscriptionManager) error {
	jsonDecoder := json.NewDecoder(r)
	for {
		// This only works if each message fits in one read buffer and a read
		// buffer contains one message.  This might only be true for web sockets
		// and i'm not 100% positive it's true for web sockets spec, but it does
		// appear to be the standard practice I've observed
		var msg SubscriptionIncoming
		if err := jsonDecoder.Decode(&msg); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if msg.Op == "" {
			conn.Send <- &SubscriptionOutgoing{
				Id:      msg.Id,
				Type:    "error",
				Payload: []byte("id and op are required properties."),
			}
			continue
		}
		recoverableBlock := func() error {
			defer func() {
				if r := recover(); r != nil {
					err, isErr := r.(error)
					if !isErr {
						err = c2.NewErr(fmt.Sprintf("%v", r))
					}
					conn.Send <- &SubscriptionOutgoing{
						Id:      msg.Id,
						Type:    "error",
						Payload: []byte(err.Error()),
					}
				}
			}()
			switch msg.Op {
			case "+":
				if msg.Path == "" || msg.Module == "" || msg.Id == "" {
					return c2.NewErr("path, id and module are all required in subscription")
				}
				return conn.newSubscription(msg)
			case "-":
				// TODO: unlisten
				if msg.Id != "" {
					conn.removeSubscription(msg.Id)
				} else {
					conn.removeGroup(msg.Group)
				}
			default:
				return c2.NewErr("Unrecognized notify operation: " + msg.Op)
			}
			return nil
		}
		if err := recoverableBlock(); err != nil {
			conn.Send <- &SubscriptionOutgoing{
				Id:      msg.Id,
				Type:    "error",
				Payload: []byte(err.Error()),
			}
		}
	}
}

func (self *SubscriptionManager) newSubscription(msg SubscriptionIncoming) error {
	// just incase there was an old one
	self.removeSubscription(msg.Id)

	sub := &Subscription{
		Id:       msg.Id,
		Group:    msg.Group,
		DeviceId: msg.Device,
		Module:   msg.Module,
		Path:     msg.Path,
		send:     self.Send,
	}
	c2.Debug.Printf("notify open %s", msg.Path)
	if err := self.factory.Subscribe(sub); err != nil {
		return err
	}

	self.subscriptions[msg.Id] = sub
	return nil
}

func (self *SubscriptionManager) Len() int {
	return len(self.subscriptions)
}

func (self *SubscriptionManager) removeSubscription(id string) {
	if sub, found := self.subscriptions[id]; found {
		sub.Close()
		delete(self.subscriptions, id)
	}
}

func (self *SubscriptionManager) removeGroup(group string) {
	for id, sub := range self.subscriptions {
		if sub.Group == group {
			self.removeSubscription(id)
		}
	}
}

func (self *SubscriptionManager) removeAllSubscriptions() {
	for id, _ := range self.subscriptions {
		self.removeSubscription(id)
	}
}
