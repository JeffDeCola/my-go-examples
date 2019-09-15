# encryptfile tool

`encryptfile` _is a useful tool for
encrypting a file with 32 byte hash key using the `crypto/aes` package._

Use my other tool
[decryptfile](https://github.com/JeffDeCola/my-go-examples/tree/master/useful-tools-i-created/decryptfile)
to decrypt.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

```bash
go run encryptfile <INPUTFILE> <OUTPUTFILE>
```

## INSTALL

```bash
go install encryptfile.go
```

## HOW IT WORKS

The Advanced Encryption Standard, or AES, is a symmetric
block cipher chosen by the U.S. government to protect classified
information and is implemented in software and hardware throughout
the world to encrypt sensitive data.

We're going to use AES encryption from the standard go
[crypto/aes](https://golang.org/pkg/crypto/aes/)
package.

### STEP 1 - LETS CREATE A HASH KEY

First you need a 32 byte key.  Instead of typing a 32
character in, lets make it simple by turning a simple paraphrase into a key.
We will use the standard go
[crypto/md5](https://golang.org/pkg/crypto/md5/)
package.

```go
hasher := md5.New()
hasher.Write([]byte(paraphrase))
hash := hex.EncodeToString(hasher.Sum(nil))
```

### STEP 2 - ENCRYPT FILE WITH 32 BYTE HASH KEY

Noe we use that hash key with the data to encrypt the file,

```go
func encrypt(data []byte, hashKey string) []byte {

    // generate a new aes cipher using our 32 byte long key
    block, err := aes.NewCipher([]byte(hashKey))
    checkErr(err)
    gcm, err := cipher.NewGCM(block)
    checkErr(err)

    // Creates a new byte array the size of the nonce
    nonce := make([]byte, gcm.NonceSize())

    // Populates our nonce with a cryptographically secure random sequence
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err.Error())
    }

    // Encrypt our text using the Seal function
    cipherText := gcm.Seal(nonce, nonce, data, nil)
    return cipherText
}
```
