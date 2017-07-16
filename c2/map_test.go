package c2

import (
	"testing"

	"reflect"
)

func Test_Merge(t *testing.T) {
	a := map[string]interface{}{
		"a": 1,
		"c": 3,
		"x": map[string]interface{}{
			"y": "a",
		},
		"f": []interface{}{
			"one", "two",
		},
		"p": []map[string]interface{}{
			map[string]interface{}{
				"a1": "b1",
			},
		},
	}
	b := map[string]interface{}{
		"b": 2,
		"c": 4,
		"x": map[string]interface{}{
			"z": "q",
		},
		"f": []interface{}{
			"uno", "dos", "tres",
		},
		"p": []map[string]interface{}{
			map[string]interface{}{
				"a2": "b2",
			},
		},
	}
	expected := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
		"x": map[string]interface{}{
			"y": "a",
			"z": "q",
		},
		"f": []interface{}{
			"one", "two", "tres",
		},
		"p": []map[string]interface{}{
			map[string]interface{}{
				"a1": "b1",
				"a2": "b2",
			},
		},
	}

	actual := MapMerge(a, b)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v", actual)
	}
}
