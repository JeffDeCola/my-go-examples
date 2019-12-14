// my-go-examples bitcoin-ledger.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

// TRANSACTION REQUEST MESSAGE

type txRequestMessageSignedStruct struct {
	TxRequestMessage txRequestMessageStruct `json:"txRequestMessage"`
	Signature        string                 `json:"signature"`
}

type txRequestMessageStruct struct {
	SourceAddress string              `json:"sourceAddress"`
	Destinations  []destinationStruct `json:"destinations"`
}

type destinationStruct struct {
	DestinationAddress string `json:"destinationAddress"`
	Value              int64  `json:"value"`
}

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
	var transaction transactionStruct

	// INCREMENT BlockID
	currentBlock.BlockID = currentBlock.BlockID + 1

	// INCREMENT TxID
	transaction.TxID = currentBlock.Transactions[0].TxID + 1
	transactionSlice = append(transactionSlice, transaction)

	// LOAD THE BLOCK
	currentBlock.Transactions = transactionSlice

}

func (trms txRequestMessageSignedStruct) processTransactionRequest() string {

	// STEP 1 - VERIFY SIGNATURE
	if !(trms.verifySignature()) {
		return "Signature failed"
	}

	// STEP 2 - CHECK BALANCE TO SEE IF YOU HAVE THE MONEY
	// Returns list of the TxID of output unspent transactions
	balance, unspentOutputTxID := getBalance(trms.TxRequestMessage.SourceAddress)
	if balance == 0 {
		return "Not enough money"
	}

	// STEP 3 - LOAD currentBlock WITH TRANSACTION and UPDATE BALANCE
	trms.loadTxRequestMessageSignedToCurrentBlock(unspentOutputTxID)

	return "Pending"

}

// STEP 1 - VERIFY SIGNATURE
func (trms txRequestMessageSignedStruct) verifySignature() bool {

	if trms.Signature == "Bad" {
		return false
	}

	return true

}

// STEP 2 - CHECK BALANCE TO SEE IF YOU HAVE THE MONEY
func getBalance(address string) (int64, []int64) {

	if address == "Founders PubKey" {
		unspentOutputTxID := []int64{0}
		return 100000000, unspentOutputTxID
	}

	if address == "Jeffs PubKey" {
		unspentOutputTxID := []int64{1, 2}
		return 80000, unspentOutputTxID
	}

	unspentOutputTxID := []int64{0}
	return 0, unspentOutputTxID

}

// STEP 3 - LOAD currentBlock WITH TRANSACTION and UPDATE BALANCE
func (trms txRequestMessageSignedStruct) loadTxRequestMessageSignedToCurrentBlock(unspentOutputTxID []int64) {

	// Check if first transaction in currentBlock
	first := false
	if currentBlock.Transactions[0].Inputs == nil {
		first = true
	}

	//-------------------------------------------------------
	// STEP 3.1 - BUILD INPUT STRUCT
	//-------------------------------------------------------
	var inputsTemp = inputsStruct{}

	inputsTemp.InID = 22222222
	inputsTemp.RefTxID = 333333333
	inputsTemp.InPubKey = trms.TxRequestMessage.SourceAddress
	inputsTemp.Signature = trms.Signature

	// Place in slice
	var inputsSlice []inputsStruct
	inputsSlice = append(inputsSlice, inputsTemp)

	//-------------------------------------------------------
	// STEP 3.2 - BUILD OUTPUT STRUCT
	//-------------------------------------------------------
	var outputsTemp = outputsStruct{}

	outputsTemp.OutPubKey = trms.TxRequestMessage.Destinations[0].DestinationAddress
	outputsTemp.Value = trms.TxRequestMessage.Destinations[0].Value

	// Place in slice
	var outputsSlice []outputsStruct
	outputsSlice = append(outputsSlice, outputsTemp)

	//-------------------------------------------------------
	// STEP 3.3 - BUILD THE TRANSACTON
	//-------------------------------------------------------
	var transactionTemp = transactionStruct{}

	// Check if first transaction
	if first {
		transactionTemp.TxID = currentBlock.Transactions[0].TxID
	} else {
		// Increment TxID
		transactionTemp.TxID = currentBlock.Transactions[len(currentBlock.Transactions)-1].TxID + 1
	}
	transactionTemp.Inputs = inputsSlice
	transactionTemp.Outputs = outputsSlice

	//-------------------------------------------------------
	// STEP 3.4 - PLACE transactionStruct IN transactionSlice
	//-------------------------------------------------------
	var transactionSlice []transactionStruct

	// Check if first transaction
	if first {
		transactionSlice = append(transactionSlice, transactionTemp)
	} else {
		transactionSlice = append(currentBlock.Transactions, transactionTemp)
	}

	//-------------------------------------------------------
	// STEP 3.5 - LOAD THE BLOCK
	//-------------------------------------------------------
	currentBlock.Transactions = transactionSlice

}

