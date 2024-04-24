package main

import (
	"fmt"

	"math/big"

	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/types"
)

func main() {
	if false {
		fmt.Println(ethash.FrontierBlockReward)
		fmt.Println(ethash.ByzantiumBlockReward)
		fmt.Println(ethash.ConstantinopleBlockReward)
	}
	if true {
		b := ethash.NewFaker()
		fmt.Println(b)
		fmt.Println(b.Author(&types.Header{Difficulty: big.NewInt(1)}))
	}
}
