package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// 生成一个新的私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 要签名的数据
	data := []byte("Hello, Ethereum!")
	data = common.HexToHash("0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef").Bytes()
	data = crypto.Keccak256([]byte("foo"))
	// 对数据进行签名
	sig, err := crypto.Sign(data, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 打印签名
	fmt.Println("Signature:", hexutil.Encode(sig))

	// 验证签名
	publicKey := privateKey.Public()
	pubKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	// publicKeyBytes := crypto.CompressPubkey(pubKeyECDSA)
	fmt.Println(len(publicKeyBytes))
	verified := crypto.VerifySignature(publicKeyBytes, data, sig[0:64])
	fmt.Println("Signature verified:", verified)
}
