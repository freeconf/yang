package restconf

import (
	"github.com/c2g/node"
	"github.com/c2g/meta"
)

func (service *Service) Manage() node.Node {
	s := &node.MyNode{Peekables: map[string]interface{}{"internal": service}}
	s.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "callHome":
			if r.New {
				service.CallHome = &CallHome{
					EndpointAddress: service.EffectiveCallbackAddress(),
					Module: service.Root().Meta().(*meta.Module),
				}
			}
			if service.CallHome != nil {
				return service.CallHome.Manage(), nil
			}
		case "webSocket":
			return node.MarshalContainer(service.socketHandler), nil
		}
		return nil, nil
	}
	s.OnRead = func(r node.FieldRequest) (*node.Value, error) {
		return node.ReadField(r.Meta, service)
	}
	s.OnWrite = func(r node.FieldRequest, v *node.Value) (err error) {
		switch r.Meta.GetIdent() {
		case "docRoot":
			service.DocRoot = v.Str
			service.SetDocRoot(&meta.FileStreamSource{Root: service.DocRoot})
		}
		return node.WriteField(r.Meta, service, v)
	}
	return s
}

