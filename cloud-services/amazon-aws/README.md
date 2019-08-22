# amazon-aws example

```txt
*** UNDER CONSTRUCTION - CHECK BACK SOON ***
```

`amazon-aws` _is an example of
interacting with amazon's
[aws](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/amazon-web-services-cheat-sheet)
using
`aws-sdk-go`
client libraries._

Documentation and reference,

* My
  [aws cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/amazon-web-services-cheat-sheet)
* [AWS Go Documentation](https://aws.amazon.com/sdk-for-go/)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

The `aws` SDK for Go provides APIs and utilities that developers can use
to build Go applications that use `aws` services.

## GO GET

get,

```bash
go get -u github.com/aws/aws-sdk-go/...
```

And import them in your go code,

```go
import "github.com/aws/aws-sdk-go/service/s3"
```

## CREDENTIALS

Open [IAM console](https://console.aws.amazon.com/iam/home?#/home)
goto user, open the security credentials tab, and then
choose create access key.
