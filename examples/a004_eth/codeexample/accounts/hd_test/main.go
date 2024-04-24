package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
)

func main() {
	if false {
		fmt.Println(accounts.DefaultRootDerivationPath)
		fmt.Println(accounts.DefaultBaseDerivationPath)
		fmt.Println(accounts.LegacyLedgerBaseDerivationPath)
	}
	if true {
		fmt.Println(accounts.ParseDerivationPath("path"))
	}
}
