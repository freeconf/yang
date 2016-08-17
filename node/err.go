package node

import (
	"fmt"

	"github.com/dhubler/c2g/c2"
	"github.com/dhubler/c2g/meta"
)

func EditNotImplemented(m meta.Meta) error {
	return c2.NewErrC(fmt.Sprintf("editing of \"%s\" not implemented", m.GetIdent()), 501)
}

func NotImplementedByName(ident string) error {
	return c2.NewErrC(fmt.Sprintf("browsing of \"%s\" not implemented", ident), 501)
}

func NotImplemented(m meta.Meta) error {
	msg := fmt.Sprintf("browsing of \"%s.%s\" not implemented", m.GetParent().GetIdent(), m.GetIdent())
	return c2.NewErrC(msg, 501)
}

func PathNotFound(path string) error {
	return c2.NewErrC(fmt.Sprintf("item identified with path \"%s\" not found", path), 404)
}

func ListItemNotFound(key string) error {
	return c2.NewErrC(fmt.Sprintf("item identified with key \"%s\" not found", key), 404)
}
