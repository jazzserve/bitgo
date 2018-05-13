package main

import (
	"log"

	"github.com/jazzserve/bitgo"
)

func main() {
	b := bitgo.New("test", "{Access token}")

	list, err := b.Coin("tbtc").ListWallets(nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, w := range list.Wallets {
		log.Println(w.ID, w.Label)
	}
}
