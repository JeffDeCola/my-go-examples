# crypto/md5 (standard) example

`crypto/md5` _standard package. An example of
getting an md5 fingerprint from an ssh key._

Refer to the
[crypto/md5](https://golang.org/pkg/crypto/md5/)
package for more info.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

Unless you use the -i flag, it
will assume you want to use `~/home/jeff/.ssh/id_rsa.pub`,

```bash
go run crypto-md5.go
go run crypto-md5.go -i ~/.ssh/id_ras.pub
```
