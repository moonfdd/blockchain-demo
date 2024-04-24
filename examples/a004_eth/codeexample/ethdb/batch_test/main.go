package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
)

func main() {
	if false {
		db := memorydb.NewWithCap(1)
		defer db.Close()
		b := db.NewBatch()
		b.Put([]byte("hello"), []byte("world"))
		b.Write()
		db.Put([]byte("hello1"), []byte("world1"))
		db.Put([]byte("hello2"), []byte("world2"))
		fmt.Println(db.Get([]byte("hello")))
	}
	if true {
		db := memorydb.NewWithCap(1)
		defer db.Close()
		batch := ethdb.HookedBatch{
			Batch: db.NewBatch(),
			OnPut: func(key []byte, value []byte) {
				//s.accountBytes += common.StorageSize(len(key) + len(value))
			},
		}
		batch.Put([]byte("hello"), []byte("world"))
		batch.Write()
		db.Put([]byte("hello1"), []byte("world1"))
		db.Put([]byte("hello2"), []byte("world2"))
		fmt.Println(db.Get([]byte("hello")))
	}
}
