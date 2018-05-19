package main

import (
	"log"
	"time"

	"github.com/jazzserve/bitgo"
)

func main() {
	b, err := bitgo.New("test", time.Minute)
	if err != nil {
		log.Fatal(err.Error())
	}

	list, err := b.Token("{Access token}").Coin("tbtc").Debug(true).ListWallets(bitgo.ListParams{
		AllTokens: true,
	})
	if err != nil {
		log.Fatalf("%#v\n", err.(bitgo.Error))
	}

	for _, w := range list.Wallets {
		log.Println(w.ID, w.Label)
	}
}
