# ecdsa-digital-signature example

_To verify who a message is from,
create and verify a digital signature using the
`crypto/ecdsa` standard package._

Refer to the
[crypto/ecdsa](https://golang.org/pkg/crypto/ecdsa/)
package for more info.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Digital signatures are based on elliptic curve cryptography (ECC) and
are important for message validation.

The Elliptic Curve Digital Signature Algorithm
is a great way to cryptographically sign a message.

This illustration may help,

![IMAGE - ecdsa-digital-signature.jpg - IMAGE](/docs/pics/ecdsa-digital-signature.jpg)

## RUN

```bash
go run ecdsa-digital-signature.go <FILENAME>
go run ecdsa-digital-signature.go test.txt
```

## HOW IT WORKS

Generate ecdsa keys,

```go
// GENERATE PRIVATE & PUBLIC KEY PAIR
curve := elliptic.P256()
privateKeyRaw, err := ecdsa.GenerateKey(curve, rand.Reader)

// EXTRACT PUBLIC KEY
publicKeyRaw := &privateKeyRaw.PublicKey
```

Create a digital signature,

```go
// CREATE SIGNATURE
r, s, err := ecdsa.Sign(rand.Reader, senderPrivateKeyRaw, signHash)

signatureByte := r.Bytes()
signatureByte = append(signatureByte, s.Bytes()...)
```

Verify the digital signature,

```go
// VERIFY SIGNATURE
verifyStatus := ecdsa.Verify(senderPublicKeyRaw, signhash, r, s)
```
