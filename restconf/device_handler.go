package restconf

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"context"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/meta"
	"golang.org/x/net/websocket"
)

type DeviceHandler struct {
	NotifyKeepaliveTimeoutMs int
	devices                  map[string]conf.Device
}

func NewDeviceHandler() *DeviceHandler {
	m := &DeviceHandler{
		devices: make(map[string]conf.Device),
	}
	return m
}

func (self *DeviceHandler) ServeDevice(d conf.Device, path string) error {
	c2.Info.Print("restconf mount ", path)
	self.devices[path] = d
	return nil
}

func (self *DeviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Access-Control-Allow-Headers", "origin, content-type, accept")
	h.Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE, PATCH")
	h.Set("Access-Control-Allow-Origin", "*")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// NOTE: Could be ipv4 or ipv6 address.  There doesn't appear to be SDK call to parse
	// ip host:port strings so looking for last colon seems to be universal
	colon := strings.LastIndex(r.RemoteAddr, ":")
	if colon > 0 {
		ctx = context.WithValue(ctx, conf.RemoteIpAddressKey, r.RemoteAddr[:colon])
	}

	if r.Method == "OPTIONS" && r.URL.Path == "/" {
		return
	}
	if self.handleStaticRoute(w, r) {
		return
	} else if path, device := self.findDevice(r.URL.Path); device != nil {
		var operation string
		operation, path = shiftOperation(path)
		if operation == "schema" {
			self.serveStream(w, device.SchemaSource(), path)
			return
		} else if operation == "ui" {
			self.serveStream(w, device.UiSource(), path)
			return
		}
		copy := *r.URL
		module, path := shiftModule(path)
		if module == "" {
			HandleError(badAddressErr, w)
		}
		copy.Path = path
		browser, err := device.Browser(module)
		if err != nil {
			HandleError(err, w)
			return
		}
		if browser == nil {
			HandleError(c2.NewErrC(module+" not found ", 404), w)
			return
		}
		hndlr := &BrowserHandler{
			Browser: browser,
			Path:    &copy,
		}
		if operation == "streams" {
			socketHandler := &WsNotifyService{
				Factory: hndlr,
				Timeout: self.NotifyKeepaliveTimeoutMs,
			}
			websocket.Handler(socketHandler.Handle).ServeHTTP(w, r)
		} else if operation == "data" {
			hndlr.ServeHTTP(ctx, w, r)
		} else {
			HandleError(badAddressErr, w)
		}
	} else if ui, mount := self.findUiStream(w, r); mount != nil {
		if ui == nil {
			return
		}
		defer meta.CloseResource(ui)
		ext := filepath.Ext(path)
		ctype := mime.TypeByExtension(ext)
		w.Header().Set("Content-Type", ctype)
		if _, err := io.Copy(w, ui); err != nil {
			HandleError(err, w)
		}
		// Eventually support this but need file seeker to do that.
		// http.ServeContent(wtr, req, path, time.Now(), &ReaderPeeker{rdr})
	} else {
		HandleError(c2.NewErrC("no mount point found", 404), w)
	}
}

func (self *DeviceHandler) serveStream(w http.ResponseWriter, s meta.StreamSource, path string) {
	rdr, err := s.OpenStream(path, "")
	if err != nil {
		HandleError(err, w)
		return
	}
	ext := filepath.Ext(path)
	ctype := mime.TypeByExtension(ext)
	w.Header().Set("Content-Type", ctype)
	if _, err := io.Copy(w, rdr); err != nil {
		HandleError(err, w)
	}
}

func (self *DeviceHandler) findUiStream(w http.ResponseWriter, r *http.Request) (meta.DataStream, conf.Device) {
	path := r.URL.Path[1:]
	for _, device := range self.devices {
		if u := device.UiSource(); u != nil {
			if rdr, err := u.OpenStream(path, ""); err != nil {
				http.Error(w, err.Error(), 500)
				return nil, device
			} else if rdr != nil {
				return rdr, device
			}
		}
	}
	return nil, nil
}

func (self *DeviceHandler) findDevice(path string) (string, conf.Device) {
	for prefix, device := range self.devices {
		if strings.HasPrefix(path, prefix) {
			return path[len(prefix):], device
		}
	}
	return "", nil
}

func shiftModule(path string) (string, string) {
	colon := strings.IndexRune(path, ':')
	if colon == -1 {
		return path, ""
	}
	return path[:colon], path[colon+1:]
}

func shiftOperation(path string) (string, string) {
	slash := strings.IndexRune(path, '/')
	if slash == -1 {
		return path, ""
	}
	return path[:slash], path[slash+1:]
}

func (self *DeviceHandler) handleStaticRoute(w http.ResponseWriter, r *http.Request) bool {
	switch r.URL.Path {
	case "/.well-known/host-meta":
		// RESTCONF Sec. 3.1
		fmt.Fprintf(w, `{ "xrd" : { "link" : { "@rel" : "restconf", "@href" : "/restconf" } } }`)
		return true
	}
	return false
}
