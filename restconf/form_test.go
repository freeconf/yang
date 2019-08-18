package restconf

import (
	"os"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"testing"

	"github.com/freeconf/gconf/meta/yang"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/nodes"
)

type handlerImpl http.HandlerFunc

func (impl handlerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	impl(w, r)
}

func TestForm(t *testing.T) {
	m := yang.RequireModuleFromString(nil, `
		module test {
			rpc x {
				input {
					leaf a {
						type string;
					}
					anydata b;					
				}
			}
		}
	`)
	done := make(chan bool, 2)
	handler := func(w http.ResponseWriter, r *http.Request) {
		b := node.NewBrowser(m, formDummyNode(t))
		input, err := requestNode(r)
		chkErr(t, err)
		resp := b.Root().Find("x").Action(input)
		chkErr(t, resp.LastErr)
		w.Write([]byte("ok"))
		done <- true
	}
	srv := &http.Server{Addr: "0.0.0.0:9999", Handler: handlerImpl(handler)}
	go func() {
		srv.ListenAndServe()
	}()
	post(t)
	<-done
	srv.Shutdown(context.TODO())
}

func chkErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func post(t *testing.T) {
	if "true" == os.Getenv("TRAVIS") {
		t.Skip()
		return
	}
	rdr, wtr := io.Pipe()
	wait := make(chan bool, 2)
	form := multipart.NewWriter(wtr)
	go func() {
		req, err := http.NewRequest("POST", "http://127.0.0.1:9999", rdr)
		chkErr(t, err)
		req.Header.Set("Content-Type", form.FormDataContentType())
		_, err = http.DefaultClient.Do(req)
		wait <- true
	}()
	dataPart, err := form.CreateFormField("a")
	chkErr(t, err)
	_, err = io.Copy(dataPart, strings.NewReader("hello"))
	chkErr(t, err)
	filePart, err := form.CreateFormFile("b", "b")
	chkErr(t, err)
	_, err = io.Copy(filePart, strings.NewReader("hello world"))
	chkErr(t, err)
	chkErr(t, form.Close())
	chkErr(t, wtr.Close())
	<-wait
}

func formDummyNode(t *testing.T) node.Node {
	return &nodes.Basic{
		OnAction: func(r node.ActionRequest) (node.Node, error) {
			v, err := r.Input.GetValue("a")
			chkErr(t, err)
			if v.String() != "hello" {
				t.Error(v.String())
			}

			v, err = r.Input.GetValue("b")
			chkErr(t, err)
			rdr, valid := v.Value().(io.ReadCloser)
			if !valid {
				panic("invalid")
			}
			actual, err := ioutil.ReadAll(rdr)
			chkErr(t, err)
			if string(actual) != "hello world" {
				t.Error(actual)
			}
			defer rdr.Close()
			fmt.Printf(string(actual))
			return nil, nil
		},
	}
}
