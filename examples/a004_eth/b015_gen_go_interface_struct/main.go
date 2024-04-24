package main

import (
	"fmt"

	"github.com/moonfdd/blockchain-demo/examples/a004_eth/b015_gen_go_interface_struct/mycore2"
)

func main() {

	if false {
		fmt.Println("生成interface代码")
		mycore2.BuildInterface(`D:\mysetup\gopath\src\dsy\blockchain-demo\examples\a004_eth\b015_gen_go_interface_struct\buildinterface\a.go`)
		return
	}
	if true {
		fmt.Println("生成struct代码")
		mycore2.BuildStruct(`D:\mysetup\gopath\src\dsy\blockchain-demo\examples\a004_eth\b015_gen_go_interface_struct\buildstruct\b.go`)
		return
	}

	// if false {

	// 	fmt.Println("接口继承")
	// 	build.Test()
	// 	return
	// }
	// if true {
	// 	fmt.Println("结构体继承")
	// 	build.Test2()
	// }
	return

}
