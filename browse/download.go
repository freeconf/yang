package browse

import (
	"context"
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
	var m meta.MetaList
	if dest, isModule := dest.(*meta.Module); isModule {
		resolve = false
		m = yangModule
	} else {
		resolve = true
		m = meta.FindByPath(yangModule, "module/definitions").(meta.MetaList)
		if meta.IsList(dest) {
			m = meta.FindByIdentExpandChoices(m, "list").(meta.MetaList)
		} else {
			m = meta.FindByIdentExpandChoices(m, "container").(meta.MetaList)
		}
	}
	destNode := node.SchemaData{Resolve: resolve}.MetaList(dest)
	if err = node.NewBrowser(m.(meta.MetaList), destNode).Root().UpsertFrom(context.Background(), in).LastErr; err != nil {
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
