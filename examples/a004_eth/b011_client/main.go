package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics/librato"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/dnsdisc"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/rpc"
)

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
	if true {
		// var res []common.Address
		// if err := cl.Client().CallContext(context.Background(), &res, "account_list"); err != nil {
		// 	fmt.Println("account_list失败:", err)
		// 	return
		// }
		// fmt.Println("account_list:", res)
	}
	if false {
		var v string
		if err := cl.Client().Call(&v, "admin_datadir"); err != nil {
			fmt.Println("admin_datadir失败:", err)
			return
		}
		fmt.Println("admin_datadir:", v)
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
	if false {
		cl2 := ethclient.NewClient(cl.Client())
		blockNumber, err := cl2.BlockNumber(context.Background())
		if err != nil {
			fmt.Println("Failed to retrieve block number:", err)
			return
		}
		fmt.Println("blockNumber:", blockNumber)
	}
	if false {
		cl2 := gethclient.New(cl.Client())
		info, err := cl2.GetNodeInfo(context.Background())
		if err != nil {
			fmt.Println("GetNodeInfo:", err)
			return
		}
		d, _ := json.MarshalIndent(info, "", "  ")
		fmt.Println(string(d))

		type Client222 struct {
			C *rpc.Client
		}

		cl3 := (*Client222)(unsafe.Pointer(cl2)).C
		_ = cl3

		info2 := &p2p.NodeInfo{
			// ID: n.ID.String(),
			//ID: `{"address":"7e5f4552091a69125d5dfcb7b8c2659029395bdf","crypto":{"cipher":"aes-128-ctr","ciphertext":"e468ad3ce2a6ae1436f8b50259f7cef84e688df2b08575f93eda7067492fec45","cipherparams":{"iv":"80697348d506166676cc27cc10aa00e0"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"806d6da06fbb4f2480e4c93f7bd82d5575f2da3ce382b469e6ed293e9e5cb52c"},"mac":"5fb6e482ee88c6a99972ace20827d547d2245249c2c3cfa8a8c8fea8bc82b771"},"id":"509c8fb9-b781-4c3b-bce7-e11b75e03ac3","version":3}`,
		}
		err = cl3.Call(&info2, "admin_nodeInfo")
		if err != nil {
			fmt.Println("cl3 admin_nodeInfo失败:", err)
			return
		}
		d2, _ := json.MarshalIndent(info2, "", "  ")
		fmt.Println("cl3 admin_nodeInfo成功:", string(d2))

	}

	if false {
		var res []common.Address
		if err := cl.Client().Call(&res, "account_list"); err != nil {
			fmt.Println("account_list err:", err)
			return
		}
		fmt.Println("account_list:", res)
	}

	if false {
		var lastNumber hexutil.Uint64
		if err := cl.Client().Call(&lastNumber, "eth_blockNumber"); err != nil {
			fmt.Println("eth_blockNumber err:", err)
			return
		}
		fmt.Println("eth_blockNumber:", lastNumber)
	}

	// 成功
	if true {
		lastNumber := false
		if err := cl.Client().Call(&lastNumber, "net_listening"); err != nil {
			fmt.Println("net_listening err:", err)
			return
		}
		fmt.Println("net_listening:", lastNumber)
	}

	if false {
		lastNumber := make([]string, 0)
		if err := cl.Client().Call(&lastNumber, "eth_accounts"); err != nil {
			fmt.Println("eth_accounts err:", err)
			return
		}
		fmt.Println("eth_accounts:", lastNumber)
	}

	if false {
		nodes := []string{
			"enr:-HW4QOFzoVLaFJnNhbgMoDXPnOvcdVuj7pDpqRvh6BRDO68aVi5ZcjB3vzQRZH2IcLBGHzo8uUN3snqmgTiE56CH3AMBgmlkgnY0iXNlY3AyNTZrMaECC2_24YYkYHEgdzxlSNKQEnHhuNAbNlMlWJxrJxbAFvA",
			"enr:-HW4QAggRauloj2SDLtIHN1XBkvhFZ1vtf1raYQp9TBW2RD5EEawDzbtSmlXUfnaHcvwOizhVYLtr7e6vw7NAf6mTuoCgmlkgnY0iXNlY3AyNTZrMaECjrXI8TLNXU0f8cthpAMxEshUyQlK-AM0PW2wfrnacNI",
			"enr:-HW4QLAYqmrwllBEnzWWs7I5Ev2IAs7x_dZlbYdRdMUx5EyKHDXp7AV5CkuPGUPdvbv1_Ms1CPfhcGCvSElSosZmyoqAgmlkgnY0iXNlY3AyNTZrMaECriawHKWdDRk2xeZkrOXBQ0dfMFLHY4eENZwdufn1S1o",
		}

		r := mapResolver{
			"n":                            "enrtree-root:v1 e=JWXYDBPXYWG6FX3GMDIBFA6CJ4 l=C7HRFPF3BLGF3YR4DY5KX3SMBE seq=1 sig=o908WmNp7LibOfPsr4btQwatZJ5URBr2ZAuxvK4UWHlsB9sUOTJQaGAlLPVAhM__XJesCHxLISo94z5Z2a463gA",
			"C7HRFPF3BLGF3YR4DY5KX3SMBE.n": "enrtree://AM5FCQLWIZX2QFPNJAP7VUERCCRNGRHWZG3YYHIUV7BVDQ5FDPRT2@morenodes.example.org",
			"JWXYDBPXYWG6FX3GMDIBFA6CJ4.n": "enrtree-branch:2XS2367YHAXJFGLZHVAWLQD4ZY,H4FHT4B454P6UXFD7JCYQ5PWDY,MHTDO6TMUBRIA2XWG5LUDACK24",
			"2XS2367YHAXJFGLZHVAWLQD4ZY.n": nodes[0],
			"H4FHT4B454P6UXFD7JCYQ5PWDY.n": nodes[1],
			"MHTDO6TMUBRIA2XWG5LUDACK24.n": nodes[2],
		}
		w := withDefaults()
		w.Resolver = r
		c := dnsdisc.NewClient(withDefaults())
		_ = c
		stree, err := c.SyncTree("enrtree://AKPYQIUQIL7PSIACI32J7FGZW56E5FKHEFCCOFHILBIMW3M6LWXS2@n")
		// stree, err := c.SyncTree("https://www.baidu.com")
		if err != nil {
			fmt.Println("SyncTree失败:", err)
			return
		}
		_ = stree

		d2, _ := json.MarshalIndent(stree, "", "  ")
		fmt.Println("SyncTree成功:", string(d2))
	}
	if false {
		c := &librato.LibratoClient{"email", "token"}
		err := c.PostMetrics(librato.Batch{})
		if err != nil {
			fmt.Println("PostMetrics失败:", err)
			return
		}
		fmt.Println("PostMetrics成功")
	}

}
func withDefaults() dnsdisc.Config {
	var cfg dnsdisc.Config
	const (
		defaultTimeout   = 5 * time.Second
		defaultRecheck   = 30 * time.Minute
		defaultRateLimit = 3
		defaultCache     = 1000
	)
	if cfg.Timeout == 0 {
		cfg.Timeout = defaultTimeout
	}
	if cfg.RecheckInterval == 0 {
		cfg.RecheckInterval = defaultRecheck
	}
	if cfg.CacheLimit == 0 {
		cfg.CacheLimit = defaultCache
	}
	if cfg.RateLimit == 0 {
		cfg.RateLimit = defaultRateLimit
	}
	if cfg.ValidSchemes == nil {
		cfg.ValidSchemes = enode.ValidSchemes
	}
	if cfg.Resolver == nil {
		cfg.Resolver = new(net.Resolver)
	}
	if cfg.Logger == nil {
		cfg.Logger = log.Root()
	}
	return cfg
}

type mapResolver map[string]string

func newMapResolver(maps ...map[string]string) mapResolver {
	mr := make(mapResolver, len(maps))
	for _, m := range maps {
		mr.add(m)
	}
	return mr
}

func (mr mapResolver) clear() {
	for k := range mr {
		delete(mr, k)
	}
}

func (mr mapResolver) add(m map[string]string) {
	for k, v := range m {
		mr[k] = v
	}
}

func (mr mapResolver) LookupTXT(ctx context.Context, name string) ([]string, error) {
	if record, ok := mr[name]; ok {
		return []string{record}, nil
	}
	return nil, errors.New("not found")
}
