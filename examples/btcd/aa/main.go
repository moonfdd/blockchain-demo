package main

import (
	// "crypto/rand"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	// "log"
	// "github.com/btcsuite/btcd/btcec"
)

// https://github.com/btcsuite/btcd/blob/v0.24.0/btcec/privkey.go
func main() {
	if true {
		// 生成随机私钥
		privateKey, err := btcec.NewPrivateKey()
		if err != nil {
			fmt.Println("Error generating private key:", err)
			return
		}

		// 获取对应的公钥
		publicKey := privateKey.PublicKey
		// publicKeyBytes := elliptic.Marshal(btcec.S256(), publicKey.X, publicKey.Y)
		// publicKey := privKey.PubKey()
		fmt.Printf("Private key: %x\n", privateKey.D)
		fmt.Printf("Public key: %x\n", publicKey)
		return
	}
	if false {
		// 生成随机私钥
		privateKey, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
		if err != nil {
			fmt.Println("Error generating private key:", err)
			return
		}

		// 获取对应的公钥
		publicKey := privateKey.PublicKey
		// publicKeyBytes := elliptic.Marshal(btcec.S256(), publicKey.X, publicKey.Y)
		// publicKey := privKey.PubKey()
		fmt.Printf("Private key: %x\n", privateKey.D)
		fmt.Printf("Public key: %x\n", publicKey)
		return
	}
	if true {
		hexString := "01" // 16进制字符串
		hexString = "36747315f61b4dfc9b5ce947301312a604248e992b51778d28db6a1f94f832e3"
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
}
