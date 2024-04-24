package main

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	if true {
		a, _ := abi.JSON(strings.NewReader(`[
			{ "type" : "event", "name" : "Balance", "inputs": [{ "name" : "in", "type": "uint256" }] },
			{ "type" : "event", "name" : "Check", "inputs": [{ "name" : "t", "type": "address" }, { "name": "b", "type": "uint256" }] },
			{ "type" : "event", "name" : "Transfer", "inputs": [{ "name": "from", "type": "address", "indexed": true }, { "name": "to", "type": "address", "indexed": true }, { "name": "value", "type": "uint256" }] }
			]`))
		for name, event := range a.Events {
			fmt.Println(name, event.String())
		}

	}
	fmt.Println("")
}
