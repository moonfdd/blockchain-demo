package main

import (
	"crypto/rand"
	"fmt"

	"github.com/ethereum/go-ethereum/beacon/light"
	"github.com/ethereum/go-ethereum/beacon/types"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
)

func main() {
	if false {
		fmt.Println(light.ErrNeedCommittee)
		fmt.Println(light.ErrInvalidUpdate)
		fmt.Println(light.ErrInvalidPeriod)
		fmt.Println(light.ErrWrongCommitteeRoot)
		fmt.Println(light.ErrCannotReorg)
	}
	// 这段注释描述了一个被动数据结构，名为CommitteeChain（委员会链），它能够验证、保存和更新一系列的信标委员会和更新。它需要至少一个在链开始时外部设置的固定委员会根，可以基于BootstrapData或信任的来源（本地信标全节点）进行设置。这使得该结构对轻客户端和轻服务器都非常有用。

	// 它始终保持以下一致性约束：

	// 如果一个委员会的根哈希与现有的固定根匹配，或者它通过上个周期的更新被证明存在，则该委员会才能存在。
	// 如果在同一周期存在一个委员会，并且更新签名有效且参与者足够，则该更新才能存在。下一个周期中的委员会（由更新证明）也应当存在。（注意，这意味着如果两者都尚未存在，则它们只能同时添加）。如果下一个周期存在固定根，则更新只能在证明相同委员会根的情况下才能存在。
	// 一旦同步到当前同步周期，CommitteeChain还可以验证签名的信标头。
	if true {
		db := memorydb.New()
		testGenesis := newTestGenesis()
		tfAlternative := newTestForks(testGenesis, types.Forks{
			&types.Fork{Epoch: 0, Version: []byte{0}},
			&types.Fork{Epoch: 0x700, Version: []byte{1}},
		})
		chain := light.NewCommitteeChain(db, &tfAlternative, 300, true)
		fmt.Println(chain)
	}
}

func newTestForks(config types.ChainConfig, forks types.Forks) types.ChainConfig {
	for _, fork := range forks {
		config.AddFork(fork.Name, fork.Epoch, fork.Version)
	}
	return config
}

func newTestGenesis() types.ChainConfig {
	var config types.ChainConfig
	rand.Read(config.GenesisValidatorsRoot[:])
	return config
}
