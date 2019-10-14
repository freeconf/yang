package doc

import "io"

type builder interface {
	generate(doc *doc, template string, out io.Writer) error
	builtinTemplate() string
}
