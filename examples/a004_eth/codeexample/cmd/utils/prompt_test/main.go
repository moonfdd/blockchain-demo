package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/cmd/utils"
)

func main() {
	if false {
		fmt.Println(utils.GetPassPhrase("123", false))
	}
	if true {
		fmt.Println(utils.GetPassPhraseWithList("123", false, 1, []string{"123", "456"}))
	}
}
