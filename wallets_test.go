package bitgo

import (
	"testing"
)

// List Wallets

func TestListWallets(t *testing.T) {
	coin, _ := getTestCoin(t)

	_, err := coin.ListWallets(&ListParams{
		Limit:     3,
		AllTokens: true,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Generate Wallet

func TestGenerateWallet(t *testing.T) {
	coin, params := getTestCoin(t)

	newLabel := randStringRunes(5)

	_, err := coin.GenerateWallet(GenerateWalletParams{
		Label:      newLabel,
		Passphrase: params.Passphrase,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Get Wallet

func TestGetWallet(t *testing.T) {
	coin, params := getTestCoin(t)

	_, err := coin.GetWallet(params.WalletId, &GetWalletParams{
		AllTokens: true,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Update Wallet

func TestUpdateWallet(t *testing.T) {
	coin, params := getTestCoin(t)

	newLabel := randStringRunes(5)

	wallet, err := coin.UpdateWallet(params.WalletId, UpdateWalletParams{
		Label: newLabel,
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	if wallet.Label != newLabel {
		t.Errorf("update label %s", newLabel)
	}
}

// Get Wallet By Address

func TestGetWalletByAddress(t *testing.T) {
	coin, params := getTestCoin(t)

	_, err := coin.GetWalletByAddress(params.Address)
	if err != nil {
		t.Fatal(err.Error())
	}
}
