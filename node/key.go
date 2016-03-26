package node

import (
	"errors"
	"fmt"
	"github.com/blitter/meta"
)

func ReadKeys(sel *Selection) (values []*Value, err error) {
	if len(sel.path.key) > 0 {
		return sel.path.key, nil
	}
	list := sel.path.meta.(*meta.List)
	values = make([]*Value, len(list.Key))
	var key *Value
	for i, keyIdent := range list.Key {
		keyMeta := meta.FindByIdent2(sel.path.meta, keyIdent).(meta.HasDataType)
		r := FieldRequest{
			Request:Request {
				Selection: sel,
			},
			Meta: keyMeta,
		}
		if key, err = sel.node.Read(r); err != nil {
			return nil, err
		}
		if key == nil {
			return nil, errors.New(fmt.Sprint("Key value is nil for ", keyIdent))
		}
		key.Type = keyMeta.GetDataType()
		values[i] = key
	}
	return
}

