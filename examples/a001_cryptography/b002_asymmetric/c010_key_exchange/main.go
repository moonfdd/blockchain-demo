package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

// 如果你想实现Golang中的密钥交换（Key Exchange）算法，可以使用Diffie-Hellman或者ECDH（Elliptic Curve Diffie-Hellman）算法。以下是一个使用ECDH算法进行密钥交换的示例代码
// 这个代码示例中，我们首先使用ecdsa.GenerateKey函数生成两对ECC私钥对，分别代表参与交换的双方（例如Alice和Bob）。然后从私钥中提取出公钥的x、y坐标。接着，通过在曲线上执行点乘运算，计算出双方各自的共享密钥。最后打印出这两个共享密钥。
// 请注意，示例中使用了P256曲线（elliptic.P256()），你也可以选择其他曲线，例如P384或者P521等。此外，实际应用中可能需要更多的错误处理和安全性措施。
func main() {
	// 生成ECC参数
	curve := elliptic.P256()
	privateKeyA, _ := ecdsa.GenerateKey(curve, rand.Reader)
	privateKeyB, _ := ecdsa.GenerateKey(curve, rand.Reader)

	// 公钥A和公钥B的x、y坐标
	pubKeyAx, pubKeyAy := privateKeyA.PublicKey.X.Bytes(), privateKeyA.PublicKey.Y.Bytes()
	pubKeyBx, pubKeyBy := privateKeyB.PublicKey.X.Bytes(), privateKeyB.PublicKey.Y.Bytes()

	// A计算共享密钥
	x1, _ := curve.ScalarMult(big.NewInt(0).SetBytes(pubKeyBx), big.NewInt(0).SetBytes(pubKeyBy), privateKeyA.D.Bytes())
	keyA := x1.Bytes()

	// B计算共享密钥
	x2, _ := curve.ScalarMult(big.NewInt(0).SetBytes(pubKeyAx), big.NewInt(0).SetBytes(pubKeyAy), privateKeyB.D.Bytes())
	keyB := x2.Bytes()

	fmt.Printf("共享密钥(A): %x\n", keyA)
	fmt.Printf("共享密钥(B): %x\n", keyB)
}
