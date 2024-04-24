package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/types"
)

func main() {
	if true {
		b := beacon.New(ethash.NewFaker())
		fmt.Println(b)
		fmt.Println(b.IsPoSHeader(&types.Header{Difficulty: big.NewInt(1)}))
		fmt.Println(b.InnerEngine())
	}
}
