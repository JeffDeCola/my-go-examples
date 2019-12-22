# bitcoin-ledger example

_Demonstrates the bitcoin ledger in a blockchain using the
**unspent transaction output model**._

Table of Contents,

* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/bitcoin-ledger#prerequisites)
* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/bitcoin-ledger#overview)
* [THIS EXAMPLE](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/bitcoin-ledger#this-example)
* [TRANSACTIONS IN LEDGER](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/bitcoin-ledger#transactions-in-ledger)
* [ADDING A TRANSACTION TO currentBlock](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/bitcoin-ledger#adding-a-transaction-to-currentblock)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/bitcoin-ledger#run)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## PREREQUISITES

```bash
go get -u -v github.com/sirupsen/logrus
```

## OVERVIEW

An **unspent transaction output model** is a method to log the transactions
of coins in a cryptocurrency. This type of bookkeeping does not keep a
running balance. Each transaction records where the value
is coming from (input) and where it is going to (output).

All inputs must be tied to other outputs, but outputs can be `"unspent"`, meaning
not connected to anything.  When you add up all the unspent outputs for an address
you have the balance.

It's a little tricky at first but makes sense if you're forced to code it.

## THIS EXAMPLE

The first block in the blockchain will contain 1 transaction.  That will be the
founders initial value of 100,000,000 (or 100,000 jeffCoins). Each jeffCoin
has a value of 1,000 addies. Meow.

* Block 0 (Genesis Block)
  * Founders start with  100,000 jeffCoins

There are 5 transaction requests thereafter,

* Block 1
  * Founders sends Jeff 80 jeffCoins
* Block 2  
  * Jeff sends Matt 50 jeffCoins and a .5 jeffCoins fee to CoinVault
* Block 3  
  * Founders sends Matt 250 jeffCoins
  * Matt sends Jill 35 jeffCoins
* currentBlock
  * Matt send Jeff 15 Coins

The `currentBlock` is pending, meaning it needs to be verified by other nodes.
So this block will not have any value until part of the blockchain.

## TRANSACTIONS IN LEDGER

This illustrations shows the ledger in hte blockchain and currentBlock,

![IMAGE - bitcoin-unspent-transactions-ledger - IMAGE](../../docs/pics/bitcoin-unspent-transactions-ledger.jpg)

This illustration shows a visual look at how the transactions relate
(input/output) to each other in the blockchain,

![IMAGE - bitcoin-unspent-transactions-ledger-flow - IMAGE](../../docs/pics/bitcoin-unspent-transactions-ledger-flow.jpg)

## ADDING A TRANSACTION TO currentBlock

Here are the steps when adding a transaction request to the currentBlock,

* STEP 1 - MOCK - VERIFY SIGNATURE (So I don't need to deal with key pairs)
* STEP 2 - CHECK BALANCE TO SEE IF YOU HAVE THE MONEY
  * STEP 2.1 - GET UNSPENT OUTPUT TRANSACTIONS  - Make unspentOutputSlice
  * STEP 2.2 - GET BALANCE from unspentOutputSlice
* STEP 3 - CHECK IF YOU HAVE ENOUGH jeffCoins
* STEP 4 - PICK THE UNSPENT OUTPUTS TO USE AND PROVIDE CHANGE
* STEP 5 - LOAD currentBlock WITH TRANSACTION and MAKE CHANGE
  * STEP 5.1 - BUILD INPUT STRUCT FOR EACH UNSPENT OUTPUT
  * STEP 5.2 - BUILD OUTPUT STRUCT
  * STEP 5.3 - BUILD THE TRANSACTION
  * STEP 5.4 - PLACE transactionStruct IN transactionSlice
  * STEP 5.5 - LOAD THE CURRENT BLOCK WITH TRANSACTION

## RUN

This will load the transactions, display the blockchain and show the
balances for each address (pub Key),

```bash
go run control.go bitcoin-ledger.go data.go
```

The balances should be,

```txt
The balance for Founders PubKey (Address) is 99670000
The balance for Jeffs PubKey (Address) is 29500
The balance for Matts PubKey (Address) is 265000
The balance for CoinVaults PubKey (Address) is 500
```
