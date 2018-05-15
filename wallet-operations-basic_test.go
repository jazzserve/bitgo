package bitgo

import (
	"testing"
)

// List Wallet Transfers

func TestListWalletTransfers(t *testing.T) {
	_, err := coin.ListWalletTransfers(walletId, &ListParams{
		Limit: 5,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
}

// Get Wallet Transfer

func TestGetWalletTransfer(t *testing.T) {
	_, err := coin.GetWalletTransfer(walletId, transferId)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

// Get Wallet Transfer By Sequence ID

func TestGetWalletTransferBySequenceID(t *testing.T) {
	// TODO
	/*
		_, err := coin.GetWalletTransferBySequenceID(walletId, sequenceId)
		if err != nil {
			t.Error(err.Error())
			return
		}
	*/
}

// Create Wallet Address

func TestCreateWalletAddress(t *testing.T) {
	newLabel := randStringRunes(5)

	_, err := coin.CreateWalletAddress(walletId, &AddressParams{
		Chain: 0,
		Label: newLabel,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
}

// Get Wallet Address

func TestGetWalletAddress(t *testing.T) {
	_, err := coin.GetWalletAddress(walletId, address)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

// Update Wallet Address

func TestUpdateWalletAddress(t *testing.T) {
	newLabel := randStringRunes(5)

	_, err := coin.UpdateWalletAddress(walletId, address, UpdateWalletAddressParams{
		Label: newLabel,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
}

// Send Transaction

// TODO

/*

func TestSendTransaction(t *testing.T) {
	_, err := coin.SendTransaction(walletId, SendParams{
		Address:          address,
		Amount:           0.01 * 1e8,
		WalletPassphrase: passphrase,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
}

*/

// Send Transaction to Many

// TODO
