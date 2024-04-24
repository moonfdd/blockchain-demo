package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func main() {
	// 输出错误类型
	// 以太坊智能合约开发中可能出现的一些情况
	if true {
		// 当调用或执行操作的目标合约不存在或没有与之关联的代码时，会返回此错误。
		fmt.Println(bind.ErrNoCode)
		// 当尝试在不支持待定状态调用的后端上执行待定状态操作时，会引发此错误。
		fmt.Println(bind.ErrNoPendingState)
		// 当尝试在不支持区块哈希状态调用的后端上执行区块哈希操作时，会引发此错误。
		fmt.Println(bind.ErrNoBlockHashState)
		// 如果合约创建后没有留下任何代码，则在等待部署过程中返回此错误。
		fmt.Println(bind.ErrNoCodeAfterDeploy)
	}
	fmt.Println("")
}
