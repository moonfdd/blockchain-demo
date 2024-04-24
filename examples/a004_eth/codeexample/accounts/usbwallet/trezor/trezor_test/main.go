package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/usbwallet/trezor"
)

func main() {
	if true {
		fmt.Println(trezor.Type(&trezor.Success{}))
	}
	if false {
		fmt.Println(trezor.Name(2))
	}
}
