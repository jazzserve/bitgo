package bitgo

import (
	"math/rand"
	"os"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var (
	b          *BitGo
	coin       *BitGo
	walletId   string
	address    string
	transferId string
	sequenceId string
	passphrase string
)

func init() {
	rand.Seed(time.Now().UnixNano())

	env := os.Getenv("ENV")
	token := os.Getenv("ACCESS_TOKEN")
	coinCode := os.Getenv("COIN")

	walletId = os.Getenv("WALLET_ID")
	address = os.Getenv("ADDRESS")
	transferId = os.Getenv("TRANSFER_ID")
	sequenceId = os.Getenv("SEQUENCE_ID")
	passphrase = os.Getenv("PASSPHRASE")

	b = New(env, token)
	coin = b.Coin(coinCode)
}
