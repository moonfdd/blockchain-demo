package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/usbwallet"
)

func main() {
	if true {
		fmt.Println(usbwallet.ErrTrezorPINNeeded)
		fmt.Println(usbwallet.ErrTrezorPassphraseNeeded)
	}
}
