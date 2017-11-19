package node

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Example:
//  DataValue(data, "foo.10.bar.blah.0")
func MapValue(container map[string]interface{}, path string) interface{} {
	segments := strings.Split(path, ".")
	var v interface{}
	v = container
	for _, seg := range segments {
		switch x := v.(type) {
		case []map[string]interface{}:
			n, _ := strconv.Atoi(seg)
			v = x[n]
		case []interface{}:
			n, _ := strconv.Atoi(seg)
			v = x[n]
		case map[string]interface{}:
			v = x[seg]
		default:
			panic(fmt.Sprintf("Bad type %s on %s", reflect.TypeOf(v), seg))
		}
		if v == nil {
			return nil
		}
	}
	return v
}
