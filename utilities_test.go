package bitgo

import (
	"testing"
)

// Estimate Transaction Fees

func TestEstimateTransactionFees(t *testing.T) {
	coin, _ := getTestCoin(t)

	_, err := coin.EstimateTransactionFees(2)
	if err != nil {
		t.Fatal(err.Error())
	}
}
