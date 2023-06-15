package nodeutil

import (
	"io"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/clbanning/mxj/v2"
	"github.com/freeconf/yang/node"
)

type XmlRdr struct {
	Rdr
}

func ReadXMLIO(rdr io.Reader) node.Node {
	reader := &XmlRdr{}
	reader.decoder = reader
	reader.In = rdr
	return reader.Node()
}

func ReadXML(data string) node.Node {
	rdr := &XmlRdr{}
	rdr.decoder = rdr
	rdr.In = strings.NewReader(data)
	return rdr.Node()
}

func (self *XmlRdr) Decode() (map[string]interface{}, error) {
	if self.values == nil {
		data, err := ioutil.ReadAll(self.In)
		if err != nil {
			return nil, err
		}
		xml_map, err := mxj.NewMapXml(data)
		if err != nil {
			return nil, err
		}

		self.removeAttributesFromXmlMap(xml_map)
		self.values = xml_map
	}
	return self.values, nil
}

func (self *XmlRdr) removeAttributesFromXmlMap(m map[string]interface{}) {
	val := reflect.ValueOf(m)
	for _, e := range val.MapKeys() {
		v := val.MapIndex(e)
		if strings.Index(e.String(), "-") == 0 {
			delete(m, e.String())
			continue
		}
		switch t := v.Interface().(type) {
		case map[string]interface{}:
			self.removeAttributesFromXmlMap(t)
		}
	}
}
