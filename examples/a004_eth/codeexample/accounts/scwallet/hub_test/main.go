package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/scwallet"
)

func main() {
	if false {
		fmt.Println(scwallet.Scheme)
	}
	if true {
		w := scwallet.Hub{}
		w.Subscribe(nil)
		fmt.Println(w)
	}
}
