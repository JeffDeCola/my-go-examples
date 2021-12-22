# md5-hash-from-file example

_Getting an md5 hash (fingerprint) from an input file using the standard
`crypto/md5` package.
I also added a flag to read in your `.ssh/id_rsa.pub` key to get your ssh fingerprint.
Your github uses this for verification._

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

Take a string `plainText` and turn into a md5 hash,

```go
plainTextBytes := []byte(plainText)

// HASH
md5HashByte := md5.Sum(plainTextBytes)

// CONVERT TO STRING
md5Hash := hex.EncodeToString(md5HashByte[:])
```
