package c2

/*
  Merge two maps using item in primary map should there be a conflict. Useful for mixing-in
  startup or details maps
*/
func MapMerge(primary map[string]interface{}, secondary map[string]interface{}) map[string]interface{} {
	copy := make(map[string]interface{})
	for key, aVal := range primary {
		if bVal, both := secondary[key]; both {
			copy[key] = mergeVal(aVal, bVal)
		} else {
			copy[key] = aVal
		}
	}
	for key, bVal := range secondary {
		if _, inA := primary[key]; !inA {
			copy[key] = bVal
		}
	}
	return copy
}

func mergeVal(primary interface{}, secondary interface{}) interface{} {
	switch x := primary.(type) {
	case map[string]interface{}:
		return MapMerge(x, secondary.(map[string]interface{}))
	case []map[string]interface{}:
		return mergeMapSlice(x, secondary.([]map[string]interface{}))
	case []interface{}:
		return mergeSlice(x, secondary.([]interface{}))
	}
	return primary
}

func mergeMapSlice(primary []map[string]interface{}, secondary []map[string]interface{}) []map[string]interface{} {
	l := max(len(primary), len(secondary))
	copy := make([]map[string]interface{}, l)
	for i := 0; i < l; i++ {
		inA := i < len(primary)
		inB := i < len(secondary)
		if inA && inB {
			copy[i] = MapMerge(primary[i], secondary[i])
		} else if inA {
			copy[i] = primary[i]
		} else {
			copy[i] = secondary[i]
		}
	}
	return copy
}

func mergeSlice(primary []interface{}, secondary []interface{}) []interface{} {
	l := max(len(primary), len(secondary))
	copy := make([]interface{}, l)
	for i := 0; i < l; i++ {
		inA := i < len(primary)
		inB := i < len(secondary)
		if inA && inB {
			copy[i] = mergeVal(primary[i], secondary[i])
		} else if inA {
			copy[i] = primary[i]
		} else {
			copy[i] = secondary[i]
		}
	}
	return copy
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
