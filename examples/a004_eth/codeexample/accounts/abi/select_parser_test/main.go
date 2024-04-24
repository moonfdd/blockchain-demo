package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	// 这段代码注释描述了一个函数 `ParseSelector` 的作用。这个函数用于将一个方法选择器（method selector）转换为一个结构体，以便可以通过 JSON 编码并被这个包中的其他函数使用。
	// 在 Ethereum 中，Solidity 编译后的函数调用通常会包含一个方法选择器，用来确定要调用的具体函数。这个函数的目的是将这个选择器解析为一个结构体的格式，这样可以方便地通过 JSON 编码，同时提供给这个包中的其他函数使用。
	// 需要注意的是，尽管 ABI 规范中并不包含大写字母，但是这个函数依然接受这种格式，因为一般性的格式是有效的。这意味着，即使 ABI 规范不强制要求，这个函数依然可以处理包含大写字母的方法选择器。
	if true {
		selector, err := abi.ParseSelector("noargs()")
		if err != nil {
			fmt.Println("ParseSelector失败", err)
		}
		fmt.Println("ParseSelector成功", selector)
	}
	if true {
		selector, err := abi.ParseSelector("simple(uint256,uint256,uint256)")
		if err != nil {
			fmt.Println("ParseSelector失败", err)
		}
		fmt.Println("ParseSelector成功", selector)
	}
}
