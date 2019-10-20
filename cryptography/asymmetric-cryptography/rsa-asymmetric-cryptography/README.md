# rsa-asymmetric-cryptography example

_Encrypt a message using a public key.
Decrypt the message using a private key.
Key pair generated using the `crypto/rsa` standard package
(RSA is a cryptosystem for public key encryption)._

Refer to the
[crypto/rsa](https://golang.org/pkg/crypto/rsa/)
package for more info.

I added a digital signature to this example
[rsa-asymmetric-cryptography-with-digital-signature](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/asymmetric-cryptography/rsa-asymmetric-cryptography-with-digital-signature).

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

RSA is a cryptosystem for public-key encryption, and is used for
securing sensitive data over unsecured networks.

This illustration may help,

![IMAGE - rsa-asymmetric-cryptography.jpg - IMAGE](/docs/pics/rsa-asymmetric-cryptography.jpg)

## RUN

```bash
go run rsa-asymmetric-cryptography.go <FILENAME>
go run rsa-asymmetric-cryptography.go test.txt
```

## HOW IT WORKS

Generate rsa keys,

```go
// GENERATE PRIVATE & PUBLIC KEY PAIR
privateKeyRaw, err := rsa.GenerateKey(rand.Reader, 2048)

// EXTRACT PUBLIC KEY
publicKeyRaw := &privateKeyRaw.PublicKey
```

Take a message `plainText` and encrypt it,

```go
// ENCRYPT
cipherTextByte, err := rsa.EncryptOAEP(
    hash,
    rand.Reader,
    publicKeyRaw,
    plainTextByte,
    label,
)
```

Take a message `cipherText` and decrypt it,

```go
// DECRYPT DATA
plainTextByte, err := rsa.DecryptOAEP(
    hash,
    rand.Reader,
    privateKeyRaw,
    cipherTextByte,
    label,
)
```
