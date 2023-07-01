# BITCOIN LEDGER EXAMPLE

_Demonstrates a bitcoin ledger in a blockchain using the
unspent transaction output model._

Table of Contents

* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/bitcoin-ledger#prerequisites)
* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/bitcoin-ledger#overview)
* [THIS EXAMPLE](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/bitcoin-ledger#this-example)
* [BLOCKCHAIN AND PENDING BLOCK](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/bitcoin-ledger#blockchain-and-pending-block)
* [TRANSACTIONS IN LEDGER](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/bitcoin-ledger#transactions-in-ledger)
* [ADDING A TRANSACTION TO THE pendingBlock](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/bitcoin-ledger#adding-a-transaction-to-the-pendingblock)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/bitcoin-ledger#run)

Documentation and Reference

* My cheat sheet on
[blockchains](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/blockchain-cheat-sheet)
* A
[blockchain with REST](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/single-node-blockchain-with-REST)
example I wrote that just stores data (no coins)
* I also wrote an entire cryptocurrency called
  [jeffCoin](https://github.com/JeffDeCola/jeffCoin)

## PREREQUISITES

```bash
go get -u -v github.com/sirupsen/logrus
```

## OVERVIEW

An **unspent transaction output model** is a method to log the transactions
of coins/value in a cryptocurrency. This type of bookkeeping does not keep a
running balance. Each transaction records where the coin/value
is coming from (input) and where it is going to (output).

All inputs must be tied to other outputs, but outputs can be `"unspent"`, meaning
not connected to anything.  When you add up all the unspent outputs for an address
you have the balance.

It's a little tricky at first but makes sense if you code it.

## THIS EXAMPLE

To keep this example really simple, blocks will only contain an ID and transactions.

```go
type blockStruct struct {
    BlockID      int64               `json:"blockID"`
    Transactions []transactionStruct `json:"transactions"`
}
```

Where a transaction is just a slice of inputs and outputs,

```go
type transactionStruct struct {
    TxID    int64           `json:"txID"`
    Inputs  []inputsStruct  `json:"inputs"`
    Outputs []outputsStruct `json:"outputs"`
}
```

The first block (genesis) in the blockchain will contain 1 transaction.
That will be the founders initial value of 100,000,000
(or 100,000 jeffCoins). Each jeffCoin
has a value of 1,000 addies. _Meow._

* **BlockID 0** _Genesis Block_
  * **TxID 0** _Founders start with  100,000 jeffCoins_

There are 7 transaction requests thereafter (One is bad).
I will be adding blocks as follows,

* **BlockID 1**
  * **TxID 1** _Founders sends Jeff 80 jeffCoins_
  * _Founders send someone 80 jeffCoins (Rejected Signature Failed)_
* **BlockID 2**  
  * **TxID 2** _Jeff sends Matt 50 jeffCoins & CoinVault .5 jeffCoin_
* **BlockID 3**  
  * **TxID 3** _Founders sends Matt 250 jeffCoins & Jeff 13 jeffCoins_
  * **TxID 4** _Matt sends Jill 35 jeffCoins_

Then all other transactions will be pending,

* **pendingBlock**
  * **TxID 5** _Matt sends Jeff 15 jeffCoins_
  * **TxID 6** _Jeff sends Jill 33 jeffCoins_

The `pendingBlock` is pending, meaning it needs to be verified by other nodes.
So this block will not have any value until its part of the blockchain.

## BLOCKCHAIN AND PENDING BLOCK

After the run the blockchain and pendingBlock should look like
[blockchain-output.txt](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/bitcoin-ledger/blockchain-output.txt).

This illustration may help,

![IMAGE - pendingBlock-and-blockchain-flow - IMAGE](../../docs/pics/blockchain/pendingBlock-and-blockchain-flow.jpg)

## TRANSACTIONS IN LEDGER

This illustrations shows the ledger in the blockchain and pendingBlock,

![IMAGE - bitcoin-transactions-inputs-and-outputs - IMAGE](../../docs/pics/blockchain/bitcoin-transactions-inputs-and-outputs.jpg)

This illustration shows a visual look at how the transactions relate
(input/output) to each other in the blockchain,

![IMAGE - bitcoin-transaction-flow-from-block-to-block - IMAGE](../../docs/pics/blockchain/bitcoin-transaction-flow-from-block-to-block.jpg)

## ADDING A TRANSACTION TO THE pendingBlock

A transaction request message is just a source address requesting to
transfer value/coins to one ore more destination addresses.

This is easily seen using the following struct,

```go
type txRequestMessageStruct struct {
    SourceAddress string              `json:"sourceAddress"`
    Destinations  []destinationStruct `json:"destinations"`
}

type destinationStruct struct {
    DestinationAddress string `json:"destinationAddress"`
    Value              int64  `json:"value"`
}
```

Then it is signed for verification,

```go
type txRequestMessageSignedStruct struct {
    TxRequestMessage txRequestMessageStruct `json:"txRequestMessage"`
    Signature        string                 `json:"signature"`
}
```

Below are the **functions/methods** for processing the
transaction request message. The ultimate goal is to load the
transaction onto the `pendingBlock` which will be eventually appended
onto the blockchain.

 I broke it down into 5 steps (verify signature is just a mockup for simplicity),

* **processTransactionRequest()**
  * STEP 1 - VERIFY SIGNATURE
    * **verifySignature()**
  * STEP 2 - GET BALANCE AND A LIST OF UNSPENT OUTPUTS
    * **getBalance()**
  * STEP 3 - CHECK IF YOU HAVE ENOUGH jeffCoins
  * STEP 4 - PICK THE UNSPENT OUTPUTS TO USE AND PROVIDE CHANGE
    * **pickUnspentOutputs()**
  * STEP 5 - ADD TRANSACTION to pendingBlock and MAKE CHANGE
    * **addTransactionToPendingBlock()**

## RUN

This will process a transaction, display the blockchain and show the
balances for each address (public Keys),

```bash
go run control.go bitcoin-ledger.go data.go
```

The blockchain and pendingBlock should look like
[blockchain-output.txt](https://github.com/JeffDeCola/my-go-examples/blob/master/blockchain/bitcoin-ledger/blockchain-output.txt).

The balances in the blockchain should be,

```txt
The balance for Founders PubKey (Address) is 99657000
The balance for Jeffs PubKey (Address) is 42500
The balance for Matts PubKey (Address) is 265000
The balance for Jills PubKey (Address) is 35000
The balance for CoinVaults PubKey (Address) is 500
```

Remember, the pendingBlock is pending so it is not part
of this calculation.
