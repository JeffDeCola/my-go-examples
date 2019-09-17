# aes-gcm example

`aes-gcm` _is an example of
_AES-256 **GCM** (**Galois/Counter Mode**) mode
is a block cipher counter mode with authentication._

I have the following AES mode examples,

* [aes](https://github.com/JeffDeCola/my-go-examples/tree/master/encryption-decryption/aes)
  No Mode
* [aes-gcm](https://github.com/JeffDeCola/my-go-examples/tree/master/encryption-decryption/aes-gcm)
  Galois/Counter Mode **(You are here)**
* [aes-cbc](https://github.com/JeffDeCola/my-go-examples/tree/master/encryption-decryption/aes-cbc)
  Cipher Block Chaining
* [aes-cfb](https://github.com/JeffDeCola/my-go-examples/tree/master/encryption-decryption/aes-cfb)
  Cipher Feedback Mode
* [aes-ctr](https://github.com/JeffDeCola/my-go-examples/tree/master/encryption-decryption/aes-ctr)
  Counter Mode
* [aes-ofb](https://github.com/JeffDeCola/my-go-examples/tree/master/encryption-decryption/aes-ofb)
  Output Feedback Mode

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

```go
run aes-gcm.go
```

Your output should be similar,

```txt
Original Text:           This is an example of AES-256 GCM.

The 32-byte Key:         myverystrongpasswordo32bitlength
Additional Data:         Jeff's additional data for authorization
The Nonce:               f796b0e6a7fabdf9cb6e6d4e

Encrypted Text:          f796b0e6a7fabdf9cb6e6d4ed4109c9538107d3b9a28ecefde2b69608798209ca9ae3932e7ddc1c1884d2bb7d3aed3d85ed56a9f60389503b9b7c08e8121
Decrypted Text:          This is an example of AES-256 GCM.
```

Note how you an see the nonce at the beginning of the encrypted text.
Its because I put it there so I can extract it during decrypt.

## HOW IT WORKS

A Counter mode effectively turns a block cipher into a stream cipher,
and therefore many of the rules for stream ciphers still apply.
GCM uses an IV (Initialization Vector) or Nonce.
I also used some Additional Data for authentication.

Encryption using your gcm block,

```go
// GET CIPHER BLOCK USING KEY
block, err := aes.NewCipher(keyByte)

// GET GCM INSTANCE THAT USES THE AES CIPHER
gcm, err := cipher.NewGCM(block)

// ENCRYPT DATA
cipherTextByte := gcm.Seal(nonce, nonce, plaintextByte, additionalDataByte)
```

Decryption using your gcm block,

```go
// GET CIPHER BLOCK USING KEY
block, err := aes.NewCipher(keyByte)

// GET GCM BLOCK
gcm, err := cipher.NewGCM(block)

// DECRYPT DATA
plainTextByte, err := gcm.Open(nil, nonce, cipherTextByte, additionalDataByte)
```

This illustration may help,

![IMAGE - aes-gcm - IMAGE](../../docs/pics/aes-gcm.jpg)
