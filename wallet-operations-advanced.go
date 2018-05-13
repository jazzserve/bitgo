package bitgo

import (
	"fmt"
)

type TotalBalances struct {
	Balance                int    `json:"balance"`
	BalanceString          string `json:"balanceString"`
	ConfirmedBalance       int    `json:"confirmedBalance"`
	ConfirmedBalanceString string `json:"confirmedBalanceString"`
	SpendableBalance       int    `json:"spendableBalance"`
	SpendableBalanceString string `json:"spendableBalanceString"`
}

// Get Total Balances

func (b *BitGo) GetTotalBalances() (total TotalBalances, err error) {
	err = b.get(
		fmt.Sprintf("%s/wallet/balances",
			b.coin),
		nil,
		&total)
	return
}

// TODO
