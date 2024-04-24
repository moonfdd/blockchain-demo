package main

import (
	"crypto/rand"
	"encoding/asn1"
	"fmt"

	"github.com/tjfoc/gmsm/sm2"
)

// 在 Golang 中，要实现 SM2 密钥分发，可以使用第三方库提供的 SM2 算法实现。一个常用的库是 github.com/tjfoc/gmsm/sm2，它提供了对 SM2 算法的支持。
func main() {
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("生成私钥失败：", err)
		return
	}

	publicKey := &privateKey.PublicKey

	message := []byte("Hello, SM2!")

	ciphertext, err := publicKey.EncryptAsn1(message, nil)
	if err != nil {
		fmt.Println("加密失败：", err)
		return
	}

	fmt.Printf("加密后的密文：%x\n", ciphertext)

	plaintext, err := privateKey.DecryptAsn1(ciphertext)
	if err != nil {
		fmt.Println("解密失败：", err)
		return
	}

	fmt.Printf("解密后的明文：%s\n", string(plaintext))

	signature, err := privateKey.Sign(rand.Reader, message, nil)
	if err != nil {
		fmt.Println("签名失败：", err)
		return
	}

	asn1Data, _ := asn1.Marshal(signature)
	fmt.Printf("数字签名：%x\n", asn1Data)

	valid := publicKey.Verify(message, signature)
	if valid {
		fmt.Println("数字签名验证通过")
	} else {
		fmt.Println("数字签名验证失败")
	}
}
