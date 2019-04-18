# amazon-aws example

`amazon-aws` _interacting with amazon's `aws` using `aws-sdk-go` client libraries._

Documentation and reference,

* My `aws` cheat sheet is
  [here](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/amazon-web-services-cheat-sheet)
* [AWS Go Documentation](https://aws.amazon.com/sdk-for-go/)

This [GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

The `aws` SDK for Go provides APIs and utilities that developers can use
to build Go applications that use `aws` services.

## GO GET

```bash
go get -u github.com/aws/aws-sdk-go/...
```

```go
import "github.com/aws/aws-sdk-go/service/s3"
```

## CREDENTIALS

Open [IAM console](https://console.aws.amazon.com/iam/home?#/home)
goto user, open the security credentials tab, and then
choose create access key.
