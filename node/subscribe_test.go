package node

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/c2stack/c2g/c2"
)

type dummySubFactory struct {
	expectErr error
	t         *testing.T
	count     int
}

func (self *dummySubFactory) Subscribe(c context.Context, sub *Subscription) error {
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
	}
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
	send(`{"op":"+","group":"foo","path":"some/path1"}`)
	if _, hasSub := subs["foo|some/path1"]; !hasSub {
		t.Errorf("Missing subscription: %v", subs)
	}
	if err := c2.CheckEqual(factory.count, 1); err != nil {
		t.Error("Matching subscribes to closes.", err)
	}
	send(`{"op":"+","group":"foo","path":"some/path2"}`)
	send(`{"op":"+","group":"foo","path":"some/path3"}`)
	if err := c2.CheckEqual(len(subs), 3); err != nil {
		t.Errorf("Wrong number of subs %s in  %v", err, subs)
	}
	send(`{"op":"-","group":"foo","path":"some/path1"}`)
	if _, hasSub := subs["foo|some/path1"]; hasSub {
		t.Errorf("Subscription wasn't removed: %v", subs)
	}
	send(`{"op":"-","group":"foo"}`)
	if len(subs) != 0 {
		t.Errorf("Expected no subs, got %v", subs)
	}
	if err := c2.CheckEqual(factory.count, 0); err != nil {
		t.Error("Matching subscribes to closes.", err)
	}
}

func TestSubscribeEncode(t *testing.T) {
	tests := []struct {
		msg      *SubscriptionMessage
		expected string
	}{
		{
			msg: &SubscriptionMessage{
				Group:   "foo",
				Path:    "xpath",
				Type:    "notify",
				Payload: []byte("hello"),
			},
			expected: `{"path":"xpath","type":"notify","group":"foo","payload":"aGVsbG8="}`,
		},
		{
			msg: &SubscriptionMessage{
				Path:    "xpath",
				Type:    "error",
				Payload: []byte("Bad boy"),
			},
			expected: `{"path":"xpath","type":"error","payload":"QmFkIGJveQ=="}`,
		},
	}
	for _, test := range tests {
		var buff bytes.Buffer
		if err := EncodeSubscriptionStream(&buff, test.msg); err != nil {
			t.Error(err)
		}
		if err := c2.CheckEqual(test.expected, strings.TrimRight(buff.String(), "\n")); err != nil {
			t.Error(err)
		}
	}
}
