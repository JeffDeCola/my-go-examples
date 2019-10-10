// my-go-examples multi-node-blockchain-with-REST-and-tcp-ip handlers.go

package webserver

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	blockchain "github.com/JeffDeCola/my-go-examples/blockchain/multi-node-blockchain-with-REST-and-tcp-ip/blockchain"
	"github.com/gorilla/mux"
)

// GET
// Return entire Blockchain
func rootHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	// GET BLOCKCHAIN
	// Not needed, but I want to go through engine
	theBlockchain := blockchain.GetBlockchain()

	// RESPOND BLOCKCHAIN
	s := "The entire Blockchain:"
	respondMessage(s, res)
	js, _ := json.MarshalIndent(theBlockchain, "", "    ")
	s = string(js)
	log.Println("WEBSERVER:      " + "Blockchain too long, not shown")
	io.WriteString(res, s+"\n")
}

// GET
func showBlockHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// GET BLOCK ID
	blockID := params["blockID"]

	// GET BLOCK - ask interface
	theblock := blockchain.GetBlock(blockID)

	// RESPOND BLOCK
	s := "The Block you requested:"
	respondMessage(s, res)
	js, _ := json.MarshalIndent(theblock, "", "    ")
	s = string(js)
	respondMessage(s, res)

}

// POST (add)
func addBlockHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	var newData Message

	// GET THE DATA TO PUT IN NEW BLOCK
	_ = json.NewDecoder(req.Body).Decode(&newData)

	// ADD NEW BLOCK TO CHAIN
	newBlock := blockchain.AddBlockToChain(newData.Data)

	// RESPOND NEWBLOCK
	//json.NewEncoder(res).Encode(todosDatabase)
	s := "Added the Following Block to the chain:"
	respondMessage(s, res)
	js, _ := json.MarshalIndent(newBlock, "", "    ")
	s = string(js)
	respondMessage(s, res)
}

func respondMessage(s string, res http.ResponseWriter) {

	log.Println("WEBSERVER:      " + s)
	io.WriteString(res, s+"\n")

}
