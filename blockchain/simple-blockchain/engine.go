// my-go-examples simple-blockchain engine.go

package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// Create a blockchain
func createBlockchain() {

	t := time.Now()
	firstBlock := BlockStruct{}

	firstBlock = BlockStruct{
		Index:     0,
		Timestamp: t.String(),
		Data:      "hi jeff",
		Hash:      calculateBlockHash(firstBlock),
		PrevHash:  "",
	}

	fmt.Printf("\nCongrats, your first Block in your blockchain is:\n\n")
	spew.Dump(firstBlock)

	Blockchain = append(Blockchain, firstBlock)
}

// Get the Blockchain
func getBlockchain() Blockchain {

	return Blockchain

}

// getBlock
func getBlock(id string) BlockStruct {

	// SEARCH DATA FOR blockID
	for _, item := range Blockchain {
		if strconv.Itoa(item.Index) == id {
			// RETURN ITEM
			return item
		}
	}

	// RETURN NOT FOUND
	return nil
}

// addBlockToChain
func addBlockToChain(newData Message) BlockStruct {

	prevBlock := Blockchain[len(Blockchain)-1]
	newBlock := addBlockToChain(prevBlock, newData.Data)

	// CHECK IF NEWBLOCK IS VALID
	if isBlockValid(newBlock, prevBlock) {
		Blockchain = append(Blockchain, newBlock)
		spew.Dump(Blockchain)
	}

	return newBlock
}
