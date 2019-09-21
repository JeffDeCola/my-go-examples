# simple-blockchain example

`simple-blockchain` _is an example of
a blockchain with a web interface to view and add data (blocks)._

Refer to my
[cheat sheet on blockchains](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/blockchain-cheat-sheet)

I got a lot of code and inspiration from mycoralhealth
[here](https://github.com/mycoralhealth/blockchain-tutorial/blob/master/main.go).

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## PREREQUISITES

* `go get -v -u github.com/davecgh/go-spew/spew`
  Rather the use fmt.Print, this will print out the structs and slices very pretty
* `go get -v -u github.com/gorilla/mux`
  Gorilla/mux is a popular package for writing web handlers

## RUN

```bash
go run simple-blockchain.go guts.go engine.go web.go
```

Then you can goto the webpage to see your first block,

[http://localhost:8080/](http://localhost:8080/)

You could also use curl from the command line,

```go
curl localhost:8080
```

To add Data use `http POST`,

```bash
curl -H "Content-Type: application/json" -X POST -d '{"data":"hi jeff"}' localhost:8080
```

Now just refresh the webpage or use curl again,

[http://localhost:8080/](http://localhost:8080/)

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

