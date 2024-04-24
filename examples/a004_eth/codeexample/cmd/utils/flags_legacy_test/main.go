package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/cmd/utils"
)

func main() {
	if false {
		fmt.Println(utils.ShowDeprecated)
		fmt.Println(utils.DeprecatedFlags)
	}
	if true {
		fmt.Println(utils.NoUSBFlag)
	}
}
