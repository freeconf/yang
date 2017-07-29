package node

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
)

// Value is a union of all possible values a leaf could hold.
type Value struct {
	Format     meta.DataFormat
	Bool       bool
	Int        int
	UInt       uint
	Int64      int64
	UInt64     uint64
	Int64list  []int64
	UInt64list []uint64
	Str        string
	Float      float64
	Intlist    []int
	UIntlist   []uint
	Strlist    []string
	Boollist   []bool
	Floatlist  []float64
	Keys       []string
	AnyData    interface{}
}

func EncodeKey(v []*Value) string {
	var s string
	// TODO: read RFC and escape chars including commas
	for i, val := range v {
		if i > 0 {
			s = "," + val.String()
		}
		s = val.String()
	}
	return s
}

func (v *Value) Value() interface{} {
	switch v.Format {
	case meta.FMT_BOOLEAN:
		return v.Bool
	case meta.FMT_BOOLEAN_LIST:
		return v.Boollist
	case meta.FMT_INT64:
		return v.Int64
	case meta.FMT_UINT64:
		return v.UInt64
	case meta.FMT_INT64_LIST:
		return v.Int64list
	case meta.FMT_UINT64_LIST:
		return v.UInt64list
	case meta.FMT_UINT32:
		return v.UInt
	case meta.FMT_INT32, meta.FMT_ENUMERATION:
		return v.Int
	case meta.FMT_INT32_LIST, meta.FMT_ENUMERATION_LIST:
		return v.Intlist
	case meta.FMT_UINT32_LIST:
		return v.UIntlist
	case meta.FMT_STRING:
		return v.Str
	case meta.FMT_STRING_LIST:
		return v.Strlist
	case meta.FMT_ANYDATA:
		return v.AnyData
	case meta.FMT_DECIMAL64:
		return v.Float
	case meta.FMT_DECIMAL64_LIST:
		return v.Floatlist
	default:
		panic("Value() not implemented for format " + v.Format.String())
	}
}

func (a *Value) Compare(b *Value) int {
	if a == nil {
		if b == nil {
			return 0
		}
		return 1
	}
	if b == nil {
		return -1
	}
	if a.Format != b.Format {
		return int(a.Format) - int(b.Format)
	}
	if meta.IsListFormat(a.Format) {
		panic("not sure how to compare arrays")
	}
	switch a.Format {
	case meta.FMT_BOOLEAN:
		if a.Bool == b.Bool {
			return 0
		}
		if a.Bool {
			return 1
		}
		return -1
	case meta.FMT_INT64:
		x := a.Int64 - b.Int64
		// we do the "long" form of comparision to avoid numeric overflow
		// where positive int64 > 2 billion is negative int32 when cast
		if x < 0 {
			return -1
		}
		if x > 0 {
			return 1
		}
		return 0
	case meta.FMT_UINT64:
		x := a.UInt64 - b.UInt64
		if x < 0 {
			return -1
		}
		if x > 0 {
			return 1
		}
		return 0
	case meta.FMT_UINT32:
		x := a.UInt - b.UInt
		if x < 0 {
			return -1
		}
		if x > 0 {
			return 1
		}
		return 0
	case meta.FMT_INT32, meta.FMT_ENUMERATION:
		return a.Int - b.Int
	case meta.FMT_STRING:
		return strings.Compare(a.Str, b.Str)
	case meta.FMT_DECIMAL64:
		x := a.Float - b.Float
		if x < 0 {
			return -1
		}
		if x > 0 {
			return 1
		}
		return 0
	default:
		panic("Not implemented")
	}
}

func (a *Value) Equal(b *Value) bool {
	if a == nil {
		if b == nil {
			return true
		}
		return false
	}
	if b == nil {
		return false
	}
	if a.Format != b.Format {
		return false
	}
	if meta.IsListFormat(a.Format) {
		return reflect.DeepEqual(a.Value(), b.Value())
	}
	return a.Value() == b.Value()
}

func NewEnumList(en meta.Enum, intlist []int) (*Value, error) {
	v := &Value{
		Format:  meta.FMT_ENUMERATION_LIST,
		Strlist: make([]string, len(intlist)),
		Intlist: intlist,
	}
	for i, n := range intlist {
		if found := en.ByValue(n); found.Nil() {
			return nil, c2.NewErr(fmt.Sprintf("%d not legal value of enumeration", n))
		} else {
			v.Strlist[i] = en[n].Label
		}
	}
	return v, nil
}

