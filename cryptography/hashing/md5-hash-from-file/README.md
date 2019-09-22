# md5-hash-from-file example

`md5-hash-from-file` _is an example of
getting an md5 hash (fingerprint) from an input file using the standard
`crypto/md5` package.

Refer to the
[crypto/md5](https://golang.org/pkg/crypto/md5/)
package for more info.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

```bash
go run md5-hash-from-file.go <FILENAME>
go run md5-hash-from-file.go test.txt
```

If you run on test.txt you should get,

```txt
950dc9055bc2eb9b1f143e92d7bee6c4
```

## HOW IT WORKS

Take string `plainText` and turn into a md5 hash,

```go
plainTextBytes := []byte(plainText)

// HASH
md5HashByte := md5.Sum(plainTextBytes)

// CONVERT TO STRING
md5Hash := hex.EncodeToString(md5HashByte[:])
```
