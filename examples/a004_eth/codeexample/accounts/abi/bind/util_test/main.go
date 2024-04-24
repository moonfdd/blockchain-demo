package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
)

var testKey, _ = crypto.HexToECDSA("0000000000000000000000000000000000000000000000000000000000000001")

func main() {
	// WaitDeployed等待一个合约部署的交易，并在交易被成功上链后返回合约的地址。当传入的ctx（上下文）被取消时，该函数会停止等待。
	// WaitMined等待指定的交易在区块链上被打包并确认
	if true {
		backend := simulated.NewBackend(
			types.GenesisAlloc{
				crypto.PubkeyToAddress(testKey.PublicKey): {Balance: big.NewInt(10000000000000000)},
			},
		)
		defer backend.Close()
		// Create the transaction
		head, err := backend.Client().HeaderByNumber(context.Background(), nil) // Should be child's, good enough
		if err != nil {
			fmt.Println("HeaderByNumber失败 = ", err)
			return
		}
		gasPrice := new(big.Int).Add(head.BaseFee, big.NewInt(params.GWei))

		tx := types.NewContractCreation(0, big.NewInt(0), 3000000, gasPrice, common.FromHex(`6060604052600a8060106000396000f360606040526008565b00`))
		tx, err = types.SignTx(tx, types.LatestSignerForChainID(big.NewInt(1337)), testKey)
		if err != nil {
			fmt.Println("SignTx失败 = ", err)
			return
		}
		// Wait for it to get mined in the background.
		var (
			address common.Address
			mined   = make(chan struct{})
			ctx     = context.Background()
		)
		go func() {
			address, err = bind.WaitDeployed(ctx, backend.Client(), tx)
			if err != nil {
				fmt.Println("WaitDeployed失败 = ", err)
				return
			}
			close(mined)
		}()

		// Send and mine the transaction.
		backend.Client().SendTransaction(ctx, tx)
		backend.Commit()
		_ = err
		_ = address
		return
	}
	// 未完成
	if false {
		backend := simulated.NewBackend(
			types.GenesisAlloc{
				crypto.PubkeyToAddress(testKey.PublicKey): {Balance: big.NewInt(10000000000000000)},
			},
		)
		defer backend.Close()
		fmt.Println(backend.Client().BlockNumber(context.Background()))
		head, err := backend.Client().HeaderByNumber(context.Background(), nil) // Should be child's, good enough
		if err != nil {
			fmt.Println("HeaderByNumber = ", err)
			return
		}
		// fmt.Println(head)

		gasPrice := new(big.Int).Add(head.BaseFee, big.NewInt(1))

		// Create a transaction to an account.
		code := "6060604052600a8060106000396000f360606040526008565b00"
		tx := types.NewTransaction(0, common.HexToAddress("0x01"), big.NewInt(0), 3000000, gasPrice, common.FromHex(code))
		tx, _ = types.SignTx(tx, types.LatestSigner(params.AllDevChainProtocolChanges), testKey)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		backend.Client().SendTransaction(ctx, tx)
		backend.Commit()
		fmt.Println(backend.Client().BlockNumber(ctx))
		notContractCreation := errors.New("tx is not contract creation")
		if _, err := bind.WaitDeployed(ctx, backend.Client(), tx); err.Error() != notContractCreation.Error() {
			fmt.Printf("error mismatch: want %q, got %q, \r\n", notContractCreation, err)
		}

		// Create a transaction that is not mined.
		tx = types.NewContractCreation(1, big.NewInt(0), 3000000, gasPrice, common.FromHex(code))
		tx, _ = types.SignTx(tx, types.LatestSigner(params.AllDevChainProtocolChanges), testKey)

		go func() {
			contextCanceled := errors.New("context canceled")
			if _, err := bind.WaitDeployed(ctx, backend.Client(), tx); err.Error() != contextCanceled.Error() {
				fmt.Printf("error mismatch: want %q, got %q, \r\n", contextCanceled, err)
			}
		}()

		backend.Client().SendTransaction(ctx, tx)
		cancel()
	}
}
