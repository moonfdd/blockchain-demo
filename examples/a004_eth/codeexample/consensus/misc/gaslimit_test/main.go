package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/consensus/misc"
)

func main() {
	fmt.Println(misc.VerifyGaslimit(3, 2))
}
