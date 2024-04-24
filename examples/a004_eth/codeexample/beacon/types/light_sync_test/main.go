package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/beacon/types"
)

func main() {
	if false {
		bd := types.BootstrapData{}
		err := bd.Validate()
		fmt.Println(err)
	}
	if false {
		bd := types.LightClientUpdate{}
		err := bd.Validate()
		fmt.Println(err)
	}
	if true {
		bd := types.UpdateScore{}
		bd.SignerCount = 1
		b := bd.BetterThan(types.UpdateScore{})
		fmt.Println(b)
	}
}
