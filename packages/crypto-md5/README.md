# crypto/md5 example

`crypto/md5` _package. An example of
getting an md5 fingerprint from an ssh key._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## RUN

Unless you use the -i flag, it
will assume you want to use `~/home/jeff/.ssh/id_rsa.pub`,

```bash
go run crypto-md5.go
go run crypto-md5.go -i ~/.ssh/id_ras.pub
```
