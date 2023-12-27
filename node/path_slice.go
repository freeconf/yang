package node

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
)

func parseUrlPath(pathStr string, m meta.Definition) ([]*Path, error) {
	var err error
	p := &Path{Meta: m}
	path := []*Path{}
	segments := strings.Split(pathStr, "/")
	for _, segment := range segments {

		// a/b/c same as a/b/c/
		if segment == "" {
			break
		}

		var ident string
		var keyStrs []string

		// next path segment
		seg := &Path{Parent: p}

		// has key
		if equalsMark := strings.Index(segment, "="); equalsMark >= 0 {
			if ident, err = url.QueryUnescape(segment[:equalsMark]); err != nil {
				return nil, err
			}
			keyStrs = strings.Split(segment[equalsMark+1:], ",")
			for i, escapedKeystr := range keyStrs {
				if keyStrs[i], err = url.QueryUnescape(escapedKeystr); err != nil {
					return nil, err
				}
			}
			// no key
		} else {
			if ident, err = url.QueryUnescape(segment); err != nil {
				return nil, err
			}
		}

		// find meta associated with path ident
		seg.Meta = meta.Find(p.Meta.(meta.HasDefinitions), ident)
		if seg.Meta == nil {
			return nil, fmt.Errorf("%w. %s not found in %s", fc.NotFoundError, ident, p.Meta.Ident())
		}

		if len(keyStrs) > 0 {
			if seg.Key, err = NewValuesByString(seg.Meta.(*meta.List).KeyMeta(), keyStrs...); err != nil {
				return nil, err
			}
		}

		path = append(path, seg)
		p = seg
	}
	return path, nil
}
