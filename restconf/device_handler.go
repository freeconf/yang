package restconf

import (
	"bytes"
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"

	"net/url"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"golang.org/x/net/websocket"
)

type DeviceHandler struct {
	Ver                      string
	NotifyKeepaliveTimeoutMs int
	main                     device.Device
	devices                  *device.Map
	notifiers                *list.List
}

func NewDeviceHandler() *DeviceHandler {
	m := &DeviceHandler{
		notifiers: list.New(),
	}
	return m
}

func (self *DeviceHandler) ServeDevices(m *device.Map) error {
	self.devices = m
	return nil
}

func (self *DeviceHandler) ServeDevice(d device.Device) error {
	self.main = d
	return nil
}

func (self *DeviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
			self.serveNotifications(w, r)
		case "data":
			self.serveData(device, w, r)
		case "ui":
			self.serveStreamSource(w, device.UiSource(), r.URL.Path)
		case "schema":
			self.serveStreamSource(w, device.SchemaSource(), r.URL.Path)
		default:
			handleErr(badAddressErr, w)
		}
	}
}

func (self *DeviceHandler) serveData(d device.Device, w http.ResponseWriter, r *http.Request) {
	if hndlr, p := self.shiftBrowserHandler(d, w, r.URL); hndlr != nil {
		r.URL = p
		hndlr.ServeHTTP(w, r)
	}
}

func (self *DeviceHandler) Subscribe(sub *node.Subscription) error {
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

func (self *DeviceHandler) serveNotifications(w http.ResponseWriter, r *http.Request) {
	socketHndlr := &wsNotifyService{
		factory: self,
		timeout: self.NotifyKeepaliveTimeoutMs,
	}
	elem := self.notifiers.PushBack(socketHndlr)
	defer self.notifiers.Remove(elem)
	websocket.Handler(socketHndlr.Handle).ServeHTTP(w, r)
}

func (self *DeviceHandler) SubscriptionCount() int {
	var c int
	p := self.notifiers.Front()
	for p != nil {
		c += p.Value.(*wsNotifyService).conn.mgr.Len()
		p = p.Next()
	}
	return c
}

func (self *DeviceHandler) serveStreamSource(w http.ResponseWriter, s meta.StreamSource, path string) {
	rdr, err := s.OpenStream(path, "")
	if err != nil {
		handleErr(err, w)
		return
	}
	ext := filepath.Ext(path)
	ctype := mime.TypeByExtension(ext)
	w.Header().Set("Content-Type", ctype)
	if _, err := io.Copy(w, rdr); err != nil {
		handleErr(err, w)
	}
}

func (self *DeviceHandler) findDevice(deviceId string) (device.Device, error) {
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

func (self *DeviceHandler) shiftOperationAndDevice(w http.ResponseWriter, orig *url.URL) (string, device.Device, *url.URL) {
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

func (self *DeviceHandler) shiftBrowserHandler(d device.Device, w http.ResponseWriter, orig *url.URL) (*browserHandler, *url.URL) {
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

func (self *DeviceHandler) serveStaticRoute(w http.ResponseWriter, r *http.Request) bool {
	op, _ := shift(r.URL, '/')
	switch op {
	case "host-meta":
		// RESTCONF Sec. 3.1
		fmt.Fprintf(w, `{ "xrd" : { "link" : { "@rel" : "restconf", "@href" : "/restconf" } } }`)
		return true
	}
	return false
}
