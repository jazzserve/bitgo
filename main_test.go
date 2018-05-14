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
	b        *BitGo
	tbtc     *BitGo
	walletId string
	address  string
)

func init() {
	rand.Seed(time.Now().UnixNano())

	env := os.Getenv("ENV")
	token := os.Getenv("TOKEN")
	walletId = os.Getenv("WALLET")
	address = os.Getenv("ADDRESS")
	coin := os.Getenv("COIN")

	b = New(env, token)
	tbtc = b.Coin(coin)
}
