package secure

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io"
	"math/big"
	"time"
)

type Generator struct {
	Country      string
	Organization string
}

type Cert struct {
	PrivateKey *rsa.PrivateKey
	Cert       *x509.Certificate
	Raw        []byte
}

func Decode(inKey io.Reader, inCert io.Reader) (*Cert, error) {
	return nil, nil
}

func (self *Cert) EncodeCert(out io.Writer) error {
	return pem.Encode(out, &pem.Block{Type: "CERTIFICATE", Bytes: self.Raw})
}

func (self *Cert) EncodeKey(out io.Writer) error {
	raw := x509.MarshalPKCS1PrivateKey(self.PrivateKey)
	return pem.Encode(out, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: raw})
}

func (self *Generator) CA() (*Cert, error) {
	pvtKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	pubKey := &pvtKey.PublicKey

	template := self.template()
	template.IsCA = true
	template.KeyUsage |= x509.KeyUsageCertSign

	// self-signed
	parent := template

	raw, err := x509.CreateCertificate(rand.Reader, template, parent, pubKey, pvtKey)
	if err != nil {
		return nil, err
	}
	return &Cert{
		PrivateKey: pvtKey,
		Cert:       template,
		Raw:        raw,
	}, nil
}

func (self *Generator) template() *x509.Certificate {
	return &x509.Certificate{
		BasicConstraintsValid: true,
		SubjectKeyId:          []byte{1, 2, 3},
		SerialNumber:          big.NewInt(1234),
		Subject: pkix.Name{
			Country:      []string{self.Country},
			Organization: []string{self.Organization},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(5, 0, 0),
		KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
	}
}

func (self *Generator) Cert(parent *Cert) (*Cert, error) {
	pvtKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	pubKey := &pvtKey.PublicKey
	template := self.template()
	template.ExtKeyUsage = append(template.ExtKeyUsage, x509.ExtKeyUsageClientAuth)

	raw, err := x509.CreateCertificate(rand.Reader, template, parent.Cert, pubKey, pvtKey)
	if err != nil {
		return nil, err
	}
	return &Cert{
		PrivateKey: pvtKey,
		Cert:       template,
		Raw:        raw,
	}, nil
}
