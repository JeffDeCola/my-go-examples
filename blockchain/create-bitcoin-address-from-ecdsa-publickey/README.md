# create-bitcoin-address-from-ecdsa-publickey example

_TBD._

Table of Contents,

* TBD

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## PREREQUISITES

```bash
go get -v -u github.com/btcsuite/btcutil/base58
go get -v -u golang.org/x/crypto/ripemd160
```

## OVERVIEW

This code is broken up as follows,

* **generateECDSAKeys()** Generate public and private keys
* **generateBitcoinAddress()**
  * hashPublicKey()
  * checksumKeyHash()
  * encodeKeyHash()

This illustration may help,

![IMAGE - create-bitcoin-address-from-ecdsa-publickey - IMAGE](https://github.com/JeffDeCola/my-go-examples/blob/master/docs/pics/create-bitcoin-address-from-ecdsa-publickey.jpg)

## RUN

```bash
go run create-bitcoin-address-from-ecdsa-publickey.go
```
