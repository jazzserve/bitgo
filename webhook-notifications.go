package bitgo

import (
	"fmt"
	"time"
)

type WebhookEvent struct {
	Hash     string `json:"hash"`
	Transfer string `json:"transfer"`
	Coin     string `json:"coin"`
	Type     string `json:"type"`
	Wallet   string `json:"wallet"`
}

type Webhook struct {
	ID       string    `json:"id"`
	Label    string    `json:"label"`
	Created  time.Time `json:"created"`
	WalletID string    `json:"walletId"`
	Coin     string    `json:"coin"`
	Type     string    `json:"type"`
	URL      string    `json:"url"`
	Version  int       `json:"version"`
}

type Webhooks []Webhook

type ListWebhooks struct {
	Webhooks Webhooks `json:"webhooks"`
}

// List Wallet Webhooks

func (b *BitGo) ListWalletWebhooks(walletId string, params GetWalletParams) (webhooks Webhooks, err error) {
	responce := ListWebhooks{}
	err = b.get(
		fmt.Sprintf("%s/wallet/%s/webhooks",
			b.coin,
			walletId),
		nil,
		&responce)
	return responce.Webhooks, err
}

// Add Wallet Webhook

// Remove Wallet Webhook

// Simulate Wallet Webhook

// List User Webhooks

//  Add User Webhook

// Remove User Webhook

// Simulate User Webhook
