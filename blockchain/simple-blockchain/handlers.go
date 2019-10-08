// my-go-examples simple-blockchain handlers.go

package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

// GET
func rootHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res, "You are in the rootHandler\n\n")

    // GET BLOCKCHAIN
    // Not needed, but I want to go through engine
    theBlockchain := getBlockchain()

	// RETURN BLOCKCHAIN
	js, _ := json.MarshalIndent(theBlockchain, "", "    ")
	res.Write(js)

}

// GET
func showBlockHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

    // GET BLOCK ID
    blockID := params["blockID"] 

    // GET BLOCK
    // Not needed, but I want to go through engine
    theblock := getBlock(blockID)

	// RETURN BLOCK
	json.NewEncoder(res).Encode(theblock)
}

// POST (add)
func addBlockHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	var newData Message

	// GET THE DATA TO PUT IN NEW BLOCK
	_ = json.NewDecoder(req.Body).Decode(&newData)

    // ADD BLOCK TO CHAIN
    newBlock := addBlockToChain(newData)

	// RETURN NEWBLOCK
	//json.NewEncoder(res).Encode(todosDatabase)
	js, _ := json.MarshalIndent(newBlock, "", "    ")
	res.Write(js)

}
