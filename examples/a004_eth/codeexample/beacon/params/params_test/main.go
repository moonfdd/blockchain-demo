package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/beacon/params"
)

func main() {
	fmt.Println(params.EpochLength)
	fmt.Println(params.SyncPeriodLength)

	fmt.Println(params.BLSSignatureSize)
	fmt.Println(params.BLSPubkeySize)

	fmt.Println(params.SyncCommitteeSize)
	fmt.Println(params.SyncCommitteeBitmaskSize)
	fmt.Println(params.SyncCommitteeSupermajority)

	fmt.Println(params.StateIndexGenesisTime)
	fmt.Println(params.StateIndexGenesisValidators)
	fmt.Println(params.StateIndexForkVersion)
	fmt.Println(params.StateIndexLatestHeader)
	fmt.Println(params.StateIndexBlockRoots)
	fmt.Println(params.StateIndexStateRoots)
	fmt.Println(params.StateIndexHistoricRoots)
	fmt.Println(params.StateIndexFinalBlock)
	fmt.Println(params.StateIndexSyncCommittee)
	fmt.Println(params.StateIndexNextSyncCommittee)
	fmt.Println(params.StateIndexExecPayload)
	fmt.Println(params.StateIndexExecHead)
}
