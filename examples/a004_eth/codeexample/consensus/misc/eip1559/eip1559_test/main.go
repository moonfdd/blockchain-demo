package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

func main() {
	if true {
		initial := new(big.Int).SetUint64(params.InitialBaseFee)
		parent := &types.Header{
			GasUsed:  10000000 / 2,
			GasLimit: 10000000,
			BaseFee:  initial,
			Number:   big.NewInt(4),
		}
		header := &types.Header{
			GasUsed:  20000000 / 2,
			GasLimit: 20000000,
			BaseFee:  initial,
			Number:   big.NewInt(4 + 1),
		}
		err := eip1559.VerifyEIP1559Header(config(), parent, header)
		fmt.Println(err)
	}
	if false {
		parent := &types.Header{
			Number:   common.Big32,
			GasLimit: 20000000,
			GasUsed:  10000000,
			BaseFee:  big.NewInt(params.InitialBaseFee),
		}
		r := eip1559.CalcBaseFee(config(), parent)
		fmt.Println(r)
	}
}
func copyConfig(original *params.ChainConfig) *params.ChainConfig {
	return &params.ChainConfig{
		ChainID:                 original.ChainID,
		HomesteadBlock:          original.HomesteadBlock,
		DAOForkBlock:            original.DAOForkBlock,
		DAOForkSupport:          original.DAOForkSupport,
		EIP150Block:             original.EIP150Block,
		EIP155Block:             original.EIP155Block,
		EIP158Block:             original.EIP158Block,
		ByzantiumBlock:          original.ByzantiumBlock,
		ConstantinopleBlock:     original.ConstantinopleBlock,
		PetersburgBlock:         original.PetersburgBlock,
		IstanbulBlock:           original.IstanbulBlock,
		MuirGlacierBlock:        original.MuirGlacierBlock,
		BerlinBlock:             original.BerlinBlock,
		LondonBlock:             original.LondonBlock,
		TerminalTotalDifficulty: original.TerminalTotalDifficulty,
		Ethash:                  original.Ethash,
		Clique:                  original.Clique,
	}
}
func config() *params.ChainConfig {
	config := copyConfig(params.TestChainConfig)
	config.LondonBlock = big.NewInt(5)
	return config
}
