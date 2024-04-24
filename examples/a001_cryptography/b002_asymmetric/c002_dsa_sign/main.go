package main

import (
	"crypto/dsa"
	"crypto/rand"
	"fmt"
)

/*
验证签名的作用：
1.保证数据的完整性
2.确保数据的来源
*/
func main() {
	//Parameters代表秘钥的域参数
	var param dsa.Parameters
	//GenerateParameters函数随机的设置合法的参数到params,
	//根据第三个参数就决定L和N的长度，长度越长，加密强度越高
	dsa.GenerateParameters(&param, rand.Reader, dsa.L1024N160)

	//生成私钥
	var priv dsa.PrivateKey
	priv.Parameters = param
	dsa.GenerateKey(&priv, rand.Reader)

	//利用私钥签名数据
	message := []byte("微信公众号【福大大架构师每日一题】")
	r, s, _ := dsa.Sign(rand.Reader, &priv, message)

	//通过私钥获取公钥
	pub := priv.PublicKey
	message1 := []byte("微信公众号【福大大架构师每日一题】")
	//使用公钥验证签名
	if dsa.Verify(&pub, message1, r, s) {
		fmt.Println("验证通过!")
	} else {
		fmt.Println("验证失败!")
	}
}
