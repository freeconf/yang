package restconf

import (
	"bytes"
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"golang.org/x/net/websocket"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/device"
	"github.com/freeconf/gconf/meta"
	"github.com/freeconf/gconf/meta/yang"
	"github.com/freeconf/gconf/nodes"
	"github.com/freeconf/gconf/secure"
	"github.com/freeconf/gconf/stock"
)

type Server struct {
	Web                      *stock.HttpServer
	CallHome                 *device.CallHome
	Auth                     secure.Auth
	Ver                      string
	NotifyKeepaliveTimeoutMs int
	main                     device.Device
	devices                  device.Map
	notifiers                *list.List
	web                      *stock.HttpServer
	ypath                    meta.StreamSource

	// Optional: Anything not handled by RESTCONF protocol can call this handler otherwise
	UnhandledRequestHandler http.HandlerFunc
}

func NewServer(d *device.Local) *Server {
	s := NewWebHandler(d)
	return s
}

func (self *Server) Close() {
	if self.web != nil {
		self.web.Server.Close()
	}
}

func NewWebHandler(d *device.Local) *Server {
	m := &Server{
		notifiers: list.New(),
		ypath:     d.SchemaSource(),
	}
	m.ServeDevice(d)

	if err := d.Add("restconf", Node(m, d.SchemaSource())); err != nil {
		panic(err)
	}

	// Required by all devices according to RFC
	if err := d.Add("ietf-yang-library", device.LocalDeviceYangLibNode(m.ModuleAddress, d)); err != nil {
		panic(err)
	}
	return m
}

func (self *Server) ModuleAddress(m *meta.Module) string {
	return fmt.Sprint("schema/", m.Ident(), ".yang")
}

func (self *Server) DeviceAddress(id string, d device.Device) string {
	return fmt.Sprint("/restconf=", id)
}

func (self *Server) ServeDevices(m device.Map) error {
	self.devices = m
	return nil
}

func (self *Server) ServeDevice(d device.Device) error {
	self.main = d
	return nil
}

func (self *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if c2.DebugLogEnabled() {
		c2.Debug.Printf("%s %s", r.Method, r.URL)
		if r.Body != nil {
			content, rerr := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if rerr != nil {
				c2.Err.Printf("error trying to log body content %s", rerr)
			} else {
				if len(content) > 0 {
					c2.Debug.Print(string(content))
					r.Body = ioutil.NopCloser(bytes.NewBuffer(content))
				}
			}
		}
	}

	h := w.Header()

	// CORS
	h.Set("Access-Control-Allow-Headers", "origin, content-type, accept")
	h.Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE, PATCH")
	h.Set("Access-Control-Allow-Origin", "*")
	if r.URL.Path == "/" {
		switch r.Method {
		case "OPTIONS":
			return
		case "GET":
			http.Redirect(w, r, "restconf/ui/index.html", http.StatusMovedPermanently)
			return
		}
	}

	op1, deviceId, p := shiftOptionalParamWithinSegment(r.URL, '=', '/')
	device, err := self.findDevice(deviceId)
	if err != nil {
		handleErr(err, w)
		return
	}
	switch op1 {
	case ".ver":
		w.Write([]byte(self.Ver))
	case ".well-known":
		self.serveStaticRoute(w, r)
	case "restconf":
		op2, p := shift(p, '/')
		r.URL = p
		switch op2 {
		case "streams":
			// no device id in url is normal. this allows client to share one websocket connection
			// to support many devices.
			self.serveNotifications(w, r)
		case "data":
			self.serveData(device, w, r)
		case "ui":
			self.serveStreamSource(w, device.UiSource(), r.URL.Path)
		case "schema":
			// Hack - parse accept header to get proper content type
			accept := r.Header.Get("Accept")
			c2.Debug.Printf("accept %s", accept)
			if strings.Contains(accept, "/json") {
				self.serveSchema(w, r, device.SchemaSource())
			} else {
				self.serveStreamSource(w, device.SchemaSource(), r.URL.Path)
			}
		default:
			handleErr(badAddressErr, w)
		}
	default:
		if self.UnhandledRequestHandler != nil {
			self.UnhandledRequestHandler(w, r)
			return
		}
	}
}

