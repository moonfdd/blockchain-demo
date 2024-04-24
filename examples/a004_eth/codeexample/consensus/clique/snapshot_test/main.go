package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/consensus/clique"
)

func main() {
	if false {
		vote := clique.Vote{}
		fmt.Println(vote)
	}
	if true {
		vote := clique.Tally{}
		fmt.Println(vote)
	}
}
