package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/usbwallet"
)

func main() {
	if false {
		fmt.Println(usbwallet.LedgerScheme)
		fmt.Println(usbwallet.TrezorScheme)
	}
	if false {
		h, err := usbwallet.NewLedgerHub()
		if err != nil {
			fmt.Println("NewLedgerHub失败 = ", err)
			return
		}
		fmt.Println(h)
	}
	if false {
		h, err := usbwallet.NewTrezorHubWithHID()
		if err != nil {
			fmt.Println("NewTrezorHubWithHID失败 = ", err)
			return
		}
		fmt.Println(h)
	}
	if true {
		h, err := usbwallet.NewTrezorHubWithWebUSB()
		if err != nil {
			fmt.Println("NewTrezorHubWithWebUSB失败 = ", err)
			return
		}
		fmt.Println(h)
	}
}
