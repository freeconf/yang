package node

import (
	"fmt"
	"meta"
	"blit"
)

func EditNotImplemented(m meta.Meta) error {
	return blit.NewErrC(fmt.Sprintf("editing of \"%s\" not implemented", m.GetIdent()),  501)
}

func NotImplementedByName(ident string) error {
	return blit.NewErrC(fmt.Sprintf("browsing of \"%s\" not implemented", ident),  501)
}

func NotImplemented(m meta.Meta) error {
	msg := fmt.Sprintf("browsing of \"%s.%s\" not implemented", m.GetParent().GetIdent(), m.GetIdent())
	return blit.NewErrC(msg,  501)
}

func PathNotFound(path string) error {
	return blit.NewErrC(fmt.Sprintf("item identified with path \"%s\" not found", path),  404)
}

func ListItemNotFound(key string) error {
	return blit.NewErrC(fmt.Sprintf("item identified with key \"%s\" not found", key),  404)
}
