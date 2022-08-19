package utils

import (
	"regexp"
	"strings"
)

var (
	HexCharacters = regexp.MustCompile(`^[A-Fa-f0-9]*$`)
	AlphaNumeric  = regexp.MustCompile(`^[A-Za-z0-9_-]*$`)
)

const (
	PUBKEY_LEN = 66
)

func ValidatePubkey(pubkey string) bool {
	if len(pubkey) != PUBKEY_LEN {
		return false
	}

	if !strings.HasPrefix(pubkey, "02") && !strings.HasPrefix(pubkey, "03") {
		return false
	}

	return HexCharacters.MatchString(pubkey)
}
