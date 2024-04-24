package main

import (
	"fmt"

	leveldb2 "github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

func main() {
	if false {
		db, err := leveldb.Open(storage.NewMemStorage(), nil)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		db.Put([]byte("hello"), []byte("world"), nil)
		fmt.Println(db.Get([]byte("hello"), nil))
	}
	if false {
		db, err := leveldb2.New("a.leveldb", 0, 0, "", false)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		db.Put([]byte("hello"), []byte("world"))
		fmt.Println(db.Get([]byte("hello")))
	}
	if true {
		db, err := leveldb2.NewCustom("a2.leveldb", "", nil)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		db.Put([]byte("hello"), []byte("world"))
		fmt.Println(db.Get([]byte("hello")))
	}
}
