# simple-blockchain example

_A simple blockchain with a REST JSON API to view (GET) the blockchain
and add (POST) a block._

Table of Contents,

* tbd

Documentation and reference,

* Refer to my
  [simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST)
  for the webserver engine
* Refer to my
  [cheat sheet on blockchains](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/blockchain-cheat-sheet)
* I got a lot of code and inspiration from mycoralhealth
  [here](https://github.com/mycoralhealth/blockchain-tutorial/blob/master/main.go).

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## PREREQUISITES

* `go get -v -u github.com/davecgh/go-spew/spew`
  Rather the use fmt.Print, this will print out the structs and slices very pretty
* `go get -v -u github.com/gorilla/mux`
  Gorilla/mux is a popular router

## RUN

```bash
go run simple-blockchain.go guts.go blockchain.go engine.go router.go routes.go handlers.go logger.go
```

### GET (view entire Blockchain)

Then you can goto the webpage to see your first block,

[localhost:1234/](http://localhost:1234/)

You could also use curl from the command line,

```go
curl localhost:1234
```

### GET (show a block)

[localhost:1234//showblock/0](http://localhost:1234/showblock/0)

### POST (add block)

To add Data use `http POST`,

```bash
curl -H "Content-Type: application/json" -X POST -d '{"data":"Add this data for new block"}' localhost:1234/addblock
```

Now just refresh the webpage or use curl again,

[localhost:8080/](http://localhost:8080/)

```bash
curl localhost:8080
```







## HOW IT WORKS

I broke this into 3 main sections,

* GUTS
* ENGINE
* WEB

But first, lets define a block using a struct,

```go
type block struct {
    Index     int
    Timestamp string
    Data      string
    Hash      string
    PrevHash  string
}
```

And the blockchain itself,

```go
var blockchain []Block
```

### GUTS

These are the functions that will do the manipulation on the blockchain,

* `calculateHash()` Calculate the hash for a block
* `isBlockValid()` Check that a block is valid
* `generateBlock()` Generate a new block

### ENGINE

These function will listen and wait for instruction on what to do,

* ???()
* ???()

### WEB

This is the user interface via http GET and POST,

* handleGetBlockchain()
* handleWriteBlock()

