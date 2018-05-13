package bitgo

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var (
	b        *BitGo
	tbtc     *BitGo
	walletId string
	address  string
)

func init() {
	rand.Seed(time.Now().UnixNano())
	b = New("test", os.Getenv("BITGO_TOKEN"))
	tbtc = b.Coin("tbtc")
}

// List Wallets

func TestListWallets(t *testing.T) {
	list, err := tbtc.ListWallets(&ListParams{
		Limit:     3,
		AllTokens: true,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(list.Wallets) == 0 {
		t.Errorf("wallets not found")
	}

	wallet := list.Wallets[0]

	walletId = wallet.ID
	if walletId == "" {
		t.Errorf("empty wallet ID")
	}
}

// Get Wallet

func TestGetWallet(t *testing.T) {
	wallet, err := tbtc.GetWallet(walletId, &GetWalletParams{
		AllTokens: true,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}

	if wallet.ID == "" {
		t.Errorf("empty wallet ID")
	}

	address = wallet.ReceiveAddress.Address
	if address == "" {
		t.Errorf("empty wallet address")
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
	wallet, err := tbtc.GetWalletByAddress(address)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if wallet.ReceiveAddress.Address != address {
		t.Errorf("")
	}
}
