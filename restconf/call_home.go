package restconf

import (
	"encoding/json"
	"fmt"
	"github.com/c2g/c2"
	"github.com/c2g/meta"
	"github.com/c2g/node"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"github.com/c2g/browse"
)

// Implements RFC Draft in spirit-only
//   https://tools.ietf.org/html/draft-ietf-netconf-call-home-17
//
// Draft calls for server-initiated registration and this implementation is client-initiated
// which may or may-not be part of the final draft.  Client-initiated registration at first
// glance appears to be more useful, but this may prove to be a wrong assumption on my part.
//
type CallHome struct {
	Module            *meta.Module
	ControllerAddress string
	EndpointAddress   string
	EndpointId        string
	Registration      *Registration
	ClientSource      browse.ClientSource
}

type Registration struct {
	Id string
}

func (self *CallHome) Manage() node.Node {
	return &node.Extend{
		Node: node.MarshalContainer(self),
		OnSelect: func(p node.Node, r node.ContainerRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "registration":
				if self.Registration != nil {
					return node.MarshalContainer(self.Registration), nil
				}
			}
			return nil, nil
		},
		OnEvent: func(p node.Node, sel *node.Selection, e node.Event) error {
			switch e.Type {
			case node.LEAVE_EDIT:
				time.AfterFunc(1*time.Second, func() {
					if err := self.Call(); err != nil {
						c2.Err.Print(err)
					}
				})
			}
			return p.Event(sel, e)
		},
	}
}

func (self *CallHome) Call() (err error) {
	var req *http.Request
	c2.Info.Printf("Registering controller %s", self.ControllerAddress)
	if req, err = http.NewRequest("POST", self.ControllerAddress, nil); err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	payload := fmt.Sprintf(`{"module":"%s","id":"%s","endpointAddress":"%s"}`, self.Module.GetIdent(),
		self.EndpointId, self.EndpointAddress)
	req.Body = ioutil.NopCloser(strings.NewReader(payload))
	client := self.ClientSource.GetHttpClient()
	resp, getErr := client.Do(req)
	if getErr != nil {
		return getErr
	}
	defer resp.Body.Close()
	respBytes, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return c2.NewErrC(string(respBytes), resp.StatusCode)
	}
	var rc map[string]interface{}
	if err = json.Unmarshal(respBytes, &rc); err != nil {
		return err
	}
	self.Registration = &Registration{
		Id: rc["id"].(string),
	}
	return nil
}
