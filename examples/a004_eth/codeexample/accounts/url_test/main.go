package main

import (
	"fmt"

	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts"
)

func main() {
	url := accounts.URL{Scheme: "https", Path: "ethereum.org"}
	fmt.Println(url)
	d, _ := url.MarshalJSON()
	fmt.Println(string(d))

	d, _ = json.MarshalIndent(url, "", "  ")
	fmt.Println(string(d))
}
