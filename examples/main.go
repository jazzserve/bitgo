package main

import (
	"log"
	"time"

	"github.com/jazzserve/bitgo"
)

func main() {
	b, err := bitgo.New("test", "{Access token}", time.Minute)
	if err != nil {
		log.Fatal(err.Error())
	}

	list, err := b.Coin("tbtc").Debug(true).ListWallets(bitgo.ListParams{
		AllTokens: true,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, w := range list.Wallets {
		log.Println(w.ID, w.Label)
	}
}
