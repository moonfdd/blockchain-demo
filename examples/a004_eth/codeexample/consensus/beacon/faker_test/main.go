package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/core/types"
)

func main() {
	if true {
		b := beacon.NewFaker()
		fmt.Println(b)
		fmt.Println(b.CalcDifficulty(nil, 2, &types.Header{}))
	}
}
