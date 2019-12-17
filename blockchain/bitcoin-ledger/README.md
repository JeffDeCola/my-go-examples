# bitcoin-ledger example

_Demonstrates the bitcoin ledger in a blockchain using the
**unspent transaction output model**._

Table of Contents,

* tbd

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## PREREQUISITES

```bash
tbd
```

## OVERVIEW

An **unspent transaction output model** is a method to log the transactions
of coins in a cryptocurrency.  It uses an input/output model and has
the following rules,

* Every input must have an output.
* tbd

## THIS EXAMPLE

The first block in the blockchain will contain 1 transaction.  That will be the
founders initial value of 100,000,000 (or 100,000 jeffCoins). Each jeffCoin
has a value of 1,000 jefftoshi. Yup.

There will be 4 transactions thereafter,

* Founders sends Jeff 80 jeffCoins
* Jeff sends Matt 50 jeffCoins and .5 jeffCoins fee to CoinVault
* Founders sends Matt 250 jeffCoins
* Matt sends Jill 275 jeffCoins

## TRANSACTIONS IN LEDGER

This illustrations shows the ledger starting point and the transaction
we're going to add,

![IMAGE - bitcoin-unspent-transactions-ledger - IMAGE](../../docs/pics/bitcoin-unspent-transactions-ledger.jpg)

This illustration shows a visual look at how the transaction relate to each other,
![IMAGE - bitcoin-unspent-transactions-ledger-flow - IMAGE](../../docs/pics/bitcoin-unspent-transactions-ledger-flow.jpg)

## ADDING A TRANSACTION

tbd

## GETTING A BALANCE

tbd

## RUN

```bash
go run bitcoin-ledger.go data.go
```
