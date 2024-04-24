package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethdb/memorydb"
)

func main() {
	// db := memorydb.New()
	db := memorydb.NewWithCap(1)
	defer db.Close()
	db.Put([]byte("hello1"), []byte("world1"))
	db.Put([]byte("hello2"), []byte("world2"))
	fmt.Println(db.Get([]byte("hello2")))
}
