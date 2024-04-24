package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	// 生成2048位的RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Failed to generate private key:", err)
		return
	}

	// 从私钥中提取公钥
	publicKey := &privateKey.PublicKey

	// 将私钥编码为PEM格式
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	privateKeyFile, err := os.Create("private.pem")
	if err != nil {
		fmt.Println("Failed to create private key file:", err)
		return
	}
	defer privateKeyFile.Close()
	pem.Encode(privateKeyFile, privateKeyPEM)

	// 将公钥编码为PEM格式
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("Failed to marshal public key:", err)
		return
	}
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicKeyFile, err := os.Create("public.pem")
	if err != nil {
		fmt.Println("Failed to create public key file:", err)
		return
	}
	defer publicKeyFile.Close()
	pem.Encode(publicKeyFile, publicKeyPEM)

	fmt.Println("RSA key pair generated successfully.")
}
