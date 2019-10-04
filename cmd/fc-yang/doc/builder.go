package doc

import "io"

type DocDefBuilder interface {
	Generate(doc *Doc, template string, out io.Writer) error
	BuiltinTemplate() string
}
