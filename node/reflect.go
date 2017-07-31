package node

import (
	"fmt"
	"reflect"

	"github.com/c2stack/c2g/val"

	"github.com/c2stack/c2g/meta"
)

func ReadField(m meta.HasDataType, obj interface{}) (val.Value, error) {
	return ReadFieldWithFieldName(meta.MetaNameToFieldName(m.GetIdent()), m, obj)
}

func ReadFieldWithFieldName(fieldName string, m meta.HasDataType, obj interface{}) (v val.Value, err error) {
	objVal := reflect.ValueOf(obj)
	if objVal.Kind() == reflect.Interface || objVal.Kind() == reflect.Ptr {
		objVal = objVal.Elem()
		if objVal.Kind() == reflect.Ptr {
			panic(fmt.Sprintf("Pointer to a pointer not legal %s on %v ", m.GetIdent(), reflect.TypeOf(obj)))
		}
	}
	value := objVal.FieldByName(fieldName)

	if !value.IsValid() {
		panic(fmt.Sprintf("Field not found: %s on %v ", m.GetIdent(), reflect.TypeOf(obj)))
	}

	// convert arrays to slices so casts work. this should not make a copy
	// of the array and therefore be efficient operation
	i, err := m.GetDataType().Info()
	if err != nil {
		return nil, err
	}

	// Turn arrays into slices to leverage more of val.Conv's ability to convert data
	if i.Format.IsList() && value.Kind() == reflect.Array {
		value = value.Slice(0, value.Len())
	}

	switch i.Format {
	case val.FmtString:
		s := value.String()
		if len(s) == 0 {
			return nil, nil
		}
		return val.String(s), nil
	case val.FmtAny:
		if value.IsNil() {
			return nil, nil
		}
	}
	return NewValue(m.GetDataType(), value.Interface())

	// case val.FmtBool:
	// 	return val.Bool(value.Bool())
	// case val.Fmtmeta.FMT_BOOLEAN_LIST:
	// 	v.Boollist = value.Interface().([]bool)
	// case meta.FMT_INT32_LIST:
	// 	v.Intlist = value.Interface().([]int)
	// case meta.FMT_INT64_LIST:
	// 	v.Int64list = value.Interface().([]int64)
	// case meta.FMT_INT32:
	// 	v.Int = int(value.Int())
	// case meta.FMT_INT64:
	// 	v.Int64 = value.Int()
	// case meta.FMT_UINT64:
	// 	v.UInt64 = value.Interface().(uint64)
	// case meta.FMT_DECIMAL64:
	// 	v.Float = value.Float()
	// case meta.FMT_DECIMAL64_LIST:
	// 	v.Floatlist = value.Interface().([]float64)
	// case meta.FMT_STRING:
	// 	v.Str = value.String()
	// 	if len(v.Str) == 0 {
	// 		return nil, nil
	// 	}
	// case meta.FMT_STRING_LIST:
	// 	v.Strlist = value.Interface().([]string)
	// case meta.FMT_ENUMERATION:
	// 	switch value.Type().Kind() {
	// 	case reflect.String:
	// 		v, err = NewEnumByLabel(i.Enum, value.String())
	// 	default:
	// 		v, err = NewEnumValue(i.Enum, int(value.Int()))
	// 	}
	// case meta.FMT_ANYDATA:
	// 	if anyData, isAnyData := value.Interface().(map[string]interface{}); isAnyData {
	// 		if value.IsNil() {
	// 			return nil, nil
	// 		}
	// 		v.AnyData = anyData
	// 	} else {
	// 		return nil, c2.NewErr("Cannot read anydata from value that doesn't implement AnyData")
	// 	}
	// default:
	// }
	return val.Conv(i.Format, value.Interface())
}

func WriteField(m meta.HasDataType, obj interface{}, v val.Value) error {
	return WriteFieldWithFieldName(meta.MetaNameToFieldName(m.GetIdent()), m, obj, v)
}

// Look for public fields that match fieldName.  Some attempt will be made to convert value to proper
// type.
//
// TODO: We only look for fields, but it would be useful to look for methods as well with pattern
// Set___(x) or the like
func WriteFieldWithFieldName(fieldName string, m meta.HasDataType, obj interface{}, v val.Value) error {
	objType := reflect.ValueOf(obj).Elem()
	if !objType.IsValid() {
		panic(fmt.Sprintf("Cannot find property \"%s\" on invalid or nil %s", fieldName, reflect.TypeOf(obj)))
	}

	value := objType.FieldByName(fieldName)
	if !value.IsValid() {
		panic(fmt.Sprintf("Invalid property \"%s\" on %s", fieldName, reflect.TypeOf(obj)))
	}
	if v == nil {
		panic(fmt.Sprintf("No value given to set %s", m.GetIdent()))
	}
	switch v.Format() {
	case val.FmtEnum:
		e := v.(val.Enum)
		switch value.Kind() {
		case reflect.Int:
			value.SetInt(int64(e.Id))
		case reflect.String:
			value.SetString(e.Label)
		}
	case val.FmtEnumList:
		el := v.(val.EnumList)
		switch value.Elem().Kind() {
		case reflect.Int:
			value.Set(reflect.ValueOf(el.Ids()))
		case reflect.String:
			value.Set(reflect.ValueOf(el.Labels()))
		}
	}
	value.Set(reflect.ValueOf(v.Value()))
	return nil
}
