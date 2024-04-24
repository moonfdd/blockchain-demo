package main

import (
	"encoding/hex"
	"fmt"

	"github.com/tjfoc/gmsm/sm3"
)

func main2() {
	hash := sm3.New()
	hash.Write([]byte("微信公众号【福大大架构师每日一题】"))
	result := hash.Sum(nil)
	fmt.Println(hex.EncodeToString(result))
	fmt.Println("length = ", len(result)*8)
}

func main() {
	result := sm3.Sm3Sum([]byte("微信公众号【福大大架构师每日一题】"))
	fmt.Println("length = ", len(result)*8)
	fmt.Println(hex.EncodeToString(result))
}
