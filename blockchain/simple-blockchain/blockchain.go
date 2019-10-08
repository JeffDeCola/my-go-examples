// my-go-examples simple-blockchain blockchain.go

package main

// BlockStruct is your block
type BlockStruct struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevhash"`
}

// Blockchain Slice
type Blockchain []BlockStruct
