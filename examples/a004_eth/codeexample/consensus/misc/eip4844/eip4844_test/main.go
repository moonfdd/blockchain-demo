package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

func main() {
	if true {
		initial := new(big.Int).SetUint64(params.InitialBaseFee)
		parent := &types.Header{
			GasUsed:  10000000 / 2,
			GasLimit: 10000000,
			BaseFee:  initial,
			Number:   big.NewInt(4),
		}
		header := &types.Header{
			GasUsed:  20000000 / 2,
			GasLimit: 20000000,
			BaseFee:  initial,
			Number:   big.NewInt(4 + 1),
		}
		err := eip4844.VerifyEIP4844Header(parent, header)
		fmt.Println(err)
	}
	if false {
		fmt.Println(eip4844.CalcExcessBlobGas(5500, 333))
	}
	if false {
		fmt.Println(eip4844.CalcBlobFee(5500))
	}
}
