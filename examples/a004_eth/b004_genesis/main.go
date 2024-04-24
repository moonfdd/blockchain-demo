package main

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/core"
)

func main() {
	if true {
		g := core.DefaultGenesisBlock()
		data, err := json.MarshalIndent(g, "", "  ")
		if err != nil {
			fmt.Println("序列化失败", err)
			return
		}
		fmt.Println(string(data))
	}
	if false {
		str := `{
		"config": {
		  "chainId": 666123,
		  "homesteadBlock": 0,
		  "eip150Block": 0,
		  "eip155Block": 0,
		  "eip158Block": 0,
		  "byzantiumBlock": 0,
		  "constantinopleBlock": 0,
		  "petersburgBlock": 0,
		  "istanbulBlock": 0,
		  "berlinBlock": 0,
		  "londonBlock": 0
		},
		"alloc": {},
		"coinbase": "0x0000000000000000000000000000000000000000",
		"difficulty": "0x20000",
		"extraData": "",
		"gasLimit": "0x2fefd8",
		"nonce": "0x0000000000000042",
		"mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
		"parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
		"timestamp": "0x00"
	  }`
		// fmt.Println(str)
		g := core.Genesis{}
		err := json.Unmarshal([]byte(str), &g)
		if err != nil {
			fmt.Println("反序列化失败", err)
			return
		}
		data, err := json.MarshalIndent(g, "", "  ")
		if err != nil {
			fmt.Println("序列化失败", err)
			return
		}
		fmt.Println(string(data))
		fmt.Println("----")
		err = g.UnmarshalJSON([]byte(str))
		if err != nil {
			fmt.Println("反序列化失败2", err)
			return
		}
		data, err = g.MarshalJSON()
		if err != nil {
			fmt.Println("序列化失败2", err)
			return
		}
		fmt.Println(string(data))
	}
	fmt.Println("Hello, World!")
}
