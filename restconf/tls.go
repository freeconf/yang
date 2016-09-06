package restconf

import "crypto/tls"

type Tls struct {
	Config tls.Config
	CertFile string
	KeyFile string
	CaCertFile string
}
