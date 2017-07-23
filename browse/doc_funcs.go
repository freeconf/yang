package browse

import (
	"fmt"
	"strings"

	"github.com/c2stack/c2g/meta"
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
		return docLink(x.Parent) + "/" + x.Meta.GetIdent()
	case *DocAction:
		return docLink(x.Def) + "/" + x.Meta.GetIdent()
	case *DocEvent:
		return docLink(x.Def) + "/" + x.Meta.GetIdent()
	}
	panic(fmt.Sprintf("not supported %T", o))
}

func docPath(def *DocDef) string {
	if def == nil || def.Parent == nil {
		return "/"
	}
	seg := def.Meta.GetIdent()
	if mlist, isList := def.Meta.(*meta.List); isList {
		seg += fmt.Sprintf("={%v}", strings.Join(mlist.Key, ","))
	}
	return docPath(def.Parent) + docTitle2(def.Meta) + "/"
}

func docTitle(m meta.Meta) string {
	title := m.GetIdent()
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
func docTitle2(m meta.Meta) string {
	if mlist, isList := m.(*meta.List); isList {
		title := m.GetIdent()
		title += fmt.Sprintf("={%v}", strings.Join(mlist.Key, ","))
		return title
	}
	return docTitle(m)
}

func docFieldType(f *DocField) string {
	var fieldType string
	if meta.IsLeaf(f.Meta) {
		leafMeta := f.Meta.(meta.HasDataType)
		fieldType = leafMeta.GetDataType().Ident
		if meta.IsListFormat(leafMeta.GetDataType().Format()) {
			fieldType = fieldType + "[]"
		}
	}
	return fieldType
}
