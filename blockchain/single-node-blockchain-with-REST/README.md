# single-node-blockchain-with-REST example

_A simple single node sha256 blockchain with a REST JSON API
(to view (GET) the blockchain and add (POST) a block)._

Table of Contents,

* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/single-node-blockchain-with-REST#prerequisites)
* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/single-node-blockchain-with-REST#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/single-node-blockchain-with-REST#run)
  * [GET (View the entire Blockchain)](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/single-node-blockchain-with-REST#get-view-the-entire-blockchain)
  * [GET (Show a Particular Block)](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/single-node-blockchain-with-REST#get-show-a-particular-block)
  * [POST (Add a Block)](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/single-node-blockchain-with-REST#post-add-a-block)
* [HOW IT WORKS](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/single-node-blockchain-with-REST#how-it-works)

Documentation and reference,

* Refer to my
  [simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST)
  for the webserver engine
* Refer to my
  [cheat sheet on blockchains](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/blockchain-cheat-sheet)
* I got a lot of inspiration from
  [here](https://github.com/nosequeldeebee/blockchain-tutorial)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## PREREQUISITES

Gorilla/mux is a popular router I use for the webserver.

```bash
go get -v -u github.com/gorilla/mux
```

## OVERVIEW

This code is broken up into three parts,

* **guts** The blockchain code
* **blockchain-interface** The interface to the blockchain
* **Webserver** The API and gui

This examples will,

* Create a Blockchain
* Hash a Block
* View the Entire Blockchain via a web GUI
* View a specific Block via a REST API
* Add a Block (with your data) to the chain via a REST API

This illustration may help,

![IMAGE - single-node-blockchain-with-REST - IMAGE](https://github.com/JeffDeCola/my-go-examples/blob/master/docs/pics/single-node-blockchain-with-REST.jpg)

## RUN

```bash
go run single-node-blockchain-with-REST.go \
       guts.go blockchain.go blockchain-interface.go \
       router.go routes.go handlers.go logger.go
```

### GET (View the entire Blockchain)

Then you can goto the webpage to see your first block,

[localhost:1234/](http://localhost:1234/)

You could also use curl from the command line,

```go
curl localhost:1234
```

### GET (Show a Particular Block)

[localhost:1234//showblock/0](http://localhost:1234/showblock/0)

```go
curl localhost:1234/showblock/0
```

### POST (Add a Block)

```bash
curl -H "Content-Type: application/json" \
     -X POST \
     -d '{"data":"Add this data for new block"}' \
     localhost:1234/addblock
```

Check,

[localhost:1234/](http://localhost:1234/)

## HOW IT WORKS

We will just be looking at the guts, ignoring the webserver and blockchain-interface.

So it actually becomes quite simple. A Block is just a struct and The Blockchain
is simply a slice of structs. That's it.

```go
type BlockStruct struct {
    Index     int    `json:"index"`
    Timestamp string `json:"timestamp"`
    Data      string `json:"data"`
    Hash      string `json:"hash"`
    PrevHash  string `json:"prevhash"`
}

type BlockchainSlice []BlockStruct

var Blockchain = BlockchainSlice{}
```

And `guts.go` contain the three basic things you want to do
to that Blockchain (slice of structs),

* calculateBlockHash()
* isBlockValid()
* createNewBlock()
