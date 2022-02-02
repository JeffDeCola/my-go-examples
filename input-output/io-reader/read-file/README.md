# read-file

_Read a file (*os.File) to a buffer._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-file#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-file#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-file#test)

Documentation and references,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Read a file input a small buffer and print.

## RUN

```bash
go run read-file.go
```

## TEST

To create `_test` files,

```bash
gotests -w -all read-file.go
```

To unit test the code,

```bash
go test -cover ./...
```
