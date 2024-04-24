package main

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"

	"github.com/tjfoc/gmsm/sm4"
)

// sm4加密
// src:明文
// key:秘钥
func EncryptSm4(src, key []byte) []byte {
	block, err := sm4.NewCipher(key)
	if err != nil {
		panic(err)
	}

	src = PaddingText(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// sm4解密
// src:密文
// key：秘钥
func DecryptSm4(src, key []byte) []byte {
	block, err := sm4.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	newText := UnPaddingText(dst)
	return newText
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

func main() {
	key := []byte("12345678abcdefgh")
	msg := []byte("微信公众号【福大大架构师每日一题】")
	encrypt_msg := EncryptSm4(msg, key)
	fmt.Println("encrypt_msg = ", hex.EncodeToString(encrypt_msg))
	decrypt_msg := DecryptSm4(encrypt_msg, key)
	fmt.Println("decrypt_msg = ", string(decrypt_msg))
}
