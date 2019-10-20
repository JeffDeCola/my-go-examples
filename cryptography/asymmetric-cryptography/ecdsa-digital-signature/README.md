# ecdsa-digital-signature example

_tbd._

Refer to the
[crypto/ecdsa](https://golang.org/pkg/crypto/ecdsa/)
package for more info.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Digital signatures are based on elliptic curve cryptography (ECC) and
are important for message validation.

The Elliptic Curve Digital Signature Algorithm
is a great way at present to cryptographically sign a message

This illustration may help,

![IMAGE - ecdsa-digital-signature.jpg - IMAGE](/docs/pics/ecdsa-digital-signature.jpg)

## RUN

```bash
go run ecdsa-digital-signature.go <FILENAME>
go run ecdsa-digital-signature.go test.txt
```

If you run on test.txt you should get,

```txt
????
```

## HOW IT WORKS

Take a string `plainText` and turn into a md5 hash,

```go
plainTextBytes := []byte(plainText)

// HASH
md5HashByte := md5.Sum(plainTextBytes)

// CONVERT TO STRING
md5Hash := hex.EncodeToString(md5HashByte[:])
```
