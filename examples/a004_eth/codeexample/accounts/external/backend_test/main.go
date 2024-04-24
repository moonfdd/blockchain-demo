package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/external"
)

const key = `{"address":"7e5f4552091a69125d5dfcb7b8c2659029395bdf","crypto":{"cipher":"aes-128-ctr","ciphertext":"e468ad3ce2a6ae1436f8b50259f7cef84e688df2b08575f93eda7067492fec45","cipherparams":{"iv":"80697348d506166676cc27cc10aa00e0"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"806d6da06fbb4f2480e4c93f7bd82d5575f2da3ce382b469e6ed293e9e5cb52c"},"mac":"5fb6e482ee88c6a99972ace20827d547d2245249c2c3cfa8a8c8fea8bc82b771"},"id":"509c8fb9-b781-4c3b-bce7-e11b75e03ac3","version":3}`

func main() {
	// 创建一个具有clef后端的交易签名者
	// clef --keystore datadirprivate/keystore --chainid 666123 --http
	if true {
		clef, err := external.NewExternalSigner("http://127.0.0.1:8550") //需要在clef客户端上确认y
		if err != nil {
			fmt.Println("NewExternalSigner失败", err)
			return
		}
		defer clef.Close()
		transactOpts := bind.NewClefTransactor(clef, clef.Accounts()[0])
		fmt.Println("NewClefTransactor成功", transactOpts)
		return
	}
	fmt.Println("Hello World")
}