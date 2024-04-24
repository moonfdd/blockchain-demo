package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := md5.New()
	hash.Write([]byte("微信公众号【福大大架构师每日一题】"))
	result := hash.Sum(nil)
	fmt.Println(hex.EncodeToString(result))
	hash.Reset()
	hash.Write([]byte("微信公众号【福大大架构师每日一题】"))
	result = hash.Sum(nil)
	fmt.Println(hex.EncodeToString(result))
}
