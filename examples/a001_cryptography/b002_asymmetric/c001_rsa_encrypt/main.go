package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
)

func RsaGenKey(bits int) error {
	//GenerateKey函数使用随机数生成器生成一对指定长度的公钥和私钥
	//rand.Reader是一个全局，共享的密码随机生成器
	privKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	//x509是通用的整数格式：序列号  签名算法  颁发者 有效时间  持有者  公钥
	//PKCS:RSA实验室与其他安全系统开发商为促进公钥密码的发展而指定的一系列标准
	priStream := x509.MarshalPKCS1PrivateKey(privKey)
	//将私钥字符串设置pem格式的块中
	/*
		pem是一种整数或私钥的格式：
		---------------BEGIN  RSA Private Key---------------


		--------------END RSA Private Key-------------------
	*/
	block := pem.Block{
		Type:  "RSA Private Key",
		Bytes: priStream,
	}
	privFile, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	defer privFile.Close()
	//将块编码到文件
	err = pem.Encode(privFile, &block)
	if err != nil {
		panic(err)
	}
	//从私钥中获取公钥
	pubKey := privKey.PublicKey
	//将公钥序列化
	pubStream := x509.MarshalPKCS1PublicKey(&pubKey)
	//将公钥设置到pem块中
	block = pem.Block{
		Type:  "RSA Public Key",
		Bytes: pubStream,
	}
	pubFile, err := os.Create("publiv.pem")
	defer pubFile.Close()
	if err != nil {
		panic(err)
	}
	err = pem.Encode(pubFile, &block)
	if err != nil {
		panic(err)
	}
	return nil
}

// 公钥加密
func RsaPublicEncrypt(src, pathName []byte) ([]byte, error) {
	file, err := os.Open(string(pathName))
	msg := []byte("")
	if err != nil {
		return msg, err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return msg, err
	}
	//创建切片，用于存储公钥
	recvBuf := make([]byte, info.Size())
	//读取公钥
	file.Read(recvBuf)
	//将得到的公钥反序列化
	//参数一：存储公钥的切片， 参数二：剩余未解密的数据
	block, _ := pem.Decode(recvBuf)
	//使用x509将编码之后的公钥解析出来
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return msg, err
	}
	msg, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

// 使用私钥解密
func RsaPrivateDecrypt(src []byte, pathName string) ([]byte, error) {
	msg := []byte("")
	file, err := os.Open(pathName)
	if err != nil {
		return msg, err
	}
	info, err := file.Stat()
	if err != nil {
		return msg, err
	}
	//创建切片，用于存储公钥
	recvBuf := make([]byte, info.Size())
	//读取公钥
	file.Read(recvBuf)
	block, _ := pem.Decode(recvBuf)
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return msg, err
	}
	msg, err = rsa.DecryptPKCS1v15(rand.Reader, privKey, src)
	if err != nil {
		panic(err)
	}
	return msg, nil
}

func main() {
	err := RsaGenKey(1024)
	if err != nil {
		fmt.Println("秘钥对生成失败!")
	} else {
		fmt.Println("秘钥对生成成功!")
	}

	src := []byte("微信公众号【福大大架构师每日一题】")
	//src := []byte("微信公众号【福大大架构师每日一题】如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：如果您不希望使用sm2.ParsePKCS8PrivateKey()方法，可以直接将固定的私钥以字节数组形式导入。以下是修改后的示例代码：")
	encrypt_msg, _ := RsaPublicEncrypt(src, []byte("publiv.pem"))
	fmt.Println("encrypt_msg = ", hex.EncodeToString(encrypt_msg))
	decrypt_msg, _ := RsaPrivateDecrypt(encrypt_msg, "private.pem")
	fmt.Println("decrypt_msg = ", string(decrypt_msg))
}
