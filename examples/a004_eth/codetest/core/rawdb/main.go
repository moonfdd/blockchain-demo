package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

func main() {
	if true {
		db := memorydb.New()
		db.Put([]byte("aa"), []byte("aa"))
		db.Put([]byte("bb"), []byte("bb"))
		db.Put([]byte("cc"), []byte("cc"))
		fmt.Println(db)
	}
	if false {
		db, err := leveldb.Open(storage.NewMemStorage(), nil)
		if err != nil {
			fmt.Println("err1 = ", err)
			return
		}
		err = db.Put([]byte("aaa"), []byte("aaa"), nil)
		db.Put([]byte("bbb"), []byte("ccc"), nil)
		if err != nil {
			fmt.Println("err2 = ", err)
			return
		}
		v, _ := db.Get([]byte("aaa"), nil)
		fmt.Println("string(v) = ", string(v))
		fmt.Println(db)
	}
	if true {

	}
	fmt.Println("aa")
}
