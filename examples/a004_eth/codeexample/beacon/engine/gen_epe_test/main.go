package main

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/beacon/engine"
)

func main() {
	a := &engine.ExecutionPayloadEnvelope{}
	a.BlockValue = big.NewInt(33)
	a.Override = true
	d, err := a.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(d))

	d, err = json.MarshalIndent(a, "", " ")
	fmt.Println(string(d))
}
