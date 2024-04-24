package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/beacon/types"
)

func main() {
	if true {
		f := types.Fork{}
		fmt.Println(f)
		fs := make([]*types.Fork, 0)
		fs = append(fs, &f)
		fs = append(fs, &types.Fork{})
		var forks types.Forks = fs
		fmt.Println(forks.Len())
		hash, err := forks.SigningRoot(types.Header{})
		if err != nil {
			panic(err)
		}
		fmt.Println(hash)

	}
	// 	这是一个Go语言中的结构体定义，表示了区块链的配置信息。其中包含了以下字段：

	// - GenesisTime：创世时间，以Unix时间戳形式表示，表示第一个槽位（slot 0）的时间。
	// - GenesisValidatorsRoot：创世验证人集合的根哈希，在签名域计算中使用。
	// - Forks：关于区块链分叉（fork）的相关配置信息。

	// 这个结构体用于存储和传递区块链配置的重要参数和数据。
	// 未完成

	if true {
		c := types.ChainConfig{}
		err := c.LoadForks("path")
		if err != nil {
			panic(err)
		}
		fmt.Println(c)
	}
}
