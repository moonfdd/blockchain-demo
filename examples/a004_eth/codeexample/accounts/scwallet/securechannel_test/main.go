package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/scwallet"
)

func main() {
	if true {
		w, err := scwallet.NewSecureChannelSession(nil, nil)
		if err != nil {
			fmt.Println("NewSecureChannelSession失败", err)
		}
		fmt.Println(w)
	}
	fmt.Println("")
}
