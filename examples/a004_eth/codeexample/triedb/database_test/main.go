package main

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/triedb"
)

func main() {
	if false {
		fmt.Println(triedb.HashDefaults)
	}
	// 未完成
	if false {
		db := triedb.NewDatabase(nil, triedb.HashDefaults)
		fmt.Println(db)
	}
	if true {
		var a json.Number = "123"
		fmt.Println(a.Int64())
	}
}
