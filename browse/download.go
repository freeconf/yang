package browse

import (
	"net/http"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

type MetaResolver func(yangPath meta.StreamSource, url string, receiver meta.MetaList) error

func DownloadMeta(yangPath meta.StreamSource, url string, dest meta.MetaList) error {
	in, err := DownloadJson(url)
	if err != nil {
		return err
	}
	yangModule := yang.RequireModule(yangPath, "yang")
	var resolve bool
	var m meta.Meta
	if dest, isModule := dest.(*meta.Module); isModule {
		resolve = false
		m = yangModule
	} else {
		resolve = true
		m, err = meta.FindByPath(yangModule, "module/definitions")
		if err != nil {
			return err
		}
		if meta.IsList(dest) {
			m, err = meta.FindByIdentExpandChoices(m, "list")
			if err != nil {
				return err
			}
		} else {
			m, err = meta.FindByIdentExpandChoices(m, "container")
			if err != nil {
				return err
			}
		}
	}
	destNode := node.SchemaData{Resolve: resolve}.MetaList(dest)
	if err = node.NewBrowser(m.(meta.MetaList), destNode).Root().UpsertFrom(in).LastErr; err != nil {
		return err
	}
	return err
}

func DownloadJson(url string) (n node.Node, err error) {
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	client := http.DefaultClient
	resp, getErr := client.Do(req)
	if getErr != nil {
		return nil, getErr
	}
	defer resp.Body.Close()
	return node.NewJsonReader(resp.Body).Node(), nil
}
