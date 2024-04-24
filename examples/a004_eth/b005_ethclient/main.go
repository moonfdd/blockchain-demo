package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	// It's negative.
	if number.IsInt64() {
		return rpc.BlockNumber(number.Int64()).String()
	}
	// It's negative and large, which is invalid.
	return fmt.Sprintf("<invalid %d>", number)
}

func main() {
	// create instance of ethclient and assign to cl
	cl, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer cl.Close()
	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		fmt.Println("获取ChainID失败:", err)
		return
	}
	fmt.Println("chainid:", chainid)
	if false {
		var res []common.Address
		if err := cl.Client().Call(&res, "account_list"); err != nil {
			fmt.Println("account_list失败:", err)
			return
		}
		fmt.Println("account_list:", res)
	}
	if false {
		var lastNumber hexutil.Uint64

		err = cl.Client().Call(&lastNumber, "eth_blockNumber")
		if err != nil {
			fmt.Println("eth_blockNumber失败:", err)
			return
		}
		fmt.Println(lastNumber)

		return
	}

	if false {
		info := &p2p.NodeInfo{
			// ID: n.ID.String(),
			//ID: `{"address":"7e5f4552091a69125d5dfcb7b8c2659029395bdf","crypto":{"cipher":"aes-128-ctr","ciphertext":"e468ad3ce2a6ae1436f8b50259f7cef84e688df2b08575f93eda7067492fec45","cipherparams":{"iv":"80697348d506166676cc27cc10aa00e0"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"806d6da06fbb4f2480e4c93f7bd82d5575f2da3ce382b469e6ed293e9e5cb52c"},"mac":"5fb6e482ee88c6a99972ace20827d547d2245249c2c3cfa8a8c8fea8bc82b771"},"id":"509c8fb9-b781-4c3b-bce7-e11b75e03ac3","version":3}`,
		}
		err = cl.Client().Call(&info, "admin_nodeInfo")
		if err != nil {
			fmt.Println("admin_nodeInfo失败:", err)
			return
		}
		d, _ := json.MarshalIndent(info, "", "  ")
		fmt.Println(string(d))

		return
	}

	addr := common.HexToAddress("0x7e5f4552091a69125d5dfcb7b8c2659029395bdf")
	nonce, err := cl.NonceAt(context.Background(), addr, big.NewInt(0))
	if err != nil {
		fmt.Println("NonceAt失败:", err)
		return
	}
	fmt.Println("nonce:", nonce)

	blockNumber, err := cl.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Failed to retrieve block number:", err)
		return
	}
	fmt.Println("blockNumber:", blockNumber)
	blockNumberBig := big.NewInt(int64(blockNumber))
	if false {
		eventSignatureBytes := []byte("Transfer(address,address,uint256)")
		eventSignaturehash := crypto.Keccak256Hash(eventSignatureBytes)

		q := ethereum.FilterQuery{
			FromBlock: new(big.Int).Sub(blockNumberBig, big.NewInt(10)),
			ToBlock:   blockNumberBig,
			Topics: [][]common.Hash{
				{eventSignaturehash},
			},
		}

		logs, err := cl.FilterLogs(context.Background(), q)
		if err != nil {
			fmt.Println("FilterLogs失败:", err)
			return
		}
		fmt.Println("FilterLogs成功:", logs)
	}

	//成功
	if true {
		// SK := "0x0000000000000000000000000000000000000000000000000000000000000001"
		ADDR := "0x7e5f4552091a69125d5dfcb7b8c2659029395bdf"
		nonce, err := cl.PendingNonceAt(context.Background(), common.HexToAddress(ADDR))
		if err != nil {
			fmt.Println("PendingNonceAt失败:", err)
			return
		}
		fmt.Println("PendingNonceAt成功:", nonce)

		value := big.NewInt(1000000000000000000) // 1 ETH in Wei
		gasLimit := uint64(21000)                // 默认gas limit为21,000
		gasPrice, err := cl.SuggestGasPrice(context.Background())
		if err != nil {
			fmt.Println("SuggestGasPrice失败:", err)
			return
		}

		toAddress := common.HexToAddress("0x2b5ad5c4795c026514f8317c7a215e218dccd6cf")

		var data []byte
		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

		privateKey, err := crypto.HexToECDSA("0000000000000000000000000000000000000000000000000000000000000001")
		if err != nil {
			fmt.Println("HexToECDSA失败:", err)
			return
		}
		// types.HomesteadSigner{}
		signedTx, err := types.SignTx(tx, types.NewCancunSigner(chainid), privateKey)
		if err != nil {
			fmt.Println("SignTx失败:", err)
			return
		}

		err = cl.SendTransaction(context.Background(), signedTx)
		if err != nil {
			fmt.Println("SendTransaction失败:", err)
			return
		}

		// 失败
		// _, err = bind.WaitDeployed(context.Background(), cl, signedTx)
		// if err != nil {
		// 	fmt.Printf("WaitDeployed: %v", err)
		// 	return
		// }

		fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	}
	//失败
	if false {
		// SK and ADDR are the secret key and sender address
		SK := "0x0000000000000000000000000000000000000000000000000000000000000001"
		ADDR := "0x7e5f4552091a69125d5dfcb7b8c2659029395bdf"
		var (
			sk       = crypto.ToECDSAUnsafe(common.FromHex(SK))
			to       = common.HexToAddress("0x2b5ad5c4795c026514f8317c7a215e218dccd6cf")
			value    = new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether))
			sender   = common.HexToAddress(ADDR)
			gasLimit = uint64(21000)
		)
		// Retrieve the chainid (needed for signer)
		// chainid, err := cl.ChainID(context.Background())
		// if err != nil {
		// 	return err
		// }
		// Retrieve the pending nonce
		nonce2, err := cl.PendingNonceAt(context.Background(), sender)
		if err != nil {
			fmt.Println("PendingNonceAt失败:", err)
			return
		}
		fmt.Println("nonce2:", nonce2)
		// Get suggested gas price
		tipCap, err := cl.SuggestGasTipCap(context.Background())
		if err != nil {
			fmt.Println("SuggestGasTipCap失败:", err)
			return
		}
		fmt.Println("tipCap:", tipCap)
		feeCap, err := cl.SuggestGasPrice(context.Background())
		if err != nil {
			fmt.Println("SuggestGasPrice失败:", err)
			return
		}
		fmt.Println("feeCap:", feeCap)
		// Create a new transaction
		tx := types.NewTx(
			&types.DynamicFeeTx{
				ChainID:   chainid,
				Nonce:     nonce2,
				GasTipCap: tipCap,
				GasFeeCap: feeCap,
				Gas:       gasLimit,
				To:        &to,
				Value:     value,
				Data:      nil,
			})
		// Sign the transaction using our keys
		// types.NewLondonSigner(chainid)

		signedTx, err := types.SignTx(tx, types.NewCancunSigner(chainid), sk)
		if err != nil {
			fmt.Println("SignTx失败:", err)
			return
		}
		// Send the transaction to our node
		err = cl.SendTransaction(context.Background(), signedTx)
		if err != nil {
			fmt.Println("SendTransaction失败:", err)
			return
		}
	}
}