func NewEnumValue(en meta.Enum, n int) (*Value, error) {
	v := &Value{
		Format: meta.FMT_ENUMERATION,
		Int:    n,
	}
	if found := en.ByValue(n); found.Nil() {
		return nil, c2.NewErr(fmt.Sprintf("%d not legal value of enumeration", n))
	} else {
		v.Str = found.Label
	}
	return v, nil
}

func NewEnumListByLabels(en meta.Enum, labels []string) (*Value, error) {
	v := &Value{
		Format:  meta.FMT_ENUMERATION_LIST,
		Intlist: make([]int, len(labels)),
		Strlist: labels,
	}
	for i, s := range labels {
		if found := en.ByLabel(s); found.Nil() {
			return nil, c2.NewErr(fmt.Sprintf("%s not legal value of enumeration", s))
		} else {
			v.Intlist[i] = found.Value
		}
	}
	return v, nil
}

func NewEnumByLabel(en meta.Enum, label string) (*Value, error) {
	v := &Value{
		Format: meta.FMT_ENUMERATION,
		Str:    label,
	}
	if found := en.ByLabel(label); found.Nil() {
		return nil, c2.NewErr(fmt.Sprintf("%s not legal value of enumeration", label))
	} else {
		v.Int = found.Value
	}
	return v, nil
}

func (v *Value) String() string {
	return fmt.Sprintf("%v", v.Value())
}

func NewValuesByString(m []meta.HasDataType, objs ...string) ([]*Value, error) {
	var err error
	vals := make([]*Value, len(m))
	for i, obj := range objs {
		vals[i], err = NewValue(m[i].GetDataType(), obj)
		if err != nil {
			return nil, err
		}
	}
	return vals, nil
}

func NewValues(m []meta.HasDataType, objs ...interface{}) ([]*Value, error) {
	var err error
	vals := make([]*Value, len(m))
	for i, obj := range objs {
		vals[i], err = NewValue(m[i].GetDataType(), obj)
		if err != nil {
			return nil, err
		}
	}
	return vals, nil
}

// Incoming value should be of appropriate type according to given data type format
func NewValue(typ *meta.DataType, val interface{}) (*Value, error) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf("%s : %s", typ.Ident, r))
		}
	}()
	if val == nil {
		return nil, nil
	}
	i, err := typ.Info()
	if err != nil {
		return nil, err
	}
	reflectVal := reflect.ValueOf(val)
	v := &Value{Format: i.Format}
	switch i.Format {
	case meta.FMT_BOOLEAN:
		v.Bool = reflectVal.Bool()
	case meta.FMT_BOOLEAN_LIST:
		v.Boollist = InterfaceToBoollist(reflectVal.Interface())
	case meta.FMT_INT32_LIST:
		v.Intlist = InterfaceToIntlist(val)
	case meta.FMT_INT32:
		switch reflectVal.Kind() {
		case reflect.String:
			i64, err := strconv.ParseInt(reflectVal.String(), 10, 32)
			if err != nil {
				return nil, err
			}
			v.Int = int(i64)
		// special case float mostly because of JSON
		case reflect.Float64:
			v.Int = int(reflectVal.Float())
		default:
			v.Int = int(reflectVal.Int())
		}
	case meta.FMT_UINT32:
		switch reflectVal.Kind() {
		case reflect.Float64:
			v.UInt = uint(reflectVal.Float())
		default:
			v.UInt = reflectVal.Interface().(uint)
		}
	case meta.FMT_DECIMAL64:
		v.Float = reflectVal.Float()
	case meta.FMT_UINT64:
		v.UInt64 = reflectVal.Interface().(uint64)
	case meta.FMT_INT64:
		switch reflectVal.Kind() {
		// special case float mostly because of JSON
		case reflect.Float64:
			v.Int64 = int64(reflectVal.Float())
		default:
			v.Int64 = reflectVal.Int()
		}
	case meta.FMT_INT64_LIST:
		v.Int64list = InterfaceToInt64list(val)
	case meta.FMT_STRING:
		v.Str = reflectVal.String()
	case meta.FMT_ENUMERATION:
		switch reflectVal.Kind() {
		case reflect.String:
			v, err = NewEnumByLabel(i.Enum, reflectVal.String())
		case reflect.Float64:
			v, err = NewEnumValue(i.Enum, int(reflectVal.Float()))
		default:
			v, err = NewEnumValue(i.Enum, int(reflectVal.Int()))
		}
	case meta.FMT_ENUMERATION_LIST:
		val := reflectVal.Interface()
		strlist := InterfaceToStrlist(val)
		if len(strlist) > 0 {
			v, err = NewEnumListByLabels(i.Enum, strlist)
		} else {
			intlist := InterfaceToIntlist(val)
			v, err = NewEnumList(i.Enum, intlist)
		}
	case meta.FMT_STRING_LIST:
		v.Strlist = InterfaceToStrlist(reflectVal.Interface())
	case meta.FMT_ANYDATA:
		v.AnyData = reflectVal.Interface()
	default:
		panic(fmt.Sprintf("Format code %s not implemented", i.Format))
	}
	return v, err
}

