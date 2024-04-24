package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
)

func main() {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}

	// Since we are using a simulated backend, we will get the chain id
	// from the same place that the simulated backend gets it.
	chainID := params.AllDevChainProtocolChanges.ChainID

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		log.Fatalf("Failed to make transactor: %v", err)
	}

	sim := simulated.NewBackend(map[common.Address]core.GenesisAccount{
		auth.From: {Balance: big.NewInt(9e18)},
	})

	time.Sleep(20 * time.Second)

	fmt.Println(sim.Client().BlockNumber(context.Background()))

	_, tx, store, err := DeployStorage(auth, sim.Client())
	if err != nil {
		log.Fatalf("Failed to deploy smart contract: %v", err)
	}

	fmt.Printf("Deploy pending: 0x%x\n", tx.Hash())

	sim.Commit()

	tx, err = store.Store(auth, big.NewInt(420))
	if err != nil {
		log.Fatalf("Failed to call store method: %v", err)
	}
	fmt.Printf("State update pending: 0x%x\n", tx.Hash())

	sim.Commit()

	val, err := store.Retrieve(nil)
	if err != nil {
		log.Fatalf("Failed to call retrieve method: %v", err)
	}
	fmt.Printf("Value: %v\n", val)

	if true {
		// var res []common.Address
		// if err := sim.Client().Call(&res, "account_list"); err != nil {
		// 	fmt.Println("account_list失败:", err)
		// 	return
		// }
		// fmt.Println("account_list:", res)
	}
}
