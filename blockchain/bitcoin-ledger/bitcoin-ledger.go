// my-go-examples blockchain bitcoin-ledger.go

package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// BLOCK

type blockStruct struct {
	BlockID      int64               `json:"blockID"`
	Transactions []transactionStruct `json:"transactions"`
}

var pendingBlock = blockStruct{}

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
	RefTxID   int64  `json:"refTxID"`
	InPubKey  string `json:"inPubKey"`
	Signature string `json:"signature"`
}

type outputsStruct struct {
	OutPubKey string `json:"outPubKey"`
	Value     int64  `json:"value"`
}

// UNSPENT OUTPUT

type unspentOutputStruct struct {
	TxID  int64 `json:"txID"`
	Value int64 `json:"value"`
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
	pendingBlock.BlockID = 0
	pendingBlock.Transactions = transactionSlice

	// ADD BLOCK TO THE BLOCKCHAIN
	blockchain = append(blockchain, pendingBlock)

}

func resetPendingBlock() {

	var transactionSlice []transactionStruct
	var transaction transactionStruct

	// INCREMENT BlockID
	pendingBlock.BlockID = pendingBlock.BlockID + 1

	// INCREMENT TxID (From last one)
	transaction.TxID = pendingBlock.Transactions[len(pendingBlock.Transactions)-1].TxID + 1
	transactionSlice = append(transactionSlice, transaction)

	// LOAD THE BLOCK
	pendingBlock.Transactions = transactionSlice

}

func (trms txRequestMessageSignedStruct) processTransactionRequest() string {

	// STEP 1 - MOCK - VERIFY SIGNATURE
	s := "STEP 1 - MOCK - VERIFY SIGNATURE"
	log.Info("processTransactionRequest()      " + s)
	if !(trms.verifySignature()) {
		return "Signature failed"
	}

	// STEP 2 - CHECK BALANCE TO SEE IF YOU HAVE THE MONEY
	// Returns entire list of the TxID of output unspent transactions
	s = "STEP 2 - CHECK BALANCE TO SEE IF YOU HAVE THE MONEY"
	log.Info("processTransactionRequest()      " + s)
	balance, unspentOutput := getBalance(trms.TxRequestMessage.SourceAddress)

	// STEP 3 - CHECK IF YOU HAVE ENOUGH jeffCoins
	s = "STEP 3 - CHECK IF YOU HAVE ENOUGH jeffCoins"
	log.Info("processTransactionRequest()      " + s)
	var value int64
	for _, destinations := range trms.TxRequestMessage.Destinations {
		value = value + destinations.Value
	}
	s = "The balance for " + trms.TxRequestMessage.SourceAddress + " is " + strconv.FormatInt(balance, 10) +
		" and value to remove is " + strconv.FormatInt(value, 10) + " from " + fmt.Sprint(unspentOutput)
	log.Info("processTransactionRequest()      " + s)
	if balance < value {
		return "Not enough money"
	}

	// STEP 4 - PICK THE UNSPENT OUTPUTS TO USE AND PROVIDE CHANGE
	s = "STEP 4 - PICK THE UNSPENT OUTPUTS TO USE AND PROVIDE CHANGE"
	log.Info("processTransactionRequest()      " + s)
	useUnspentOutput, change := pickUnspentOutputs(unspentOutput, value)
	s = "You are using unspent outputs " + fmt.Sprint(useUnspentOutput)
	log.Info("processTransactionRequest()      " + s)
	s = "The change will be " + strconv.FormatInt(change, 10)
	log.Info("processTransactionRequest()      " + s)

	// STEP 5 - ADD TRANSACTION to pendingBlock and MAKE CHANGE
	s = "STEP 5 - ADD TRANSACTION to pendingBlock and MAKE CHANGE"
	log.Info("processTransactionRequest()      " + s)
	trms.addTransactionToPendingBlock(useUnspentOutput, change)

	return "Pending"

}

// STEP 1 - MOCK - VERIFY SIGNATURE
func (trms txRequestMessageSignedStruct) verifySignature() bool {

	if trms.Signature == "Bad" {
		return false
	}

	return true

}

// STEP 2 - CHECK BALANCE TO SEE IF YOU HAVE THE MONEY
// Returns entire list of the TxID of output unspent transactions
func getBalance(address string) (int64, []unspentOutputStruct) {

	unspentOutputMap := make(map[int64]int64)

	var unspentOutput unspentOutputStruct
	unspentOutputSlice := []unspentOutputStruct{}

	// STEP 2.1 - GET UNSPENT OUTPUT TRANSACTIONS - Make unspentOutputSlice
	s := "STEP 2.1 - GET UNSPENT OUTPUT TRANSACTIONS  - Make unspentOutputSlice"
	log.Info("getBalance()                     " + s)

	// LETS ITERATE OVER ALL TRANSACTIONS IN BLOCK CHAIN
	for _, blocks := range blockchain {

		for _, transaction := range blocks.Transactions {

			// ITERATE OVER INPUTS - find inputs with address
			for _, input := range transaction.Inputs {
				if address == input.InPubKey {
					// DID AN OUTPUT USE THIS? If so, delete from map.
					refTxID := input.RefTxID
					if unspentOutputMap[refTxID] != 0 {
						delete(unspentOutputMap, refTxID)
					}
				}

			}

			// ITERATE OVER OUTPUTS - find outputs with address
			for _, output := range transaction.Outputs {
				if address == output.OutPubKey {
					// Place in map
					unspentOutputMap[transaction.TxID] = output.Value
				}
			}
		}
	}

	// PUT MAP unspentOutputMap IN SLICE unspentOutputSlice
	for k, v := range unspentOutputMap {
		unspentOutput = unspentOutputStruct{k, v}
		unspentOutputSlice = append(unspentOutputSlice, unspentOutput)
	}

	// SORT SLICE FROM TxID (LOW TO HIGH)
	sort.Slice(unspentOutputSlice, func(i, j int) bool {
		return unspentOutputSlice[i].TxID < unspentOutputSlice[j].TxID
	})

	// STEP 2.2 - GET BALANCE from unspentOutputSlice
	s = "STEP 2.2 - GET BALANCE from unspentOutputSlice"
	log.Info("getBalance()                     " + s)
	var balance int64
	balance = 0
	for _, unspentOutput := range unspentOutputSlice {
		balance = balance + unspentOutput.Value
	}

	return balance, unspentOutputSlice

}

