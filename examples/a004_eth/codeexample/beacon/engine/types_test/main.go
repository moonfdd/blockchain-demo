package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/beacon/engine"
)

func main() {
	if false {
		fmt.Println(engine.PayloadV1)
		fmt.Println(engine.PayloadV2)
		fmt.Println(engine.PayloadV3)

		fmt.Println(engine.ClientCode)
		fmt.Println(engine.ClientName)
	}
	if true {
		v := engine.ClientVersionV1{}
		v.Code = "1"
		v.Name = "2"
		v.Version = "3"
		v.Commit = "4"
		fmt.Println(v)
	}

}
