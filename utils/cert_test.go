package utils

import (
	"testing"
)

func TestObtainCert(t *testing.T) {
	cert := ObtainCert("www.arnes.si:443")
	if cert == "" {
		t.Fatalf("Cert could not be obtained")
		return
	}
}
