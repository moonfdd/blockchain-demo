package main

import (
	"fmt"
	"strings"

	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	if false {
		fmt.Println(abi.MaxUint256)
		fmt.Println(abi.MaxInt256)
	}
	if false {
		//Unpack
		testdef := `[{ "type": "bool" }]`
		packed := "0000000000000000000000000000000000000000000000000000000000000001"
		def := fmt.Sprintf(`[{ "name" : "method", "type": "function", "outputs": %s}]`, testdef)
		abi, err := abi.JSON(strings.NewReader(def))
		if err != nil {
			fmt.Printf("invalid ABI definition %s: %v", def, err)
			return
		}
		encb, err := hex.DecodeString(packed)
		if err != nil {
			fmt.Printf("invalid hex %s: %v", packed, err)
			return
		}
		out, err := abi.Unpack("method", encb)
		if err != nil {
			fmt.Printf("test %d (%v) failed: %v", 0, testdef, err)
			return
		}
		_ = out
		fmt.Println(out)
	}
	fmt.Println("")
}
