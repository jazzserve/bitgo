package bitgo

import (
	"testing"
)

// List Wallets

func TestListWallets(t *testing.T) {
	_, err := tbtc.ListWallets(&ListParams{
		Limit:     3,
		AllTokens: true,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
}

// Get Wallet

func TestGetWallet(t *testing.T) {
	_, err := tbtc.GetWallet(walletId, &GetWalletParams{
		AllTokens: true,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
}

// Update Wallet

func TestUpdateWallet(t *testing.T) {
	newLabel := randStringRunes(5)

	wallet, err := tbtc.UpdateWallet(walletId, UpdateWalletParams{
		Label: newLabel,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}

	if wallet.Label != newLabel {
		t.Errorf("update label %s", newLabel)
	}
}

// Get Wallet By Address

func TestGetWalletByAddress(t *testing.T) {
	_, err := tbtc.GetWalletByAddress(address)
	if err != nil {
		t.Error(err.Error())
		return
	}
}
