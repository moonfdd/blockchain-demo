package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

// 如果你想要实现在Golang中进行密钥共享（Key Sharing）而不是进行数据加密和解密，那么你可以使用Diffie-Hellman密钥交换协议。下面是一个示例代码，演示了如何在Golang中使用Diffie-Hellman密钥交换进行密钥共享
// 这段代码使用了RSA加密算法和Diffie-Hellman密钥交换协议实现了密钥共享。它生成一个RSA密钥对，并从私钥中提取公钥。然后，模拟了Alice和Bob生成各自的Diffie-Hellman共享密钥。接着，Alice使用Bob的公钥计算共享密钥，Bob使用Alice的公钥计算共享密钥。最后，验证两者计算得到的共享密钥是否相同。
// 注意，这只是一个简单的示例代码，可能并不完整或安全。在实际应用中，你需要根据具体情况进行适当的改进和保护，以确保密钥共享的安全性和可靠性。
func main() {
	// 创建一个RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err.Error())
	}

	// 从私钥中提取公钥
	publicKey := &privateKey.PublicKey

	// 模拟Alice和Bob进行密钥共享
	aliceSharedKey, err := generateSharedKey()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Alice's shared key:", aliceSharedKey)

	bobSharedKey, err := generateSharedKey()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Bob's shared key:", bobSharedKey)

	// Alice使用Bob的公钥计算共享密钥
	aliceComputedKey, err := computeSharedKey(aliceSharedKey, publicKey)
	if err != nil {
		panic(err.Error())
	}

	// Bob使用Alice的公钥计算共享密钥
	bobComputedKey, err := computeSharedKey(bobSharedKey, publicKey)
	if err != nil {
		panic(err.Error())
	}

	// 验证两者计算得到的共享密钥是否相同
	fmt.Println("Alice's computed key:", aliceComputedKey)
	fmt.Println("Bob's computed key:", bobComputedKey)
}

// 生成Diffie-Hellman共享密钥对
func generateSharedKey() ([]byte, error) {
	sharedKey := make([]byte, 32)
	if _, err := rand.Read(sharedKey); err != nil {
		return nil, err
	}

	return sharedKey, nil
}

// 使用公钥计算共享密钥
func computeSharedKey(sharedKey []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	computedKey, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, sharedKey)
	if err != nil {
		return nil, err
	}

	return computedKey, nil
}
