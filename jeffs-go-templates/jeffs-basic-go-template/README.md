# jeffs-basic-go-template example

_A simple go template with flags, logging & error handling._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/jeffs-go-templates/jeffs-basic-go-template#overview)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/jeffs-go-templates/jeffs-basic-go-template#prerequisites)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/jeffs-go-templates/jeffs-basic-go-template#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/jeffs-go-templates/jeffs-basic-go-template#test)
* [USAGE](https://github.com/JeffDeCola/my-go-examples/tree/master/jeffs-go-templates/jeffs-basic-go-template#usage)
  * [-h](https://github.com/JeffDeCola/my-go-examples/tree/master/jeffs-go-templates/jeffs-basic-go-template#-h)
  * [-v](https://github.com/JeffDeCola/my-go-examples/tree/master/jeffs-go-templates/jeffs-basic-go-template#-v)
  * [-loglevel string](https://github.com/JeffDeCola/my-go-examples/tree/master/jeffs-go-templates/jeffs-basic-go-template#-loglevel-string)

Documentation and references,

* Check out my [flag](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/flag),
  [logrus](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus)
  and
  [errors](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/errors)
  examples
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

This is a very simple program that will ask a user for 2 integers
and add them together.

## PREREQUISITES

```bash
go get -u -v github.com/sirupsen/logrus
go get -u -v github.com/golang/mock/gomock
go get -u -v github.com/golang/mock/mockgen
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

Since there are external dependencies like user input, we need to mock that.

```bash
mockgen -source=jeffs-basic-go-template.go -destination=jeffs-basic-go-template_mock.go -package=main
```

To unit test the code,

```bash
go test -cover ./... | tee test_coverage.txt
cat test_coverage.txt
```

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