func showbalance(pubKey string) int64 {
	return 675
}

func main() {

	// GENISIS BLOCKCHAIN
	// LOAD BLOCK 0 WITH 0000 TRANSACTION
	genesisBlockchain(genesisTransactionString)

	// RESET currentBlock
	resetCurrentBlock()

	// SHOW THE currentBlock
	fmt.Printf("\nThe currentBlock is:\n\n")
	js, _ := json.MarshalIndent(currentBlock, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	// RECEIVED TRANSACTION MESSAGE txRequestMessageSignedDataStringN
	// Place transaction Request Message data in transaction Request Message struct
	var trms txRequestMessageSignedStruct
	txRequestMessageSignedDataStringByte := []byte(txRequestMessageSignedDataString1)
	err := json.Unmarshal(txRequestMessageSignedDataStringByte, &trms)
	checkErr(err)
	// Check is message valid, get balance and add to currentBlock
	status := trms.processTransactionRequest()
	fmt.Printf("\nThe transaction message from %v to %v is %v\n\n",
		trms.TxRequestMessage.SourceAddress, trms.TxRequestMessage.Destinations, status)

	// SHOW THE currentBlock
	fmt.Printf("\nThe currentBlock is:\n\n")
	js, _ = json.MarshalIndent(currentBlock, "", "    ")
	fmt.Printf("%v\n\n", string(js))
	os.Exit(1)

	// RECEIVED TRANSACTION MESSAGE txRequestMessageSignedDataStringBad
	// Place transaction Request Message data in transaction Request Message struct
	txRequestMessageSignedDataStringByte = []byte(txRequestMessageSignedDataStringBad)
	err = json.Unmarshal(txRequestMessageSignedDataStringByte, &trms)
	checkErr(err)
	// Check is message valid, get balance and add to currentBlock
	status = trms.processTransactionRequest()
	fmt.Printf("\nThe transaction message from %v to %v is %v\n\n",
		trms.TxRequestMessage.SourceAddress, trms.TxRequestMessage.Destinations, status)

	// RECEIVED TRANSACTION MESSAGE txRequestMessageSignedDataStringN
	// Place transaction Request Message data in transaction Request Message struct
	txRequestMessageSignedDataStringByte = []byte(txRequestMessageSignedDataString2)
	err = json.Unmarshal(txRequestMessageSignedDataStringByte, &trms)
	checkErr(err)
	// Check is message valid, get balance and add to currentBlock
	status = trms.processTransactionRequest()
	fmt.Printf("\nThe transaction message from %v to %v is %v\n\n",
		trms.TxRequestMessage.SourceAddress, trms.TxRequestMessage.Destinations, status)

	// ADD currentBlock TO THE blockchain
	blockchain = append(blockchain, currentBlock)

	// SHOW THE blockchain
	fmt.Printf("\nThe blockchain is:\n\n")
	js, _ = json.MarshalIndent(blockchain, "", "    ")
	fmt.Printf("%v\n\n", string(js))
	os.Exit(1)

	// RESET currentBlock
	resetCurrentBlock()

	// RECEIVED TRANSACTION MESSAGE txRequestMessageSignedDataStringN
	// Place transaction Request Message data in transaction Request Message struct
	txRequestMessageSignedDataStringByte = []byte(txRequestMessageSignedDataString3)
	err = json.Unmarshal(txRequestMessageSignedDataStringByte, &trms)
	checkErr(err)
	// Check is message valid, get balance and add to currentBlock
	status = trms.processTransactionRequest()
	fmt.Printf("\nThe transaction message from %v to %v is %v\n\n",
		trms.TxRequestMessage.SourceAddress, trms.TxRequestMessage.Destinations, status)

	// RECEIVED TRANSACTION MESSAGE txRequestMessageSignedDataStringN
	// Place transaction Request Message data in transaction Request Message struct
	txRequestMessageSignedDataStringByte = []byte(txRequestMessageSignedDataString4)
	err = json.Unmarshal(txRequestMessageSignedDataStringByte, &trms)
	checkErr(err)
	// Check is message valid, get balance and add to currentBlock
	status = trms.processTransactionRequest()
	fmt.Printf("\nThe transaction message from %v to %v is %v\n\n",
		trms.TxRequestMessage.SourceAddress, trms.TxRequestMessage.Destinations, status)

	// ADD currentBlock TO THE blockchain
	blockchain = append(blockchain, currentBlock)

	// RESET currentBlock
	resetCurrentBlock()

	// SHOW THE blockchain
	fmt.Printf("\nThe blockchain is:\n\n")
	js, _ = json.MarshalIndent(blockchain, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	// balance := showbalance(mattPubKey)
	//fmt.Printf("The balance for %s is %d\n\n", mattPubKey, balance)

}
