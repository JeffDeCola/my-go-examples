// my-go-examples simple-blockchain engine.go

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var mutex = &sync.Mutex{}

// Create a blockchain
func createBlockchain() {

	t := time.Now()
	firstBlock := Block{}

	firstBlock = Block{
		Index:     0,
		Timestamp: t.String(),
		Data:      "hi jeff",
		Hash:      calculateBlockHash(firstBlock),
		PrevHash:  "",
	}

	fmt.Printf("\nCongrats, your first Block in your blockchain is:\n\n")
	spew.Dump(firstBlock)

	mutex.Lock()
	Blockchain = append(Blockchain, firstBlock)
	mutex.Unlock()
}
