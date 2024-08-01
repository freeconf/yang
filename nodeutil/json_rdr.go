package nodeutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/val"

	"github.com/freeconf/yang/meta"
)

type JSONRdr struct {
	In     io.Reader
	values map[string]interface{}
}

func ReadJSONIO(rdr io.Reader) (node.Node, error) {
	jrdr := &JSONRdr{In: rdr}
	return jrdr.Node()
}

func ReadJSONValues(values map[string]interface{}) (node.Node, error) {
	jrdr := &JSONRdr{values: values}
	return jrdr.Node()
}

func ReadJSON(data string) (node.Node, error) {
	rdr := &JSONRdr{In: strings.NewReader(data)}
	return rdr.Node()
}

func (self *JSONRdr) Node() (node.Node, error) {
	var err error
	if self.values == nil {
		self.values, err = self.decode()
		if err != nil {
			return nil, err
		}
	}
	return JsonContainerReader(self.values), nil
}

// This function inspects the JSON payload against YANG schema,
// looking for missing or unexpected keys.
func (self *JSONRdr) Validate(selection *node.Selection) error {
	n, err := self.Node()
	if err != nil {
		return err
	}

	err = validate(self.values, selection.Split(n))
	if err != nil {
		return fmt.Errorf("%w: %s", fc.BadRequestError, err)
	}

	return nil
}

func validate(value interface{}, selection *node.Selection) error {
	m := selection.Meta()
	if meta.IsContainer(m) {
		containerValue, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("expected a container, got: %+v", value)
		}
		return validateContainer(containerValue, selection)
	} else if meta.IsList(m) {
		listValue, ok := value.([]interface{})
		if !ok {
			return fmt.Errorf("expected a list, got: %+v", value)
		}
		return validateList(listValue, selection)
	} else {
		// TODO: no validation for leaves, choices, and the rest
	}

	return nil
}

func validateList(elements []interface{}, selection *node.Selection) error {
	elementSelection, err := selection.First()

	for i := range elements {
		for err != nil {
			return fmt.Errorf("error selecting list element %d: %s", i, err)
		}
		path := elementSelection.Selection.Path.String()

		element, ok := elements[i].(map[string]interface{})
		if !ok {
			return fmt.Errorf("expected a map for path %s, got %+v", path, elements[i])
		}
		if err := validateChildNodes(element, elementSelection.Selection); err != nil {
			return err
		}
		elementSelection, err = elementSelection.Next()
	}

	return nil
}

func validateChildNodes(values map[string]interface{}, selection *node.Selection) error {
	m := selection.Meta()
	path := selection.Path.String()
	metaChildren := map[string]struct{}{}
	hd := m.(meta.HasDataDefinitions)

	for _, child := range hd.DataDefinitions() {
		id := child.Ident()
		metaChildren[id] = struct{}{}
		details := child.(meta.HasDetails)

		value, ok := values[id]
		if !ok {
			if details.Mandatory() {
				return fmt.Errorf("missing mandatory node: %s/%s", path, id)
			}
		} else {
			newSelection, err := selection.Find(id)
			if err != nil {
				return fmt.Errorf("error finding: %s/%s: %s", path, id, err)
			}
			if err := validate(value, newSelection); err != nil {
				return err
			}
		}
	}

	for k := range values {
		if _, ok := metaChildren[k]; !ok {
			return fmt.Errorf("unexpected node: %s/%s", path, k)
		}
	}

	return nil
}

func validateContainer(values map[string]interface{}, selection *node.Selection) error {
	return validateChildNodes(values, selection)
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

func leafOrLeafListJsonReader(m meta.Leafable, data interface{}) (v val.Value, err error) {
	return node.NewValue(m.Type(), data)
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
				keyFields := r.Meta.KeyMeta()
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
				if len(r.Meta.KeyMeta()) > 0 {
					keyData := make([]interface{}, len(r.Meta.KeyMeta()))
					for i, kmeta := range r.Meta.KeyMeta() {
						// Key may legitimately not exist when inserting new data
						keyData[i] = fqkGetOrNil(kmeta, container)
					}
					if key, err = node.NewValues(r.Meta.KeyMeta(), keyData...); err != nil {
						return nil, nil, err
					}
				}
				return JsonContainerReader(container), key, nil
			}
		}
		return nil, nil, nil
	}
	return s
}

func fqkGetOrNil(m meta.Definition, container map[string]interface{}) interface{} {
	v, _ := fqkGet(m, container)
	return v
}

func fqkGet(m meta.Definition, container map[string]interface{}) (interface{}, bool) {
	v, found := container[m.Ident()]
	if !found {
		mod := meta.OriginalModule(m)
		v, found = container[fmt.Sprintf("%s:%s", mod.Ident(), m.Ident())]
	}
	return v, found
}

func JsonContainerReader(container map[string]interface{}) node.Node {
	s := &Basic{}
	var divertedList node.Node
	s.OnChoose = func(state *node.Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
		// go thru each case and if there are any properties in the data that are not
		// part of the meta, that disqualifies that case and we move onto next case
		// until one case aligns with data.  If no cases align then input in inconclusive
		// i.e. non-discriminating and we should error out.
		for _, kase := range choice.Cases() {
			for _, prop := range kase.DataDefinitions() {
				if _, found := fqkGet(prop, container); found {
					return kase, nil
				}
				// just because you didn't find a property doesnt
				// mean it's invalid, it's only if you don't find any
				// of the properties of a case
			}
		}
		// just because you didn't find any properties of any cases doesn't
		// mean it's invalid, just that *none* of the cases are there.
		return nil, nil
	}
	s.OnChild = func(r node.ChildRequest) (child node.Node, e error) {
		if r.New {
			panic("cannot write to JSON reader")
		}
		if value, found := fqkGet(r.Meta, container); found {
			if meta.IsList(r.Meta) {
				return JsonListReader(value.([]interface{})), nil
			}
			return JsonContainerReader(value.(map[string]interface{})), nil
		}
		return
	}
	s.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		if r.Write {
			panic("cannot write to JSON reader")
		}
		if val, found := fqkGet(r.Meta, container); found {
			hnd.Val, err = leafOrLeafListJsonReader(r.Meta, val)
		}
		return
	}
	s.OnNext = func(r node.ListRequest) (node.Node, []val.Value, error) {
		if divertedList != nil {
			return nil, nil, nil
		}
		// divert to list handler
		foundValues, found := fqkGet(r.Meta, container)
		list, ok := foundValues.([]interface{})
		if len(container) != 1 || !found || !ok {
			msg := fmt.Sprintf("Expected { %s: [] }", r.Meta.Ident())
			return nil, nil, errors.New(msg)
		}
		divertedList = JsonListReader(list)
		s.OnNext = divertedList.Next
		return divertedList.Next(r)
	}
	return s
}

func jsonKeyMatches(keyFields []meta.Leafable, candidate map[string]interface{}, key []val.Value) bool {
	for i, field := range keyFields {
		if fqkGetOrNil(field, candidate) != key[i].String() {
			return false
		}
	}
	return true
}
