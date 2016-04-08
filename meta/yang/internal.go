package yang

import (
	"github.com/c2g/meta"
	"strings"
	"fmt"
)

type YangPool map[string]string

var internal = make(YangPool)

func (self YangPool) OpenStream(streamId string) (meta.DataStream, error) {
	if s, found := self[streamId]; found {
		return strings.NewReader(s), nil
	}
	return nil, nil
}

func InternalModule(name string) *meta.Module {
	// TODO: performance - return deep copy of cached copy
	inlineYang, err := LoadModule(InternalYang(), name)
	if err != nil {
		msg := fmt.Sprintf("Error parsing %s yang, %s", name, err.Error())
		panic(msg)
	}
	return inlineYang
}

func InternalYang() YangPool {
	return internal
}

