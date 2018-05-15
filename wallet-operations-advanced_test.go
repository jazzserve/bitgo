package bitgo

import (
	"testing"
)

// Get Total Balances

func TestGetTotalBalances(t *testing.T) {
	_, err := coin.GetTotalBalances()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
