package yang

import (
	"embed"

	"github.com/freeconf/yang/source"
)

//go:embed yang/*.yang
var internal embed.FS

// Access to fc-yang and fc-doc yang definitions.
var InternalYPath = source.EmbedDir(internal, "yang")
