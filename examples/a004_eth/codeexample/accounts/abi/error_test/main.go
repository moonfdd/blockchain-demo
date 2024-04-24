package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	err := abi.NewError("aa", []abi.Argument{})
	fmt.Println("err = ", err)
	fmt.Println(err.Unpack([]byte("哈哈hehe")))
}
