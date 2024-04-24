package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/beacon/types"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	if false {
		fmt.Println(types.SerializedSyncCommitteeSize)
	}
	if false {
		var tt types.SerializedSyncCommittee
		d, err := tt.MarshalJSON()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(d))
		fmt.Println(tt.Root())
	}
	// 未完成
	if false {
		var tt types.SyncCommittee
		b := tt.VerifySignature(common.Hash{}, nil)
		fmt.Println(b)
	}
	// go:generate go run github.com/fjl/gencodec -type SyncAggregate -field-override syncAggregateMarshaling -out gen_syncaggregate_json.go
	if true {
		var tt types.SyncAggregate
		fmt.Println(tt.SignerCount())
	}
}
