package doc

import (
	"fmt"
	"strings"

	"github.com/freeconf/yang/meta"
)

// functions useful in templates function maps

func escape(chars string, escChar string) func(string) string {
	charReplace := make([]string, len(chars)*2)
	for i, r := range chars {
		j := 2 * i
		charReplace[j] = string(r)
		charReplace[j+1] = escChar + string(r)
	}
	return strings.NewReplacer(charReplace...).Replace
}

func docLink(d *def) string {
	var link string
	for d.Parent != nil {
		ident := d.Meta.(meta.Identifiable).Ident()
		if link == "" {
			link = ident
		} else {
			link = ident + "/" + link
		}
		d = d.Parent
	}
	return link
}

func docPath(def *def) string {
	if def == nil || def.Parent == nil {
		return "/"
	}
	seg := def.Meta.Ident()
	if mlist, isList := def.Meta.(*meta.List); isList {
		seg += fmt.Sprintf("={%v}", docKeyId(mlist))
	}
	return docPath(def.Parent) + docTitle2(def.Meta) + "/"
}

func docKeyId(mlist *meta.List) string {
	var keyId string
	for i, k := range mlist.KeyMeta() {
		if i > 0 {
			keyId += ","
		}
		keyId += k.Ident()
	}
	return keyId
}

func docTitle(m meta.Identifiable) string {
	title := m.Ident()
	if meta.IsList(m.(meta.Meta)) {
		// ellipsis
		title += "[\u2026]"
	} else if _, isModule := m.(*meta.Module); isModule {
		// Modules should not show up as they don't show
		// up in data.
		return ""
	}
	return title
}

// Only difference between title is that list items show keys
func docTitle2(m meta.Identifiable) string {
	if mlist, isList := m.(*meta.List); isList {
		title := m.Ident()
		title += fmt.Sprintf("={%v}", docKeyId(mlist))
		return title
	}
	return docTitle(m)
}

func docFieldType(f *def) string {
	var fieldType string
	if meta.IsLeaf(f.Meta) {
		dt := f.Meta.(meta.HasType).Type()
		fieldType = dt.Ident()
		if dt.Format().IsList() {
			fieldType = fieldType + "[]"
		}
	}
	return fieldType
}
