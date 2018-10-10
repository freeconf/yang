package restconf

import (
	"fmt"
	"mime"
	"net/http"

	"context"

	"github.com/freeconf/gconf/device"
	"github.com/freeconf/gconf/meta"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/nodes"
)

type browserHandler struct {
	browser  *node.Browser
	subCount int
}

var subscribeCount int

func (self *browserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var payload node.Node
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if r.RemoteAddr != "" {
		host, _ := ipAddrSplitHostPort(r.RemoteAddr)
		ctx = context.WithValue(ctx, device.RemoteIpAddressKey, host)
	}
	sel := self.browser.RootWithContext(ctx)
	if sel = sel.FindUrl(r.URL); sel.LastErr == nil {
		hdr := w.Header()
		if sel.IsNil() {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		if handleErr(err, w) {
			return
		}
		switch r.Method {
		case "DELETE":
			// CRUD - Delete
			err = sel.Delete()
		case "GET":

			// compliance note : decided to support notifictions on get by devilering
			// first event, then closing connection.  Spec calls for SSE
			if meta.IsNotification(sel.Meta()) {

				// Event

				hdr.Set("Content-Type", "text/event-stream")
				hdr.Set("Cache-Control", "no-cache")
				hdr.Set("Connection", "keep-alive")
				hdr.Set("Access-Control-Allow-Origin", "*")
				// default is chunked and web browsers don't know to read after each
				// flush
				hdr.Set("Transfer-Encoding", "identity")

				var sub node.NotifyCloser
				flusher, hasFlusher := w.(http.Flusher)
				if !hasFlusher {
					panic("invalid response writer")
				}
				closer, hasCloser := w.(http.CloseNotifier)
				if !hasCloser {
					panic("invalid response writer")
				}
				subscribeCount++
				defer func() {
					subscribeCount--
				}()
				sub, err := sel.Notifications(func(msg node.Selection) {

					// According to SSE Spec, each event needs following format:
					// data: {payload}\n\n

					fmt.Fprint(w, "data: ")
					jout := &nodes.JSONWtr{Out: w}
					if err := msg.InsertInto(jout.Node()).LastErr; err != nil {
						handleErr(err, w)
						return
					}
					fmt.Fprint(w, "\n\n")
					flusher.Flush()
				})
				if err != nil {
					handleErr(err, w)
					return
				}
				defer sub()
				<-closer.CloseNotify()
			} else {

				// CRUD - Read

				hdr.Set("Content-Type", mime.TypeByExtension(".json"))
				jout := &nodes.JSONWtr{Out: w}
				err = sel.InsertInto(jout.Node()).LastErr
			}
		case "PUT":
			// CRUD - Update
			input, err := requestNode(r)
			if err != nil {
				handleErr(err, w)
				return
			}
			err = sel.UpsertFrom(input).LastErr
		case "POST":
			if meta.IsAction(sel.Meta()) {

				// RPC

				a := sel.Meta().(*meta.Rpc)
				var input node.Node
				if a.Input() != nil {
					if input, err = requestNode(r); err != nil {
						handleErr(err, w)
						return
					}
				}
				if outputSel := sel.Action(input); !outputSel.IsNil() && a.Output() != nil {
					hdr.Set("Content-Type", mime.TypeByExtension(".json"))
					jout := &nodes.JSONWtr{Out: w}
					err = outputSel.InsertInto(jout.Node()).LastErr
				} else {
					err = outputSel.LastErr
				}
			} else {

				// CRUD - Insert

				payload = nodes.ReadJSONIO(r.Body)
				err = sel.InsertFrom(payload).LastErr
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
		handleErr(err, w)
	}
}

func requestNode(r *http.Request) (node.Node, error) {
	if isMultiPartForm(r.Header) {
		return formNode(r)
	}
	return nodes.ReadJSONIO(r.Body), nil
}
