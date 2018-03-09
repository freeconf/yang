package val

import (
	"fmt"
	"strconv"

	"github.com/freeconf/gconf/c2"
)

func ConvOneOf(f []Format, val interface{}) (Value, Format, error) {
	for _, f := range f {
		if v, err := Conv(f, val); err == nil {
			return v, f, nil
		}
	}
	return nil, 0, c2.NewErr(fmt.Sprintf("Could not convert %v to any of the allowed types", val))
}

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
	case FmtInt8:
		if x, err := toInt8(val); err != nil {
			return nil, err
		} else {
			return Int8(x), nil
		}
	case FmtInt16:
		if x, err := toInt16(val); err != nil {
			return nil, err
		} else {
			return Int16(x), nil
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
	case FmtInt64:
		if x, err := toInt64(val); err != nil {
			return nil, err
		} else {
			return Int64(x), nil
		}
	case FmtInt64List:
		if x, err := toInt64List(val); err != nil {
			return nil, err
		} else {
			return Int64List(x), nil
		}
	case FmtUInt8:
		if x, err := toUInt8(val); err != nil {
			return nil, err
		} else {
			return UInt64(x), nil
		}
	case FmtUInt16:
		if x, err := toUInt16(val); err != nil {
			return nil, err
		} else {
			return UInt16(x), nil
		}
	case FmtUInt32:
		if x, err := toUInt32(val); err != nil {
			return nil, err
		} else {
			return UInt32(x), nil
		}
	case FmtUInt64:
		if x, err := toUInt64(val); err != nil {
			return nil, err
		} else {
			return UInt64(x), nil
		}
	case FmtUInt64List:
		if x, err := toUInt64List(val); err != nil {
			return nil, err
		} else {
			return UInt64List(x), nil
		}
	case FmtDecimal64:
		if x, err := toDecimal64(val); err != nil {
			return nil, err
		} else {
			return Decimal64(x), nil
		}
	case FmtDecimal64List:
		if x, err := toDecimal64List(val); err != nil {
			return nil, err
		} else {
			return Decimal64List(x), nil
		}
	case FmtAny:
		return Any{Thing: val}, nil
	case FmtString:
		if x, err := toString(val); err != nil {
			return nil, err
		} else {
			return String(x), nil
		}
	case FmtStringList:
		if x, err := toStringList(val); err != nil {
			return nil, err
		} else {
			return StringList(x), nil
		}
	}

	return nil, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to %s value", val, f.String()))
}

func toDecimal64(val interface{}) (float64, error) {
	switch x := val.(type) {
	case int:
		return float64(x), nil
	case int64:
		return float64(x), nil
	case float32:
		return float64(x), nil
	case float64:
		return x, nil
	case string:
		return strconv.ParseFloat(x, 64)
	}
	return 0, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to float64", val))
}

