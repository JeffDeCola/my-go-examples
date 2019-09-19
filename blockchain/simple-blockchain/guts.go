// my-go-examples simple-blockchain guts.go

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// GET HASH
// SHA256 hasing
func calculateBlockHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// CHECK THAT A NEW BLOCK IS VALID
func isBlockValid(checkBlock, oldBlock Block) bool {

	// Check index
	if oldBlock.Index+1 != checkBlock.Index {
		return false
	}

	// Compare the hash matches
	if oldBlock.Hash != checkBlock.PrevHash {
		return false
	}

	// Recalculate Hash to check
	if calculateBlockHash(checkBlock) != checkBlock.Hash {
		return false
	}

	return true
}

// ADD BLOCK TO CHAIN
func addBlockToChain(oldBlock Block, data string) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateBlockHash(newBlock)

	return newBlock
}
