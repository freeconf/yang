package nodeutil

import (
	"fmt"
	"reflect"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
)

type actionHandler struct {
	src  reflect.Value
	rpc  *meta.Rpc
	opts NodeOptions
	m    reflect.Method
}

func newActionHandler(src reflect.Value, rpc *meta.Rpc, opts NodeOptions) (*actionHandler, error) {
	def := &actionHandler{
		src:  src,
		rpc:  rpc,
		opts: opts,
	}
	candidate := def.opts.Ident
	if candidate == "" {
		candidate = MetaNameToFieldName(def.rpc.Ident())
	}
	var found bool
	def.m, found = def.src.Type().MethodByName(candidate)
	if !found {
		return nil, fmt.Errorf("could not find method '%s' in %s for action", candidate, def.src.Type())
	}
	return def, nil
}

func (def *actionHandler) do(n *Node, m *meta.Rpc, in *node.Selection) (node.Node, error) {
	inVal, err := def.newInput(in != nil)
	if err != nil {
		return nil, err
	}
	if in != nil && m.Input() != nil && inVal.IsValid() {
		inNode, err := n.New(m.Input(), inVal.Interface())
		if err != nil {
			return nil, err
		}
		if err = in.UpsertInto(inNode); err != nil {
			return nil, err
		}
	}
	respVal, err := def.invoke(inVal)
	if err != nil {
		return nil, err
	}

	if respVal.IsValid() {
		return n.New(m.Output(), respVal.Interface())
	}

	return nil, nil
}

func (def *actionHandler) invoke(in reflect.Value) (reflect.Value, error) {
	var empty reflect.Value

	// encode input
	inArgs := []reflect.Value{def.src} // starts with reciever (i.e. self aka this)
	if in.IsValid() {
		if def.opts.ActionInputExploded {
			args := in.Interface().(map[string]interface{})
			for _, argMeta := range def.rpc.Input().DataDefinitions() {
				if argVal, found := args[argMeta.Ident()]; found {
					// TODO: candidate for coersion
					inArgs = append(inArgs, reflect.ValueOf(argVal))
				} else {
					// no argument given so fill in with zero value
					inArgs = append(inArgs, reflect.Value{})
				}
			}
		} else {
			inArgs = append(inArgs, replacePointerWithStruct(def.m.Type.In(1), in))
		}
	}

	out := def.m.Func.Call(inArgs)

	// decode output
	goOut := len(out)
	var err error
	var outVal reflect.Value
	if goOut > 0 {
		lastVal := out[goOut-1]
		lastIsErr := isErrType(lastVal)
		if def.rpc.Output() != nil {
			yangOut := len(def.rpc.Output().DataDefinitions())
			if def.opts.ActionOutputExploded {
				if goOut != yangOut {
					if (goOut-1) == yangOut && lastIsErr {
						if !lastVal.IsNil() {
							return empty, lastVal.Interface().(error)
						}
					} else {
						return empty, fmt.Errorf("%s.%s expected %d outputs with optional error output as well and instead found %d outputs, you might want to consider using non exploded output option", def.src.Type(), def.m.Name, yangOut, goOut)
					}
				}
				args := make(map[string]any, yangOut)
				for i, outMeta := range def.rpc.Output().DataDefinitions() {
					if out[i].IsValid() {
						if canNil(out[i].Type().Kind()) && out[i].IsNil() {
							continue
						}
						out := replaceStructWithPointer(out[i])
						args[outMeta.Ident()] = out.Interface()
					}
				}
				outVal = reflect.ValueOf(args)
			} else {
				if goOut != 1 {
					if goOut == 2 && lastIsErr {
						if !lastVal.IsNil() {
							return empty, lastVal.Interface().(error)
						}
					} else {
						return empty, fmt.Errorf("%s.%s expected a single output with optional error output as well and instead found %d outputs, you might want to consider using exploded output option", def.src.Type(), def.m.Name, goOut)
					}
				}
				outVal = replaceStructWithPointer(out[0])
			}
		} else if lastIsErr && !lastVal.IsNil() {
			return empty, lastVal.Interface().(error)
		}
	}

	return outVal, err
}

// If method returns a struct, allocate a new object and copy in struct
// contents as structs value types are not used/useful in this nodeutil
// module.
func replaceStructWithPointer(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Struct {
		ptr := reflect.New(v.Type())
		ptr.Elem().Set(v)
		return ptr
	}
	return v
}

func replacePointerWithStruct(t reflect.Type, v reflect.Value) reflect.Value {
	if t.Kind() == reflect.Struct {
		return v.Elem()
	}
	return v
}

func (def *actionHandler) newInput(hasAnyInput bool) (reflect.Value, error) {
	var empty reflect.Value
	goIn := def.m.Type.NumIn() - 1 // substract 1 for the receiver
	yangIn := 0
	if def.rpc.Input() != nil {
		yangIn = len(def.rpc.Input().DataDefinitions())
	}
	if goIn == 0 && yangIn == 0 { // allow exploded input aligning with yang input (or nil on both)
		return empty, nil
	}
	if def.opts.ActionInputExploded {
		if len(def.rpc.Input().DataDefinitions()) != goIn {
			return empty, fmt.Errorf("%s.%s found %d params but %d defined", def.src.Type(), def.m.Name, goIn, yangIn)
		}
		return reflect.ValueOf(make(map[string]interface{})), nil
	}
	if goIn != 1 {
		return empty, fmt.Errorf("%s.%s requires a single struct to hold all input parameters but %d were found", def.src.Type(), def.m.Name, goIn)
	}

	t0 := def.m.Type.In(1) // 0 is the receiever
	isPtr := t0.Kind() == reflect.Pointer
	kind := t0.Kind()
	if isPtr {
		kind = t0.Elem().Kind()
	}
	if isPtr && !hasAnyInput {
		return empty, nil
	}
	switch kind {
	case reflect.Struct:
		if isPtr {
			return reflect.New(t0.Elem()), nil
		}
		return reflect.New(t0), nil
	case reflect.Map:
		v := reflect.MakeMap(t0)
		return v, nil
	}
	return empty, fmt.Errorf("%s.%s input parameters needs to be a struct or a map, instead found %s", def.src.Type(), def.m.Name, t0)
}
