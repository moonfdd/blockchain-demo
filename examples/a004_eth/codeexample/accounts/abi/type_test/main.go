package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	if true {
		typ, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{{Name: "a", Type: "int64"}})
		if err != nil {
			fmt.Println("NewType失败", err)
			return
		}
		fmt.Println("typ:", typ)
		fmt.Println(typ.GetType())
	}
}
