package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/ethereum/go-ethereum/signer/fourbyte"
)

func main() {
	if false {
		hexdata := "a9059cbb000000000000000000000000ea0e2dc7d65a50e77fc7e84bff3fd2a9e781ff5c0000000000000000000000000000000000000000000000015af1d78b58c40000"
		data, err := hex.DecodeString(strings.TrimPrefix(hexdata, "0x"))
		if err != nil {
			panic(err)
		}
		db, err := fourbyte.New()
		if err != nil {
			panic(err)
		}
		messages := apitypes.ValidationMessages{}
		db.ValidateCallData(nil, data, &messages)
		for _, m := range messages.Messages {
			fmt.Printf("%v---- %v\n", m.Typ, m.Message)
		}
	}
	if true {
		hexdata := "0x608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100a1565b60405180910390f35b610073600480360381019061006e91906100ed565b61007e565b005b60008054905090565b8060008190555050565b6000819050919050565b61009b81610088565b82525050565b60006020820190506100b66000830184610092565b92915050565b600080fd5b6100ca81610088565b81146100d557600080fd5b50565b6000813590506100e7816100c1565b92915050565b600060208284031215610103576101026100bc565b5b6000610111848285016100d8565b9150509291505056fea2646970667358221220d930a0e982648452bbdb8c6bcff7d8f073432d0f1b9155afb04547759545c4dc64736f6c63430008130033"
		data, err := hex.DecodeString(strings.TrimPrefix(hexdata, "0x"))
		if err != nil {
			panic(err)
		}
		db, err := fourbyte.New()
		if err != nil {
			panic(err)
		}
		messages := apitypes.ValidationMessages{}
		db.ValidateCallData(nil, data, &messages)
		for _, m := range messages.Messages {
			fmt.Printf("%v---- %v\n", m.Typ, m.Message)
		}
	}
}
