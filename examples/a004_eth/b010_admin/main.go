package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// create instance of ethclient and assign to cl
	cl, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer cl.Close()
	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		fmt.Println("获取ChainID失败:", err)
		return
	}
	fmt.Println("chainid:", chainid)

	addr := common.HexToAddress("0x7e5f4552091a69125d5dfcb7b8c2659029395bdf")
	nonce, err := cl.NonceAt(context.Background(), addr, big.NewInt(0))
	if err != nil {
		fmt.Println("NonceAt失败:", err)
		return
	}
	fmt.Println("nonce:", nonce)

	blockNumber, err := cl.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Failed to retrieve block number:", err)
		return
	}
	fmt.Println("blockNumber:", blockNumber)
	blockNumberBig := big.NewInt(int64(blockNumber))
	if false {
		eventSignatureBytes := []byte("Transfer(address,address,uint256)")
		eventSignaturehash := crypto.Keccak256Hash(eventSignatureBytes)

		q := ethereum.FilterQuery{
			FromBlock: new(big.Int).Sub(blockNumberBig, big.NewInt(10)),
			ToBlock:   blockNumberBig,
			Topics: [][]common.Hash{
				{eventSignaturehash},
			},
		}

		logs, err := cl.FilterLogs(context.Background(), q)
		if err != nil {
			fmt.Println("FilterLogs失败:", err)
			return
		}
		fmt.Println("FilterLogs成功:", logs)
	}

}
