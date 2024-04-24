package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/cmd/utils"
)

func main() {
	if false {
		fmt.Println(utils.DataDirFlag)
		fmt.Println(utils.RemoteDBFlag)
		// 未完成
		utils.CheckExclusive(nil, nil, nil)
		utils.SetEthConfig(nil, nil, nil)
	}
	if false {
		utils.CheckExclusive(nil, nil, nil)
	}
	if false {
		utils.SetDataDir(nil, nil)
	}
	if true {
		fmt.Println(utils.SplitTagsFlag("host=localhost,bzzkey=123"))
	}
}
