package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func encrypt(plainText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCBCEncrypter(block, iv)

	plainData := []byte(plainText)
	plainData = pkcs7Padding(plainData, block.BlockSize())
	cipherText := make([]byte, len(plainData))

	stream.CryptBlocks(cipherText, plainData)

	return hex.EncodeToString(cipherText), nil
}

func decrypt(cipherText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherData, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCBCDecrypter(block, iv)

	plainData := make([]byte, len(cipherData))
	stream.CryptBlocks(plainData, cipherData)

	plainData = pkcs7Unpadding(plainData)

	return string(plainData), nil
}

func main() {
	key := []byte("0123456789abcdef0123456789abcdef") // 32 bytes AES-256 key

	plainText := "Hello, World!"
	fmt.Println("Plain Text:", plainText)

	cipherText, err := encrypt(plainText, key)
	if err != nil {
		fmt.Println("Encryption Error:", err)
		return
	}
	fmt.Println("Encrypted Text:", cipherText)

	decryptedText, err := decrypt(cipherText, key)
	if err != nil {
		fmt.Println("Decryption Error:", err)
		return
	}
	fmt.Println("Decrypted Text:", decryptedText)
}