func toDecimal64List(val interface{}) ([]float64, error) {
	switch x := val.(type) {
	case []float64:
		return x, nil
	case []interface{}:
		l := make([]float64, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toDecimal64(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	case []string:
		l := make([]float64, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toDecimal64(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	default:
		// TODO: Use reflection on general array type
		if i, notSingle := toDecimal64(val); notSingle == nil {
			return []float64{i}, nil
		}
	}
	return nil, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to []int", val))
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
		// TODO: Use reflection on general array type

		if i, notSingle := toInt32(val); notSingle == nil {
			return []int{i}, nil
		}
	}
	return nil, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to []int", val))
}

func toInt8(val interface{}) (int8, error) {
	switch x := val.(type) {
	case int8:
		return x, nil
	default:
		i, err := toInt32(val)
		max := 2 ^ 7
		if err == nil && i > -max && i < max {
			return int8(i), nil
		}
	}
	return 0, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to int8", val))
}

func toInt16(val interface{}) (int16, error) {
	switch x := val.(type) {
	case int16:
		return x, nil
	default:
		i, err := toInt32(val)
		max := 2 ^ 15
		if err == nil && i > -max && i < max {
			return int16(i), nil
		}
	}
	return 0, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to int16", val))
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

func toInt64List(val interface{}) ([]int64, error) {
	switch x := val.(type) {
	case []int:
		l := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			l[i] = int64(x[i])
		}
		return l, nil
	case []int64:
		return x, nil
	case []interface{}:
		l := make([]int64, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toInt64(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	case []float64:
		l := make([]int64, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toInt64(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	case []string:
		l := make([]int64, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toInt64(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	default:
		// TODO: Use reflection on general array type

		if i, notSingle := toInt64(val); notSingle == nil {
			return []int64{i}, nil
		}
	}
	return nil, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to []int", val))
}

func toInt64(val interface{}) (int64, error) {
	switch x := val.(type) {
	case int:
		return int64(x), nil
	case int64:
		return x, nil
	case string:
		return strconv.ParseInt(x, 10, 64)
	case float64:
		return int64(x), nil
	case float32:
		return int64(x), nil
	}
	return 0, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to int64", val))
}

func toUInt8(val interface{}) (uint8, error) {
	switch x := val.(type) {
	case uint8:
		return uint8(x), nil
	default:
		i, err := toInt32(val)
		if err == nil && i > 0 && i < 2^8 {
			return uint8(i), nil
		}
	}
	return 0, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to uint8", val))
}

func toUInt16(val interface{}) (uint16, error) {
	switch x := val.(type) {
	case uint16:
		return uint16(x), nil
	default:
		i, err := toInt32(val)
		if err == nil && i > 0 && i < 2^16 {
			return uint16(i), nil
		}
	}
	return 0, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to uint16", val))
}

func toUInt32(val interface{}) (uint32, error) {
	switch x := val.(type) {
	case uint32:
		return uint32(x), nil
	default:
		i, err := toInt64(val)
		if err == nil && i > 0 && i < 2^32 {
			return uint32(i), nil
		}
	}
	return 0, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to uint32", val))
}

func toUInt64(val interface{}) (uint64, error) {
	switch x := val.(type) {
	case int:
		return uint64(x), nil
	case int64:
		return uint64(x), nil
	case uint64:
		return x, nil
	case string:
		i, err := strconv.ParseInt(x, 10, 64)
		return uint64(i), err
	case float64:
		return uint64(x), nil
	case float32:
		return uint64(x), nil
	}
	return 0, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to uint", val))
}

func toUInt64List(val interface{}) ([]uint64, error) {
	switch x := val.(type) {
	case []int:
		l := make([]uint64, len(x))
		for i := 0; i < len(x); i++ {
			l[i] = uint64(x[i])
		}
		return l, nil
	case []uint64:
		return x, nil
	case []interface{}:
		l := make([]uint64, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toUInt64(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	case []float64:
		l := make([]uint64, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toUInt64(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	case []string:
		l := make([]uint64, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toUInt64(x[i]); err != nil {
				return nil, err
			}
		}
		return l, nil
	default:
		// TODO: Use reflection on general array type

		if i, notSingle := toUInt64(val); notSingle == nil {
			return []uint64{i}, nil
		}
	}
	return nil, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to int64 array", val))
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

func toString(val interface{}) (string, error) {
	switch x := val.(type) {
	case float64:
		// wrong format, truncating decimals as most likely mistake but
		// will not please everyone.  Get input in correct format by placing
		// quotes around data.
		return strconv.FormatFloat(x, 'f', 0, 64), nil
	}
	return fmt.Sprintf("%v", val), nil
}

func toStringList(val interface{}) ([]string, error) {
	switch x := val.(type) {
	case []string:
		return x, nil
	case []float64:
		l := make([]string, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toString(x[i]); err != nil {
				return nil, err
			}
		}
		return l, err
	case []interface{}:
		l := make([]string, len(x))
		var err error
		for i := 0; i < len(x); i++ {
			if l[i], err = toString(x[i]); err != nil {
				return nil, err
			}
		}
		return l, err
	}
	return nil, c2.NewErr(fmt.Sprintf("cannot coerse '%v' to []string", val))
}
