package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"io/ioutil"
	"strings"
)

const KEY_SIZE int = 24

// 填充最后一个分组
// src:待填充的数据，   blockSize：分组大小
func PaddingText(src []byte, blockSize int) []byte {
	//求出最后一个分组需要填充的字节数
	padding := blockSize - len(src)%blockSize
	//创建新的切片，切片的字节数为padding。
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	//将新创建的切片和待填充的数据进行拼接
	nextText := append(src, padText...)
	return nextText
}

// 删除尾部填充数据
func UnPaddingText(src []byte) []byte {
	//获取待处理的数据的长度
	len := len(src)
	//去除最后一个字符
	number := int(src[len-1])
	newText := src[:len-number]
	return newText
}

// 使用3des加密
func Encrypt3(src, key []byte) []byte {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	src = PaddingText(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// 解密
func Decrypt3(src, key []byte) []byte {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	dst = UnPaddingText(dst)
	return dst
}

func genKey3(key []byte) []byte {
	//用于存储最终的秘钥
	kkey := make([]byte, 0, KEY_SIZE)
	//获取原始秘钥的长度
	length := len(key)
	if length > KEY_SIZE {
		kkey = append(kkey, key[:KEY_SIZE]...)
	} else {
		//用指定的长度对实际秘钥长度进行求商
		div := KEY_SIZE / length
		//用指定的长度对实际秘钥长度进行求余
		mod := KEY_SIZE % length
		for i := 0; i < div; i++ {
			kkey = append(kkey, key...)
		}
		kkey = append(kkey, key[:mod]...)
	}
	return kkey
}

func main() {
	var command string
	var filename string
	fmt.Print("请输入命令(加密|解密):")
	fmt.Scanln(&command)
	if command == "加密" {
		fmt.Print("请输入需要被加密的文件的路径:")
		fmt.Scanln(&filename)
		fmt.Print("请输入秘钥:")
		var password string
		fmt.Scanln(&password)
		fmt.Print("请输入确认密码:")
		var confirmpassword string
		fmt.Scanln(&confirmpassword)
		if !bytes.Equal([]byte(password), []byte(confirmpassword)) {
			fmt.Println("两次输入的密码不一致，请重新输入!")
		} else {
			key := genKey3([]byte(password))
			info, _ := ioutil.ReadFile(filename)
			dst := Encrypt3(info, key)
			index := strings.LastIndex(filename, ".")
			newfilename := filename[:index] + "_encrypted" + filename[index:]
			ioutil.WriteFile(newfilename, dst, 0777)
			fmt.Println("已生成加密问价" + newfilename + ",请妥善保管您的秘钥!")
		}
	} else if command == "解密" {
		fmt.Print("请输入需要解密的文件的路径:")
		fmt.Scanln(&filename)
		fmt.Print("请输入秘钥:")
		var password string
		fmt.Scanln(&password)
		key := genKey3([]byte(password))
		info, _ := ioutil.ReadFile(filename)
		src := Decrypt3(info, key)
		if len(src) == 0 {
			fmt.Println("秘钥不对，请重新输入!")
		} else {
			index := strings.LastIndex(filename, ".")
			newfilename := filename[:index] + "_decrypted" + filename[index:]
			ioutil.WriteFile(newfilename, src, 0777)
			fmt.Println("已生成解密文件" + newfilename + ",请妥善保管您的秘钥!")
		}
	}
}
