package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/beacon/merkle"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// 未完成
	if true {
		var m merkle.Value
		err := merkle.VerifyProof(common.Hash{}, 1, nil, m)
		fmt.Println(err)
	}
}
