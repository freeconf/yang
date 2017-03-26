package restconf

import (
	"mime"
	"net/http"
	"net/url"

	"context"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

type BrowserHandler struct {
	Browser *node.Browser
	Path    *url.URL
}

func HandleError(err error, w http.ResponseWriter) {
	if httpErr, ok := err.(c2.HttpError); ok {
		if httpErr.HttpCode() >= 500 {
			c2.Err.Print(httpErr.Error() + "\n" + httpErr.Stack())
		}
		http.Error(w, httpErr.Error(), httpErr.HttpCode())
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (self *BrowserHandler) Subscribe(c context.Context, sub *node.Subscription) error {
	if sel := self.Browser.Root().Find(sub.Path); sel.LastErr == nil {
		closer, err := sel.NotificationsCntx(c, sub.Notify)
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

func (self *BrowserHandler) ServeHTTP(c context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	var payload node.Node
	sel := self.Browser.Root()
	// Noisey, but very useful and acts as Access log
	c2.Info.Printf("%s %s", r.Method, self.Path)

	if sel = sel.FindUrl(self.Path); sel.LastErr == nil {
		if sel.IsNil() {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		if err != nil {
			HandleError(err, w)
			return
		}
		switch r.Method {
		case "DELETE":
			err = sel.Delete()
		case "GET":
			w.Header().Set("Content-Type", mime.TypeByExtension(".json"))
			output := node.NewJsonWriter(w).Node()
			err = sel.InsertIntoCntx(c, output).LastErr
		case "PUT":
			err = sel.UpsertFromCntx(c, node.NewJsonReader(r.Body).Node()).LastErr
		case "POST":
			if meta.IsAction(sel.Meta()) {
				a := sel.Meta().(*meta.Rpc)
				var input node.Node
				if a.Input != nil {
					input = node.NewJsonReader(r.Body).Node()
				}
				if outputSel := sel.ActionCntx(c, input); !outputSel.IsNil() && a.Output != nil {
					w.Header().Set("Content-Type", mime.TypeByExtension(".json"))
					err = outputSel.InsertIntoCntx(c, node.NewJsonWriter(w).Node()).LastErr
				} else {
					err = outputSel.LastErr
				}
			} else {
				payload = node.NewJsonReader(r.Body).Node()
				err = sel.InsertFromCntx(c, payload).LastErr
			}
		case "OPTIONS":
			// NOP
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	} else {
		err = sel.LastErr
	}

	if err != nil {
		HandleError(err, w)
	}
}
