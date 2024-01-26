package car

import (
	"embed"

	"github.com/freeconf/yang/source"
)

//go:embed *.yang car.go manage.go
var internal embed.FS

// Access to car.yang
var YPath = source.EmbedDir(internal, ".")
