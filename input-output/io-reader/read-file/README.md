# READ FILE EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Read a file (*os.File) to a buffer._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-file#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-file#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-file#test)

## OVERVIEW

Read a file input a small buffer and print.

## RUN

Run,

```bash
go run main.go
```

## TEST

To create `_test` files,

```bash
gotests -w -all main.go
```

To unit test the code,

```bash
go test -cover ./...
```
