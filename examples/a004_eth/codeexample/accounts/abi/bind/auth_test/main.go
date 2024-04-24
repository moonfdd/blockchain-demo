package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/external"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const key = `{"address":"7e5f4552091a69125d5dfcb7b8c2659029395bdf","crypto":{"cipher":"aes-128-ctr","ciphertext":"e468ad3ce2a6ae1436f8b50259f7cef84e688df2b08575f93eda7067492fec45","cipherparams":{"iv":"80697348d506166676cc27cc10aa00e0"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"806d6da06fbb4f2480e4c93f7bd82d5575f2da3ce382b469e6ed293e9e5cb52c"},"mac":"5fb6e482ee88c6a99972ace20827d547d2245249c2c3cfa8a8c8fea8bc82b771"},"id":"509c8fb9-b781-4c3b-bce7-e11b75e03ac3","version":3}`

func main() {
	// 输出错误类型
	if false {
		// 用户未指定链ID时返回的错误
		fmt.Println(bind.ErrNoChainID)
		// 账户未正确解锁时返回的错误
		fmt.Println(bind.ErrNotAuthorized)
		return
	}
	// 从加密的JSON密钥流和相关的密码短语中创建一个交易签名者。
	if false {
		chainID := big.NewInt(666123)
		transactOpts, err := bind.NewTransactorWithChainID(strings.NewReader(key), "123456", chainID)
		if err != nil {
			fmt.Println("NewTransactorWithChainID失败", err)
			return
		}
		fmt.Println("NewTransactorWithChainID成功", transactOpts)
		return
	}
	// 从解密的keystore中轻松创建交易签名者。否决
	if false {
		chainID := big.NewInt(666123)
		ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
		transactOpts, err := bind.NewKeyStoreTransactorWithChainID(ks, ks.Accounts()[0], chainID)
		if err != nil {
			fmt.Println("NewKeyStoreTransactorWithChainID失败", err)
			return
		}
		fmt.Println("NewKeyStoreTransactorWithChainID成功", transactOpts)
		return
	}
	// 从单个私钥快速创建一个交易签名者。
	if false {
		chainID := big.NewInt(666123)
		privateKeyHex := "0x0000000000000000000000000000000000000000000000000000000000000001"
		privateKeyBytes := common.FromHex(privateKeyHex)
		privateKey, err := crypto.ToECDSA(privateKeyBytes)
		if err != nil {
			log.Fatal(err)
		}

		transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			fmt.Println("NewKeyedTransactorWithChainID失败", err)
			return
		}
		fmt.Println("NewKeyedTransactorWithChainID成功", transactOpts)
		return
	}
	// 创建一个具有clef后端的交易签名者
	// clef --keystore datadirprivate/keystore --chainid 666123 --http
	if true {
		clef, err := external.NewExternalSigner("http://127.0.0.1:8550") //需要在clef客户端上确认y
		if err != nil {
			fmt.Println("NewExternalSigner失败", err)
			return
		}
		transactOpts := bind.NewClefTransactor(clef, clef.Accounts()[0])
		fmt.Println("NewClefTransactor成功", transactOpts)
		return
	}
	fmt.Println("Hello World")
}
