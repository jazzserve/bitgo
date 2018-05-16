package bitgo

//	"testing"

/*

// List Wallet Transfers

func TestListWalletTransfers(t *testing.T) {
	_, err := coin.ListWalletTransfers(walletId, &ListParams{
		Limit: 5,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Get Wallet Transfer

func TestGetWalletTransfer(t *testing.T) {
	_, err := coin.GetWalletTransfer(walletId, transferId)
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Get Wallet Transfer By Sequence ID

func TestGetWalletTransferBySequenceID(t *testing.T) {
	// TODO
		_, err := coin.GetWalletTransferBySequenceID(walletId, sequenceId)
		if err != nil {
			t.Fatal(err.Error())
		}
}

// Create Wallet Address

func TestCreateWalletAddress(t *testing.T) {
	newLabel := randStringRunes(5)

	_, err := coin.CreateWalletAddress(walletId, &AddressParams{
		Chain: 0,
		Label: newLabel,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Get Wallet Address

func TestGetWalletAddress(t *testing.T) {
	_, err := coin.GetWalletAddress(walletId, address)
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Update Wallet Address

func TestUpdateWalletAddress(t *testing.T) {
	newLabel := randStringRunes(5)

	_, err := coin.UpdateWalletAddress(walletId, address, UpdateWalletAddressParams{
		Label: newLabel,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Send Transaction

func TestSendTransaction(t *testing.T) {
	err := b.Unlock(UnlockParams{
		Otp: "0000000",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = coin.SendTransaction(walletId, SendParams{
		Address:          address,
		Amount:           0.0001 * 1e8,
		WalletPassphrase: passphrase,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Send Transaction to Many

func TestSendTransactionToMany(t *testing.T) {
	err := b.Unlock(UnlockParams{
		Otp: "0000000",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = coin.SendTransactionToMany(walletId, SendToManyParams{
		Recipients: []Recipient{
			Recipient{
				Address: address,
				Amount:  0.0001 * 1e8,
			},
		},
		WalletPassphrase: passphrase,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

*/
