package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	if false {
		fmt.Println(keystore.StandardScryptN)
		fmt.Println(keystore.StandardScryptP)
		fmt.Println(keystore.LightScryptN)
		fmt.Println(keystore.LightScryptP)
	}
	// 创建keystore文件
	if true {
		account, err := keystore.StoreKey("./", "123456", keystore.StandardScryptN, keystore.StandardScryptP)
		// account, err := keystore.StoreKey("./", "123456", keystore.LightScryptN, keystore.LightScryptP)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(account)
	}
}
