// my-go-examples multi-node-blockchain-with-REST-and-tcp-ip guts.go

package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// GET HASH
// SHA256 hasing
func calculateBlockHash(block BlockStruct) string {

	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	log.Println("GUTS:           Calculated Block Hash")
	return hex.EncodeToString(hashed)

}

// CHECK THAT A NEW BLOCK IS VALID
func isBlockValid(checkBlock, oldBlock BlockStruct) bool {

	log.Println("GUTS:           Checking if Block is valid")

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

// CREATE NEW BLOCK
func createNewBlock(oldBlock BlockStruct, data string) BlockStruct {

	var newBlock BlockStruct

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateBlockHash(newBlock)

	log.Println("GUTS:           Created New Block")
	return newBlock
}

// REPLACE CHAIN WITH LONGER ONE
func replaceChain(newBlock BlockchainSlice) {

	if len(newBlock) > len(Blockchain) {
		log.Println("GUTS:           New Block added to chain")
		Blockchain = newBlock
	} else {
		log.Println("GUTS:           New Block NOT added to chain")
	}
}
