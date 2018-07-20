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

// Build Transaction

type BuildTransactionParams struct {
	Recipients []Recipient `json:"recipients"`
}

type Transaction struct {
	TxHex  string `json:"txHex"`
	TxInfo struct {
		NP2SHInputs   int `json:"nP2SHInputs"`
		NSegwitInputs int `json:"nSegwitInputs"`
		NOutputs      int `json:"nOutputs"`
		Unspents      []struct {
			Chain        int    `json:"chain"`
			Index        int    `json:"index"`
			RedeemScript string `json:"redeemScript"`
			ID           string `json:"id"`
			Address      string `json:"address"`
			Value        int    `json:"value"`
		} `json:"unspents"`
		ChangeAddresses []string `json:"changeAddresses"`
	} `json:"txInfo"`
	FeeInfo struct {
		Size           int    `json:"size"`
		Fee            int    `json:"fee"`
		FeeRate        int    `json:"feeRate"`
		PayGoFee       int    `json:"payGoFee"`
		PayGoFeeString string `json:"payGoFeeString"`
	} `json:"feeInfo"`
}

func (b *BitGo) BuildTransaction(walletId string, params BuildTransactionParams) (transaction Transaction, err error) {
	err = b.post(
		fmt.Sprintf("%s/wallet/%s/tx/build",
			b.coin,
			walletId),
		params,
		&transaction)
	return
}

// Sign Transaction

type Keychain struct {
	EncryptedPrv string `json:"encryptedPrv"`
}

type SignTransactionParams struct {
	TxPrebuild       Transaction `json:"txPrebuild"`
	Keychain         Keychain    `json:"keychain"`
	WalletPassphrase string      `json:"walletPassphrase"`
}

type SignedTransaction struct {
	TxHex string `json:"txHex"`
}

func (b *BitGo) SignTransaction(walletId string, params SignTransactionParams) (signedTransaction SignedTransaction, err error) {
	err = b.post(
		fmt.Sprintf("%s/wallet/%s/signtx",
			b.coin,
			walletId),
		params,
		&signedTransaction)
	return
}

// Send Transaction

// TODO
