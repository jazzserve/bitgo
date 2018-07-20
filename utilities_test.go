package bitgo

import (
	"testing"

	"fmt"
)

// Estimate Transaction Fees

func TestEstimateTransactionFees(t *testing.T) {
	coin, _ := getTestCoin(t)

	a, err := coin.EstimateTransactionFees(2)
	if err != nil {
		t.Fatal(err.Error())
	}
}
