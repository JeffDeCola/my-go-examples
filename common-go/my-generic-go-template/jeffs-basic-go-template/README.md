# JEFF BASIC GO TEMPLATE

_A simple go template with flags, logging & error handling._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#overview)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#prerequisites)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#test)
* [USAGE](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#usage)
  * [-h](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#-h)
  * [-v](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#-v)
  * [-loglevel string](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#-loglevel-string)

Documentation and References

* Check out my
  [flags](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flags),
  [jeffs-logger](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/jeffs-logger)
  and
  [error-example](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/error-example)
  examples

## OVERVIEW

This is a very simple program that will ask a user for 2 integers
and adds them together.

## PREREQUISITES

You will need the following go package,

```txt
go install github.com/cweill/gotests/...@latest
go get -u github.com/sirupsen/logrus
```

## RUN

```bash
go run main.go
```

Things to try,

* Add numbers together
* Enter a string instead of a number
* Use flags

## TEST

To create `_test` files,

```bash
gotests -w -all main.go
```

To unit test the code,

```bash
go test -cover ./...
```

Note, to test the function `getUserInput()` I am using an io.Reader to provide
a way to mimic user input and test.

```go
func getUserInput(r io.Reader, askUser string) (string, error) {
...
    _, err := fmt.Fscan(r, &nString)
...
```

I do the same thing in the function `getNumbers()`.

## USAGE

```text
go run main.go {-h|-v} -loglevel {trace|info|error}
```

### -h

Help,

```bash
go run main.go -h
```

### -v

Version,

```bash
go run main.go -v
```

### -loglevel string

```bash
go run main.go -loglevel trace
go run main.go -loglevel info
go run main.go -loglevel error
```
