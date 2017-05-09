package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
)

type JsonReader struct {
	In     io.Reader
	values map[string]interface{}
}

func NewJsonReader(in io.Reader) *JsonReader {
	r := &JsonReader{In: in}
	return r
}

func ReadJson(data string) Node {
	return NewJsonReader(strings.NewReader(data)).Node()
}

func (self *JsonReader) Node() Node {
	var err error
	if self.values == nil {
		self.values, err = self.decode()
		if err != nil {
			return ErrorNode{Err: err}
		}
	}
	return JsonContainerReader(self.values)
}

func (self *JsonReader) decode() (map[string]interface{}, error) {
	if self.values == nil {
		d := json.NewDecoder(self.In)
		if err := d.Decode(&self.values); err != nil {
			return nil, err
		}
	}
	return self.values, nil
}

func leafOrLeafListJsonReader(m meta.HasDataType, data interface{}) (v *Value, err error) {
	// TODO: Consider using CoerseValue
	v = &Value{Type: m.GetDataType()}
	switch v.Type.Format() {
	case meta.FMT_INT64:
		valF64, err := asFloat64(data)
		if err != nil {
			return nil, err
		}
		v.Int64 = int64(valF64)
	case meta.FMT_UINT64:
		valF64, err := asFloat64(data)
		if err != nil {
			return nil, err
		}
		v.UInt64 = uint64(valF64)
	case meta.FMT_INT64_LIST:
		a := data.([]interface{})
		v.Int64list = make([]int64, len(a))
		for i, f := range a {
			valF64, err := asFloat64(f)
			if err != nil {
				return nil, err
			}
			v.Int64list[i] = int64(valF64)
		}
	case meta.FMT_INT32:
		valF64, err := asFloat64(data)
		if err != nil {
			return nil, err
		}
		v.Int = int(valF64)
	case meta.FMT_UINT32:
		valF64, err := asFloat64(data)
		if err != nil {
			return nil, err
		}
		v.UInt = uint(valF64)
	case meta.FMT_INT32_LIST:
		a := data.([]interface{})
		v.Intlist = make([]int, len(a))
		for i, f := range a {
			valF64, err := asFloat64(f)
			if err != nil {
				return nil, err
			}
			v.Intlist[i] = int(valF64)
		}
	case meta.FMT_DECIMAL64:
		valF64, err := asFloat64(data)
		if err != nil {
			return nil, err
		}
		v.Float = valF64
	case meta.FMT_DECIMAL64_LIST:
		a := data.([]interface{})
		v.Floatlist = make([]float64, len(a))
		for i, f := range a {
			valF64, err := asFloat64(f)
			if err != nil {
				return nil, err
			}
			v.Floatlist[i] = valF64
		}
	case meta.FMT_STRING:
		switch vdata := data.(type) {
		case float64:
			// wrong format, truncating decimals as most likely mistake but
			// will not please everyone.  Get input in correct format by placing
			// quotes around data.
			v.Str = strconv.FormatFloat(vdata, 'f', 0, 64)
		case bool:
			if vdata {
				v.Str = "true"
			} else {
				v.Str = "false"
			}
		case string:
			v.Str = data.(string)
		}
	case meta.FMT_STRING_LIST:
		v.Strlist = asStringArray(data.([]interface{}))
	case meta.FMT_BOOLEAN:
		switch vdata := data.(type) {
		case string:
			s := data.(string)
			v.Bool = ("true" == s)
		case bool:
			v.Bool = vdata
		}
	case meta.FMT_BOOLEAN_LIST:
		a := data.([]interface{})
		v.Boollist = make([]bool, len(a))
		for i, data := range a {
			switch vdata := data.(type) {
			case string:
				s := data.(string)
				v.Boollist[i] = ("true" == s)
			case bool:
				v.Boollist[i] = vdata
			}
		}
	case meta.FMT_ENUMERATION:
		v.SetEnumByLabel(data.(string))
	case meta.FMT_ENUMERATION_LIST:
		strlist := InterfaceToStrlist(data)
		if len(strlist) > 0 {
			v.SetEnumListByLabels(strlist)
		} else {
			intlist := InterfaceToIntlist(data)
			v.SetEnumList(intlist)
		}
	case meta.FMT_ANYDATA:
		v.AnyData = data
	default:
		msg := fmt.Sprint("JSON reading value type not implemented ", m.GetDataType().Format())
		return nil, errors.New(msg)
	}
	return
}

func asFloat64(data interface{}) (val float64, err error) {
	var valF64 float64
	valF64, ok := data.(float64)
	if !ok {
		valStr, ok := data.(string)
		if !ok {
			msg := fmt.Sprint("JSON reading value could not parse %v as int64", data)
			return 0.0, errors.New(msg)
		}
		valF64, err = strconv.ParseFloat(valStr, 64)
		if err != nil {
			msg := fmt.Sprint("JSON reading value could not parse %v as int64: %s", data, err.Error())
			return 0.0, errors.New(msg)
		}
	}

	return valF64, nil
}

func asStringArray(data []interface{}) []string {
	s := make([]string, len(data))
	for i, d := range data {
		s[i] = d.(string)
	}
	return s
}

func JsonListReader(list []interface{}) Node {
	s := &MyNode{Label: "JSON Read List"}
	s.OnNext = func(r ListRequest) (next Node, key []*Value, err error) {
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
						key = SetValues(r.Meta.KeyMeta(), keyData)
					}
				}
				return JsonContainerReader(container), key, nil
			}
		}
		return nil, nil, nil
	}
	return s
}

func JsonContainerReader(container map[string]interface{}) Node {
	s := &MyNode{Label: "JSON Read Container"}
	var divertedList Node
	s.OnChoose = func(state Selection, choice *meta.Choice) (m *meta.ChoiceCase, err error) {
		// go thru each case and if there are any properties in the data that are not
		// part of the meta, that disqualifies that case and we move onto next case
		// until one case aligns with data.  If no cases align then input in inconclusive
		// i.e. non-discriminating and we should error out.
		cases := meta.NewMetaListIterator(choice, false)
		for cases.HasNextMeta() {
			kase := cases.NextMeta().(*meta.ChoiceCase)
			props := meta.NewMetaListIterator(kase, true)
			for props.HasNextMeta() {
				prop := props.NextMeta()
				if _, found := container[prop.GetIdent()]; found {
					return kase, nil
				}
				// just because you didn't find a property doesnt
				// mean it's invalid, it's only if you don't find any
				// of the properties of a case
			}
		}
		msg := fmt.Sprintf("No discriminating data for choice meta %s ", state.String())
		return nil, c2.NewErrC(msg, 400)
	}
	s.OnChild = func(r ChildRequest) (child Node, e error) {
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
	s.OnField = func(r FieldRequest, hnd *ValueHandle) (err error) {
		if r.Write {
			panic("Cannot write to JSON reader")
		}
		if val, found := container[r.Meta.GetIdent()]; found {
			hnd.Val, err = leafOrLeafListJsonReader(r.Meta, val)
		}
		return
	}
	s.OnNext = func(r ListRequest) (Node, []*Value, error) {
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

func jsonKeyMatches(keyFields []string, candidate map[string]interface{}, key []*Value) bool {
	for i, field := range keyFields {
		if candidate[field] != key[i].String() {
			return false
		}
	}
	return true
}
