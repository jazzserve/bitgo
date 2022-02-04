package bitgo

import (
	"fmt"
	"time"
)

type Transfer struct {
	ID            string    `json:"id"`
	Coin          string    `json:"coin"`
	Wallet        string    `json:"wallet"`
	Txid          string    `json:"txid"`
	Height        int       `json:"height"`
	Date          time.Time `json:"date"`
	Confirmations int       `json:"confirmations"`
	Value         int       `json:"value"`
	BitgoFee      int       `json:"bitgoFee"`
	Usd           float64   `json:"usd"`
	UsdRate       float64   `json:"usdRate"`
	State         string    `json:"state"`
	VSize         int       `json:"vSize"`
	NSegwitInputs int       `json:"nSegwitInputs"`
	Tags          []string  `json:"tags"`
	SequenceID    string    `json:"sequenceId"`
	History       []struct {
		Date   time.Time `json:"date"`
		Action string    `json:"action"`
	} `json:"history"`
	Entries []struct {
		Address string `json:"address"`
		Value   int    `json:"value"`
		Wallet  string `json:"wallet,omitempty"`
	} `json:"entries"`
	Outputs []struct {
		ID          string `json:"id"`
		Address     string `json:"address"`
		Value       int    `json:"value"`
		ValueString string `json:"valueString"`
		Wallet      string `json:"wallet"`
		Chain       int    `json:"chain"`
		Index       int    `json:"index"`
	} `json:"outputs"`
	Inputs []struct {
		ID          string `json:"id"`
		Address     string `json:"address"`
		Value       int    `json:"value"`
		ValueString string `json:"valueString"`
		Wallet      string `json:"wallet"`
		Chain       int    `json:"chain"`
		Index       int    `json:"index"`
	} `json:"inputs"`
	ConfirmedTime time.Time `json:"confirmedTime"`
	CreatedTime   time.Time `json:"createdTime"`
}

type TransferList struct {
	NextBatchPrevId string     `json:"nextBatchPrevId"`
	Coin            string     `json:"coin"`
	Count           int        `json:"count"`
	Transfers       []Transfer `json:"transfers"`
}

type Address struct {
	ID           string `json:"id"`
	Address      string `json:"address"`
	Chain        int    `json:"chain"`
	Index        int    `json:"index"`
	Coin         string `json:"coin"`
	Wallet       string `json:"wallet"`
	CoinSpecific struct {
		RedeemScript string `json:"redeemScript"`
	} `json:"coinSpecific"`
	Label   string `json:"label"`
	Balance struct {
		Updated       time.Time `json:"updated"`
		NumTx         int       `json:"numTx"`
		NumUnspents   int       `json:"numUnspents"`
		TotalReceived int       `json:"totalReceived"`
		TotalSent     int       `json:"totalSent"`
	} `json:"balance"`
}

type TransactionDescription struct {
	Txid   string `json:"txid"`
	Status string `json:"status"`
}

// List Wallet Transfers

func (b *BitGo) ListWalletTransfers(walletId string, params ListParams) (list TransferList, err error) {
	err = b.get(
		fmt.Sprintf("%s/wallet/%s/transfer",
			b.coin,
			walletId),
		params,
		&list)
	return
}

// Get Wallet Transfer

func (b *BitGo) GetWalletTransfer(walletId string, transferId string) (transfer Transfer, err error) {
	err = b.get(
		fmt.Sprintf("%s/wallet/%s/transfer/%s",
			b.coin,
			walletId,
			transferId),
		nil,
		&transfer)
	return
}

// Get Wallet Transfer By Sequence ID

func (b *BitGo) GetWalletTransferBySequenceID(walletId string, sequenceId string) (transfer Transfer, err error) {
	err = b.get(
		fmt.Sprintf("%s/wallet/%s/transfer/sequenceId/%s",
			b.coin,
			walletId,
			sequenceId),
		nil,
		&transfer)
	return
}

// Create Wallet Address
type ForwarderVersionType int

const (
	FORWARDER_VERSION_0 ForwarderVersionType = 0
	FORWARDER_VERSION_1                      = 1
)

type AddressParams struct {
	Chain            int                  `json:"chain,omitempty"`
	Label            string               `json:"label,omitempty"`
	ForwarderVersion ForwarderVersionType `json:"forwarderVersion,omitempty"`
}

func (b *BitGo) CreateWalletAddress(walletId string, params *AddressParams) (address Address, err error) {
	err = b.post(
		fmt.Sprintf("%s/wallet/%s/address",
			b.coin,
			walletId),
		params,
		&address)
	return
}

