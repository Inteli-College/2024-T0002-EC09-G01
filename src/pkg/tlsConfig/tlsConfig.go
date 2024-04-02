package tlsconfig

import (
	"crypto/tls"
	"crypto/x509"
	"os"
)

func NewTLSConfig() *tls.Config {
	certpool := x509.NewCertPool()

	pemCerts, err := os.ReadFile("../../config/intermediate.cert.pem")
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	cert, err := tls.LoadX509KeyPair("../../config/mqtt-client.cert.pem", "../../config/mqtt-client.key.pem")

	return &tls.Config{
		RootCAs:    certpool,
		ClientAuth: tls.NoClientCert,
		ClientCAs:  nil,
		InsecureSkipVerify: true,
		Certificates: []tls.Certificate{cert},
	}
}
