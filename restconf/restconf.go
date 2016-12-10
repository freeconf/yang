package restconf

import (
	"crypto/tls"
	"fmt"
	"io"
	"mime"
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"golang.org/x/net/websocket"
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

func NewService(yangPath meta.StreamSource, root *node.Browser) *Service {
	service := &Service{
		Path:     "/restconf/",
		Root:     root,
		mux:      http.NewServeMux(),
		yangPath: yangPath,
	}
	service.mux.HandleFunc("/.well-known/host-meta", service.resources)
	service.mux.Handle("/restconf/", http.StripPrefix("/restconf/", service))
	service.mux.HandleFunc("/meta/", service.meta)

	service.socketHandler = &WebSocketService{Factory: service}
	service.mux.Handle("/restsock/", websocket.Handler(service.socketHandler.Handle))
	return service
}

type Auth interface {
	ConstrainRoot(r *http.Request, constraints *node.Constraints) error
}

type Service struct {
	Path            string
	yangPath        meta.StreamSource
	Root            *node.Browser
	mux             *http.ServeMux
	docrootSource   *docRootImpl
	DocRoot         string
	Port            string
	Iface           string
	CallbackAddress string
	CallHome        *CallHome
	ReadTimeout     int
	WriteTimeout    int
	socketHandler   *WebSocketService
	Tls             *Tls
	Auth            Auth
}

// SetRootRedirect will direct "/" to any url you want. Useful when you want to server your user
// interface.  Example: SetRootRedirect("/ui/index.html")
func (service *Service) SetRootRedirect(path string) {
	service.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, path, 301)
	})
}

func (service *Service) SetAppVersion(ver string) {
	service.mux.HandleFunc("/.ver", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(ver))
	})
}

func (service *Service) EffectiveCallbackAddress() string {
	if len(service.CallbackAddress) > 0 {
		return service.CallbackAddress
	}
	if len(service.Iface) == 0 {
		panic("No iface given for management port")
	}
	ip := c2.GetIpForIface(service.Iface)
	proto := "http://"
	if service.Tls != nil {
		proto = "https://"
	}
	return fmt.Sprintf("%s%s%s/", proto, ip, service.Port)
}

func (service *Service) GetHttpClient() *http.Client {
	var client *http.Client
	if service.Tls != nil {
		tlsConfig := &tls.Config{
			Certificates: service.Tls.Config.Certificates,
			RootCAs:      service.Tls.Config.RootCAs,
		}
		transport := &http.Transport{TLSClientConfig: tlsConfig}
		client = &http.Client{Transport: transport}
	} else {
		client = http.DefaultClient
	}
	return client
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
	c2.Info.Printf("Starting RESTCONF interface on port %s", service.Port)
	s := &http.Server{
		Addr:           service.Port,
		Handler:        service.mux,
		ReadTimeout:    time.Duration(service.ReadTimeout) * time.Millisecond,
		WriteTimeout:   time.Duration(service.WriteTimeout) * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}
	if service.Tls != nil {
		s.TLSConfig = &service.Tls.Config
		conn, err := net.Listen("tcp", s.Addr)
		if err != nil {
			panic(err)
		}
		tlsListener := tls.NewListener(conn, &service.Tls.Config)
		c2.Err.Fatal(s.Serve(tlsListener))
	} else {
		c2.Err.Fatal(s.ListenAndServe())
	}
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
	if rdr, err := service.docroot.OpenStream(path, ""); err != nil {
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
	m := service.Root.Meta.(*meta.Module)
	_, noexpand := r.URL.Query()["noexpand"]

	sel := node.SelectModule(service.yangPath, m, !noexpand).Root()
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
