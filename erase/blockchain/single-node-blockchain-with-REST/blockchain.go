// my-go-examples single-node-blockchain-with-REST blockchain.go

package main

// BlockStruct is your block
type BlockStruct struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevhash"`
}

// BlockchainSlice is my type
type BlockchainSlice []BlockStruct

// Blockchain is the blockchain
var Blockchain = BlockchainSlice{}
