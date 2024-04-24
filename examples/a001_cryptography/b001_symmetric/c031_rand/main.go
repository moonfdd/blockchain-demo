//随机数生成器

package main

import (
	"bytes"
	"crypto/dsa"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	// 随机数生成
	if false {
		lenb := 10
		b := make([]byte, lenb)
		_, err := rand.Read(b)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		// The slice should now contain random bytes instead of only zeroes.
		fmt.Println(bytes.Equal(b, make([]byte, lenb)))

		// Output:
		// false
		return
	}

	// rsa私钥生成
	if false {
		_, err := rsa.GenerateKey(rand.Reader, 1024)
		if err != nil {
			fmt.Println("rsa私钥生成失败", err)
			return
		}
		fmt.Println("rsa私钥生成成功")
		return
	}
	// dsa私钥生成
	if true {
		//Parameters代表秘钥的域参数
		var param dsa.Parameters
		//GenerateParameters函数随机的设置合法的参数到params,
		//根据第三个参数就决定L和N的长度，长度越长，加密强度越高
		err := dsa.GenerateParameters(&param, rand.Reader, dsa.L1024N160)
		if err != nil {
			fmt.Println("dsa中GenerateParameters失败")
			return
		}

		//生成私钥
		var priv dsa.PrivateKey
		priv.Parameters = param
		err = dsa.GenerateKey(&priv, rand.Reader)
		if err != nil {
			fmt.Println("dsa私钥生成失败", err)
			return
		}
		fmt.Println("dsa私钥生成成功")
	}
}
