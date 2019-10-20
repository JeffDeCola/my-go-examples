# ecdsa-asymmetric-cryptography example

_tbd._

Refer to the
[crypto/ecdsa](https://golang.org/pkg/crypto/ecdsa/)
package for more info.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Elliptic curve cryptography (ECC) 

This illustration may help,

![IMAGE - ecdsa-asymmetric-cryptography.jpg - IMAGE](/docs/pics/ecdsa-asymmetric-cryptography.jpg)

## RUN

```bash
go run ecdsa-asymmetric-cryptography.go <FILENAME>
go run ecdsa-asymmetric-cryptography.go test.txt
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
