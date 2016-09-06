package restconf

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"io/ioutil"
)

func ServiceNode(service *Service) node.Node {
	s := &node.MyNode{Peekable: service}
	s.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "callHome":
			if r.New {
				service.CallHome = &CallHome{
					EndpointAddress: service.EffectiveCallbackAddress(),
					Module:          service.Root.Meta.(*meta.Module),
					ClientSource:    service,
				}
			}
			if service.CallHome != nil {
				return service.CallHome.Manage(), nil
			}
		case "webSocket":
			return node.MarshalContainer(service.socketHandler), nil
		case "tls":
			if r.New {
				service.Tls = &Tls{}
			}
			if service.Tls != nil {
				return TlsNode(service.Tls), nil
			}
		}
		return nil, nil
	}
	s.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		if r.Write {
			switch r.Meta.GetIdent() {
			case "docRoot":
				service.DocRoot = hnd.Val.Str
				service.SetDocRoot(&meta.FileStreamSource{Root: service.DocRoot})
			default:
				err = node.WriteField(r.Meta, service, hnd.Val)
			}
		} else {
			hnd.Val, err = node.ReadField(r.Meta, service)
		}
		return
	}
	return s
}

func TlsNode(config *Tls) node.Node {
	return &node.Extend{
		Node: node.MarshalContainer(&config.Config),
		OnSelect: func(p node.Node, r node.ContainerRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "ca":
				if r.New {
					config.Config.RootCAs = x509.NewCertPool()

					// assertion - harmless if not used, but useful if is used.
					config.Config.ClientCAs = config.Config.RootCAs
					config.Config.ClientAuth = tls.VerifyClientCertIfGiven
				}
				if config.Config.RootCAs != nil {
					return CertificateAuthorityNode(config), nil
				}
			case "cert":
				if r.New {
					config.Config.Certificates = make([]tls.Certificate, 1)
				}
				if len(config.Config.Certificates) > 0 {
					return CertificateNode(config), nil
				}
			}
			return p.Select(r)
		},
	}
}

func CertificateAuthorityNode(config *Tls) node.Node {
	n := &node.MyNode{}
	n.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) error {
		switch r.Meta.GetIdent() {
		case "certFile":
			if r.Write {
				config.CaCertFile = hnd.Val.Str
				pemData, err := ioutil.ReadFile(hnd.Val.Str)
				if err != nil {
					return err
				}
				config.Config.RootCAs.AppendCertsFromPEM(pemData)
			} else {
				hnd.Val = &node.Value{Str :config.CaCertFile}
			}
		}
		return nil
	}
	return n
}

func CertificateNode(config *Tls) node.Node {
	n := &node.MyNode{}
	n.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		switch r.Meta.GetIdent() {
		case "certFile":
			if r.Write {
				config.CertFile = hnd.Val.Str
			} else {
				hnd.Val = &node.Value{Str:config.CertFile}
			}
		case "keyFile":
			if r.Write {
				config.KeyFile = hnd.Val.Str
			} else {
				hnd.Val = &node.Value{Str:config.KeyFile}
			}
		}
		return nil
	}
	n.OnEvent = func(sel node.Selection, e node.Event) error {
		switch e.Type {
		case node.NEW:
			var err error
			config.Config.Certificates[0], err = tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return n
}
