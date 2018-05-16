package bitgo

import (
	"math/rand"
	"os"
	"testing"
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

type TestParams struct {
	Env        string
	Token      string
	Coin       string
	WalletId   string
	Address    string
	TransferId string
	SequenceId string
	Passphrase string
	Email      string
}

func getTestParams() (params *TestParams) {
	return &TestParams{
		Env:        os.Getenv("ENV"),
		Token:      os.Getenv("ACCESS_TOKEN"),
		Coin:       os.Getenv("COIN"),
		WalletId:   os.Getenv("WALLET_ID"),
		Address:    os.Getenv("ADDRESS"),
		TransferId: os.Getenv("TRANSFER_ID"),
		SequenceId: os.Getenv("SEQUENCE_ID"),
		Email:      os.Getenv("EMAIL"),
		Passphrase: os.Getenv("PASSPHRASE"),
	}
}

func getTestBitGo(t *testing.T) (b *BitGo, params *TestParams) {
	params = getTestParams()

	b, err := New(params.Env, params.Token)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func getTestCoin(t *testing.T) (coin *BitGo, params *TestParams) {
	b, params := getTestBitGo(t)
	coin = b.Coin(params.Coin)
	return
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
