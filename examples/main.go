package main

import (
	"log"

	"github.com/jazzserve/bitgo"
)

func main() {
	b := &bitgo.BitGo{
		Host:  "https://test.bitgo.com/api/v2",
		Token: "{Access token}",
	}

	list, err := b.Coin("tbtc").ListWallets(nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, w := range list.Wallets {
		log.Println(w.ID, w.Label)
	}
}
