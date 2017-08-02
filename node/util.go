package node

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/c2stack/c2g/meta"
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

func RenameMeta(m meta.Meta, rename string) {
	switch m := m.(type) {
	case *meta.Container:
		m.Ident = rename
	case *meta.Module:
		m.Ident = rename
	case *meta.List:
		m.Ident = rename
	case *meta.Leaf:
		m.Ident = rename
	case *meta.LeafList:
		m.Ident = rename
	case *meta.Choice:
		m.Ident = rename
	case *meta.ChoiceCase:
		m.Ident = rename
	case *meta.Any:
		m.Ident = rename
	default:
		panic("rename not supported on " + reflect.TypeOf(m).Name())
	}
}

func PathModule(path *Path) meta.MetaList {
	p := path
	for {
		parent := p.Parent()
		if parent == nil {
			return p.Meta().(meta.MetaList)
		}
		p = parent
	}
}
