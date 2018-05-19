package bitgo

import (
	"testing"
)

// List Wallet Webhooks

func TestListWalletWebhooks(t *testing.T) {
	coin, params := getTestCoin(t)

	_, err := coin.ListWalletWebhooks(params.WalletId, GetWalletParams{AllTokens: true})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Add Wallet Webhook

// Remove Wallet Webhook

// Simulate Wallet Webhook

// List User Webhooks

//  Add User Webhook

// Remove User Webhook

// Simulate User Webhook
