package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/consensus"
)

func main() {
	if true {
		fmt.Println(consensus.ErrUnknownAncestor)
		fmt.Println(consensus.ErrPrunedAncestor)
		fmt.Println(consensus.ErrFutureBlock)
		fmt.Println(consensus.ErrInvalidNumber)
		fmt.Println(consensus.ErrInvalidTerminalBlock)
	}
}
