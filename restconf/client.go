package restconf

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"golang.org/x/net/websocket"
)

// NewClient interfaces with a remote RESTCONF server.  This also implements conf.Device
// making it appear like a local device and is important architecturaly.  Code that uses
// this in a node.Browser context would not know the difference from a remote or local device
// with one minor exceptions. Peek() wouldn't work.
type Client struct {
	YangPath meta.StreamSource
}

func NewClient(ypath meta.StreamSource) Client {
	return Client{YangPath: ypath}
}

func (self Client) NewDevice(address string) (conf.Device, error) {
	// remove trailing '/' if there is one to prepare for appending
	if address[len(address)-1] == '/' {
		address = address[:len(address)-1]
	}
	remoteSchemaPath := httpStream{
		client: http.DefaultClient,
		url:    address + "/schema/",
	}
	c := &client{
		address:       address,
		yangPath:      self.YangPath,
		schemaPath:    meta.MultipleSources(self.YangPath, remoteSchemaPath),
		client:        http.DefaultClient,
		subscriptions: make(map[string]*clientSubscription),
	}
	d := &clientNode{support: c}
	modules, err := conf.LoadModules(self.YangPath, d.node())
	if err != nil {
		return nil, err
	}
	c.modules = modules
	return c, nil
}

var badAddressErr = c2.NewErr("Expected format: http://server/restconf[=device]/operation/module:path")

type client struct {
	address       string
	yangPath      meta.StreamSource
	schemaPath    meta.StreamSource
	client        *http.Client
	origin        string
	_ws           *websocket.Conn
	subscriptions map[string]*clientSubscription
	modules       map[string]*meta.Module
}

func (self *client) SchemaSource() meta.StreamSource {
	return self.schemaPath
}

func (self *client) UiSource() meta.StreamSource {
	return httpStream{
		client: http.DefaultClient,
		url:    self.address + "/ui/",
	}
}

func (self *client) Browser(module string) (*node.Browser, error) {
	d := &clientNode{support: self}
	m, err := self.module(module)
	if err != nil {
		return nil, err
	}
	return node.NewBrowser(m, d.node()), nil
}

func (self *client) Close() {
	if self._ws != nil {
		self._ws.Close()
		self._ws = nil
	}
}

func (self *client) Modules() map[string]*meta.Module {
	return self.modules
}

func (self *client) clientSocket() (io.Writer, error) {
	// lazy start websocket connection to be more efficient if it's never used
	// but I have no data how how much resources this saves
	if self._ws == nil {
		urlParts, err := url.Parse(self.address)
		if err != nil {
			return nil, err
		}
		wsUrl := "ws://" + urlParts.Host + "/restconf/streams"
		origin := self.origin
		if origin == "" {
			origin = "http://" + urlParts.Host
		}
		if self._ws, err = websocket.Dial(wsUrl, "", origin); err != nil {
			return nil, err
		}
		go self.watch(self._ws)
	}
	return self._ws, nil
}

func (self *client) watch(ws io.Reader) {
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
		idVal := notification["id"]
		if l := self.subscriptions[idVal.(string)]; l != nil {
			n := node.NewJsonReader(strings.NewReader(payload)).Node()
			l.notify(l.sel.Split(n))
		} else {
			c2.Info.Printf("no listener found with id %s", idVal)
		}
	}
}

func (self *client) clientSubscriptions() map[string]*clientSubscription {
	return self.subscriptions
}

func (self *client) module(module string) (*meta.Module, error) {
	// caching module, but should replace w/cache that can refresh on stale
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

func (self *client) clientDo(method string, params string, p *node.Path, payload io.Reader) (node.Node, error) {
	var req *http.Request
	var err error
	mod := getModule(p)
	fullUrl := self.address + "/data/" + mod.GetIdent() + ":" + p.StringNoModule()
	if params != "" {
		fullUrl += "?" + params
	}
	if req, err = http.NewRequest(method, fullUrl, payload); err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	c2.Info.Printf("=> %s %s", method, fullUrl)
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
