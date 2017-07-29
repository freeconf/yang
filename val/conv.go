package val

import (
	"fmt"
	"strconv"

	"github.com/c2stack/c2g/c2"
)

func Conv(f Format, val interface{}) (Value, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = c2.NewErr(fmt.Sprintf("Could not convert %v to type  %s", val, f))
		}
	}()
	if val == nil {
		return nil, nil
	}
	switch f {
	case FmtBool:
		if x, err := toBool(val); err != nil {
			return nil, err
		} else {
			return Bool(x), nil
		}
	case FmtBoolList:
		if x, err := toBoolList(val); err != nil {
			return nil, err
		} else {
			return BoolList(x), nil
		}
	case FmtInt32:
		if x, err := toInt32(val); err != nil {
			return nil, err
		} else {
			return Int32(x), nil
		}
	case FmtInt32List:
		if x, err := toInt32List(val); err != nil {
			return nil, err
		} else {
			return Int32List(x), nil
		}
	}

	return nil, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to %s value", val, f.String()))
}

func toInt32List(val interface{}) ([]int, error) {
	switch x := val.(type) {
	case []int:
		return x, nil
	case []interface{}:
		l := make([]int, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toInt32(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	case []float64:
		l := make([]int, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toInt32(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	case []string:
		l := make([]int, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toInt32(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	default:
		if i, notSingle := toInt32(val); notSingle == nil {
			return []int{i}, nil
		}
	}
	return nil, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to []int", val))
}

func toInt32(val interface{}) (int, error) {
	switch x := val.(type) {
	case int:
		return x, nil
	case int64:
		return int(x), nil
	case string:
		return strconv.Atoi(x)
	case float64:
		return int(x), nil
	case float32:
		return int(x), nil
	}
	return 0, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to int", val))
}

func toBoolList(val interface{}) ([]bool, error) {
	switch x := val.(type) {
	case []bool:
		return x, nil
	case []interface{}:
		l := make([]bool, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toBool(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	case []string:
		l := make([]bool, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toBool(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	default:
		if b, canBool := toBool(val); canBool == nil {
			return []bool{b}, nil
		}
	}
	return nil, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to boolean array", val))
}

func toBool(val interface{}) (bool, error) {
	switch x := val.(type) {
	case bool:
		return x, nil
	case string:
		switch x {
		case "1", "true", "yes":
			return true, nil
		case "0", "false", "np":
			return false, nil
		}
	}
	return false, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to boolean value", val))
}

type Union struct {
}