func InterfaceToStrlist(o interface{}) (strlist []string) {
	switch arrayValues := o.(type) {
	case []string:
		return arrayValues
	case []interface{}:
		found := make([]string, len(arrayValues))
		for i, arrayValue := range arrayValues {
			if reflect.TypeOf(arrayValue).Kind() != reflect.String {
				return
			}
			found[i] = arrayValue.(string)
		}
		strlist = found
	}
	return
}

func InterfaceToBoollist(o interface{}) (boollist []bool) {
	switch arrayValues := o.(type) {
	case []bool:
		return arrayValues
	case []interface{}:
		boollist = make([]bool, len(arrayValues))
		for i, arrayValue := range arrayValues {
			boollist[i] = arrayValue.(bool)
		}
	}
	return
}

func InterfaceToInt64list(o interface{}) (intlist []int64) {
	switch arrayValues := o.(type) {
	case []int64:
		return arrayValues
	case []interface{}:
		intlist = make([]int64, len(arrayValues))
		for i, arrayValue := range arrayValues {
			intlist[i] = arrayValue.(int64)
		}
	}
	return
}

func InterfaceToIntlist(o interface{}) (intlist []int) {
	switch arrayValues := o.(type) {
	case []int:
		return arrayValues
	case []interface{}:
		intlist = make([]int, len(arrayValues))
		for i, arrayValue := range arrayValues {
			switch n := arrayValue.(type) {
			case int:
				intlist[i] = n
			case float32:
				intlist[i] = int(n)
			case float64:
				intlist[i] = int(n)
			case int64:
				intlist[i] = int(n)
			}
		}
	}
	return
}

// func (v *Value) CoerseStrValue(s string) error {
// 	switch v.Format {
// 	case meta.FMT_BOOLEAN:
// 		v.Bool = s == "true"
// 	case meta.FMT_UINT64:
// 		var err error
// 		v.UInt64, err = strconv.ParseUint(s, 10, 64)
// 		if err != nil {
// 			return err
// 		}
// 	case meta.FMT_INT64:
// 		var err error
// 		v.Int64, err = strconv.ParseInt(s, 10, 64)
// 		if err != nil {
// 			return err
// 		}
// 	case meta.FMT_UINT32:
// 		i64, err := strconv.ParseUint(s, 10, 32)
// 		if err != nil {
// 			return err
// 		}
// 		v.UInt = uint(i64)
// 	case meta.FMT_INT32:
// 		var err error
// 		v.Int, err = strconv.Atoi(s)
// 		if err != nil {
// 			return err
// 		}
// 	case meta.FMT_DECIMAL64:
// 		var err error
// 		v.Float, err = strconv.ParseFloat(s, 64)
// 		if err != nil {
// 			return err
// 		}
// 	case meta.FMT_STRING:
// 		v.Str = s
// 	case meta.FMT_ENUMERATION:
// 		eid, err := strconv.Atoi(s)
// 		if err == nil {
// 			v.SetEnum(eid)
// 			return nil
// 		}
// 		if !v.SetEnumByLabel(s) {
// 			return c2.NewErr("Not an allowed enumation: " + s)
// 		}
// 	default:
// 		panic(fmt.Sprintf("Coersion not supported from data format " + v.Type.Format().String()))
// 	}
// 	return nil
// }
