package nodes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
)

type JSONRdr struct {
	In     io.Reader
	values map[string]interface{}
}

func ReadJSONIO(rdr io.Reader) node.Node {
	jrdr := &JSONRdr{In: rdr}
	return jrdr.Node()
}

func ReadJSON(data string) node.Node {
	rdr := &JSONRdr{In: strings.NewReader(data)}
	return rdr.Node()
}

func (self *JSONRdr) Node() node.Node {
	var err error
	if self.values == nil {
		self.values, err = self.decode()
		if err != nil {
			return node.ErrorNode{Err: err}
		}
	}
	return JsonContainerReader(self.values)
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

func leafOrLeafListJsonReader(m meta.HasDataType, data interface{}) (v val.Value, err error) {
	return node.NewValue(m.GetDataType(), data)
}

func JsonListReader(list []interface{}) node.Node {
	s := &Basic{}
	s.OnNext = func(r node.ListRequest) (next node.Node, key []val.Value, err error) {
		key = r.Key
		if r.New {
			panic("Cannot write to JSON reader")
		}
		if len(r.Key) > 0 {
			if r.First {
				keyFields := r.Meta.Key
				for i := 0; i < len(list); i++ {
					candidate := list[i].(map[string]interface{})
					if jsonKeyMatches(keyFields, candidate, key) {
						return JsonContainerReader(candidate), r.Key, nil
					}
				}
			}
		} else {
			if r.Row < len(list) {
				container := list[r.Row].(map[string]interface{})
				if len(r.Meta.Key) > 0 {
					// TODO: compound keys
					if keyData, hasKey := container[r.Meta.Key[0]]; hasKey {
						// Key may legitimately not exist when inserting new data
						if key, err = node.NewValues(r.Meta.KeyMeta(), keyData); err != nil {
							return nil, nil, err
						}
					}
				}
				return JsonContainerReader(container), key, nil
			}
		}
		return nil, nil, nil
	}
	return s
}

func JsonContainerReader(container map[string]interface{}) node.Node {
	s := &Basic{}
	var divertedList node.Node
	s.OnChoose = func(state node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
		// go thru each case and if there are any properties in the data that are not
		// part of the meta, that disqualifies that case and we move onto next case
		// until one case aligns with data.  If no cases align then input in inconclusive
		// i.e. non-discriminating and we should error out.
		cases := meta.ChildrenNoResolve(choice)
		for cases.HasNext() {
			mkase, err := cases.Next()
			if err != nil {
				return nil, err
			}
			kase := mkase.(*meta.ChoiceCase)
			props := meta.Children(kase)
			for props.HasNext() {
				prop, err := props.Next()
				if err != nil {
					return nil, err
				}
				if _, found := container[prop.GetIdent()]; found {
					return kase, nil
				}
				// just because you didn't find a property doesnt
				// mean it's invalid, it's only if you don't find any
				// of the properties of a case
			}
		}
		msg := fmt.Sprintf("No discriminating data for choice meta %s ", state.Path)
		return nil, c2.NewErrC(msg, 400)
	}
	s.OnChild = func(r node.ChildRequest) (child node.Node, e error) {
		if r.New {
			panic("Cannot write to JSON reader")
		}
		if value, found := container[r.Meta.GetIdent()]; found {
			if meta.IsList(r.Meta) {
				return JsonListReader(value.([]interface{})), nil
			} else {
				return JsonContainerReader(value.(map[string]interface{})), nil
			}
		}
		return
	}
	s.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		if r.Write {
			panic("Cannot write to JSON reader")
		}
		if val, found := container[r.Meta.GetIdent()]; found {
			hnd.Val, err = leafOrLeafListJsonReader(r.Meta, val)
		}
		return
	}
	s.OnNext = func(r node.ListRequest) (node.Node, []val.Value, error) {
		if divertedList != nil {
			return nil, nil, nil
		}
		// divert to list handler
		foundValues, found := container[r.Meta.GetIdent()]
		list, ok := foundValues.([]interface{})
		if len(container) != 1 || !found || !ok {
			msg := fmt.Sprintf("Expected { %s: [] }", r.Meta.GetIdent())
			return nil, nil, errors.New(msg)
		}
		divertedList = JsonListReader(list)
		s.OnNext = divertedList.Next
		return divertedList.Next(r)
	}
	return s
}

func jsonKeyMatches(keyFields []string, candidate map[string]interface{}, key []val.Value) bool {
	for i, field := range keyFields {
		if candidate[field] != key[i].String() {
			return false
		}
	}
	return true
}
