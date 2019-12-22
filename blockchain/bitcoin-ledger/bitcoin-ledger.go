// my-go-examples blockchain bitcoin-ledger.go

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"

	log "github.com/sirupsen/logrus"
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

	// INCREMENT TxID (From last one)
	transaction.TxID = currentBlock.Transactions[len(currentBlock.Transactions)-1].TxID + 1
	transactionSlice = append(transactionSlice, transaction)

	// LOAD THE BLOCK
	currentBlock.Transactions = transactionSlice

}

func receivingTransaction(txRequestMessageSignedDataString string) {

	s := "------------------------------------------------------------------"
	log.Info("receivingTransaction()           " + s)

	var trms txRequestMessageSignedStruct

	// RECEIVED TRANSACTION MESSAGE txRequestMessageSignedDataStringBad
	// Place transaction Request Message data in transaction Request Message struct
	txRequestMessageSignedDataStringByte := []byte(txRequestMessageSignedDataString)
	err := json.Unmarshal(txRequestMessageSignedDataStringByte, &trms)
	checkErr(err)

	// Check is message valid, get balance and add to currentBlock
	s = "Received transaction message from " + trms.TxRequestMessage.SourceAddress + " to " +
		fmt.Sprint(trms.TxRequestMessage.Destinations)
	log.Info("receivingTransaction()           " + s)
	status := trms.processTransactionRequest()
	s = "The status of transaction message from " + trms.TxRequestMessage.SourceAddress + " to " +
		fmt.Sprint(trms.TxRequestMessage.Destinations) + " is " + status
	log.Info("receivingTransaction()           " + s)
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

	// STEP 5 - LOAD currentBlock WITH TRANSACTION and MAKE CHANGE
	s = "STEP 5 - LOAD currentBlock WITH TRANSACTION and MAKE CHANGE"
	log.Info("processTransactionRequest()      " + s)
	trms.loadTRMSignedToCurrentBlock(useUnspentOutput, change)

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

// STEP 5 - LOAD currentBlock WITH TRANSACTION and MAKE CHANGE
func (trms txRequestMessageSignedStruct) loadTRMSignedToCurrentBlock(unspentOutputSlice []unspentOutputStruct, change int64) {

	// Check if first transaction in currentBlock
	first := false
	if currentBlock.Transactions[0].Inputs == nil {
		first = true
	}

	//-------------------------------------------------------
	// STEP 5.1 - BUILD INPUT STRUCT FOR EACH UNSPENT OUTPUT
	//-------------------------------------------------------
	s := "STEP 5.1 - BUILD INPUT STRUCT FOR EACH UNSPENT OUTPUT"
	log.Info("loadTRMSignedToCurrentBlock()    " + s)
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
	log.Info("loadTRMSignedToCurrentBlock()    " + s)
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
	log.Info("loadTRMSignedToCurrentBlock()    " + s)
	var transactionTemp = transactionStruct{}

	// Check if first transaction
	if first {
		transactionTemp.TxID = currentBlock.Transactions[0].TxID
	} else {
		transactionTemp.TxID = currentBlock.Transactions[len(currentBlock.Transactions)-1].TxID + 1
	}
	transactionTemp.Inputs = inputsSlice
	transactionTemp.Outputs = outputsSlice

	s = "The transactionTemp is " + fmt.Sprint(transactionTemp)
	log.Info("loadTRMSignedToCurrentBlock()    " + s)

	//-------------------------------------------------------
	// STEP 5.4 - PLACE transactionStruct IN transactionSlice
	//-------------------------------------------------------
	s = "STEP 5.4 - PLACE transactionStruct IN transactionSlice"
	log.Info("loadTRMSignedToCurrentBlock()    " + s)
	var transactionSlice []transactionStruct

	// Check if first transaction
	if first {
		transactionSlice = append(transactionSlice, transactionTemp)
	} else {
		transactionSlice = append(currentBlock.Transactions, transactionTemp)
	}

	//-------------------------------------------------------
	// STEP 5.5 - LOAD THE CURRENT BLOCK WITH TRANSACTION
	//-------------------------------------------------------
	s = "STEP 5.5 - LOAD THE CURRENT BLOCK WITH TRANSACTION"
	log.Info("loadTRMSignedToCurrentBlock()    " + s)

	currentBlock.Transactions = transactionSlice

}

func showbalance(pubKey string) int64 {
	return 675
}

func init() {

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

	// SET LOG LEVEL
	log.SetLevel(log.TraceLevel)
	// log.SetLevel(log.InfoLevel)

}

func main() {

	// GENISIS BLOCKCHAIN
	// LOAD BLOCK 0 WITH 0000 TRANSACTION
	s := "GENISIS BLOCKCHAIN / LOAD BLOCK 0 WITH 0000 TRANSACTION"
	log.Info("main()                           " + s)
	genesisBlockchain(genesisTransactionString)

	// RESET currentBlock
	s = "RESET currentBlock"
	log.Info("main()                           " + s)
	resetCurrentBlock()

	// RECEIVING SOME TRANSACTION REQUEST MESSAGES
	s = "RECEIVING SOME TRANSACTION REQUEST MESSAGES"
	log.Info("main()                           " + s)
	receivingTransaction(txRequestMessageSignedDataString1)
	receivingTransaction(txRequestMessageSignedDataStringBad)

	// ADD currentBlock TO THE blockchain
	s = "ADD currentBlock TO THE blockchain. Adding block number " + fmt.Sprint(currentBlock.BlockID)
	log.Info("main()                           " + s)
	blockchain = append(blockchain, currentBlock)

	// RESET currentBlock
	s = "RESET currentBlock"
	log.Info("main()                           " + s)
	resetCurrentBlock()

	receivingTransaction(txRequestMessageSignedDataString2)

	// ADD currentBlock TO THE blockchain
	s = "ADD currentBlock TO THE blockchain. Adding block number " + fmt.Sprint(currentBlock.BlockID)
	log.Info("main()                           " + s)
	blockchain = append(blockchain, currentBlock)

	// RESET currentBlock
	s = "RESET currentBlock"
	log.Info("main()                           " + s)
	resetCurrentBlock()

	// RECEIVING SOME TRANSACTION REQUEST MESSAGES
	s = "RECEIVING SOME TRANSACTION REQUEST MESSAGES"
	log.Info("main()                           " + s)
	receivingTransaction(txRequestMessageSignedDataString3)
	receivingTransaction(txRequestMessageSignedDataString4)

	// ADD currentBlock TO THE blockchain
	s = "ADD currentBlock TO THE blockchain. Adding block number " + fmt.Sprint(currentBlock.BlockID)
	log.Info("main()                           " + s)
	blockchain = append(blockchain, currentBlock)

	// RESET currentBlock
	s = "RESET currentBlock"
	log.Info("main()                           " + s)
	resetCurrentBlock()

	// RECEIVING SOME TRANSACTION REQUEST MESSAGES
	s = "RECEIVING SOME TRANSACTION REQUEST MESSAGES"
	log.Info("main()                           " + s)
	receivingTransaction(txRequestMessageSignedDataString5)

	// RECEIVING SOME TRANSACTION REQUEST MESSAGES
	s = "RECEIVING SOME TRANSACTION REQUEST MESSAGES"
	log.Info("main()                           " + s)
	receivingTransaction(txRequestMessageSignedDataString6)

	// SHOW THE blockchain
	fmt.Printf("\nThe blockchain is:\n\n")
	js, _ := json.MarshalIndent(blockchain, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	// SHOW THE currentBlock
	fmt.Printf("\nThe currentBlock is:\n\n")
	js, _ = json.MarshalIndent(currentBlock, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	// SHOW BALANCES
	balance, _ := getBalance(foundersPubKey)
	fmt.Printf("The balance for %s (Address) is %d\n\n", foundersPubKey, balance)
	balance, _ = getBalance(jeffPubKey)
	fmt.Printf("The balance for %s (Address) is %d\n\n", jeffPubKey, balance)
	balance, _ = getBalance(mattPubKey)
	fmt.Printf("The balance for %s (Address) is %d\n\n", mattPubKey, balance)
	balance, _ = getBalance(jillPubKey)
	fmt.Printf("The balance for %s (Address) is %d\n\n", jillPubKey, balance)
	balance, _ = getBalance(coinVaultPubKey)
	fmt.Printf("The balance for %s (Address) is %d\n\n", coinVaultPubKey, balance)

}
