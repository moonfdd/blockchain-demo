package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

// 密钥派生Key Derivation
func main() {
	password := []byte("supersecretpassword")
	salt := make([]byte, 16) // 生成一个随机的 salt
	_, err := rand.Read(salt)
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	key := pbkdf2.Key(password, salt, 10000, 32, sha256.New)
	fmt.Printf("Derived key: %x\n", key)

	key = pbkdf2.Key(password, salt, 10000, 32, sha256.New)
	fmt.Printf("Derived key: %x\n", key)
}
