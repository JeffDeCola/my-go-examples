# aes example

`aes` _is an example of
AES-256 **No Mode**
can only encrypt/decrypt 16 bytes of data._

I have the following AES-256 mode examples,

* [aes-256](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/symmetric-cryptography/aes-256)
  No Mode **(You are here)**
* [aes-256-cbc](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/symmetric-cryptography/aes-256-cbc)
  Cipher Block Chaining
* [aes-256-cfb](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/symmetric-cryptography/aes-256-cfb)
  Cipher FeedBack Mode
* [aes-256-ctr](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/symmetric-cryptography/aes-256-ctr)
  Counter Mode
* [aes-256-gcm](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/symmetric-cryptography/aes-256-gcm)
  Galois/Counter Mode **(I like this one)**
* [aes-256-ofb](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/symmetric-cryptography/aes-256-ofb)
  Output FeedBack Mode

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

```go
run aes.go
```

Your output should be,

```txt
Original Text:           Hello Jeff, only 16 Bytes of this will be encrypted.

The 32-byte Key:         myverystrongpasswordo32bitlength

Encrypted Text:          1d3288a55c1b74826395bd823d7cf0a2
Decrypted Text:          Hello Jeff, only
```

## HOW IT WORKS

This example is simple and very limiting and has `no mode`.
It will only encrypt 16 bytes of data.

Encryption,

```go
// ENCRYPT DATA
block.Encrypt(cipherTextByte, plaintextByte)
```

Decryption,

```go
// DeCRYPT DATA
block.Decrypt(plainTextByte, cipherTextByte)
```

This illustration may help,

![IMAGE - aes - IMAGE](../../../docs/pics/aes.jpg)
