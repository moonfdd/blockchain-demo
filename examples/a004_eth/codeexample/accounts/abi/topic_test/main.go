package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	// 这段代码片段的注释描述了一个函数的功能。这个函数名为 `MakeTopics`，用于将一个过滤器（filter）查询参数列表转换为一个过滤器主题集合（filter topic set）。
	// 在区块链和智能合约开发中，过滤器通常用于订阅特定事件。过滤器可以包含各种查询参数来过滤所订阅的事件，而过滤器主题是用于标识事件的关键信息。因此，这个函数的作用是将传入的过滤器查询参数列表转换为对应的过滤器主题集合，以便进一步对事件进行过滤和订阅。
	if true {
		hashes, err := abi.MakeTopics([]interface{}{"Transfer", "from", "to"})
		if err != nil {
			fmt.Println("MakeTopics失败", err)
			return
		}
		fmt.Println("hashes:", hashes)
	}
}
