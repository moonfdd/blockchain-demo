package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	// ABI Method Event Error Argument
	// ArgumentMarshaling
	//abi文件本质是json，用json.Unmarshal也每问题
	// abi入参是元组
	if false {
		// Method with tuple arguments
		s, _ := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
			{Name: "a", Type: "int256"},
			{Name: "b", Type: "int256[]"},
			{Name: "c", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
				{Name: "x", Type: "int256"},
				{Name: "y", Type: "int256"},
			}},
			{Name: "d", Type: "tuple[2]", Components: []abi.ArgumentMarshaling{
				{Name: "x", Type: "int256"},
				{Name: "y", Type: "int256"},
			}},
		})
		String, _ := abi.NewType("string", "", nil)
		m := abi.NewMethod("foo", "foo", abi.Function, "", false, false, []abi.Argument{{"s", s, false}, {"bar", String, false}}, nil)
		exp := "foo((int256,int256[],(int256,int256)[],(int256,int256)[2]),string)"
		fmt.Println(m.Sig)
		fmt.Println(exp)
	}
	// 首字母大写。帕斯卡命名法
	if true {
		str := abi.ToCamelCase("foo_bar")
		fmt.Println(str)
		str = abi.ToCamelCase("fooBar")
		fmt.Println(str)
		str = abi.ToCamelCase("fo_Fu")
		fmt.Println(str)
	}
}
