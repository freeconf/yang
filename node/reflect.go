package node

import (
	"fmt"
	"reflect"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
)

func ReadField(m meta.HasDataType, obj interface{}) (*Value, error) {
	return ReadFieldWithFieldName(meta.MetaNameToFieldName(m.GetIdent()), m, obj)
}

func ReadFieldWithFieldName(fieldName string, m meta.HasDataType, obj interface{}) (v *Value, err error) {
	objVal := reflect.ValueOf(obj)
	if objVal.Kind() == reflect.Interface || objVal.Kind() == reflect.Ptr {
		objVal = objVal.Elem()
	}
	value := objVal.FieldByName(fieldName)

	if !value.IsValid() {
		panic(fmt.Sprintf("Field not found: %s on %v ", m.GetIdent(), reflect.TypeOf(obj)))
	}

	// convert arrays to slices so casts work. this should not make a copy
	// of the array and therefore be efficient operation
	if meta.IsListFormat(m.GetDataType().Format()) && value.Kind() == reflect.Array {
		value = value.Slice(0, value.Len())
	}

	v = &Value{Type: m.GetDataType()}
	switch v.Type.Format() {
	case meta.FMT_BOOLEAN:
		v.Bool = value.Bool()
	case meta.FMT_BOOLEAN_LIST:
		v.Boollist = value.Interface().([]bool)
	case meta.FMT_INT32_LIST:
		v.Intlist = value.Interface().([]int)
	case meta.FMT_INT64_LIST:
		v.Int64list = value.Interface().([]int64)
	case meta.FMT_INT32:
		v.Int = int(value.Int())
	case meta.FMT_INT64:
		v.Int64 = value.Int()
	case meta.FMT_UINT64:
		v.UInt64 = value.Interface().(uint64)
	case meta.FMT_DECIMAL64:
		v.Float = value.Float()
	case meta.FMT_DECIMAL64_LIST:
		v.Floatlist = value.Interface().([]float64)
	case meta.FMT_STRING:
		v.Str = value.String()
		if len(v.Str) == 0 {
			return nil, nil
		}
	case meta.FMT_STRING_LIST:
		v.Strlist = value.Interface().([]string)
	case meta.FMT_ENUMERATION:
		switch value.Type().Kind() {
		case reflect.String:
			v.SetEnumByLabel(value.String())
		default:
			v.SetEnum(int(value.Int()))
		}
	case meta.FMT_ANYDATA:
		if anyData, isAnyData := value.Interface().(map[string]interface{}); isAnyData {
			if value.IsNil() {
				return nil, nil
			}
			v.AnyData = anyData
		} else {
			return nil, c2.NewErr("Cannot read anydata from value that doesn't implement AnyData")
		}
	default:
		panic(fmt.Sprintf("Format code %d not implemented", m.GetDataType().Format))
	}
	return
}

func WriteField(m meta.HasDataType, obj interface{}, v *Value) error {
	return WriteFieldWithFieldName(meta.MetaNameToFieldName(m.GetIdent()), m, obj, v)
}

func WriteFieldWithFieldName(fieldName string, m meta.HasDataType, obj interface{}, v *Value) error {
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
	if v.Type == nil {
		panic(fmt.Sprintf("No type or format found %s", m.GetIdent()))
	}
	switch v.Type.Format() {
	case meta.FMT_BOOLEAN_LIST:
		value.Set(reflect.ValueOf(v.Boollist))
	case meta.FMT_BOOLEAN:
		value.SetBool(v.Bool)
	case meta.FMT_INT32_LIST:
		value.Set(reflect.ValueOf(v.Intlist))
	case meta.FMT_INT32:
		value.SetInt(int64(v.Int))
	case meta.FMT_INT64_LIST:
		value.Set(reflect.ValueOf(v.Int64list))
	case meta.FMT_INT64:
		value.SetInt(v.Int64)
	case meta.FMT_DECIMAL64:
		value.Set(reflect.ValueOf(v.Float))
	case meta.FMT_DECIMAL64_LIST:
		value.Set(reflect.ValueOf(v.Floatlist))
	case meta.FMT_STRING_LIST:
		value.Set(reflect.ValueOf(v.Strlist))
	case meta.FMT_STRING:
		value.SetString(v.Str)
	case meta.FMT_ENUMERATION:
		switch value.Type().Kind() {
		case reflect.String:
			value.SetString(v.Str)
		default:
			value.SetInt(int64(v.Int))
		}
	case meta.FMT_ANYDATA:
		// could support writing to string as well
		value.Set(reflect.ValueOf(v.AnyData))

	// TODO: Enum list
	default:
		panic(m.GetIdent() + " not implemented")
	}
	return nil
}
