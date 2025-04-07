# CREATE BITCOIN ADDRESS FROM ECDSA PUBLICKEY EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Create a bitcoin address from your ecdsa public key
using the
[crypto/ecdsa](https://pkg.go.dev/crypto/ecdsa)
standard package._

Table of Contents

* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/create-bitcoin-address-from-ecdsa-publickey#prerequisites)
* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/create-bitcoin-address-from-ecdsa-publickey#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/architectures/blockchain/create-bitcoin-address-from-ecdsa-publickey#run)

Documentation and Reference

* [blockchains](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/blockchain-cheat-sheet)

## PREREQUISITES

```bash
go get -v -u github.com/btcsuite/btcutil/base58
go get -v -u golang.org/x/crypto/ripemd160
```

## OVERVIEW

This code is broken up as follows,

* **generateECDSAKeys()**
  * encodeKeys()
  * decodeKeys()
* **generateBitcoinAddress()**
  * hashPublicKey()
  * checksumKeyHash()
  * encodeKeyHash()

This illustration may help,

![IMAGE - create-bitcoin-address-from-ecdsa-publickey - IMAGE](https://github.com/JeffDeCola/my-go-examples/blob/master/docs/pics/blockchain/create-bitcoin-address-from-ecdsa-publickey.jpg)

## RUN

```bash
go run create-bitcoin-address-from-ecdsa-publickey.go
```

A website I used to test my code
[here](http://gobittest.appspot.com/Address).
