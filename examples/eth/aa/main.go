package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {

	if false {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			fmt.Println("Error generating private key:", err)
			return
		}

		publicKey := privateKey.PublicKey
		fmt.Printf("Private key: %x\n", privateKey.D)
		fmt.Printf("Public key: %x\n", crypto.FromECDSAPub(&publicKey))
		return
	}
	if true {
		privateKeyHex := "0000000000000000000000000000000000000000000000000000000000000001"
		privateKeyBytes, err := hex.DecodeString(privateKeyHex)
		if err != nil {
			log.Fatal(err)
		}

		privateKey, err := crypto.ToECDSA(privateKeyBytes)
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public().(*ecdsa.PublicKey)
		publicKeyBytes := crypto.FromECDSAPub(publicKey)

		fmt.Printf("Private key: %s\n", privateKeyHex)
		fmt.Printf("Public key: %x\n", publicKeyBytes)
	}
}
