package main

import (
	"fmt"

	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
	pebble2 "github.com/ethereum/go-ethereum/ethdb/pebble"
)

func main() {
	if false {
		db, err := pebble.Open("", &pebble.Options{
			FS: vfs.NewMem(),
		})
		if err != nil {
			panic(err)
		}
		defer db.Close()
		db.Set([]byte("hello"), []byte("world"), nil)
		fmt.Println(db.Get([]byte("hello")))
	}
	if true {
		db, err := pebble2.New("a.pebble2", 0, 0, "", false, false)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		db.Put([]byte("hello"), []byte("world"))
		fmt.Println(db.Get([]byte("hello")))
	}
}
