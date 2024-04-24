package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	if false {
		account, err := keystore.StoreKey("./", "123456", keystore.StandardScryptN, keystore.StandardScryptP)
		// account, err := keystore.StoreKey("./", "123456", keystore.LightScryptN, keystore.LightScryptP)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(account)
	}
	if true {
		ks := keystore.NewKeyStore("./", keystore.StandardScryptN, keystore.StandardScryptP)
		fmt.Println(ks.Accounts())
	}
}
