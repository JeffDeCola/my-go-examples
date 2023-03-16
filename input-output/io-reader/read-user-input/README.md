# READ USER INPUT EXAMPLE

_Read user input (os.Stdin) to a buffer (using Read method)
and string (using Fscan)._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-user-input#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-user-input#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-user-input#test)

## OVERVIEW

Read user input until "stop".

## RUN

User one word. This is just an example, not for production.

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
