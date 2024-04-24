package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/consensus/misc"
)

func main() {
	if true {
		fmt.Println(misc.ErrBadProDAOExtra)
		fmt.Println(misc.ErrBadNoDAOExtra)
		// fmt.Println(misc.VerifyDAOHeaderExtraData(nil, nil))
		// misc.ApplyDAOHardFork(nil)
	}
}
