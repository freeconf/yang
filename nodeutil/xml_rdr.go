package nodeutil

import (
	"encoding/xml"
	"io"
	"strings"

	"github.com/freeconf/yang/node"
)

type XmlRdr struct {
	Rdr
}

func ReadXMLIO(rdr io.Reader) node.Node {
	reader := &XmlRdr{}
	reader.In = rdr
	return reader.Node()
}

func ReadXMLValues(values map[string]interface{}) node.Node {
	reader := &XmlRdr{}
	reader.values = values
	return reader.Node()
}
func ReadXML(data string) node.Node {
	rdr := &XmlRdr{}
	rdr.In = strings.NewReader(data)
	return rdr.Node()
}

func (self *XmlRdr) decode() (map[string]interface{}, error) {
	if self.values == nil {
		d := xml.NewDecoder(self.In)
		if err := d.Decode(&self.values); err != nil {
			return nil, err
		}
	}
	return self.values, nil
}
