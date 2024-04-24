package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	key := []byte("moonfdd1moonfdd1moonfdd1") //24个字符 16,24,32
	data := []byte("微信公众号【福大大架构师每日一题】")
	fmt.Println("原文：", string(data))
	encrypt_msg := EncryptAes(data, key)
	fmt.Println("encrypt_msg = ", hex.EncodeToString(encrypt_msg))
	decrypt_msg := DecryptAes(encrypt_msg, key)
	fmt.Println("decrypt_msg = ", string(decrypt_msg))
}

// aes加密
func EncryptAes(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	src = PaddingText(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// aes解密
func DecryptAes(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	dst = UnPaddingText(dst)
	return dst
}

// 填充数据
func PaddingText(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	nextText := append(src, padText...)
	return nextText
}

// 删除尾部填充数据
func UnPaddingText(src []byte) []byte {
	len := len(src)
	number := int(src[len-1])
	newText := src[:len-number]
	return newText
}
