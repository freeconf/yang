package nodeutil

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/freeconf/yang/meta"
)

type structAsContainer struct {
	ref    *Node
	src    reflect.Value
	fields map[meta.Definition]reflectFieldHandler
}

type reflectFieldHandler interface {
	clear() error
	get() (reflect.Value, error)
	set(reflect.Value) error
	fieldType() reflect.Type
}

func newStructAsContainer(ref *Node, src reflect.Value) *structAsContainer {
	if src.Kind() == reflect.Struct {
		panic(fmt.Sprintf("struct %s not allowed, need pointer to struct", src.Type()))
	}
	return &structAsContainer{
		ref:    ref,
		src:    src,
		fields: make(map[meta.Definition]reflectFieldHandler),
	}
}

func (def *structAsContainer) elem() reflect.Value {
	if def.src.Kind() == reflect.Pointer {
		return def.src.Elem()
	}
	return def.src
}

func (def *structAsContainer) clear(m meta.Definition) error {
	// TODO: catch panic
	h, err := def.getHandler(m)
	if err != nil {
		return err
	}
	h.clear()
	return nil
}

func (def *structAsContainer) exists(m meta.Definition) bool {
	// TODO: catch panic and return false
	if h, herr := def.getHandler(m); herr == nil {
		v, err := h.get()
		return err == nil && v.IsValid() && !v.IsZero()
	}
	return false
}

func (def *structAsContainer) getHandler(m meta.Definition) (reflectFieldHandler, error) {
	h, exists := def.fields[m]
	if !exists {
		var err error
		if h, err = def.newHandler(m); err != nil {
			return nil, err
		}
		def.fields[m] = h
	}
	return h, nil
}

func (def *structAsContainer) newHandler(m meta.Definition) (reflectFieldHandler, error) {
	opts := def.ref.options(m)
	if x := findReflectByField(def.src, m, opts); x != nil {
		return x, nil
	}
	return nil, fmt.Errorf("could not find field for '%s' on %s using reflection", m.Ident(), def.elem().Type())
}

func (def *structAsContainer) newChild(m meta.HasDataDefinitions) (reflect.Value, error) {
	var empty reflect.Value
	h, err := def.getHandler(m)
	if err != nil {
		return empty, err
	}
	return def.ref.NewObject(h.fieldType(), m, false)
}

func (def *structAsContainer) get(m meta.Definition) (reflect.Value, error) {
	var empty reflect.Value
	h, err := def.getHandler(m)
	if err != nil {
		return empty, err
	}
	v, err := h.get()
	if err != nil {
		return empty, err
	}
	if !v.IsValid() {
		return empty, nil
	}
	if canNil(v.Kind()) && v.IsNil() {
		return empty, nil
	}
	if dt, valid := m.(meta.HasType); valid {
		// Turn arrays into slices to leverage more of val.Conv's ability to convert data
		if dt.Type().Format().IsList() && v.Kind() == reflect.Array {
			v = v.Slice(0, v.Len())
		}
	}

	return v, nil
}

func (def *structAsContainer) getType(m meta.Definition) (reflect.Type, error) {
	var empty reflect.Type
	h, err := def.getHandler(m)
	if err != nil {
		return empty, err
	}
	return h.fieldType(), nil
}

func canNil(k reflect.Kind) bool {
	switch k {
	// see https://pkg.go.dev/reflect@go1.20.7#Value.IsNil
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return true
	}
	return false
}

func (def *structAsContainer) set(m meta.Definition, v reflect.Value) error {
	// TODO: catch panic and return err
	h, err := def.getHandler(m)
	if err != nil {
		return err
	}
	return h.set(v)
}

func reflectFieldCandidates(o NodeOptions, m meta.Definition) []string {
	if o.Ident != "" {
		return []string{MetaNameToFieldName(o.Ident)}
	}
	candidates := []string{MetaNameToFieldName(m.Ident())}
	if o.TryPluralOnLists {
		_, isLeafList := m.(*meta.LeafList)
		if meta.IsList(m) || isLeafList {
			candidates = append(candidates, candidates[0]+"s")
		}
	}
	return candidates
}

func reflectAccessorCandidates(o NodeOptions, m meta.Definition, prefix string) []string {
	fieldCandidates := reflectFieldCandidates(o, m)
	var candidates []string
	// put Get???s first as it is more conclusive if found
	for _, x := range fieldCandidates {
		candidates = append(candidates, prefix+x)
	}
	candidates = append(candidates, fieldCandidates...)
	return candidates
}