// Get Wallet Address

func (b *BitGo) GetWalletAddress(walletId string, addressOrId string) (address Address, err error) {
	err = b.get(
		fmt.Sprintf("%s/wallet/%s/address/%s",
			b.coin,
			walletId,
			addressOrId),
		nil,
		&address)
	return
}

// Update Wallet Address

type UpdateWalletAddressParams struct {
	Label string `json:"label,omitempty"`
}

func (b *BitGo) UpdateWalletAddress(walletId string, addressOrId string, params UpdateWalletAddressParams) (address Address, err error) {
	err = b.put(
		fmt.Sprintf("%s/wallet/%s/address/%s",
			b.coin,
			walletId,
			addressOrId),
		params,
		&address)
	return
}

// Send Transaction

type SendParams struct {
	Address string `json:"address" valid:"required"`
	Amount  string `json:"amount" valid:"required"`

	WalletPassphrase string `json:"walletPassphrase" valid:"required"`
	Prv              string `json:"prv,omitempty"`
	NumBlocks        int    `json:"numBlocks,omitempty"`
	FeeRate          int    `json:"feeRate,omitempty"`
	Comment          string `json:"comment,omitempty"`
	//	Unspents                    array  `json:"unspents,omitempty"`
	MinConfirms                 int    `json:"minConfirms,omitempty"`
	EnforceMinConfirmsForChange bool   `json:"enforceMinConfirmsForChange,omitempty"`
	TargetWalletUnspents        int    `json:"targetWalletUnspents,omitempty"`
	NoSplitChange               bool   `json:"noSplitChange,omitempty"`
	MinValue                    int    `json:"minValue,omitempty"`
	MaxValue                    int    `json:"maxValue,omitempty"`
	GasPrice                    int    `json:"gasPrice,omitempty"`
	GasLimit                    int    `json:"gasLimit,omitempty"`
	SequenceId                  int    `json:"sequenceId,omitempty"`
	Segwit                      bool   `json:"segwit,omitempty"`
	LastLedgerSequence          int    `json:"lastLedgerSequence,omitempty"`
	LedgerSequenceDelta         string `json:"ledgerSequenceDelta,omitempty"`
}

func (b *BitGo) SendTransaction(walletId string, params SendParams) (transactionDescription TransactionDescription, err error) {
	err = b.post(
		fmt.Sprintf("%s/wallet/%s/sendcoins",
			b.coin,
			walletId),
		params,
		&transactionDescription)
	return
}

// Send Transaction to Many

type Recipient struct {
	Address  string `json:"address" valid:"required"`
	Amount   int    `json:"amount" valid:"required"`
	GasPrice int    `json:"gasPrice,omitempty"`
}

type SendToManyParams struct {
	Recipients []Recipient `json:"recipients" valid:"required"`

	WalletPassphrase string `json:"walletPassphrase" valid:"required"`
	Prv              string `json:"prv,omitempty"`
	NumBlocks        int    `json:"numBlocks,omitempty"`
	FeeRate          int    `json:"feeRate,omitempty"`
	Comment          string `json:"comment,omitempty"`
	//	Unspents                    array  `json:"unspents,omitempty"`
	MinConfirms                 int    `json:"minConfirms,omitempty"`
	EnforceMinConfirmsForChange bool   `json:"enforceMinConfirmsForChange,omitempty"`
	TargetWalletUnspents        int    `json:"targetWalletUnspents,omitempty"`
	NoSplitChange               bool   `json:"noSplitChange,omitempty"`
	MinValue                    int    `json:"minValue,omitempty"`
	MaxValue                    int    `json:"maxValue,omitempty"`
	GasPrice                    int    `json:"gasPrice,omitempty"`
	GasLimit                    int    `json:"gasLimit,omitempty"`
	SequenceId                  int    `json:"sequenceId,omitempty"`
	Segwit                      bool   `json:"segwit,omitempty"`
	LastLedgerSequence          int    `json:"lastLedgerSequence,omitempty"`
	LedgerSequenceDelta         string `json:"ledgerSequenceDelta,omitempty"`
}

func (b *BitGo) SendTransactionToMany(walletId string, params SendToManyParams) (transactionDescription TransactionDescription, err error) {
	err = b.post(
		fmt.Sprintf("%s/wallet/%s/sendmany",
			b.coin,
			walletId),
		params,
		&transactionDescription)
	return
}
