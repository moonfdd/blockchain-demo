package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = `{"address":"7e5f4552091a69125d5dfcb7b8c2659029395bdf","crypto":{"cipher":"aes-128-ctr","ciphertext":"e468ad3ce2a6ae1436f8b50259f7cef84e688df2b08575f93eda7067492fec45","cipherparams":{"iv":"80697348d506166676cc27cc10aa00e0"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"806d6da06fbb4f2480e4c93f7bd82d5575f2da3ce382b469e6ed293e9e5cb52c"},"mac":"5fb6e482ee88c6a99972ace20827d547d2245249c2c3cfa8a8c8fea8bc82b771"},"id":"509c8fb9-b781-4c3b-bce7-e11b75e03ac3","version":3}`

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100a1565b60405180910390f35b610073600480360381019061006e91906100ed565b61007e565b005b60008054905090565b8060008190555050565b6000819050919050565b61009b81610088565b82525050565b60006020820190506100b66000830184610092565b92915050565b600080fd5b6100ca81610088565b81146100d557600080fd5b50565b6000813590506100e7816100c1565b92915050565b600060208284031215610103576101026100bc565b5b6000610111848285016100d8565b9150509291505056fea2646970667358221220d930a0e982648452bbdb8c6bcff7d8f073432d0f1b9155afb04547759545c4dc64736f6c63430008130033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

func main() {
	// 部署合约
	if false {
		// Create an IPC based RPC connection to a remote node and an authorized transactor
		conn, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}
		fmt.Println("Dial成功")
		chainID, err := conn.ChainID(context.Background())
		if err != nil {
			log.Panic(err)
		}
		_ = chainID

		auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), "123456", chainID)
		// auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
		if err != nil {
			log.Fatalf("Failed to create authorized transactor: %v", err)
		}
		fmt.Println("NewTransactor成功")
		_ = auth

		parsed, err := StorageMetaData.GetAbi()
		if err != nil {
			log.Fatalf("StorageMetaData.GetAbi: %v", err)
		}
		if parsed == nil {
			log.Fatalf("parsed == nil")
		}
		address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), conn)
		if err != nil {
			log.Fatalf("DeployContract: %v", err)
		}
		_ = contract

		fmt.Println("DeployStorage成功")
		fmt.Printf("Contract pending deploy: 0x%x\n", address)
		fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

		// WaitDeployed成功
		_, err = bind.WaitDeployed(context.Background(), conn, tx)
		if err != nil {
			log.Fatalf("WaitDeployed: %v", err)
		}
		fmt.Println("WaitDeployed成功")

	}
	// 调用合约
	if true {
		// Create an IPC based RPC connection to a remote node and an authorized transactor
		conn, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}
		fmt.Println("Dial成功")
		chainID, err := conn.ChainID(context.Background())
		if err != nil {
			log.Panic(err)
		}
		_ = chainID

		var result bool
		conn.Client().CallContext(context.Background(), &result, "personal_unlockAccount", common.HexToAddress("0x7e5f4552091a69125d5dfcb7b8c2659029395bdf"), "123456", big.NewInt(100000000))

		auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), "123456", chainID)
		// auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
		if err != nil {
			log.Fatalf("Failed to create authorized transactor: %v", err)
		}
		fmt.Println("NewTransactor成功")
		_ = auth
		parsed, err := StorageMetaData.GetAbi()
		if err != nil {
			if err != nil {
				log.Fatalf("StorageMetaData.GetAbi: %v", err)
			}
		}
		_ = parsed
		contract := bind.NewBoundContract(common.HexToAddress("0xf5b492adbc2ef7ae61825a69e34dc75e2e3a83e4"), *parsed, conn, conn, conn)
		_ = contract
		var out []interface{}
		contract.Call(nil, &out, "retrieve")
		out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		fmt.Println("retrieve = ", out0)

		tx, err := contract.Transact(auth, "store", big.NewInt(25))
		if err != nil {
			log.Fatalf("store.Store: %v", err)
		}
		_ = tx
		fmt.Println("store.Store ok")

		// WaitMined成功
		_, err = bind.WaitMined(context.Background(), conn, tx)
		if err != nil {
			log.Fatalf("WaitMined: %v", err)
		}
		fmt.Println("WaitMined")

		contract.Call(nil, &out, "retrieve")
		out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		fmt.Println("retrieve = ", out0)

	}
	fmt.Println("")
}
