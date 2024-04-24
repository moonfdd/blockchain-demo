package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

// 密钥协商Key Agreement
// 在Go语言中实现密钥协商可以使用Diffie-Hellman算法或者其他加密算法。下面是一个基于Diffie-Hellman算法的示例代码：
func main() {
	// 选择椭圆曲线
	curve := elliptic.P256()

	// 生成私钥
	privKeyA, _ := ecdsa.GenerateKey(curve, rand.Reader)

	// 获取公钥坐标点
	pubKeyAx, pubKeyAy := privKeyA.PublicKey.X, privKeyA.PublicKey.Y

	// 生成私钥
	privKeyB, _ := ecdsa.GenerateKey(curve, rand.Reader)

	// 获取公钥坐标点
	pubKeyBx, pubKeyBy := privKeyB.PublicKey.X, privKeyB.PublicKey.Y

	// 计算共享密钥
	sharedKeyAx, _ := curve.ScalarMult(pubKeyBx, pubKeyBy, privKeyA.D.Bytes())
	sharedKeyBx, _ := curve.ScalarMult(pubKeyAx, pubKeyAy, privKeyB.D.Bytes())

	// 检查双方计算的共享密钥是否一致
	if sharedKeyAx.Cmp(sharedKeyBx) == 0 {
		fmt.Println("双方计算的共享密钥相同")
		// 将共享密钥用于进一步的加密通信
	} else {
		fmt.Println("双方计算的共享密钥不相同")
	}
}
