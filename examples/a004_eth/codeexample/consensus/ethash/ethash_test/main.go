package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/consensus/ethash"
)

func main() {
	if false {
		f := ethash.NewFaker()
		fmt.Println(f)
	}
	if false {
		f := ethash.NewFakeFailer(33)
		fmt.Println(f)
	}
	if false {
		f := ethash.NewFakeDelayer(33)
		fmt.Println(f)
	}
	if true {
		f := ethash.NewFullFaker()
		fmt.Println(f)
	}
}
