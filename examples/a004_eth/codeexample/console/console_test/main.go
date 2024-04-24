package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/console"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	if false {
		fmt.Println(console.HistoryFile)
		fmt.Println(console.DefaultPrompt)
	}
	if true {
		conn, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}
		fmt.Println("Dial成功")
		printer := new(bytes.Buffer)
		c, err := console.New(console.Config{
			DataDir:  ".",
			DocRoot:  "testdata",
			Client:   conn.Client(),
			Prompter: nil,
			Printer:  printer,
			// Preload:  []string{"preload.js"},
			Preload: nil,
		})
		if err != nil {
			fmt.Println("console.New失败", err)
			return
		}
		c.Welcome()
		c.Evaluate("1+1")
		c.Evaluate("eth.accounts")

		fmt.Println(string(printer.Bytes()))

	}
}
