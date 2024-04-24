package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

func main() {
	key := []byte("moonfdd1moonfdd1moonfdd1") //24个字符
	data := []byte("微信公众号【福大大架构师每日一题】")
	fmt.Println("原文：", string(data))
	encrypt_msg := Encrypt3DES(data, key)
	fmt.Println("encrypt_msg = ", hex.EncodeToString(encrypt_msg))
	decrypt_msg := Decrypt3DES(encrypt_msg, key)
	fmt.Println("decrypt_msg = ", string(decrypt_msg))
}

// 使用3des加密
func Encrypt3DES(src, key []byte) []byte {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	src = PaddingText(src, block.BlockSize())
	iv := []byte("12345678")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// 解密
func Decrypt3DES(src, key []byte) []byte {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)
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
