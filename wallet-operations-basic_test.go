package bitgo

import (
	"testing"
)

// Create Wallet Address

func TestCreateWalletAddress(t *testing.T) {
	newLabel := randStringRunes(5)

	_, err := tbtc.CreateWalletAddress(walletId, &AddressParams{
		Chain: 0,
		Label: newLabel,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
}
