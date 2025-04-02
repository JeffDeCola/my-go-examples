# ERROR EXAMPLE

_Error Handling using the standard
[errors](https://pkg.go.dev/errors)
library._

tl;dr

```go
var ErrFilenameEmpty = errors.New("filename can not be empty")

return fmt.Errorf("firstLevel: %w", err)
return fmt.Errorf("secondLevel: %w", err)
return fmt.Errorf("thirdLevel: %w", ErrFilenameEmpty)
return fmt.Errorf("thirdLevel: %w, %s", ErrFilenameEmpty, otherInfo)
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example#overview)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example#prerequisites)
* [DEFINE YOUR OWN ERROR](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example#define-your-own-error)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example#test)

Documentation and Reference

* Refer to the
  [errors](https://pkg.go.dev/errors)
  package for more info

## OVERVIEW

Go doesn’t have exceptions, so it doesn’t have try, catch or anything similar.
So how do we handle errors? **With Multiple return values**.

## PREREQUISITES

You will need the following go packages,

```bash
go install github.com/cweill/gotests/...@latest
```

## DEFINE YOUR OWN ERROR

It's very easy. For example,

```go
var ErrFilenameEmpty = errors.New("filename can not be empty")

if filename == "" {
    return fmt.Errorf("thirdLevel: %w", ErrFilenameEmpty)
}
```

## RUN

The programs **propagates** the error up the function calls.

Run,

```bash
go run main.go
```

The output may look like,

```bash
[ERROR] Got an Error (error: firstLevel: secondLevel: thirdLevel: filename can not be empty)
Optional
[ERROR] Caused by: (chain: firstLevel: secondLevel: thirdLevel: filename can not be empty)
[ERROR] Caused by: (chain: secondLevel: thirdLevel: filename can not be empty)
[ERROR] Caused by: (chain: thirdLevel: filename can not be empty)
[ERROR] Caused by: (chain: filename can not be empty)
[ERROR] Root Cause: (root: filename can not be empty)
```

Notice how it traces the error back to your original function.

## TEST

To create _test files,

```bash
gotests -w -all main.go
```

To unit test the code,

```bash
go test -cover ./...
```
