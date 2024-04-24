package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/ethdb/remotedb"
)

func main() {
	c, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	db := remotedb.New(c)
	defer db.Close()
	fmt.Println(db.Get([]byte("hello")))
	fmt.Println(db.Ancients())
}
