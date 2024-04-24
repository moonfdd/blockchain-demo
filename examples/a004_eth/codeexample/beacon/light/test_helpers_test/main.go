package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/beacon/light"
)

func main() {
	if false {
		fmt.Println(light.ErrNeedCommittee)
		fmt.Println(light.ErrInvalidUpdate)
		fmt.Println(light.ErrInvalidPeriod)
		fmt.Println(light.ErrWrongCommitteeRoot)
		fmt.Println(light.ErrCannotReorg)
	}
	if true {
		ttee := light.GenerateTestCommittee()
		fmt.Println(ttee)
	}

}
