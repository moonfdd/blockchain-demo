package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = `{"address":"7e5f4552091a69125d5dfcb7b8c2659029395bdf","crypto":{"cipher":"aes-128-ctr","ciphertext":"e468ad3ce2a6ae1436f8b50259f7cef84e688df2b08575f93eda7067492fec45","cipherparams":{"iv":"80697348d506166676cc27cc10aa00e0"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"806d6da06fbb4f2480e4c93f7bd82d5575f2da3ce382b469e6ed293e9e5cb52c"},"mac":"5fb6e482ee88c6a99972ace20827d547d2245249c2c3cfa8a8c8fea8bc82b771"},"id":"509c8fb9-b781-4c3b-bce7-e11b75e03ac3","version":3}`

func main() {
	//部署合约
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

		auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), "123456", chainID)
		// auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
		if err != nil {
			log.Fatalf("Failed to create authorized transactor: %v", err)
		}
		fmt.Println("NewTransactor成功")
		// Deploy the contract passing the newly created `auth` and `conn` vars
		address, tx, instance, err := DeployStorage(auth, conn)
		// address, tx, instance, err := DeployStorage(auth, conn)
		if err != nil {
			log.Fatalf("Failed to deploy new storage contract: %v", err)
		}
		fmt.Println("DeployStorage成功")
		fmt.Printf("Contract pending deploy: 0x%x\n", address)
		fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

		// WaitDeployed成功
		_, err = bind.WaitDeployed(context.Background(), conn, tx)
		if err != nil {
			log.Fatalf("WaitDeployed: %v", err)
		}
		fmt.Println("WaitDeployed成功")

		time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P
		_ = instance
		// function call on `instance`. Retrieves pending name
	}
	//调用合约
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
		// Instantiate the contract and display its name
		// NOTE update the deployment address!
		// 0x44399fe087f1955701001dff92ce00e2c1f0aa449ed630b23a5cb6ba44c9cbbe
		store, err := NewStorage(common.HexToAddress("0x8913265b93d7e3d9dcc1913616e1f874b741512d"), conn)
		if err != nil {
			log.Fatalf("Failed to instantiate Storage contract: %v", err)
		}

		auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), "123456", chainID)
		// auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
		if err != nil {
			log.Fatalf("Failed to create authorized transactor: %v", err)
		}
		fmt.Println("auth.Value = ", auth.Value)

		value, err := store.Retrieve(nil)
		if err != nil {
			log.Fatalf("Failed to retrieve value1: %v", err)
		}
		fmt.Println("Value1: ", value)

		var tx *types.Transaction
		tx, err = store.Store(auth, big.NewInt(56))
		if err != nil {
			log.Fatalf("store.Store: %v", err)
		}
		fmt.Println("store.Store ok")
		_ = tx

		fmt.Println("tx = ", tx)

		// 失败
		if false {
			_, err = bind.WaitDeployed(context.Background(), conn, tx)
			if err != nil {
				log.Fatalf("WaitDeployed: %v", err)
			}
		}

		value, err = store.Retrieve(nil)
		if err != nil {
			log.Fatalf("Failed to retrieve value2: %v", err)
		}
		fmt.Println("Value2: ", value)
	}
	// 合约session
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

		store, err := NewStorage(common.HexToAddress("0x21e6fc92f93c8a1bb41e2be64b4e1f88a54d3576"), conn)
		if err != nil {
			log.Fatalf("Failed to instantiate a Storage contract: %v", err)
		}

		auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), "123456", chainID)
		// auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
		if err != nil {
			log.Fatalf("Failed to create authorized transactor: %v", err)
		}

		// Wrap the Storage contract instance into a session
		session := &StorageSession{
			Contract: store,
			CallOpts: bind.CallOpts{
				Pending: true,
			},
			TransactOpts: bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasLimit: 3141592,
			},
		}
		// Call the previous methods without the option parameters
		_, err = session.Store(big.NewInt(70))
		if err != nil {
			log.Fatalf("session.Store: %v", err)
		}
		fmt.Println("session.Store成功")

		_, err = session.Store(big.NewInt(71))
		if err != nil {
			log.Fatalf("session.Store: %v", err)
		}
		fmt.Println("session.Store成功")

		value, err := session.Retrieve()
		if err != nil {
			log.Fatalf("session.Store: %v", err)
		}
		fmt.Println("session.Retrieve:", value)
	}

}
