package restconf

import (
	"fmt"
	"mime"
	"net/http"
	"strings"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/stock"
	"golang.org/x/net/websocket"
)

type WebServer interface {
	HandleFunc(string, http.HandlerFunc)
	Handle(string, http.Handler)
}

func NewService(yangPath meta.StreamSource, root *node.Browser) *Service {
	return &Service{
		Path:     "/restconf/",
		Root:     root,
		yangPath: yangPath,
	}
}

func (service *Service) SetWebServer(src WebServer) {
	src.HandleFunc("/.well-known/host-meta", service.resources)
	src.Handle(service.Path, http.StripPrefix(service.Path, service))
	src.HandleFunc("/meta/", service.meta)

	service.socketHandler = &WsNotifyService{Factory: service}
	src.Handle("/restsock/", websocket.Handler(service.socketHandler.Handle))
}

type Auth interface {
	ConstrainRoot(r *http.Request, constraints *node.Constraints) error
}

type ServiceOptions struct {
	stock.HttpServerOptions
	NotifyKeepaliveTimeoutMs int
	Path                     string
}

type Management struct {
	options  ServiceOptions
	Handler  *Service
	mux      *http.ServeMux
	Web      *stock.HttpServer
	CallHome *CallHome
}

func NewManagement(yangPath meta.StreamSource, root *node.Browser) *Management {
	m := &Management{
		Handler: NewService(yangPath, root),
		Web:     stock.NewHttpServer(),
		mux:     http.NewServeMux(),
	}
	m.Handler.SetWebServer(m)
	return m
}

func (self *Management) Options() ServiceOptions {
	return self.options
}

func (self *Management) ApplyOptions(options ServiceOptions) {
	if options == self.options {
		return
	}
	self.options = options
	self.Handler.socketHandler.Timeout = self.options.NotifyKeepaliveTimeoutMs
	self.Handler.Path = self.options.Path
	self.Web.ApplyOptions(options.HttpServerOptions, self.mux)
}

func (self *Management) Stop() {
	self.Web.Stop()
}

func (service *Management) HandleFunc(path string, handler http.HandlerFunc) {
	service.mux.HandleFunc(path, handler)
}

func (service *Management) Handle(path string, handler http.Handler) {
	service.mux.Handle(path, handler)
}

type Service struct {
	Path          string
	yangPath      meta.StreamSource
	Root          *node.Browser
	socketHandler *WsNotifyService
	Auth          Auth
}

func (service *Service) handleError(err error, w http.ResponseWriter) {
	if httpErr, ok := err.(c2.HttpError); ok {
		if httpErr.HttpCode() >= 500 {
			c2.Err.Print(httpErr.Error() + "\n" + httpErr.Stack())
		}
		http.Error(w, httpErr.Error(), httpErr.HttpCode())
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (self *Service) Subscribe(sub *node.Subscription) error {
	if sel := self.Root.Root().Find(sub.Path); sel.LastErr == nil {
		closer, notifSel := sel.Notifications(sub)
		if notifSel.LastErr != nil {
			return notifSel.LastErr
		}
		sub.Notification = notifSel.Meta().(*meta.Notification)
		sub.Closer = closer
	} else {
		return sel.LastErr
	}
	return nil
}

func (self *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Access-Control-Allow-Headers", "origin, content-type, accept")
	h.Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE, PATCH")
	h.Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		return
	}
	var err error
	var payload node.Node
	sel := self.Root.Root()
	if self.Auth != nil {
		if err = self.Auth.ConstrainRoot(r, sel.Constraints); err != nil {
			self.handleError(err, w)
			return
		}
	}

	// Noisey, but very useful and acts as Access log
	c2.Info.Printf("%s %s", r.Method, r.URL)

	if sel = sel.FindUrl(r.URL); sel.LastErr == nil {
		if sel.IsNil() {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		if err != nil {
			self.handleError(err, w)
			return
		}
		switch r.Method {
		case "DELETE":
			err = sel.Delete()
		case "GET":
			w.Header().Set("Content-Type", mime.TypeByExtension(".json"))
			output := node.NewJsonWriter(w).Node()
			err = sel.InsertInto(output).LastErr
		case "PUT":
			err = sel.UpsertFrom(node.NewJsonReader(r.Body).Node()).LastErr
		case "POST":
			if meta.IsAction(sel.Meta()) {
				input := node.NewJsonReader(r.Body).Node()
				if outputSel := sel.Action(input); !outputSel.IsNil() {
					w.Header().Set("Content-Type", mime.TypeByExtension(".json"))
					err = outputSel.InsertInto(node.NewJsonWriter(w).Node()).LastErr
				} else {
					err = outputSel.LastErr
				}
			} else {
				payload = node.NewJsonReader(r.Body).Node()
				err = sel.InsertFrom(payload).LastErr
			}
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	} else {
		err = sel.LastErr
	}

	if err != nil {
		self.handleError(err, w)
	}
}

func (service *Service) meta(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	if p := strings.TrimPrefix(r.URL.Path, "/meta/"); len(p) < len(r.URL.Path) {
		r.URL.Path = p
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	m := service.Root.Meta.(*meta.Module)
	_, noexpand := r.URL.Query()["noexpand"]

	sel := node.SelectModule(m, !noexpand).Root()
	if sel = sel.FindUrl(r.URL); sel.LastErr != nil {
		service.handleError(sel.LastErr, w)
		return
	} else if sel.IsNil() {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else {
		w.Header().Set("Content-Type", mime.TypeByExtension(".json"))
		output := node.NewJsonWriter(w).Node()
		if err := sel.InsertInto(output).LastErr; err != nil {
			service.handleError(err, w)
			return
		}
	}
}

func (service *Service) resources(w http.ResponseWriter, r *http.Request) {
	// RESTCONF Sec. 3.1
	fmt.Fprintf(w, `"xrd" : { "link" : { "@rel" : "restconf", "@href" : "/restconf" } } }`)
}
