package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/beacon/types"
)

func main() {
	if false {
		h := types.Header{}
		fmt.Println(h.Hash())
		fmt.Println(h.SyncPeriod())
	}
	if false {
		fmt.Println(types.SyncPeriodStart(99))
		fmt.Println(types.SyncPeriod(99))
	}
	// 这是一个Go语言中的结构体定义，表示了由同步委员会签名的信标链（beacon chain）头部。

	// 这个结构体包含以下字段：

	// - Header：被签名的信标链头部。它是一个类型为Header的结构体，包含了信标链头部的各种信息。
	// - Signature：同步委员会BLS签名的聚合值。它是一个类型为SyncAggregate的结构体，用于存储同步委员会对该头部进行签名的BLS聚合签名。
	// - SignatureSlot：创建该签名的槽位号（Slot）。这个槽位号通常比Header中记录的槽位号要新，它决定了进行签名操作的同步委员会。

	// 这个结构体用于存储和传递由同步委员会签署的信标链头部及其相关信息。

	if true {
		h := types.SignedHeader{}
		fmt.Println(h)
	}
}
