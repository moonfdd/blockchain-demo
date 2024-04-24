package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/tjfoc/gmsm/sm2"
)

func main() {

	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}
	// d, _ := json.MarshalIndent(privateKey, "", "  ")
	// fmt.Println(string(d))
	// return
	privateKey.D = big.NewInt(1)
	privateKey.D.SetString("58813022008286889280726917052263921292796008422545733079686076267308335241271", 10)
	privateKey.X = big.NewInt(1)
	privateKey.X.SetString("74426446515384143246782476669202054294874984675849763417103443808257273376614", 10)
	privateKey.Y = big.NewInt(1)
	privateKey.Y.SetString("106493256329545608248660242228701043221716914976457764892567176785153111996293", 10)
	// d, _ := json.MarshalIndent(privateKey, "", "  ")
	// fmt.Println(string(d))
	// return
	publicKey := &privateKey.PublicKey
	// fmt.Println(publicKey)

	if true {
		// 要加密的数据
		plaintext := []byte("Hello, SM2!")
		// 加密
		ciphertext, err := publicKey.EncryptAsn1(plaintext, nil)
		if err != nil {
			fmt.Println("Error encrypting:", err)
			return
		}

		fmt.Println("Encrypted data:", hex.EncodeToString(ciphertext))

		// 解密
		decrypted, err := privateKey.DecryptAsn1(ciphertext)
		if err != nil {
			fmt.Println("Error decrypting:", err)
			return
		}

		fmt.Println("Decrypted data:", string(decrypted))
	}
	if true {
		// 要加密的数据
		plaintext := []byte("Hello, SM2!")
		// 加密
		ciphertext, err := publicKey.EncryptAsn1(plaintext, nil)
		if err != nil {
			fmt.Println("Error encrypting:", err)
			return
		}

		fmt.Println("Encrypted data:", hex.EncodeToString(ciphertext))

		// 解密
		decrypted, err := privateKey.DecryptAsn1(ciphertext)
		if err != nil {
			fmt.Println("Error decrypting:", err)
			return
		}

		fmt.Println("Decrypted data:", string(decrypted))
	}

}
