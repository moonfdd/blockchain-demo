package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/beacon/engine"
)

func main() {
	if true {
		fmt.Println(engine.VALID)
		fmt.Println(engine.INVALID)
		fmt.Println(engine.SYNCING)
		fmt.Println(engine.ACCEPTED)

		fmt.Println(engine.GenericServerError)
		fmt.Println(engine.UnknownPayload)
		fmt.Println(engine.InvalidForkChoiceState)
		fmt.Println(engine.InvalidPayloadAttributes)
		fmt.Println(engine.TooLargeRequest)
		fmt.Println(engine.InvalidParams)
		fmt.Println(engine.UnsupportedFork)

		fmt.Println(engine.STATUS_INVALID)
		fmt.Println(engine.STATUS_SYNCING)
		fmt.Println(engine.INVALID_TERMINAL_BLOCK)
	}

}
