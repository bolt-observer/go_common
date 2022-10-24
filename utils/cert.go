package utils

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/pem"
	"strings"
	"time"

	"github.com/golang/glog"
)

func ObtainCert(endpoint string) string {
	conf := &tls.Config{
		MinVersion:         tls.VersionTLS11,
		InsecureSkipVerify: true, // on purpose
	}

	split := strings.Split(endpoint, ":")
	if len(split) <= 2 {
		conf.ServerName = split[0]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	d := tls.Dialer{
		Config: conf,
	}

	conn, err := d.DialContext(ctx, "tcp", endpoint)
	cancel()
	if err != nil {
		glog.Errorf("Could not connect to endpoint: %v", err)
		return ""
	}

	defer conn.Close()

	tlsConn := conn.(*tls.Conn)
	cert := tlsConn.ConnectionState().PeerCertificates[0]

	result := string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}))

	return base64.StdEncoding.EncodeToString([]byte(result))
}
