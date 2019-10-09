# multi-node-blockchain-with-REST-and-tcp-ip example

_A multi node sha256 blockchain with a REST JSON API
(to view (GET) the blockchain and add (POST) a block)
and a tcp server to communicate between the nodes over ip.

Table of Contents,

* tbd

Documentation and reference,

* The blockchain is built from my
  [simple-blockchain-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/simple-blockchain-with-REST)
* The webserver is built from my
  [simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST)
* The tcp server is built from my
  [simple-tcp-ip-server](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-tcp-ip-server)
* Refer to my
  [cheat sheet on blockchains](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/blockchain-cheat-sheet)
* I got a lot of inspiration from
  [here](https://github.com/mycoralhealth/blockchain-tutorial/blob/master/main.go)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

This code is broken up into four parts,

* **guts** The blockchain code
* **blockchain-interface** The interface to the blockchain for th web and tcp server
* **Webserver** The API and gui
* **TCP Server** Top communicate between he nodes

This examples will,

* Create a Blockchain
* Hash a Block
* View the Entire Blockchain via a web GUI
* View a specific Block via a REST API
* Add a Block (with your data) to the chain via a REST API

This illustration may help,

![IMAGE - multi-node-blockchain-with-REST-and-tcp-ip - IMAGE](https://github.com/JeffDeCola/my-go-examples/blob/master/docs/pics/multi-node-blockchain-with-REST-and-tcp-ip.jpg)

## RUN

```bash
go run multi-node-blockchain-with-REST-and-tcp-ip.go
```

## WEBSERVER AND API

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

## TCP SERVER

Open a connection,

```txt
netcat -q -1 localhost 3333
```

