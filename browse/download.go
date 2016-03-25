package browse

import (
	"node"
	"meta"
	"net/http"
	"blit"
)

func DownloadMeta(url string, dest meta.MetaList) (error) {
	in, err := DownloadJson(url)
	if err != nil {
		return err
	}
	c := node.NewContext()
	yangModule := node.YangModule()
	var resolve bool
	var goober meta.MetaList
	if dest, isModule := dest.(*meta.Module); isModule {
		resolve = false
		goober = yangModule
	} else {
		resolve = true
		goober = meta.FindByPath(yangModule, "module/definitions").(meta.MetaList)
		if meta.IsList(dest) {
			goober = meta.FindByIdentExpandChoices(goober, "list").(meta.MetaList)
		} else {
			goober = meta.FindByIdentExpandChoices(goober, "container").(meta.MetaList)
		}
	}
	destNode := node.SchemaData{Resolve:resolve}.MetaList(dest)
	if err = c.Select(goober.(meta.MetaList), destNode).UpsertFrom(in).LastErr; err != nil {
		return err
	}
	return err
}

func DownloadJson(url string) (n node.Node, err error) {
	var req *http.Request
	blit.Info.Printf("Downloading meta %s", url)
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	client := http.DefaultClient
	resp, getErr := client.Do(req)
	if getErr != nil {
		return nil, getErr
	}
	//defer resp.Body.Close()
	return node.NewJsonReader(resp.Body).Node(), nil
}