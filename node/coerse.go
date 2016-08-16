package node

import (
	"fmt"
	"errors"
	"github.com/dhubler/c2g/meta"
)

var NO_KEYS = make([]*Value, 0)

func CoerseKeys(list *meta.List, keyStrs []string) ([]*Value, error) {
	var err error
	if len(keyStrs) == 0 {
		return NO_KEYS, nil
	}
	if len(list.Key) != len(keyStrs) {
		return NO_KEYS, errors.New("Missing keys on " + list.GetIdent())
	}
	values := make([]*Value, len(keyStrs))
	for i, keyStr := range keyStrs {
		keyProp := meta.FindByIdent2(list, list.Key[i])
		if keyProp == nil {
			return nil, errors.New(fmt.Sprintf("no key prop %s on %s", list.Key[i], list.GetIdent()))
		}
		values[i] = &Value{
			Type: keyProp.(meta.HasDataType).GetDataType(),
		}
		err = values[i].CoerseStrValue(keyStr)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func SerializeKey(key []*Value) []string {
	keyStr := make([]string, len(key))
	for i, v := range key {
		keyStr[i] = v.String()
	}
	return keyStr
}
