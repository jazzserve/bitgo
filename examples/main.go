package main

import (
	"log"

	"github.com/jazzserve/bitgo"
)

func main() {
	b, err := bitgo.New("test", "{Access token}")
	if err != nil {
		log.Fatal(err.Error())
	}

	list, err := b.Coin("tbtc").ListWallets(nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, w := range list.Wallets {
		log.Println(w.ID, w.Label)
	}
}
