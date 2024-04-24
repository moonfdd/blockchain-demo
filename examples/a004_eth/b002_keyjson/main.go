package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EncryptedKeyJSONV3 struct {
	Address string              `json:"address"`
	Crypto  keystore.CryptoJSON `json:"crypto"`
	Id      string              `json:"id"`
	Version int                 `json:"version"`
}

func main() {
	if true {
		keyjson, err := os.ReadFile("d.json")
		if err != nil {
			log.Fatal(err)
		}
		var data = new(EncryptedKeyJSONV3)
		err = json.Unmarshal(keyjson, data)
		if err != nil {
			fmt.Println("反序列化失败")
			return
		}
		dataJson, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println("序列化失败")
			return
		}
		fmt.Println(string(dataJson))
		password := "123456"
		key, err := keystore.DecryptKey(keyjson, password)
		if err != nil {
			log.Fatal(err)
		}
		privateKey := key.PrivateKey
		publicKey := privateKey.PublicKey
		publicKeyBytes := crypto.FromECDSAPub(&publicKey)
		fmt.Printf("Private key: %s\n", privateKey.D.Text(16))
		fmt.Printf("Public key: %x\n", publicKeyBytes)
		fmt.Println("Address:", crypto.PubkeyToAddress(publicKey))
		fmt.Println("----")
	}
	if false {
		keyjson, err := os.ReadFile("a.json")
		if err != nil {
			log.Fatal(err)
		}
		password := ""
		address := common.HexToAddress("45dea0fb0bba44f4fcf290bba71fd57d7117cbb8")

		// Do a few rounds of decryption and encryption
		for i := 0; i < 3; i++ {
			// Try a bad password first
			if _, err := keystore.DecryptKey(keyjson, password+"bad"); err == nil {
				fmt.Errorf("test %d: json key decrypted with bad password", i)
			}
			// Decrypt with the correct password
			key, err := keystore.DecryptKey(keyjson, password)
			if err != nil {
				log.Fatalf("test %d: json key failed to decrypt: %v", i, err)
			}
			if key.Address != address {
				fmt.Errorf("test %d: key address mismatch: have %x, want %x", i, key.Address, address)
			}
			// Recrypt with a new password and start over
			password += "new data appended" // nolint: gosec
			// if keyjson, err = keystore.EncryptKey(key, password, veryLightScryptN, veryLightScryptP); err != nil {
			// 	fmt.Errorf("test %d: failed to re-encrypt key %v", i, err)
			// }
			fmt.Println("string(keyjson) = ", string(keyjson))
			fmt.Println("password = ", password)
			fmt.Println("key = ", key)
			fmt.Println(key.PrivateKey)
			fmt.Println("----")
		}
	}
}
