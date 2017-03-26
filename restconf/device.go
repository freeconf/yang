package restconf

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"context"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"golang.org/x/net/websocket"
)

func NewDeviceByHostAndPort(yangPath meta.StreamSource, host string, port string) (conf.Device, error) {
	return NewDevice(yangPath, fmt.Sprintf("https://%s:%s/restconf/", host, port))
}

func NewInsecureDeviceByHostAndPort(yangPath meta.StreamSource, host string, port string) (conf.Device, error) {
	return NewDevice(yangPath, fmt.Sprintf("http://%s:%s/restconf/", host, port))
}

func NewDevice(yangPath meta.StreamSource, address string) (conf.Device, error) {
	remoteSchemaPath := httpStream{
		client: http.DefaultClient,
		url:    address + "schema/",
	}
	return &Device{
		address:       address,
		yangPath:      yangPath,
		schemaPath:    meta.MultipleSources(yangPath, remoteSchemaPath),
		client:        http.DefaultClient,
		subscriptions: make(map[string]*driverSub),
		modules:       make(map[string]*meta.Module),
	}, nil
}

var badAddressErr = c2.NewErr("Expected format: http://server/mount/module:path")

// SplitAddress takes a complete address and breaks it into pieces according
// to RESTCONF standards so you can use each piece in appropriate API call
// Example:
//   http://server:port/restconf/module:path/some=x/where
//
func SplitAddress(url string) (address string, module string, path string, err error) {
	eoSlashSlash := strings.Index(url, "//") + 2
	if eoSlashSlash < 2 {
		err = badAddressErr
		return
	}
	eoSlash := eoSlashSlash + strings.IndexRune(url[eoSlashSlash:], '/') + 1
	if eoSlash <= eoSlashSlash {
		err = badAddressErr
		return
	}
	colon := eoSlash + strings.IndexRune(url[eoSlash:], ':')
	if colon <= eoSlash {
		err = badAddressErr
		return
	}
	moduleBegin := strings.LastIndex(url[:colon], "/")
	address = url[:moduleBegin+1]
	module = url[moduleBegin+1 : colon]
	path = url[colon+1:]
	return
}

type Device struct {
	address       string
	yangPath      meta.StreamSource
	schemaPath    meta.StreamSource
	client        *http.Client
	origin        string
	_ws           *websocket.Conn
	subscriptions map[string]*driverSub
	modules       map[string]*meta.Module
}

func (self *Device) SchemaSource() meta.StreamSource {
	return self.schemaPath
}

func (self *Device) UiSource() meta.StreamSource {
	return httpStream{
		client: http.DefaultClient,
		url:    self.address + "ui/",
	}
}

func (self *Device) Browser(module string) (*node.Browser, error) {
	d := &driver{support: self}
	m, err := self.module(module)
	if err != nil {
		return nil, err
	}
	return node.NewBrowser(m, d.node()), nil
}

func (self *Device) Close() {
	if self._ws != nil {
		self._ws.Close()
		self._ws = nil
	}
}

func (self *Device) driverWebsocket() (io.Writer, error) {
	// lazy start websocket connection to be more efficient if it's never used
	// but I have no data how how much resources this saves
	if self._ws == nil {
		wsUrl := self.address + "/stream/"
		origin := self.origin
		if origin == "" {
			urlParts, err := url.Parse(wsUrl)
			if err != nil {
				return nil, err
			}
			origin = urlParts.Host
		}
		var err error
		if self._ws, err = websocket.Dial(wsUrl, "", origin); err != nil {
			return nil, err
		}
		self.watch(self._ws)
	}
	return self._ws, nil
}

func (self *Device) watch(ws io.Reader) {
	for {
		var notification map[string]interface{}
		if err := json.NewDecoder(ws).Decode(&notification); err != nil {
			c2.Err.Print(err)
			continue
		}
		var payload string
		if payloadData, exists := notification["payload"]; !exists {
			c2.Err.Print("No payload found")
			continue
		} else {
			if payloadDecoded, err := base64.StdEncoding.DecodeString(payloadData.(string)); err != nil {
				c2.Err.Print(err)
				continue
			} else {
				payload = string(payloadDecoded)
			}
		}
		if notification["type"] == "error" {
			c2.Err.Print(payload)
			continue
		}
		pathData := notification["path"]
		if l := self.subscriptions[pathData.(string)]; l != nil {
			n := node.NewJsonReader(strings.NewReader(payload)).Node()
			l.notify(context.Background(), l.sel.Split(n))
		}
	}
}

func (self *Device) driverSubs() map[string]*driverSub {
	return self.subscriptions
}

func (self *Device) module(module string) (*meta.Module, error) {
	m := self.modules[module]
	if m == nil {
		var err error
		if m, err = yang.LoadModule(self.schemaPath, module); err != nil {
			return nil, err
		}
		self.modules[module] = m
	}
	return m, nil
}

// ClientSchema downloads schema and implements yang.StreamSource so it can transparently
// be used in a YangPath.
type httpStream struct {
	client *http.Client
	url    string
}

// OpenStream implements meta.StreamSource
func (self httpStream) OpenStream(name string, ext string) (meta.DataStream, error) {
	fullUrl := self.url + name + ext
	resp, err := self.client.Get(fullUrl)
	if resp != nil {
		return resp.Body, err
	}
	return nil, err
}

func getModule(p *node.Path) *meta.Module {
	seg := p
	for seg != nil {
		if mod, isModule := seg.Meta().(*meta.Module); isModule {
			return mod
		}
		seg = seg.Parent()
	}
	panic("No module found, illegal path")
}

func (self *Device) driverDo(method string, params string, p *node.Path, payload io.Reader) (node.Node, error) {
	var req *http.Request
	var err error
	mod := getModule(p)
	fullUrl := self.address + "data/" + mod.GetIdent() + ":" + p.StringNoModule()
	if params != "" {
		fullUrl += "&" + params
	}
	if req, err = http.NewRequest(method, fullUrl, payload); err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	c2.Info.Printf("%s %s", method, fullUrl)
	resp, getErr := self.client.Do(req)
	if getErr != nil || resp.Body == nil {
		return nil, getErr
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, c2.NewErrC(string(msg), resp.StatusCode)
	}
	return node.NewJsonReader(resp.Body).Node(), nil
}
