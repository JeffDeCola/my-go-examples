# rsa-asymmetric-cryptography-with-digital-signature example

_Lets verify who the message is from.
Create and verify a digital signature using the
`crypto/rsa` standard package._

Refer to the
[crypto/rsa](https://golang.org/pkg/crypto/rsa/)
package for more info.

I added a digital signature to this example
[rsa-asymmetric-cryptography](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/asymmetric-cryptography/rsa-asymmetric-cryptography).

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

RSA is a cryptosystem for public-key encryption, and is used for
securing sensitive data over unsecured networks. It is also used to add
digital signatures for verification.

This illustration may help,

![IMAGE - rsa-asymmetric-cryptography-with-digital-signature.jpg - IMAGE](/docs/pics/rsa-asymmetric-cryptography-with-digital-signature.jpg)

## RUN

```bash
go run rsa-asymmetric-cryptography-with-digital-signature.go <FILENAME>
go run rsa-asymmetric-cryptography-with-digital-signature.go test.txt
```

## HOW IT WORKS

Generate rsa keys,

```go
// GENERATE PRIVATE & PUBLIC KEY PAIR
privateKeyRaw, err := rsa.GenerateKey(rand.Reader, 2048)

// EXTRACT PUBLIC KEY
publicKeyRaw := &privateKeyRaw.PublicKey
```

Create a digital signature,

```go
// CREATE SIGNATURE
signatureByte, err := rsa.SignPSS(
    rand.Reader,
    senderPrivateKeyRaw,
    newhash,
    hashed,
    &opts,
)
```

Verify the digital signature,

```go
// VERIFY SIGNATURE
err := rsa.VerifyPSS(
    senderPublicKeyRaw,
    newhash,
    hashed,
    signatureByte,
    &opts,
)
```
