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
