package node

import (
	"errors"
	"fmt"

	"github.com/dhubler/c2g/meta"
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
			Request: Request{
				Selection: sel,
			},
			Meta: keyMeta,
		}
		var hnd ValueHandle
		if err = sel.node.Field(r, &hnd); err != nil {
			return nil, err
		}
		if hnd.Val == nil {
			return nil, errors.New(fmt.Sprint("Key value is nil for ", keyIdent))
		}
		key.Type = keyMeta.GetDataType()
		values[i] = hnd.Val
	}
	return
}