// STEP 4 - PICK THE UNSPENT OUTPUTS TO USE
func pickUnspentOutputs(pickUnspentOutputSlice []unspentOutputStruct, value int64) ([]unspentOutputStruct, int64) {

	var unspentOutputStructTemp = unspentOutputStruct{}
	var useUnspentOutputSlice []unspentOutputStruct

	var change int64
	var runningTotal int64

	// Once you hit the value, stop
	for _, unspentOutput := range pickUnspentOutputSlice {

		unspentOutputStructTemp.TxID = unspentOutput.TxID
		unspentOutputStructTemp.Value = unspentOutput.Value

		// Place in slice
		useUnspentOutputSlice = append(useUnspentOutputSlice, unspentOutputStructTemp)

		runningTotal = runningTotal + unspentOutput.Value

		// did you get enough - If yes, provide change
		if value < runningTotal {
			change = runningTotal - value
			break
		}

	}

	return useUnspentOutputSlice, change

}

// STEP 5 - ADD TRANSACTION to pendingBlock and MAKE CHANGE
func (trms txRequestMessageSignedStruct) addTransactionToPendingBlock(unspentOutputSlice []unspentOutputStruct, change int64) {

	// Check if first transaction in pendingBlock
	first := false
	if pendingBlock.Transactions[0].Inputs == nil {
		first = true
	}

	//-------------------------------------------------------
	// STEP 5.1 - BUILD INPUT STRUCT FOR EACH UNSPENT OUTPUT
	//-------------------------------------------------------
	s := "STEP 5.1 - BUILD INPUT STRUCT FOR EACH UNSPENT OUTPUT"
	log.Info("addTransactionToPendingBlock()   " + s)
	var inputsTemp = inputsStruct{}
	var inputsSlice []inputsStruct

	// Using the following unspent outputs
	for _, unspentOutput := range unspentOutputSlice {
		inputsTemp.RefTxID = unspentOutput.TxID
		inputsTemp.InPubKey = trms.TxRequestMessage.SourceAddress
		inputsTemp.Signature = trms.Signature

		// Place in slice
		inputsSlice = append(inputsSlice, inputsTemp)
	}

	//-------------------------------------------------------
	// STEP 5.2 - BUILD OUTPUT STRUCT
	//-------------------------------------------------------
	s = "STEP 5.2 - BUILD OUTPUT STRUCT"
	log.Info("addTransactionToPendingBlock()   " + s)
	var outputsTemp = outputsStruct{}
	var outputsSlice []outputsStruct

	// Using the following destinations
	for _, destination := range trms.TxRequestMessage.Destinations {

		outputsTemp.OutPubKey = destination.DestinationAddress
		outputsTemp.Value = destination.Value

		// Place in slice
		outputsSlice = append(outputsSlice, outputsTemp)
	}

	// PROVIDE CHANGE if > 0
	if change > 0 {

		outputsTemp.OutPubKey = trms.TxRequestMessage.SourceAddress
		outputsTemp.Value = change

		// Place in slice
		outputsSlice = append(outputsSlice, outputsTemp)

	}

	//-------------------------------------------------------
	// STEP 5.3 - BUILD THE TRANSACTION
	//-------------------------------------------------------
	s = "STEP 5.3 - BUILD THE TRANSACTION"
	log.Info("addTransactionToPendingBlock()   " + s)
	var transactionTemp = transactionStruct{}

	// Check if first transaction
	if first {
		transactionTemp.TxID = pendingBlock.Transactions[0].TxID
	} else {
		transactionTemp.TxID = pendingBlock.Transactions[len(pendingBlock.Transactions)-1].TxID + 1
	}
	transactionTemp.Inputs = inputsSlice
	transactionTemp.Outputs = outputsSlice

	s = "The transactionTemp is " + fmt.Sprint(transactionTemp)
	log.Info("addTransactionToPendingBlock()   " + s)

	//-------------------------------------------------------
	// STEP 5.4 - PLACE transactionStruct IN transactionSlice
	//-------------------------------------------------------
	s = "STEP 5.4 - PLACE transactionStruct IN transactionSlice"
	log.Info("addTransactionToPendingBlock()   " + s)
	var transactionSlice []transactionStruct

	// Check if first transaction
	if first {
		transactionSlice = append(transactionSlice, transactionTemp)
	} else {
		transactionSlice = append(pendingBlock.Transactions, transactionTemp)
	}

	//-------------------------------------------------------
	// STEP 5.5 - LOAD THE PENDING BLOCK WITH TRANSACTION
	//-------------------------------------------------------
	s = "STEP 5.5 - LOAD THE PENDING BLOCK WITH TRANSACTION"
	log.Info("addTransactionToPendingBlock()   " + s)

	pendingBlock.Transactions = transactionSlice

}