func (self *Server) serveSchema(w http.ResponseWriter, r *http.Request, ypath meta.StreamSource) {
	modName, p := shift(r.URL, '/')
	r.URL = p
	m, err := yang.LoadModule(ypath, modName)
	if err != nil {
		handleErr(err, w)
		return
	}
	ylib, err := yang.LoadModule(ypath, "yang")
	if err != nil {
		handleErr(err, w)
		return
	}
	b := nodes.Schema(ylib, m)
	hndlr := &browserHandler{browser: b}
	hndlr.ServeHTTP(w, r)
}

func (self *Server) serveData(d device.Device, w http.ResponseWriter, r *http.Request) {
	if hndlr, p := self.shiftBrowserHandler(d, w, r.URL); hndlr != nil {
		r.URL = p
		hndlr.ServeHTTP(w, r)
	}
}

func (self *Server) Subscribe(sub *Subscription) error {
	device, err := self.findDevice(sub.DeviceId)
	if err != nil {
		return err
	}
	b, err := device.Browser(sub.Module)
	if err != nil {
		return err
	} else if b == nil {
		return c2.NewErrC("No module found:"+sub.Module, 404)
	}
	if sel := b.Root().Find(sub.Path); sel.LastErr == nil {
		closer, err := sel.Notifications(sub.Notify)
		if err != nil {
			return err
		}
		sub.Notification = sel.Meta().(*meta.Notification)
		sub.Closer = closer
	} else {
		return sel.LastErr
	}
	return nil
}

func (self *Server) serveNotifications(w http.ResponseWriter, r *http.Request) {
	socketHndlr := &wsNotifyService{
		factory: self,
		timeout: self.NotifyKeepaliveTimeoutMs,
	}
	elem := self.notifiers.PushBack(socketHndlr)
	defer self.notifiers.Remove(elem)
	websocket.Handler(socketHndlr.Handle).ServeHTTP(w, r)
}

func (self *Server) SubscriptionCount() int {
	var c int
	p := self.notifiers.Front()
	for p != nil {
		c += p.Value.(*wsNotifyService).conn.mgr.Len()
		p = p.Next()
	}
	return c
}

func (self *Server) serveStreamSource(w http.ResponseWriter, s meta.StreamSource, path string) {
	rdr, err := s.OpenStream(path, "")
	if err != nil {
		handleErr(err, w)
		return
	} else if rdr == nil {
		handleErr(c2.NewErrC("not found", 404), w)
		return
	}
	ext := filepath.Ext(path)
	ctype := mime.TypeByExtension(ext)
	w.Header().Set("Content-Type", ctype)
	if _, err := io.Copy(w, rdr); err != nil {
		handleErr(err, w)
	}
}

func (self *Server) findDevice(deviceId string) (device.Device, error) {
	if deviceId == "" {
		return self.main, nil
	}
	device, err := self.devices.Device(deviceId)
	if err != nil {
		return nil, err
	}
	if device == nil {
		return nil, c2.NewErrC("device not found "+deviceId, 404)
	}
	return device, nil
}

func (self *Server) shiftOperationAndDevice(w http.ResponseWriter, orig *url.URL) (string, device.Device, *url.URL) {
	//  operation[=deviceId]/...
	op, deviceId, p := shiftOptionalParamWithinSegment(orig, '=', '/')
	if op == "" {
		handleErr(c2.NewErrC("no operation found in path", 404), w)
		return op, nil, orig
	}
	device, err := self.findDevice(deviceId)
	if err != nil {
		handleErr(err, w)
		return "", nil, orig
	}
	return op, device, p
}

func (self *Server) shiftBrowserHandler(d device.Device, w http.ResponseWriter, orig *url.URL) (*browserHandler, *url.URL) {
	if module, p := shift(orig, ':'); module != "" {
		if browser, err := d.Browser(module); browser != nil {
			return &browserHandler{
				browser: browser,
			}, p
		} else if err != nil {
			handleErr(err, w)
			return nil, orig
		}
	}

	handleErr(c2.NewErrC("no module found in path", 404), w)
	return nil, orig
}

func (self *Server) serveStaticRoute(w http.ResponseWriter, r *http.Request) bool {
	op, _ := shift(r.URL, '/')
	switch op {
	case "host-meta":
		// RESTCONF Sec. 3.1
		fmt.Fprintf(w, `{ "xrd" : { "link" : { "@rel" : "restconf", "@href" : "/restconf" } } }`)
		return true
	}
	return false
}
