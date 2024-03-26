package main

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
)

// https://github.com/btcsuite/btcd/blob/v0.24.0/btcec/privkey.go
func main() {
	hexString := "ff" // 16进制字符串
	privateKeyBytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("解码出错：", err)
		return
	}
	_, publicKey := btcec.PrivKeyFromBytes(privateKeyBytes)
	// publicKey := privKey.PubKey()
	fmt.Println(publicKey.X().Text(16), publicKey.Y().Text(16))
	return
}
