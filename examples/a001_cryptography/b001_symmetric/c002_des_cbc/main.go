package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

func main() {
	src := []byte("微信公众号【福大大架构师每日一题】")
	fmt.Println("明文：", string(src))
	key := []byte("moonfdd1") //64bit
	ret := EncryptDES(src, key)
	fmt.Println("密文：", hex.EncodeToString(ret))
	ret = DecryptDes(ret, key)
	fmt.Println("明文：", string(ret))
}

// 使用des算法进行加密
// src:待加密的明文   key:秘钥
func EncryptDES(src, key []byte) []byte {
	//创建cipher.Block.接口，其对应的就是一个加密块
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//获取每个块的大小
	length := block.BlockSize()
	//对最后一组明文进行填充
	src = PaddingText(src, length)
	//初始化向量
	iv := []byte("12345678")
	//创建cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//创建切片，用于存储加密之后的数据
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// 使用des进行解密
// src:待解密的密文    key:秘钥
func DecryptDes(src, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678")
	//创建CBC解密模式
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
