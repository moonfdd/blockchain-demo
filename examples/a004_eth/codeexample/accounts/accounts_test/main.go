package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
)

func main() {
	if false {
		fmt.Println(accounts.MimetypeDataWithValidator)
		fmt.Println(accounts.MimetypeTypedData)
		fmt.Println(accounts.MimetypeClique)
		fmt.Println(accounts.MimetypeTextPlain)

		fmt.Println(accounts.WalletArrived)
		fmt.Println(accounts.WalletOpened)
		fmt.Println(accounts.WalletDropped)
	}
	if true {
		hash := accounts.TextHash([]byte("Hello Joe"))
		fmt.Printf("hash: %x", hash)
	}
}
