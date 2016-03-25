package node

import (
	"fmt"
	"sort"
	"strings"
)

type KeyListBuilder struct {
	set      map[string]struct{}
	prefix   string
	keyStart int
}

func NewKeyListBuilder(listPath string) *KeyListBuilder {
	prefix := fmt.Sprint(listPath, "=")
	return &KeyListBuilder{
		set:      make(map[string]struct{}, 10),
		prefix:   prefix,
		keyStart: len(prefix),
	}
}

func (klb *KeyListBuilder) ParseKey(path string) bool {
	// TODO: performance - most efficient way? sort first?
	if strings.HasPrefix(path, klb.prefix) {
		keyEnd := strings.IndexRune(path[klb.keyStart:], '/')
		var key string
		if keyEnd < 0 {
			key = path[klb.keyStart:]
		} else {
			key = path[klb.keyStart : klb.keyStart+keyEnd]
		}
		klb.set[key] = struct{}{}
		return true
	}
	return false
}

func (klb *KeyListBuilder) List() []string {
	keys := make([]string, len(klb.set))
	var i int
	for k, _ := range klb.set {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
