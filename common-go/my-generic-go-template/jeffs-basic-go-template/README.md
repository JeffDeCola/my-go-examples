# jeffs-basic-go-template

_A simple go template with flags, logging & error handling._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#overview)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#prerequisites)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#test)
* [USAGE](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#usage)
  * [-h](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#-h)
  * [-v](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#-v)
  * [-loglevel string](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/my-generic-go-template/jeffs-basic-go-template#-loglevel-string)
  
Documentation and references,

* Check out my
  [flag](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flag),
  [logrus](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/logging/logrus)
  and
  [errors](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/error-reporting/errors)
  examples
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

This is a very simple program that will ask a user for 2 integers
and adds them together.

## PREREQUISITES

Need,

```txt
github.com/sirupsen/logrus
```

Run,

```bash
go mod tidy
```

## RUN

```bash
go run jeffs-basic-go-template.go
```

Things to try,

* Add numbers together
* Enter a string instead of a number
* Use flags

## TEST

To create `_test` files,

```bash
gotests -w -all jeffs-basic-go-template.go
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

```bash
go run jeffs-basic-go-template.go {-h|-v} -loglevel {trace|info|error}
```

### -h

Help,

```bash
go run jeffs-basic-go-template.go -h
```

### -v

Version,

```bash
go run jeffs-basic-go-template.go -v
```

### -loglevel string

```bash
go run jeffs-basic-go-template.go -loglevel trace
go run jeffs-basic-go-template.go -loglevel info
go run jeffs-basic-go-template.go -loglevel error
```
