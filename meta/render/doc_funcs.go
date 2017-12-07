package render

import (
	"fmt"
	"strings"

	"github.com/freeconf/c2g/meta"
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

func docLink(o interface{}) string {
	switch x := o.(type) {
	case *DocDef:
		if x.Parent == nil {
			return ""
		}
		return docLink(x.Parent) + "/" + x.Meta.Ident()
	case *DocAction:
		return docLink(x.Def) + "/" + x.Meta.Ident()
	case *DocEvent:
		return docLink(x.Def) + "/" + x.Meta.Ident()
	}
	panic(fmt.Sprintf("not supported %T", o))
}

func docPath(def *DocDef) string {
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
	if meta.IsList(m) {
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

func docFieldType(f *DocField) string {
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
