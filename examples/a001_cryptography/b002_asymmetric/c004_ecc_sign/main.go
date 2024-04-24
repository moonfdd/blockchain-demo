package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	message := []byte("微信公众号【福大大架构师每日一题】")
	//参数一：曲线类型
	privateKey, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	//获取公钥
	pub := privateKey.PublicKey

	//散列明文
	digest := sha256.Sum256(message)
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, digest[:])

	message1 := []byte("微信公众号【福大大架构师每日一题】")
	digest = sha256.Sum256(message1)
	flag := ecdsa.Verify(&pub, digest[:], r, s)
	if flag {
		fmt.Println("验证成功")
	} else {
		fmt.Println("验证失败")
	}

}
