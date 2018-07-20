package bitgo

import (
	"testing"
)

// Get Total Balances

func TestGetTotalBalances(t *testing.T) {
	coin, _ := getTestCoin(t)

	_, err := coin.GetTotalBalances()
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Build Transaction

func TestBuildTransaction(t *testing.T) {
	coin, params := getTestCoin(t)

	var recipients []Recipient
	recipients = append(recipients, Recipient{
		Address: params.Address,
		Amount:  0.0001 * 1e8,
	})

	_, err := coin.BuildTransaction(params.WalletId, BuildTransactionParams{Recipients: recipients})
	if err != nil {
		t.Fatal(err.Error())
	}
}
