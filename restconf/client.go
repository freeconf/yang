package restconf

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/device"
	"github.com/freeconf/gconf/meta"
	"github.com/freeconf/gconf/meta/yang"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/nodes"
)

// NewClient interfaces with a remote RESTCONF server.  This also implements device.Device
// making it appear like a local device and is important architecturaly.  Code that uses
// this in a node.Browser context would not know the difference from a remote or local device
// with one minor exceptions. Peek() wouldn't work.
type Client struct {
	YangPath meta.StreamSource
}

func ProtocolHandler(ypath meta.StreamSource) device.ProtocolHandler {
	c := Client{YangPath: ypath}
	return c.NewDevice
}

type Address struct {
	Base     string
	Data     string
	Stream   string
	Ui       string
	Schema   string
	DeviceId string
	Host     string
	Origin   string
}

func NewAddress(urlAddr string) (Address, error) {
	// remove trailing '/' if there is one to prepare for appending
	if urlAddr[len(urlAddr)-1] != '/' {
		urlAddr = urlAddr + "/"
	}

	urlParts, err := url.Parse(urlAddr)
	if err != nil {
		return Address{}, err
	}

	return Address{
		Base:     urlAddr,
		Data:     urlAddr + "data/",
		Schema:   urlAddr + "schema/",
		Ui:       urlAddr + "ui/",
		Origin:   "http://" + urlParts.Host,
		DeviceId: findDeviceIdInUrl(urlAddr),
	}, nil
}

func findDeviceIdInUrl(addr string) string {
	segs := strings.SplitAfter(addr, "/restconf=")
	if len(segs) == 2 {
		post := segs[1]
		return post[:len(post)-1]
	}
	return ""
}

func (self Client) NewDevice(url string) (device.Device, error) {
	address, err := NewAddress(url)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	remoteSchemaPath := httpStream{
		ypath:  self.YangPath,
		client: httpClient,
		url:    address.Schema,
	}
	c := &client{
		address:    address,
		yangPath:   self.YangPath,
		schemaPath: meta.MultipleSources(self.YangPath, remoteSchemaPath),
		client:     httpClient,
	}
	d := &clientNode{support: c, device: address.DeviceId}
	m := yang.RequireModule(self.YangPath, "ietf-yang-library")
	b := node.NewBrowser(m, d.node())
	modules, err := device.LoadModules(b, remoteSchemaPath)
	c2.Debug.Printf("loaded modules %v", modules)
	if err != nil {
		return nil, fmt.Errorf("could not load modules. %s", err)
	}
	c.modules = modules
	return c, nil
}

var badAddressErr = c2.NewErr("Expected format: http://server/restconf[=device]/operation/module:path")

type client struct {
	address    Address
	yangPath   meta.StreamSource
	schemaPath meta.StreamSource
	client     *http.Client
	origin     string
	modules    map[string]*meta.Module
}

func (self *client) SchemaSource() meta.StreamSource {
	return self.schemaPath
}

func (self *client) UiSource() meta.StreamSource {
	return httpStream{
		client: self.client,
		url:    self.address.Ui,
	}
}

func (self *client) Browser(module string) (*node.Browser, error) {
	d := &clientNode{support: self, device: self.address.DeviceId}
	m, err := self.module(module)
	if err != nil {
		return nil, err
	}
	return node.NewBrowser(m, d.node()), nil
}

func (self *client) Close() {

}

func (self *client) Modules() map[string]*meta.Module {
	return self.modules
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
	ypath  meta.StreamSource
	client *http.Client
	url    string
}

func (self httpStream) ResolveModuleHnd(hnd device.ModuleHnd) (*meta.Module, error) {
	m, _ := yang.LoadModule(self.ypath, hnd.Name)
	if m != nil {
		return m, nil
	}
	return yang.LoadModule(self, hnd.Name)
}

// OpenStream implements meta.StreamSource
func (self httpStream) OpenStream(name string, ext string) (meta.DataStream, error) {
	fullUrl := self.url + name + ext
	c2.Debug.Printf("httpStream url %s, name=%s, ext=%s", fullUrl, name, ext)
	resp, err := self.client.Get(fullUrl)
	if resp != nil {
		return resp.Body, err
	}
	return nil, err
}

func (self *client) clientDo(method string, params string, p *node.Path, payload io.Reader) (node.Node, error) {
	var req *http.Request
	var err error
	mod := meta.Root(p.Meta())
	fullUrl := fmt.Sprint(self.address.Data, mod.Ident(), ":", p.StringNoModule())
	if params != "" {
		fullUrl = fmt.Sprint(fullUrl, "?", params)
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
	return nodes.ReadJSONIO(resp.Body), nil
}

func (self *client) clientStream(params string, p *node.Path, ctx context.Context) (<-chan node.Node, error) {
	mod := meta.Root(p.Meta())
	fullUrl := fmt.Sprint(self.address.Data, mod.Ident(), ":", p.StringNoModule())
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/event-stream")
	c2.Info.Printf("<=> SSE %s", fullUrl)
	resp, err := self.client.Do(req)
	if err != nil {
		return nil, err
	}
	stream := make(chan node.Node)
	go func() {
		events := decodeSse(resp.Body)
		defer resp.Body.Close()
		for {
			select {
			case event := <-events:
				stream <- nodes.ReadJSONIO(bytes.NewReader(event))
			case <-ctx.Done():
				return
			}
		}
	}()

	return stream, nil
}
