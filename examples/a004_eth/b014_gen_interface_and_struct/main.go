package main

import (
	"encoding/json"
	"fmt"
	"io"

	ethereum_174b608b422ef534a95e8bc58ac60f15 "github.com/ethereum/go-ethereum"
	rpc_6be02959c527ca36913a567aaf6fe515 "github.com/ethereum/go-ethereum/rpc"
	"github.com/moonfdd/blockchain-demo/examples/a004_eth/b014_gen_interface_and_struct/build"
	"github.com/moonfdd/blockchain-demo/examples/a004_eth/b014_gen_interface_and_struct/mycore"
)

type B interface {
	A
	Test2()
}

type A interface {
	Test11()
	Test12()
	Test13()
}

func main() {
	if false {
		var a A = struct{ A }{}
		var b B = struct{ B }{}
		if _, ok := a.(B); ok {
			fmt.Println("ok")
		} else {
			fmt.Println("nook")
		}
		if _, ok := b.(A); ok {
			fmt.Println("ok")
		} else {
			fmt.Println("nook")
		}

		if false {

			type S struct {
				ethereum_174b608b422ef534a95e8bc58ac60f15.Subscription
			}

			var d ethereum_174b608b422ef534a95e8bc58ac60f15.Subscription = S{}

			if _, ok := d.(ethereum_174b608b422ef534a95e8bc58ac60f15.ChainReader); ok {
				fmt.Println("ok")
			} else {
				fmt.Println("nook")
			}
		}
		if true {

			type T1 = rpc_6be02959c527ca36913a567aaf6fe515.Error
			type S struct {
				T1
			}

			var c rpc_6be02959c527ca36913a567aaf6fe515.Error = &S{}

			if _, ok := c.(T1); ok {
				fmt.Println("ok")
			} else {
				fmt.Println("nook")
			}
		}

		return
	}
	// D:\mysetup\gopath\pkg\mod\github.com\ethereum\go-ethereum@v1.13.14
	// github.com/ethereum/go-ethereum

	// D:\mysetup\gopath\src\go-ethereum\common\lru\basiclru.go
	// D:\mysetup\gopath\src\go-ethereum\accounts\accounts.go
	if false {
		a, b := mycore.LoadDir(`D:\gvm\.g\go\src\`, "")
		count := 0
		for _, v := range a {
			count++
			fmt.Println(v.Name, v.Alias, v.Package)
			if count >= 2000 {
				break
			}
		}
		fmt.Println("")
		fmt.Println(len(a))
		fmt.Println(len(b))
		return
	}
	if false {
		a, b := mycore.LoadDir(`D:\mysetup\gopath\pkg\mod\github.com\ethereum\go-ethereum@v1.13.14`, "github.com/ethereum/go-ethereum")
		fmt.Println(len(a))
		fmt.Println(len(b))
		return
	}
	if false {
		mycore.LoadPath(`D:\mysetup\gopath\src\dsy\blockchain-demo\examples\a004_eth\b014_gen_interface_and_struct\main.go`)
		return
	}

	if false {
		fmt.Println("生成interface代码")
		mycore.BuildInterface(`D:\mysetup\gopath\src\dsy\blockchain-demo\examples\a004_eth\b014_gen_interface_and_struct\build\a.go`)
		return
	}
	if false {
		fmt.Println("生成struct代码")
		mycore.BuildStruct(`D:\mysetup\gopath\src\dsy\blockchain-demo\examples\a004_eth\b014_gen_interface_and_struct\build\b.go`)
		return
	}
	if false {
		// var b = struct {
		// 	rpc_6be02959c527ca36913a567aaf6fe515.Error
		// }{}
		// var c rpc_6be02959c527ca36913a567aaf6fe515.Error = b
		// _ = c
		// //  = struct {
		// // 	rpc_6be02959c527ca36913a567aaf6fe515.Error
		// // }{}
		// _ = b
		return
	}
	if false {

		fmt.Println("接口继承")
		build.Test()
		return
	}
	if true {
		fmt.Println("结构体继承")
		build.Test2()
	}
	return
	dir := `D:\mysetup\gopath\pkg\mod\github.com\ethereum\go-ethereum@v1.13.14`
	packageNamePrefix := `github.com/ethereum/go-ethereum`
	a, b := mycore.LoadDir(dir, packageNamePrefix)
	d, _ := json.MarshalIndent(a, "", "  ")
	_ = b
	fmt.Println(string(d))
	// d, _ = json.MarshalIndent(b, "", "  ")
	// fmt.Println(string(d))
	fmt.Println(len(a))
	fmt.Println(len(b))

	var r io.Reader
	var i interface{}
	r, _ = i.(io.Reader)
	i, _ = r.(interface{})
	return
	fmt.Println(mycore.LoadPathList([]string{
		// `D:\mysetup\gopath\src\go-ethereum\common\lru\basiclru.go`,
		`D:\mysetup\gopath\src\go-ethereum\accounts\accounts.go`,
		// `D:\mysetup\gopath\src\go-ethereum\consensus\beacon\consensus.go`,
		// `D:\mysetup\gopath\src\go-ethereum\consensus\merger.go`
	}))
	fmt.Println("")
}
