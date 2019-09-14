package secure

import (
	"crypto/x509"
	"fmt"

	"github.com/freeconf/yang/stock"
)

type CertHandler struct {
	Authority *stock.Tls
}

func (self *CertHandler) VerifyRequest(certs []*x509.Certificate) error {
	// certs := r.TLS.PeerCertificates
	opts := x509.VerifyOptions{
		Roots: self.Authority.Config.RootCAs,
	}
	var err error
	var valid *x509.Certificate
	for _, cert := range certs {
		if _, invalid := cert.Verify(opts); invalid != nil {
			err = invalid
		} else {
			valid = cert
		}
	}

	if err != nil {
		return err
	}
	fmt.Printf("Valid! %v", valid)

	return nil
}
