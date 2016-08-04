package restconf

import (
	"github.com/c2g/node"
	"github.com/c2g/meta"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

type Api struct {
}

func (self Api) Manage(service *Service) node.Node {
	s := &node.MyNode{Peekable: service}
	s.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "callHome":
			if r.New {
				service.CallHome = &CallHome{
					EndpointAddress: service.EffectiveCallbackAddress(),
					Module: service.Root.Meta.(*meta.Module),
				}
			}
			if service.CallHome != nil {
				return service.CallHome.Manage(), nil
			}
		case "webSocket":
			return node.MarshalContainer(service.socketHandler), nil
		case "tls":
			if r.New {
				service.Tls = &tls.Config{}
			}
			if service.Tls != nil {
				return self.Tls(service.Tls), nil
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

func (self Api) Tls(config *tls.Config) node.Node {
	return &node.Extend{
		Node: node.MarshalContainer(config),
		OnSelect : func(p node.Node, r node.ContainerRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "ca":
				if r.New {
					config.RootCAs = x509.NewCertPool()
				}
				if config.RootCAs != nil {
					return self.CertificateAuthority(config.RootCAs), nil
				}
			case "cert":
				if r.New {
					config.Certificates = make([]tls.Certificate, 1)
				}
				if len(config.Certificates) > 0 {
					return self.Certificate(&config.Certificates[0]), nil
				}
			}
			return p.Select(r)
		},
	}
}

func (self Api) CertificateAuthority(pool *x509.CertPool) node.Node {
	n := &node.MyNode{}
	n.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) error {
		switch r.Meta.GetIdent() {
		case "certFile":
			if r.Write {
				pemData, err := ioutil.ReadFile(hnd.Val.Str)
				if err != nil {
					return err
				}
				pool.AppendCertsFromPEM(pemData)
			}
		}
		return nil
	}
	return n
}

func (self Api) Certificate(cert *tls.Certificate) node.Node {
	n := &node.MyNode{}
	var certFile string
	var keyFile string
	n.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) (err error) {
		if r.Write {
			switch r.Meta.GetIdent() {
			case "certFile":
				certFile = hnd.Val.Str
			case "keyFile":
				keyFile = hnd.Val.Str
			}
		}
		// else nop = not readable back
		return nil
	}
	n.OnEvent = func(sel *node.Selection, e node.Event) error {
		switch e.Type {
		case node.NEW:
			var err error
			*cert, err = tls.LoadX509KeyPair(certFile, keyFile)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return n
}
