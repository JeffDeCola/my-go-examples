# gotests-complex-function example

`gotests-complex-function` _is an example of
testing a function with complex inputs and outputs._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## WHAT IS GOTESTS

`gotests` makes writing go tests easy. It's a golang cli that generates
table driven tests based on its target source files' function
and method signatures.

Install,

```bash
go get -u github.com/cweill/gotests/...
```

## THE FUNCTION TO TEST

The function `Check()` has the following input/outputs.

`func Check(input InputJSON, logger *log.Logger) (checkOutputJSON, error) {`

where,

```go
version struct {
    Ref string `json:"ref"`
}

InputJSON struct {
    Params  map[string]string `json:"params"`
    Source  map[string]string `json:"source"`
    Version version           `json:"version"`
}

checkOutputJSON []version
```

## HOW I GENERATED THE TEST FILE

```bash
gotests -w -all complex-function.go
```

Refer to this file to see how I added tests.

## RUN TEST & COVERAGE

```bash
go test -v -cover .
go test -coverprofile coverage.out
go tool cover -html=coverage.out -o coverage.html
```
