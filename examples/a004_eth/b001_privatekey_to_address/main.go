package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	if false {
		privateKeyHex := "0x000000000000000000000000000000000000000000000000000000000000000a"
		privateKeyBytes := common.FromHex(privateKeyHex)
		privateKey, err := crypto.ToECDSA(privateKeyBytes)
		if err != nil {
			log.Fatal(err)
		}
		publicKey := privateKey.PublicKey
		publicKeyBytes := crypto.FromECDSAPub(&publicKey)
		fmt.Printf("Private key: %s\n", privateKeyHex)
		fmt.Printf("Public key: %x\n", publicKeyBytes)
		fmt.Println("Address:", crypto.PubkeyToAddress(publicKey))
		fmt.Println("----")
		return
	}
	if false {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			fmt.Println("Error generating private key:", err)
			return
		}

		publicKey := privateKey.PublicKey
		fmt.Printf("Private key: %x\n", privateKey.D)
		fmt.Printf("Public key: %x\n", crypto.FromECDSAPub(&publicKey))
		fmt.Println("Address:", crypto.PubkeyToAddress(publicKey))
		fmt.Println("----")
	}

	if true {
		if true {
			privateKeyHex := "0x0000000000000000000000000000000000000000000000000000000000000001"
			privateKeyBytes := common.FromHex(privateKeyHex)
			privateKey, err := crypto.ToECDSA(privateKeyBytes)
			if err != nil {
				log.Fatal(err)
			}
			publicKey := privateKey.PublicKey
			publicKeyBytes := crypto.FromECDSAPub(&publicKey)
			fmt.Printf("Private key: %s\n", privateKeyHex)
			fmt.Printf("Public key: %x\n", publicKeyBytes)
			fmt.Println("Address:", crypto.PubkeyToAddress(publicKey))
			fmt.Println("----") //0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF
			//"0xac8dd6b5c1b0a5dcb1ad173d4042708ee2b537c3dd05c7343c0072cfc7cb1929"
			//"0x40b1e4b7baf054e0fbce31533079af1a9f7138ee8db4a3626498288323b90c35"
			//0x6813Eb9362372EEF6200f3b1dbC3f819671cBA69
		}
		if true {
			privateKeyHex := "0x0000000000000000000000000000000000000000000000000000000000000002"
			privateKeyBytes := common.FromHex(privateKeyHex)
			privateKey, err := crypto.ToECDSA(privateKeyBytes)
			if err != nil {
				log.Fatal(err)
			}
			publicKey := privateKey.PublicKey
			publicKeyBytes := crypto.FromECDSAPub(&publicKey)
			fmt.Printf("Private key: %s\n", privateKeyHex)
			fmt.Printf("Public key: %x\n", publicKeyBytes)
			fmt.Println("Address:", crypto.PubkeyToAddress(publicKey))
			fmt.Println("----")
		}
		if true {
			privateKeyHex := "0x0000000000000000000000000000000000000000000000000000000000000003"
			privateKeyBytes := common.FromHex(privateKeyHex)
			privateKey, err := crypto.ToECDSA(privateKeyBytes)
			if err != nil {
				log.Fatal(err)
			}
			publicKey := privateKey.PublicKey
			publicKeyBytes := crypto.FromECDSAPub(&publicKey)
			fmt.Printf("Private key: %s\n", privateKeyHex)
			fmt.Printf("Public key: %x\n", publicKeyBytes)
			fmt.Println("Address:", crypto.PubkeyToAddress(publicKey))
			fmt.Println("----")
		}
		if true {
			privateKeyHex := "0x0000000000000000000000000000000000000000000000000000000000000004"
			privateKeyBytes := common.FromHex(privateKeyHex)
			privateKey, err := crypto.ToECDSA(privateKeyBytes)
			if err != nil {
				log.Fatal(err)
			}
			publicKey := privateKey.PublicKey
			publicKeyBytes := crypto.FromECDSAPub(&publicKey)
			fmt.Printf("Private key: %s\n", privateKeyHex)
			fmt.Printf("Public key: %x\n", publicKeyBytes)
			fmt.Println("Address:", crypto.PubkeyToAddress(publicKey))
			fmt.Println("----")
		}
		if true {
			privateKeyHex := "0x0000000000000000000000000000000000000000000000000000000000000005"
			privateKeyBytes := common.FromHex(privateKeyHex)
			privateKey, err := crypto.ToECDSA(privateKeyBytes)
			if err != nil {
				log.Fatal(err)
			}
			publicKey := privateKey.PublicKey
			publicKeyBytes := crypto.FromECDSAPub(&publicKey)
			fmt.Printf("Private key: %s\n", privateKeyHex)
			fmt.Printf("Public key: %x\n", publicKeyBytes)
			fmt.Println("Address:", crypto.PubkeyToAddress(publicKey))
			fmt.Println("----")
		}
	}
}
