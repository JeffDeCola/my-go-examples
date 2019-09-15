# decryptfile tool

`decryptfile` _is a useful tool for
decrypting a file with 32 byte hash key using the crypto package._

Use my other tool
[encryptfile](https://github.com/JeffDeCola/my-go-examples/tree/master/useful-tools-i-created/encryptfile)
to encrypt.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

```bash
go run decryptfile <INPUTFILE> <OUTPUTFILE>
```

## INSTALL

```bash
go install decryptfile.go
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

### STEP 2 - DECRYPT FILE WITH 32 BYTE HASH KEY

Noe we use that hash key with the data to decrypt the file,

```go
func decrypt(data []byte, hashKey string) []byte {

    key := []byte(hashKey)
    block, err := aes.NewCipher(key)
    checkErr(err)

    gcm, err := cipher.NewGCM(block)
    checkErr(err)

    nonceSize := gcm.NonceSize()
    nonce, cipherText := data[:nonceSize], data[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
    checkErr(err)

    return plaintext
}
```
