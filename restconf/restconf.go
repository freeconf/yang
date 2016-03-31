package restconf

import (
	"github.com/c2g/c2"
	"github.com/c2g/node"
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"github.com/c2g/meta"
	"time"
	"strings"
	"github.com/c2g/meta/yang"
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
	}
	service.mux.HandleFunc("/.well-known/host-meta", service.resources)
	service.mux.Handle("/restconf/", http.StripPrefix("/restconf/", service))
	service.mux.HandleFunc("/meta/", service.meta)
	return service
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

func (service *Service) Manage() node.Node {
	s := &node.MyNode{Peekables: map[string]interface{}{"internal": service}}
	s.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "callHome":
			if r.New {
				service.CallHome = &CallHome{
					EndpointAddress: service.EffectiveCallbackAddress(),
					Module: service.Root.Select().Meta().(*meta.Module),
				}
			}
			if service.CallHome != nil {
				return service.CallHome.Manage(), nil
			}
		}
		return nil, nil
	}
	s.OnRead = func(r node.FieldRequest) (*node.Value, error) {
		return node.ReadField(r.Meta, service)
	}
	s.OnWrite = func(r node.FieldRequest, v *node.Value) (err error) {
		switch r.Meta.GetIdent() {
		case "docRoot":
			service.DocRoot = v.Str
			service.SetDocRoot(&meta.FileStreamSource{Root: service.DocRoot})
		}
		return node.WriteField(r.Meta, service, v)
	}
	return s
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

func (service *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func init() {
	yang.InternalYang()["restconf"] = `
module restconf {
	namespace "http://org.conf2/ns/management";
	prefix "restconf";
	revision 0;

    grouping management {
        leaf port {
            type string;
        }
        /* looks at first ip address for iface, use callbackAddress to explicitly set callback */
        leaf iface {
            type string;
            default "eth0";
        }
        /* optional, will determine callback automatically based on iface ip */
        leaf callbackAddress {
            type string;
        }
        leaf docRoot {
            type string;
        }
        leaf path {
            type string;
            default "/restconf/";
        }
        container callHome {
            leaf controllerAddress {
                type string;
            }
            /*
             optional, will determine automatically otherwise based on
             restconf's ip address and port
            */
            leaf endpointAddress {
                type string;
            }
            leaf endpointId {
                type string;
            }
            container registration {
                config "false";
                leaf id {
                    type string;
                }
            }
        }
    }
}
`
}