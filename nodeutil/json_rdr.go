package nodeutil

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/freeconf/yang/node"
)

type JSONRdr struct {
	Rdr
}

func ReadJSONIO(rdr io.Reader) node.Node {
	jrdr := &JSONRdr{}
	jrdr.In = rdr
	return jrdr.Node()
}

func ReadJSONValues(values map[string]interface{}) node.Node {
	jrdr := &JSONRdr{}
	jrdr.values = values
	return jrdr.Node()
}

func ReadJSON(data string) node.Node {
	rdr := &JSONRdr{}
	rdr.In = strings.NewReader(data)
	return rdr.Node()
}

func (self *JSONRdr) decode() (map[string]interface{}, error) {
	if self.values == nil {
		d := json.NewDecoder(self.In)
		if err := d.Decode(&self.values); err != nil {
			return nil, err
		}
	}
	return self.values, nil
}
