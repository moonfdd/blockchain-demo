package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/rawdb"
)

func main() {
	if true {
		merger := consensus.NewMerger(rawdb.NewMemoryDatabase())
		fmt.Println(merger)
		fmt.Println(merger.TDDReached())
		fmt.Println(merger.PoSFinalized())
	}
}
