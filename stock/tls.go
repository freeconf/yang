package stock

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"
)

type Tls struct {
	Config     tls.Config
	CertFile   string
	KeyFile    string
	CaCertFile string
}

func TlsNode(config *Tls) node.Node {
	return &node.Extend{
		Node: node.ReflectNode(&config.Config),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
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
			return p.Child(r)
		},
	}
}

func CertificateAuthorityNode(config *Tls) node.Node {
	n := &node.MyNode{}
	n.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) error {
		switch r.Meta.GetIdent() {
		case "certFile":
			if r.Write {
				config.CaCertFile = hnd.Val.String()
				pemData, err := ioutil.ReadFile(hnd.Val.String())
				if err != nil {
					return err
				}
				config.Config.RootCAs.AppendCertsFromPEM(pemData)
			} else {
				hnd.Val = val.String(config.CaCertFile)
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
				config.CertFile = hnd.Val.String()
			} else {
				hnd.Val = val.String(config.CertFile)
			}
		case "keyFile":
			if r.Write {
				config.KeyFile = hnd.Val.String()
			} else {
				hnd.Val = val.String(config.KeyFile)
			}
		}
		return nil
	}
	n.OnEndEdit = func(r node.NodeRequest) error {
		var err error
		if r.New {
			config.Config.Certificates[0], err = tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
		}
		return err
	}
	return n
}
