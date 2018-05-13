package bitgo

import (
	"testing"
)

// Get Total Balances

func TestGetTotalBalances(t *testing.T) {
	total, err := tbtc.GetTotalBalances()
	if err != nil {
		t.Error(err.Error())
		return
	}

	if total.Balance == 0 {
		t.Errorf("balance = 0")
	}
}
