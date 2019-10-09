// my-go-examples multi-node-blockchain-with-REST-and-tcp-ip handlers.go

package tcpserver

import (
	"bufio"
	"encoding/json"
	"strings"

	blockchain "github.com/JeffDeCola/my-go-examples/blockchain/multi-node-blockchain-with-REST-and-tcp-ip/blockchain"
)

func handleAddNewBlock(rw *bufio.ReadWriter) {

	s := "Please enter the new Data"
	returnMessage(s, rw)

	// WAITING FOR DATA
	data, err := rw.ReadString('\n')
	checkErr(err)
	data = strings.Trim(data, "\n ")
	s = "Received DATA: " + data
	returnMessage(s, rw)

	// MAKE A NEW BLOCK
	// ADD NEW BLOCK TO CHAIN
	newBlock := blockchain.AddBlockToChain(data)
	js, _ := json.MarshalIndent(newBlock, "", "    ")
	s = "Added block to blockchain:\n" + string(js)
	returnMessage(s, rw)

}
