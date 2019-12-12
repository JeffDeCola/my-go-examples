// my-go-examples bitcoin-unspent-transaction-output-model.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	foundersPubKey  = "Founders PubKey"
	jeffPubKey      = "Jeffs PubKey"
	mattPubKey      = "Matts PubKey"
	coinVaultPubKey = "CoinVaults PubKey"
	jillPubKey      = "Jills PubKey"
	foundersSig     = "Founders Signature"
	jeffSig         = "Jeffs Signature"
	mattSig         = "Matts Signature"
)

// BLOCK

type blockStruct struct {
	BlockID      int64               `json:"blockID"`
	Transactions []transactionStruct `json:"transactions"`
}

var currentBlock = blockStruct{}

// BLOCKCHAIN

type blockchainSlice []blockStruct

var blockchain = blockchainSlice{}

// TRANSACTIONS

type transactionStruct struct {
	TxID    int64           `json:"txID"`
	Inputs  []inputsStruct  `json:"inputs"`
	Outputs []outputsStruct `json:"outputs"`
}

type inputsStruct struct {
	InID      int64  `json:"inID"`
	RefTxID   int64  `json:"refTxID"`
	InPubKey  string `json:"inPubKey"`
	Signature string `json:"signature"`
}

type outputsStruct struct {
	OutPubKey string `json:"outPubKey"`
	Value     int64  `json:"value"`
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func genesisBlockchain(transactionDataString string) {

	// LOAD THE TRANSACTION - Place transaction data in transaction struct
	var tempTransaction transactionStruct
	transactionDataByte := []byte(transactionDataString)
	err := json.Unmarshal(transactionDataByte, &tempTransaction)
	checkErr(err)

	// PLACE transaction IN transactionSlice
	var transactionSlice []transactionStruct
	transactionSlice = append(transactionSlice, tempTransaction)

	// LOAD THE GENESIS BLOCK with transactionSlice
	currentBlock.BlockID = 0
	currentBlock.Transactions = transactionSlice

	// ADD BLOCK TO THE BLOCKCHAIN
	blockchain = append(blockchain, currentBlock)

}

func resetCurrentBlock() {

	var transactionSlice []transactionStruct

	// LOAD THE BLOCK (with a null transaction)
	currentBlock.BlockID = currentBlock.BlockID + 1
	currentBlock.Transactions = transactionSlice

}

func loadTransactionToBlock(transactionDataString string) {

	// LOAD THE TRANSACTION - Place transaction data in transaction struct
	var tempTransaction transactionStruct
	transactionDataByte := []byte(transactionDataString)
	err := json.Unmarshal(transactionDataByte, &tempTransaction)
	checkErr(err)

	var transactionSlice []transactionStruct

	// PLACE transaction IN transactionSlice
	transactionSlice = append(currentBlock.Transactions, tempTransaction)

	// LOAD THE BLOCK
	currentBlock.Transactions = transactionSlice

}

func showbalance(pubKey string) int64 {
	return 675
}

func main() {

	// GENISIS BLOCKCHAIN
	// LOAD BLOCK WITH 0000 TRANSACTION
	genesisBlockchain(genesisDataString)

	// REEST BLOCK
	resetCurrentBlock()

	// LOAD BLOCK WITH 0001 TRANSACTION
	loadTransactionToBlock(transaction1DataString)

	// LOAD BLOCK WITH 0002 TRANSACTION
	loadTransactionToBlock(transaction2DataString)

	// ADD BLOCK TO THE BLOCKCHAIN
	blockchain = append(blockchain, currentBlock)

	// RESET BLOCK
	resetCurrentBlock()

	// LOAD BLOCK WITH 0003 TRANSACTION
	loadTransactionToBlock(transaction3DataString)

	// LOAD BLOCK WITH 0004 TRANSACTION
	loadTransactionToBlock(transaction4DataString)

	// SHOW THE BLOCKCHAIN
	//fmt.Printf("\nThe blockchain is:\n\n")
	//js, _ := json.MarshalIndent(blockchain, "", "    ")
	//fmt.Printf("%v\n\n", string(js))

	// SHOW THE BLOCK
	//fmt.Printf("\nThe block is:\n\n")
	//js, _ = json.MarshalIndent(currentBlock, "", "    ")
	//fmt.Printf("%v\n\n", string(js))

	balance := showbalance(mattPubKey)
	fmt.Printf("The balance for %s is %d\n\n", mattPubKey, balance)

}
