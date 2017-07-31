package val

import "reflect"

func Equal(a Value, b Value) bool {
	if a == nil {
		if b == nil {
			return true
		}
		return false
	}
	if b == nil {
		return false
	}
	if a.Format() != b.Format() {
		return false
	}
	if a.Format().IsList() {
		return reflect.DeepEqual(a.Value(), b.Value())
	}
	return a.(Comparable).Compare(b.(Comparable)) == 0
}

type Reducer func(index int, v Value, data interface{}) interface{}

type Eacher func(index int, v Value)

func Reduce(v Value, initial interface{}, f Reducer) interface{} {
	result := initial
	if l, listable := v.(Listable); listable {
		len := l.Len()
		for i := 0; i < len; i++ {
			result = f(i, l.Item(i), result)
		}
	} else {
		result = f(0, v, result)
	}
	return result
}

func ForEach(v Value, f Eacher) {
	Reduce(v, nil, func(index int, item Value, data interface{}) interface{} {
		f(index, item)
		return nil
	})
}
