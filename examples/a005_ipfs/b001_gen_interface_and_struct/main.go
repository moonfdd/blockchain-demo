package main

import (
	"fmt"

	"github.com/moonfdd/blockchain-demo/examples/a005_ipfs/b001_gen_interface_and_struct/mycore2"
)

func main() {
	if false {
		fmt.Println("生成interface代码")
		mycore2.BuildInterface(`D:\mysetup\gopath\src\dsy\blockchain-demo\examples\a005_ipfs\b001_gen_interface_and_struct\buildinterface\a.go`)
		return
	}
	if true {
		fmt.Println("生成struct代码")
		mycore2.BuildStruct(`D:\mysetup\gopath\src\dsy\blockchain-demo\examples\a005_ipfs\b001_gen_interface_and_struct\buildstruct\b.go`)
		return
	}
}
