package utils

import "testing"

func TestValidatePubkey(t *testing.T) {

	if ValidatePubkey("") {
		t.Fatalf("empty pubkey is invalid")
	}
	if ValidatePubkey("222222222222222222222222222222222222222222222222222222222222222222") {
		t.Fatalf("wrong prefix for pubkey")
	}
	if ValidatePubkey("02222222222222222222222222222222222222222222222222222222222222222") {
		t.Fatalf("to short pubkey")
	}
	if ValidatePubkey("02xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx") {
		t.Fatalf("invalid char for pubkey")
	}
	if !ValidatePubkey("022222222222222222222222222222222222222222222222222222222222222222") {
		t.Fatalf("valid pubkey")
	}
}
