// my-go-examples multi-node-blockchain-with-REST-and-tcp-ip blockchain-interface.go

package blockchain

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// CreateBlockchain - Create a blockchain
func CreateBlockchain() {

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
	js, _ := json.MarshalIndent(firstBlock, "", "    ")

	fmt.Printf("%v\n", string(js))

	Blockchain = append(Blockchain, firstBlock)
}

// GetBlockchain - Get the Blockchain
func GetBlockchain() BlockchainSlice {

	return Blockchain

}

// GetBlock - Get a Block from the chain
func GetBlock(id string) BlockStruct {

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

// AddBlockToChain - Add a Block to the chain
func AddBlockToChain(newData string) BlockStruct {

	prevBlock := Blockchain[len(Blockchain)-1]
	newBlock := createNewBlock(prevBlock, newData)

	// CHECK IF NEWBLOCK IS VALID
	if isBlockValid(newBlock, prevBlock) {
        log.Println("Block is valid")
        newBlockchain := append(Blockchain, newBlock)
        // REPLACE WITH LONGER ONE
		replaceChain(newBlockchain)
	}

	fmt.Printf("\nAdding new Block:\n%v\n\n", newBlock)
	return newBlock
}
