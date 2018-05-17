package bitgo

import (
	"testing"
)

// List Wallet Transfers

func TestListWalletTransfers(t *testing.T) {
	coin, params := getTestCoin(t)

	_, err := coin.ListWalletTransfers(params.WalletId, &ListParams{
		Limit: 5,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Get Wallet Transfer

func TestGetWalletTransfer(t *testing.T) {
	coin, params := getTestCoin(t)

	_, err := coin.GetWalletTransfer(params.WalletId, params.TransferId)
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Get Wallet Transfer By Sequence ID

/*

func TestGetWalletTransferBySequenceID(t *testing.T) {
	// TODO
	_, err := coin.GetWalletTransferBySequenceID(walletId, sequenceId)
	if err != nil {
		t.Fatal(err.Error())
	}
}

*/

// Create Wallet Address

func TestCreateWalletAddress(t *testing.T) {
	coin, params := getTestCoin(t)

	newLabel := randStringRunes(5)

	_, err := coin.CreateWalletAddress(params.WalletId, &AddressParams{
		Chain: 0,
		Label: newLabel,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Get Wallet Address

func TestGetWalletAddress(t *testing.T) {
	coin, params := getTestCoin(t)

	_, err := coin.GetWalletAddress(params.WalletId, params.Address)
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Update Wallet Address

func TestUpdateWalletAddress(t *testing.T) {
	coin, params := getTestCoin(t)

	newLabel := randStringRunes(5)

	_, err := coin.UpdateWalletAddress(params.WalletId, params.Address, UpdateWalletAddressParams{
		Label: newLabel,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Send Transaction

func TestSendTransaction(t *testing.T) {
	coin, params := getTestCoin(t)

	err := coin.Unlock(UnlockParams{
		Otp: "0000000",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = coin.SendTransaction(params.WalletId, SendParams{
		Address:          params.Address,
		Amount:           0.0001 * 1e8,
		WalletPassphrase: params.Passphrase,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Send Transaction to Many

func TestSendTransactionToMany(t *testing.T) {
	coin, params := getTestCoin(t)

	err := coin.Unlock(UnlockParams{
		Otp: "0000000",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = coin.SendTransactionToMany(params.WalletId, SendToManyParams{
		Recipients: []Recipient{
			Recipient{
				Address: params.Address,
				Amount:  0.0001 * 1e8,
			},
		},
		WalletPassphrase: params.Passphrase,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
