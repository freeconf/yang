package restconf

import (
	"errors"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/freeconf/gconf/meta"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/nodes"
	"github.com/freeconf/gconf/val"
)

func isMultiPartForm(hdrs http.Header) bool {
	return strings.HasPrefix(hdrs.Get("Content-Type"), "multipart/form-data")
}

func formNode(req *http.Request) (node.Node, error) {
	err := req.ParseMultipartForm(10000)
	if err != nil {
		return nil, err
	}
	return &nodes.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			entry, found := req.MultipartForm.File[r.Meta.Ident()]
			if !found || len(entry) == 0 {
				return nil, nil
			}
			if meta.IsList(r.Meta) {
				return formListNode(entry), nil
			}
			if len(entry) != 1 {
				return nil, errors.New("invalid number of form files for structure, expected 0 or 1")
			}

			return nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			sval := req.FormValue(r.Meta.Ident())
			if sval != "" {
				var err error
				hnd.Val, err = node.NewValue(r.Meta.Type(), sval)
				return err
			}
			entry, found := req.MultipartForm.File[r.Meta.Ident()]
			if found {
				if len(entry) == 0 {
					return nil
				}

				// Can type any be a leaf-list? spec says yes
				// if r.Meta.Type().Format().IsList()

				if len(entry) != 1 {
					return errors.New("invalid number of form files for field, expected 0 or 1")
				}
				f, err := entry[0].Open()
				if err != nil {
					return err
				}
				hnd.Val = val.Any{Thing: f}
				return nil
			}

			return nil
		},
	}, nil
}

func formChildNode(f *multipart.FileHeader) (node.Node, error) {
	rdr, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rdr.Close()
	return nodes.ReadJSONIO(rdr), nil
}

func formListNode(files []*multipart.FileHeader) node.Node {
	return &nodes.Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			if r.Row >= len(files) {
				return nil, nil, nil
			}
			n, err := formChildNode(files[r.Row])
			return n, nil, err
		},
	}
}
