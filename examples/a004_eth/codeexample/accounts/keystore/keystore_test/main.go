package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	if true {
		ks := keystore.NewKeyStore("./", keystore.StandardScryptN, keystore.StandardScryptP)

		fmt.Println(ks.Accounts())
	}
}
