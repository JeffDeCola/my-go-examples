// my-go-examples blockchain control.go

package main

import (
	"encoding/json"
	"fmt"
	"os"

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

func receivingTransaction(txRequestMessageSignedDataString string) {

	s := "------------------------------------------------------------------"
	log.Info("receivingTransaction()           " + s)

	var trms txRequestMessageSignedStruct

	// RECEIVED TRANSACTION MESSAGE txRequestMessageSignedDataStringBad
	// Place transaction Request Message data in transaction Request Message struct
	txRequestMessageSignedDataStringByte := []byte(txRequestMessageSignedDataString)
	err := json.Unmarshal(txRequestMessageSignedDataStringByte, &trms)
	checkErr(err)

	// Check is message valid, get balance and add to pendingBlock
	s = "Received transaction message from " + trms.TxRequestMessage.SourceAddress + " to " +
		fmt.Sprint(trms.TxRequestMessage.Destinations)
	log.Info("receivingTransaction()           " + s)
	status := trms.processTransactionRequest()
	s = "The status of transaction message from " + trms.TxRequestMessage.SourceAddress + " to " +
		fmt.Sprint(trms.TxRequestMessage.Destinations) + " is " + status
	log.Info("receivingTransaction()           " + s)
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

	// RESET pendingBlock
	s = "RESET pendingBlock"
	log.Info("main()                           " + s)
	resetPendingBlock()

	// RECEIVING SOME TRANSACTION REQUEST MESSAGES
	s = "RECEIVING SOME TRANSACTION REQUEST MESSAGES"
	log.Info("main()                           " + s)
	receivingTransaction(txRequestMessageSignedDataString1)
	receivingTransaction(txRequestMessageSignedDataStringBad)

	// ADD pendingBlock TO THE blockchain
	s = "ADD pendingBlock TO THE blockchain. Adding block number " + fmt.Sprint(pendingBlock.BlockID)
	log.Info("main()                           " + s)
	blockchain = append(blockchain, pendingBlock)

	// RESET pendingBlock
	s = "RESET pendingBlock"
	log.Info("main()                           " + s)
	resetPendingBlock()

	receivingTransaction(txRequestMessageSignedDataString2)

	// ADD pendingBlock TO THE blockchain
	s = "ADD pendingBlock TO THE blockchain. Adding block number " + fmt.Sprint(pendingBlock.BlockID)
	log.Info("main()                           " + s)
	blockchain = append(blockchain, pendingBlock)

	// RESET pendingBlock
	s = "RESET pendingBlock"
	log.Info("main()                           " + s)
	resetPendingBlock()

	// RECEIVING SOME TRANSACTION REQUEST MESSAGES
	s = "RECEIVING SOME TRANSACTION REQUEST MESSAGES"
	log.Info("main()                           " + s)
	receivingTransaction(txRequestMessageSignedDataString3)
	receivingTransaction(txRequestMessageSignedDataString4)

	// ADD pendingBlock TO THE blockchain
	s = "ADD pendingBlock TO THE blockchain. Adding block number " + fmt.Sprint(pendingBlock.BlockID)
	log.Info("main()                           " + s)
	blockchain = append(blockchain, pendingBlock)

	// RESET pendingBlock
	s = "RESET pendingBlock"
	log.Info("main()                           " + s)
	resetPendingBlock()

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

	// SHOW THE pendingBlock
	fmt.Printf("\nThe pendingBlock is:\n\n")
	js, _ = json.MarshalIndent(pendingBlock, "", "    ")
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
