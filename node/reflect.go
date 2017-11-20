package node

import (
	"fmt"
	"reflect"
	"unicode"

	"github.com/freeconf/c2g/val"

	"github.com/freeconf/c2g/meta"
)

func ReadField(m meta.HasDataType, obj interface{}) (val.Value, error) {
	return ReadFieldWithFieldName(MetaNameToFieldName(m.Ident()), m, obj)
}

func ReadFieldWithFieldName(fieldName string, m meta.HasDataType, obj interface{}) (v val.Value, err error) {
	objVal := reflect.ValueOf(obj)
	if objVal.Kind() == reflect.Interface || objVal.Kind() == reflect.Ptr {
		objVal = objVal.Elem()
		if objVal.Kind() == reflect.Ptr {
			panic(fmt.Sprintf("Pointer to a pointer not legal %s on %v ", m.Ident(), reflect.TypeOf(obj)))
		}
	}
	value := objVal.FieldByName(fieldName)

	if !value.IsValid() {
		panic(fmt.Sprintf("Field not found: %s on %v ", m.Ident(), reflect.TypeOf(obj)))
	}

	// convert arrays to slices so casts work. this should not make a copy
	// of the array and therefore be efficient operation
	dt := m.DataType()

	// Turn arrays into slices to leverage more of val.Conv's ability to convert data
	if dt.Format().IsList() && value.Kind() == reflect.Array {
		value = value.Slice(0, value.Len())
	}

	switch dt.Format() {
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
	return NewValue(dt, value.Interface())
}

func WriteField(m meta.HasDataType, obj interface{}, v val.Value) error {
	return WriteFieldWithFieldName(MetaNameToFieldName(m.Ident()), m, obj, v)
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
		panic(fmt.Sprintf("No value given to set %s", m.Ident()))
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

func MetaNameToFieldName(in string) string {
	// assumes fix is always shorter because char can be dropped and not added
	fixed := make([]rune, len(in))
	cap := true
	j := 0
	for _, r := range in {
		if r == '-' {
			cap = true
		} else {
			if cap {
				fixed[j] = unicode.ToUpper(r)
			} else {
				fixed[j] = r
			}
			j += 1
			cap = false
		}
	}
	return string(fixed[:j])
}
