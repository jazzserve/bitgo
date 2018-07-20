package bitgo

import (
	"fmt"
)

// Estimate Transaction Fees

type TransactionFees struct {
	FeePerKb  int `json:"feePerKb"`
	NumBlocks int `json:"numBlocks"`
}

func (b *BitGo) EstimateTransactionFees(numBlocks int) (transactionFees TransactionFees, err error) {
	params := struct {
		NumBlocks int `json:"numBlocks"`
	}{
		NumBlocks: numBlocks,
	}

	err = b.get(
		fmt.Sprintf("%s/tx/fee",
			b.coin),
		params,
		&transactionFees)
	return
}

// TODO