func findReflectByField(src reflect.Value, m meta.Definition, opts NodeOptions) *reflectByField {

	x := &reflectByField{src: src, m: m, opts: opts}

	// we don't check types because we relyin on value coersion of node.NewValue

	// look for yang tags first as that might want to override things picked up
	// by default. explicit tags are definitive, no need for heuristics below
	for i := 0; i < x.elem().Type().NumField(); i++ {
		f := x.elem().Type().Field(i)
		if tag, ok := f.Tag.Lookup("yang"); ok {
			name, _, _ := strings.Cut(tag, ",")
			if name == m.Ident() {
				x.f = f
				return x
			}
		}
	}

	for _, candidate := range reflectFieldCandidates(opts, m) {
		if f, found := x.elem().Type().FieldByName(candidate); found {
			x.f = f
			// keep going in case we find an explicit getter
		}
	}

	getPrefix := opts.GetterPrefix
	if getPrefix == "" {
		getPrefix = "Get"
	}
	for _, candidate := range reflectAccessorCandidates(opts, m, getPrefix) {
		if m, found := x.src.Type().MethodByName(candidate); found {
			// litmus test: returns something and takes nothing except receiever
			if m.Type.NumOut() >= 1 && m.Type.NumIn() == 1 {
				x.getter = m
				break
			}
		}
	}
	setPrefix := opts.SetterPrefix
	if setPrefix == "" {
		setPrefix = "Set"
	}
	for _, candidate := range reflectAccessorCandidates(opts, m, setPrefix) {
		if m, found := x.src.Type().MethodByName(candidate); found {
			// litmus test: returns nothing or one (err) and takes exactly one plus reciever
			if m.Type.NumOut() <= 1 && m.Type.NumIn() == 2 {
				x.setter = m
				break
			}
		}
	}

	if x.f.Name != "" || x.getter.Name != "" || x.setter.Name != "" {
		return x
	}
	return nil
}

type reflectByField struct {
	src    reflect.Value
	m      meta.Definition
	f      reflect.StructField
	getter reflect.Method
	setter reflect.Method
	opts   NodeOptions
}

func (fdef *reflectByField) elem() reflect.Value {
	if fdef.src.Kind() == reflect.Pointer {
		return fdef.src.Elem()
	}
	return fdef.src
}

func (fdef *reflectByField) clear() error {
	fdef.elem().FieldByIndex(fdef.f.Index).SetZero()
	return nil
}

func (fdef *reflectByField) get() (reflect.Value, error) {
	var v reflect.Value
	if fdef.f.Name != "" {
		v = fdef.elem().FieldByIndex(fdef.f.Index)

		// Important note.  Cannot do anything with a struct, need a pointer
		// to a struct otherwise cannot set any fields so we implicity grab
		// a pointer while we have the source object.
		if v.Kind() == reflect.Struct {
			v = v.Addr()
		}

	} else if fdef.getter.Name != "" {
		input := []reflect.Value{fdef.src}
		resp := fdef.getter.Func.Call(input)
		if len(resp) == 0 {
			return v, fmt.Errorf("%s had no response", fdef.getter.Name)
		}
		v = resp[0]
		if len(resp) == 2 {
			if !isErrType(resp[1]) {
				return v, fmt.Errorf("%s return 2 items and 2nd arg was allowed to be an err type", fdef.getter.Name)
			}
			if !resp[1].IsNil() {
				return v, resp[1].Interface().(error)
			}
		}
	} else {
		return v, fmt.Errorf("%s has no recognized way to get value", fdef.m.Ident())
	}
	if fdef.opts.IgnoreEmpty && reflectIsEmpty(v) {
		return reflect.Value{}, nil
	}
	return v, nil
}

var errorInterface = reflect.TypeOf((*error)(nil)).Elem()

func isErrType(v reflect.Value) bool {
	return v.Type().Implements(errorInterface)
}

func (fdef *reflectByField) set(v reflect.Value) error {
	if fdef.f.Name != "" {
		fdef.elem().FieldByIndex(fdef.f.Index).Set(v)
		return nil
	}
	if fdef.setter.Name != "" {
		input := []reflect.Value{fdef.src, v}
		resp := fdef.setter.Func.Call(input)
		if len(resp) == 0 {
			return nil
		}
		if len(resp) == 1 {
			var err error
			if !isErrType(resp[0]) {
				return fmt.Errorf("%s returns item and seconds was expected to be err", fdef.setter.Name)
			}
			if !resp[0].IsNil() {
				err = resp[0].Interface().(error)
			}
			return err
		}
		return fmt.Errorf("%s expected 0 or 1 items and got %d", fdef.setter.Name, len(resp))
	}
	return fmt.Errorf("%s has no recognized way to set value", fdef.m.Ident())
}

func (fdef *reflectByField) fieldType() reflect.Type {
	if fdef.f.Name != "" {
		return fdef.elem().FieldByIndex(fdef.f.Index).Type()
	}
	if fdef.getter.Name != "" {
		return fdef.getter.Func.Type().Out(0)
	}
	return fdef.setter.Func.Type().In(0)
}
