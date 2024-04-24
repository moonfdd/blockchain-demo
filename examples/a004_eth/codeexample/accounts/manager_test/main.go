package main

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	if true {
		//创建账户管理器
		ks := keystore.NewKeyStore("test/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
		am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
		data, _ := json.MarshalIndent(am, "", "  ")
		fmt.Println(string(data))
	}
}
