package main

import (
	"crypto/sha256"
	"fmt"
)

/*
go语言sha256包中实现了两种哈希函数，分别是sha256hesha224
*/

func main() {
	hash := sha256.New()
	hash.Write([]byte("微信公众号【福大大架构师每日一题】"))
	result := hash.Sum(nil)
	fmt.Printf("%x\n", result)
}
