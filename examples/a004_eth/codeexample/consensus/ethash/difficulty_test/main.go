package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/types"
)

func main() {
	if true {
		fmt.Println(ethash.CalcDifficultyFrontierU256(1, &types.Header{Difficulty: big.NewInt(1)}))
	}
}
