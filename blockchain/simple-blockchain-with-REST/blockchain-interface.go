// my-go-examples simple-blockchain-with-REST blockchain-interface.go

package main

import (
	"fmt"
	"strconv"
	"time"
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
	fmt.Printf("%v\n", firstBlock)

	Blockchain = append(Blockchain, firstBlock)
}

// Get the Blockchain
func getBlockchain() BlockchainSlice {

	return Blockchain

}

// Get a Block from the chain
func getBlock(id string) BlockStruct {

	var item BlockStruct

	// SEARCH DATA FOR blockID
	for _, item := range Blockchain {
		if strconv.Itoa(item.Index) == id {
			// RETURN ITEM
			return item
		}
	}

	// RETURN NOT FOUND
	return item
}

// Add a Block to the chain
func addBlockToChain(newData string) BlockStruct {

	prevBlock := Blockchain[len(Blockchain)-1]
	newBlock := createNewBlock(prevBlock, newData)

	// CHECK IF NEWBLOCK IS VALID
	if isBlockValid(newBlock, prevBlock) {
		Blockchain = append(Blockchain, newBlock)
	}

	fmt.Printf("\nAdding new Block:\n%v\n\n", newBlock)
	return newBlock
}
