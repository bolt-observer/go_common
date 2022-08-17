package go_common

type Data struct {
	PubKey            string `json:"pubkey"`
	MacaroonHex       string `json:"macaroon_hex"`
	CertificateBase64 string `json:"certificate_base64,omitempty"`
	Endpoint          string `json:"endpoint"`
	Tags              string `json:"tags,omitempty"`
	ApiType           *int   `json:"api_type,omitempty"`
}
