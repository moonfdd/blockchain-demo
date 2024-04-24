package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	// ABI Method Event Error Argument
	// ArgumentMarshaling
	//abi文件本质是json，用json.Unmarshal也每问题
	if true {
		j := `[{"inputs":[],"name":"retrieve","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"num","type":"uint256"}],"name":"store","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
		a, err := abi.JSON(strings.NewReader(j))
		if err != nil {
			fmt.Println("JSON失败", err)
		}
		d, _ := json.MarshalIndent(a, "", "  ")
		fmt.Println("JSON成功", string(d))
		json.Unmarshal([]byte(j), &a)
		d, _ = json.MarshalIndent(a, "", "  ")
		fmt.Println("JSON成功2", string(d))

		// 失败
		var packed []byte
		var i interface{}
		packed, err = a.Pack("method", &i)
		fmt.Println(packed, err)
	}

}
