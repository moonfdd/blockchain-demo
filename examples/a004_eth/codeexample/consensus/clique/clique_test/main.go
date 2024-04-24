package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

func main() {
	if false {
		var (
			db     = rawdb.NewMemoryDatabase()
			key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
			addr   = crypto.PubkeyToAddress(key.PublicKey)
			engine = clique.New(params.AllCliqueProtocolChanges.Clique, db)
			signer = new(types.HomesteadSigner)
		)
		fmt.Println(addr, engine, signer)
	}
	if true {
		hash := clique.SealHash(&types.Header{
			Difficulty: new(big.Int),
			Number:     new(big.Int),
			Extra:      make([]byte, 32+65),
			BaseFee:    new(big.Int),
		})
		fmt.Println(hash)
	}
}
