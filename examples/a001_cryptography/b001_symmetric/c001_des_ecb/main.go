package main

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	common.Report("1", 2, "3.0")
	return
	src := []byte("微信公众号【福大大架构师每日一题】")
	fmt.Println("明文：", string(src))
	key := []byte("moonfdd1") //64bit
	ret := EncryptDES(src, key)
	fmt.Println("密文：", hex.EncodeToString(ret))
	ret = DecryptDes(ret, key)
	fmt.Println("明文：", string(ret))
}

// 加密
func EncryptDES(src, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	length := block.BlockSize()
	src = PaddingText(src, length)
	dst := make([]byte, len(src))
	out := dst
	for len(src) > 0 {
		//每次加密8字节
		block.Encrypt(out, src[:length])
		//去除已被加密的数据
		src = src[length:]
		out = out[length:]
	}
	return dst
}

// 解密
func DecryptDes(src, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	length := block.BlockSize()
	dst := make([]byte, len(src))
	out := dst
	for len(src) > 0 {
		block.Decrypt(out, src[:length])
		src = src[length:]
		out = out[length:]
	}
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
