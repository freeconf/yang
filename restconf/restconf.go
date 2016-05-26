package restconf

import (
	"fmt"
	"github.com/c2g/c2"
	"github.com/c2g/meta"
	"github.com/c2g/node"
	"golang.org/x/net/websocket"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type restconfError struct {
	Code int
	Msg  string
}

func (err *restconfError) Error() string {
	return err.Msg
}

func (err *restconfError) HttpCode() int {
	return err.Code
}

func NewService(root node.Data) *Service {
	service := &Service{
		Path: "/restconf/",
		Root: root,
		mux:  http.NewServeMux(),
		webSocket: &WebSocket{},
	}
	service.mux.HandleFunc("/.well-known/host-meta", service.resources)
	service.mux.Handle("/restconf/", http.StripPrefix("/restconf/", service))
	service.mux.HandleFunc("/meta/", service.meta)
	service.mux.Handle("/restsock/", websocket.Handler(service.socketHandler))
	return service
}

func (self *Service) socketHandler(ws *websocket.Conn) {
	self.webSocket.ConnectionHandler(ws, self)
}

type Service struct {
	Path            string
	Root            node.Data
	mux             *http.ServeMux
	docrootSource   *docRootImpl
	DocRoot         string
	Port            string
	Iface           string
	CallbackAddress string
	CallHome        *CallHome
	webSocket       *WebSocket
}

func (service *Service) EffectiveCallbackAddress() string {
	if len(service.CallbackAddress) > 0 {
		return service.CallbackAddress
	}
	if len(service.Iface) == 0 {
		panic("No iface given for management port")
	}
	ip := c2.GetIpForIface(service.Iface)
	return fmt.Sprintf("http://%s%s/", ip, service.Port)
}

type registration struct {
	browser node.Data
}

func (service *Service) handleError(err error, w http.ResponseWriter) {
	if httpErr, ok := err.(c2.HttpError); ok {
		c2.Err.Print(httpErr.Error() + "\n" + httpErr.Stack())
		http.Error(w, httpErr.Error(), httpErr.HttpCode())
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (self *Service) NewChannel(channel *node.NotifyChannel, url string) {
	c := node.NewContext()
	if sel := c.Selector(self.Root.Select()).Find(url); sel.LastErr == nil {
		notifSel := sel.Notifications(channel)
		if notifSel.LastErr != nil {
			panic(notifSel.LastErr)
		}
		channel.Notification = notifSel.Selection.Meta().(*meta.Notification)
	} else {
		panic(sel.LastErr)
	}
}

func (service *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Access-Control-Allow-Headers", "origin, content-type, accept")
	h.Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE, PATCH")
	h.Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		return
	}
	var err error
	var payload node.Node
	var sel node.Selector
	c := node.NewContext()
	if sel = c.Selector(service.Root.Select()).FindUrl(r.URL); sel.LastErr == nil {
		if sel.Selection == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		if err != nil {
			service.handleError(err, w)
			return
		}
		switch r.Method {
		case "DELETE":
			err = sel.Selection.Delete()
		case "GET":
			w.Header().Set("Content-Type", mime.TypeByExtension(".json"))
			output := node.NewJsonWriter(w).Node()
			err = sel.InsertInto(output).LastErr
		case "PUT":
			err = sel.UpsertFrom(node.NewJsonReader(r.Body).Node()).LastErr
		case "POST":
			if meta.IsAction(sel.Selection.Meta()) {
				input := node.NewJsonReader(r.Body).Node()
				if outputSel := sel.Action(input); outputSel.Selection != nil {
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
		service.handleError(err, w)
	}
}

type docRootImpl struct {
	docroot meta.StreamSource
}

func (service *Service) SetDocRoot(docroot meta.StreamSource) {
	service.docrootSource = &docRootImpl{docroot: docroot}
	service.mux.Handle("/ui/", http.StripPrefix("/ui/", service.docrootSource))
}

func (service *Service) AddHandler(pattern string, handler http.Handler) {
	service.mux.Handle(pattern, http.StripPrefix(pattern, handler))
}

func (service *Service) Listen() {
	s := &http.Server{
		Addr:           service.Port,
		Handler:        service.mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	c2.Info.Println("Starting RESTCONF interface")
	c2.Err.Fatal(s.ListenAndServe())
}

func (service *Service) Stop() {
	if service.docrootSource != nil && service.docrootSource.docroot != nil {
		meta.CloseResource(service.docrootSource.docroot)
	}
	// TODO - actually stop service
}

func (service *docRootImpl) ServeHTTP(wtr http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if path == "" {
		path = "index.html"
	}
	if rdr, err := service.docroot.OpenStream(path); err != nil {
		http.Error(wtr, err.Error(), http.StatusInternalServerError)
	} else {
		defer meta.CloseResource(rdr)
		ext := filepath.Ext(path)
		ctype := mime.TypeByExtension(ext)
		wtr.Header().Set("Content-Type", ctype)
		if _, err = io.Copy(wtr, rdr); err != nil {
			http.Error(wtr, err.Error(), http.StatusInternalServerError)
		}
		// Eventually support this but need file seeker to do that.
		// http.ServeContent(wtr, req, path, time.Now(), &ReaderPeeker{rdr})
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
	m := service.Root.Select().Meta().(*meta.Module)
	_, noexpand := r.URL.Query()["noexpand"]

	c := node.NewContext()
	sel := c.Selector(node.SelectModule(m, !noexpand))
	if sel = sel.FindUrl(r.URL); sel.LastErr != nil {
		service.handleError(sel.LastErr, w)
		return
	} else if sel.Selection == nil {
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
