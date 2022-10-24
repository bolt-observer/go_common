package utils

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"testing"
)

func TestObtainCert(t *testing.T) {
	cert := ObtainCert("www.arnes.si:443")
	if cert == "" {
		t.Fatalf("Cert could not be obtained")
		return
	}
}

func TestObtainCertSNI(t *testing.T) {
	const HOST = "btcpay.skunkworks.si"
	cert := ObtainCert(fmt.Sprintf("%s:443", HOST))
	if cert == "" {
		t.Fatalf("Cert could not be obtained")
		return
	}

	actual, err := base64.StdEncoding.DecodeString(cert)

	if err != nil {
		t.Fatalf("Error decoding certificate: %v", err)
		return
	}

	decoded, _ := pem.Decode(actual)

	x, err := x509.ParseCertificate(decoded.Bytes)
	if err != nil {
		t.Fatalf("Error parsing certificate: %v", err)
		return
	}

	if x.Subject.CommonName != HOST {
		t.Fatalf("Common name was %s not %s\n", x.Subject.CommonName, HOST)
		return
	}
}
