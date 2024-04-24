package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/scwallet"
)

func main() {
	if false {
		fmt.Println(scwallet.ErrPairingPasswordNeeded)
		fmt.Println(scwallet.ErrPINNeeded)
		fmt.Println(scwallet.ErrPINUnblockNeeded)
		fmt.Println(scwallet.ErrAlreadyOpen)
		fmt.Println(scwallet.ErrPubkeyMismatch)
		fmt.Println(scwallet.DerivationSignatureHash)

		fmt.Println(scwallet.P1DeriveKeyFromMaster)
		fmt.Println(scwallet.P1DeriveKeyFromParent)
		fmt.Println(scwallet.P1DeriveKeyFromCurrent)
	}
	if true {
		w := scwallet.NewWallet(nil, nil)
		fmt.Println(w)
		s := scwallet.Session{}
		fmt.Println(s.Wallet)
	}
}
