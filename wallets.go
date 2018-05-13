package bitgo

import (
	"fmt"
	"time"
)

type ListWallets struct {
	NextBatchPrevId string `json:"nextBatchPrevId"`
	Wallets         []struct {
		ID    string `json:"id"`
		Users []struct {
			User        string   `json:"user"`
			Permissions []string `json:"permissions"`
		} `json:"users"`
		Coin                            string   `json:"coin"`
		Label                           string   `json:"label"`
		M                               int      `json:"m"`
		N                               int      `json:"n"`
		Keys                            []string `json:"keys"`
		Tags                            []string `json:"tags"`
		DisableTransactionNotifications bool     `json:"disableTransactionNotifications"`
		Freeze                          struct {
		} `json:"freeze"`
		Deleted           bool `json:"deleted"`
		ApprovalsRequired int  `json:"approvalsRequired"`
		CoinSpecific      struct {
		} `json:"coinSpecific"`
	} `json:"wallets"`
}

type Wallet struct {
	ID    string `json:"id"`
	Users []struct {
		User        string   `json:"user"`
		Permissions []string `json:"permissions"`
	} `json:"users"`
	Coin                            string   `json:"coin"`
	Label                           string   `json:"label"`
	M                               int      `json:"m"`
	N                               int      `json:"n"`
	Keys                            []string `json:"keys"`
	Tags                            []string `json:"tags"`
	DisableTransactionNotifications bool     `json:"disableTransactionNotifications"`
	Freeze                          struct {
	} `json:"freeze"`
	Deleted           bool `json:"deleted"`
	ApprovalsRequired int  `json:"approvalsRequired"`
	CoinSpecific      struct {
	} `json:"coinSpecific"`
	Balance                int    `json:"balance"`
	ConfirmedBalance       int    `json:"confirmedBalance"`
	SpendableBalance       int    `json:"spendableBalance"`
	BalanceString          string `json:"balanceString"`
	ConfirmedBalanceString string `json:"confirmedBalanceString"`
	SpendableBalanceString string `json:"spendableBalanceString"`
	ReceiveAddress         struct {
		Address      string `json:"address"`
		Chain        int    `json:"chain"`
		Index        int    `json:"index"`
		Coin         string `json:"coin"`
		Wallet       string `json:"wallet"`
		CoinSpecific struct {
			RedeemScript string `json:"redeemScript"`
		} `json:"coinSpecific"`
	} `json:"receiveAddress"`
	Admin struct {
		Policy struct {
			ID      string    `json:"id"`
			Label   string    `json:"label"`
			Version int       `json:"version"`
			Date    time.Time `json:"date"`
			Rules   []struct {
				ID     string `json:"id"`
				Coin   string `json:"coin"`
				Type   string `json:"type"`
				Action struct {
					Type string `json:"type"`
				} `json:"action"`
				Condition struct {
					AmountString string        `json:"amountString"`
					TimeWindow   int           `json:"timeWindow"`
					GroupTags    []string      `json:"groupTags"`
					ExcludeTags  []interface{} `json:"excludeTags"`
				} `json:"condition"`
			} `json:"rules"`
		} `json:"policy"`
	} `json:"admin"`
}

type GeneratedWallet struct {
	Wallet struct {
		Wallet struct {
			ID    string `json:"id"`
			Users []struct {
				User        string   `json:"user"`
				Permissions []string `json:"permissions"`
			} `json:"users"`
			Coin                            string   `json:"coin"`
			Label                           string   `json:"label"`
			M                               int      `json:"m"`
			N                               int      `json:"n"`
			Keys                            []string `json:"keys"`
			Tags                            []string `json:"tags"`
			DisableTransactionNotifications bool     `json:"disableTransactionNotifications"`
			Freeze                          struct {
			} `json:"freeze"`
			Deleted           bool `json:"deleted"`
			ApprovalsRequired int  `json:"approvalsRequired"`
			IsCold            bool `json:"isCold"`
			CoinSpecific      struct {
			} `json:"coinSpecific"`
			Balance                int    `json:"balance"`
			ConfirmedBalance       int    `json:"confirmedBalance"`
			SpendableBalance       int    `json:"spendableBalance"`
			BalanceString          string `json:"balanceString"`
			ConfirmedBalanceString string `json:"confirmedBalanceString"`
			SpendableBalanceString string `json:"spendableBalanceString"`
			ReceiveAddress         struct {
				Address      string `json:"address"`
				Chain        int    `json:"chain"`
				Index        int    `json:"index"`
				Coin         string `json:"coin"`
				Wallet       string `json:"wallet"`
				CoinSpecific struct {
					RedeemScript string `json:"redeemScript"`
				} `json:"coinSpecific"`
			} `json:"receiveAddress"`
			PendingApprovals []interface{} `json:"pendingApprovals"`
		} `json:"_wallet"`
	} `json:"wallet"`
}

// List Wallets

func (b *BitGo) ListWallets(params *ListParams) (list ListWallets, err error) {
	err = b.get(
		fmt.Sprintf("%s/wallet",
			b.coin),
		params,
		&list)
	return
}

// Generate Wallet

type GenerateWalletParams struct {
	Label                           string `json:"label"`
	Passphrase                      string `json:"passphrase"`
	UserKey                         string `json:"userKey,omitempty"`
	BackupXpub                      string `json:"backupXpub,omitempty"`
	BackupXpubProvider              string `json:"backupXpubProvider,omitempty"`
	Enterprise                      string `json:"enterprise,omitempty"`
	DisableTransactionNotifications bool   `json:"disableTransactionNotifications,omitempty"`
	GasPrice                        int    `json:"gasPrice,omitempty"`
	PasscodeEncryptionCode          string `json:"passcodeEncryptionCode,omitempty"`
}

func (b *BitGo) GenerateWallet(params *GenerateWalletParams) (wallet GeneratedWallet, err error) {
	err = b.post(
		fmt.Sprintf("%s/wallet/generate",
			b.coin),
		params,
		&wallet)
	return
}

// Get Wallet

type GetWalletParams struct {
	AllTokens bool `url:"allTokens,omitempty"`
}

func (b *BitGo) GetWallet(walletId string, params *GetWalletParams) (wallet Wallet, err error) {
	err = b.get(
		fmt.Sprintf("%s/wallet/%s",
			b.coin,
			walletId),
		params,
		&wallet)
	return
}

// Update Wallet

type UpdateWalletParams struct {
	Label                           string  `json:"label,omitempty"`
	DisableTransactionNotifications bool    `json:"disableTransactionNotifications,omitempty"`
	ApprovalsRequired               float64 `json:"approvalsRequired,omitempty"`
}

func (b *BitGo) UpdateWallet(walletId string, params UpdateWalletParams) (wallet Wallet, err error) {
	err = b.put(
		fmt.Sprintf("%s/wallet/%s",
			b.coin,
			walletId),
		params,
		&wallet)
	return
}

// Get Wallet By Address

func (b *BitGo) GetWalletByAddress(address string) (wallet Wallet, err error) {
	err = b.get(
		fmt.Sprintf("%s/wallet/address/%s",
			b.coin,
			address),
		nil,
		&wallet)
	return
}
