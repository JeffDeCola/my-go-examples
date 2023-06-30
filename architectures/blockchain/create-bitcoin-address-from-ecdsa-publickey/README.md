# create-bitcoin-address-from-ecdsa-publickey example

_Create a bitcoin address from your ecdsa public key
using the `crypto/ecdsa` standard package._

Table of Contents,

* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/create-bitcoin-address-from-ecdsa-publickey#prerequisites)
* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/create-bitcoin-address-from-ecdsa-publickey#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/create-bitcoin-address-from-ecdsa-publickey#run)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

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
