# sha256-hash-from-file example

_Get an sha256 hash (fingerprint) from an input file using the standard
`crypto/sha256` package._

Refer to the
[crypto/sha256](https://golang.org/pkg/crypto/sha256/)
package for more info.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

```bash
go run sha256-hash-from-file.go <FILENAME>
go run sha256-hash-from-file.go test.txt
```

If you run on test.txt you should get,

```txt
7df64288ab2d0bcd15a9c78985965326d4b9d61202d1bc7b55b4abc89698c599
```

## HOW IT WORKS

Take a string `plainText` and turn into a sha256 hash,

```go
plainTextBytes := []byte(plainText)

// HASH
sha256HashByte := sha256.Sum256(plainTextBytes)

// CONVERT TO STRING
sha256Hash := hex.EncodeToString(sha256HashByte[:])
```
