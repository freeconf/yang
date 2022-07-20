package fc

import "fmt"

// MapValue used in unit tests, not intended for use outsite this golang
// module so this function may be subject to change w/o warning
func MapValue(parent interface{}, key ...interface{}) interface{} {
	var child interface{}
	switch x := parent.(type) {
	case []interface{}:
		child = x[key[0].(int)]
	case []map[string]interface{}:
		child = x[key[0].(int)]
	case map[string]interface{}:
		child = x[key[0].(string)]
	case map[interface{}]interface{}:
		child = x[key[0]]
	default:
		panic(fmt.Sprintf("unsupported map type %T", parent))
	}
	if child != nil && len(key) > 1 {
		return MapValue(child, key[1:]...)
	}
	return child
}
