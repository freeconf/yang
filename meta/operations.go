package meta

import "container/list"

type operations struct {
	compile compileFunc
}

type compileFunc func(*list.List) error
