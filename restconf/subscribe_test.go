package restconf

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/freeconf/yang/c2"
)

type dummySubFactory struct {
	expectErr error
	t         *testing.T
	count     int
}

func (self *dummySubFactory) Subscribe(sub *Subscription) error {
	sub.Closer = func() error {
		self.count--
		return nil
	}
	self.count++
	return nil
}

func TestSubscribeDecode(t *testing.T) {
	subs := make(map[string]*Subscription)
	factory := &dummySubFactory{}
	c := &SubscriptionManager{
		factory:       factory,
		subscriptions: subs,
		Send:          make(chan *SubscriptionOutgoing),
	}
	go func() {
		// drain outgoing msg bus
		for {
			msg := <-c.Send
			t.Errorf("unexpected message %v", msg)
		}

	}()
	send := func(msg string) {
		r, w := io.Pipe()
		go func() {
			w.Write([]byte(msg))
			w.Close()
		}()
		if err := DecodeSubscriptionStream(r, c); err != nil {
			t.Error(err)
		}
	}
	send(`{"op":"+","id":"1","group":"foo","module":"x","path":"some/path1"}`)
	if _, hasSub := subs["1"]; !hasSub {
		t.Errorf("Missing subscription: %v", subs)
	}
	c2.AssertEqual(t, factory.count, 1)
	send(`{"op":"+","id":"2","group":"foo","module":"x","path":"some/path2"}`)
	send(`{"op":"+","id":"3","group":"foo","module":"x","path":"some/path3"}`)
	c2.AssertEqual(t, len(subs), 3)
	send(`{"op":"-","id":"1"}`)
	if _, hasSub := subs["1"]; hasSub {
		t.Errorf("Subscription wasn't removed: %v", subs)
	}
	send(`{"op":"-","group":"foo"}`)
	if len(subs) != 0 {
		t.Errorf("Expected no subs, got %v", subs)
	}
	c2.AssertEqual(t, factory.count, 0)
}

func TestSubscribeEncode(t *testing.T) {
	tests := []struct {
		msg      *SubscriptionOutgoing
		expected string
	}{
		{
			msg: &SubscriptionOutgoing{
				Id:      "x",
				Type:    "notify",
				Payload: []byte("hello"),
			},
			expected: `{"id":"x","type":"notify","payload":"aGVsbG8="}`,
		},
		{
			msg: &SubscriptionOutgoing{
				Id:      "x",
				Type:    "error",
				Payload: []byte("Bad boy"),
			},
			expected: `{"id":"x","type":"error","payload":"QmFkIGJveQ=="}`,
		},
	}
	for _, test := range tests {
		var buff bytes.Buffer
		if err := EncodeSubscriptionStream(&buff, test.msg); err != nil {
			t.Error(err)
		}
		c2.AssertEqual(t, test.expected, strings.TrimRight(buff.String(), "\n"))
	}
}
