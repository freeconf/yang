package conf

import "github.com/c2stack/c2g/node"
import "context"
import "github.com/c2stack/c2g/c2"

type ProxyContextKey int

const RemoteIpAddressKey ProxyContextKey = 0

func ProxyNode(proxy *Proxy) node.Node {
	return &node.MyNode{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			return nil, nil
		},
		OnAction: func(r node.ActionRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "register":
				reg, err := registrationRequest(r.Context, r.Input)
				if err != nil {
					return nil, err
				}
				if reg.Address == "" {
					reg.Address = r.Context.Value(RemoteIpAddressKey).(string)
				}
				c2.Debug.Printf("register %v", reg)
				return nil, proxy.Register(reg.Id, reg.Address, reg.Port)
			}
			return nil, nil
		},
	}
}

type RegistrationRequest struct {
	Address string
	Port    string
	Id      string
}

func registrationRequest(c context.Context, s node.Selection) (RegistrationRequest, error) {
	var reg RegistrationRequest
	if err := s.InsertInto(c, node.ReflectNode(&reg)).LastErr; err != nil {
		return reg, err
	}
	return reg, nil
}
