package main

import "github.com/ethereum/go-ethereum/beacon/types"

func main() {
	h := types.SyncAggregate{}
	d, err := h.MarshalJSON()
	if err != nil {
		panic(err)
	}
	println(string(d))
}
