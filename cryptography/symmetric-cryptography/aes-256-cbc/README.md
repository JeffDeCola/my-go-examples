# aes-cbc example

_AES-256 **CBC** (**Cipher Block Chaining**) mode
where a block of plaintext is XORed with the previous cipherText block before
being encrypted._

I have the following AES-256 mode examples,

* [aes-256](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/symmetric-cryptography/aes-256)
  No Mode
* [aes-256-cbc](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/symmetric-cryptography/aes-256-cbc)
  Cipher Block Chaining **(You are here)**
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
run aes-cbc.go
```

You output should be,

```txt
Original Text:           This is AES-256 CBC (32 Bytes)!!

The 32-byte Key:         myverystrongpasswordo32bitlength
The Nonce:               dcde73a548a63e3d0073ad870fd21d7b

Encrypted Text:          8eddd047a775fe81ce5343edb4d5684c16465192ee85f3765bbfb35ddb219e50
Decrypted Text:          This is AES-256 CBC (32 Bytes)!!
```

## HOW IT WORKS

* plainText should be multiple of AES block size,
  else padding may be used
* The IV value should be equal to AES block size
* Encryption but not message integrity

For simplicity I did not include the nonce in the cipherText.

Encryption,

```go
// GET CIPHER BLOCK USING KEY
block, err := aes.NewCipher(keyByte)

// GET CBC ENCRYPTER
cbc := cipher.NewCBCEncrypter(block, nonce)

// ENCRYPT DATA
cbc.CryptBlocks(cipherTextByte, plaintextByte)

// RETURN HEX
cipherText := hex.EncodeToString(cipherTextByte)
```

Decryption,

```go
// GET CIPHER BLOCK USING KEY
block, err := aes.NewCipher(keyByte)
checkErr(err)

// GET CBC DECRYPTER
cbc := cipher.NewCBCDecrypter(block, nonce)

// DECRYPT DATA
cbc.CryptBlocks(plainTextByte, cipherTextByte)

// RETURN STRING
plainText := string(plainTextByte[:])
```

This illustration may help,

![IMAGE - aes-cbc - IMAGE](../../../docs/pics/aes-cbc.jpg)
