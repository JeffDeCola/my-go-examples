// my-go-examples multi-node-blockchain-with-REST-and-tcp-ip handlers.go

package webserver

import (
	"encoding/json"
	"io"
	"net/http"

	blockchain "github.com/JeffDeCola/my-go-examples/blockchain/multi-node-blockchain-with-REST-and-tcp-ip/blockchain"
	"github.com/gorilla/mux"
)

// GET
// Return entire Blockchain
func rootHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res, "Welcome, this is the entire Blockchain:\n\n")

	// GET BLOCKCHAIN
	// Not needed, but I want to go through engine
	theBlockchain := blockchain.GetBlockchain()

	// RESPOND BLOCKCHAIN
	js, _ := json.MarshalIndent(theBlockchain, "", "    ")
	res.Write(js)
	io.WriteString(res, "\n\n")

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
	io.WriteString(res, "The Block you requested:\n\n")
	js, _ := json.MarshalIndent(theblock, "", "    ")
	res.Write(js)
	io.WriteString(res, "\n\n")

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
	io.WriteString(res, "Added the Following Block to the chain:\n\n")
	js, _ := json.MarshalIndent(newBlock, "", "    ")
	res.Write(js)
	io.WriteString(res, "\n\n")

}
