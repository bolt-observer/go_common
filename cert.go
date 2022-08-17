package go_common

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/pem"

	"github.com/golang/glog"
)

func ObtainCert(endpoint string) string {
	conf := &tls.Config{
		InsecureSkipVerify: true, // on purpose
	}

	conn, err := tls.Dial("tcp", endpoint, conf)
	if err != nil {
		glog.Errorf("Could not connect to endpoint: %v", err)
		return ""
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]

	result := string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}))

	return base64.StdEncoding.EncodeToString([]byte(result))
}
