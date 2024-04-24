package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	if true {
		err := &keystore.AmbiguousAddrError{Addr: common.Address{97}, Matches: make([]accounts.Account, 5)}
		err.Matches[0].Address = common.Address{97}
		err.Matches[1].Address = common.Address{98}
		err.Matches[2].Address = common.Address{99}
		err.Matches[3].Address = common.Address{100}
		err.Matches[4].Address = common.Address{101}
		fmt.Println(err)
	}
}
