package main

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/beacon/engine"
)

func main() {
	a := &engine.ExecutableData{}
	a.Timestamp = 22
	d, err := a.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(d))

	d, err = json.MarshalIndent(a, "", " ")
	fmt.Println(string(d))
}
