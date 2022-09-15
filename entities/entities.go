package entities

import (
	"strconv"
	"strings"
	"time"
)

type Data struct {
	PubKey            string `json:"pubkey"`
	MacaroonHex       string `json:"macaroon_hex"`
	CertificateBase64 string `json:"certificate_base64,omitempty"`
	Endpoint          string `json:"endpoint"`
	Tags              string `json:"tags,omitempty"`
	ApiType           *int   `json:"api_type,omitempty"`
}

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *JsonTime) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q, 0)
	return
}
